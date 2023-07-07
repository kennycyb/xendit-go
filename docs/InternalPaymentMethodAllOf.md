# InternalPaymentMethodAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ewallet** | Pointer to [**NullableInternalEwallet**](InternalEwallet.md) |  | [optional] 
**Card** | Pointer to [**NullableInternalCard**](InternalCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**NullableInternalCrypto**](InternalCrypto.md) |  | [optional] 
**DirectBankTransfer** | Pointer to [**NullableInternalDirectBankTransfer**](InternalDirectBankTransfer.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullableInternalDirectDebit**](InternalDirectDebit.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullableInternalOverTheCounter**](InternalOverTheCounter.md) |  | [optional] 
**VirtualAccount** | Pointer to [**NullableInternalVirtualAccount**](InternalVirtualAccount.md) |  | [optional] 
**QrCode** | Pointer to [**NullableInternalQRCode**](InternalQRCode.md) |  | [optional] 
**XCallbackUrl** | Pointer to **NullableString** |  | [optional] 
**ClientType** | Pointer to **NullableString** |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInternalPaymentMethodAllOf

`func NewInternalPaymentMethodAllOf() *InternalPaymentMethodAllOf`

NewInternalPaymentMethodAllOf instantiates a new InternalPaymentMethodAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalPaymentMethodAllOfWithDefaults

`func NewInternalPaymentMethodAllOfWithDefaults() *InternalPaymentMethodAllOf`

NewInternalPaymentMethodAllOfWithDefaults instantiates a new InternalPaymentMethodAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEwallet

`func (o *InternalPaymentMethodAllOf) GetEwallet() InternalEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *InternalPaymentMethodAllOf) GetEwalletOk() (*InternalEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *InternalPaymentMethodAllOf) SetEwallet(v InternalEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *InternalPaymentMethodAllOf) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *InternalPaymentMethodAllOf) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *InternalPaymentMethodAllOf) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetCard

`func (o *InternalPaymentMethodAllOf) GetCard() InternalCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *InternalPaymentMethodAllOf) GetCardOk() (*InternalCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *InternalPaymentMethodAllOf) SetCard(v InternalCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *InternalPaymentMethodAllOf) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *InternalPaymentMethodAllOf) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *InternalPaymentMethodAllOf) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCryptocurrency

`func (o *InternalPaymentMethodAllOf) GetCryptocurrency() InternalCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *InternalPaymentMethodAllOf) GetCryptocurrencyOk() (*InternalCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *InternalPaymentMethodAllOf) SetCryptocurrency(v InternalCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *InternalPaymentMethodAllOf) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### SetCryptocurrencyNil

`func (o *InternalPaymentMethodAllOf) SetCryptocurrencyNil(b bool)`

 SetCryptocurrencyNil sets the value for Cryptocurrency to be an explicit nil

### UnsetCryptocurrency
`func (o *InternalPaymentMethodAllOf) UnsetCryptocurrency()`

UnsetCryptocurrency ensures that no value is present for Cryptocurrency, not even an explicit nil
### GetDirectBankTransfer

`func (o *InternalPaymentMethodAllOf) GetDirectBankTransfer() InternalDirectBankTransfer`

GetDirectBankTransfer returns the DirectBankTransfer field if non-nil, zero value otherwise.

### GetDirectBankTransferOk

`func (o *InternalPaymentMethodAllOf) GetDirectBankTransferOk() (*InternalDirectBankTransfer, bool)`

GetDirectBankTransferOk returns a tuple with the DirectBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectBankTransfer

`func (o *InternalPaymentMethodAllOf) SetDirectBankTransfer(v InternalDirectBankTransfer)`

SetDirectBankTransfer sets DirectBankTransfer field to given value.

### HasDirectBankTransfer

`func (o *InternalPaymentMethodAllOf) HasDirectBankTransfer() bool`

HasDirectBankTransfer returns a boolean if a field has been set.

### SetDirectBankTransferNil

`func (o *InternalPaymentMethodAllOf) SetDirectBankTransferNil(b bool)`

 SetDirectBankTransferNil sets the value for DirectBankTransfer to be an explicit nil

### UnsetDirectBankTransfer
`func (o *InternalPaymentMethodAllOf) UnsetDirectBankTransfer()`

UnsetDirectBankTransfer ensures that no value is present for DirectBankTransfer, not even an explicit nil
### GetDirectDebit

`func (o *InternalPaymentMethodAllOf) GetDirectDebit() InternalDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *InternalPaymentMethodAllOf) GetDirectDebitOk() (*InternalDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *InternalPaymentMethodAllOf) SetDirectDebit(v InternalDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *InternalPaymentMethodAllOf) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *InternalPaymentMethodAllOf) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *InternalPaymentMethodAllOf) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetOverTheCounter

