package orderbook

import (
	"container/list"
	"context"
	"crypto/ecdsa"
	"errors"
	"time"
	"math/big"
	"github.com/shopspring/decimal"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	margin "github.com/pareto-xyz/pareto-orderbook-v1/contract/margin"
	oracle "github.com/pareto-xyz/pareto-orderbook-v1/contract/oracle"
	controller "github.com/pareto-xyz/pareto-orderbook-v1/controller"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// Book - Data structure for an order book
// Notes:
// 	A "Book" contains two "Derivatives", one for calls and one for puts
// 	A "Derivative" contains two "Sides", one to buy and one to sell
// 	A "Side" is made up many "Managers", organized by unique option
// 	A "Manager" is made up of many "Queues", organized by price
// 	A "Queue" is made up of many "Orders", sorted by time
// 	An "Order" contains an "Option"
type Book struct {
	// The underlying asset for this book
	Underlying shared.Underlying
	// Mapping from identifier to order pointer. These orders are across
	// all derivatives, sides, managers, and queues
	Orders map[uuid.UUID]*list.Element
	// Mapping from user address to open order UUIDs. Given a UUID, the pointer
	// can be found in "orders"
	UserOrderIds map[common.Address]map[uuid.UUID]bool
	// Storage of user addresses that enter in positions in the current round
	UsersWithPositions map[common.Address]bool
	// Pointers to the two sides
	Calls *Derivative
	Puts *Derivative
	// True if the book is paused
	Paused bool
	// True if the book is setteld
	Settled bool
	// Margin contract data
	MarginContract *MarginContractData
	// Oracle contract data
	OracleContract *OracleContractData
	// Volatility data
	VolManager *controller.VolatilityManager
}

// MarginContractData - Stores data from smart contracts
type MarginContractData struct {
	// Address of contract 
	DeployedAddress common.Address
	// The active expiry for new options
	ActiveExpiry int64
	// The strikes for the current round
	RoundStrikes [11]decimal.Decimal
	// Minimum order size
	MinQuantity decimal.Decimal
	// Collateral decimals
	CollateralDecimals uint8
}

// OracleContractData - Stores data from smart contracts
type OracleContractData struct {
	// Address of contract 
	DeployedAddress common.Address
}

// PriceLevel - Data structure used in reporting depth
type PriceLevel struct {
	Price decimal.Decimal `json:"price"`
	Quantity decimal.Decimal `json:"quantity"`
}

/***********************************
 * Create function
 ***********************************/

// CreateBook - Creates new book. Returns a pointer
func CreateBook(
	underlying shared.Underlying,
	marginContractAddress common.Address,
	oracleContractAddress common.Address,
) (*Book, error) {
	book := &Book{
		Underlying: underlying,
		Orders: map[uuid.UUID]*list.Element{},
		UserOrderIds: map[common.Address]map[uuid.UUID]bool{},
		UsersWithPositions: map[common.Address]bool{},
		Calls: CreateDerivative(),
		Puts: CreateDerivative(),
		Paused: false,
		Settled: false,
		MarginContract: &MarginContractData{
			DeployedAddress: marginContractAddress,
		},
		OracleContract: &OracleContractData{
			DeployedAddress: oracleContractAddress,
		},
	}

	// Fetch updated contract data
	if err := book.UpdateMarginContractData(); err != nil {
		return nil, err
	}

	// Build a volatility surface
	// Notes:
	// - We set the minMoneyness to 0.01 and the maxMoneyness to 2
	// meaning we do not expect the moneyness ratio: (spot/strike)
	// to be <0.01 or >2
	// - We initialize the volatility surface to be 50% everywhere
	// - The expiry is fetched from an oracle
	volManager, err := controller.CreateVolatilityManager(
		underlying,
		book.GetActiveExpiry(),
		0.5,
		0.01,
		2,
	)
	if err != nil {
		return nil, err
	}

	// Save surface to book
	book.VolManager = volManager

	return book, nil
}

/***********************************
 * Getter functions 
 ***********************************/

// GetActiveExpiry - Get the current active expiry
func (book *Book) GetActiveExpiry() (uint64) {
	return uint64(book.MarginContract.ActiveExpiry)
}

// GetActiveAnnualizedTau - Get the (annualized) time to expiry
func (book *Book) GetActiveAnnualizedTau() (float64) {
	now := uint64(time.Now().Unix())
	expiry := uint64(book.MarginContract.ActiveExpiry)
	var tau uint64
	if expiry >= now {
		tau = expiry - now
	} else {
		tau = 0
	}
	return float64(tau) / 31556952.0
}

// GetMinQuantity - Get the minimum order quantity
func (book *Book) GetMinQuantity() (decimal.Decimal) {
	return book.MarginContract.MinQuantity
}

// GetOrderByID - Returns an order by identifier
func (book *Book) GetOrderByID(orderID uuid.UUID) (*Order, error) {
	ptr, ok := book.Orders[orderID]
	if !ok {
		return nil, errors.New("book.GetOrderByID: identifier doesn't exist")
	}
	return ptr.Value.(*Order), nil
}

// GetOrders - Return unmatched order by user
// Arguments:
// 	user - Address of user to look up
func (book *Book) GetOrders(user common.Address) (orders []*Order, err error) {
	orderIds, ok := book.UserOrderIds[user]
	if !ok {
		return orders, nil
	}
	for uuid, ok := range orderIds {
		if ok {
			ptr := book.Orders[uuid]
			order := ptr.Value.(*Order)
			orders = append(orders, order)
		}
	}
	return orders, nil
}

// GetPositions - Return positions (matched orders) by user
// Arguments:
// 	user - Hex string of address to look up
func (book *Book) GetPositions(user common.Address) (positions []margin.DerivativeOrder, err error) {
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return positions, err
	}
	// Make call from instance
	callOpts := bind.CallOpts{From: user}
	positions, err = instance.GetPositions(&callOpts, book.Underlying.UInt8())
	if err != nil {
		return positions, err
	}
	return positions, err
}

