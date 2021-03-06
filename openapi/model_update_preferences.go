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
// UpdatePreferences struct for UpdatePreferences
type UpdatePreferences struct {
	AuthTokenTimeout string `json:"authTokenTimeout,omitempty"`
	DefaultAdvancedToolLaunch string `json:"defaultAdvancedToolLaunch,omitempty"`
	DefaultEquityOrderDuration Duration `json:"defaultEquityOrderDuration,omitempty"`
	DefaultEquityOrderLegInstruction string `json:"defaultEquityOrderLegInstruction,omitempty"`
	DefaultEquityOrderMarketSession Session `json:"defaultEquityOrderMarketSession,omitempty"`
	DefaultEquityOrderPriceLinkType EquityPriceLinkType `json:"defaultEquityOrderPriceLinkType,omitempty"`
	DefaultEquityOrderType OrderType `json:"defaultEquityOrderType,omitempty"`
	DefaultEquityQuantity int32 `json:"defaultEquityQuantity,omitempty"`
	DirectEquityRouting bool `json:"directEquityRouting,omitempty"`
	DirectOptionsRouting bool `json:"directOptionsRouting,omitempty"`
	EquityTaxLotMethod TaxLotMethod `json:"equityTaxLotMethod,omitempty"`
	ExpressTrading bool `json:"expressTrading,omitempty"`
	MutualFundTaxLotMethod TaxLotMethod `json:"mutualFundTaxLotMethod,omitempty"`
	OptionTaxLotMethod TaxLotMethod `json:"optionTaxLotMethod,omitempty"`
}
