package orderbook

import (
	"os"
	"testing"
	"time"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/ethereum/go-ethereum/common"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// Use this option pointer for all the tests below
var option = CreateOption(
	shared.ETH,
	shared.ATM,
	uint64(time.Now().Unix()) + 604800,
	true,
)

// Test creating a new book
func TestCreateBook(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, err := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	require.Equal(t, err, nil)
	t.Log(book)
}

// Test pinging contract data
func TestUpdateMarginContractData(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	// Check that the active expiry is not zero
	require.Equal(t, book.MarginContract.ActiveExpiry > 0, true)
	for i := 0; i < 11; i++ {
		require.Equal(t, book.MarginContract.RoundStrikes[i].GreaterThan(decimal.Zero), true)
		if i > 0 {
			require.Equal(
				t, 
				book.MarginContract.RoundStrikes[i].GreaterThan(book.MarginContract.RoundStrikes[i-1]),
				true,
			)
		}
	}
}

// Test getting an order in empty book
func TestBookGetOrderEmpty(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	_, err := book.GetOrderByID(uuid.New())
	// Should error given no order with id "test" exists
	require.Equal(t, err == nil, false)
}

// Test getting depth in empty book
func TestBookGetDepthEmpty(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	// Should error given empty
	require.Equal(t, asks == nil, true)
	require.Equal(t, bids == nil, true)
}

// Test getting price in empty book
func TestBookGetPriceEmpty(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	_, err := book.GetPrice(true, decimal.New(10, 0), option)
	require.Equal(t, err == nil, false)
}

// Test creating a single limit order
func TestBookCreateLimitOrder(t *testing.T) {
	// creator, _, _, _ := auth.CreateWallet()
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	// Place first order in book
	partial, quantityLeft, err := book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	require.Equal(t, err, nil)
	// All 100 should be leftover since first order
	require.Equal(t, quantityLeft.Equal(decimal.New(1, 0)), true)

	order, err2 := book.GetOrderByID(partial.ID)
	require.Equal(t, err2 == nil, true)
	require.Equal(t, order, partial)

	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	require.Equal(t, len(asks), 0)
	require.Equal(t, len(bids), 1)
}

// Test creating a order that is too large
func TestBookTooLargeLimitOrder(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	// Place first order in book
	_, _, err := book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1000, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	require.Equal(t, err != nil, true)
}

// Test creating a single market order (empty book)
func TestBookEmptyBookCreateMarketOrder(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	// Place market order without any to fill
	quantityLeft, err := book.CreateMarketOrder(
		creator,
		true,
		decimal.New(1, 0),
		option,
		privateKeyHex,
		31337,
	)
	// Check that none of it was filled
	require.Equal(t, quantityLeft.Equal(decimal.NewFromInt(1)), true)
	require.Equal(t, err != nil, true)
}

// Test creating a single market order against limit order
func TestBookCreateMarketOrder(t *testing.T) {
	// TEST_DEPOSITOR_ADDRESS needs to be one that has deposited
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	quantityLeft, err := book.CreateMarketOrder(
		creator,
		false,
		decimal.New(1, 0),
		option,
		privateKeyHex,
		31337,
	)
	// Check that all of it was filled
	require.Equal(t, quantityLeft.IsZero(), true)
	require.Equal(t, err, nil)
}

// Create one buy and one sell limit order (unmatched)
func TestBookCreateTwoLimitOrders(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	book.CreateLimitOrder(
		creator,
		false,
		decimal.New(1, 0),
		decimal.New(200, 0),
		option,
		privateKeyHex,
		31337,
	)
	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	require.Equal(t, len(asks), 1)
	require.Equal(t, len(bids), 1)
}

// Create one buy and one sell limit order (matched)
func TestBookCreateMatchedLimitOrders(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.CreateLimitOrder(
		creator,
		true,
		decimal.New(2, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	book.CreateLimitOrder(
		creator,
		false,
		decimal.New(1, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)

	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	require.Equal(t, len(asks), 0)
	require.Equal(t, len(bids), 1)
}

// Add a bunch of orders (of the same option)
func simulateDepth(
	book *Book, 
	option *Option,
	quantity decimal.Decimal,
) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	for i := 50; i < 100; i = i + 10 {
		book.CreateLimitOrder(
			creator,
			true,
			quantity, 
			decimal.New(int64(i), 0),
			option,
			privateKeyHex,
			31337,
		)
	}
	for i := 100; i < 150; i = i + 10 {
		book.CreateLimitOrder(
			creator,
			false,
			quantity,
			decimal.New(int64(i), 0),
			option,
			privateKeyHex,
			31337,
		)
	}
}

// Create many buy and sell limit order (unmatched)
func TestBookCreateManyLimitOrders(t *testing.T) {
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	simulateDepth(book, option, decimal.New(1, 0))
	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	require.Equal(t, len(asks), 5)
	require.Equal(t, len(bids), 5)
}

// Test book pruning - all orders should disappear
func TestPruneBook(t *testing.T) {
	// Create an expired option
	option := CreateOption(
		shared.ETH,
		shared.ATM,
		uint64(time.Now().Unix()) - 1,
		true,
	)
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	simulateDepth(book, option, decimal.New(1, 0))
	book.PruneOrders()
	asks, bids, err := book.GetDepth(option)
	require.Equal(t, err, nil)
	require.Equal(t, len(asks), 0)
	require.Equal(t, len(bids), 0)
}

// Test book pausing
func TestPauseBook(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.PauseOrders()
	_, _, err := book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0), 
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	// There should be an error
	require.Equal(t, err != nil, true)
}

// Test book unpausing
func TestUnpauseBook(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.PauseOrders()
	book.UnpauseOrders()
	_, _, err := book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0), 
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	// There should be not an error
	require.Equal(t, err, nil)
}

// Test book loading
func TestBookJSON(t *testing.T) {
	creator := common.HexToAddress(os.Getenv("TEST_DEPOSITOR_ADDRESS"))
	privateKeyHex := os.Getenv("TEST_DEPOSITOR_PRIVATE_KEY")
	marginContractAddress := common.HexToAddress(os.Getenv("MARGIN_CONTRACT"))
	oracleContractAddress := common.HexToAddress(os.Getenv("ORACLE_CONTRACT"))
	book, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	book.CreateLimitOrder(
		creator,
		true,
		decimal.New(1, 0),
		decimal.New(100, 0),
		option,
		privateKeyHex,
		31337,
	)
	data, err := json.Marshal(book)
	require.Equal(t, err, nil)
	
	recon, _ := CreateBook(shared.ETH, marginContractAddress, oracleContractAddress)
	err = json.Unmarshal(data, &recon)
	require.Equal(t, err, nil)
}