// GetDepth - For a specific option, returns all unique prices and volumes
func (book *Book) GetDepth(option *Option) (
	asks, bids []*PriceLevel, 
	err error,
) {
	// Find manager for the ask (sell) side
	ask, atLeastOneAsk := book.GetAskManager(option)
	// Find manager with the right option on the bid (buy) side
	bid, atLeastOneBid := book.GetBidManager(option)
	// If we can't find either
	if (!atLeastOneAsk) && (!atLeastOneBid) {
		return asks, bids, nil
	}
	// Loop through queues and append price and volume
	level := ask.MaxPriceQueue()
	for level != nil {
		asks = append(
			asks,
			&PriceLevel{
				Price: level.Price,
				Quantity: level.Volume,
			},
		)
		level = ask.LeftNeighborQueue(level.Price)
	}
	// Loop through queues and append price and volume
	level = bid.MaxPriceQueue()
	for level != nil {
		bids = append(
			bids,
			&PriceLevel{
				Price: level.Price,
				Quantity: level.Volume,
			},
		)
		level = bid.LeftNeighborQueue(level.Price)
	}

	return
}

// IsEmpty - Check if the book is empty
func (book *Book) IsEmpty() bool {
	return len(book.Orders) == 0
}

// GetUnderlying - Returns the underlying asset address
func (book *Book) GetUnderlying() shared.Underlying {
	return book.Underlying
}

// GetPrice - Gets market price for requested quantity of asset
// Dev:
// 	- Simulate prices you would get over orders
// 	- This is critically different than the Black Scholes mark price
func (book *Book) GetPrice(
	isBuy bool,
	quantity decimal.Decimal,
	option *Option,
) (
	price decimal.Decimal,
	err error,
) {
	var (
		level *Queue
		getQueue func(decimal.Decimal) *Queue
	)

	if isBuy {
		// Find manager on "ask" (sell) side of book
		// Get lowest price someone is willing to sell at
		ask, ok := book.GetAskManager(option)
		if !ok {
			return decimal.Zero, errors.New("Book.GetPrice: option not found")
		}
		// Get lowest price
		level = ask.MinPriceQueue()
		getQueue = ask.RightNeighborQueue
	} else {
		// Find manager on "bid" (buy) side of book
		// Get highest price someone is willing to buy at
		bid, ok := book.GetBidManager(option)
		if !ok {
			return decimal.Zero, errors.New("Book.GetPrice: option not found")
		}
		// Get highest price
		level = bid.MaxPriceQueue()
		getQueue = bid.LeftNeighborQueue
	}

	// Loop until either quantity is zero (request is all fulfilled)
	// or the level is null, meaning we ran out of queues
	for quantity.Sign() > 0 && level != nil {
		levelVolume := level.Volume
		levelPrice := level.Price

		if quantity.GreaterThanOrEqual(levelVolume) {
			price = price.Add(levelPrice.Mul(levelVolume))
			quantity = quantity.Sub(levelVolume)
			// Get next queue - do not error here
			level = getQueue(levelPrice)
		} else {
			// In this branch, we fulfill the order
			// Note: use `quantity` which is < `levelVolume`
			price = price.Add(levelPrice.Mul(quantity))
			// Break loop
			quantity = decimal.Zero
		}
	}

	if quantity.Sign() > 0 {
		return decimal.Zero, errors.New("Book.GetPrice: leftover quantity")
	}
	return
}

// GetSpot - Get the spot price for underlying in the book
// Calls the oracle smart contract through a read request
// Returns:
// 	price (decimal.Decimal) - Median price over spot oracles
func (book *Book) GetSpot() (decimal.Decimal, error) {
	instance, err := book.GetOracleContract()
	if err != nil {
		return decimal.Zero, err
	}
	_, spotBn, _, decimals, err := instance.LatestRoundSpot(nil)
	if err != nil {
		return decimal.Zero, err
	}
	spot := decimal.NewFromBigInt(spotBn, -int32(decimals))
	return spot, nil
}

// GetInterestRate - Get the interest rate for USDC in the book
// Calls the oracle smart contract through a read request
// 
// Returns:
// 	rate (decimal.Decimal) - Median rate over data sources
func (book *Book) GetInterestRate() (decimal.Decimal, error) {
	instance, err := book.GetOracleContract()
	if err != nil {
		return decimal.Zero, err
	}
	_, rateBn, _, decimals, err := instance.LatestRoundRate(nil)
	if err != nil {
		return decimal.Zero, err
	}
	rate := decimal.NewFromBigInt(rateBn, -int32(decimals))
	return rate, nil
}

// GetSigma - Get the implied volatility for underlying in book
// Arguments:
// 	strikeLevel (StrikeLevel) - Chosen strike level
// Returns:
// 	sigma (float64) - Implied volatility from nearest point in surface
// Dev:
// 	Spot price is fetched using the "SpotManager"
// 	Expiry is fetched from the active expiry
func (book *Book) GetSigma(strikeLevel shared.StrikeLevel) (float64) {
	spot, _ := book.GetSpot()
	strike := book.LevelToStrike(strikeLevel)
	tauAnnualized := book.VolManager.TauAnnualized()
	sigma := book.VolManager.Query(spot, strike, tauAnnualized)
	return sigma
}

// GetMark - Get the Black Scholes mark price
// Arguments:
// 	isCall (bool) - Pricing for call or put?
// 	strikeLevel (StrikeLevel) - Chosen strike level
// 	spot (decimal.Decimal) - Spot price
// 	interestRate (decimal.Decimal) - Interest rate
// Returns:
// 	mark (decimal.Decimal) - Mark price
// Dev:
// 	Spot price is fetched using the "SpotManager"
// 	Expiry is fetched from the active expiry
func (book *Book) GetMark(
	isCall bool,
	strikeLevel shared.StrikeLevel,
	spot decimal.Decimal,
	interestRate decimal.Decimal,
) (decimal.Decimal) {
	strike := book.LevelToStrike(strikeLevel)
	tauAnnualized := book.VolManager.TauAnnualized()
	sigma := book.VolManager.Query(spot, strike, tauAnnualized)
	markFloat := controller.GetMarkPrice(
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
		isCall,
	)
	mark := decimal.NewFromFloat(markFloat)
	return mark
}

