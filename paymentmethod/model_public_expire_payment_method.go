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

// checks if the PublicExpirePaymentMethod type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicExpirePaymentMethod{}

// PublicExpirePaymentMethod struct for PublicExpirePaymentMethod
type PublicExpirePaymentMethod struct {
	// URL where the end customer is redirected if the unlinking authorization is successful.
	SuccessReturnUrl NullableString `json:"success_return_url,omitempty"`
	// URL where the end customer is redirected if the unlinking authorization is failed.
	FailureReturnUrl NullableString `json:"failure_return_url,omitempty"`
}

// NewPublicExpirePaymentMethod instantiates a new PublicExpirePaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicExpirePaymentMethod() *PublicExpirePaymentMethod {
	this := PublicExpirePaymentMethod{}
	return &this
}

// NewPublicExpirePaymentMethodWithDefaults instantiates a new PublicExpirePaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicExpirePaymentMethodWithDefaults() *PublicExpirePaymentMethod {
	this := PublicExpirePaymentMethod{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicExpirePaymentMethod) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl.Get()
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicExpirePaymentMethod) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessReturnUrl.Get(), o.SuccessReturnUrl.IsSet()
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *PublicExpirePaymentMethod) HasSuccessReturnUrl() bool {
	if o != nil && o.SuccessReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given NullableString and assigns it to the SuccessReturnUrl field.
func (o *PublicExpirePaymentMethod) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl.Set(&v)
}
// SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil
func (o *PublicExpirePaymentMethod) SetSuccessReturnUrlNil() {
	o.SuccessReturnUrl.Set(nil)
}

// UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
func (o *PublicExpirePaymentMethod) UnsetSuccessReturnUrl() {
	o.SuccessReturnUrl.Unset()
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicExpirePaymentMethod) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl.Get()
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicExpirePaymentMethod) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReturnUrl.Get(), o.FailureReturnUrl.IsSet()
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *PublicExpirePaymentMethod) HasFailureReturnUrl() bool {
	if o != nil && o.FailureReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given NullableString and assigns it to the FailureReturnUrl field.
func (o *PublicExpirePaymentMethod) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl.Set(&v)
}
// SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil
func (o *PublicExpirePaymentMethod) SetFailureReturnUrlNil() {
	o.FailureReturnUrl.Set(nil)
}

// UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
func (o *PublicExpirePaymentMethod) UnsetFailureReturnUrl() {
	o.FailureReturnUrl.Unset()
}

func (o PublicExpirePaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicExpirePaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SuccessReturnUrl.IsSet() {
		toSerialize["success_return_url"] = o.SuccessReturnUrl.Get()
	}
	if o.FailureReturnUrl.IsSet() {
		toSerialize["failure_return_url"] = o.FailureReturnUrl.Get()
	}
	return toSerialize, nil
}

type NullablePublicExpirePaymentMethod struct {
	value *PublicExpirePaymentMethod
	isSet bool
}

func (v NullablePublicExpirePaymentMethod) Get() *PublicExpirePaymentMethod {
	return v.value
}

func (v *NullablePublicExpirePaymentMethod) Set(val *PublicExpirePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicExpirePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicExpirePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicExpirePaymentMethod(val *PublicExpirePaymentMethod) *NullablePublicExpirePaymentMethod {
	return &NullablePublicExpirePaymentMethod{value: val, isSet: true}
}

func (v NullablePublicExpirePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicExpirePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


