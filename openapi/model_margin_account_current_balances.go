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
// MarginAccountCurrentBalances struct for MarginAccountCurrentBalances
type MarginAccountCurrentBalances struct {
	AccruedInterest float32 `json:"accruedInterest,omitempty"`
	AvailableFunds float32 `json:"availableFunds,omitempty"`
	AvailableFundsNonMarginableTrade float32 `json:"availableFundsNonMarginableTrade,omitempty"`
	BondValue float32 `json:"bondValue,omitempty"`
	BuyingPower float32 `json:"buyingPower,omitempty"`
	BuyingPowerNonMarginableTrade float32 `json:"buyingPowerNonMarginableTrade,omitempty"`
	CashBalance float32 `json:"cashBalance,omitempty"`
	CashReceipts float32 `json:"cashReceipts,omitempty"`
	DayTradingBuyingPower float32 `json:"dayTradingBuyingPower,omitempty"`
	DayTradingBuyingPowerCall float32 `json:"dayTradingBuyingPowerCall,omitempty"`
	Equity float32 `json:"equity,omitempty"`
	EquityPercentage float32 `json:"equityPercentage,omitempty"`
	IsInCall bool `json:"isInCall,omitempty"`
	LiquidationValue float32 `json:"liquidationValue,omitempty"`
	LongMarginValue float32 `json:"longMarginValue,omitempty"`
	LongMarketValue float32 `json:"longMarketValue,omitempty"`
	LongOptionMarketValue float32 `json:"longOptionMarketValue,omitempty"`
	MaintenanceCall float32 `json:"maintenanceCall,omitempty"`
	MaintenanceRequirement float32 `json:"maintenanceRequirement,omitempty"`
	MarginBalance float32 `json:"marginBalance,omitempty"`
	MoneyMarketFund float32 `json:"moneyMarketFund,omitempty"`
	MutualFundValue float32 `json:"mutualFundValue,omitempty"`
	OptionBuyingPower float32 `json:"optionBuyingPower,omitempty"`
	PendingDeposits float32 `json:"pendingDeposits,omitempty"`
	RegTCall float32 `json:"regTCall,omitempty"`
	Savings float32 `json:"savings,omitempty"`
	ShortBalance float32 `json:"shortBalance,omitempty"`
	ShortMarginValue float32 `json:"shortMarginValue,omitempty"`
	ShortMarketValue float32 `json:"shortMarketValue,omitempty"`
	ShortOptionMarketValue float32 `json:"shortOptionMarketValue,omitempty"`
	Sma float32 `json:"sma,omitempty"`
	StockBuyingPower float32 `json:"stockBuyingPower,omitempty"`
}
