# BasePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BusinessId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PaymentMethodType**](PaymentMethodType.md) |  | [optional] 
**Country** | Pointer to [**PaymentMethodCountry**](PaymentMethodCountry.md) |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**Reusability** | Pointer to [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | [optional] 
**Actions** | Pointer to [**[]PaymentMethodAction**](PaymentMethodAction.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BillingInformation** | Pointer to [**NullableBillingInformation**](BillingInformation.md) |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBasePaymentMethod

`func NewBasePaymentMethod(id string, ) *BasePaymentMethod`

NewBasePaymentMethod instantiates a new BasePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasePaymentMethodWithDefaults

`func NewBasePaymentMethodWithDefaults() *BasePaymentMethod`

NewBasePaymentMethodWithDefaults instantiates a new BasePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasePaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasePaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasePaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetBusinessId

`func (o *BasePaymentMethod) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *BasePaymentMethod) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *BasePaymentMethod) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *BasePaymentMethod) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetType

`func (o *BasePaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasePaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasePaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *BasePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *BasePaymentMethod) GetCountry() PaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BasePaymentMethod) GetCountryOk() (*PaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BasePaymentMethod) SetCountry(v PaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BasePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomerId

`func (o *BasePaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BasePaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BasePaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *BasePaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *BasePaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *BasePaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *BasePaymentMethod) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *BasePaymentMethod) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *BasePaymentMethod) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *BasePaymentMethod) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *BasePaymentMethod) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *BasePaymentMethod) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetReferenceId

`func (o *BasePaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *BasePaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *BasePaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *BasePaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *BasePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BasePaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BasePaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *BasePaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BasePaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BasePaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BasePaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReusability

`func (o *BasePaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *BasePaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *BasePaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *BasePaymentMethod) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetActions

`func (o *BasePaymentMethod) GetActions() []PaymentMethodAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BasePaymentMethod) GetActionsOk() (*[]PaymentMethodAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BasePaymentMethod) SetActions(v []PaymentMethodAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BasePaymentMethod) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *BasePaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BasePaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BasePaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BasePaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BasePaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BasePaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBillingInformation

`func (o *BasePaymentMethod) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *BasePaymentMethod) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *BasePaymentMethod) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *BasePaymentMethod) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *BasePaymentMethod) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *BasePaymentMethod) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
### GetFailureCode

`func (o *BasePaymentMethod) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *BasePaymentMethod) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *BasePaymentMethod) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *BasePaymentMethod) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *BasePaymentMethod) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *BasePaymentMethod) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCreated

`func (o *BasePaymentMethod) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BasePaymentMethod) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BasePaymentMethod) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BasePaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *BasePaymentMethod) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BasePaymentMethod) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BasePaymentMethod) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BasePaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


