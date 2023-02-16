package controller

import (
	"math"
	"gonum.org/v1/gonum/stat/distuv"
)

// GetDelta - Compute delta of an option
// Formula:
// 	delta = N(d1) if call
// 	delta = N(d1) - 1 if put
// Notes:
// 	https://en.wikipedia.org/wiki/Greeks_(finance)#Vega
// 	http://www.columbia.edu/~mh2078/FoundationsFE/BlackScholes.pdf
// Arguments:
// 	isCall - whether this is a call or a put 
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	delta - The delta for an option
func GetDelta(
	isCall bool,
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, _ := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	var delta float64
	if isCall {
		delta = distuv.UnitNormal.CDF(d1)
	} else {
		delta = distuv.UnitNormal.CDF(d1) - 1
	}
	return delta
}

// GetGamma - Compute gamma of an option
// Formula:
// 	gamma = N'(d1) / (sigma * spot * sqrt{tau})
// Notes:
// 	https://en.wikipedia.org/wiki/Greeks_(finance)#Vega
// 	http://www.columbia.edu/~mh2078/FoundationsFE/BlackScholes.pdf
// 	Call and put gamma are equal
// Arguments:
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	gamma - The gamma for an option
func GetGamma(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, _ := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	factor := spot * sigma * math.Sqrt(tau)
	factor = 1.0 / factor
	gamma := factor * distuv.UnitNormal.Prob(d1)
	return gamma
}

// GetTheta - Compute rho of an option
// Notes:
// 	https://en.wikipedia.org/wiki/Greeks_(finance)#Theta
// 	http://www.columbia.edu/~mh2078/FoundationsFE/BlackScholes.pdf
// Arguments:
// 	isCall - whether this is a call or a put 
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	theta - The theta for an option
func GetTheta(
	isCall bool,
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, d2 := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	term1 := -spot * distuv.UnitNormal.Prob(d1) * sigma / (2 * math.Sqrt(tau))
	var (
		theta float64
		term2 float64
	)
	if isCall {
		term2 = -rate * strike * math.Exp(-rate * tau) * distuv.UnitNormal.CDF(d2)
	} else {
		term2 = rate * strike * math.Exp(-rate * tau) * distuv.UnitNormal.CDF(-d2)
	}
	theta = term1 + term2
	return theta
}

// GetVega - Compute vega of an option
// Formula:
// 	vega = S * sqrt{tau} * N'(d1)
// Notes:
// 	https://en.wikipedia.org/wiki/Greeks_(finance)#Vega
// 	http://www.columbia.edu/~mh2078/FoundationsFE/BlackScholes.pdf
// 	Call and put vega are equal
// Arguments:
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	vega - The vega for an option
func GetVega(
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	d1, _ := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	factor := spot * math.Sqrt(tau)
	gamma := factor * distuv.UnitNormal.Prob(d1)
	return gamma
}

// GetRho - Compute rho of an option
// Notes:
// 	https://en.wikipedia.org/wiki/Greeks_(finance)#Vega
// 	http://www.columbia.edu/~mh2078/FoundationsFE/BlackScholes.pdf
// Arguments:
// 	isCall - whether this is a call or a put 
// 	spot - Spot price
// 	strike - Strike price
// 	sigma - Annualized implied volatility
// 	tau - Time to expiry (annualized)
// 	rate - The risk-free rate
// Returns:
// 	rho - The rho for an option
func GetRho(
	isCall bool,
	spot float64,
	strike float64,
	sigma float64,
	tau float64,
	rate float64,
) (float64) {
	_, d2 := GetProbabilityFactors(spot, strike, sigma, tau, rate)
	factor := strike * tau * math.Exp(-rate * tau)
	var rho float64
	if isCall {
		rho = factor * distuv.UnitNormal.CDF(d2)
	} else {
		rho = -factor * distuv.UnitNormal.CDF(-d2)
	}
	return rho
}