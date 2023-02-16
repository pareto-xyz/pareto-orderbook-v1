package main

import (
	"encoding/json"
	"errors"
	"net"
	"bytes"
	"net/http"
	"os"
	"io"
	"strconv"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	auth "github.com/pareto-xyz/pareto-orderbook-v1/auth"
	controller "github.com/pareto-xyz/pareto-orderbook-v1/controller"
	pareto "github.com/pareto-xyz/pareto-orderbook-v1/orderbook"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
	"github.com/shopspring/decimal"
)

var ethBook, _ = pareto.CreateBook(
	0,
	common.HexToAddress(os.Getenv("MARGIN_CONTRACT")), 
	common.HexToAddress(os.Getenv("ORACLE_CONTRACT")),
)
// Books - Map from underyling to book
var Books = map[shared.Underlying]*pareto.Book { shared.ETH: ethBook }

// RateLimitManager - Create data structures for rate limiting
var RateLimitManager, _ = auth.CreateRateLimitManager(1, 1)

/**
 * Primary entry function for the API
 */
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message": "pong" })
	})
	// Public endpoints are open to everyone
	public := router.Group("/public")
	{
		/***********************************
		* Public GET requests
		***********************************/
		public.GET("/depth/:underlying", RateMiddleware(), GetDepth)
		public.GET("/expiry/:underlying", RateMiddleware(), GetActiveExpiry)
		public.GET("/sigma/:underlying", RateMiddleware(), GetSigma)

		// Public endpoints for getting orderbook prices
		price := public.Group("/price")
		{
			price.GET("/market/:underlying", RateMiddleware(), GetPrice)
			price.GET("/breakeven/:underlying", RateMiddleware(), GetBreakeven)
			price.GET("/strikes/:underlying", RateMiddleware(), GetStrikes)
			price.GET("/mark/:underlying", RateMiddleware(), GetMark)
			price.GET("/greeks/:underlying", RateMiddleware(), GetGreeks)
			price.GET("/margin/:underlying", RateMiddleware(), GetInitialMarginNewOrder)
		}

		// Public endpoints for checking margin
		margin := public.Group("/margin")
		{
			margin.GET("/check/:underlying", RateMiddleware(), CheckMargin)
			margin.GET("/check/all/:underlying", RateMiddleware(), CheckMarginAll)
			margin.POST("/liquidate/:underlying", RateMiddleware(), LiquidateOpenOrders)
		}
	}
	// Endpoints specific to a single user's margin account
	user := router.Group("/user")
	{
		/***********************************
		* Private GET requests
		* Requests "address" and "signature" to be specified in the request header
		***********************************/
		// Get information about your own account
		user.GET("/order/:underlying/:id", AuthMiddleware(), RateMiddleware(), GetOrderByID)
		user.GET("/orders/:underlying", AuthMiddleware(), RateMiddleware(), GetOrders)
		user.GET("/positions/:underlying", AuthMiddleware(), RateMiddleware(), GetPositions)
		
		account := user.Group("/account")
		{
			account.GET("/summary/:underlying", AuthMiddleware(), RateMiddleware(), GetAccountSummary)
			account.GET("/openinterest/:underlying", AuthMiddleware(), RateMiddleware(), GetOpenInterest)
		}

		/***********************************
		* Private POST requests
		* Requests "address" and "signature" to be specified in the request header
		***********************************/
		// Create new orders in orderbook
		user.POST("/create/market/:underlying", AuthMiddleware(), RateMiddleware(), CreateMarketOrder)
		user.POST("/create/limit/:underlying", AuthMiddleware(), RateMiddleware(), CreateLimitOrder)

		// Cancel new orders in orderbook
		user.POST("/cancel/:underlying/:id", AuthMiddleware(), RateMiddleware(), CancelOrderByID)
		user.POST("/cancel/batch/:underlying", AuthMiddleware(), RateMiddleware(), CancelBatch)
		user.POST("/cancel/all/:underlying", AuthMiddleware(), RateMiddleware(), CancelOrders)
	}
	// Endpoints specific to administrators
	admin := router.Group("/admin") 
	{
		/***********************************
		 * ADMIN GET requests
		 * Require special API keys in the header
		 * We do not rate limit admin endpoints
		 ***********************************/
		// For taking a snapshot of the book
		// TODO: post this to firebase
		admin.GET("/snapshot/:underlying", AdminMiddleware(), AuthMiddleware(), GetSnapshot)
		/***********************************
		 * ADMIN POST requests
		 * Require special API keys in the header
		 * We do not rate limit admin endpoints
		 ***********************************/
		admin.POST("/pause/:underlying", AdminMiddleware(), AuthMiddleware(), PauseOrders)
		admin.POST("/unpause/:underlying", AdminMiddleware(), AuthMiddleware(), UnpauseOrders)
		admin.POST("/settle/:underlying", AdminMiddleware(), AuthMiddleware(), Settle)
		// For cleaning the order book after an expiry
		// Also calls a smart contract function online
		admin.POST("/rollover/:underlying", AdminMiddleware(), AuthMiddleware(), Rollover)
		// For updating the orderbook with new strike prices
		admin.POST("/sync/:underlying", AdminMiddleware(), AuthMiddleware(), SyncWithContract)
	}
	
	// Run the server
	router.Run("localhost:8080")
}

