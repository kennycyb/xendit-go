# PublicOverTheCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ChannelCode** | [**OverTheCounterChannelCode**](OverTheCounterChannelCode.md) |  | 
**ChannelProperties** | [**OverTheCounterChannelProperties**](OverTheCounterChannelProperties.md) |  | 

## Methods

### NewPublicOverTheCounter

`func NewPublicOverTheCounter(channelCode OverTheCounterChannelCode, channelProperties OverTheCounterChannelProperties, ) *PublicOverTheCounter`

NewPublicOverTheCounter instantiates a new PublicOverTheCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicOverTheCounterWithDefaults

`func NewPublicOverTheCounterWithDefaults() *PublicOverTheCounter`

NewPublicOverTheCounterWithDefaults instantiates a new PublicOverTheCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PublicOverTheCounter) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicOverTheCounter) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicOverTheCounter) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PublicOverTheCounter) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PublicOverTheCounter) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PublicOverTheCounter) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *PublicOverTheCounter) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PublicOverTheCounter) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PublicOverTheCounter) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PublicOverTheCounter) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *PublicOverTheCounter) GetChannelCode() OverTheCounterChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicOverTheCounter) GetChannelCodeOk() (*OverTheCounterChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicOverTheCounter) SetChannelCode(v OverTheCounterChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *PublicOverTheCounter) GetChannelProperties() OverTheCounterChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicOverTheCounter) GetChannelPropertiesOk() (*OverTheCounterChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicOverTheCounter) SetChannelProperties(v OverTheCounterChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