// GetGreeks - Get the Black Scholes greeks
// Arguments:
// 	isCall (bool) - Pricing for call or put?
// 	strikeLevel (StrikeLevel) - Chosen strike level
// Returns:
// 	delta (decimal.Decimal)
// 	gamma (decimal.Decimal)
// 	theta (decimal.Decimal)
// 	vega (decimal.Decimal)
// 	rho (decimal.Decimal)
// Dev:
// 	Spot price is fetched using the "SpotManager"
// 	Expiry is fetched from the active expiry
func (book *Book) GetGreeks(
	isCall bool,
	strikeLevel shared.StrikeLevel,
) (
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	error,
) {
	spot, err := book.GetSpot()
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	interestRate, err := book.GetInterestRate()
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	strike := book.LevelToStrike(strikeLevel)
	tauAnnualized := book.VolManager.TauAnnualized()
	sigma := book.VolManager.Query(spot, strike, tauAnnualized)
	delta64 := controller.GetDelta(
		isCall,
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
	)
	delta := decimal.NewFromFloat(delta64)
	gamma64 := controller.GetGamma(
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
	)
	gamma := decimal.NewFromFloat(gamma64)
	theta64 := controller.GetTheta(
		isCall,
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
	)
	theta := decimal.NewFromFloat(theta64)
	vega64 := controller.GetVega(
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
	)
	vega := decimal.NewFromFloat(vega64)
	rho64 := controller.GetRho(
		isCall,
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
	)
	rho := decimal.NewFromFloat(rho64)

	return delta, gamma, theta, vega, rho, nil
}

// GetStrikes - Get strike prices
// Returns:
// 	strikes ([11]decimal.Decimal) - 11 strike prices for current expiry
func (book *Book) GetStrikes() ([11]decimal.Decimal) {
	return book.MarginContract.RoundStrikes
}

/***********************************
 * Setter functions 
 ***********************************/

// ProcessOrder - Helper function to process an order
// Notes:
// 	Called by `CreateMarketOrder` and `CreateLimitOrder`
// 	Performs a single match between request and book. If the top order does 
// 	not fully satisfy the request, it does NOT automatically match next order.
// 	One should loop and call this function repeatedly if they wish to fulfill fully.
//  A matched order in the book must also match the requested option
// Arguments:
// 	requestor - Owner of either a market or limit order that was matched with an existing order
// 	isBuy - Whether the requestor is buying or selling
// 	isTaker - Whether the requestor is making a limit or market order
// 	queue - Queue or orders at a single price 
// 	quantityRequested - Amount requested to be fulfilled
// 	option - A pointer to an Option object
// 	privateKeyHex - String of the master private key
// 	chainID - Ethereum chain ID
// Returns:
// 	partial - If `quantityLeft > 0`, returns pointer to a new order that replaces the matched order.
// 	 Set to nil if matched order is fulfilled fully
// 	quantityLeft - Amount in the matched order not used up by request
// 	 Set to zero if matched order is fulfilled fully
// 	err - not nil if something went wrong
// Dev:
// 	Updates the volatility surface
func (book *Book) ProcessOrder(
	requestor common.Address,
	isBuy bool,
	isTaker bool,
	queue *Queue,
	quantityRequested decimal.Decimal,
	option *Option,
	privateKeyHex string,
	chainID int64,
) (
	partial *Order,
	quantityLeft decimal.Decimal,
	err error,
) {
	// Check if `quantityRequested` is negative
	if quantityRequested.Sign() <= 0 {
		err := errors.New("Book.ProcessOrder: Quantity is negative")
		return nil, decimal.Zero, err
	}

	// Check queue is not nil
	if queue == nil {
		err := errors.New("Book.ProcessOrder: Queue is undefined")
		return nil, decimal.Zero, err
	}

	// Check queue is not empty
	if queue.Len() == 0 {
		err := errors.New("Book.ProcessOrder: Queue is empty")
		return nil, decimal.Zero, err
	}

	// Get top order from queue (which we know is not empty)
	var orderPtr *list.Element
	for queue.Len() > 0 {
		orderPtr = queue.Head()
		order := orderPtr.Value.(*Order)
		expired := order.IsExpired()

		if expired {
			// Remove the stale order
			book.CancelOrder(order.ID)
		} else {
			// Found a valid order pointer
			break
		}
	}
	if queue.Len() == 0 && orderPtr == nil {
		// We didn't find one 
		err := errors.New("Book.ProcessOrder: Queue is empty post-prune")
		return nil, decimal.Zero, err
	}
	order := orderPtr.Value.(*Order)

	// If the requested quantity < order size, then we will
	// need to edit the matched order
	if quantityRequested.LessThan(order.Quantity) {
		// Create a new order to replace existing order
		partial = CreateOrder(
			order.ID,  // copies the id!
			requestor,
			order.IsBuy,
			// Reduce quantity to show remaining amount
			order.Quantity.Sub(quantityRequested),
			order.Price,
			order.Timestamp,
			order.Option,
		)

		// Replace matched order with this one
		_, err = queue.Update(orderPtr, partial)
		
		if err != nil {
			return nil, decimal.Zero, err
		}

		// Full quantity has been satisfied
		quantityLeft = decimal.Zero
	} else {  
		// Otherwise, no partial order is needed
		quantityLeft = quantityRequested.Sub(order.Quantity)

		// Cancel order since completely filled
		_, err := book.CancelOrder(order.ID)
		
		if err != nil {
			return nil, decimal.Zero, err
		}
	}

	// Update the volatility smile
	// If this fails, we still proceed but an error will be returned
	spot, _ := book.GetSpot()
	interestRate, _ := book.GetInterestRate()
	book.UpdateVolatility(spot, interestRate, order)

	// Post order on chain
	err = book.PostMatchedOrderOnchain(
		requestor,
		isBuy,
		isTaker,
		quantityRequested,
		order,
		privateKeyHex,
		chainID,
	)

	if err != nil {
		return nil, decimal.Zero, err
	}

	// Add the users to our records
	book.UsersWithPositions[requestor] = true
	book.UsersWithPositions[order.Creator] = true

	return
}

