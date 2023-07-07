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

// checks if the InternalDirectBankTransferAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalDirectBankTransferAllOf{}

// InternalDirectBankTransferAllOf struct for InternalDirectBankTransferAllOf
type InternalDirectBankTransferAllOf struct {
	LinkedTokenId NullableString `json:"linked_token_id,omitempty"`
}

// NewInternalDirectBankTransferAllOf instantiates a new InternalDirectBankTransferAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDirectBankTransferAllOf() *InternalDirectBankTransferAllOf {
	this := InternalDirectBankTransferAllOf{}
	return &this
}

// NewInternalDirectBankTransferAllOfWithDefaults instantiates a new InternalDirectBankTransferAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDirectBankTransferAllOfWithDefaults() *InternalDirectBankTransferAllOf {
	this := InternalDirectBankTransferAllOf{}
	return &this
}

// GetLinkedTokenId returns the LinkedTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectBankTransferAllOf) GetLinkedTokenId() string {
	if o == nil || utils.IsNil(o.LinkedTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedTokenId.Get()
}

// GetLinkedTokenIdOk returns a tuple with the LinkedTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectBankTransferAllOf) GetLinkedTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedTokenId.Get(), o.LinkedTokenId.IsSet()
}

// HasLinkedTokenId returns a boolean if a field has been set.
func (o *InternalDirectBankTransferAllOf) HasLinkedTokenId() bool {
	if o != nil && o.LinkedTokenId.IsSet() {
		return true
	}

	return false
}

// SetLinkedTokenId gets a reference to the given NullableString and assigns it to the LinkedTokenId field.
func (o *InternalDirectBankTransferAllOf) SetLinkedTokenId(v string) {
	o.LinkedTokenId.Set(&v)
}
// SetLinkedTokenIdNil sets the value for LinkedTokenId to be an explicit nil
func (o *InternalDirectBankTransferAllOf) SetLinkedTokenIdNil() {
	o.LinkedTokenId.Set(nil)
}

// UnsetLinkedTokenId ensures that no value is present for LinkedTokenId, not even an explicit nil
func (o *InternalDirectBankTransferAllOf) UnsetLinkedTokenId() {
	o.LinkedTokenId.Unset()
}

func (o InternalDirectBankTransferAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDirectBankTransferAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkedTokenId.IsSet() {
		toSerialize["linked_token_id"] = o.LinkedTokenId.Get()
	}
	return toSerialize, nil
}

type NullableInternalDirectBankTransferAllOf struct {
	value *InternalDirectBankTransferAllOf
	isSet bool
}

func (v NullableInternalDirectBankTransferAllOf) Get() *InternalDirectBankTransferAllOf {
	return v.value
}

func (v *NullableInternalDirectBankTransferAllOf) Set(val *InternalDirectBankTransferAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDirectBankTransferAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDirectBankTransferAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDirectBankTransferAllOf(val *InternalDirectBankTransferAllOf) *NullableInternalDirectBankTransferAllOf {
	return &NullableInternalDirectBankTransferAllOf{value: val, isSet: true}
}

func (v NullableInternalDirectBankTransferAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDirectBankTransferAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

