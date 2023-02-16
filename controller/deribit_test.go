package controller

import (
	"testing"
	"time"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestExpiryToInstrumentName(t *testing.T) {
	expiry := uint64(1665457899)
	underlying := shared.ETH
	strike, _ := decimal.NewFromString("1000")
	isCall := true
	name := ExpiryToInstrumentName(underlying, expiry, strike, isCall)
	require.Equal(t, name, "ETH-11OCT22-1000-C")
}

func TestInstrumentNameToID(t *testing.T) {
	expiry := uint64(GetNextFriday(time.Now().Unix()))
	underlying := shared.ETH
	strike, _ := decimal.NewFromString("1500")
	isCall := true
	name := ExpiryToInstrumentName(underlying, expiry, strike, isCall)
	ID, err := InstrumentNameToID(name)
	require.Equal(t, err, nil)
	require.Equal(t, ID > 0, true)
}

func TestGetDeribitSigma(t *testing.T) {
	expiry := uint64(GetNextFriday(time.Now().Unix()))
	underlying := shared.ETH
	strike, _ := decimal.NewFromString("1500")
	isCall := true
	name := ExpiryToInstrumentName(underlying, expiry, strike, isCall)
	_, _, _, err := GetDeribitSigma(name)
	require.Equal(t, err, nil)
}

// GetNextFriday - Get the timestamp for the next expiry
func GetNextFriday(timestamp int64) int64 {
	// dayOfWeek = 0 (sunday) - 6 (saturday)
	dayOfWeek := ((timestamp / 86400) + 4) % 7
	nextFriday := timestamp + ((7 + 5 - dayOfWeek) % 7) * 86400
	friday8am := nextFriday - (nextFriday % 86400) + 28800

	// If the passed timestamp is day=Friday hour>8am, we simply
    // increment it by a week to next Friday
	if timestamp >= friday8am {
		friday8am += 604800
	}
	return friday8am
}