// CancelOrder - Removes order with id from order book
// The `orderID` must be the identifier of the order to cancel
// Orders can be cancelled when paused
/// @dev Since IM on open orders is purely additive (no netting)
/// we do not need to perform margin checks
func (book *Book) CancelOrder(orderID uuid.UUID) (*Order, error) {
	ptr, ok := book.Orders[orderID]

	if !ok {
		err := errors.New("Book.CancelOrder: Incorrect identifier")
		return nil, err
	}

	order := ptr.Value.(*Order)

	// Remove from "book.Orders"
	delete(book.Orders, orderID)

	// Remove from "book.UserOrderIds"
	delete(book.UserOrderIds[order.Creator], orderID)
	if len(book.UserOrderIds[order.Creator]) == 0 {
		delete(book.UserOrderIds, order.Creator)
	}

	var (
		popped *Order
		err error
	)

	// Remove from book's storage
	if order.Option.IsCall {
		popped, err = book.Calls.Remove(ptr)
		if err != nil {
			return nil, err
		}
	} else {
		popped, err = book.Puts.Remove(ptr)
		if err != nil {
			return nil, err
		}
	}
	return popped, nil
}

// CreateMarketOrder - A market order immediately fulfills a limit order on the book
// Notes:
//  Finds the best price: buy @ lowest or sell @ highest
//  Keeps matching until order fully fulfilled
// Dev:
// 	Calls `ProcessOrder` under the hood
// Arguments:
// 	creator - Address for the creator 
// 	isBuy - true (buy) or false (sell)
// 	quantity - Amount to buy or sell
// 	option - A pointer to an Option object
// 	privateKeyHex - String of the master private key
// 	chainID - Ethereum chain ID
// Returns:
// 	quantityLeft - Amount of quantity in order that is not fulfilled
// 	err - Not nil if something went wrong internally
func (book *Book) CreateMarketOrder(
	creator common.Address,
	isBuy bool,
	quantity decimal.Decimal,
	option *Option,
	privateKeyHex string,
	chainID int64,
) (
	quantityLeft decimal.Decimal,
	err error,
) {
	// Initialize the amount to be matched to the whole quantity
	quantityLeft = quantity

	// Can't do anything if book is paused
	if book.Paused {
		err = errors.New("Book.CreateMarketOrder: Book is paused")
		return
	}
	// Check if quantity is negative
	if quantity.Sign() <= 0 {
		err = errors.New("Book.CreateMarketOrder: Quantity is negative")
		return
	}
	// Check that option is valid
	if err = book.IsValidOption(option); err != nil {
		return
	}
	// Check if user can handle new order
	satisfied, _, err := book.CheckInitialMarginNewOrder(creator, isBuy, quantity, option)
	if err != nil {
		return
	}
	if !satisfied {
		err = errors.New("Book.CreateMarketOrder: fails initial margin check")
		return
	}

	// Create variables that differ depending on `side`
	var (
		getQueue func() *Queue
		manager *Manager
		ok bool
	)

	// Get the queue from the book side
	if isBuy {
		// Find sell order at the lowest price
		manager, ok = book.GetAskManager(option)
		if !ok {
			err = errors.New("Book.CreateMarketOrder: Market side empty")
			return
		}
		getQueue = manager.MinPriceQueue
	} else {
		manager, ok = book.GetBidManager(option)
		if !ok {
			err = errors.New("Book.CreateMarketOrder: Market side empty")
			return
		}
		getQueue = manager.MaxPriceQueue
	}

	if manager.Len() == 0 {
		err = errors.New("Book.CreateMarketOrder: Market side empty")
		return
	}

	// Process the order & check for error
	for quantityLeft.Sign() > 0 && manager.Len() > 0 {
		queue := getQueue()
		_, quantityLeft, err = book.ProcessOrder(
			creator,
			isBuy,
			true,
			queue,
			quantityLeft,
			option,
			privateKeyHex,
			chainID,
		)
		if err != nil {
			return
		}
	}

	// If we reached here, then return `quantityLeft` as is 
	return
}

