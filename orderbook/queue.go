package orderbook

import (
	// The `list.List` object stores pointers for each entry so it is 
	// fast for insertion/deletion
	"encoding/json"
	"container/list"
	"errors"
	"github.com/shopspring/decimal"
)

// Queue - A queue will store orders at a single price & on one side of the market
type Queue struct {
	// List of orders sorted by timestamp
	Orders *list.List `json:"orders"`
	// Total quantity in list
	Volume decimal.Decimal `json:"volume"`
	// Price of orders in the queue
	Price decimal.Decimal `json:"price"`
}

/***********************************
 * Create function
 ***********************************/

// CreateQueue - Create a new `Queue` object, returns pointer
func CreateQueue(price decimal.Decimal) *Queue {
	return &Queue {
		Orders: list.New(),
		Volume: decimal.Zero,
		Price: price,
	}
}

/***********************************
 * Getter functions
 ***********************************/

// Len - Returns number of orders in queue
func (queue *Queue) Len() int {
	return queue.Orders.Len()
}

// Head - Returns pointer to top order in queue (earliest)
func (queue *Queue) Head() *list.Element {
	return queue.Orders.Front()
}

// Tail - Returns pointer to bottom order in queue (latest)
func (queue *Queue) Tail() *list.Element {
	return queue.Orders.Back()
}

/***********************************
 * Setter functions
 ***********************************/

// Append - Append new order to queue. Returns pointer to new element
// Returns a second argument that is nil if no errors
func (queue *Queue) Append(order *Order) (*list.Element, error) {
	// Check that order is valid
	if !order.IsValid() {
		return nil, errors.New("Queue.Append: Order is not valid")
	}
	// Update volume in queue
	queue.Volume = queue.Volume.Add(order.Quantity)
	// https://pkg.go.dev/container/list#List.PushBack
	return queue.Orders.PushBack(order), nil
}

// Update - Update list pointer with new order
// Returns a second argument that is nil if no errors
func (queue *Queue) Update(
	ptr *list.Element, 
	order *Order,
) (*list.Element, error) {
	// Check that order is valid
	if !order.IsValid() {
		return nil, errors.New("Queue.Update: Order is not valid")
	}
	// Subtract quantity from old pointer
	queue.Volume = queue.Volume.Sub(ptr.Value.(*Order).Quantity)
	// Add quantity from new pointer
	queue.Volume = queue.Volume.Add(order.Quantity)
	// Update list pointer
	ptr.Value = order
	return ptr, nil
}

// Remove - Removes order from queue using pointer
// Returns a second argument that is nil if no errors
func (queue *Queue) Remove(ptr *list.Element) (*Order, error) {
	quantity := ptr.Value.(*Order).Quantity
	if (queue.Volume.LessThan(quantity)) {
		return nil, errors.New(
			"Queue.Remove: Order quantity is larger than queue volume")
	}
	// Remove quantity from pointer to be deleted
	queue.Volume = queue.Volume.Sub(quantity)
	return queue.Orders.Remove(ptr).(*Order), nil
}

/***********************************
 * I/O functions
 ***********************************/

// MarshalJSON - Function to serialize queue into JSON
func (queue *Queue) MarshalJSON() ([]byte, error) {
	iter := queue.Head()
	var orders []*Order
	for iter != nil {
		orders = append(orders, iter.Value.(*Order))
		iter = iter.Next()
	}
	return json.Marshal(
		&struct {
			Volume decimal.Decimal `json:"volume"`
			Price decimal.Decimal `json:"price"`
			Orders []*Order `json:"orders"`
		}{
			Volume: queue.Volume,
			Price: queue.Price,
			Orders: orders,
		},
	)
}

// UnmarshalJSON - Function to de-serialize queue from JSON
func (queue *Queue) UnmarshalJSON(data []byte) error {
	obj := struct {
		Volume decimal.Decimal `json:"volume"`
		Price decimal.Decimal `json:"price"`
		Orders []*Order `json:"orders"`
	}{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	queue.Volume = obj.Volume
	queue.Price = obj.Price
	queue.Orders = list.New()
	for _, order := range obj.Orders {
		queue.Orders.PushBack(order)
	}
	return nil
}
