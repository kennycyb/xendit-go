# PublicCrypto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**NullableCryptoChannelCode**](CryptoChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**CryptoChannelProperties**](CryptoChannelProperties.md) |  | [optional] 
**Wallet** | Pointer to [**NullableCryptoWallet**](CryptoWallet.md) |  | [optional] 

## Methods

### NewPublicCrypto

`func NewPublicCrypto() *PublicCrypto`

NewPublicCrypto instantiates a new PublicCrypto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCryptoWithDefaults

`func NewPublicCryptoWithDefaults() *PublicCrypto`

NewPublicCryptoWithDefaults instantiates a new PublicCrypto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PublicCrypto) GetChannelCode() CryptoChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicCrypto) GetChannelCodeOk() (*CryptoChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicCrypto) SetChannelCode(v CryptoChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *PublicCrypto) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *PublicCrypto) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *PublicCrypto) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *PublicCrypto) GetChannelProperties() CryptoChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicCrypto) GetChannelPropertiesOk() (*CryptoChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicCrypto) SetChannelProperties(v CryptoChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PublicCrypto) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetWallet

`func (o *PublicCrypto) GetWallet() CryptoWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *PublicCrypto) GetWalletOk() (*CryptoWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *PublicCrypto) SetWallet(v CryptoWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *PublicCrypto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### SetWalletNil

`func (o *PublicCrypto) SetWalletNil(b bool)`

 SetWalletNil sets the value for Wallet to be an explicit nil

### UnsetWallet
`func (o *PublicCrypto) UnsetWallet()`

UnsetWallet ensures that no value is present for Wallet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


