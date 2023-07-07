# InternalEwallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**EwalletChannelCode**](EwalletChannelCode.md) |  | 
**ChannelProperties** | Pointer to [**EwalletChannelProperties**](EwalletChannelProperties.md) |  | [optional] 
**Account** | Pointer to [**EwalletAccount**](EwalletAccount.md) |  | [optional] 
**LinkedAccountTokenId** | Pointer to **NullableString** |  | [optional] 
**LinkedAccountId** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInternalEwallet

`func NewInternalEwallet(channelCode EwalletChannelCode, ) *InternalEwallet`

NewInternalEwallet instantiates a new InternalEwallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalEwalletWithDefaults

`func NewInternalEwalletWithDefaults() *InternalEwallet`

NewInternalEwalletWithDefaults instantiates a new InternalEwallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *InternalEwallet) GetChannelCode() EwalletChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalEwallet) SetChannelCode(v EwalletChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *InternalEwallet) GetChannelProperties() EwalletChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalEwallet) GetChannelPropertiesOk() (*EwalletChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalEwallet) SetChannelProperties(v EwalletChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *InternalEwallet) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAccount

`func (o *InternalEwallet) GetAccount() EwalletAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InternalEwallet) GetAccountOk() (*EwalletAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InternalEwallet) SetAccount(v EwalletAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InternalEwallet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinkedAccountTokenId

`func (o *InternalEwallet) GetLinkedAccountTokenId() string`

GetLinkedAccountTokenId returns the LinkedAccountTokenId field if non-nil, zero value otherwise.

### GetLinkedAccountTokenIdOk

`func (o *InternalEwallet) GetLinkedAccountTokenIdOk() (*string, bool)`

GetLinkedAccountTokenIdOk returns a tuple with the LinkedAccountTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountTokenId

`func (o *InternalEwallet) SetLinkedAccountTokenId(v string)`

SetLinkedAccountTokenId sets LinkedAccountTokenId field to given value.

### HasLinkedAccountTokenId

`func (o *InternalEwallet) HasLinkedAccountTokenId() bool`

HasLinkedAccountTokenId returns a boolean if a field has been set.

### SetLinkedAccountTokenIdNil

`func (o *InternalEwallet) SetLinkedAccountTokenIdNil(b bool)`

 SetLinkedAccountTokenIdNil sets the value for LinkedAccountTokenId to be an explicit nil

### UnsetLinkedAccountTokenId
`func (o *InternalEwallet) UnsetLinkedAccountTokenId()`

UnsetLinkedAccountTokenId ensures that no value is present for LinkedAccountTokenId, not even an explicit nil
### GetLinkedAccountId

`func (o *InternalEwallet) GetLinkedAccountId() string`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *InternalEwallet) GetLinkedAccountIdOk() (*string, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *InternalEwallet) SetLinkedAccountId(v string)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *InternalEwallet) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### SetLinkedAccountIdNil

`func (o *InternalEwallet) SetLinkedAccountIdNil(b bool)`

 SetLinkedAccountIdNil sets the value for LinkedAccountId to be an explicit nil

### UnsetLinkedAccountId
`func (o *InternalEwallet) UnsetLinkedAccountId()`

UnsetLinkedAccountId ensures that no value is present for LinkedAccountId, not even an explicit nil
### GetAccessToken

`func (o *InternalEwallet) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *InternalEwallet) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *InternalEwallet) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *InternalEwallet) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *InternalEwallet) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *InternalEwallet) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


