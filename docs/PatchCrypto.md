# PatchCrypto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**NullableCryptoChannelCode**](CryptoChannelCode.md) |  | [optional] 
**Wallet** | Pointer to [**NullableCryptoWallet**](CryptoWallet.md) |  | [optional] 

## Methods

### NewPatchCrypto

`func NewPatchCrypto() *PatchCrypto`

NewPatchCrypto instantiates a new PatchCrypto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCryptoWithDefaults

`func NewPatchCryptoWithDefaults() *PatchCrypto`

NewPatchCryptoWithDefaults instantiates a new PatchCrypto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PatchCrypto) GetChannelCode() CryptoChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PatchCrypto) GetChannelCodeOk() (*CryptoChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PatchCrypto) SetChannelCode(v CryptoChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *PatchCrypto) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *PatchCrypto) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *PatchCrypto) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetWallet

`func (o *PatchCrypto) GetWallet() CryptoWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *PatchCrypto) GetWalletOk() (*CryptoWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *PatchCrypto) SetWallet(v CryptoWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *PatchCrypto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### SetWalletNil

`func (o *PatchCrypto) SetWalletNil(b bool)`

 SetWalletNil sets the value for Wallet to be an explicit nil

### UnsetWallet
`func (o *PatchCrypto) UnsetWallet()`

UnsetWallet ensures that no value is present for Wallet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


