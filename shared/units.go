package shared

import (
	"math/big"
	"github.com/shopspring/decimal"
)

// GetScaleFactor - Computes 10**decimals
func GetScaleFactor(decimals int64) *big.Int {
	factor := big.NewInt(1)
	factor = factor.Exp(big.NewInt(10), big.NewInt(decimals), nil)
	return factor
}

// BigPow - Returns a ** b as a big integer.
func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

// ToQuantityScaleFactor - times 10**4
func ToQuantityScaleFactor(x decimal.Decimal) *big.Int {
	factorInt := BigPow(10, 4)
	factor := big.NewFloat(0).SetInt(factorInt)
	r := big.NewFloat(1.0)
	r = r.Mul(x.BigFloat(), factor)
	rInt, _ := r.Int(nil)
	return rInt
}

// ToPriceScaleFactor - times 10**4
func ToPriceScaleFactor(x decimal.Decimal, decimals uint8) *big.Int {
	factorInt := BigPow(10, int64(decimals))
	factor := big.NewFloat(0).SetInt(factorInt)
	r := big.NewFloat(1.0)
	r = r.Mul(x.BigFloat(), factor)
	rInt, _ := r.Int(nil)
	return rInt
}