package controller

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGetProbabilityFactors(t *testing.T) {
	var (
		d1 float64
		d2 float64
	)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, d1, -1.3418479364191978, 0.0001)
	require.InDelta(t, d2, -1.4110904450392643, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.5, 0.5, 0.0191781, 0)
	require.InDelta(t, d1, -5.821104024693793, 0.0001)
	require.InDelta(t, d2, -5.89034653331386, 0.0001)
	d1, d2 = GetProbabilityFactors(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, d1, 1.4110904450392647, 0.0001)
	require.InDelta(t, d2, 1.3418479364191982, 0.0001)
	d1, d2 = GetProbabilityFactors(1.5, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, d1, 5.890346533313859, 0.0001)
	require.InDelta(t, d2, 5.821104024693792, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.1, 0.0191781, 0)
	require.InDelta(t, d1, -6.8754217027841475, 0.0001)
	require.InDelta(t, d2, -6.889270204508161, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.9, 0.0191781, 0)
	require.InDelta(t, d1, -0.7023868482026239, 0.0001)
	require.InDelta(t, d2, -0.8270233637187437, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.5, 0.0833334, 0)
	require.InDelta(t, d1, -0.5881592189510927, 0.0001)
	require.InDelta(t, d2, -0.7324968439835144, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.5, 0.5, 0)
	require.InDelta(t, d1, -0.09280120252635217, 0.0001)
	require.InDelta(t, d2, -0.44635459311962594, 0.0001)
	d1, d2 = GetProbabilityFactors(1, 1.1, 0.5, 0.0191781, 0.001)
	require.InDelta(t, d1, -1.3274059408888652, 0.0001)
	require.InDelta(t, d2, -1.3966484495089317, 0.0001)
}

func TestGetCallPrice(t *testing.T) {
	var price float64
	price = GetCallPrice(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.0028027819781911295, 1e-6)
	price = GetCallPrice(1, 1.5, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 3.260876116140812e-11, 1e-11)
	price = GetCallPrice(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.10280278197819115, 0.0001)
	price = GetCallPrice(1.5, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.5, 0.0001)
	price = GetCallPrice(1, 1.1, 0.1, 0.0191781, 0)
	require.InDelta(t, price, 5.974057983838765e-15, 1e-15)
	price = GetCallPrice(1, 1.1, 0.9, 0.0191781, 0)
	require.InDelta(t, price, 0.016695871936353607, 0.0001)
	price = GetCallPrice(1, 1.1, 0.5, 0.0833334, 0)
	require.InDelta(t, price, 0.02308676113002428, 0.0001)
	price = GetCallPrice(1, 1.1, 0.5, 0.5, 0)
	require.InDelta(t, price, 0.10259312807045418, 0.0001)
	price = GetCallPrice(1, 1.1, 0.5, 0.0191781, 0.001)
	require.InDelta(t, price, 0.002803309731630746, 1e-6)
}

func TestGetPutPrice(t *testing.T) {
	var price float64
	price = GetPutPrice(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.10280278197819115, 0.0001)
	price = GetPutPrice(1, 1.5, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.5, 0.0001)
	price = GetPutPrice(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 0.0028027819781911295, 1e-6)
	price = GetPutPrice(1.5, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, price, 3.2608761161429214e-11, 1e-11)
	price = GetPutPrice(1, 1.1, 0.1, 0.0191781, 0)
	require.InDelta(t, price, 0.1, 0.0001)
	price = GetPutPrice(1, 1.1, 0.9, 0.0191781, 0)
	require.InDelta(t, price, 0.11669587193635367, 0.0001)
	price = GetPutPrice(1, 1.1, 0.5, 0.0833334, 0)
	require.InDelta(t, price, 0.12308676113002437, 0.0001)
	price = GetPutPrice(1, 1.1, 0.5, 0.5, 0)
	require.InDelta(t, price, 0.20259312807045438, 0.0001)
	price = GetPutPrice(1, 1.1, 0.5, 0.0191781, 0.001)
	require.InDelta(t, price, 0.1027822140239193, 0.0001)
}

func TestGetMarkPrice(t *testing.T) {
	call1 := GetCallPrice(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, call1, 0.0028027819781911295, 1e-6)
	call2 := GetCallPrice(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, call2, 0.10280278197819115, 0.0001)
	put1 := GetPutPrice(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, put1, 0.10280278197819115, 0.0001)
	put2 := GetPutPrice(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, put2, 0.0028027819781911295, 1e-6)
}

func TestGetVega(t *testing.T) {
	var vega float64
	vega = GetVega(1, 1.1, 0.5, 0.0191781, 0)
	require.InDelta(t, vega, 0.022455726633627923, 1e-6)
	vega = GetVega(1, 1.5, 0.5, 0.0191781, 0)
	require.InDelta(t, vega, 2.422277304549266e-09, 1e-10)
	vega = GetVega(1.1, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, vega, 0.022455726633627913, 1e-6)
	vega = GetVega(1.5, 1, 0.5, 0.0191781, 0)
	require.InDelta(t, vega, 2.4222773045492676e-09, 1e-10)
	vega = GetVega(1, 1.1, 0.1, 0.0191781, 0)
	require.InDelta(t, vega, 3.0022969604494246e-12, 1e-12)
	vega = GetVega(1, 1.1, 0.9, 0.0191781, 0)
	require.InDelta(t, vega, 0.043170179413510514, 1e-6)
	vega = GetVega(1, 1.1, 0.5, 0.0833334, 0)
	require.InDelta(t, vega, 0.0968727428185703, 1e-6)
	vega = GetVega(1, 1.1, 0.5, 0.5, 0)
	require.InDelta(t, vega, 0.28088269422571965, 1e-6)
	vega = GetVega(1, 1.1, 0.5, 0.0191781, 0.001)
	require.InDelta(t, vega, 0.022892751723790735, 1e-6)
}

func TestGetSigmaByBisection(t *testing.T) {
	actual := 0.5
	price := GetCallPrice(1, 1.1, actual, 0.0191781, 0)
	sigma, err := GetSigmaByBisection(1, 1.1, 0.0191781, 0, price, true, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)

	actual = 0.9
	price = GetCallPrice(1, 1.1, actual, 0.0191781, 0)
	sigma, err = GetSigmaByBisection(1, 1.1, 0.0191781, 0, price, true, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)

	actual = 0.7
	price = GetCallPrice(1, 1.1, actual, 0.0191781, 0.001)
	sigma, err = GetSigmaByBisection(1, 1.1, 0.0191781, 0.001, price, true, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)

	actual = 0.5
	price = GetPutPrice(1.1, 1, actual, 0.0191781, 0)
	sigma, err = GetSigmaByBisection(1.1, 1, 0.0191781, 0, price, false, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)

	actual = 0.9
	price = GetPutPrice(1.1, 1, actual, 0.0191781, 0)
	sigma, err = GetSigmaByBisection(1.1, 1, 0.0191781, 0, price, false, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)

	actual = 0.7
	price = GetPutPrice(1.1, 1, actual, 0.0191781, 0.001)
	sigma, err = GetSigmaByBisection(1.1, 1, 0.0191781, 0.001, price, false, 10000, 1e-4)
	require.Equal(t, err, nil)
	require.InDelta(t, sigma, actual, 5e-3)
}