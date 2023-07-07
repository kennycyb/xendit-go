# MutableEwallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**EwalletChannelCode**](EwalletChannelCode.md) |  | 
**ChannelProperties** | Pointer to [**EwalletChannelProperties**](EwalletChannelProperties.md) |  | [optional] 
**Account** | Pointer to [**EwalletAccount**](EwalletAccount.md) |  | [optional] 

## Methods

### NewMutableEwallet

`func NewMutableEwallet(channelCode EwalletChannelCode, ) *MutableEwallet`

NewMutableEwallet instantiates a new MutableEwallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableEwalletWithDefaults

`func NewMutableEwalletWithDefaults() *MutableEwallet`

NewMutableEwalletWithDefaults instantiates a new MutableEwallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *MutableEwallet) GetChannelCode() EwalletChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *MutableEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *MutableEwallet) SetChannelCode(v EwalletChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *MutableEwallet) GetChannelProperties() EwalletChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *MutableEwallet) GetChannelPropertiesOk() (*EwalletChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *MutableEwallet) SetChannelProperties(v EwalletChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *MutableEwallet) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAccount

`func (o *MutableEwallet) GetAccount() EwalletAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MutableEwallet) GetAccountOk() (*EwalletAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MutableEwallet) SetAccount(v EwalletAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MutableEwallet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


