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

// Test creation of side
func TestCreateSide(t *testing.T) {
	side := CreateSide()
	t.Log(side)
}

// Test default length is zero
func TestGroupDefaultLengthZero(t *testing.T) {
	side := CreateSide()
	require.Equal(t, side.Len(), 0)
}

// Test appending an order
func TestGroupAppendOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	side := CreateSide()
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
	_, err := side.Append(order)

	require.Equal(t, err, nil)
}

// Test length is one
func TestGroupLengthTwoOrdersSameOption(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	side := CreateSide()
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	order1 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option,
	)
	order2 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option,
	)
	_, err1 := side.Append(order1)
	_, err2 := side.Append(order2)

	require.Equal(t, err1, nil)
	require.Equal(t, err2, nil)

	require.Equal(t, side.Len(), 1)
}

// Test length is two
func TestGroupLengthTwoOrdersDiffOption(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	side := CreateSide()
	option1 := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	option2 := CreateOption(
		shared.ETH,
		shared.ITM1,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	order1 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option1,
	)
	order2 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option2,
	)
	_, err1 := side.Append(order1)
	_, err2 := side.Append(order2)

	require.Equal(t, err1, nil)
	require.Equal(t, err2, nil)

	require.Equal(t, side.Len(), 2)
}

// Test remove order 
func TestGroupRemoveOrderSameOption(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	price := decimal.New(100, 0)
	side := CreateSide()
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	order1 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option,
	)
	order2 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option,
	)
	head, _ := side.Append(order1)
	side.Append(order2)

	_, err := side.Remove(head)
	require.Equal(t, err, nil)

	require.Equal(t, side.Len(), 1)
}

func TestGroupRemoveOrderDiffOption(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	price := decimal.New(100, 0)
	side := CreateSide()
	option1 := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	option2 := CreateOption(
		shared.ETH,
		shared.ITM1,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	order1 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option1,
	)
	order2 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option2,
	)
	head, _ := side.Append(order1)
	side.Append(order2)

	_, err := side.Remove(head)
	require.Equal(t, err, nil)

	require.Equal(t, side.Len(), 1)
}

// Test JSON
func TestSideJSON(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	side := CreateSide()
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
	side.Append(order)

	data, err := json.Marshal(side)
	require.Equal(t, err, nil)

	recon := CreateSide()
	err = json.Unmarshal(data, &recon)
	require.Equal(t, err, nil)
}
