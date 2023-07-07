# PublicDirectDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**DirectDebitChannelCode**](DirectDebitChannelCode.md) |  | 
**ChannelProperties** | [**NullableDirectDebitChannelProperties**](DirectDebitChannelProperties.md) |  | 
**Type** | [**DirectDebitType**](DirectDebitType.md) |  | 
**BankAccount** | Pointer to [**NullableDirectDebitBankAccount**](DirectDebitBankAccount.md) |  | [optional] 
**DebitCard** | Pointer to [**NullableDirectDebitDebitCard**](DirectDebitDebitCard.md) |  | [optional] 

## Methods

### NewPublicDirectDebit

`func NewPublicDirectDebit(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, type_ DirectDebitType, ) *PublicDirectDebit`

NewPublicDirectDebit instantiates a new PublicDirectDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDirectDebitWithDefaults

`func NewPublicDirectDebitWithDefaults() *PublicDirectDebit`

NewPublicDirectDebitWithDefaults instantiates a new PublicDirectDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PublicDirectDebit) GetChannelCode() DirectDebitChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicDirectDebit) GetChannelCodeOk() (*DirectDebitChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicDirectDebit) SetChannelCode(v DirectDebitChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *PublicDirectDebit) GetChannelProperties() DirectDebitChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicDirectDebit) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicDirectDebit) SetChannelProperties(v DirectDebitChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *PublicDirectDebit) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *PublicDirectDebit) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetType

`func (o *PublicDirectDebit) GetType() DirectDebitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicDirectDebit) GetTypeOk() (*DirectDebitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicDirectDebit) SetType(v DirectDebitType)`

SetType sets Type field to given value.


### GetBankAccount

`func (o *PublicDirectDebit) GetBankAccount() DirectDebitBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *PublicDirectDebit) GetBankAccountOk() (*DirectDebitBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *PublicDirectDebit) SetBankAccount(v DirectDebitBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *PublicDirectDebit) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### SetBankAccountNil

`func (o *PublicDirectDebit) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *PublicDirectDebit) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetDebitCard

`func (o *PublicDirectDebit) GetDebitCard() DirectDebitDebitCard`

GetDebitCard returns the DebitCard field if non-nil, zero value otherwise.

### GetDebitCardOk

`func (o *PublicDirectDebit) GetDebitCardOk() (*DirectDebitDebitCard, bool)`

GetDebitCardOk returns a tuple with the DebitCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitCard

`func (o *PublicDirectDebit) SetDebitCard(v DirectDebitDebitCard)`

SetDebitCard sets DebitCard field to given value.

### HasDebitCard

`func (o *PublicDirectDebit) HasDebitCard() bool`

HasDebitCard returns a boolean if a field has been set.

### SetDebitCardNil

`func (o *PublicDirectDebit) SetDebitCardNil(b bool)`

 SetDebitCardNil sets the value for DebitCard to be an explicit nil

### UnsetDebitCard
`func (o *PublicDirectDebit) UnsetDebitCard()`

UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


