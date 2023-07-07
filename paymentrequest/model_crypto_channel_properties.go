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

// checks if the CryptoChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CryptoChannelProperties{}

// CryptoChannelProperties Cryptocurrency Channel Properties
type CryptoChannelProperties struct {
	// URL where the end-customer is redirected if the authorization is successful
	SuccessReturnUrl NullableString `json:"success_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization failed
	FailureReturnUrl NullableString `json:"failure_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization cancelled
	CancelReturnUrl NullableString `json:"cancel_return_url,omitempty"`
}

// NewCryptoChannelProperties instantiates a new CryptoChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoChannelProperties() *CryptoChannelProperties {
	this := CryptoChannelProperties{}
	return &this
}

// NewCryptoChannelPropertiesWithDefaults instantiates a new CryptoChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoChannelPropertiesWithDefaults() *CryptoChannelProperties {
	this := CryptoChannelProperties{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptoChannelProperties) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl.Get()
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptoChannelProperties) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessReturnUrl.Get(), o.SuccessReturnUrl.IsSet()
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *CryptoChannelProperties) HasSuccessReturnUrl() bool {
	if o != nil && o.SuccessReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given NullableString and assigns it to the SuccessReturnUrl field.
func (o *CryptoChannelProperties) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl.Set(&v)
}
// SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil
func (o *CryptoChannelProperties) SetSuccessReturnUrlNil() {
	o.SuccessReturnUrl.Set(nil)
}

// UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
func (o *CryptoChannelProperties) UnsetSuccessReturnUrl() {
	o.SuccessReturnUrl.Unset()
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptoChannelProperties) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl.Get()
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptoChannelProperties) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReturnUrl.Get(), o.FailureReturnUrl.IsSet()
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *CryptoChannelProperties) HasFailureReturnUrl() bool {
	if o != nil && o.FailureReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given NullableString and assigns it to the FailureReturnUrl field.
func (o *CryptoChannelProperties) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl.Set(&v)
}
// SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil
func (o *CryptoChannelProperties) SetFailureReturnUrlNil() {
	o.FailureReturnUrl.Set(nil)
}

// UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
func (o *CryptoChannelProperties) UnsetFailureReturnUrl() {
	o.FailureReturnUrl.Unset()
}

// GetCancelReturnUrl returns the CancelReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptoChannelProperties) GetCancelReturnUrl() string {
	if o == nil || utils.IsNil(o.CancelReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CancelReturnUrl.Get()
}

// GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptoChannelProperties) GetCancelReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelReturnUrl.Get(), o.CancelReturnUrl.IsSet()
}

// HasCancelReturnUrl returns a boolean if a field has been set.
func (o *CryptoChannelProperties) HasCancelReturnUrl() bool {
	if o != nil && o.CancelReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetCancelReturnUrl gets a reference to the given NullableString and assigns it to the CancelReturnUrl field.
func (o *CryptoChannelProperties) SetCancelReturnUrl(v string) {
	o.CancelReturnUrl.Set(&v)
}
// SetCancelReturnUrlNil sets the value for CancelReturnUrl to be an explicit nil
func (o *CryptoChannelProperties) SetCancelReturnUrlNil() {
	o.CancelReturnUrl.Set(nil)
}

// UnsetCancelReturnUrl ensures that no value is present for CancelReturnUrl, not even an explicit nil
func (o *CryptoChannelProperties) UnsetCancelReturnUrl() {
	o.CancelReturnUrl.Unset()
}

func (o CryptoChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SuccessReturnUrl.IsSet() {
		toSerialize["success_return_url"] = o.SuccessReturnUrl.Get()
	}
	if o.FailureReturnUrl.IsSet() {
		toSerialize["failure_return_url"] = o.FailureReturnUrl.Get()
	}
	if o.CancelReturnUrl.IsSet() {
		toSerialize["cancel_return_url"] = o.CancelReturnUrl.Get()
	}
	return toSerialize, nil
}

type NullableCryptoChannelProperties struct {
	value *CryptoChannelProperties
	isSet bool
}

func (v NullableCryptoChannelProperties) Get() *CryptoChannelProperties {
	return v.value
}

func (v *NullableCryptoChannelProperties) Set(val *CryptoChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoChannelProperties(val *CryptoChannelProperties) *NullableCryptoChannelProperties {
	return &NullableCryptoChannelProperties{value: val, isSet: true}
}

func (v NullableCryptoChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


