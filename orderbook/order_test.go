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

// Test creating a new order
func TestCreateOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	t.Log(order)
}

// Test calling `order.Id()`
func TestOrderGetId(t *testing.T) {
	id := uuid.New()
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		id,
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.ID, id)
}

// Test calling `order.Side()`
func TestOrderGetSide(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsBuy, false)
}

// Test calling `order.Quantity()`
func TestOrderGetQuantity(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	if quantity := order.Quantity; !quantity.Equal(decimal.New(100, 0)) {
		t.Errorf("order.Quantity() = %s; want 100", quantity.String())
	}
}

// Test calling `order.Price()`
func TestOrderGetPrice(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	if price := order.Price; !price.Equal(decimal.New(100, 0)) {
		t.Errorf("order.Price() = %s; want 100", price.String())
	}
}

// Test calling `order.Timestamp()`
func TestOrderGetTimestamp(t *testing.T) {
	ts := uint64(time.Now().Unix())
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		ts,
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.Timestamp, ts)
}

func TestOrderGetDerivative(t *testing.T) {
	expiry := uint64(time.Now().Unix()) + 604800
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		expiry,
		true,
	)
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		expiry,
		option,
	)
	require.Equal(t, order.Option, option)
}

// Test `IsValid` function with proper order
func TestOrderIsValid(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), true)
}

// Test `IsValid` function with bad name
func TestOrderIsValidBadName(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.Nil,
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), false)
}

// Test `IsValid` function with negative quantity
func TestOrderIsValidNegQuantity(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(-100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), false)
}

// Test `IsValid` function with negative price
func TestOrderIsValidNegPrice(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(-100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), false)
}

// Test `IsValid` function with zero quantity
func TestOrderIsValidZeroQuantity(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(0, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), false)
}

// Test `IsValid` function with zero price
func TestOrderIsValidZeroPrice(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(0, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	require.Equal(t, order.IsValid(), false)
}

// Test order serialization
func TestOrderJson(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(0, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) + 604800,
			true,
		),
	)
	data, err := json.Marshal(order)
	require.Equal(t, err, nil)

	recon := &Order{};
	err2 := json.Unmarshal(data, &recon)
	require.Equal(t, err2, nil)

	// Check actual objects are equal
	require.Equal(t, recon, order)
}

// Test expired order
func TestOrderExpired(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	order := CreateOrder(
		uuid.New(),
		creator,
		false,
		decimal.New(100, 0),
		decimal.New(0, 0),
		uint64(time.Now().Unix()),
		CreateOption(
			shared.ETH,
			shared.ATM,
			uint64(time.Now().Unix()) - 1,
			true,
		),
	);
	expired := order.IsExpired()
	require.Equal(t, expired, true)

	expired2 := order.Option.IsExpired()
	require.Equal(t, expired, expired2)
}