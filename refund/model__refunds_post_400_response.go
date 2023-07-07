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

// checks if the RefundsPost400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundsPost400Response{}

// RefundsPost400Response struct for RefundsPost400Response
type RefundsPost400Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewRefundsPost400Response instantiates a new RefundsPost400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundsPost400Response() *RefundsPost400Response {
	this := RefundsPost400Response{}
	return &this
}

// NewRefundsPost400ResponseWithDefaults instantiates a new RefundsPost400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundsPost400ResponseWithDefaults() *RefundsPost400Response {
	this := RefundsPost400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RefundsPost400Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundsPost400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RefundsPost400Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RefundsPost400Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RefundsPost400Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundsPost400Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RefundsPost400Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RefundsPost400Response) SetMessage(v string) {
	o.Message = &v
}

func (o RefundsPost400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundsPost400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableRefundsPost400Response struct {
	value *RefundsPost400Response
	isSet bool
}

func (v NullableRefundsPost400Response) Get() *RefundsPost400Response {
	return v.value
}

func (v *NullableRefundsPost400Response) Set(val *RefundsPost400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundsPost400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundsPost400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundsPost400Response(val *RefundsPost400Response) *NullableRefundsPost400Response {
	return &NullableRefundsPost400Response{value: val, isSet: true}
}

func (v NullableRefundsPost400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundsPost400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


