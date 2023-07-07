/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.79.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package paymentmethod

import (
	"encoding/json"
	
	
	utils "github.com/kennycyb/xendit-go/utils"
	
)

// checks if the CardToken type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardToken{}

// CardToken Card Token
type CardToken struct {
	Id string `json:"id"`
	PaymentMethodId string `json:"payment_method_id"`
	ReferenceId string `json:"reference_id"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	RedirectUrl NullableString `json:"redirect_url,omitempty"`
	MaskedCardNumber NullableString `json:"masked_card_number,omitempty"`
	ExpiryMonth NullableString `json:"expiry_month,omitempty"`
	ExpiryYear NullableString `json:"expiry_year,omitempty"`
	CardInfo NullableCardTokenInfo `json:"card_info,omitempty"`
	BillingDetails NullableCardBillingDetails `json:"billing_details,omitempty"`
}

// NewCardToken instantiates a new CardToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardToken(id string, paymentMethodId string, referenceId string, status string, currency string) *CardToken {
	this := CardToken{}
	this.Id = id
	this.PaymentMethodId = paymentMethodId
	this.ReferenceId = referenceId
	this.Status = status
	this.Currency = currency
	return &this
}

// NewCardTokenWithDefaults instantiates a new CardToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTokenWithDefaults() *CardToken {
	this := CardToken{}
	return &this
}

// GetId returns the Id field value
func (o *CardToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CardToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CardToken) SetId(v string) {
	o.Id = v
}

// GetPaymentMethodId returns the PaymentMethodId field value
func (o *CardToken) GetPaymentMethodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
func (o *CardToken) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodId, true
}

// SetPaymentMethodId sets field value
func (o *CardToken) SetPaymentMethodId(v string) {
	o.PaymentMethodId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *CardToken) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CardToken) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CardToken) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetStatus returns the Status field value
func (o *CardToken) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CardToken) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CardToken) SetStatus(v string) {
	o.Status = v
}

// GetCurrency returns the Currency field value
func (o *CardToken) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CardToken) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CardToken) SetCurrency(v string) {
	o.Currency = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetRedirectUrl() string {
	if o == nil || utils.IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *CardToken) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *CardToken) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *CardToken) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *CardToken) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetMaskedCardNumber returns the MaskedCardNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetMaskedCardNumber() string {
	if o == nil || utils.IsNil(o.MaskedCardNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MaskedCardNumber.Get()
}

// GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetMaskedCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedCardNumber.Get(), o.MaskedCardNumber.IsSet()
}

// HasMaskedCardNumber returns a boolean if a field has been set.
func (o *CardToken) HasMaskedCardNumber() bool {
	if o != nil && o.MaskedCardNumber.IsSet() {
		return true
	}

	return false
}

// SetMaskedCardNumber gets a reference to the given NullableString and assigns it to the MaskedCardNumber field.
func (o *CardToken) SetMaskedCardNumber(v string) {
	o.MaskedCardNumber.Set(&v)
}
// SetMaskedCardNumberNil sets the value for MaskedCardNumber to be an explicit nil
func (o *CardToken) SetMaskedCardNumberNil() {
	o.MaskedCardNumber.Set(nil)
}

// UnsetMaskedCardNumber ensures that no value is present for MaskedCardNumber, not even an explicit nil
func (o *CardToken) UnsetMaskedCardNumber() {
	o.MaskedCardNumber.Unset()
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetExpiryMonth() string {
	if o == nil || utils.IsNil(o.ExpiryMonth.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth.Get()
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetExpiryMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryMonth.Get(), o.ExpiryMonth.IsSet()
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *CardToken) HasExpiryMonth() bool {
	if o != nil && o.ExpiryMonth.IsSet() {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given NullableString and assigns it to the ExpiryMonth field.
func (o *CardToken) SetExpiryMonth(v string) {
	o.ExpiryMonth.Set(&v)
}
// SetExpiryMonthNil sets the value for ExpiryMonth to be an explicit nil
func (o *CardToken) SetExpiryMonthNil() {
	o.ExpiryMonth.Set(nil)
}

// UnsetExpiryMonth ensures that no value is present for ExpiryMonth, not even an explicit nil
func (o *CardToken) UnsetExpiryMonth() {
	o.ExpiryMonth.Unset()
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetExpiryYear() string {
	if o == nil || utils.IsNil(o.ExpiryYear.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiryYear.Get()
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetExpiryYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryYear.Get(), o.ExpiryYear.IsSet()
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *CardToken) HasExpiryYear() bool {
	if o != nil && o.ExpiryYear.IsSet() {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given NullableString and assigns it to the ExpiryYear field.
func (o *CardToken) SetExpiryYear(v string) {
	o.ExpiryYear.Set(&v)
}
// SetExpiryYearNil sets the value for ExpiryYear to be an explicit nil
func (o *CardToken) SetExpiryYearNil() {
	o.ExpiryYear.Set(nil)
}

// UnsetExpiryYear ensures that no value is present for ExpiryYear, not even an explicit nil
func (o *CardToken) UnsetExpiryYear() {
	o.ExpiryYear.Unset()
}

// GetCardInfo returns the CardInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetCardInfo() CardTokenInfo {
	if o == nil || utils.IsNil(o.CardInfo.Get()) {
		var ret CardTokenInfo
		return ret
	}
	return *o.CardInfo.Get()
}

// GetCardInfoOk returns a tuple with the CardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetCardInfoOk() (*CardTokenInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardInfo.Get(), o.CardInfo.IsSet()
}

// HasCardInfo returns a boolean if a field has been set.
func (o *CardToken) HasCardInfo() bool {
	if o != nil && o.CardInfo.IsSet() {
		return true
	}

	return false
}

// SetCardInfo gets a reference to the given NullableCardTokenInfo and assigns it to the CardInfo field.
func (o *CardToken) SetCardInfo(v CardTokenInfo) {
	o.CardInfo.Set(&v)
}
// SetCardInfoNil sets the value for CardInfo to be an explicit nil
func (o *CardToken) SetCardInfoNil() {
	o.CardInfo.Set(nil)
}

// UnsetCardInfo ensures that no value is present for CardInfo, not even an explicit nil
func (o *CardToken) UnsetCardInfo() {
	o.CardInfo.Unset()
}

// GetBillingDetails returns the BillingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardToken) GetBillingDetails() CardBillingDetails {
	if o == nil || utils.IsNil(o.BillingDetails.Get()) {
		var ret CardBillingDetails
		return ret
	}
	return *o.BillingDetails.Get()
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardToken) GetBillingDetailsOk() (*CardBillingDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingDetails.Get(), o.BillingDetails.IsSet()
}

// HasBillingDetails returns a boolean if a field has been set.
func (o *CardToken) HasBillingDetails() bool {
	if o != nil && o.BillingDetails.IsSet() {
		return true
	}

	return false
}

// SetBillingDetails gets a reference to the given NullableCardBillingDetails and assigns it to the BillingDetails field.
func (o *CardToken) SetBillingDetails(v CardBillingDetails) {
	o.BillingDetails.Set(&v)
}
// SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil
func (o *CardToken) SetBillingDetailsNil() {
	o.BillingDetails.Set(nil)
}

// UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
func (o *CardToken) UnsetBillingDetails() {
	o.BillingDetails.Unset()
}

func (o CardToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["payment_method_id"] = o.PaymentMethodId
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["status"] = o.Status
	toSerialize["currency"] = o.Currency
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	if o.MaskedCardNumber.IsSet() {
		toSerialize["masked_card_number"] = o.MaskedCardNumber.Get()
	}
	if o.ExpiryMonth.IsSet() {
		toSerialize["expiry_month"] = o.ExpiryMonth.Get()
	}
	if o.ExpiryYear.IsSet() {
		toSerialize["expiry_year"] = o.ExpiryYear.Get()
	}
	if o.CardInfo.IsSet() {
		toSerialize["card_info"] = o.CardInfo.Get()
	}
	if o.BillingDetails.IsSet() {
		toSerialize["billing_details"] = o.BillingDetails.Get()
	}
	return toSerialize, nil
}

type NullableCardToken struct {
	value *CardToken
	isSet bool
}

func (v NullableCardToken) Get() *CardToken {
	return v.value
}

func (v *NullableCardToken) Set(val *CardToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCardToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCardToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardToken(val *CardToken) *NullableCardToken {
	return &NullableCardToken{value: val, isSet: true}
}

func (v NullableCardToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


