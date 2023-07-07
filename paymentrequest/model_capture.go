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

// checks if the Capture type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Capture{}

// Capture struct for Capture
type Capture struct {
	Id string `json:"id"`
	PaymentRequestId string `json:"payment_request_id"`
	PaymentId string `json:"payment_id"`
	ReferenceId string `json:"reference_id"`
	Currency string `json:"currency"`
	AuthorizedAmount float64 `json:"authorized_amount"`
	CapturedAmount float64 `json:"captured_amount"`
	Status string `json:"status"`
	PaymentMethod PublicPaymentMethod `json:"payment_method"`
	FailureCode NullableString `json:"failure_code"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// NewCapture instantiates a new Capture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapture(id string, paymentRequestId string, paymentId string, referenceId string, currency string, authorizedAmount float64, capturedAmount float64, status string, paymentMethod PublicPaymentMethod, failureCode NullableString, metadata map[string]interface{}, created string, updated string) *Capture {
	this := Capture{}
	this.Id = id
	this.PaymentRequestId = paymentRequestId
	this.PaymentId = paymentId
	this.ReferenceId = referenceId
	this.Currency = currency
	this.AuthorizedAmount = authorizedAmount
	this.CapturedAmount = capturedAmount
	this.Status = status
	this.PaymentMethod = paymentMethod
	this.FailureCode = failureCode
	this.Metadata = metadata
	this.Created = created
	this.Updated = updated
	return &this
}

// NewCaptureWithDefaults instantiates a new Capture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureWithDefaults() *Capture {
	this := Capture{}
	return &this
}

// GetId returns the Id field value
func (o *Capture) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Capture) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Capture) SetId(v string) {
	o.Id = v
}

// GetPaymentRequestId returns the PaymentRequestId field value
func (o *Capture) GetPaymentRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentRequestId
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentRequestId, true
}

// SetPaymentRequestId sets field value
func (o *Capture) SetPaymentRequestId(v string) {
	o.PaymentRequestId = v
}

// GetPaymentId returns the PaymentId field value
func (o *Capture) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *Capture) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *Capture) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *Capture) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetCurrency returns the Currency field value
func (o *Capture) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Capture) SetCurrency(v string) {
	o.Currency = v
}

// GetAuthorizedAmount returns the AuthorizedAmount field value
func (o *Capture) GetAuthorizedAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AuthorizedAmount
}

// GetAuthorizedAmountOk returns a tuple with the AuthorizedAmount field value
// and a boolean to check if the value has been set.
func (o *Capture) GetAuthorizedAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizedAmount, true
}

// SetAuthorizedAmount sets field value
func (o *Capture) SetAuthorizedAmount(v float64) {
	o.AuthorizedAmount = v
}

// GetCapturedAmount returns the CapturedAmount field value
func (o *Capture) GetCapturedAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CapturedAmount
}

// GetCapturedAmountOk returns a tuple with the CapturedAmount field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCapturedAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAmount, true
}

// SetCapturedAmount sets field value
func (o *Capture) SetCapturedAmount(v float64) {
	o.CapturedAmount = v
}

// GetStatus returns the Status field value
func (o *Capture) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Capture) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Capture) SetStatus(v string) {
	o.Status = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *Capture) GetPaymentMethod() PublicPaymentMethod {
	if o == nil {
		var ret PublicPaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentMethodOk() (*PublicPaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *Capture) SetPaymentMethod(v PublicPaymentMethod) {
	o.PaymentMethod = v
}

// GetFailureCode returns the FailureCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Capture) GetFailureCode() string {
	if o == nil || o.FailureCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// SetFailureCode sets field value
func (o *Capture) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Capture) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Capture) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Capture) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Capture) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Capture) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Capture) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Capture) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreated returns the Created field value
func (o *Capture) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Capture) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Capture) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Capture) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Capture) SetUpdated(v string) {
	o.Updated = v
}

func (o Capture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Capture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["payment_request_id"] = o.PaymentRequestId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["currency"] = o.Currency
	toSerialize["authorized_amount"] = o.AuthorizedAmount
	toSerialize["captured_amount"] = o.CapturedAmount
	toSerialize["status"] = o.Status
	toSerialize["payment_method"] = o.PaymentMethod
	toSerialize["failure_code"] = o.FailureCode.Get()
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

type NullableCapture struct {
	value *Capture
	isSet bool
}

func (v NullableCapture) Get() *Capture {
	return v.value
}

func (v *NullableCapture) Set(val *Capture) {
	v.value = val
	v.isSet = true
}

func (v NullableCapture) IsSet() bool {
	return v.isSet
}

func (v *NullableCapture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapture(val *Capture) *NullableCapture {
	return &NullableCapture{value: val, isSet: true}
}

func (v NullableCapture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