// CreateLimitOrder - A limit order potentially adds a new order to the book. 
// We must check that no match exists first. If so, fulfill before adding
// Notes:
// 	https://www.schwab.com/learn/story/3-order-types-market-limit-and-stop-orders
// Dev:
// 	Calls `ProcessOrder` under the hood
// Arguments:
// 	creator - Address for the creator 
// 	isBuy - true (buy) or false (sell)
// 	quantity - Amount to buy or sell
// 	price - Limit price: either the maximum price to be paid or the 
// 	        minimum price to be received. If such a price does not exist,
// 	        then no match will be executed
// 	option - A pointer to an Option object
// 	privateKeyHex - String of the master private key
// 	chainID - Ethereum chain ID
// Returns:
// 	quantityLeft - Amount of quantity in order that is not automatically fulfilled.
// 	 This is the amount of quantity actually placed in the limit order
// 	err - Not nil if something went wrong internally
func (book *Book) CreateLimitOrder(
	creator common.Address,
	isBuy bool,
	quantity decimal.Decimal,
	price decimal.Decimal,
	option *Option,
	privateKeyHex string,
	chainID int64,
) (
	order *Order,
	quantityLeft decimal.Decimal,
	err error,
) {
	// Initialize the amount to be matched to the whole quantity
	quantityLeft = quantity

	// Can't do anything if book is paused
	if book.Paused {
		err = errors.New("Book.CreateLimitOrder: Book is paused")
		return
	}
	id := uuid.New()
	if _, ok := book.Orders[id]; ok {
		err = errors.New("Book.CreateLimitOrder: Order already exists")
		return
	}
	if len(id) == 0 {
		err = errors.New("Book.CreateLimitOrder: Empty identifier")
		return
	}
	if quantity.Sign() <= 0 {
		err = errors.New("Book.CreateLimitOrder: Quantity <= 0")
		return
	}
	if price.Sign() <= 0 {
		err = errors.New("Book.CreateLimitOrder: Price <= 0")
		return
	}
	// Check that option is valid
	if err = book.IsValidOption(option); err != nil {
		return
	}
	// Check if user can handle new order
	satisfied, _, err := book.CheckInitialMarginNewOrder(creator, isBuy, quantity, option)
	if err != nil {
		return
	}
	if !satisfied {
		err = errors.New("Book.CreateLimitOrder: fails initial margin check")
		return
	}

	var (
 		comparator func(decimal.Decimal) bool
		getQueue func() *Queue
		manager *Manager
		other *Manager
		ok bool
	)

	if isBuy {
		manager, ok = book.GetAskManager(option)
		if !ok {
			manager = CreateManager()
			book.SetAskManager(option, manager)
		}
		other, ok = book.GetBidManager(option)
		if !ok {
			other = CreateManager()
			book.SetBidManager(option, other)
		}
		comparator = price.GreaterThanOrEqual
		getQueue = manager.MinPriceQueue
	} else {
		manager, ok = book.GetBidManager(option)
		if !ok {
			manager = CreateManager()
			book.SetBidManager(option, manager)
		}
		other, ok = book.GetAskManager(option)
		if !ok {
			other = CreateManager()
			book.SetAskManager(option, other)
		}
		comparator = price.LessThanOrEqual
		getQueue = manager.MaxPriceQueue
	}

	// Init queue since used in condition
	queue := getQueue()

	// Keep matching until either quantity is fulfilled or queue is out
	// `quantityLeft` gets reduced as the order is partially filled
	for quantityLeft.Sign() > 0 && manager.Len() > 0 && comparator(queue.Price) {
		// Note this sets `quantityLeft` to leftover quantity
		_, quantityLeft, err = book.ProcessOrder(
			creator,
			isBuy,
			false,
			queue,
			quantityLeft,
			option,
			privateKeyHex,
			chainID,
		)
		if err != nil {
			return
		}

		// Prepare for next round
		queue = getQueue()
	}

	// If we unable to fulfill total request, make a new limit order
	if quantityLeft.Sign() > 0 {
		order = CreateOrder(
			id,
			creator,
			isBuy,
			quantityLeft,
			price,
			uint64(time.Now().Unix()),
			option,
		)

		var orderPtr *list.Element
		// Update group, manager, and book with new order
		// This will update depth, volume, etc.
		orderPtr, err = other.Append(order)
		if err != nil {
			return
		}

		// Add to book's collection orders
		book.Orders[id] = orderPtr

		// Add to book's user collection
		if _, ok := book.UserOrderIds[creator]; !ok {
			book.UserOrderIds[creator] = map[uuid.UUID]bool{}
		}
		book.UserOrderIds[creator][id] = true
	}

	return
}

/***********************************
 * Margin functions 
 ***********************************/

// CheckInitialMarginNewOrder - Check if the user has enough margin: available 
// balance must be higher than the margin requirements for this new order
// Arguments:
// 	creator - Address of the option creator
// 	isBuy - true if buying, else false if selling
// 	quantity - Amount of the option the user is trying to buy/sell
// 	option - Option under consideration
func (book *Book) CheckInitialMarginNewOrder(
	creator common.Address, 
	isBuy bool,
	quantity decimal.Decimal,
	option *Option,
) (bool, decimal.Decimal, error) {
	// Compute the available balance
	_, availableBalance, err := book.CheckInitialMargin(creator)
	if err != nil {
		return false, decimal.Zero, err
	}
	// Get IM for new order
	marginOrder, err := book.GetInitialMarginNewOrder(isBuy, quantity, option)
	if err != nil {
		return false, decimal.Zero, err
	}
	// Balance - Margin
	remaining := availableBalance.Sub(marginOrder)
	satisfied := (remaining.Sign() != -1)

	return satisfied, remaining, nil
}

// GetInitialMarginNewOrder - Get the initial margin requirements for new order
// Arguments:
// 	isBuy - true if buying, else false if selling
// 	quantity - Amount of the option the user is trying to buy/sell
// 	option - Option under consideration
func (book *Book) GetInitialMarginNewOrder(
	isBuy bool,
	quantity decimal.Decimal,
	option *Option,
) (decimal.Decimal, error) {
	// Get option parameters
	spot, _ := book.GetSpot()
	interestRate, _ := book.GetInterestRate()
	strike := book.LevelToStrike(option.Strike)
	tauAnnualized := option.TauAnnualized()
	sigma := book.VolManager.Query(spot, strike, tauAnnualized)

	// Boot up the smart contract
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return decimal.Zero, err
	}

	// Fetch the minimum margin percentage from the contract
	minMarginPerc, err := book.getMinMarginPerc(instance)
	if err != nil {
		return decimal.Zero, err
	}

	// Compute margin on a single unit
	marginUnit64 := controller.GetInitialMargin(
		isBuy,
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		sigma,
		tauAnnualized,
		interestRate.InexactFloat64(),
		option.IsCall,
		minMarginPerc.InexactFloat64(),
	)
	marginUnit := decimal.NewFromFloat(marginUnit64)
	// Amount of margin for `quantity` units
	marginOrder := marginUnit.Mul(quantity)

	return marginOrder, nil
}

// CheckInitialMargin - Check if the user has enough balance in their margin account 
// on the smart contract for their open orders and positions
// Returns:
// 	satisifed - true (passes check) or false (fails check)
// 	availableBalance - Amount of extra balance minus unrealized losses and margin reqs
func (book *Book) CheckInitialMargin(user common.Address) (bool, decimal.Decimal, error) {
	_, _, _, _, _, availableBalance, _, err := book.GetAccount(user)
	if err != nil {
		return false, decimal.Zero, err
	}
	// availableBalance must be >= 0
	satisfied := (availableBalance.Sign() != -1)
	return satisfied, availableBalance, nil
}

