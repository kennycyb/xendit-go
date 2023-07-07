# PublicPaymentMethodAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ewallet** | Pointer to [**NullablePublicEwallet**](PublicEwallet.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullablePublicDirectDebit**](PublicDirectDebit.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullablePublicOverTheCounter**](PublicOverTheCounter.md) |  | [optional] 
**Card** | Pointer to [**NullablePublicCard**](PublicCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**NullablePublicCrypto**](PublicCrypto.md) |  | [optional] 
**QrCode** | Pointer to [**NullablePublicQRCode**](PublicQRCode.md) |  | [optional] 

## Methods

### NewPublicPaymentMethodAllOf

`func NewPublicPaymentMethodAllOf() *PublicPaymentMethodAllOf`

NewPublicPaymentMethodAllOf instantiates a new PublicPaymentMethodAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentMethodAllOfWithDefaults

`func NewPublicPaymentMethodAllOfWithDefaults() *PublicPaymentMethodAllOf`

NewPublicPaymentMethodAllOfWithDefaults instantiates a new PublicPaymentMethodAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEwallet

`func (o *PublicPaymentMethodAllOf) GetEwallet() PublicEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PublicPaymentMethodAllOf) GetEwalletOk() (*PublicEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PublicPaymentMethodAllOf) SetEwallet(v PublicEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PublicPaymentMethodAllOf) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *PublicPaymentMethodAllOf) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *PublicPaymentMethodAllOf) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetDirectDebit

`func (o *PublicPaymentMethodAllOf) GetDirectDebit() PublicDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PublicPaymentMethodAllOf) GetDirectDebitOk() (*PublicDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PublicPaymentMethodAllOf) SetDirectDebit(v PublicDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PublicPaymentMethodAllOf) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *PublicPaymentMethodAllOf) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *PublicPaymentMethodAllOf) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetOverTheCounter

`func (o *PublicPaymentMethodAllOf) GetOverTheCounter() PublicOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PublicPaymentMethodAllOf) GetOverTheCounterOk() (*PublicOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PublicPaymentMethodAllOf) SetOverTheCounter(v PublicOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PublicPaymentMethodAllOf) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *PublicPaymentMethodAllOf) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *PublicPaymentMethodAllOf) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetCard

`func (o *PublicPaymentMethodAllOf) GetCard() PublicCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PublicPaymentMethodAllOf) GetCardOk() (*PublicCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PublicPaymentMethodAllOf) SetCard(v PublicCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *PublicPaymentMethodAllOf) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *PublicPaymentMethodAllOf) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *PublicPaymentMethodAllOf) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCryptocurrency

`func (o *PublicPaymentMethodAllOf) GetCryptocurrency() PublicCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *PublicPaymentMethodAllOf) GetCryptocurrencyOk() (*PublicCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *PublicPaymentMethodAllOf) SetCryptocurrency(v PublicCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *PublicPaymentMethodAllOf) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### SetCryptocurrencyNil

`func (o *PublicPaymentMethodAllOf) SetCryptocurrencyNil(b bool)`

 SetCryptocurrencyNil sets the value for Cryptocurrency to be an explicit nil

### UnsetCryptocurrency
`func (o *PublicPaymentMethodAllOf) UnsetCryptocurrency()`

UnsetCryptocurrency ensures that no value is present for Cryptocurrency, not even an explicit nil
### GetQrCode

`func (o *PublicPaymentMethodAllOf) GetQrCode() PublicQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PublicPaymentMethodAllOf) GetQrCodeOk() (*PublicQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PublicPaymentMethodAllOf) SetQrCode(v PublicQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PublicPaymentMethodAllOf) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *PublicPaymentMethodAllOf) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PublicPaymentMethodAllOf) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


