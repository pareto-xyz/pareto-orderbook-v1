package shared

// Underlying - Enum for supported underlying tokens
type Underlying uint8
const (
	// ETH = Ethereum
	ETH Underlying = 0
)

// String - Converts underlying to string
func (underlying Underlying) String() string {
	switch underlying {
	case ETH:
		return "eth"
	}
	return ""
}

// UInt8 - Converts underlying to an unsigned 8-bit integer
func (underlying Underlying) UInt8() uint8 {
	return uint8(underlying)
}