// GetAccount - Get various statistics on account health
// Arguments:
// 	user (common.Address) - Address for the account to measure health of
// Returns:
// 	balance (decimal.Decimal) - Balance in margin account
// 	upnl (decimal.Decimal) - Unrealized P&L
// 	orderIM (decimal.Decimal) - Initial margin for orders
// 	positionIM (decimal.Decimal) - Initial margin for positions
// 	positionMM (deimal.Decimal) - Maintenance margin for positions
// 	availableBalance (decimal.Decimal) - Available balance for new orderes
// 	liquidationBuffer (decimal.Decimal) - Buffer until liquidation
func (book *Book) GetAccount(user common.Address) (
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	decimal.Decimal,
	error,
) {
	// Boot up the smart contract
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	decimals := book.MarginContract.CollateralDecimals
	// Fetch the balance for the user from the smart contract
	balance, err := book.GetMarginAccountBalance(instance, user, decimals)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	// Fetch the unrealized losses from the smart contract for user's positions
	upnl, err := book.GetUnrealizedPnL(instance, user, decimals)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	// Compute the IM requirements for all open orders
	orderIM, err := book.GetInitialMarginOrders(instance, user)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	// Fetch the IM requirements for user's positions
	positionIM, err := book.GetInitialMarginPositions(instance, user, decimals)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	// Fetch the MM requirements for user's positions
	positionMM, err := book.GetMaintenanceMarginPositions(instance, user, decimals)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	// Check if balance + UPnL > IM(orders) + IM(positions)
	availableBalance := balance.Add(upnl).Sub(orderIM.Add(positionIM))

	// Check if balance + UPnL > MM(positions)
	liquidationBuffer := balance.Add(upnl).Sub(positionMM)

	return balance, upnl, orderIM, positionIM, positionMM, availableBalance, liquidationBuffer, nil
}

// GetMarginAccountBalance - Get the amount of USDC in user's margin account
// @dev Requires a smart contract call
// Arguments:
// 	instance - Margin account instance
// 	user - Account user
func (book *Book) GetMarginAccountBalance(
	instance *margin.Margin,
	user common.Address,
	decimals uint8,
) (decimal.Decimal, error) {
	// https://geth.ethereum.org/docs/dapp/native-bindings
	balanceBn, err := instance.GetBalanceOf(nil, user)
	if err != nil {
		return decimal.Zero, err
	}
	balance := decimal.NewFromBigInt(balanceBn, -int32(decimals))
	return balance, nil
}

// getMinMarginPerc - Get the minimum margin percentage
func (book *Book) getMinMarginPerc(instance *margin.Margin) (decimal.Decimal, error) {
	// Fetch the minimum margin percentage from the contract
	minMarginPercBn, err := instance.MinMarginPerc(nil)
	var minMarginPerc decimal.Decimal
	if err != nil {
		return decimal.Zero, err
	}
	minMarginPerc = decimal.NewFromBigInt(minMarginPercBn, -int32(4))
	return minMarginPerc, nil
}

// GetInitialMarginOrders - Get initial margin of all unmatched orders in orderbook
// These are NOT netted on purpose: it is expensive to open many orders. Only 
// positions have netted margining
func (book *Book) GetInitialMarginOrders(
	instance *margin.Margin,
	user common.Address,
) (decimal.Decimal, error) {
	// Fetch the latest interest rate and spot 
	spot, _ := book.GetSpot()
	interestRate, _ := book.GetInterestRate()

	// Fetch the minimum margin percentage from the contract
	minMarginPerc, err :=  book.getMinMarginPerc(instance)
	if err != nil {
		return decimal.Zero, err
	}

	// Initialize to zero and we will add to it
	margin := decimal.Zero

	// If no entry, then no orders
	if orders, ok := book.UserOrderIds[user]; ok {
		for uuid, check := range orders {
			// If "check" is false, then the order is "turned off"
			if !check {
				continue
			}
			ptr := book.Orders[uuid]
			order := ptr.Value.(*Order)

			// Ignore any orders with zero quantity
			if order.Quantity.Sign() <= 0 {
				continue
			}

			strike := book.LevelToStrike(order.Option.Strike)
			tauAnnualized := order.Option.TauAnnualized()
			sigma := book.VolManager.Query(spot, strike, tauAnnualized)

			// Compute margin on a single unit
			marginUnit64 := controller.GetInitialMargin(
				order.IsBuy,
				spot.InexactFloat64(),
				strike.InexactFloat64(),
				sigma,
				tauAnnualized,
				interestRate.InexactFloat64(),
				order.Option.IsCall,
				minMarginPerc.InexactFloat64(),
			)
			marginUnit := decimal.NewFromFloat(marginUnit64)
			marginOrder := marginUnit.Mul(order.Quantity)

			// Add this margin to the total 
			margin = margin.Add(marginOrder)
		}
	}
	return margin, nil
}

// GetInitialMarginPositions - Get initial margin of all matched orders (positions)
// @dev Requires a smart contract call
func (book *Book) GetInitialMarginPositions(
	instance *margin.Margin,
	user common.Address,
	decimals uint8,
) (decimal.Decimal, error) {
	// The last argument being true means we compute initial margin
	marginBn, err := instance.GetMargin(nil, user, true)
	if err != nil {
		return decimal.Zero, err
	}
	margin := decimal.NewFromBigInt(marginBn, -int32(decimals))
	return margin, nil
}

// GetMaintenanceMarginPositions - Get maintenance margin of all matched orders (positions)
// @dev Requires a smart contract call
func (book *Book) GetMaintenanceMarginPositions(
	instance *margin.Margin,
	user common.Address,
	decimals uint8,
) (decimal.Decimal, error) {
	// The last argument being false means we compute maintenance margin
	marginBn, err := instance.GetMargin(nil, user, false)
	if err != nil {
		return decimal.Zero, err
	}
	margin := decimal.NewFromBigInt(marginBn, -int32(decimals))
	return margin, nil
}