/***********************************
 * Middleware
 ***********************************/

// AuthMiddleware - Authenticates that user has a valid wallet 
// A proper request must put the wallet address and a valid signature
// request header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch address from the request header
		addressHex := c.Request.Header.Get("pareto-ethereum-address")
		if len(addressHex) == 0 {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Fetch signature (as a hex string) from request header
		signature := c.Request.Header.Get("pareto-signature")
		if len(signature) == 0 {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Fetch timestamp from request header
		timestampRaw := c.Request.Header.Get("pareto-timestamp")
		if len(timestampRaw) == 0 {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		timestamp, err := strconv.Atoi(timestampRaw)
		if err != nil {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Get the current timestamp
		now := time.Now().Unix()
		timeDiff := now - int64(timestamp)

		// Timestamp cannot be in the future, and the signature is only valid for 30 seconds
		if (timeDiff > 30) || (timeDiff < 0) {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Create an authenticator class
		address := common.HexToAddress(addressHex)
		authenticator, _ := auth.CreateAuthenticator(address)

		// Get the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			body = []byte{}
		}
		// Replenish the body stream for a later read
		// https://stackoverflow.com/questions/62736851/go-gin-read-request-body-many-times
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		// Create a message to sign
		message, err := auth.CreateMessage(
			c.Request.Method,
			c.Request.RequestURI,
			string(body),
			int64(timestamp),
		)
		if err != nil {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		msg, err := json.Marshal(message)
		if err != nil {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Pass it to the authenticator to verify
		if !authenticator.Verify(signature, msg) {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Call next context
		c.Next()
	}
}

// AdminMiddleware - Check header contains admin X-API-KEYs
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch address from the request header
		addressHex := c.Request.Header.Get("PARETO-ETHEREUM-ADDRESS")
		// Check address is not empty and is the expected admin address
		if (addressHex != os.Getenv("PARETO_ADMIN_ADDRESS")) || (len(addressHex) == 0) {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("authentication failed"),
			)
			return
		}
		// Call next context
		c.Next()
	}
}

// RateMiddleware - Rate limit users by IP address
func RateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get ip address for current user
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status": http.StatusInternalServerError,
					"message": "Unable to get IP address",
				},
			)
			return
		}
		// Get limit
		limit := RateLimitManager.GetVisitor(ip)
		if !limit.Allow() {
			c.JSON(
				http.StatusTooManyRequests,
				gin.H{
					"status": http.StatusTooManyRequests,
					"message": "Too many requests",
				},
			)
			return
		}
		c.Next()
	}
}

/***********************************
 * Shared utility functions
 ***********************************/

// GetUnderlyingURI - URI for only underlying
type GetUnderlyingURI struct {
	Underlying string `uri:"underlying" binding:"required"` 
}

// GetUnderlyingIDURI - URI for only underlying & ID
type GetUnderlyingIDURI struct {
	Underlying string `uri:"underlying" binding:"required"`
	ID string `uri:"id" binding:"required"`
}

// ProcessUnderlying - Preprocess underlying 
// Arguments:
// 	underlyingStr (string) - Underlying in string format
// Returns:
// 	underlying (shared.Underlying) - Underlying enum
// 	error (error)
func ProcessUnderlying(underlyingStr string) (shared.Underlying, error) {
	underlyingInt, err := strconv.Atoi(underlyingStr)
	if err != nil {
		return shared.ETH, err
	}
	underlying := shared.Underlying(underlyingInt)
	return underlying, nil
}

// ProcessUnderlyingID - Preprocess underlying & Id
// Arguments:
// 	underlyingStr (string) - Underlying in string format
// 	idStr (string) - Id in string format
// Returns:
// 	underlying (shared.Underlying) - Underlying enum
// 	id (uuid.UUID) - Identifier
// 	error (error)
func ProcessUnderlyingID(
	underlyingStr string,
	idStr string,
) (shared.Underlying, uuid.UUID, error) {
	underlying, err := ProcessUnderlying(underlyingStr)
	if err != nil {
		return underlying, uuid.UUID{}, err
	}
	id , err:= uuid.Parse(idStr)
	if err != nil {
		return underlying, uuid.UUID{}, err
	}
	return underlying, id, nil
}

