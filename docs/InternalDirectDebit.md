# InternalDirectDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**DirectDebitChannelCode**](DirectDebitChannelCode.md) |  | 
**ChannelProperties** | [**NullableDirectDebitChannelProperties**](DirectDebitChannelProperties.md) |  | 
**Type** | [**DirectDebitType**](DirectDebitType.md) |  | 
**BankAccount** | Pointer to [**NullableDirectDebitBankAccount**](DirectDebitBankAccount.md) |  | [optional] 
**DebitCard** | Pointer to [**NullableDirectDebitDebitCard**](DirectDebitDebitCard.md) |  | [optional] 
**LinkedAccountTokenId** | Pointer to **NullableString** |  | [optional] 
**LinkedAccountId** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInternalDirectDebit

`func NewInternalDirectDebit(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, type_ DirectDebitType, ) *InternalDirectDebit`

NewInternalDirectDebit instantiates a new InternalDirectDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDirectDebitWithDefaults

`func NewInternalDirectDebitWithDefaults() *InternalDirectDebit`

NewInternalDirectDebitWithDefaults instantiates a new InternalDirectDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *InternalDirectDebit) GetChannelCode() DirectDebitChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalDirectDebit) GetChannelCodeOk() (*DirectDebitChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalDirectDebit) SetChannelCode(v DirectDebitChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *InternalDirectDebit) GetChannelProperties() DirectDebitChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalDirectDebit) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalDirectDebit) SetChannelProperties(v DirectDebitChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *InternalDirectDebit) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *InternalDirectDebit) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetType

`func (o *InternalDirectDebit) GetType() DirectDebitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalDirectDebit) GetTypeOk() (*DirectDebitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalDirectDebit) SetType(v DirectDebitType)`

SetType sets Type field to given value.


### GetBankAccount

`func (o *InternalDirectDebit) GetBankAccount() DirectDebitBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *InternalDirectDebit) GetBankAccountOk() (*DirectDebitBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *InternalDirectDebit) SetBankAccount(v DirectDebitBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *InternalDirectDebit) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### SetBankAccountNil

`func (o *InternalDirectDebit) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *InternalDirectDebit) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetDebitCard

`func (o *InternalDirectDebit) GetDebitCard() DirectDebitDebitCard`

GetDebitCard returns the DebitCard field if non-nil, zero value otherwise.

### GetDebitCardOk

`func (o *InternalDirectDebit) GetDebitCardOk() (*DirectDebitDebitCard, bool)`

GetDebitCardOk returns a tuple with the DebitCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitCard

`func (o *InternalDirectDebit) SetDebitCard(v DirectDebitDebitCard)`

SetDebitCard sets DebitCard field to given value.

### HasDebitCard

`func (o *InternalDirectDebit) HasDebitCard() bool`

HasDebitCard returns a boolean if a field has been set.

### SetDebitCardNil

`func (o *InternalDirectDebit) SetDebitCardNil(b bool)`

 SetDebitCardNil sets the value for DebitCard to be an explicit nil

### UnsetDebitCard
`func (o *InternalDirectDebit) UnsetDebitCard()`

UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil
### GetLinkedAccountTokenId

`func (o *InternalDirectDebit) GetLinkedAccountTokenId() string`

GetLinkedAccountTokenId returns the LinkedAccountTokenId field if non-nil, zero value otherwise.

### GetLinkedAccountTokenIdOk

`func (o *InternalDirectDebit) GetLinkedAccountTokenIdOk() (*string, bool)`

GetLinkedAccountTokenIdOk returns a tuple with the LinkedAccountTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountTokenId

`func (o *InternalDirectDebit) SetLinkedAccountTokenId(v string)`

SetLinkedAccountTokenId sets LinkedAccountTokenId field to given value.

### HasLinkedAccountTokenId

`func (o *InternalDirectDebit) HasLinkedAccountTokenId() bool`

HasLinkedAccountTokenId returns a boolean if a field has been set.

### SetLinkedAccountTokenIdNil

`func (o *InternalDirectDebit) SetLinkedAccountTokenIdNil(b bool)`

 SetLinkedAccountTokenIdNil sets the value for LinkedAccountTokenId to be an explicit nil

### UnsetLinkedAccountTokenId
`func (o *InternalDirectDebit) UnsetLinkedAccountTokenId()`

UnsetLinkedAccountTokenId ensures that no value is present for LinkedAccountTokenId, not even an explicit nil
### GetLinkedAccountId

`func (o *InternalDirectDebit) GetLinkedAccountId() string`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *InternalDirectDebit) GetLinkedAccountIdOk() (*string, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *InternalDirectDebit) SetLinkedAccountId(v string)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *InternalDirectDebit) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### SetLinkedAccountIdNil

`func (o *InternalDirectDebit) SetLinkedAccountIdNil(b bool)`

 SetLinkedAccountIdNil sets the value for LinkedAccountId to be an explicit nil

### UnsetLinkedAccountId
`func (o *InternalDirectDebit) UnsetLinkedAccountId()`

UnsetLinkedAccountId ensures that no value is present for LinkedAccountId, not even an explicit nil
### GetAccessToken

`func (o *InternalDirectDebit) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *InternalDirectDebit) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *InternalDirectDebit) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *InternalDirectDebit) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *InternalDirectDebit) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *InternalDirectDebit) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


