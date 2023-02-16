package controller

import (
	"io"
	"fmt"
	"errors"
	"strings"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/shopspring/decimal"
	shared "github.com/pareto-xyz/pareto-orderbook-v1/shared"
)

// ExpiryToInstrumentName - Convert an expiry timestamp to instrument name
// Arguments:
// 	expiry - Timestamp for expiry
// Returns:
// 	name - Name of the instrument on Deribit
func ExpiryToInstrumentName(
	underlying shared.Underlying,
	expiry uint64,
	strike decimal.Decimal,
	isCall bool,
) (name string) {
	date := time.Unix(int64(expiry), 0)
	asset := underlying.String()
	strikeStr := strike.Round(0).String()
	month := date.Month().String()[:3]
	year := strconv.Itoa(date.Year())[2:]
	day := date.Day()
	var typeStr string
	if isCall {
		typeStr = "C"
	} else {
		typeStr = "P"
	}
	name = fmt.Sprintf("%s-%v%v%v-%s-%s", asset, day, month, year, strikeStr, typeStr);
	name = strings.ToUpper(name)
	return
}

// InstrumentNameToID - Get Instrument ID from Name using deribits API
func InstrumentNameToID(name string) (uint64, error) {
	endpoint := fmt.Sprintf(
		"https://test.deribit.com/api/v2/public/get_instrument?instrument_name=%s",
		name,
	)
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("DeribitManager.InstrumentNameToID: status code is not 200")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var object map[string]interface{}
	err = json.Unmarshal(body, &object)	
	if err != nil {
		return 0, err
	}
	result := object["result"].(map[string]interface{})
	ID := uint64(result["instrument_id"].(float64))
	return ID, nil
}


// GetDeribitSigma Get IV From Instrument ID
// Arguments:
// 	ID - Instrument ID. See `InstrumentNameToID`
// Returns:
// 	askSigma - Implied volatility of the best ask price
// 	bidSigma - Implied volatility of the best bid price
// 	markSigma - Implied volatility of the best mark price
// 	err - Error. Nil if no error
func GetDeribitSigma(name string) (
	askSigma float64, 
	bidSigma float64, 
	markSigma float64, 
	err error,
) {
	endpoint := fmt.Sprintf(
		"https://test.deribit.com/api/v2/public/get_order_book?depth=1&instrument_name=%s",
		name,
	)
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, 0, 0, err
	}
	if resp.StatusCode != 200 {
		return 0, 0, 0, errors.New("DeribitManager.GetDeribitSigma: status code is not 200")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, err
	}
	var object map[string]interface{}
	err = json.Unmarshal(body, &object)	
	if err != nil {
		return 0, 0, 0, err
	}
	result := object["result"].(map[string]interface{})
	askSigma = result["ask_iv"].(float64) / 100.0
	bidSigma = result["bid_iv"].(float64) / 100.0
	markSigma = result["mark_iv"].(float64) / 100.0

	return
}
