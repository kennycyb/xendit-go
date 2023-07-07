# InternalPatchPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**PatchCrypto**](PatchCrypto.md) |  | [optional] 
**OverTheCounter** | Pointer to [**PatchOverTheCounter**](PatchOverTheCounter.md) |  | [optional] 

## Methods

### NewInternalPatchPaymentMethod

`func NewInternalPatchPaymentMethod() *InternalPatchPaymentMethod`

NewInternalPatchPaymentMethod instantiates a new InternalPatchPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalPatchPaymentMethodWithDefaults

`func NewInternalPatchPaymentMethodWithDefaults() *InternalPatchPaymentMethod`

NewInternalPatchPaymentMethodWithDefaults instantiates a new InternalPatchPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InternalPatchPaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalPatchPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalPatchPaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalPatchPaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCryptocurrency

`func (o *InternalPatchPaymentMethod) GetCryptocurrency() PatchCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *InternalPatchPaymentMethod) GetCryptocurrencyOk() (*PatchCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *InternalPatchPaymentMethod) SetCryptocurrency(v PatchCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *InternalPatchPaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *InternalPatchPaymentMethod) GetOverTheCounter() PatchOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *InternalPatchPaymentMethod) GetOverTheCounterOk() (*PatchOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *InternalPatchPaymentMethod) SetOverTheCounter(v PatchOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *InternalPatchPaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


