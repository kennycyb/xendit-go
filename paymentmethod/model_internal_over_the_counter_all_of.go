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

// checks if the InternalOverTheCounterAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalOverTheCounterAllOf{}

// InternalOverTheCounterAllOf struct for InternalOverTheCounterAllOf
type InternalOverTheCounterAllOf struct {
	PaymentCodeId *string `json:"payment_code_id,omitempty"`
}

// NewInternalOverTheCounterAllOf instantiates a new InternalOverTheCounterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalOverTheCounterAllOf() *InternalOverTheCounterAllOf {
	this := InternalOverTheCounterAllOf{}
	return &this
}

// NewInternalOverTheCounterAllOfWithDefaults instantiates a new InternalOverTheCounterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalOverTheCounterAllOfWithDefaults() *InternalOverTheCounterAllOf {
	this := InternalOverTheCounterAllOf{}
	return &this
}

// GetPaymentCodeId returns the PaymentCodeId field value if set, zero value otherwise.
func (o *InternalOverTheCounterAllOf) GetPaymentCodeId() string {
	if o == nil || utils.IsNil(o.PaymentCodeId) {
		var ret string
		return ret
	}
	return *o.PaymentCodeId
}

// GetPaymentCodeIdOk returns a tuple with the PaymentCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOverTheCounterAllOf) GetPaymentCodeIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentCodeId) {
		return nil, false
	}
	return o.PaymentCodeId, true
}

// HasPaymentCodeId returns a boolean if a field has been set.
func (o *InternalOverTheCounterAllOf) HasPaymentCodeId() bool {
	if o != nil && !utils.IsNil(o.PaymentCodeId) {
		return true
	}

	return false
}

// SetPaymentCodeId gets a reference to the given string and assigns it to the PaymentCodeId field.
func (o *InternalOverTheCounterAllOf) SetPaymentCodeId(v string) {
	o.PaymentCodeId = &v
}

func (o InternalOverTheCounterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalOverTheCounterAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PaymentCodeId) {
		toSerialize["payment_code_id"] = o.PaymentCodeId
	}
	return toSerialize, nil
}

type NullableInternalOverTheCounterAllOf struct {
	value *InternalOverTheCounterAllOf
	isSet bool
}

func (v NullableInternalOverTheCounterAllOf) Get() *InternalOverTheCounterAllOf {
	return v.value
}

func (v *NullableInternalOverTheCounterAllOf) Set(val *InternalOverTheCounterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalOverTheCounterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalOverTheCounterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalOverTheCounterAllOf(val *InternalOverTheCounterAllOf) *NullableInternalOverTheCounterAllOf {
	return &NullableInternalOverTheCounterAllOf{value: val, isSet: true}
}

func (v NullableInternalOverTheCounterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalOverTheCounterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


