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

// checks if the PublicAuthPaymentMethod type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicAuthPaymentMethod{}

// PublicAuthPaymentMethod struct for PublicAuthPaymentMethod
type PublicAuthPaymentMethod struct {
	AuthCode string `json:"auth_code"`
}

// NewPublicAuthPaymentMethod instantiates a new PublicAuthPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAuthPaymentMethod(authCode string) *PublicAuthPaymentMethod {
	this := PublicAuthPaymentMethod{}
	this.AuthCode = authCode
	return &this
}

// NewPublicAuthPaymentMethodWithDefaults instantiates a new PublicAuthPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAuthPaymentMethodWithDefaults() *PublicAuthPaymentMethod {
	this := PublicAuthPaymentMethod{}
	return &this
}

// GetAuthCode returns the AuthCode field value
func (o *PublicAuthPaymentMethod) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *PublicAuthPaymentMethod) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *PublicAuthPaymentMethod) SetAuthCode(v string) {
	o.AuthCode = v
}

func (o PublicAuthPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAuthPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_code"] = o.AuthCode
	return toSerialize, nil
}

type NullablePublicAuthPaymentMethod struct {
	value *PublicAuthPaymentMethod
	isSet bool
}

func (v NullablePublicAuthPaymentMethod) Get() *PublicAuthPaymentMethod {
	return v.value
}

func (v *NullablePublicAuthPaymentMethod) Set(val *PublicAuthPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAuthPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAuthPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAuthPaymentMethod(val *PublicAuthPaymentMethod) *NullablePublicAuthPaymentMethod {
	return &NullablePublicAuthPaymentMethod{value: val, isSet: true}
}

func (v NullablePublicAuthPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAuthPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

