# InternalCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **NullableString** |  | 
**ChannelProperties** | [**NullableCardChannelProperties**](CardChannelProperties.md) |  | 
**CardInformation** | Pointer to [**TokenizedCardInformation**](TokenizedCardInformation.md) |  | [optional] 
**CardVerificationResults** | Pointer to [**NullableCardVerificationResults**](CardVerificationResults.md) |  | [optional] 

## Methods

### NewInternalCard

`func NewInternalCard(currency NullableString, channelProperties NullableCardChannelProperties, ) *InternalCard`

NewInternalCard instantiates a new InternalCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalCardWithDefaults

`func NewInternalCardWithDefaults() *InternalCard`

NewInternalCardWithDefaults instantiates a new InternalCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InternalCard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalCard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalCard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *InternalCard) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InternalCard) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetChannelProperties

`func (o *InternalCard) GetChannelProperties() CardChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalCard) GetChannelPropertiesOk() (*CardChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalCard) SetChannelProperties(v CardChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *InternalCard) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *InternalCard) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetCardInformation

`func (o *InternalCard) GetCardInformation() TokenizedCardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *InternalCard) GetCardInformationOk() (*TokenizedCardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *InternalCard) SetCardInformation(v TokenizedCardInformation)`

SetCardInformation sets CardInformation field to given value.

### HasCardInformation

`func (o *InternalCard) HasCardInformation() bool`

HasCardInformation returns a boolean if a field has been set.

### GetCardVerificationResults

`func (o *InternalCard) GetCardVerificationResults() CardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *InternalCard) GetCardVerificationResultsOk() (*CardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *InternalCard) SetCardVerificationResults(v CardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *InternalCard) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *InternalCard) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *InternalCard) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


