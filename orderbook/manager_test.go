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

// Test creating a new manager
func TestCreateManager(t *testing.T) {
	manager := CreateManager()
	t.Log(manager)
}

// Test default length is zero
func TestManagerDefaultLengthZero(t *testing.T) {
	manager := CreateManager()
	require.Equal(t, manager.Len(), 0)
}

// Test default depth is zero
func TestManagerDefaultDepthZero(t *testing.T) {
	manager := CreateManager()
	require.Equal(t, manager.Depth, 0)
}

// Test default volume is zero
func TestManagerDefaultVolumeZero(t *testing.T) {
	manager := CreateManager()
	require.Equal(t, manager.Volume.IsZero(), true)
}

// Test default max price queue
func TestManagerDefaultMaxPriceQueue(t *testing.T) {
	manager := CreateManager()
	out := manager.MaxPriceQueue()
	require.Equal(t, out == nil, true)
}

// Test default min price queue
func TestManagerDefaultMinPriceQueue(t *testing.T) {
	manager := CreateManager()
	out := manager.MinPriceQueue()
	require.Equal(t, out == nil, true)
}

// Test default left neighbor queue
func TestManagerDefaultLeftNeighborQueue(t *testing.T) {
	manager := CreateManager()
	price := decimal.New(100, 0)
	out := manager.LeftNeighborQueue(price)
	require.Equal(t, out == nil, true)
}

// Test default right neighbor queue
func TestManagerDefaultRightNeighborQueue(t *testing.T) {
	manager := CreateManager()
	price := decimal.New(100, 0)
	out := manager.RightNeighborQueue(price)
	require.Equal(t, out == nil, true)
}

// Test length is two
func TestManagerLengthTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
	manager.Append(order1)
	manager.Append(order2)

	require.Equal(t, manager.Len(), 2)
}

// Test volume is non-zero
func TestManagerVolumeTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
	manager.Append(order1)
	manager.Append(order2)

	if volume := manager.Volume; !volume.Equal(decimal.New(200, 0)) {
		t.Errorf("manager.Volume = %s; want 200", volume.String())
	}
}

// Test appending queues to manager
func TestManagerAppendOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
	_, err := manager.Append(order)

	require.Equal(t, err, nil)
}

// Test remove queues to manager
func TestManagerRemoveOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	price := decimal.New(100, 0)
	manager := CreateManager()
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) + 604800,
		true,
	)
	orderID1 := uuid.New()
	orderID2 := uuid.New()
	order1 := CreateOrder(
		orderID1,
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option,
	)
	order2 := CreateOrder(
		orderID2,
		creator,
		true,
		decimal.New(100, 0),
		price,
		uint64(time.Now().Unix()),
		option,
	)
	head, _ := manager.Append(order1)
	manager.Append(order2)

	_, err := manager.Remove(head)
	require.Equal(t, err, nil)

	// Try to find queue at the price
	queue, ok := manager.Prices[price.String()]
	require.Equal(t, ok, true)

	// Fetch head again
	newOrder := queue.Head().Value.(*Order)

	// Check id is as expected
	require.Equal(t, newOrder.ID, orderID2)
}

// Test fetching max price on two queues
func TestManagerMaxPriceQueueTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
		decimal.New(200, 0),
		uint64(time.Now().Unix()),
		option,
	)
	manager.Append(order1)
	manager.Append(order2)

	query := decimal.New(200, 0)
	queue, ok := manager.Prices[query.String()]
	require.Equal(t, ok, true)

	maxQueue := manager.MaxPriceQueue()
	require.Equal(t, queue, maxQueue)
}

// Test fetching min price on two queues
func TestManagerMinPriceQueueTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
		decimal.New(200, 0),
		uint64(time.Now().Unix()),
		option,
	)
	manager.Append(order1)
	manager.Append(order2)

	query := decimal.New(100, 0)
	queue, ok := manager.Prices[query.String()]
	require.Equal(t, ok, true)

	minQueue := manager.MinPriceQueue()
	require.Equal(t, queue, minQueue)
}

// Test fetching left neighbor on a query price
func TestManagerLeftNeighborQueueTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
		decimal.New(200, 0),
		uint64(time.Now().Unix()),
		option,
	)
	manager.Append(order1)
	manager.Append(order2)

	// Query right in the middle
	query := decimal.New(150, 0)

	// Expected neighbor should be @100
	priceQuery := decimal.New(100, 0)
	queue, ok := manager.Prices[priceQuery.String()]
	require.Equal(t, ok, true)

	// Fetch actual left neighbor
	leftQueue := manager.LeftNeighborQueue(query)
	require.Equal(t, queue, leftQueue)
}

// Test fetching right neighbor on a query price
func TestManagerRightNeighborQueueTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
		decimal.New(200, 0),
		uint64(time.Now().Unix()),
		option,
	)
	manager.Append(order1)
	manager.Append(order2)

	// Query right in the middle
	query := decimal.New(150, 0)

	// Expected neighbor should be @200
	priceQuery := decimal.New(200, 0)
	queue, ok := manager.Prices[priceQuery.String()]
	require.Equal(t, ok, true)

	// Fetch actual right neighbor
	rightQueue := manager.RightNeighborQueue(query)
	require.Equal(t, queue, rightQueue)
}

// Test JSON
func TestManagerJSON(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	manager := CreateManager()
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
	manager.Append(order)

	data, err := json.Marshal(manager)
	require.Equal(t, err, nil)

	recon := CreateManager()
	err = json.Unmarshal(data, &recon)
	require.Equal(t, err, nil)
}
