package orderbook

import (
	"fmt"
	"time"
	"encoding/json"
	"crypto/sha256"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// Option - Stores option information
type Option struct {
	// Enum of underlying
	Underlying shared.Underlying `json:"underlying"`
	// Level chosen by user
	Strike shared.StrikeLevel `json:"strike"`
	// Expiry timestamp in epoch time
	Expiry uint64 `json:"expiry"`
	// Store if this is a put or call option
	IsCall bool `json:"isCall"`
}

/***********************************
 * Create function
 ***********************************/

// CreateOption - Create a new `Option` object, returns pointer
func CreateOption(
	underlying shared.Underlying,
	strike shared.StrikeLevel,
	expiry uint64,
	isCall bool,
) *Option {
	return &Option{
		Underlying: underlying,
		Strike: strike,
		Expiry: expiry,
		IsCall: isCall,
	}
}

/***********************************
 * Match function
 ***********************************/

// Matches - Check if the two options match (and are valid)
func (option *Option) Matches(op *Option) bool {
	if !option.IsValid() || !op.IsValid() {
		return false
	}
	if option.Underlying != op.Underlying {
		return false
	}
	if option.Expiry != op.Expiry {
		return false
	}
	if option.IsCall != op.IsCall {
		return false
	}
	return true
}

/***********************************
 * Helper functions
 ***********************************/

// IsValid - Check parameters of option are valid
func (option *Option) IsValid() bool {
	// Expiry must be > 0
	if expiry := option.Expiry; expiry <= 0 {
		return false
	}
	return true
}

// IsExpired - Check if an option is expired
func (option *Option) IsExpired() bool {
	ts := uint64(time.Now().Unix())  // make unsigned
	return option.Expiry < ts
}

// Tau - Return the time to expiry
// If < 0, returns 0
func (option *Option) Tau() uint64 {
	ts := uint64(time.Now().Unix())  // make unsigned
	var tau uint64
	if option.Expiry >= ts {
		tau = option.Expiry - ts
	} else {
		tau = 0
	}
	return tau
}

// TauAnnualized - Return time to expiry in years
// If < 0, returns 0
func (option *Option) TauAnnualized() float64 {
	tau := option.Tau()
	return float64(tau) / 31556952.0
}

// Hash - String representation for an option object
func (option *Option) Hash() (string, error) {
	data, err := json.Marshal(option)
	if err != nil {
		return "", err
	}
	hashBytes := sha256.Sum256(data)
	return fmt.Sprintf("%x", hashBytes), nil
}