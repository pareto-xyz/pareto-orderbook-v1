package orderbook

import (
	"testing"
	"time"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
	auth "github.com/pareto-xyz/pareto-orderbook-v1/auth"
)

// Test JSON
func TestDerivativeJSON(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	derivative := CreateDerivative()
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	order := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option,
	)
	derivative.Append(order)

	data, err := json.Marshal(derivative)
	require.Equal(t, err, nil)

	recon := CreateDerivative()
	err = json.Unmarshal(data, &recon)
	require.Equal(t, err, nil)
}
