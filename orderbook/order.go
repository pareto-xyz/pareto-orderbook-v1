package orderbook

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/ethereum/go-ethereum/common"
)

// Order - Stores order information.
// Wraps around an Option
type Order struct {
	// Address of the creator
	Creator common.Address `json:"creator"`
	// Denotes buy or sell
	IsBuy bool `json:"isBuy"`
	// Unique identifier for the order
	ID uuid.UUID `json:"id"`
	// Time the order was placed in epoch time
	Timestamp uint64 `json:"timestamp"`
	// Amount in order
	Quantity decimal.Decimal `json:"quantity"`
	// Price within the order
	Price decimal.Decimal `json:"price"`
	// Contains option information
	Option *Option `json:"option"`
	// Is the order expired
	expired bool
}

/***********************************
 * Create function
 ***********************************/

// CreateOrder - Create a new order containing an option
func CreateOrder(
	id uuid.UUID,
	creator common.Address,
	isBuy bool,
	quantity decimal.Decimal,
	price decimal.Decimal,
	timestamp uint64,
	option *Option,
) *Order {
	return &Order{
		ID: id,
		Creator: creator,
		IsBuy: isBuy,
		Quantity: quantity,
		Price: price,
		Timestamp: timestamp,
		Option: option,
		expired: false,
	}
}

/***********************************
 * Utility functions
 ***********************************/

// IsValid - Check parameters of order are valid
func (order *Order) IsValid() bool {
	// Identifier cannot be empty
	if id := order.ID; id == uuid.Nil {
		return false
	}
	// Quantity must be > 0
	if quantity := order.Quantity; !quantity.IsPositive() {
		return false
	}
	// Price must be > 0
	if price := order.Price; !price.IsPositive() {
		return false
	}
	// Timestamp must be > 0
	if timestamp := order.Timestamp; timestamp == 0 {
		return false
	}
	// Check option is valid
	if !order.Option.IsValid() {
		return false
	}
	return true
}

// IsExpired - Check if the option in this order is expired
func (order *Order) IsExpired() (bool) {
	// Once expired, the order is already expired
	// We cache the boolean to save lookups
	if order.expired {
		return true
	}
	var expired bool = order.Option.IsExpired()
	if expired {
		order.expired = expired
	}
	return expired
}

/***********************************
 * Helper functions
 ***********************************/

// GetHash - Compute hash on an option
func GetHash(option *Option) (string, error) {
	var hash string
	hash, err := option.Hash()
	if err != nil {
		return "", err
	}
	return hash, nil
}