`func (o *InternalPaymentMethodAllOf) GetOverTheCounter() InternalOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *InternalPaymentMethodAllOf) GetOverTheCounterOk() (*InternalOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *InternalPaymentMethodAllOf) SetOverTheCounter(v InternalOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *InternalPaymentMethodAllOf) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *InternalPaymentMethodAllOf) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *InternalPaymentMethodAllOf) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetVirtualAccount

`func (o *InternalPaymentMethodAllOf) GetVirtualAccount() InternalVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *InternalPaymentMethodAllOf) GetVirtualAccountOk() (*InternalVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *InternalPaymentMethodAllOf) SetVirtualAccount(v InternalVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *InternalPaymentMethodAllOf) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### SetVirtualAccountNil

`func (o *InternalPaymentMethodAllOf) SetVirtualAccountNil(b bool)`

 SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil

### UnsetVirtualAccount
`func (o *InternalPaymentMethodAllOf) UnsetVirtualAccount()`

UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
### GetQrCode

`func (o *InternalPaymentMethodAllOf) GetQrCode() InternalQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *InternalPaymentMethodAllOf) GetQrCodeOk() (*InternalQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *InternalPaymentMethodAllOf) SetQrCode(v InternalQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *InternalPaymentMethodAllOf) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *InternalPaymentMethodAllOf) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *InternalPaymentMethodAllOf) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
### GetXCallbackUrl

`func (o *InternalPaymentMethodAllOf) GetXCallbackUrl() string`

GetXCallbackUrl returns the XCallbackUrl field if non-nil, zero value otherwise.

### GetXCallbackUrlOk

`func (o *InternalPaymentMethodAllOf) GetXCallbackUrlOk() (*string, bool)`

GetXCallbackUrlOk returns a tuple with the XCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCallbackUrl

`func (o *InternalPaymentMethodAllOf) SetXCallbackUrl(v string)`

SetXCallbackUrl sets XCallbackUrl field to given value.

### HasXCallbackUrl

`func (o *InternalPaymentMethodAllOf) HasXCallbackUrl() bool`

HasXCallbackUrl returns a boolean if a field has been set.

### SetXCallbackUrlNil

`func (o *InternalPaymentMethodAllOf) SetXCallbackUrlNil(b bool)`

 SetXCallbackUrlNil sets the value for XCallbackUrl to be an explicit nil

### UnsetXCallbackUrl
`func (o *InternalPaymentMethodAllOf) UnsetXCallbackUrl()`

UnsetXCallbackUrl ensures that no value is present for XCallbackUrl, not even an explicit nil
### GetClientType

`func (o *InternalPaymentMethodAllOf) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *InternalPaymentMethodAllOf) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *InternalPaymentMethodAllOf) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *InternalPaymentMethodAllOf) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### SetClientTypeNil

`func (o *InternalPaymentMethodAllOf) SetClientTypeNil(b bool)`

 SetClientTypeNil sets the value for ClientType to be an explicit nil

### UnsetClientType
`func (o *InternalPaymentMethodAllOf) UnsetClientType()`

UnsetClientType ensures that no value is present for ClientType, not even an explicit nil
### GetInternalMetadata

`func (o *InternalPaymentMethodAllOf) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalPaymentMethodAllOf) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalPaymentMethodAllOf) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalPaymentMethodAllOf) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalPaymentMethodAllOf) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalPaymentMethodAllOf) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


