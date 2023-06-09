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

// checks if the PaymentListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentListResponse{}

// PaymentListResponse struct for PaymentListResponse
type PaymentListResponse struct {
	Data []Payment `json:"data"`
	HasMore bool `json:"has_more"`
	Links []Link `json:"links"`
}

// NewPaymentListResponse instantiates a new PaymentListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentListResponse(data []Payment, hasMore bool, links []Link) *PaymentListResponse {
	this := PaymentListResponse{}
	this.Data = data
	this.HasMore = hasMore
	this.Links = links
	return &this
}

// NewPaymentListResponseWithDefaults instantiates a new PaymentListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentListResponseWithDefaults() *PaymentListResponse {
	this := PaymentListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentListResponse) GetData() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentListResponse) GetDataOk() ([]Payment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaymentListResponse) SetData(v []Payment) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *PaymentListResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaymentListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaymentListResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetLinks returns the Links field value
func (o *PaymentListResponse) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaymentListResponse) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaymentListResponse) SetLinks(v []Link) {
	o.Links = v
}

func (o PaymentListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullablePaymentListResponse struct {
	value *PaymentListResponse
	isSet bool
}

func (v NullablePaymentListResponse) Get() *PaymentListResponse {
	return v.value
}

func (v *NullablePaymentListResponse) Set(val *PaymentListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentListResponse(val *PaymentListResponse) *NullablePaymentListResponse {
	return &NullablePaymentListResponse{value: val, isSet: true}
}

func (v NullablePaymentListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


