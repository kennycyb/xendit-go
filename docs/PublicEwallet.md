# PublicEwallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**EwalletChannelCode**](EwalletChannelCode.md) |  | 
**ChannelProperties** | Pointer to [**EwalletChannelProperties**](EwalletChannelProperties.md) |  | [optional] 
**Account** | Pointer to [**EwalletAccount**](EwalletAccount.md) |  | [optional] 

## Methods

### NewPublicEwallet

`func NewPublicEwallet(channelCode EwalletChannelCode, ) *PublicEwallet`

NewPublicEwallet instantiates a new PublicEwallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEwalletWithDefaults

`func NewPublicEwalletWithDefaults() *PublicEwallet`

NewPublicEwalletWithDefaults instantiates a new PublicEwallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PublicEwallet) GetChannelCode() EwalletChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicEwallet) SetChannelCode(v EwalletChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *PublicEwallet) GetChannelProperties() EwalletChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicEwallet) GetChannelPropertiesOk() (*EwalletChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicEwallet) SetChannelProperties(v EwalletChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PublicEwallet) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAccount

`func (o *PublicEwallet) GetAccount() EwalletAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PublicEwallet) GetAccountOk() (*EwalletAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PublicEwallet) SetAccount(v EwalletAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PublicEwallet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


