package controller

import (
	"math"
	"errors"
	"gonum.org/v1/gonum/stat/distuv"
)

// GetProbabilityFactors - Compute d1 and d2 in Black Scholes
// Formula:
// 	d1 = (log(S/K) + (r + sigma^2/2*tau)) / (sigma*sqrt(tau))
// 	tau = T - t = time to maturity
// 	r = rate of return
// 	S = spot price
// 	K = strike price
// 	sigma = implied volatility
// Arguments:
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	d1 Probability factor one 
// 	d2 Probability factor one 
func GetProbabilityFactors(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64, float64) {
	logRatio := math.Log(spot / strike)
	vol := sigma * math.Sqrt(tau)
	d1 := (logRatio + (rate + tau * math.Pow(sigma, 2) / 2)) / vol
	d2 := d1 - vol
	return d1, d2
}

// GetMarkPrice - Helper function to return call or put prices
// Arguments: 
// 	See `GetCallPrice` and `GetPutPrice`
// 	isCall - Boolean that is true if call else put
func GetMarkPrice(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
	isCall bool,
) (float64) {
	var price float64
	if (isCall) {
		price = GetCallPrice(spot, strike, sigma, tau, rate)
	} else {
		price = GetPutPrice(spot, strike, sigma, tau, rate)
	}
	return price
}

// GetCallPrice - Compute Black Scholes price of call
// Formula:
// 	C = SN(d1)-Ke^{-rt}N(d2)
func GetCallPrice(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, d2 := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	expTerm := math.Exp(-rate * tau)
	price := spot * distuv.UnitNormal.CDF(d1) - strike * expTerm * distuv.UnitNormal.CDF(d2)
	return price
}

// GetPutPrice - Compute Black Scholes price of put
// Formula:
// 	P = Ke^{-rt}N(-d2)-SN(-d1)
func GetPutPrice(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, d2 := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	expTerm := math.Exp(-rate * tau)
	price := strike * expTerm * distuv.UnitNormal.CDF(-d2) - spot * distuv.UnitNormal.CDF(-d1)
	return price
}

// GetSigmaByBisection - Solve for volatility from call price iteratively using Bisection method
// Resources:
// 	https://en.wikipedia.org/wiki/Bisection_method
// Notes:
// 	We found this to be more stable and efficient than Newton Raphson
// 	which struggles with Vega -> 0 for strikes from the spot
// Arguments:
// 	isCall - boolean that is true if option is a call
// 	maxIter - Maximum number of iterations for convergence
// 	tolerance - Maximum discrepancy to define convergence
func GetSigmaByBisection(
	spot float64,
	strike float64,
	tau float64,
	rate float64,
	price float64,
	isCall bool,
	maxIter int,
	tolerance float64,
) (float64, error) {
	if (price > strike) {
		return 0, errors.New("GetSigmaByBisection: price must be < strike")
	}

	// Initialize left and right bounds
	var leftSigma float64 = 0.001
	var rightSigma float64 = 10
	var midSigma float64 = (leftSigma + rightSigma) / 2


	for i := 0; i < maxIter; i++ {
		// Compute Black Scholes price
		leftPrice := GetMarkPrice(spot, strike, leftSigma, tau, rate, isCall)
		midPrice := GetMarkPrice(spot, strike, midSigma, tau, rate, isCall)

		// Calculate difference to trade price
		leftDiff := leftPrice - price
		midDiff := midPrice - price

		if math.Abs(midDiff) < tolerance {
			break;
		}

		// Check if the signs are the same
		if (midDiff >= 0) == (leftDiff >= 0) {
			leftSigma = midSigma;
		} else {
			rightSigma = midSigma;
		}

		// Update mid point
		midSigma = (leftSigma + rightSigma) / 2
	}

	if (midSigma <= 0) {
		return 0, errors.New("GetSigmaByBisection: sigma must be > 0")
	}
	return midSigma, nil
}
