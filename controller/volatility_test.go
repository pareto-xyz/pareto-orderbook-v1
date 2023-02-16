package controller

import (
	"testing"
	"time"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreateVolatilityManager(t *testing.T) {
	manager, err := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	require.Equal(t, err, nil)
	require.Equal(t, manager != nil, true)
}

func TestCorrectDims(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
  require.Equal(t, manager != nil, true)
	require.Equal(t, len(manager.Surface), 20)
	require.Equal(t, len(manager.Surface[0]), 20)
}

func TestCorrectInitMoneynessPoints(t * testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	require.Equal(t, manager.MoneynessPoints[0], 0.001)
	require.InDelta(t, manager.MoneynessPoints[1], .10621052631578948, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[2], .21142105263157895, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[3], .31663157894736843, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[4], .4218421052631579, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[5], .5270526315789473, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[6], .6322631578947369, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[7], .7374736842105263, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[8], .8426842105263158, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[9], .9478947368421052, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[10], 1.0531052631578945, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[11], 1.1583157894736842, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[12], 1.2635263157894736, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[13], 1.3687368421052633, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[14], 1.4739473684210525, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[15], 1.579157894736842, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[16], 1.6843684210526315, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[17], 1.7895789473684212, 0.0001)
	require.InDelta(t, manager.MoneynessPoints[18], 1.8947894736842104, 0.0001)
	require.Equal(t, manager.MoneynessPoints[19], 2.0)
}

func TestCorrectInitTauPoints(t * testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	require.InDelta(t, manager.TauPoints[19], 0.019165349048919554, 0.0001)
	require.InDelta(t, manager.TauPoints[18], 0.018156646467397473, 0.0001)
	require.InDelta(t, manager.TauPoints[17], 0.01714794388587539, 0.0001)
	require.InDelta(t, manager.TauPoints[16], 0.01613924130435331, 0.0001)
	require.InDelta(t, manager.TauPoints[15], 0.015130538722831228, 0.0001)
	require.InDelta(t, manager.TauPoints[14], 0.014121836141309147, 0.0001)
	require.InDelta(t, manager.TauPoints[13], 0.013113133559787063, 0.0001)
	require.InDelta(t, manager.TauPoints[12], 0.01210443097826498, 0.0001)
	require.InDelta(t, manager.TauPoints[11], 0.0110957283967429, 0.0001)
	require.InDelta(t, manager.TauPoints[10], 0.01008702581522082, 0.0001)
	require.InDelta(t, manager.TauPoints[9], 0.009078323233698737, 0.0001)
	require.InDelta(t, manager.TauPoints[8], 0.008069620652176654, 0.0001)
	require.InDelta(t, manager.TauPoints[7], 0.007060918070654574, 0.0001)
	require.InDelta(t, manager.TauPoints[6], 0.00605221548913249, 0.0001)
	require.InDelta(t, manager.TauPoints[5], 0.00504351290761041, 0.0001)
	require.InDelta(t, manager.TauPoints[4], 0.004034810326088327, 0.0001)
	require.InDelta(t, manager.TauPoints[3], 0.0030261077445662464, 0.0001)
	require.InDelta(t, manager.TauPoints[2], 0.0020174051630441636, 0.0001)
	require.InDelta(t, manager.TauPoints[1], 0.0010087025815220829, 0.0001)
	require.Equal(t, manager.TauPoints[0], 0.0)
}

func TestGetTau(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	require.Equal(t, manager.Tau(), uint64(604800))
}

func TestGetTauAnnualized(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	require.InDelta(t, manager.TauAnnualized(), 0.01916534904, 0.0001)
}

func TestGetTauAnnualizedOneYear(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 31556952,
		0.5,
		0.001,
		2,
	)
	require.InDelta(t, manager.TauAnnualized(), 1, 0.0001)
}

func TestFindClosestTwoMoneyness1(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(1)
	require.Equal(t, i1, 9)
	require.Equal(t, i2, 10)
}

func TestFindClosestTwoMoneyness2(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(1.8)
	require.Equal(t, i1, 17)
	require.Equal(t, i2, 18)
}

func TestFindClosestTwoMoneyness3(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(0.3)
	require.Equal(t, i1, 2)
	require.Equal(t, i2, 3)
}

