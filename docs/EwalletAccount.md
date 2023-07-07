# EwalletAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the eWallet account holder. The value is null if unavailableName of the eWallet account holder. The value is null if unavailable | [optional] 
**AccountDetails** | Pointer to **NullableString** | Identifier from eWallet provider e.g. phone number. The value is null if unavailable | [optional] 
**Balance** | Pointer to **NullableFloat64** | The main balance amount on eWallet account provided from eWallet provider. The value is null if unavailable | [optional] 
**PointBalance** | Pointer to **NullableFloat64** | The point balance amount on eWallet account. Applicable only on some eWallet provider that has point system. The value is null if unavailabl | [optional] 

## Methods

### NewEwalletAccount

`func NewEwalletAccount() *EwalletAccount`

NewEwalletAccount instantiates a new EwalletAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEwalletAccountWithDefaults

`func NewEwalletAccountWithDefaults() *EwalletAccount`

NewEwalletAccountWithDefaults instantiates a new EwalletAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EwalletAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EwalletAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EwalletAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EwalletAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EwalletAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EwalletAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccountDetails

`func (o *EwalletAccount) GetAccountDetails() string`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *EwalletAccount) GetAccountDetailsOk() (*string, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *EwalletAccount) SetAccountDetails(v string)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *EwalletAccount) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### SetAccountDetailsNil

`func (o *EwalletAccount) SetAccountDetailsNil(b bool)`

 SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil

### UnsetAccountDetails
`func (o *EwalletAccount) UnsetAccountDetails()`

UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
### GetBalance

`func (o *EwalletAccount) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EwalletAccount) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EwalletAccount) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EwalletAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *EwalletAccount) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *EwalletAccount) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetPointBalance

`func (o *EwalletAccount) GetPointBalance() float64`

GetPointBalance returns the PointBalance field if non-nil, zero value otherwise.

### GetPointBalanceOk

`func (o *EwalletAccount) GetPointBalanceOk() (*float64, bool)`

GetPointBalanceOk returns a tuple with the PointBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointBalance

`func (o *EwalletAccount) SetPointBalance(v float64)`

SetPointBalance sets PointBalance field to given value.

### HasPointBalance

`func (o *EwalletAccount) HasPointBalance() bool`

HasPointBalance returns a boolean if a field has been set.

### SetPointBalanceNil

`func (o *EwalletAccount) SetPointBalanceNil(b bool)`

 SetPointBalanceNil sets the value for PointBalance to be an explicit nil

### UnsetPointBalance
`func (o *EwalletAccount) UnsetPointBalance()`

UnsetPointBalance ensures that no value is present for PointBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


