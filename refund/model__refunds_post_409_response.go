/*
Refund Service

This API is used for the unified refund service

API version: 1.0.7
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package refund

import (
	"encoding/json"
	
	
	utils "github.com/kennycyb/xendit-go/utils"
	
)

// checks if the RefundsPost409Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundsPost409Response{}

// RefundsPost409Response struct for RefundsPost409Response
type RefundsPost409Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewRefundsPost409Response instantiates a new RefundsPost409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundsPost409Response() *RefundsPost409Response {
	this := RefundsPost409Response{}
	return &this
}

// NewRefundsPost409ResponseWithDefaults instantiates a new RefundsPost409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundsPost409ResponseWithDefaults() *RefundsPost409Response {
	this := RefundsPost409Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RefundsPost409Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundsPost409Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RefundsPost409Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RefundsPost409Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RefundsPost409Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundsPost409Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RefundsPost409Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RefundsPost409Response) SetMessage(v string) {
	o.Message = &v
}

func (o RefundsPost409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundsPost409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableRefundsPost409Response struct {
	value *RefundsPost409Response
	isSet bool
}

func (v NullableRefundsPost409Response) Get() *RefundsPost409Response {
	return v.value
}

func (v *NullableRefundsPost409Response) Set(val *RefundsPost409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundsPost409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundsPost409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundsPost409Response(val *RefundsPost409Response) *NullableRefundsPost409Response {
	return &NullableRefundsPost409Response{value: val, isSet: true}
}

func (v NullableRefundsPost409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundsPost409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


