# InternalV1CreatePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | [**PaymentMethodType**](PaymentMethodType.md) |  | 
**Country** | Pointer to [**PaymentMethodCountry**](PaymentMethodCountry.md) |  | [optional] 
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**BusinessId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Reusability** | Pointer to [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | [optional] 
**Ewallet** | Pointer to [**NullableInternalEwallet**](InternalEwallet.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullableInternalDirectDebit**](InternalDirectDebit.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Actions** | Pointer to [**[]PaymentMethodAction**](PaymentMethodAction.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInternalV1CreatePaymentMethod

`func NewInternalV1CreatePaymentMethod(type_ PaymentMethodType, ) *InternalV1CreatePaymentMethod`

NewInternalV1CreatePaymentMethod instantiates a new InternalV1CreatePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalV1CreatePaymentMethodWithDefaults

`func NewInternalV1CreatePaymentMethodWithDefaults() *InternalV1CreatePaymentMethod`

NewInternalV1CreatePaymentMethodWithDefaults instantiates a new InternalV1CreatePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalV1CreatePaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalV1CreatePaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalV1CreatePaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InternalV1CreatePaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InternalV1CreatePaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalV1CreatePaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalV1CreatePaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetCountry

`func (o *InternalV1CreatePaymentMethod) GetCountry() PaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalV1CreatePaymentMethod) GetCountryOk() (*PaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalV1CreatePaymentMethod) SetCountry(v PaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalV1CreatePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStatus

`func (o *InternalV1CreatePaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalV1CreatePaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalV1CreatePaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalV1CreatePaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBusinessId

`func (o *InternalV1CreatePaymentMethod) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *InternalV1CreatePaymentMethod) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *InternalV1CreatePaymentMethod) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *InternalV1CreatePaymentMethod) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *InternalV1CreatePaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternalV1CreatePaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternalV1CreatePaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InternalV1CreatePaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InternalV1CreatePaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InternalV1CreatePaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetReusability

`func (o *InternalV1CreatePaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *InternalV1CreatePaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *InternalV1CreatePaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *InternalV1CreatePaymentMethod) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetEwallet

`func (o *InternalV1CreatePaymentMethod) GetEwallet() InternalEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *InternalV1CreatePaymentMethod) GetEwalletOk() (*InternalEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *InternalV1CreatePaymentMethod) SetEwallet(v InternalEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *InternalV1CreatePaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *InternalV1CreatePaymentMethod) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *InternalV1CreatePaymentMethod) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetDirectDebit

`func (o *InternalV1CreatePaymentMethod) GetDirectDebit() InternalDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *InternalV1CreatePaymentMethod) GetDirectDebitOk() (*InternalDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *InternalV1CreatePaymentMethod) SetDirectDebit(v InternalDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *InternalV1CreatePaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *InternalV1CreatePaymentMethod) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *InternalV1CreatePaymentMethod) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetMetadata

`func (o *InternalV1CreatePaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalV1CreatePaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalV1CreatePaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalV1CreatePaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalV1CreatePaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalV1CreatePaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetInternalMetadata

`func (o *InternalV1CreatePaymentMethod) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalV1CreatePaymentMethod) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalV1CreatePaymentMethod) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalV1CreatePaymentMethod) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalV1CreatePaymentMethod) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalV1CreatePaymentMethod) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil
### GetActions

`func (o *InternalV1CreatePaymentMethod) GetActions() []PaymentMethodAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InternalV1CreatePaymentMethod) GetActionsOk() (*[]PaymentMethodAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InternalV1CreatePaymentMethod) SetActions(v []PaymentMethodAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InternalV1CreatePaymentMethod) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCreated

`func (o *InternalV1CreatePaymentMethod) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InternalV1CreatePaymentMethod) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InternalV1CreatePaymentMethod) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InternalV1CreatePaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *InternalV1CreatePaymentMethod) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InternalV1CreatePaymentMethod) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InternalV1CreatePaymentMethod) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *InternalV1CreatePaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


