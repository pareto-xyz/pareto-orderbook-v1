package controller

import "math"

// GetInitialMargin - Compute the initial margin requirements
// Arguments:
// 	isBuy - If the margin is for buyer or seller
// 	spot - Spot price
// 	strike - Strike price (not level)
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualiezd)
// 	rate - Risk-free interest rate
// 	isCall - If the option is a call or put
// 	minMarginPerc - Alternative minimum percentage
func GetInitialMargin(
	isBuy bool,
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
	isCall bool,
	minMarginPerc float64,
) (margin float64) {
	// No margin if option already expired
	if tau <= 0 {
		return 0
	}
	if (isBuy) {
		// Case 1: long position
		// min(100% of mark price, 10% of spot)
		mark := GetMarkPrice(spot, strike, sigma, tau, rate, isCall)
		margin = math.Min(mark, 0.1 * spot)
	} else {
		// Case 2: short call position 
		// max((20% * spot - OTM Amount), 12.5% * spot)
		margin = math.Max(0.2 * spot - GetAmountOTM(spot, strike, isCall), 0.125 * spot)

		if (!isCall) {
			// Case 3: short put position
			// min(max((20% - OTM Amount/spot)*spot, 12.5% * spot), 50% of strike)
			margin = math.Min(margin, 0.5 * strike)
		}
	}

	// The minimum margin must be the alternative minimum
	minMargin := GetAlternativeMinimum(spot, minMarginPerc)
	margin = math.Max(minMargin, margin)

	return
}

// GetAmountOTM - Compute the amount out the money 
// Cannot be a negative amount, if ITM, returns 0
func GetAmountOTM(spot float64, strike float64, isCall bool) (amount float64) {
	if (isCall) {
		amount = math.Max(strike - spot, 0)
	} else {
		amount = math.Max(spot - strike, 0)
	}
	return
}

// GetAlternativeMinimum - Compute alternative minimum for 1 naked calls / puts.
func GetAlternativeMinimum(spot float64, percent float64) (float64) {
	return percent * spot
}
