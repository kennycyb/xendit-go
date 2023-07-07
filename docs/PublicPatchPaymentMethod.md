# PublicPatchPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Reusability** | Pointer to [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | [optional] 
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**OverTheCounter** | Pointer to [**PatchOverTheCounter**](PatchOverTheCounter.md) |  | [optional] 
**VirtualAccount** | Pointer to [**PatchVirtualAccount**](PatchVirtualAccount.md) |  | [optional] 

## Methods

### NewPublicPatchPaymentMethod

`func NewPublicPatchPaymentMethod() *PublicPatchPaymentMethod`

NewPublicPatchPaymentMethod instantiates a new PublicPatchPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPatchPaymentMethodWithDefaults

`func NewPublicPatchPaymentMethodWithDefaults() *PublicPatchPaymentMethod`

NewPublicPatchPaymentMethodWithDefaults instantiates a new PublicPatchPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PublicPatchPaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicPatchPaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicPatchPaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicPatchPaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReferenceId

`func (o *PublicPatchPaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PublicPatchPaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PublicPatchPaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PublicPatchPaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReusability

`func (o *PublicPatchPaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PublicPatchPaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PublicPatchPaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *PublicPatchPaymentMethod) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetStatus

`func (o *PublicPatchPaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicPatchPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicPatchPaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicPatchPaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *PublicPatchPaymentMethod) GetOverTheCounter() PatchOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PublicPatchPaymentMethod) GetOverTheCounterOk() (*PatchOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PublicPatchPaymentMethod) SetOverTheCounter(v PatchOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PublicPatchPaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *PublicPatchPaymentMethod) GetVirtualAccount() PatchVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PublicPatchPaymentMethod) GetVirtualAccountOk() (*PatchVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PublicPatchPaymentMethod) SetVirtualAccount(v PatchVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PublicPatchPaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


