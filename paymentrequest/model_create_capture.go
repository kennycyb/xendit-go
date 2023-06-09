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

// checks if the CreateCapture type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateCapture{}

// CreateCapture struct for CreateCapture
type CreateCapture struct {
	ReferenceId NullableString `json:"reference_id,omitempty"`
	CaptureAmount float64 `json:"capture_amount"`
}

// NewCreateCapture instantiates a new CreateCapture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCapture(captureAmount float64) *CreateCapture {
	this := CreateCapture{}
	this.CaptureAmount = captureAmount
	return &this
}

// NewCreateCaptureWithDefaults instantiates a new CreateCapture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaptureWithDefaults() *CreateCapture {
	this := CreateCapture{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCapture) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCapture) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *CreateCapture) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *CreateCapture) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *CreateCapture) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *CreateCapture) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetCaptureAmount returns the CaptureAmount field value
func (o *CreateCapture) GetCaptureAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CaptureAmount
}

// GetCaptureAmountOk returns a tuple with the CaptureAmount field value
// and a boolean to check if the value has been set.
func (o *CreateCapture) GetCaptureAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptureAmount, true
}

// SetCaptureAmount sets field value
func (o *CreateCapture) SetCaptureAmount(v float64) {
	o.CaptureAmount = v
}

func (o CreateCapture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCapture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	toSerialize["capture_amount"] = o.CaptureAmount
	return toSerialize, nil
}

type NullableCreateCapture struct {
	value *CreateCapture
	isSet bool
}

func (v NullableCreateCapture) Get() *CreateCapture {
	return v.value
}

func (v *NullableCreateCapture) Set(val *CreateCapture) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCapture) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCapture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCapture(val *CreateCapture) *NullableCreateCapture {
	return &NullableCreateCapture{value: val, isSet: true}
}

func (v NullableCreateCapture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCapture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


