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

// checks if the InternalPaymentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalPaymentRequest{}

// InternalPaymentRequest struct for InternalPaymentRequest
type InternalPaymentRequest struct {
	Id string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	ReferenceId string `json:"reference_id"`
	BusinessId string `json:"business_id"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	Customer map[string]interface{} `json:"customer,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Country *PaymentRequestCountry `json:"country,omitempty"`
	Currency PaymentRequestCurrency `json:"currency"`
	PaymentMethod PublicPaymentMethod `json:"payment_method"`
	Description NullableString `json:"description,omitempty"`
	FailureCode NullableString `json:"failure_code,omitempty"`
	CaptureMethod NullablePaymentRequestCaptureMethod `json:"capture_method,omitempty"`
	Initiator NullablePaymentRequestInitiator `json:"initiator,omitempty"`
	CardVerificationResults NullablePaymentRequestCardVerificationResults `json:"card_verification_results,omitempty"`
	Status PaymentRequestStatus `json:"status"`
	Actions []PaymentRequestAction `json:"actions,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ShippingInformation NullablePaymentRequestShippingInformation `json:"shipping_information,omitempty"`
	Items []PaymentRequestBasketItem `json:"items,omitempty"`
	InternalMetadata map[string]interface{} `json:"internal_metadata,omitempty"`
	InstrumentId *string `json:"instrument_id,omitempty"`
}

// NewInternalPaymentRequest instantiates a new InternalPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalPaymentRequest(id string, created string, updated string, referenceId string, businessId string, currency PaymentRequestCurrency, paymentMethod PublicPaymentMethod, status PaymentRequestStatus) *InternalPaymentRequest {
	this := InternalPaymentRequest{}
	this.Id = id
	this.Created = created
	this.Updated = updated
	this.ReferenceId = referenceId
	this.BusinessId = businessId
	this.Currency = currency
	this.PaymentMethod = paymentMethod
	this.Status = status
	return &this
}

// NewInternalPaymentRequestWithDefaults instantiates a new InternalPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalPaymentRequestWithDefaults() *InternalPaymentRequest {
	this := InternalPaymentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *InternalPaymentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InternalPaymentRequest) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *InternalPaymentRequest) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *InternalPaymentRequest) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *InternalPaymentRequest) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *InternalPaymentRequest) SetUpdated(v string) {
	o.Updated = v
}

// GetReferenceId returns the ReferenceId field value
func (o *InternalPaymentRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *InternalPaymentRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetBusinessId returns the BusinessId field value
func (o *InternalPaymentRequest) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *InternalPaymentRequest) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *InternalPaymentRequest) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *InternalPaymentRequest) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *InternalPaymentRequest) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCustomer returns the Customer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetCustomer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetCustomerOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return map[string]interface{}{}, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasCustomer() bool {
	if o != nil && utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given map[string]interface{} and assigns it to the Customer field.
func (o *InternalPaymentRequest) SetCustomer(v map[string]interface{}) {
	o.Customer = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InternalPaymentRequest) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *InternalPaymentRequest) SetAmount(v float64) {
	o.Amount = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *InternalPaymentRequest) GetCountry() PaymentRequestCountry {
	if o == nil || utils.IsNil(o.Country) {
		var ret PaymentRequestCountry
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetCountryOk() (*PaymentRequestCountry, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given PaymentRequestCountry and assigns it to the Country field.
func (o *InternalPaymentRequest) SetCountry(v PaymentRequestCountry) {
	o.Country = &v
}

// GetCurrency returns the Currency field value
func (o *InternalPaymentRequest) GetCurrency() PaymentRequestCurrency {
	if o == nil {
		var ret PaymentRequestCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InternalPaymentRequest) SetCurrency(v PaymentRequestCurrency) {
	o.Currency = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *InternalPaymentRequest) GetPaymentMethod() PublicPaymentMethod {
	if o == nil {
		var ret PublicPaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetPaymentMethodOk() (*PublicPaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *InternalPaymentRequest) SetPaymentMethod(v PublicPaymentMethod) {
	o.PaymentMethod = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InternalPaymentRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InternalPaymentRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InternalPaymentRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode.Get()) {
		var ret string
		return ret
	}
	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// HasFailureCode returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasFailureCode() bool {
	if o != nil && o.FailureCode.IsSet() {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given NullableString and assigns it to the FailureCode field.
func (o *InternalPaymentRequest) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}
// SetFailureCodeNil sets the value for FailureCode to be an explicit nil
func (o *InternalPaymentRequest) SetFailureCodeNil() {
	o.FailureCode.Set(nil)
}

// UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
func (o *InternalPaymentRequest) UnsetFailureCode() {
	o.FailureCode.Unset()
}

// GetCaptureMethod returns the CaptureMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod {
	if o == nil || utils.IsNil(o.CaptureMethod.Get()) {
		var ret PaymentRequestCaptureMethod
		return ret
	}
	return *o.CaptureMethod.Get()
}

// GetCaptureMethodOk returns a tuple with the CaptureMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaptureMethod.Get(), o.CaptureMethod.IsSet()
}

// HasCaptureMethod returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasCaptureMethod() bool {
	if o != nil && o.CaptureMethod.IsSet() {
		return true
	}

	return false
}

// SetCaptureMethod gets a reference to the given NullablePaymentRequestCaptureMethod and assigns it to the CaptureMethod field.
func (o *InternalPaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod) {
	o.CaptureMethod.Set(&v)
}
// SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil
func (o *InternalPaymentRequest) SetCaptureMethodNil() {
	o.CaptureMethod.Set(nil)
}

// UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
func (o *InternalPaymentRequest) UnsetCaptureMethod() {
	o.CaptureMethod.Unset()
}

// GetInitiator returns the Initiator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetInitiator() PaymentRequestInitiator {
	if o == nil || utils.IsNil(o.Initiator.Get()) {
		var ret PaymentRequestInitiator
		return ret
	}
	return *o.Initiator.Get()
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Initiator.Get(), o.Initiator.IsSet()
}

// HasInitiator returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasInitiator() bool {
	if o != nil && o.Initiator.IsSet() {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given NullablePaymentRequestInitiator and assigns it to the Initiator field.
func (o *InternalPaymentRequest) SetInitiator(v PaymentRequestInitiator) {
	o.Initiator.Set(&v)
}
// SetInitiatorNil sets the value for Initiator to be an explicit nil
func (o *InternalPaymentRequest) SetInitiatorNil() {
	o.Initiator.Set(nil)
}

// UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
func (o *InternalPaymentRequest) UnsetInitiator() {
	o.Initiator.Unset()
}

// GetCardVerificationResults returns the CardVerificationResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetCardVerificationResults() PaymentRequestCardVerificationResults {
	if o == nil || utils.IsNil(o.CardVerificationResults.Get()) {
		var ret PaymentRequestCardVerificationResults
		return ret
	}
	return *o.CardVerificationResults.Get()
}

// GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetCardVerificationResultsOk() (*PaymentRequestCardVerificationResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardVerificationResults.Get(), o.CardVerificationResults.IsSet()
}

// HasCardVerificationResults returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasCardVerificationResults() bool {
	if o != nil && o.CardVerificationResults.IsSet() {
		return true
	}

	return false
}

// SetCardVerificationResults gets a reference to the given NullablePaymentRequestCardVerificationResults and assigns it to the CardVerificationResults field.
func (o *InternalPaymentRequest) SetCardVerificationResults(v PaymentRequestCardVerificationResults) {
	o.CardVerificationResults.Set(&v)
}
// SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil
func (o *InternalPaymentRequest) SetCardVerificationResultsNil() {
	o.CardVerificationResults.Set(nil)
}

// UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
func (o *InternalPaymentRequest) UnsetCardVerificationResults() {
	o.CardVerificationResults.Unset()
}

// GetStatus returns the Status field value
func (o *InternalPaymentRequest) GetStatus() PaymentRequestStatus {
	if o == nil {
		var ret PaymentRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InternalPaymentRequest) SetStatus(v PaymentRequestStatus) {
	o.Status = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *InternalPaymentRequest) GetActions() []PaymentRequestAction {
	if o == nil || utils.IsNil(o.Actions) {
		var ret []PaymentRequestAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetActionsOk() ([]PaymentRequestAction, bool) {
	if o == nil || utils.IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasActions() bool {
	if o != nil && !utils.IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []PaymentRequestAction and assigns it to the Actions field.
func (o *InternalPaymentRequest) SetActions(v []PaymentRequestAction) {
	o.Actions = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InternalPaymentRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetShippingInformation returns the ShippingInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetShippingInformation() PaymentRequestShippingInformation {
	if o == nil || utils.IsNil(o.ShippingInformation.Get()) {
		var ret PaymentRequestShippingInformation
		return ret
	}
	return *o.ShippingInformation.Get()
}

// GetShippingInformationOk returns a tuple with the ShippingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingInformation.Get(), o.ShippingInformation.IsSet()
}

// HasShippingInformation returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasShippingInformation() bool {
	if o != nil && o.ShippingInformation.IsSet() {
		return true
	}

	return false
}

// SetShippingInformation gets a reference to the given NullablePaymentRequestShippingInformation and assigns it to the ShippingInformation field.
func (o *InternalPaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation) {
	o.ShippingInformation.Set(&v)
}
// SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil
func (o *InternalPaymentRequest) SetShippingInformationNil() {
	o.ShippingInformation.Set(nil)
}

// UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
func (o *InternalPaymentRequest) UnsetShippingInformation() {
	o.ShippingInformation.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetItems() []PaymentRequestBasketItem {
	if o == nil {
		var ret []PaymentRequestBasketItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetItemsOk() ([]PaymentRequestBasketItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasItems() bool {
	if o != nil && utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentRequestBasketItem and assigns it to the Items field.
func (o *InternalPaymentRequest) SetItems(v []PaymentRequestBasketItem) {
	o.Items = v
}

// GetInternalMetadata returns the InternalMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalPaymentRequest) GetInternalMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InternalMetadata
}

// GetInternalMetadataOk returns a tuple with the InternalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalPaymentRequest) GetInternalMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.InternalMetadata) {
		return map[string]interface{}{}, false
	}
	return o.InternalMetadata, true
}

// HasInternalMetadata returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasInternalMetadata() bool {
	if o != nil && utils.IsNil(o.InternalMetadata) {
		return true
	}

	return false
}

// SetInternalMetadata gets a reference to the given map[string]interface{} and assigns it to the InternalMetadata field.
func (o *InternalPaymentRequest) SetInternalMetadata(v map[string]interface{}) {
	o.InternalMetadata = v
}

// GetInstrumentId returns the InstrumentId field value if set, zero value otherwise.
func (o *InternalPaymentRequest) GetInstrumentId() string {
	if o == nil || utils.IsNil(o.InstrumentId) {
		var ret string
		return ret
	}
	return *o.InstrumentId
}

// GetInstrumentIdOk returns a tuple with the InstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPaymentRequest) GetInstrumentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InstrumentId) {
		return nil, false
	}
	return o.InstrumentId, true
}

// HasInstrumentId returns a boolean if a field has been set.
func (o *InternalPaymentRequest) HasInstrumentId() bool {
	if o != nil && !utils.IsNil(o.InstrumentId) {
		return true
	}

	return false
}

// SetInstrumentId gets a reference to the given string and assigns it to the InstrumentId field.
func (o *InternalPaymentRequest) SetInstrumentId(v string) {
	o.InstrumentId = &v
}

func (o InternalPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["business_id"] = o.BusinessId
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["currency"] = o.Currency
	toSerialize["payment_method"] = o.PaymentMethod
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.FailureCode.IsSet() {
		toSerialize["failure_code"] = o.FailureCode.Get()
	}
	if o.CaptureMethod.IsSet() {
		toSerialize["capture_method"] = o.CaptureMethod.Get()
	}
	if o.Initiator.IsSet() {
		toSerialize["initiator"] = o.Initiator.Get()
	}
	if o.CardVerificationResults.IsSet() {
		toSerialize["card_verification_results"] = o.CardVerificationResults.Get()
	}
	toSerialize["status"] = o.Status
	if !utils.IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ShippingInformation.IsSet() {
		toSerialize["shipping_information"] = o.ShippingInformation.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.InternalMetadata != nil {
		toSerialize["internal_metadata"] = o.InternalMetadata
	}
	if !utils.IsNil(o.InstrumentId) {
		toSerialize["instrument_id"] = o.InstrumentId
	}
	return toSerialize, nil
}

type NullableInternalPaymentRequest struct {
	value *InternalPaymentRequest
	isSet bool
}

func (v NullableInternalPaymentRequest) Get() *InternalPaymentRequest {
	return v.value
}

func (v *NullableInternalPaymentRequest) Set(val *InternalPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalPaymentRequest(val *InternalPaymentRequest) *NullableInternalPaymentRequest {
	return &NullableInternalPaymentRequest{value: val, isSet: true}
}

func (v NullableInternalPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

