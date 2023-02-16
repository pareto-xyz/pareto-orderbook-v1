package orderbook

import (
	"container/list"
	"errors"
	"encoding/json"
	"github.com/shopspring/decimal"
	// Use standard Go red-black tree implementations
	rbtx "github.com/emirpasic/gods/examples/redblacktreeextended"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

// Manager - A manager is responsible for an entire side of the market
// Controls the queues at all prices for a side
type Manager struct {
	// Pointer to a tree to surface min/max prices
	priceTree *rbtx.RedBlackTreeExtended
	// Map from price (as a string) to queue
	Prices map[string]*Queue `json:"prices"`
	// Total volume in all queues
	Volume decimal.Decimal `json:"volume"`
	// Total number of orders in all queues
	NumOrders int `json:"numOrders"`
	// Amount of unique prices (size of prices map)
	Depth int `json:"depth"`
}

// NodeComparator - Order function for RB tree
func NodeComparator(a, b interface{}) int {
	return a.(decimal.Decimal).Cmp(b.(decimal.Decimal))
}

/***********************************
 * Create function
 ***********************************/

// CreateManager - Creates new manager. Returns a pointer
func CreateManager() *Manager {
	return &Manager {
		priceTree: &rbtx.RedBlackTreeExtended{
			Tree: rbt.NewWith(NodeComparator),
		},
		Prices: map[string]*Queue{},
		Volume: decimal.Zero,
		NumOrders: 0,
		Depth: 0,
	}
}

/***********************************
 * Getter functions
 * Note we do not surface the `prices` map
 ***********************************/

// Len - Returns number of orders owned by manager
func (manager *Manager) Len() int {
	return manager.NumOrders
}

// MaxPriceQueue - Returns queue of orders at the max price
// Returns nil when nothing is found
func (manager *Manager) MaxPriceQueue() *Queue {
	if manager.Depth > 0 {
		if value, found := manager.priceTree.GetMax(); found {
			return value.(*Queue)
		}
	}
	return nil
}

// MinPriceQueue - Returns queue of orders at the min price
// Returns nil when nothing is found
func (manager *Manager) MinPriceQueue() *Queue {
	if manager.Depth > 0 {
		if value, found := manager.priceTree.GetMin(); found {
			return value.(*Queue)
		}
	}
	return nil
}

// LeftNeighborQueue - Returns nearest queue with price less than threshold
func (manager *Manager) LeftNeighborQueue(
	price decimal.Decimal,
) *Queue {
	tree := manager.priceTree.Tree
	node := tree.Root

	var floor *rbt.Node

	// Start at root and search
	for node != nil {
		if tree.Comparator(price, node.Key) > 0  {
			floor = node
			node = node.Right
		} else {
			node = node.Left
		}
	}

	if floor != nil {
		return floor.Value.(*Queue)
	}
	// No queue found
	return nil
}

// RightNeighborQueue - Returns nearest queue with price greater than threshold
func (manager *Manager) RightNeighborQueue(
	price decimal.Decimal,
) *Queue {
	tree := manager.priceTree.Tree
	node := tree.Root

	var ceiling *rbt.Node

	// Start at root and search
	for node != nil {
		if tree.Comparator(price, node.Key) < 0  {
			ceiling = node
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if ceiling != nil {
		return ceiling.Value.(*Queue)
	}
	return nil
}

/***********************************
 * Setter functions
 ***********************************/

// Append - Add new order at a certain price level
func (manager *Manager) Append(order *Order) (*list.Element, error) {
	if !order.IsValid() {
		return nil, errors.New("Manager.Append: Invalid order")
	}

	price := order.Price

	// Try to find queue at the price
	queue, ok := manager.Prices[price.String()]

	if !ok {
		// No queue found, so make one
		queue = CreateQueue(price)
		manager.Prices[price.String()] = queue
		manager.priceTree.Put(price, queue)
		// Increase Depth since new price
		manager.Depth++
	}
	// Increase order count
	manager.NumOrders++

	// Update Volume
	manager.Volume = manager.Volume.Add(order.Quantity)

	// Appending returns a list pointer
	out, err := queue.Append(order)
	// Check for appending error
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Remove - Remove an existing order at a certain price level
func (manager *Manager) Remove(ptr *list.Element) (*Order, error) {
	price := ptr.Value.(*Order).Price

	// Don't need to check `ok` since `ptr` comes from a queue
	queue := manager.Prices[price.String()]

	// Remove the pointer
	popped, err := queue.Remove(ptr)

	// Check for removal error
	if err != nil {
		return nil, err
	}

	// Quantity of the popped Order
	quantity := popped.Quantity

	// Check that quantity is not too big
	if (manager.Volume.LessThan(quantity)) {
		return nil, errors.New(
			"Manager.Remove: Order quantity is larger than manager Volume")
	}

	// If queue is empty now, delete
	if queue.Len() == 0 {
		delete(manager.Prices, price.String())
		manager.priceTree.Remove(price)
		// Decrease Depth since this price is no longer offered
		manager.Depth--
	}

	// Update Volume and order count
	manager.NumOrders--
	manager.Volume = manager.Volume.Sub(popped.Quantity)

	return popped, nil
}

// MarshalJSON - Function to serialize queue into JSON
func (manager *Manager) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Prices map[string]*Queue `json:"prices"`
			Volume decimal.Decimal `json:"volume"`
			NumOrders int `json:"numOrders"`
			Depth int `json:"depth"`
		}{
			Prices: manager.Prices,
			Volume: manager.Volume,
			NumOrders: manager.NumOrders,
			Depth: manager.Depth,
		},
	)
}

// UnmarshalJSON - Function to de-serialize manager from JSON
func (manager *Manager) UnmarshalJSON(data []byte) error {
	obj := struct {
		Prices map[string]*Queue `json:"prices"`
		Volume decimal.Decimal `json:"volume"`
		NumOrders int `json:"numOrders"`
		Depth int `json:"depth"`
	}{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	manager.Prices = obj.Prices
	manager.Volume = obj.Volume
	manager.NumOrders = obj.NumOrders
	manager.Depth = obj.Depth
	manager.priceTree = &rbtx.RedBlackTreeExtended{
		Tree: rbt.NewWith(NodeComparator),
	}
	for price, queue := range manager.Prices {
		manager.priceTree.Put(decimal.RequireFromString(price), queue)
	}
	return nil
}
