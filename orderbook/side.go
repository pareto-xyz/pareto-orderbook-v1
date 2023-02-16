package orderbook

import (
	"container/list"
	"errors"
)

// Side - Data structure to hold managers of different expiries/strikes
type Side struct {
	Managers map[string] *Manager `json:"managers"`
}

/***********************************
 * Create function
 ***********************************/

// CreateSide - Creates new side. Returns a pointer
func CreateSide() *Side {
	return &Side {
		// Map from hash to manager
		Managers: map[string]*Manager{},
	}
}

/***********************************
 * Getter functions
 ***********************************/

// Len - Return number of unique managers
func (side *Side) Len() int {
	return len(side.Managers)
}

// GetManager - Given an option, look up the manager for it
func (side *Side) GetManager(option *Option) (*Manager, bool) {
	hash, _ := GetHash(option)
	manager, ok := side.Managers[hash]
	return manager, ok
}

// SetManager - Set the manager for a side
func (side *Side) SetManager(option *Option, manager *Manager) {
	hash, _ := GetHash(option)
	side.Managers[hash] = manager
}

/***********************************
 * Setter functions
 ***********************************/

// Append - Add new order to side, and the correct manager underneath
func (side *Side) Append(order *Order) (*list.Element, error) {
	if !order.IsValid() {
		return nil, errors.New("Side.Append: Invalid order")
	}

	// Get hash of an option
	hash, err := GetHash(order.Option)
	if err != nil {
		return nil, err
	}
	
	// Try to find the manager for this option
	// Separate managers will be created for puts and calls
	manager, ok := side.Managers[hash]
	if !ok {
		// No manager found, so make one & store it
		manager = CreateManager()
		side.Managers[hash] = manager
	}

	// Fetch pointer from manager's `Append` function
	out, err := manager.Append(order)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Remove - Remove an existing order owned by a single manager
func (side *Side) Remove(ptr *list.Element) (*Order, error) {
	order := ptr.Value.(*Order)
	
	hash, err := GetHash(order.Option)
	if err != nil {
		return nil, err
	}

	manager := side.Managers[hash]

	out, err := manager.Remove(ptr)
	if manager.Len() == 0 {
		delete(side.Managers, hash)
	}

	if err != nil {
		return nil, err
	}
	return out, nil
}
