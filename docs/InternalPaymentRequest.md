# InternalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Created** | **string** |  | 
**Updated** | **string** |  | 
**ReferenceId** | **string** |  | 
**BusinessId** | **string** |  | 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to **map[string]interface{}** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Country** | Pointer to [**PaymentRequestCountry**](PaymentRequestCountry.md) |  | [optional] 
**Currency** | [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  | 
**PaymentMethod** | [**PublicPaymentMethod**](PublicPaymentMethod.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**CaptureMethod** | Pointer to [**NullablePaymentRequestCaptureMethod**](PaymentRequestCaptureMethod.md) |  | [optional] 
**Initiator** | Pointer to [**NullablePaymentRequestInitiator**](PaymentRequestInitiator.md) |  | [optional] 
**CardVerificationResults** | Pointer to [**NullablePaymentRequestCardVerificationResults**](PaymentRequestCardVerificationResults.md) |  | [optional] 
**Status** | [**PaymentRequestStatus**](PaymentRequestStatus.md) |  | 
**Actions** | Pointer to [**[]PaymentRequestAction**](PaymentRequestAction.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ShippingInformation** | Pointer to [**NullablePaymentRequestShippingInformation**](PaymentRequestShippingInformation.md) |  | [optional] 
**Items** | Pointer to [**[]PaymentRequestBasketItem**](PaymentRequestBasketItem.md) |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**InstrumentId** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalPaymentRequest

`func NewInternalPaymentRequest(id string, created string, updated string, referenceId string, businessId string, currency PaymentRequestCurrency, paymentMethod PublicPaymentMethod, status PaymentRequestStatus, ) *InternalPaymentRequest`

NewInternalPaymentRequest instantiates a new InternalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalPaymentRequestWithDefaults

`func NewInternalPaymentRequestWithDefaults() *InternalPaymentRequest`

NewInternalPaymentRequestWithDefaults instantiates a new InternalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalPaymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalPaymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalPaymentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *InternalPaymentRequest) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InternalPaymentRequest) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InternalPaymentRequest) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *InternalPaymentRequest) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InternalPaymentRequest) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InternalPaymentRequest) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetReferenceId

`func (o *InternalPaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalPaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalPaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetBusinessId

`func (o *InternalPaymentRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *InternalPaymentRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *InternalPaymentRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCustomerId

`func (o *InternalPaymentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternalPaymentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternalPaymentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InternalPaymentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InternalPaymentRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InternalPaymentRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *InternalPaymentRequest) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalPaymentRequest) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalPaymentRequest) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalPaymentRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *InternalPaymentRequest) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *InternalPaymentRequest) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetAmount

`func (o *InternalPaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalPaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalPaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalPaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCountry

`func (o *InternalPaymentRequest) GetCountry() PaymentRequestCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalPaymentRequest) GetCountryOk() (*PaymentRequestCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalPaymentRequest) SetCountry(v PaymentRequestCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalPaymentRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalPaymentRequest) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalPaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalPaymentRequest) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *InternalPaymentRequest) GetPaymentMethod() PublicPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InternalPaymentRequest) GetPaymentMethodOk() (*PublicPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InternalPaymentRequest) SetPaymentMethod(v PublicPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDescription

`func (o *InternalPaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalPaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalPaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalPaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternalPaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternalPaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFailureCode

`func (o *InternalPaymentRequest) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *InternalPaymentRequest) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *InternalPaymentRequest) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *InternalPaymentRequest) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *InternalPaymentRequest) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *InternalPaymentRequest) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCaptureMethod

`func (o *InternalPaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *InternalPaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *InternalPaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *InternalPaymentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### SetCaptureMethodNil

`func (o *InternalPaymentRequest) SetCaptureMethodNil(b bool)`

 SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil

### UnsetCaptureMethod
`func (o *InternalPaymentRequest) UnsetCaptureMethod()`

UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
### GetInitiator

`func (o *InternalPaymentRequest) GetInitiator() PaymentRequestInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InternalPaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InternalPaymentRequest) SetInitiator(v PaymentRequestInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InternalPaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *InternalPaymentRequest) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *InternalPaymentRequest) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetCardVerificationResults

`func (o *InternalPaymentRequest) GetCardVerificationResults() PaymentRequestCardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *InternalPaymentRequest) GetCardVerificationResultsOk() (*PaymentRequestCardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *InternalPaymentRequest) SetCardVerificationResults(v PaymentRequestCardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *InternalPaymentRequest) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *InternalPaymentRequest) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *InternalPaymentRequest) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
### GetStatus

`func (o *InternalPaymentRequest) GetStatus() PaymentRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalPaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalPaymentRequest) SetStatus(v PaymentRequestStatus)`

SetStatus sets Status field to given value.


### GetActions

`func (o *InternalPaymentRequest) GetActions() []PaymentRequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InternalPaymentRequest) GetActionsOk() (*[]PaymentRequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InternalPaymentRequest) SetActions(v []PaymentRequestAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InternalPaymentRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalPaymentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalPaymentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalPaymentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalPaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalPaymentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalPaymentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetShippingInformation

`func (o *InternalPaymentRequest) GetShippingInformation() PaymentRequestShippingInformation`

GetShippingInformation returns the ShippingInformation field if non-nil, zero value otherwise.

### GetShippingInformationOk

`func (o *InternalPaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool)`

GetShippingInformationOk returns a tuple with the ShippingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInformation

`func (o *InternalPaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation)`

SetShippingInformation sets ShippingInformation field to given value.

### HasShippingInformation

`func (o *InternalPaymentRequest) HasShippingInformation() bool`

HasShippingInformation returns a boolean if a field has been set.

### SetShippingInformationNil

`func (o *InternalPaymentRequest) SetShippingInformationNil(b bool)`

 SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil

### UnsetShippingInformation
`func (o *InternalPaymentRequest) UnsetShippingInformation()`

UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
### GetItems

`func (o *InternalPaymentRequest) GetItems() []PaymentRequestBasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalPaymentRequest) GetItemsOk() (*[]PaymentRequestBasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalPaymentRequest) SetItems(v []PaymentRequestBasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *InternalPaymentRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *InternalPaymentRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *InternalPaymentRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetInternalMetadata

`func (o *InternalPaymentRequest) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalPaymentRequest) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalPaymentRequest) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalPaymentRequest) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalPaymentRequest) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalPaymentRequest) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil
### GetInstrumentId

`func (o *InternalPaymentRequest) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *InternalPaymentRequest) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *InternalPaymentRequest) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *InternalPaymentRequest) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


