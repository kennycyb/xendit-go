# MutableCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**ChannelProperties** | Pointer to [**NullableCardChannelProperties**](CardChannelProperties.md) |  | [optional] 
**CardInformation** | Pointer to [**TokenizedCardInformation**](TokenizedCardInformation.md) |  | [optional] 
**CardVerificationResults** | Pointer to [**NullableCardVerificationResults**](CardVerificationResults.md) |  | [optional] 
**Token** | Pointer to [**CardToken**](CardToken.md) |  | [optional] 

## Methods

### NewMutableCard

`func NewMutableCard(currency string, ) *MutableCard`

NewMutableCard instantiates a new MutableCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableCardWithDefaults

`func NewMutableCardWithDefaults() *MutableCard`

NewMutableCardWithDefaults instantiates a new MutableCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *MutableCard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MutableCard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MutableCard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChannelProperties

`func (o *MutableCard) GetChannelProperties() CardChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *MutableCard) GetChannelPropertiesOk() (*CardChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *MutableCard) SetChannelProperties(v CardChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *MutableCard) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *MutableCard) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *MutableCard) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetCardInformation

`func (o *MutableCard) GetCardInformation() TokenizedCardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *MutableCard) GetCardInformationOk() (*TokenizedCardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *MutableCard) SetCardInformation(v TokenizedCardInformation)`

SetCardInformation sets CardInformation field to given value.

### HasCardInformation

`func (o *MutableCard) HasCardInformation() bool`

HasCardInformation returns a boolean if a field has been set.

### GetCardVerificationResults

`func (o *MutableCard) GetCardVerificationResults() CardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *MutableCard) GetCardVerificationResultsOk() (*CardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *MutableCard) SetCardVerificationResults(v CardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *MutableCard) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *MutableCard) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *MutableCard) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
### GetToken

`func (o *MutableCard) GetToken() CardToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MutableCard) GetTokenOk() (*CardToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MutableCard) SetToken(v CardToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *MutableCard) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