// GetUnrealizedPnL - Get realized profit and losses (losses only)
// @dev Requires a smart contract call
func (book *Book) GetUnrealizedPnL(
	instance *margin.Margin,
	user common.Address,
	decimals uint8,
) (decimal.Decimal, error) {
	// The last argument ignores any profits aggregated by strike
	payoffBn, err := instance.GetPayoff(nil, user, true)
	if err != nil {
		return decimal.Zero, err
	}
	margin := decimal.NewFromBigInt(payoffBn, -int32(decimals))
	return margin, nil
}

// CheckMaintenanceMarginAllAccounts - Get the liquidation buffers for all active accounts
// Returns
// 	users - List of active users
// 	buffers - List of liquidation buffers for each user
// 	checks - List of whether user passes checks
// 	err - Error
func (book *Book) CheckMaintenanceMarginAllAccounts() (
	users []common.Address, 
	buffers []decimal.Decimal,
	checks []bool,
	err error,
) {
	// Create contract instance
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return
	}
	// Get all active users in round
	users = book.GetActiveUsers()
	// If no users, end early
	if len(users) == 0 {
		return
	}
	// Check the margin of all users 
	results, err := instance.CheckMarginBatch(nil, users, false)
	if err != nil {
		return
	}
	checks = results.Satisfieds
	
	decimals := book.MarginContract.CollateralDecimals
	// Convert differences from big ints to decimals
	for i := 0; i < len(results.Diffs); i++ {
		userBuffer := decimal.NewFromBigInt(results.Diffs[i], -int32(decimals))
		buffers = append(buffers, userBuffer)
	}
	return
}

// CheckMaintenanceMargin - Get the liquidation buffers for a single user
// Returns
// 	buffer - Liquidation buffers for each user
// 	check - Whether user passes checks
// 	err - Error
func (book *Book) CheckMaintenanceMargin(user common.Address) (
	buffer decimal.Decimal,
	check bool,
	err error,
) {
	// Create contract instance
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return
	}
	bufferRaw, check, err := instance.CheckMargin(nil, user, false)
	if err != nil {
		return
	}
	decimals := book.MarginContract.CollateralDecimals
	buffer = decimal.NewFromBigInt(bufferRaw, -int32(decimals))
	return
}

/***********************************
 * Smart contract GET functions 
 ***********************************/

// GetMarginContract - Return an instance of the margin contract
func (book *Book) GetMarginContract() (*margin.Margin, *ethclient.Client, error) {
	// Create a connection to the ETH chain
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return nil, nil, err
	}
	// Find the deployed contract
	instance, err := margin.NewMargin(book.MarginContract.DeployedAddress, conn)
	if err != nil {
		return nil, nil, err
	}
	return instance, conn, nil
}

// GetOracleContract - Return an instance of the oracle contract
func (book *Book) GetOracleContract() (*oracle.Oracle, error) {
	// Create a connection to the ETH chain
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return nil, err
	}
	instance, err := oracle.NewOracle(book.OracleContract.DeployedAddress, conn)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

/***********************************
 * Smart contract POST functions 
 ***********************************/

// PostMatchedOrderOnchain - Call smart contract function to post on-chain
// Arguments:
// 	requestor - Address of the requestor
// 	isBuy - Is the request to buy or to sell
// 	isTaker - Did the request make a market order (true) or limit order (false)
// 	matchedQuantity - Amount of the order that was matched
// 	matchedOrder - The `order` object that was matched with request 
// 	order - Matched order to be posted on chain
// 	privateKeyHex - String of the master private key to post on-chain
// 	chainID - Which chain to post to (int64)
func (book *Book) PostMatchedOrderOnchain(
	requestor common.Address,
	isBuy bool,
	isTaker bool,
	matchedQuantity decimal.Decimal,
	matchedOrder *Order,
	privateKeyHex string,
	chainID int64,
) error {
	if isBuy && matchedOrder.IsBuy {
		return errors.New("PostMatchedOrderOnchain: matched two of the same side")
	}
	if matchedQuantity.GreaterThan(matchedOrder.Quantity) {
		return errors.New("PostMatchedOrderOnchain: not enough to fulfill")
	}
	var (
		buyer common.Address
		seller common.Address
		isBuyerTaker bool
		isSellerTaker bool
	)
	if isBuy {
		buyer = requestor
		seller = matchedOrder.Creator 
		isBuyerTaker = isTaker
		isSellerTaker = false
	} else {
		seller = requestor
		buyer = matchedOrder.Creator 
		isBuyerTaker = false
		isSellerTaker = isTaker
	}
	// Get contract instance
	instance, client, err := book.GetMarginContract()
	if err != nil {
		return err
	}
	// Derivate address from private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Get the nonce 
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}
	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	// Create a keyed transactor & set standard tx options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	// If too low, this will error out
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	// Create an ID for the matched order
	id := uuid.New()

	// Convert to right decimals
	price := shared.ToPriceScaleFactor(
		matchedOrder.Price,
		book.MarginContract.CollateralDecimals,
	)
	quantity := shared.ToQuantityScaleFactor(matchedQuantity)

	// Create a matched order to send
	position := margin.DerivativePositionParams{
		Id: id.String(),
		Buyer: buyer,
		Seller: seller,
		TradePrice: price,
		Quantity: quantity,
		IsCall: matchedOrder.Option.IsCall,
		StrikeLevel: uint8(matchedOrder.Option.Strike),
		Underlying: uint8(matchedOrder.Option.Underlying),
		IsBuyerMaker: isBuyerTaker,
		IsSellerMaker: isSellerTaker,
	}

	// Call smart contract function with info
	_, err = instance.AddPosition(auth, position)
	if err != nil {
		return err
	}

	return nil
}

/***********************************
 * Helper functions 
 ***********************************/

// IsValidOption - Check the option object has expected underlying
func (book *Book) IsValidOption(option *Option) error {
	if book.Underlying != option.Underlying { 
		return errors.New(
			"Book.CreateLimitOrder: wrong underlying")
	}
	return nil
}

