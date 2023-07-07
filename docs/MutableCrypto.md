# MutableCrypto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**NullableCryptoChannelCode**](CryptoChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**CryptoChannelProperties**](CryptoChannelProperties.md) |  | [optional] 
**Wallet** | Pointer to [**NullableCryptoWallet**](CryptoWallet.md) |  | [optional] 

## Methods

### NewMutableCrypto

`func NewMutableCrypto() *MutableCrypto`

NewMutableCrypto instantiates a new MutableCrypto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableCryptoWithDefaults

`func NewMutableCryptoWithDefaults() *MutableCrypto`

NewMutableCryptoWithDefaults instantiates a new MutableCrypto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *MutableCrypto) GetChannelCode() CryptoChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *MutableCrypto) GetChannelCodeOk() (*CryptoChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *MutableCrypto) SetChannelCode(v CryptoChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *MutableCrypto) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *MutableCrypto) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *MutableCrypto) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *MutableCrypto) GetChannelProperties() CryptoChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *MutableCrypto) GetChannelPropertiesOk() (*CryptoChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *MutableCrypto) SetChannelProperties(v CryptoChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *MutableCrypto) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetWallet

`func (o *MutableCrypto) GetWallet() CryptoWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *MutableCrypto) GetWalletOk() (*CryptoWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *MutableCrypto) SetWallet(v CryptoWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *MutableCrypto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### SetWalletNil

`func (o *MutableCrypto) SetWalletNil(b bool)`

 SetWalletNil sets the value for Wallet to be an explicit nil

### UnsetWallet
`func (o *MutableCrypto) UnsetWallet()`

UnsetWallet ensures that no value is present for Wallet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


