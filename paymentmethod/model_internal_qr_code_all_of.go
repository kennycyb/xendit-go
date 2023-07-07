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

// checks if the InternalQRCodeAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalQRCodeAllOf{}

// InternalQRCodeAllOf struct for InternalQRCodeAllOf
type InternalQRCodeAllOf struct {
	QrCodeId *string `json:"qr_code_id,omitempty"`
}

// NewInternalQRCodeAllOf instantiates a new InternalQRCodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalQRCodeAllOf() *InternalQRCodeAllOf {
	this := InternalQRCodeAllOf{}
	return &this
}

// NewInternalQRCodeAllOfWithDefaults instantiates a new InternalQRCodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalQRCodeAllOfWithDefaults() *InternalQRCodeAllOf {
	this := InternalQRCodeAllOf{}
	return &this
}

// GetQrCodeId returns the QrCodeId field value if set, zero value otherwise.
func (o *InternalQRCodeAllOf) GetQrCodeId() string {
	if o == nil || utils.IsNil(o.QrCodeId) {
		var ret string
		return ret
	}
	return *o.QrCodeId
}

// GetQrCodeIdOk returns a tuple with the QrCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalQRCodeAllOf) GetQrCodeIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.QrCodeId) {
		return nil, false
	}
	return o.QrCodeId, true
}

// HasQrCodeId returns a boolean if a field has been set.
func (o *InternalQRCodeAllOf) HasQrCodeId() bool {
	if o != nil && !utils.IsNil(o.QrCodeId) {
		return true
	}

	return false
}

// SetQrCodeId gets a reference to the given string and assigns it to the QrCodeId field.
func (o *InternalQRCodeAllOf) SetQrCodeId(v string) {
	o.QrCodeId = &v
}

func (o InternalQRCodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalQRCodeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.QrCodeId) {
		toSerialize["qr_code_id"] = o.QrCodeId
	}
	return toSerialize, nil
}

type NullableInternalQRCodeAllOf struct {
	value *InternalQRCodeAllOf
	isSet bool
}

func (v NullableInternalQRCodeAllOf) Get() *InternalQRCodeAllOf {
	return v.value
}

func (v *NullableInternalQRCodeAllOf) Set(val *InternalQRCodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalQRCodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalQRCodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalQRCodeAllOf(val *InternalQRCodeAllOf) *NullableInternalQRCodeAllOf {
	return &NullableInternalQRCodeAllOf{value: val, isSet: true}
}

func (v NullableInternalQRCodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalQRCodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


