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

// checks if the MutableCrypto type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MutableCrypto{}

// MutableCrypto struct for MutableCrypto
type MutableCrypto struct {
	ChannelCode NullableCryptoChannelCode `json:"channel_code,omitempty"`
	ChannelProperties *CryptoChannelProperties `json:"channel_properties,omitempty"`
	Wallet NullableCryptoWallet `json:"wallet,omitempty"`
}

// NewMutableCrypto instantiates a new MutableCrypto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutableCrypto() *MutableCrypto {
	this := MutableCrypto{}
	return &this
}

// NewMutableCryptoWithDefaults instantiates a new MutableCrypto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutableCryptoWithDefaults() *MutableCrypto {
	this := MutableCrypto{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MutableCrypto) GetChannelCode() CryptoChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode.Get()) {
		var ret CryptoChannelCode
		return ret
	}
	return *o.ChannelCode.Get()
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MutableCrypto) GetChannelCodeOk() (*CryptoChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelCode.Get(), o.ChannelCode.IsSet()
}

// HasChannelCode returns a boolean if a field has been set.
func (o *MutableCrypto) HasChannelCode() bool {
	if o != nil && o.ChannelCode.IsSet() {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given NullableCryptoChannelCode and assigns it to the ChannelCode field.
func (o *MutableCrypto) SetChannelCode(v CryptoChannelCode) {
	o.ChannelCode.Set(&v)
}
// SetChannelCodeNil sets the value for ChannelCode to be an explicit nil
func (o *MutableCrypto) SetChannelCodeNil() {
	o.ChannelCode.Set(nil)
}

// UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
func (o *MutableCrypto) UnsetChannelCode() {
	o.ChannelCode.Unset()
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *MutableCrypto) GetChannelProperties() CryptoChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret CryptoChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableCrypto) GetChannelPropertiesOk() (*CryptoChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *MutableCrypto) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given CryptoChannelProperties and assigns it to the ChannelProperties field.
func (o *MutableCrypto) SetChannelProperties(v CryptoChannelProperties) {
	o.ChannelProperties = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MutableCrypto) GetWallet() CryptoWallet {
	if o == nil || utils.IsNil(o.Wallet.Get()) {
		var ret CryptoWallet
		return ret
	}
	return *o.Wallet.Get()
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MutableCrypto) GetWalletOk() (*CryptoWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wallet.Get(), o.Wallet.IsSet()
}

// HasWallet returns a boolean if a field has been set.
func (o *MutableCrypto) HasWallet() bool {
	if o != nil && o.Wallet.IsSet() {
		return true
	}

	return false
}

// SetWallet gets a reference to the given NullableCryptoWallet and assigns it to the Wallet field.
func (o *MutableCrypto) SetWallet(v CryptoWallet) {
	o.Wallet.Set(&v)
}
// SetWalletNil sets the value for Wallet to be an explicit nil
func (o *MutableCrypto) SetWalletNil() {
	o.Wallet.Set(nil)
}

// UnsetWallet ensures that no value is present for Wallet, not even an explicit nil
func (o *MutableCrypto) UnsetWallet() {
	o.Wallet.Unset()
}

func (o MutableCrypto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutableCrypto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelCode.IsSet() {
		toSerialize["channel_code"] = o.ChannelCode.Get()
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if o.Wallet.IsSet() {
		toSerialize["wallet"] = o.Wallet.Get()
	}
	return toSerialize, nil
}

type NullableMutableCrypto struct {
	value *MutableCrypto
	isSet bool
}

func (v NullableMutableCrypto) Get() *MutableCrypto {
	return v.value
}

func (v *NullableMutableCrypto) Set(val *MutableCrypto) {
	v.value = val
	v.isSet = true
}

func (v NullableMutableCrypto) IsSet() bool {
	return v.isSet
}

func (v *NullableMutableCrypto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutableCrypto(val *MutableCrypto) *NullableMutableCrypto {
	return &NullableMutableCrypto{value: val, isSet: true}
}

func (v NullableMutableCrypto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutableCrypto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