/***********************************
 * Public GET functions
 ***********************************/

// GetDepthQuery - Query parameters for `GetDepth` function
type GetDepthQuery struct {
	Strike shared.StrikeLevel `form:"strike" binding:"required,gte=0,lte=10"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetDepth - Get the depth of the book
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Query Parameters:
// 	strike (shared.StrikeLevel) - level for the strike (integer 0 -> 11)
// 	isCall (bool) - true (call) or false (put)
// Returns:
// 	depth (map[string][]*PriceLevel): Map with two keys, "asks" and "bids",
// 		each mapping to a list of prices and quantities
// Ex:
// 	/public/depth/0
func GetDepth(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetDepthQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// strike := shared.StrikeLevel(query.Strike)
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetDepth: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check if book is empty
	if book.IsEmpty() {
		err := errors.New("GetDepth: empty book")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Create an "query" option
	option := pareto.CreateOption(
		underlying, 
		query.Strike,
		book.GetActiveExpiry(), 
		*query.IsCall,
	)
	// Get depth with option
	asks, bids, err := book.GetDepth(option)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	result := map[string][]*pareto.PriceLevel {"asks": asks, "bids": bids}
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// GetActiveExpiry - Get active expiry
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/public/expiry/0
func GetActiveExpiry(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetActiveExpiry: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get active expiry
	expiry := book.GetActiveExpiry()
	c.JSON(http.StatusOK, gin.H{ "message": expiry })
}

// GetPriceQuery - Query for `GetPrice` function
type GetPriceQuery struct {
	IsBuy *bool `form:"isBuy" binding:"required"`
	Quantity decimal.Decimal `form:"quantity" binding:"required"`
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetPrice - Get the market price
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Query Parameters:
// 	isBuy (bool): true (buy) or false (sell)
// 	quantity (decimal.Decimal): Amount to quote
// 	strike (shared.StrikeLevel) - level for the strike (integer 0 -> 11)
// 	isCall (bool) - true (call) or false (put)
// Returns:
// 	price (decimal.Decimal): Best price to fulfill order
// Notes:
// 	Returns best bid price if buying, and best ask price with selling
// 	Factors in quantity of order requested
// Ex:
// 	/public/price/market/0
func GetPrice(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetPriceQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetPrice: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check if book is empty
	if book.IsEmpty() {
		err := errors.New("GetPrice: empty book")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Create an "query" option
	option := pareto.CreateOption(
		underlying, 
		query.Strike,
		book.GetActiveExpiry(), 
		*query.IsCall,
	)
	// Get the mark price of option
	price, err := book.GetPrice(*query.IsBuy, query.Quantity, option)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return 
	c.JSON(http.StatusOK, gin.H{ "message": price })
}

// GetStrikes - Get the strike prices for the current expiry
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Returns:
// 	strikes ([11]decimal.Decimal): 11 strike prices for the current expiry
// Ex:
// 	/public/price/strikes/0
func GetStrikes(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetStrikes: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	strikes := book.GetStrikes()
	// Return object
	c.JSON(http.StatusOK, gin.H{ "message": strikes })
}

// GetBreakevenQuery - Query for `GetBreakeven` function
type GetBreakevenQuery struct {
	IsBuy *bool `form:"isBuy" binding:"required"`
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetBreakeven - Get the breakeven price for a unit of an option
// 	BUY CALL = Strike + Call Option lowest ASK price
// 	SELL CALL = Strike + Call Options highest BID price
// 	BUY PUT = Strike - Put option lowest ASK price
// 	SELL PUT = Strike - Put option highest BID price
// URI Parameters:
// 	underlying (shared.Underlying) - Enum of the underlying token
// Query Parameters:
// 	isBuy (bool) - true (buy) or false (sell)
// 	strike (shared.StrikeLevel) - One of 11 strike levels
// 	isCall (bool) - true (call option) or false (put option)
// Notes:
// 	Assumes a single unit of quantity
// Ex:
// 	/public/price/breakeven/0
func GetBreakeven(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetBreakevenQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetBreakeven: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check if book is empty
	if book.IsEmpty() {
		err := errors.New("GetBreakeven: empty book")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get strike price 
	strikePrice := book.LevelToStrike(query.Strike)
	// Create an "query" option
	option := pareto.CreateOption(
		underlying, 
		query.Strike,
		book.GetActiveExpiry(), 
		*query.IsCall,
	)
	// Compute price with unit quantity
	unitQuantity := decimal.NewFromFloat(1.0)
	// Get the premium for an option of the strike
	premium, err := book.GetPrice(*query.IsBuy, unitQuantity, option)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var breakevenPrice decimal.Decimal
	if *query.IsCall {
		breakevenPrice = strikePrice.Add(premium)
	} else {
		breakevenPrice = strikePrice.Sub(premium)
	}
	// Return 
	c.JSON(http.StatusOK, gin.H{ "message": breakevenPrice })
}

// GetSigmaQuery - Query parameters for `GetSigma` function
type GetSigmaQuery struct {
	IsBuy *bool `form:"isBuy" binding:"required"`
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetSigma - Compute sigma 
// URI Parameters:
// 	underlying (shared.Underlying) - Enum of the underlying token
// Query Parameters:
// 	isBuy (bool) - true (buy) or false (sell)
// 	strike (shared.StrikeLevel) - One of 11 strike levels
// 	isCall (bool) - true (call option) or false (put option)
// Notes:
// 	BUY CALL = Implied volatility of the lowest ASK price
// 	SELL CALL = Implied volatility of the highest BID price
// 	BUY PUT = Implied volatility of the lowest ASK price
// 	SELL PUT = Implied volatility of the highest BID price
// 	We do not use the volatility surface but rather solve for 
// 	sigma based on market price
// Ex:
// 	/public/sigma/0
func GetSigma(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetSigmaQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetSigma: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get spot price
	spot, err := book.GetSpot()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	rate, err := book.GetInterestRate()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Create an "query" option
	option := pareto.CreateOption(
		underlying, 
		query.Strike,
		book.GetActiveExpiry(), 
		*query.IsCall,
	)
	// Compute premium of a single unit
	unitQuantity := decimal.NewFromFloat(1.0)
	premium, err := book.GetPrice(*query.IsBuy, unitQuantity, option)
	if err != nil {
		// Fail safe by looking up sigma in surface
		sigma := book.GetSigma(query.Strike)
		c.JSON(http.StatusOK, gin.H{ "message": sigma })
		return
	}
	// Get strike price 
	strikePrice := book.LevelToStrike(query.Strike)
	// Get annualized time-to-expiry 
	tauAnnualized := book.GetActiveAnnualizedTau()
	// Solve for sigma
	sigma, err := controller.GetSigmaByBisection(
		spot.InexactFloat64(),
		strikePrice.InexactFloat64(),
		tauAnnualized,
		rate.InexactFloat64(),
		premium.InexactFloat64(),
		*query.IsCall,
		10000,
		1e-4,
	)
	if err != nil {
		// Fail safe by looking up sigma in surface
		sigma := book.GetSigma(query.Strike)
		c.JSON(http.StatusOK, gin.H{ "message": sigma })
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": sigma })
}

// GetGreeksQuery - Query parameters for `GetGreeks` function
type GetGreeksQuery struct {
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetGreeks - Compute delta, gamma, theta, vega, rho for an option
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Query Parameters:
// 	strike (shared.StrikeLevel) - One of 11 strike levels
// 	isCall (bool) - true (call option) or false (put option)
// Notes:
// 	Uses market price, not ask/bid price
// 	Uses the current spot and interest rate
// Ex:
// 	/public/price/greeks/0
func GetGreeks(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetGreeksQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetGreeks: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Compute the greeks
	delta, gamma, theta, vega, rho, err := book.GetGreeks(*query.IsCall, query.Strike)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	result := map[string]any {
		"delta": delta,
		"gamma": gamma,
		"theta": theta,
		"vega": vega,
		"rho": rho,
	}
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// GetInitialMarginNewOrderURI - URI parameters for `GetInitialMarginNewOrder` function
type GetInitialMarginNewOrderURI struct {
	Underlying string `uri:"underlying" binding:"required"` 
}
// GetInitialMarginNewOrderQuery - Query parameters for `GetInitialMarginNewOrder` function
type GetInitialMarginNewOrderQuery struct {
	IsBuy *bool `form:"isBuy" binding:"required"`
	Quantity decimal.Decimal `form:"quantity" binding:"required"`
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetInitialMarginNewOrder - Compute IM for a new order
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Query Parameters:
// 	isBuy (bool): Buy (true) or Sell (false)
// 	quantity (decimal.Decimal): Amount for order
// 	strike (shared.StrikeLevel) - tier for the strike (integer 0 -> 11)
//  isCall (bool): true (call) or false (put)
// Ex:
// 	/public/price/margin/0
func GetInitialMarginNewOrder(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := GetInitialMarginNewOrderQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetInitialMarginNewOrder: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Create a temporary option
	option := pareto.CreateOption(
		underlying,
		query.Strike,
		book.GetActiveExpiry(),
		*query.IsCall,
	)
	// Compute IM for this new order
	margin, err := book.GetInitialMarginNewOrder(*query.IsBuy, query.Quantity, option)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{ "message": margin })
}

/***********************************
 * Private GET functions
 ***********************************/

// GetOrderByID - Fetch the order from the book
// Must be owner of ID to fetch it
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// 	id (string): Identifier for the order
// Header Parameters:
// 	address (common.Address): Caller wallet address
// Ex: 
// 	/user/order/0/<id>
func GetOrderByID(c *gin.Context) {
	uri := GetUnderlyingIDURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, ID, err := ProcessUnderlyingID(uri.Underlying, uri.ID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get the book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetOrderByID: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Query for the order and error if not found
	order, err := book.GetOrderByID(ID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if order.Creator != address {
		c.AbortWithError(
			http.StatusUnauthorized,
			errors.New("GetOrderByID: not the creator of order"),
		)
	}
	// Send the order back
	c.JSON(http.StatusOK, gin.H{ "message": order })
}

// GetOrders - Fetch the open unmatched orders owned by user
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Header Parameters:
// 	address (common.Address): Caller wallet address
// Ex: 
// 	/orders/0
func GetOrders(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get underlying book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Fetch orders
	orders, err := book.GetOrders(address)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Send the order back
	c.JSON(http.StatusOK, gin.H{ "message": orders })
}

// GetPositions - Fetch the matched orders from on-chain owned by user
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Ex: 
// 	/user/positions/0
func GetPositions(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get underlying book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Fetch orders
	positions, err := book.GetPositions(address)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Send the order back
	c.JSON(http.StatusOK, gin.H{ "message": positions })
}

// GetOpenInterestQuery - Query for `GetOpenInterest` function
type GetOpenInterestQuery struct {
	IsBuy *bool `form:"isBuy" binding:"required"`
	Strike shared.StrikeLevel `form:"strike" binding:"required"`
	IsCall *bool `form:"isCall" binding:"required"`
}

// GetOpenInterest - Fetch total open interest for user of a particular contract type
// Notes:
// 	Sum of notional for matched orders (i.e. positions) of a type
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Header Parameters:
// 	address (common.Address): Caller wallet address
// Ex: 
// 	/user/account/openinterest/0
func GetOpenInterest(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get query parameters
	query := GetOpenInterestQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetOpenInterest: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Fetch all positions
	positions, err := book.GetPositions(address)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Get spot price
	spot, err := book.GetSpot()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	total := decimal.Zero
	for i := 0; i < len(positions); i++ {
		position := positions[i]
		// Check if position matches query parameters
		if ((position.Option.IsCall == *query.IsCall) && 
			(position.Option.StrikeLevel == uint8(query.Strike)) &&
			((position.Buyer == address) == *query.IsBuy)) {
			// TODO: replace 4 with # from contract? 
			quantity := decimal.NewFromBigInt(position.Quantity, -int32(4))
			total = total.Add(quantity)
		}
	}
	// Multiply total quantity by spot for notional
	notional := total.Mul(spot)
	// Send the order back
	c.JSON(http.StatusOK, gin.H{ "message": notional })
}

// GetAccountSummary - Get various account info
// Notes:
// 	Account Balance - GET USDC in an individual’s margin account 
// 	Available Balance - Balance + uPnL - IM(orders) - IM(positions)
// 	Unrealized P&L - Netted P&L over positions
// 	Total Maintainence Margin - Sum of MM for all positions
// 	Total Initial Margin - Sum of IM for all orders
// 	Liquidation Buffer - Account balance - total MM
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Header Parameters:
// 	address (common.Address): Caller wallet address
// Ex: 
// 	/user/account/summary/0
func GetAccountSummary(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetAccount: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get account info
	accountBal, upnl, orderIM, positionIM, positionMM, availableBal, liquidationBuffer, err := book.GetAccount(address)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	result := map[string]any {
		"accountBalance": accountBal,
		"availableBalance": availableBal,
		"liquidationBuffer": liquidationBuffer,
		"unrealizedPnL": upnl,
		"orderIM": orderIM,
		"positionIM": positionIM,
		"positionMM": positionMM,
	}
	// Send the result back
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// CheckMarginQuery - Query parameters for `CheckMargin` function
type CheckMarginQuery struct {
	Address string `form:"address" binding:"required"`
}

// CheckMargin - Check margin of a specific account
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Query Parameters:
// 	address (string): Public address to check margin for
// Ex: 
// 	/public/margin/check/0
func CheckMargin(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get query parameters
	query := CheckMarginQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	address := common.HexToAddress(query.Address)
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CheckMargin: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Create contract instance
	buffer, check, err := book.CheckMaintenanceMargin(address)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Return result
	result := map[string]any {
		"user": query.Address,
		"buffer": buffer,
		"check": check,
	}
	// Send the result back
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// CheckMarginAll - Check margin of all active accounts i.e. those with positions
// Intended for use by liquidation bots but publically accessible.
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Ex: 
// 	/public/margin/check/all/0
func CheckMarginAll(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CheckMarginAll: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check user margins
	users, buffers, checks, err := book.CheckMaintenanceMarginAllAccounts()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Return result
	result := map[string]any {
		"users": users,
		"buffers": buffers,
		"checks": checks,
	}
	// Send the result back
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// LiquidateOpenOrdersBody - Body parameters for `LiquidateOpenOrders`
type LiquidateOpenOrdersBody struct {
	Address string `json:"address" binding:"required"`
}

// LiquidateOpenOrders - Cancel all open orders for user below margin
// Notes:
// 	This is public for anyone to call. However it requires the user to fail the 
// 	margin check i.e. liquidation buffer below 0. 
// URI Parameters:
// 	underlying (shared.Underlying): Enum for the underlying token
// Body Parameters:
// 	address (string): Public address to check margin for
// Ex: 
// 	/public/margin/liquidate/0
func LiquidateOpenOrders(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get body parameters
	body := LiquidateOpenOrdersBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	address := common.HexToAddress(body.Address)
	// Find the right book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("LiquidateOpenOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check margin
	_, check, err := book.CheckMaintenanceMargin(address)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if !check {
		err := errors.New("LiquidateOpenOrders: user not below margin")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// If we are here, then user is below margin. Cancel all orders
	orders, err := book.GetOrders(address)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Stored uuids of canceled orders
	var cancelled []uuid.UUID
	for i := 0; i < len(orders); i++ {
		id := orders[i].ID
		_, err := book.CancelOrder(id)
		if err != nil {
			continue
		}
		cancelled = append(cancelled, id)
	}
	success := len(cancelled) == len(orders)
	result := map[string]any { "success": success, "cancelled": cancelled }
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

/***********************************
 * Private POST functions
 ***********************************/

// CreateMarketOrderBody - Body parameters for `CreateMarketOrder`
type CreateMarketOrderBody struct {
	IsBuy *bool `json:"isBuy" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" binding:"required"`
	Strike shared.StrikeLevel `json:"strike" binding:"required"`
	IsCall *bool `json:"isCall" binding:"required"` 
}

