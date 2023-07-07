# InternalCrypto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**NullableCryptoChannelCode**](CryptoChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**CryptoChannelProperties**](CryptoChannelProperties.md) |  | [optional] 
**Wallet** | Pointer to [**NullableCryptoWallet**](CryptoWallet.md) |  | [optional] 

## Methods

### NewInternalCrypto

`func NewInternalCrypto() *InternalCrypto`

NewInternalCrypto instantiates a new InternalCrypto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalCryptoWithDefaults

`func NewInternalCryptoWithDefaults() *InternalCrypto`

NewInternalCryptoWithDefaults instantiates a new InternalCrypto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *InternalCrypto) GetChannelCode() CryptoChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalCrypto) GetChannelCodeOk() (*CryptoChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalCrypto) SetChannelCode(v CryptoChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *InternalCrypto) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *InternalCrypto) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *InternalCrypto) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *InternalCrypto) GetChannelProperties() CryptoChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalCrypto) GetChannelPropertiesOk() (*CryptoChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalCrypto) SetChannelProperties(v CryptoChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *InternalCrypto) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetWallet

`func (o *InternalCrypto) GetWallet() CryptoWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *InternalCrypto) GetWalletOk() (*CryptoWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *InternalCrypto) SetWallet(v CryptoWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *InternalCrypto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### SetWalletNil

`func (o *InternalCrypto) SetWalletNil(b bool)`

 SetWalletNil sets the value for Wallet to be an explicit nil

### UnsetWallet
`func (o *InternalCrypto) UnsetWallet()`

UnsetWallet ensures that no value is present for Wallet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


