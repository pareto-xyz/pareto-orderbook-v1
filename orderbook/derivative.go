package orderbook

import (
	"container/list"
	"errors"
)

// Derivative - Data structure representing calls and puts
// A single derivative holds asks and bids ("Sides")
type Derivative struct {
	Asks *Side
	Bids *Side
}

/***********************************
 * Create function
 ***********************************/

// CreateDerivative - Creates a new derivative. Returns a pointer
// A "Derivative" object stores two sides: asks and bids
func CreateDerivative() *Derivative {
	return &Derivative{
		Asks: CreateSide(),
		Bids: CreateSide(),
	}
}

/***********************************
 * Getter functions
 ***********************************/

// GetManager - Retrieve a manager from the right side 
func (derivative *Derivative) GetManager(isBuy bool, option *Option) (*Manager, bool) {
	var side *Side
	if isBuy {
		side = derivative.Bids
	} else {
		side = derivative.Asks
	}
	return side.GetManager(option)
}

// SetManager - Set the manager
func (derivative *Derivative) SetManager(isBuy bool, option *Option, manager *Manager) {
	var side *Side
	if isBuy {
		side = derivative.Bids
	} else {
		side = derivative.Asks
	}
	side.SetManager(option, manager)
}

/***********************************
 * Setter functions
 ***********************************/

// Append - Add new order to derivative, and to the correct side underneat
func (derivative *Derivative) Append(order *Order) (*list.Element, error) {
	if !order.IsValid() {
		return nil, errors.New("Derivative.Append: Invalid order")
	}

	var (
		ptr *list.Element
		err error
	)
	// Punts functionality to the correct side
	if order.IsBuy {
		ptr, err = derivative.Bids.Append(order)
	} else {
		ptr, err = derivative.Asks.Append(order)
	}
	return ptr, err
}

// Remove - Remove an existing order
func (derivative *Derivative) Remove(ptr *list.Element) (*Order, error) {
	order := ptr.Value.(*Order)
	var (
		orderPtr *Order
		err error
	)
	if order.IsBuy {
		orderPtr, err = derivative.Bids.Remove(ptr)
	} else {
		orderPtr, err = derivative.Asks.Remove(ptr)
	}
	return orderPtr, err
}