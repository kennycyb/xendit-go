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

// checks if the PaymentRequestAction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestAction{}

// PaymentRequestAction struct for PaymentRequestAction
type PaymentRequestAction struct {
	Action *string `json:"action,omitempty"`
	Method *string `json:"method,omitempty"`
	Url *string `json:"url,omitempty"`
	UrlType *string `json:"url_type,omitempty"`
}

// NewPaymentRequestAction instantiates a new PaymentRequestAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestAction() *PaymentRequestAction {
	this := PaymentRequestAction{}
	return &this
}

// NewPaymentRequestActionWithDefaults instantiates a new PaymentRequestAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestActionWithDefaults() *PaymentRequestAction {
	this := PaymentRequestAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PaymentRequestAction) GetAction() string {
	if o == nil || utils.IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetActionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PaymentRequestAction) HasAction() bool {
	if o != nil && !utils.IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PaymentRequestAction) SetAction(v string) {
	o.Action = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentRequestAction) GetMethod() string {
	if o == nil || utils.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentRequestAction) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentRequestAction) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PaymentRequestAction) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PaymentRequestAction) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PaymentRequestAction) SetUrl(v string) {
	o.Url = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *PaymentRequestAction) GetUrlType() string {
	if o == nil || utils.IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetUrlTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *PaymentRequestAction) HasUrlType() bool {
	if o != nil && !utils.IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *PaymentRequestAction) SetUrlType(v string) {
	o.UrlType = &v
}

func (o PaymentRequestAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.UrlType) {
		toSerialize["url_type"] = o.UrlType
	}
	return toSerialize, nil
}

type NullablePaymentRequestAction struct {
	value *PaymentRequestAction
	isSet bool
}

func (v NullablePaymentRequestAction) Get() *PaymentRequestAction {
	return v.value
}

func (v *NullablePaymentRequestAction) Set(val *PaymentRequestAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestAction(val *PaymentRequestAction) *NullablePaymentRequestAction {
	return &NullablePaymentRequestAction{value: val, isSet: true}
}

func (v NullablePaymentRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


