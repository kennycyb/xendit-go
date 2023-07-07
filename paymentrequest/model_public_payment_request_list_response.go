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

// checks if the PublicPaymentRequestListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicPaymentRequestListResponse{}

// PublicPaymentRequestListResponse struct for PublicPaymentRequestListResponse
type PublicPaymentRequestListResponse struct {
	Data []PublicPaymentRequest `json:"data"`
	HasMore bool `json:"has_more"`
}

// NewPublicPaymentRequestListResponse instantiates a new PublicPaymentRequestListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPaymentRequestListResponse(data []PublicPaymentRequest, hasMore bool) *PublicPaymentRequestListResponse {
	this := PublicPaymentRequestListResponse{}
	this.Data = data
	this.HasMore = hasMore
	return &this
}

// NewPublicPaymentRequestListResponseWithDefaults instantiates a new PublicPaymentRequestListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPaymentRequestListResponseWithDefaults() *PublicPaymentRequestListResponse {
	this := PublicPaymentRequestListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PublicPaymentRequestListResponse) GetData() []PublicPaymentRequest {
	if o == nil {
		var ret []PublicPaymentRequest
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PublicPaymentRequestListResponse) GetDataOk() ([]PublicPaymentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PublicPaymentRequestListResponse) SetData(v []PublicPaymentRequest) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *PublicPaymentRequestListResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PublicPaymentRequestListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PublicPaymentRequestListResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o PublicPaymentRequestListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicPaymentRequestListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

type NullablePublicPaymentRequestListResponse struct {
	value *PublicPaymentRequestListResponse
	isSet bool
}

func (v NullablePublicPaymentRequestListResponse) Get() *PublicPaymentRequestListResponse {
	return v.value
}

func (v *NullablePublicPaymentRequestListResponse) Set(val *PublicPaymentRequestListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPaymentRequestListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPaymentRequestListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPaymentRequestListResponse(val *PublicPaymentRequestListResponse) *NullablePublicPaymentRequestListResponse {
	return &NullablePublicPaymentRequestListResponse{value: val, isSet: true}
}

func (v NullablePublicPaymentRequestListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPaymentRequestListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

