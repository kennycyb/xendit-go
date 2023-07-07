# PublicPaymentRequest

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

## Methods

### NewPublicPaymentRequest

`func NewPublicPaymentRequest(id string, created string, updated string, referenceId string, businessId string, currency PaymentRequestCurrency, paymentMethod PublicPaymentMethod, status PaymentRequestStatus, ) *PublicPaymentRequest`

NewPublicPaymentRequest instantiates a new PublicPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentRequestWithDefaults

`func NewPublicPaymentRequestWithDefaults() *PublicPaymentRequest`

NewPublicPaymentRequestWithDefaults instantiates a new PublicPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicPaymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicPaymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicPaymentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *PublicPaymentRequest) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicPaymentRequest) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicPaymentRequest) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *PublicPaymentRequest) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PublicPaymentRequest) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PublicPaymentRequest) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetReferenceId

`func (o *PublicPaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PublicPaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PublicPaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetBusinessId

`func (o *PublicPaymentRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PublicPaymentRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PublicPaymentRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCustomerId

`func (o *PublicPaymentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PublicPaymentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PublicPaymentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PublicPaymentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PublicPaymentRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PublicPaymentRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *PublicPaymentRequest) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *PublicPaymentRequest) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *PublicPaymentRequest) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *PublicPaymentRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *PublicPaymentRequest) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *PublicPaymentRequest) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetAmount

`func (o *PublicPaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicPaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicPaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PublicPaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCountry

`func (o *PublicPaymentRequest) GetCountry() PaymentRequestCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicPaymentRequest) GetCountryOk() (*PaymentRequestCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicPaymentRequest) SetCountry(v PaymentRequestCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicPaymentRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *PublicPaymentRequest) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PublicPaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PublicPaymentRequest) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *PublicPaymentRequest) GetPaymentMethod() PublicPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PublicPaymentRequest) GetPaymentMethodOk() (*PublicPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PublicPaymentRequest) SetPaymentMethod(v PublicPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDescription

`func (o *PublicPaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicPaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicPaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicPaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PublicPaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PublicPaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFailureCode

`func (o *PublicPaymentRequest) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PublicPaymentRequest) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PublicPaymentRequest) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PublicPaymentRequest) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PublicPaymentRequest) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PublicPaymentRequest) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCaptureMethod

`func (o *PublicPaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *PublicPaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *PublicPaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *PublicPaymentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### SetCaptureMethodNil

`func (o *PublicPaymentRequest) SetCaptureMethodNil(b bool)`

 SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil

### UnsetCaptureMethod
`func (o *PublicPaymentRequest) UnsetCaptureMethod()`

UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
### GetInitiator

`func (o *PublicPaymentRequest) GetInitiator() PaymentRequestInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *PublicPaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *PublicPaymentRequest) SetInitiator(v PaymentRequestInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *PublicPaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *PublicPaymentRequest) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *PublicPaymentRequest) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetCardVerificationResults

`func (o *PublicPaymentRequest) GetCardVerificationResults() PaymentRequestCardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *PublicPaymentRequest) GetCardVerificationResultsOk() (*PaymentRequestCardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *PublicPaymentRequest) SetCardVerificationResults(v PaymentRequestCardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *PublicPaymentRequest) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *PublicPaymentRequest) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *PublicPaymentRequest) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
### GetStatus

`func (o *PublicPaymentRequest) GetStatus() PaymentRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicPaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicPaymentRequest) SetStatus(v PaymentRequestStatus)`

SetStatus sets Status field to given value.


### GetActions

`func (o *PublicPaymentRequest) GetActions() []PaymentRequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PublicPaymentRequest) GetActionsOk() (*[]PaymentRequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PublicPaymentRequest) SetActions(v []PaymentRequestAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PublicPaymentRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *PublicPaymentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicPaymentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicPaymentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicPaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PublicPaymentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PublicPaymentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetShippingInformation

`func (o *PublicPaymentRequest) GetShippingInformation() PaymentRequestShippingInformation`

GetShippingInformation returns the ShippingInformation field if non-nil, zero value otherwise.

### GetShippingInformationOk

`func (o *PublicPaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool)`

GetShippingInformationOk returns a tuple with the ShippingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInformation

`func (o *PublicPaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation)`

SetShippingInformation sets ShippingInformation field to given value.

### HasShippingInformation

`func (o *PublicPaymentRequest) HasShippingInformation() bool`

HasShippingInformation returns a boolean if a field has been set.

### SetShippingInformationNil

`func (o *PublicPaymentRequest) SetShippingInformationNil(b bool)`

 SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil

### UnsetShippingInformation
`func (o *PublicPaymentRequest) UnsetShippingInformation()`

UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
### GetItems

`func (o *PublicPaymentRequest) GetItems() []PaymentRequestBasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PublicPaymentRequest) GetItemsOk() (*[]PaymentRequestBasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PublicPaymentRequest) SetItems(v []PaymentRequestBasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PublicPaymentRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *PublicPaymentRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *PublicPaymentRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


