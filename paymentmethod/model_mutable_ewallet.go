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

// checks if the MutableEwallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MutableEwallet{}

// MutableEwallet struct for MutableEwallet
type MutableEwallet struct {
	ChannelCode EwalletChannelCode `json:"channel_code"`
	ChannelProperties *EwalletChannelProperties `json:"channel_properties,omitempty"`
	Account *EwalletAccount `json:"account,omitempty"`
}

// NewMutableEwallet instantiates a new MutableEwallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutableEwallet(channelCode EwalletChannelCode) *MutableEwallet {
	this := MutableEwallet{}
	this.ChannelCode = channelCode
	return &this
}

// NewMutableEwalletWithDefaults instantiates a new MutableEwallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutableEwalletWithDefaults() *MutableEwallet {
	this := MutableEwallet{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *MutableEwallet) GetChannelCode() EwalletChannelCode {
	if o == nil {
		var ret EwalletChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *MutableEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *MutableEwallet) SetChannelCode(v EwalletChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *MutableEwallet) GetChannelProperties() EwalletChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret EwalletChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableEwallet) GetChannelPropertiesOk() (*EwalletChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *MutableEwallet) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given EwalletChannelProperties and assigns it to the ChannelProperties field.
func (o *MutableEwallet) SetChannelProperties(v EwalletChannelProperties) {
	o.ChannelProperties = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *MutableEwallet) GetAccount() EwalletAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret EwalletAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableEwallet) GetAccountOk() (*EwalletAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *MutableEwallet) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given EwalletAccount and assigns it to the Account field.
func (o *MutableEwallet) SetAccount(v EwalletAccount) {
	o.Account = &v
}

func (o MutableEwallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutableEwallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableMutableEwallet struct {
	value *MutableEwallet
	isSet bool
}

func (v NullableMutableEwallet) Get() *MutableEwallet {
	return v.value
}

func (v *NullableMutableEwallet) Set(val *MutableEwallet) {
	v.value = val
	v.isSet = true
}

func (v NullableMutableEwallet) IsSet() bool {
	return v.isSet
}

func (v *NullableMutableEwallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutableEwallet(val *MutableEwallet) *NullableMutableEwallet {
	return &NullableMutableEwallet{value: val, isSet: true}
}

func (v NullableMutableEwallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutableEwallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


