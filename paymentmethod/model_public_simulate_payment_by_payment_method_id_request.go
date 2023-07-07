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

// checks if the PublicSimulatePaymentByPaymentMethodIdRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicSimulatePaymentByPaymentMethodIdRequest{}

// PublicSimulatePaymentByPaymentMethodIdRequest struct for PublicSimulatePaymentByPaymentMethodIdRequest
type PublicSimulatePaymentByPaymentMethodIdRequest struct {
	Amount *float64 `json:"amount,omitempty"`
}

// NewPublicSimulatePaymentByPaymentMethodIdRequest instantiates a new PublicSimulatePaymentByPaymentMethodIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSimulatePaymentByPaymentMethodIdRequest() *PublicSimulatePaymentByPaymentMethodIdRequest {
	this := PublicSimulatePaymentByPaymentMethodIdRequest{}
	return &this
}

// NewPublicSimulatePaymentByPaymentMethodIdRequestWithDefaults instantiates a new PublicSimulatePaymentByPaymentMethodIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSimulatePaymentByPaymentMethodIdRequestWithDefaults() *PublicSimulatePaymentByPaymentMethodIdRequest {
	this := PublicSimulatePaymentByPaymentMethodIdRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PublicSimulatePaymentByPaymentMethodIdRequest) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSimulatePaymentByPaymentMethodIdRequest) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PublicSimulatePaymentByPaymentMethodIdRequest) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PublicSimulatePaymentByPaymentMethodIdRequest) SetAmount(v float64) {
	o.Amount = &v
}

func (o PublicSimulatePaymentByPaymentMethodIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSimulatePaymentByPaymentMethodIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullablePublicSimulatePaymentByPaymentMethodIdRequest struct {
	value *PublicSimulatePaymentByPaymentMethodIdRequest
	isSet bool
}

func (v NullablePublicSimulatePaymentByPaymentMethodIdRequest) Get() *PublicSimulatePaymentByPaymentMethodIdRequest {
	return v.value
}

func (v *NullablePublicSimulatePaymentByPaymentMethodIdRequest) Set(val *PublicSimulatePaymentByPaymentMethodIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSimulatePaymentByPaymentMethodIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSimulatePaymentByPaymentMethodIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSimulatePaymentByPaymentMethodIdRequest(val *PublicSimulatePaymentByPaymentMethodIdRequest) *NullablePublicSimulatePaymentByPaymentMethodIdRequest {
	return &NullablePublicSimulatePaymentByPaymentMethodIdRequest{value: val, isSet: true}
}

func (v NullablePublicSimulatePaymentByPaymentMethodIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSimulatePaymentByPaymentMethodIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


