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

// checks if the PaymentRequestCardVerificationResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestCardVerificationResults{}

// PaymentRequestCardVerificationResults struct for PaymentRequestCardVerificationResults
type PaymentRequestCardVerificationResults struct {
	ThreeDSecure PaymentRequestCardVerificationResultsThreeDeeSecure `json:"three_d_secure"`
	CvvResult *string `json:"cvv_result,omitempty"`
	AddressVerificationResult *string `json:"address_verification_result,omitempty"`
}

// NewPaymentRequestCardVerificationResults instantiates a new PaymentRequestCardVerificationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestCardVerificationResults(threeDSecure PaymentRequestCardVerificationResultsThreeDeeSecure) *PaymentRequestCardVerificationResults {
	this := PaymentRequestCardVerificationResults{}
	this.ThreeDSecure = threeDSecure
	return &this
}

// NewPaymentRequestCardVerificationResultsWithDefaults instantiates a new PaymentRequestCardVerificationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestCardVerificationResultsWithDefaults() *PaymentRequestCardVerificationResults {
	this := PaymentRequestCardVerificationResults{}
	return &this
}

// GetThreeDSecure returns the ThreeDSecure field value
func (o *PaymentRequestCardVerificationResults) GetThreeDSecure() PaymentRequestCardVerificationResultsThreeDeeSecure {
	if o == nil {
		var ret PaymentRequestCardVerificationResultsThreeDeeSecure
		return ret
	}

	return o.ThreeDSecure
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResults) GetThreeDSecureOk() (*PaymentRequestCardVerificationResultsThreeDeeSecure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreeDSecure, true
}

// SetThreeDSecure sets field value
func (o *PaymentRequestCardVerificationResults) SetThreeDSecure(v PaymentRequestCardVerificationResultsThreeDeeSecure) {
	o.ThreeDSecure = v
}

// GetCvvResult returns the CvvResult field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResults) GetCvvResult() string {
	if o == nil || utils.IsNil(o.CvvResult) {
		var ret string
		return ret
	}
	return *o.CvvResult
}

// GetCvvResultOk returns a tuple with the CvvResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResults) GetCvvResultOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CvvResult) {
		return nil, false
	}
	return o.CvvResult, true
}

// HasCvvResult returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResults) HasCvvResult() bool {
	if o != nil && !utils.IsNil(o.CvvResult) {
		return true
	}

	return false
}

// SetCvvResult gets a reference to the given string and assigns it to the CvvResult field.
func (o *PaymentRequestCardVerificationResults) SetCvvResult(v string) {
	o.CvvResult = &v
}

// GetAddressVerificationResult returns the AddressVerificationResult field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResult() string {
	if o == nil || utils.IsNil(o.AddressVerificationResult) {
		var ret string
		return ret
	}
	return *o.AddressVerificationResult
}

// GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResultOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AddressVerificationResult) {
		return nil, false
	}
	return o.AddressVerificationResult, true
}

// HasAddressVerificationResult returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResults) HasAddressVerificationResult() bool {
	if o != nil && !utils.IsNil(o.AddressVerificationResult) {
		return true
	}

	return false
}

// SetAddressVerificationResult gets a reference to the given string and assigns it to the AddressVerificationResult field.
func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResult(v string) {
	o.AddressVerificationResult = &v
}

func (o PaymentRequestCardVerificationResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestCardVerificationResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["three_d_secure"] = o.ThreeDSecure
	if !utils.IsNil(o.CvvResult) {
		toSerialize["cvv_result"] = o.CvvResult
	}
	if !utils.IsNil(o.AddressVerificationResult) {
		toSerialize["address_verification_result"] = o.AddressVerificationResult
	}
	return toSerialize, nil
}

type NullablePaymentRequestCardVerificationResults struct {
	value *PaymentRequestCardVerificationResults
	isSet bool
}

func (v NullablePaymentRequestCardVerificationResults) Get() *PaymentRequestCardVerificationResults {
	return v.value
}

func (v *NullablePaymentRequestCardVerificationResults) Set(val *PaymentRequestCardVerificationResults) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCardVerificationResults) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCardVerificationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCardVerificationResults(val *PaymentRequestCardVerificationResults) *NullablePaymentRequestCardVerificationResults {
	return &NullablePaymentRequestCardVerificationResults{value: val, isSet: true}
}

func (v NullablePaymentRequestCardVerificationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCardVerificationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


