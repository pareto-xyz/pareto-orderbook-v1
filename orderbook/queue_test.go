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

// Test creating a new queue
func TestCreateQueue(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	t.Log(queue)
}

// Test default length is zero
func TestQueueDefaultLengthZero(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	require.Equal(t, queue.Len(), 0)
}

// Test default volume is zero
func TestQueueDefaultVolumeZero(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	require.Equal(t, queue.Volume.IsZero(), true)
}

// Test default price is correct
func TestQueueDefaultPriceZero(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	if price := queue.Price; !price.Equal(decimal.New(100, 0)) {
		t.Errorf("queue.Price() = %s; want 100", price.String())
	}
}

// Test default head is nil
func TestQueueDefaultHeadNil(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	require.Equal(t, queue.Head() == nil, true)
}

// Test default tail is nil
func TestQueueDefaultTailNil(t *testing.T) {
	queue := CreateQueue(decimal.New(100, 0))
	require.Equal(t, queue.Tail() == nil, true)
}

// Test we can add two orders to queue
func TestQueueAppendOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	_, err1 := queue.Append(order1)
	_, err2 := queue.Append(order2)

	require.Equal(t, err1, nil)
	require.Equal(t, err2, nil)
}

// Test length is two
func TestQueueLengthTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	queue.Append(order1)
	queue.Append(order2)
	require.Equal(t, queue.Len(), 2)
}

// Test volume is non-zero
func TestQueueVolumeTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	queue.Append(order1)
	queue.Append(order2)

	if volume := queue.Volume; !volume.Equal(decimal.New(200, 0)) {
		t.Errorf("queue.Volume() = %s; want 200", volume.String())
	}
}

// Test head is first of two orders
func TestQueueHeadTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	head, _ := queue.Append(order1)
	queue.Append(order2)

	require.Equal(t, queue.Head(), head)
}

// Test tail is last of two orders
func TestQueueTailTwoOrders(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	queue.Append(order1)
	tail, _ := queue.Append(order2)

	require.Equal(t, queue.Tail(), tail)
}

// Test we can update an order
func TestQueueUpdateOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	head, _ := queue.Append(order1)
	queue.Append(order2)

	order3 := CreateOrder(
		uuid.New(),
		creator,
		true,
		decimal.New(200, 0),
		decimal.New(200, 0),
		uint64(time.Now().Unix()),
		option,
	)

	newHead, err := queue.Update(head, order3)
	require.Equal(t, err, nil)

	// Check volume updated
	if volume := queue.Volume; !volume.Equal(decimal.New(300, 0)) {
		t.Errorf("queue.Volume() = %s; want 300", volume.String())
	}

	// Fetch head again
	newOrder := newHead.Value.(*Order)

	// Check price is updated
	if price := newOrder.Price; !price.Equal(decimal.New(200, 0)) {
		t.Errorf("head.Price() = %s; want 200", price.String())
	}

	// Check quantity is updated
	if quantity := newOrder.Quantity; !quantity.Equal(decimal.New(200, 0)) {
		t.Errorf("head.Quantity() = %s; want 200", quantity.String())
	}
}

// Test we can remove an order
func TestQueueRemoveOrder(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option,
	)
	order2 := CreateOrder(
		orderID2,
		creator,
		true,
		decimal.New(100, 0),
		decimal.New(100, 0),
		uint64(time.Now().Unix()),
		option,
	)
	head, _ := queue.Append(order1)
	queue.Append(order2)

	_, err := queue.Remove(head)
	require.Equal(t, err, nil)

	// Fetch head again
	newOrder := queue.Head().Value.(*Order)

	// Check id is as expected
	require.Equal(t, newOrder.ID, orderID2)
}

// Test serialize and de-serializing a queue 
func TestQueueJSON(t *testing.T) {
	creator, _, _, _ := auth.CreateWallet()
	queue := CreateQueue(decimal.New(100, 0))
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
	queue.Append(order)

	data, err := json.Marshal(queue)
	require.Equal(t, err, nil)

	recon := &Queue{}
	err = json.Unmarshal(data, &recon)
	require.Equal(t, err, nil)

	require.Equal(t, recon, queue)
}