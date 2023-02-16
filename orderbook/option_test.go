package orderbook

import (
	"testing"
	"time"
	"encoding/json"
	"github.com/stretchr/testify/require"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// Test creating a new option
func TestCreateOption(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	t.Log(option)
}

func TestOptionGetUnderlying(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	require.Equal(t, option.Underlying, shared.ETH)
}

func TestOptionGetStrike(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	require.Equal(t, option.Strike, shared.ATM)
}

func TestOptionGetExpiry(t *testing.T) {
	expiry := uint64(time.Now().Unix()) + 60480
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		expiry,
		true,
	)
	require.Equal(t, option.Expiry, expiry)
}

func TestOptionGetIsCall(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	require.Equal(t, option.IsCall, true)
}

// Test `IsValid` function with proper option
func TestOptionIsValid(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	require.Equal(t, option.IsValid(), true)
}

// Test option serialization
func TestOptionJson(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	data, err := json.Marshal(option)
	require.Equal(t, err, nil)

	recon := &Option{};
	err2 := json.Unmarshal(data, &recon)
	require.Equal(t, err2, nil)

	// Check objects are equal
	require.Equal(t, recon, option)
}

// Test option hashing
func TestOptionHash(t *testing.T) {
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	_, err := option.Hash()
	require.Equal(t, err == nil, true)
}