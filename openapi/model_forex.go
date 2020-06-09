/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Forex struct for Forex
type Forex struct {
	Var52WkHighInDouble float32 `json:"52WkHighInDouble,omitempty"`
	Var52WkLowInDouble float32 `json:"52WkLowInDouble,omitempty"`
	AskPriceInDouble float32 `json:"askPriceInDouble,omitempty"`
	BidPriceInDouble float32 `json:"bidPriceInDouble,omitempty"`
	ChangeInDouble float32 `json:"changeInDouble,omitempty"`
	ClosePriceInDouble float32 `json:"closePriceInDouble,omitempty"`
	Description string `json:"description,omitempty"`
	Digits float32 `json:"digits,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	ExchangeName string `json:"exchangeName,omitempty"`
	HighPriceInDouble float32 `json:"highPriceInDouble,omitempty"`
	IsTradable bool `json:"isTradable,omitempty"`
	LastPriceInDouble float32 `json:"lastPriceInDouble,omitempty"`
	LowPriceInDouble float32 `json:"lowPriceInDouble,omitempty"`
	Mark float32 `json:"mark,omitempty"`
	MarketMaker string `json:"marketMaker,omitempty"`
	OpenPriceInDouble float32 `json:"openPriceInDouble,omitempty"`
	PercentChange float32 `json:"percentChange,omitempty"`
	Product string `json:"product,omitempty"`
	SecurityStatus string `json:"securityStatus,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Tick float32 `json:"tick,omitempty"`
	TickAmount float32 `json:"tickAmount,omitempty"`
	TradingHours string `json:"tradingHours,omitempty"`
}
