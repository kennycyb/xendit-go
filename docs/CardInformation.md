# CardInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **string** |  | [optional] 
**MaskedCardNumber** | Pointer to **string** | 1st 6 and last 4 digits of the card | [optional] 
**ExpiryMonth** | Pointer to **string** | Card expiry month in MM format | [optional] 
**ExpiryYear** | Pointer to **string** | Card expiry month in YY format | [optional] 
**CardholderName** | Pointer to **NullableString** | Cardholder name | [optional] 

## Methods

### NewCardInformation

`func NewCardInformation() *CardInformation`

NewCardInformation instantiates a new CardInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInformationWithDefaults

`func NewCardInformationWithDefaults() *CardInformation`

NewCardInformationWithDefaults instantiates a new CardInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *CardInformation) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CardInformation) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CardInformation) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CardInformation) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetMaskedCardNumber

`func (o *CardInformation) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *CardInformation) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *CardInformation) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.

### HasMaskedCardNumber

`func (o *CardInformation) HasMaskedCardNumber() bool`

HasMaskedCardNumber returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *CardInformation) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CardInformation) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CardInformation) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *CardInformation) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *CardInformation) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CardInformation) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CardInformation) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *CardInformation) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetCardholderName

`func (o *CardInformation) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *CardInformation) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *CardInformation) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *CardInformation) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *CardInformation) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *CardInformation) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


