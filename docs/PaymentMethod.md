# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PaymentMethodType**](PaymentMethodType.md) |  | 
**Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Card** | Pointer to [**NullableMutableCard**](MutableCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**NullableMutableCrypto**](MutableCrypto.md) |  | [optional] 
**DirectBankTransfer** | Pointer to [**NullableMutableDirectBankTransfer**](MutableDirectBankTransfer.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullableMutableDirectDebit**](MutableDirectDebit.md) |  | [optional] 
**Ewallet** | Pointer to [**NullableMutableEwallet**](MutableEwallet.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullableMutableOverTheCounter**](MutableOverTheCounter.md) |  | [optional] 
**VirtualAccount** | Pointer to [**NullableMutableVirtualAccount**](MutableVirtualAccount.md) |  | [optional] 
**QrCode** | Pointer to [**NullableMutableQRCode**](MutableQRCode.md) |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(type_ PaymentMethodType, reusability PaymentMethodReusability, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetReusability

`func (o *PaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetDescription

`func (o *PaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReferenceId

`func (o *PaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetCard

`func (o *PaymentMethod) GetCard() MutableCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethod) GetCardOk() (*MutableCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethod) SetCard(v MutableCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *PaymentMethod) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *PaymentMethod) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCryptocurrency

`func (o *PaymentMethod) GetCryptocurrency() MutableCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *PaymentMethod) GetCryptocurrencyOk() (*MutableCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *PaymentMethod) SetCryptocurrency(v MutableCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *PaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### SetCryptocurrencyNil

`func (o *PaymentMethod) SetCryptocurrencyNil(b bool)`

 SetCryptocurrencyNil sets the value for Cryptocurrency to be an explicit nil

### UnsetCryptocurrency
`func (o *PaymentMethod) UnsetCryptocurrency()`

UnsetCryptocurrency ensures that no value is present for Cryptocurrency, not even an explicit nil
### GetDirectBankTransfer

`func (o *PaymentMethod) GetDirectBankTransfer() MutableDirectBankTransfer`

GetDirectBankTransfer returns the DirectBankTransfer field if non-nil, zero value otherwise.

### GetDirectBankTransferOk

`func (o *PaymentMethod) GetDirectBankTransferOk() (*MutableDirectBankTransfer, bool)`

GetDirectBankTransferOk returns a tuple with the DirectBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectBankTransfer

`func (o *PaymentMethod) SetDirectBankTransfer(v MutableDirectBankTransfer)`

SetDirectBankTransfer sets DirectBankTransfer field to given value.

### HasDirectBankTransfer

`func (o *PaymentMethod) HasDirectBankTransfer() bool`

HasDirectBankTransfer returns a boolean if a field has been set.

### SetDirectBankTransferNil

`func (o *PaymentMethod) SetDirectBankTransferNil(b bool)`

 SetDirectBankTransferNil sets the value for DirectBankTransfer to be an explicit nil

### UnsetDirectBankTransfer
`func (o *PaymentMethod) UnsetDirectBankTransfer()`

UnsetDirectBankTransfer ensures that no value is present for DirectBankTransfer, not even an explicit nil
### GetDirectDebit

`func (o *PaymentMethod) GetDirectDebit() MutableDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethod) GetDirectDebitOk() (*MutableDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethod) SetDirectDebit(v MutableDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *PaymentMethod) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *PaymentMethod) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetEwallet

`func (o *PaymentMethod) GetEwallet() MutableEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PaymentMethod) GetEwalletOk() (*MutableEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PaymentMethod) SetEwallet(v MutableEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *PaymentMethod) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *PaymentMethod) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetOverTheCounter

`func (o *PaymentMethod) GetOverTheCounter() MutableOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PaymentMethod) GetOverTheCounterOk() (*MutableOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PaymentMethod) SetOverTheCounter(v MutableOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *PaymentMethod) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *PaymentMethod) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetVirtualAccount

`func (o *PaymentMethod) GetVirtualAccount() MutableVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PaymentMethod) GetVirtualAccountOk() (*MutableVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PaymentMethod) SetVirtualAccount(v MutableVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### SetVirtualAccountNil

`func (o *PaymentMethod) SetVirtualAccountNil(b bool)`

 SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil

### UnsetVirtualAccount
`func (o *PaymentMethod) UnsetVirtualAccount()`

UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
### GetQrCode

`func (o *PaymentMethod) GetQrCode() MutableQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentMethod) GetQrCodeOk() (*MutableQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentMethod) SetQrCode(v MutableQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *PaymentMethod) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PaymentMethod) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


