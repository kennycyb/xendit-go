/*
Charge

This API is used for Charge

API version: 1.31.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package paymentrequest

import (
	"encoding/json"
	
	
	utils "github.com/kennycyb/xendit-go/utils"
	
)

// checks if the MutableVirtualAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MutableVirtualAccount{}

// MutableVirtualAccount struct for MutableVirtualAccount
type MutableVirtualAccount struct {
	Amount NullableFloat32 `json:"amount,omitempty"`
	Currency *PaymentRequestCurrency `json:"currency,omitempty"`
	ChannelCode VirtualAccountChannelCode `json:"channel_code"`
	ChannelProperties VirtualAccountChannelProperties `json:"channel_properties"`
	// Alternative display requested for the virtual account
	AlternativeDisplayTypes []string `json:"alternative_display_types,omitempty"`
}

// NewMutableVirtualAccount instantiates a new MutableVirtualAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutableVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties) *MutableVirtualAccount {
	this := MutableVirtualAccount{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	return &this
}

// NewMutableVirtualAccountWithDefaults instantiates a new MutableVirtualAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutableVirtualAccountWithDefaults() *MutableVirtualAccount {
	this := MutableVirtualAccount{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MutableVirtualAccount) GetAmount() float32 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MutableVirtualAccount) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *MutableVirtualAccount) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *MutableVirtualAccount) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *MutableVirtualAccount) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *MutableVirtualAccount) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MutableVirtualAccount) GetCurrency() PaymentRequestCurrency {
	if o == nil || utils.IsNil(o.Currency) {
		var ret PaymentRequestCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableVirtualAccount) GetCurrencyOk() (*PaymentRequestCurrency, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MutableVirtualAccount) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PaymentRequestCurrency and assigns it to the Currency field.
func (o *MutableVirtualAccount) SetCurrency(v PaymentRequestCurrency) {
	o.Currency = &v
}

// GetChannelCode returns the ChannelCode field value
func (o *MutableVirtualAccount) GetChannelCode() VirtualAccountChannelCode {
	if o == nil {
		var ret VirtualAccountChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *MutableVirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *MutableVirtualAccount) SetChannelCode(v VirtualAccountChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *MutableVirtualAccount) GetChannelProperties() VirtualAccountChannelProperties {
	if o == nil {
		var ret VirtualAccountChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *MutableVirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *MutableVirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties) {
	o.ChannelProperties = v
}

// GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field value if set, zero value otherwise.
func (o *MutableVirtualAccount) GetAlternativeDisplayTypes() []string {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		var ret []string
		return ret
	}
	return o.AlternativeDisplayTypes
}

// GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableVirtualAccount) GetAlternativeDisplayTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		return nil, false
	}
	return o.AlternativeDisplayTypes, true
}

// HasAlternativeDisplayTypes returns a boolean if a field has been set.
func (o *MutableVirtualAccount) HasAlternativeDisplayTypes() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplayTypes) {
		return true
	}

	return false
}

// SetAlternativeDisplayTypes gets a reference to the given []string and assigns it to the AlternativeDisplayTypes field.
func (o *MutableVirtualAccount) SetAlternativeDisplayTypes(v []string) {
	o.AlternativeDisplayTypes = v
}

func (o MutableVirtualAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutableVirtualAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties
	if !utils.IsNil(o.AlternativeDisplayTypes) {
		toSerialize["alternative_display_types"] = o.AlternativeDisplayTypes
	}
	return toSerialize, nil
}

type NullableMutableVirtualAccount struct {
	value *MutableVirtualAccount
	isSet bool
}

func (v NullableMutableVirtualAccount) Get() *MutableVirtualAccount {
	return v.value
}

func (v *NullableMutableVirtualAccount) Set(val *MutableVirtualAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableMutableVirtualAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableMutableVirtualAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutableVirtualAccount(val *MutableVirtualAccount) *NullableMutableVirtualAccount {
	return &NullableMutableVirtualAccount{value: val, isSet: true}
}

func (v NullableMutableVirtualAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutableVirtualAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

