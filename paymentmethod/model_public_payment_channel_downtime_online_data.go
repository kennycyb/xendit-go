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

// checks if the PublicPaymentChannelDowntimeOnlineData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicPaymentChannelDowntimeOnlineData{}

// PublicPaymentChannelDowntimeOnlineData struct for PublicPaymentChannelDowntimeOnlineData
type PublicPaymentChannelDowntimeOnlineData struct {
	// Type of payment channel
	Type *string `json:"type,omitempty"`
	// Payment channel unique code
	ChannelCode *string `json:"channel_code,omitempty"`
	// Two letter country code of the payment channel
	Country *string `json:"country,omitempty"`
	// Name description of the payment channel
	ChannelName *string `json:"channel_name,omitempty"`
	// Payment channel downtime status code (Only available when downtime event)
	Status *string `json:"status,omitempty"`
}

// NewPublicPaymentChannelDowntimeOnlineData instantiates a new PublicPaymentChannelDowntimeOnlineData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPaymentChannelDowntimeOnlineData() *PublicPaymentChannelDowntimeOnlineData {
	this := PublicPaymentChannelDowntimeOnlineData{}
	return &this
}

// NewPublicPaymentChannelDowntimeOnlineDataWithDefaults instantiates a new PublicPaymentChannelDowntimeOnlineData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPaymentChannelDowntimeOnlineDataWithDefaults() *PublicPaymentChannelDowntimeOnlineData {
	this := PublicPaymentChannelDowntimeOnlineData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PublicPaymentChannelDowntimeOnlineData) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PublicPaymentChannelDowntimeOnlineData) SetType(v string) {
	o.Type = &v
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *PublicPaymentChannelDowntimeOnlineData) GetChannelCode() string {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret string
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) GetChannelCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given string and assigns it to the ChannelCode field.
func (o *PublicPaymentChannelDowntimeOnlineData) SetChannelCode(v string) {
	o.ChannelCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PublicPaymentChannelDowntimeOnlineData) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PublicPaymentChannelDowntimeOnlineData) SetCountry(v string) {
	o.Country = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *PublicPaymentChannelDowntimeOnlineData) GetChannelName() string {
	if o == nil || utils.IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) GetChannelNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) HasChannelName() bool {
	if o != nil && !utils.IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *PublicPaymentChannelDowntimeOnlineData) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicPaymentChannelDowntimeOnlineData) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicPaymentChannelDowntimeOnlineData) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PublicPaymentChannelDowntimeOnlineData) SetStatus(v string) {
	o.Status = &v
}

func (o PublicPaymentChannelDowntimeOnlineData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicPaymentChannelDowntimeOnlineData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.ChannelName) {
		toSerialize["channel_name"] = o.ChannelName
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePublicPaymentChannelDowntimeOnlineData struct {
	value *PublicPaymentChannelDowntimeOnlineData
	isSet bool
}

func (v NullablePublicPaymentChannelDowntimeOnlineData) Get() *PublicPaymentChannelDowntimeOnlineData {
	return v.value
}

func (v *NullablePublicPaymentChannelDowntimeOnlineData) Set(val *PublicPaymentChannelDowntimeOnlineData) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPaymentChannelDowntimeOnlineData) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPaymentChannelDowntimeOnlineData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPaymentChannelDowntimeOnlineData(val *PublicPaymentChannelDowntimeOnlineData) *NullablePublicPaymentChannelDowntimeOnlineData {
	return &NullablePublicPaymentChannelDowntimeOnlineData{value: val, isSet: true}
}

func (v NullablePublicPaymentChannelDowntimeOnlineData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPaymentChannelDowntimeOnlineData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


