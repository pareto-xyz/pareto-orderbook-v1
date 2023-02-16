package shared

// StrikeLevel - Enum for 11 strikes available for options
type StrikeLevel uint8
const (
	// ITM5 = ATM - 5 * increment
	ITM5 StrikeLevel = 0
	// ITM4 = ATM - 4 * increment
	ITM4 StrikeLevel = 1
	// ITM3 = ATM - 3 * increment
	ITM3 StrikeLevel = 2
	// ITM2 = ATM - 2 * increment
	ITM2 StrikeLevel = 3
	// ITM1 = ATM - 1 * increment
	ITM1 StrikeLevel = 4
	// ATM - at the money (strike = spot)
	ATM StrikeLevel = 5
	// OTM1 = ATM + 1 * increment
	OTM1 StrikeLevel = 6
	// OTM2 = ATM + 2 * increment
	OTM2 StrikeLevel = 7
	// OTM3 = ATM + 3 * increment
	OTM3 StrikeLevel = 8
	// OTM4 = ATM + 4 * increment
	OTM4 StrikeLevel = 9
	// OTM5 = ATM + 5 * increment
	OTM5 StrikeLevel = 10
)
