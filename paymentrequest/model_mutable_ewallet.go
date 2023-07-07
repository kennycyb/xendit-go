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

// checks if the MutableEwallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MutableEwallet{}

// MutableEwallet struct for MutableEwallet
type MutableEwallet struct {
	ChannelCode *EwalletChannelCode `json:"channel_code,omitempty"`
	ChannelProperties *EwalletChannelProperties `json:"channel_properties,omitempty"`
}

// NewMutableEwallet instantiates a new MutableEwallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutableEwallet() *MutableEwallet {
	this := MutableEwallet{}
	return &this
}

// NewMutableEwalletWithDefaults instantiates a new MutableEwallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutableEwalletWithDefaults() *MutableEwallet {
	this := MutableEwallet{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *MutableEwallet) GetChannelCode() EwalletChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret EwalletChannelCode
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutableEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *MutableEwallet) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given EwalletChannelCode and assigns it to the ChannelCode field.
func (o *MutableEwallet) SetChannelCode(v EwalletChannelCode) {
	o.ChannelCode = &v
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

func (o MutableEwallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutableEwallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
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


