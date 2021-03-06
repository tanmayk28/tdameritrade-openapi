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
// Order struct for Order
type Order struct {
	AccountId float32 `json:"accountId,omitempty"`
	ActivationPrice float32 `json:"activationPrice,omitempty"`
	CancelTime CashAccountCancelTime `json:"cancelTime,omitempty"`
	Cancelable bool `json:"cancelable,omitempty"`
	ChildOrderStrategies []OrderChildOrderStrategies `json:"childOrderStrategies,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
	ComplexOrderStrategyType ComplexOrderStrategyType `json:"complexOrderStrategyType,omitempty"`
	DestinationLinkName string `json:"destinationLinkName,omitempty"`
	Duration Duration `json:"duration,omitempty"`
	Editable bool `json:"editable,omitempty"`
	EnteredTime string `json:"enteredTime,omitempty"`
	FilledQuantity float32 `json:"filledQuantity,omitempty"`
	OrderActivityCollection []Execution `json:"orderActivityCollection,omitempty"`
	OrderId float32 `json:"orderId,omitempty"`
	OrderLegCollection []OrderOrderLegCollection `json:"orderLegCollection,omitempty"`
	OrderStrategyType OrderStrategyType `json:"orderStrategyType,omitempty"`
	OrderType OrderType `json:"orderType,omitempty"`
	Price float32 `json:"price,omitempty"`
	PriceLinkBasis PriceLinkBasis `json:"priceLinkBasis,omitempty"`
	PriceLinkType PriceLinkType `json:"priceLinkType,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
	ReleaseTime string `json:"releaseTime,omitempty"`
	RemainingQuantity float32 `json:"remainingQuantity,omitempty"`
	ReplacingOrderCollection []map[string]interface{} `json:"replacingOrderCollection,omitempty"`
	RequestedDestination DestinationExchange `json:"requestedDestination,omitempty"`
	Session Session `json:"session,omitempty"`
	SpecialInstruction SpecialInstruction `json:"specialInstruction,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
	StatusDescription string `json:"statusDescription,omitempty"`
	StopPrice float32 `json:"stopPrice,omitempty"`
	StopPriceLinkBasis PriceLinkBasis `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType PriceLinkType `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset float32 `json:"stopPriceOffset,omitempty"`
	StopType StopType `json:"stopType,omitempty"`
	Tag string `json:"tag,omitempty"`
	TaxLotMethod TaxLotMethod `json:"taxLotMethod,omitempty"`
}