// CreateMarketOrder - Create a market order in the book
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Header Parameters:
// 	pareto-ethereum-address (string): Public address for the options creator
// Body Parameters
// 	isBuy (bool): Buy (true) or Sell (false)
// 	quantity (decimal.Decimal): Amount for order
// 	strike (shared.StrikeLevel) - tier for the strike (integer 0 -> 11)
//  isCall (bool): true (call) or false (put)
// Returns:
// 	partialMatch (boolean): If the order was fully matched (true) or only partially matched (false)
// 	quantityLeft (decimal.Decimal): Amount unfilled. Must be less than or equal to `quantity`
// Dev:
// 	Requires address and signature to be set in the request header
// Ex:
// 	/user/create/market/0
func CreateMarketOrder(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get body parameters
	body := CreateMarketOrderBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get the underlying to select a book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CreateMarketOrder: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Quit if paused
	if book.Paused {
		err := errors.New("CreateMarketOrder: book is paused")
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	// Create an options object
	option := pareto.CreateOption(
		underlying, 
		body.Strike,
		book.GetActiveExpiry(), 
		*body.IsCall,
	)
	// Round quantity to the nearest tick
	body.Quantity = body.Quantity.Round(2)
	// Create a market order
	quantityLeft, err := book.CreateMarketOrder(
		address,
		*body.IsBuy,
		body.Quantity,
		option,
		os.Getenv("PARETO_ADMIN_PRIVATE_KEY"),
		31337,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	// Return result
	result := map[string]any {
		"partialMatch": quantityLeft.Sign() == 1,
		"quantityLeft": quantityLeft,
		"quantityRequested": body.Quantity,
	}
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// CreateLimitOrderBody - Body parameters for `CreateLimitOrder`
type CreateLimitOrderBody struct {
	IsBuy *bool `json:"isBuy" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required"`
	Strike shared.StrikeLevel `json:"strike" binding:"required"`
	IsCall *bool `json:"isCall" binding:"required"` 
}

// CreateLimitOrder - Create a limit order in the book
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Header Parameters:
// 	pareto-ethereum-address (string): Public address for the options creator
// Body Parameters
// 	isBuy (bool): Buy (true) or Sell (false)
// 	quantity (decimal.Decimal): Amount for order
// 	price (decimal.Decimal): Limit price that bounds amount paid or received
// 	strike (shared.StrikeLevel) - tier for the strike (integer 0 -> 11)
//  isCall (bool): true (call) or false (put)
// Returns:
//	order (Order): Order that was created that exists in the book
// 	partialMatch (boolean): If the order was fully matched (true) or only partially matched (false)
// 	quantityLeft (decimal.Decimal): Amount unfilled. Must be less than or equal to `quantity`
// Dev:
// 	Requires address and signature to be set in the request header
// Ex:
// 	/user/create/limit/0
func CreateLimitOrder(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlyingInt, err := strconv.Atoi(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	underlying := shared.Underlying(underlyingInt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get body parameters
	body := CreateLimitOrderBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get the underlying to select a book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CreateLimitOrder: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Quit if paused
	if book.Paused {
		err := errors.New("CreateLimitOrder: book is paused")
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	// Check if the quantity is above the minimum quantity
	if body.Quantity.LessThan(book.GetMinQuantity()) {
		err := errors.New("CreateLimitOrder: quantity below minimum requirement")
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	// Round quantity to the nearest tick
	body.Quantity = body.Quantity.Round(2)
	// Round price to the nearest tick
	body.Price = body.Price.Round(2)
	// Create an options object
	option := pareto.CreateOption(
		underlying, 
		body.Strike,
		book.GetActiveExpiry(), 
		*body.IsCall,
	)
	// Create a limit order
	order, quantityLeft, err := book.CreateLimitOrder(
		address,
		*body.IsBuy,
		body.Quantity,
		body.Price,
		option,
		os.Getenv("PARETO_ADMIN_PRIVATE_KEY"),
		31337,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	result := map[string]any {
		"fulfilled": quantityLeft.Sign() == 0,
		"quantityLeft": quantityLeft,
	}
	if quantityLeft.Sign() == 1 {
		result["orderId"] = order.ID
	}
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// CancelOrderByID - Cancels an order from orderbook
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// 	id (string): Identifier for order to cancel
// Header Parameters:
// 	pareto-ethereum-address (string): Address for caller
// Dev:
// 	Requires address and signature to be set in the request header
// 	Allows cancellation even when book is paused
// Ex:
// 	/user/cancel/0/<id>
func CancelOrderByID(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingIDURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, ID, err := ProcessUnderlyingID(uri.Underlying, uri.ID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get a book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CancelOrder: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check order is user's
	order, err := book.GetOrderByID(ID)
	if err != nil {
		err := errors.New("CancelOrder: order not found")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if order.Creator != address {
		err := errors.New("CancelOrder: cannot cancel order you do not own")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Attempt to cancel order
	_, err = book.CancelOrder(ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// CancelBatchBody - Body parameters for `CancelBatch`
type CancelBatchBody struct {
	IDs []string `json:"ids" binding:"required"`
}

// CancelBatch - Cancel multiple orders from orderbook
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Header Parameters:
// 	pareto-ethereum-address (string): Address for caller
// Ex:
// 	/user/cancel/batch/0
func CancelBatch(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get a book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CancelBatch: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get body parameters
	body := CancelBatchBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Cannot have no IDs
	if len(body.IDs) == 0 {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Stored uuids of canceled orders
	var cancelled []uuid.UUID
	for i := 0; i < len(body.IDs); i++ {
		id, err := uuid.Parse(body.IDs[i])
		// Skip if invalid id
		if err != nil {
			continue
		}
		order, err := book.GetOrderByID(id)
		// Skip if order not found
		if err != nil {
			continue
		}
		// Skip if order is not owned by caller
		if order.Creator != address {
			continue
		}
		_, err = book.CancelOrder(id)
		if err != nil {
			continue
		}
		cancelled = append(cancelled, id)
	}
	success := len(cancelled) == len(body.IDs)
	result := map[string]any { "success": success, "cancelled": cancelled }
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

// CancelOrders - Cancel all orders
// Header Parameters:
// 	pareto-ethereum-address (string): Public address for the options creator
// Ex:
// 	/user/cancel/all/0
func CancelOrders(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get header address
	address := common.HexToAddress(c.Request.Header.Get("pareto-ethereum-address"))
	// Get a book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("CancelOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get ids
	orders, err := book.GetOrders(address)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Stored uuids of canceled orders
	var cancelled []uuid.UUID
	for i := 0; i < len(orders); i++ {
		id := orders[i].ID
		_, err := book.CancelOrder(id)
		if err != nil {
			continue
		}
		cancelled = append(cancelled, id)
	}
	success := len(cancelled) == len(orders)
	result := map[string]any { "success": success, "cancelled": cancelled }
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

/***********************************
 * Public Oracle functions
 ***********************************/

// GetSpot - Get the spot price for an underlying asset
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/public/price/spot/0
func GetSpot(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get the underlying book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetSpot: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get spot price
	spot, err := book.GetSpot()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{ "message": spot })
}

// GetMarkQuery - Query parameters for `GetMark` function
// Note that these are not required
type GetMarkQuery struct {
	Spot decimal.Decimal `form:"spot"`
	InterestRate decimal.Decimal `form:"rate"`
}

// GetMark - Get the mark price at a spot for a call or put using Black Scholes
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Query Parameters:
// 	spot (decimal.Decimal, optional): Spot price. If not given, current spot is used
// 	interestRate (decimal.Decimal, optional): Interest rate. If not given, current rate is used
// Notes:
// 	Exposes the spot and interest rate for the oracle-bot
// Ex:
// 	/public/price/mark/0
func GetMark(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get the underlying book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("GetMark: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get Query parameters
	query := GetMarkQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Handle case if spot is not supplied
	if query.Spot.IsZero() {
		spot, err := book.GetSpot()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		query.Spot = spot
	}
	// Handle case if interest rate is not supplied
	if query.InterestRate.IsZero() {
		rate, err := book.GetInterestRate()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		query.InterestRate = rate
	}
	// Fill out the mark prices
	var (
		callMarks [11]decimal.Decimal
		putMarks [11]decimal.Decimal
	)
	for i := 0; i < 11; i++ {
		callMarks[i] = book.GetMark(true, shared.StrikeLevel(i), query.Spot, query.InterestRate)
		putMarks[i] = book.GetMark(false, shared.StrikeLevel(i), query.Spot, query.InterestRate)
	}
	// Return result
	result := map[string]any { "call": callMarks, "put": putMarks }
	c.JSON(http.StatusOK, gin.H{ "message": result })
}

/***********************************
 * Admin functions
 ***********************************/

// Settle - Settle the book.
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/settle/0
func Settle(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("Settle: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Check if the expiry has passed
	expiry := book.GetActiveExpiry()
	now := uint64(time.Now().Unix())
	if now < expiry {
		err := errors.New("Settle: options not expired yet")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Call smart contract
	err = book.SettleOnchain(
		os.Getenv("PARETO_ADMIN_PRIVATE_KEY"),
		31337,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	book.Settled = true
	// Return success
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// Rollover - Prune orders that have exceeded the expiry. 
// Must be restricted to administrators
// Dev:
// 	Expensive operation as it has to loop through all orders
// 	Orders can be cancelled in three ways:
// 	 1. By the owner before expiry
// 	 2. By API checks within code. See `Book.ProcessOrder`
// 	 3. By admin through this pruning function
// 	Calls the rollover function on-chain
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/prune/0
func Rollover(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("Rollover: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// If the book is not settled, cannot rollover
	if !book.Settled {
		err := errors.New("Rollover: not settled yet")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Build unique list of users from map
	roundUsers := book.GetActiveUsers()
	// Attempt to prune orders (actually makes deletes in book)
	err = book.PruneOrders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Reset other properties in books
	book.UserOrderIds = map[common.Address]map[uuid.UUID]bool{}
	book.UsersWithPositions = map[common.Address]bool{}
	// Call smart contract to rollover
	err = book.RolloverOnchain(
		roundUsers,
		os.Getenv("PARETO_ADMIN_PRIVATE_KEY"),
		31337,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// PauseOrders - Pause the orderbook - orders can be cancelled but not added
// Must be restricted to adminstrators
// StatusInternalServerError
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/pause/0
func PauseOrders(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("PauseOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Attempt to pause orders
	err = book.PauseOrders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// UnpauseOrders - Unpause the orderbook - allow new orders to be made
// Must be restricted to adminstrators
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/unpause/0
func UnpauseOrders(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("UnpauseOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Attempt to unpause orders
	err = book.UnpauseOrders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// SyncWithContract - Sync book with smart contract data. This is important 
// to fetch the next expiry and the strike menu
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/sync/0
func SyncWithContract(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("UnpauseOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Attempt to update internal state
	err = book.UpdateMarginContractData()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": true })
}

// GetSnapshot - Download a snapshot of the orderbook
// URI Parameters:
// 	underlying (derivative.Underlying): Enum for the underlying token
// Ex:
// 	/admin/snapshot/0
func GetSnapshot(c *gin.Context) {
	// Get URI parameters
	uri := GetUnderlyingURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Parse URI parameters
	underlying, err := ProcessUnderlying(uri.Underlying)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Get book
	book, ok := Books[underlying]
	if !ok {
		err := errors.New("UnpauseOrders: unsupported underlying")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	snapshot, err := json.Marshal(book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Return result
	c.JSON(http.StatusOK, gin.H{ "message": snapshot })
}