func TestFindClosestTwoMoneyness4(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(0.001)
	require.Equal(t, i1, 0)
	require.Equal(t, i2, 0)
}

func TestFindClosestTwoMoneyness5(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(0.0005)
	require.Equal(t, i1, 0)
	require.Equal(t, i2, 0)
}

func TestFindClosestTwoMoneyness6(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(2.0)
	require.Equal(t, i1, 19)
	require.Equal(t, i2, 19)
}

func TestFindClosestTwoMoneyness7(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoMoneyness(5.0)
	require.Equal(t, i1, 19)
	require.Equal(t, i2, 19)
}

func TestFindClosestMoneyness1(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(1.05)
	require.Equal(t, i, 10)
}

func TestFindClosestMoneyness2(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(1.8)
	require.Equal(t, i, 17)
}

func TestFindClosestMoneyness3(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(0.3)
	require.Equal(t, i, 3)
}

func TestFindClosestMoneyness4(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(0.001)
	require.Equal(t, i, 0)
}

func TestFindClosestMoneyness5(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(0.0005)
	require.Equal(t, i, 0)
}

func TestFindClosestMoneyness6(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(2.0)
	require.Equal(t, i, 19)
}

func TestFindClosestMoneyness7(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestMoneyness(5.0)
	require.Equal(t, i, 19)
}

func TestFindClosestTwoTau1(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoTau(0.01)
	require.Equal(t, i1, 9)
	require.Equal(t, i2, 10)
}

func TestFindClosestTwoTau2(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoTau(0.0035)
	require.Equal(t, i1, 3)
	require.Equal(t, i2, 4)
}

func TestFindClosestTwoTau3(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoTau(0.016)
	require.Equal(t, i1, 15)
	require.Equal(t, i2, 16)
}

func TestFindClosestTwoTau4(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoTau(0.2)
	require.Equal(t, i1, 19)
	require.Equal(t, i2, 19)
}

func TestFindClosestTwoTau5(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i1, i2 := manager.FindClosestTwoTau(0.0)
	require.Equal(t, i1, 0)
	require.Equal(t, i2, 0)
}

func TestFindClosestTau1(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestTau(0.007)
	require.Equal(t, i, 7)
}

func TestFindClosestTau2(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestTau(0.012)
	require.Equal(t, i, 12)
}

func TestFindClosestTau3(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestTau(0.018)
	require.Equal(t, i, 18)
}

func TestFindClosestTau4(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestTau(0.0)
	require.Equal(t, i, 0)
}

func TestFindClosestTau5(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	i := manager.FindClosestTau(0.02)
	require.Equal(t, i, 19)
}

func TestInterpolate1(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	res := manager.Interpolate(1, 0.01)
	require.Equal(t, res, 0.5)
}

func TestInterpolate2(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	res := manager.Interpolate(5, 0.03)
	require.Equal(t, res, 0.5)
}

func TestInterpolate3(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	res := manager.Interpolate(0, 0)
	require.Equal(t, res, 0.5)
}

func TestInterpolate4(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	res := manager.Interpolate(1.2, 0.005)
	require.Equal(t, res, 0.5)
}

func TestQuery(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	sigma := manager.Query(
		decimal.NewFromFloat(1000.0),
		decimal.NewFromFloat(1500.0),
		0.009,
	)
	require.Equal(t, sigma, 0.5)
}

func TestUpdate(t *testing.T) {
	manager, _ := CreateVolatilityManager(
		shared.ETH,
		uint64(time.Now().Unix()) + 604800,
		0.5,
		0.001,
		2,
	)
	err := manager.Update(
		decimal.NewFromFloat(1000.0),
		decimal.NewFromFloat(1000.0),
		0.9,
		0.8,
	)
	require.Equal(t, err, nil)

	sigma := manager.Query(
		decimal.NewFromFloat(1000.0),
		decimal.NewFromFloat(1000.0),
		manager.TauAnnualized(),
	)

	// Won't exactly be 0.9 b/c interpolation
	require.Equal(t, sigma != 0.5, true)

	sigma = manager.Query(
		decimal.NewFromFloat(1000.0),
		decimal.NewFromFloat(1000.0),
		0.001,
	)
	// Far away it is still 0.019
	require.Equal(t, sigma, 0.5)
}
