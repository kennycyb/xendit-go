# PublicDirectDebitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DirectDebitType**](DirectDebitType.md) |  | 
**BankAccount** | Pointer to [**NullableDirectDebitBankAccount**](DirectDebitBankAccount.md) |  | [optional] 
**DebitCard** | Pointer to [**NullableDirectDebitDebitCard**](DirectDebitDebitCard.md) |  | [optional] 

## Methods

### NewPublicDirectDebitAllOf

`func NewPublicDirectDebitAllOf(type_ DirectDebitType, ) *PublicDirectDebitAllOf`

NewPublicDirectDebitAllOf instantiates a new PublicDirectDebitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDirectDebitAllOfWithDefaults

`func NewPublicDirectDebitAllOfWithDefaults() *PublicDirectDebitAllOf`

NewPublicDirectDebitAllOfWithDefaults instantiates a new PublicDirectDebitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicDirectDebitAllOf) GetType() DirectDebitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicDirectDebitAllOf) GetTypeOk() (*DirectDebitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicDirectDebitAllOf) SetType(v DirectDebitType)`

SetType sets Type field to given value.


### GetBankAccount

`func (o *PublicDirectDebitAllOf) GetBankAccount() DirectDebitBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *PublicDirectDebitAllOf) GetBankAccountOk() (*DirectDebitBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *PublicDirectDebitAllOf) SetBankAccount(v DirectDebitBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *PublicDirectDebitAllOf) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### SetBankAccountNil

`func (o *PublicDirectDebitAllOf) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *PublicDirectDebitAllOf) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetDebitCard

`func (o *PublicDirectDebitAllOf) GetDebitCard() DirectDebitDebitCard`

GetDebitCard returns the DebitCard field if non-nil, zero value otherwise.

### GetDebitCardOk

`func (o *PublicDirectDebitAllOf) GetDebitCardOk() (*DirectDebitDebitCard, bool)`

GetDebitCardOk returns a tuple with the DebitCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitCard

`func (o *PublicDirectDebitAllOf) SetDebitCard(v DirectDebitDebitCard)`

SetDebitCard sets DebitCard field to given value.

### HasDebitCard

`func (o *PublicDirectDebitAllOf) HasDebitCard() bool`

HasDebitCard returns a boolean if a field has been set.

### SetDebitCardNil

`func (o *PublicDirectDebitAllOf) SetDebitCardNil(b bool)`

 SetDebitCardNil sets the value for DebitCard to be an explicit nil

### UnsetDebitCard
`func (o *PublicDirectDebitAllOf) UnsetDebitCard()`

UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


