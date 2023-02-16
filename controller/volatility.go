package controller

import (
	"time"
	"math"
	"errors"
	"github.com/shopspring/decimal"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// VolatilityManager - A struct to store volatility surfaces
type VolatilityManager struct {
	// A single surface will be for a single underlying
	Underlying shared.Underlying
	// A single surface will be for a single expiry
	Expiry uint64
	// A surface is stored as N moneyness points 
	// And N points for time until expiry (or tau)
	Surface [20][20]float64
	// Moneyness points (sorted)
	MoneynessPoints [20]float64
	// Minimum moneyness 
	MinMoneyness float64
	// Maximum moneyness 
	MaxMoneyness float64
	// Time to expiry points (sorted & annualized)
	TauPoints [20]float64
	// Minimum time to expiry (annualized)
	MinTau float64
	// Max time to expiry (annualized)
	MaxTau float64
}

// CreateVolatilityManager - Create `VolatilityManager` instance
// Arguments:
// 	underlying - Address for the underlying asset
// 	expiry - Expiry timestamp since epoch; unique identifiser
// 	initSigma - Value to initialize the surface to (e.g. 0)
// 	minMoneyness - Minimum moneyness in surface
// 	maxMoneyness - Maximum moneyness in surface
func CreateVolatilityManager(
	underlying shared.Underlying,
	expiry uint64,
	initSigma float64,
	minMoneyness float64,
	maxMoneyness float64,
) (*VolatilityManager, error) {
	now := uint64(time.Now().Unix())
	if expiry <= now {
		return nil, errors.New("CreateVolatilityManager: past expiry")
	}
	expiryYears := float64(expiry - now) / 31556952.0
	var manager VolatilityManager = VolatilityManager{
		Underlying: underlying,
		Expiry: expiry,
		MinMoneyness: minMoneyness,
		MaxMoneyness: maxMoneyness,
		MinTau: 0.0,
		MaxTau: expiryYears,
	}
	n := len(manager.Surface)
	m := len(manager.Surface[0])
	for i := 0; i < n; i++ {
		// if i = 0, t = 0; if i = (n-1), t = 1
		t := float64(i) / float64(n-1)
		manager.MoneynessPoints[i] = minMoneyness + (maxMoneyness - minMoneyness) * t
		for j := 0; j < m; j++ {
			manager.Surface[i][j] = initSigma
			if i == 0 {
				// if j = 0, s = expiryYears; if j = (m-1), t = 0
				s := (float64(j) / float64(m-1))
				manager.TauPoints[j] =  expiryYears * s
			}
		}
	}
	return &manager, nil
}

// Tau - Return the time to expiry
// If < 0, returns 0
func (manager *VolatilityManager) Tau() uint64 {
	ts := uint64(time.Now().Unix())  // make unsigned
	var tau uint64
	if manager.Expiry > ts {
		tau = manager.Expiry - ts
	} else {
		tau = 0
	}
	return tau
}

// TauAnnualized - Return time to expiry in years
// If < 0, returns 0
func (manager *VolatilityManager) TauAnnualized() float64 {
	tau := manager.Tau()
	return float64(tau) / 31556952.0
}

// Query - Given a moneyness, and time until expiry, query for IV
// Arguments:
// 	spot - Current spot price
//  strike - Current strike price
// 	tauAnnualized - Current time to expiry in years
// 	useBackstop - If true, use Deribit's API as a backstop. This will update our own surface
// Returns:
// 	sigma - Implied volatility (annualized) as used in Black-Scholes
// Notes:
// 	The IV is computed as a weighted average over the four closest points
func (manager *VolatilityManager) Query(
	spot decimal.Decimal,
	strike decimal.Decimal,
	tauAnnualized float64,
) float64 {
	moneyness := (spot.Div(strike)).InexactFloat64()
	sigma := manager.Interpolate(moneyness, tauAnnualized)
	backstop, err := manager.Backstop(strike)
	if err == nil {
		diff := math.Abs(sigma - backstop) / sigma
		// If the difference is higher than 2%, default to deribit
		if diff > 0.02 {
			sigma = backstop
			// Move our surface to be closer to vol
			manager.Update(spot, strike, backstop, 0.8)
		}
	}
	return sigma
}

// Backstop - Get deribits implied volatility so we don't end up serving something bad
// Arguments:
// 	strike - Current strike price
// 	isCall - Whether option is a call or put. This doesn't matter that much
// Returns:
// 	sigma - Implied volatility (annualized) as used in Black-Scholes
func (manager *VolatilityManager) Backstop(strike decimal.Decimal) (float64, error) {
	// It doesnt' actually matter is call/put, the IV for mark price stays the same
	name := ExpiryToInstrumentName(manager.Underlying, manager.Expiry, strike, true)
	_, _, sigma, err := GetDeribitSigma(name)
	if err != nil {
		return 0, err
	}
	return sigma, nil
}

// Update - Given a new order, update the surface 
// Notes:
// 	This will perform a local smoothing operation rather than just update a single point
// Arguments:
// 	underlying - Enum for underlying
// 	spot - Spot price in decimals
// 	strike - Strike price in decimals
// 	sigma - Implied volatility (annualized) as used in Black-Scholes
// 	weight - Weight to give new value. DEfault choice is 0.8
func (manager *VolatilityManager) Update(
	spot decimal.Decimal,
	strike decimal.Decimal,
	sigma float64,
	weight float64,
) (error) {
	moneyness := (spot.Div(strike)).InexactFloat64()

	// Find the closest i and j indices
	i := manager.FindClosestMoneyness(moneyness)

	// We compute time-to-expiry using current time
	j := manager.FindClosestTau(manager.TauAnnualized())

	// Save to surface as a weighted rolling mean
	manager.Surface[i][j] = 0.8 * sigma + 0.2 * manager.Surface[i][j]

	return nil
}

// Interpolate - Find the weighted average between the four nearest points
// Resources:
// 	https://en.wikipedia.org/wiki/Bilinear_interpolation
func (manager *VolatilityManager) Interpolate(
	moneyness float64,
	tau float64,
) (float64) {
	i1, i2 := manager.FindClosestTwoMoneyness(moneyness)
	j1, j2 := manager.FindClosestTwoTau(tau)

	// Rename variables
	x := moneyness
	y := float64(tau)

	// Get x values at indices
	x1 := manager.MoneynessPoints[i1]
	x2 := manager.MoneynessPoints[i2]

	// Get y values at indices
	y1 := float64(manager.TauPoints[j1])
	y2 := float64(manager.TauPoints[j2])

	// Get values at the four points
	f11 := manager.Surface[i1][j1]
	f12 := manager.Surface[i1][j2]
	f21 := manager.Surface[i2][j1]
	f22 := manager.Surface[i2][j2]

	// Bilinear interpolation
	var (
		f1 float64
		f2 float64
		res float64
	)
	if x1 == x2 {
		f1 = (f11 + f21) / 2.0
		f2 = (f12 + f22) / 2.0
	} else {
		f1 = (x2 - x) / (x2 - x1) * f11 + (x - x1) / (x2 - x1) * f21
		f2 = (x2 - x) / (x2 - x1) * f12 + (x - x1) / (x2 - x1) * f22
	}
	if y1 == y2 {
		res = (f1 + f2) / 2.0
	} else {
		res = (y2 - y) / (y2 - y1) * f1 + (y - y1) / (y2 - y1) * f2
	}
	return res
}

// FindClosestTwoMoneyness - Find the index of the two closest moneyness
// Notes:
// 	If query is exactly one of the moneyness points, then return the same index twice
//	Otherwise return two indices of the closest neighbors
func (manager *VolatilityManager) FindClosestTwoMoneyness(query float64) (int, int) {
	if query < manager.MoneynessPoints[0] {
		return 0, 0
	}
	n := len(manager.Surface)
	if query > manager.MoneynessPoints[n-1] {
		return n-1, n-1
	}

	var (
		indexA int
		indexB int
	)

	for i := 0; i < n; i++ {
		if query == manager.MoneynessPoints[i] {
			return i, i
		} else if query < manager.MoneynessPoints[i] {
			// Once we find the first entry larger than query, quit
			indexB = i
			return indexA, indexB
		} else if query > manager.MoneynessPoints[i] {
			indexA = i
		}
	}
	return indexA, indexB
}

// FindClosestMoneyness - Find the closest index to a query value for moneyness
func (manager *VolatilityManager) FindClosestMoneyness(query float64) (index int) {
	i1, i2 := manager.FindClosestTwoMoneyness(query)
	m1 := manager.MoneynessPoints[i1]
	m2 := manager.MoneynessPoints[i2]

	if math.Abs(m1 - query) < math.Abs(m2 - query) {
		index = i1
	} else {
		index = i2
	}
	return
}

// FindClosestTwoTau - Find the index of the (two) closest times to expiries
// Notes:
// 	If query is exactly one of the tau points, then return the same index twice
//	Otherwise return two indices of the closest neighbors
//	All tau terms below are annualilzed
func (manager *VolatilityManager) FindClosestTwoTau(query float64) (int, int) {
	if query < manager.TauPoints[0] {
		return 0, 0
	}
	m := len(manager.Surface[0])
	if query > manager.TauPoints[m-1] {
		return m-1, m-1
	}

	var (
		indexA int
		indexB int
	)

	for j := 0; j < m; j++ {
		if query == manager.TauPoints[j] {
			return j, j
		} else if query < manager.TauPoints[j] {
			// Once we find the first entry larger than query, quit
			indexB = j
			return indexA, indexB
		} else if query > manager.TauPoints[j] {
			indexA = j
		}
	}
	return indexA, indexB
}

// FindClosestTau - Find the closest index to a query value for time-to-expiry
func (manager *VolatilityManager) FindClosestTau(query float64) (index int) {
	j1, j2 := manager.FindClosestTwoTau(query)
	t1 := manager.TauPoints[j1]
	t2 := manager.TauPoints[j2]

	if math.Abs(float64(t1 - query)) < math.Abs(float64(t2 - query)) {
		index = j1
	} else {
		index = j2
	}
	return
}

// IsEqual - Check if two volatility manager are equal
func (manager *VolatilityManager) IsEqual(other *VolatilityManager) bool {
	if manager.Underlying != other.Underlying {
		return false
	}
	if manager.Expiry != other.Expiry {
		return false
	}
	if manager.Surface != other.Surface {
		return false
	}
	if manager.MoneynessPoints != other.MoneynessPoints {
		return false
	}
	if manager.MinMoneyness != other.MinMoneyness {
		return false
	}
	if manager.MaxMoneyness != other.MaxMoneyness {
		return false
	}
	if manager.TauPoints != other.TauPoints {
		return false
	}
	if manager.MinTau != other.MinTau {
		return false
	}
	if manager.MaxTau != other.MaxTau {
		return false
	}
	return true
}