// GetActiveUsers - Get all the active users
func (book *Book) GetActiveUsers() []common.Address {
	users := []common.Address{}
	for k, v := range book.UsersWithPositions {
		if v {
			users = append(users, k)
		}
	}
	return users
}

// GetBidManager - Retrieves the manager for bids from underlying derivative 
func (book * Book) GetBidManager(option *Option) (*Manager, bool) {
	var derivative *Derivative
	if option.IsCall {
		derivative = book.Calls
	} else {
		derivative = book.Puts
	}
	return derivative.GetManager(true, option)
}

// GetAskManager - Retrieves the manager for asks from underlying derivative 
func (book * Book) GetAskManager(option *Option) (*Manager, bool) {
	var derivative *Derivative
	if option.IsCall {
		derivative = book.Calls
	} else {
		derivative = book.Puts
	}
	return derivative.GetManager(false, option)
}

// SetBidManager - Sets the manager for a certain bid
func (book *Book) SetBidManager(option *Option, manager *Manager) {
	var derivative *Derivative
	if option.IsCall {
		derivative = book.Calls
	} else {
		derivative = book.Puts
	}
	derivative.SetManager(true, option, manager)
}

// SetAskManager - Sets the manager for a certain ask
func (book *Book) SetAskManager(option *Option, manager *Manager) {
	var derivative *Derivative
	if option.IsCall {
		derivative = book.Calls
	} else {
		derivative = book.Puts
	}
	derivative.SetManager(false, option, manager)
}

// LevelToStrike - Convert strike level to decimal price
func (book *Book) LevelToStrike(level shared.StrikeLevel) decimal.Decimal {
	return book.MarginContract.RoundStrikes[level]
}

// UpdateVolatility - Update the volatility surface
func (book *Book) UpdateVolatility(
	spot decimal.Decimal, 
	interestRate decimal.Decimal,
	order *Order,
) error {
	option := order.Option
	strike := book.LevelToStrike(option.Strike)
	tauInYear := option.TauAnnualized()
	sigma, err := controller.GetSigmaByBisection(
		spot.InexactFloat64(),
		strike.InexactFloat64(),
		tauInYear,
		interestRate.InexactFloat64(),
		order.Price.InexactFloat64(),
		option.IsCall,
		10000,
		1e-4,
	)
	if sigma == 0 {
		return errors.New("book.UpdateVolatility: sigma solved as zero")
	}
	if err != nil {
		return err
	}
	err = book.VolManager.Update(spot, strike, sigma, 0.8)
	if err != nil {
		return err
	}
	return nil
}

/***********************************
 * Admin functions 
 ***********************************/

// PruneOrders - Loop through and remove expired orders
func (book *Book) PruneOrders() error {
	var expired bool
	for uuid, ptr := range book.Orders {
		expired = false  // default is false
		order := ptr.Value.(*Order)
		expired = order.IsExpired()
		if expired {
			// Perform the actual cancelling
			book.CancelOrder(uuid)
		}
	}
	return nil
}

// PauseOrders - Pause the order book
func (book *Book) PauseOrders() error {
	if book.Paused {
		return errors.New("PauseOrders: book is already paused")
	}
	book.Paused = true
	return nil
}

// UnpauseOrders - Unpause the order book
func (book *Book) UnpauseOrders() error {
	if !book.Paused {
		return errors.New("PauseOrders: book is already unpaused")
	}
	book.Paused = false
	return nil
}

// RolloverOnchain - Call `rollover` function on-chain
func (book *Book) RolloverOnchain(
	users []common.Address,
	privateKeyHex string,
	chainID int64,
) error {
	instance, client, err := book.GetMarginContract()
	if err != nil {
		return err
	}
	// Derivate address from private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Get the nonce 
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}
	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	// Create a keyed transactor & set standard tx options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Call smart contract function with info
	_, err = instance.Rollover(auth, users)
	if err != nil {
		return err
	}

	return nil
}

// SettleOnchain - Call `settle` function on-chain
func (book *Book) SettleOnchain(privateKeyHex string, chainID int64) error {
	instance, client, err := book.GetMarginContract()
	if err != nil {
		return err
	}
	// Derivate address from private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Get the nonce 
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}
	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	// Create a keyed transactor & set standard tx options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Call smart contract function with info
	_, err = instance.Settle(auth)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMarginContractData - Call margin contract to update active expiry & strike selection
// The book has no internal storage of either
func (book *Book) UpdateMarginContractData() (error) {
	instance, _, err := book.GetMarginContract()
	if err != nil {
		return err
	}
	// Get the active expiry
	activeExpiry, err := instance.ActiveExpiry(nil)
	if err != nil {
		return err
	}
	book.MarginContract.ActiveExpiry = activeExpiry.Int64()
	// Get collateral decimals
	decimals, err := instance.GetCollateralDecimals(nil)
	if err != nil {
		return err
	}
	book.MarginContract.CollateralDecimals = decimals
	// Get strike menu
	strikes, err := instance.GetStrikes(nil, book.Underlying.UInt8())
	if err != nil {
		return err
	}
	for i := 0; i < 11; i++ {
		// Divide by 10**18
		book.MarginContract.RoundStrikes[i] = decimal.NewFromBigInt(strikes[i], -int32(decimals))
	}
	// Get the minimum order size
	minQuantity, err := instance.MinQuantityPerUnderlying(nil, book.Underlying.UInt8())
	if err != nil {
		return err
	}
	book.MarginContract.MinQuantity = decimal.NewFromBigInt(minQuantity, -4)

	return nil
}

// ResetVolManager - Hard reset for volatility manager (and the surface)
// By default, volatility surfaces are propagated to the next expiry
func (book *Book) ResetVolManager(
	initSigma float64,
	minMoneyness float64,
	maxMoneyness float64,
) (error) {
	volManager, err := controller.CreateVolatilityManager(
		book.Underlying,
		book.GetActiveExpiry(),
		initSigma,
		minMoneyness,
		maxMoneyness,
	)
	if (err != nil) {
		return err
	}
	// Overwrite volatility manager
	book.VolManager = volManager
	return nil
}
