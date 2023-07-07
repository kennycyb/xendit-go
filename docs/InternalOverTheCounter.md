# InternalOverTheCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ChannelCode** | [**OverTheCounterChannelCode**](OverTheCounterChannelCode.md) |  | 
**ChannelProperties** | [**OverTheCounterChannelProperties**](OverTheCounterChannelProperties.md) |  | 
**PaymentCodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalOverTheCounter

`func NewInternalOverTheCounter(channelCode OverTheCounterChannelCode, channelProperties OverTheCounterChannelProperties, ) *InternalOverTheCounter`

NewInternalOverTheCounter instantiates a new InternalOverTheCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalOverTheCounterWithDefaults

`func NewInternalOverTheCounterWithDefaults() *InternalOverTheCounter`

NewInternalOverTheCounterWithDefaults instantiates a new InternalOverTheCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalOverTheCounter) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalOverTheCounter) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalOverTheCounter) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalOverTheCounter) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InternalOverTheCounter) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InternalOverTheCounter) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *InternalOverTheCounter) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalOverTheCounter) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalOverTheCounter) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InternalOverTheCounter) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *InternalOverTheCounter) GetChannelCode() OverTheCounterChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalOverTheCounter) GetChannelCodeOk() (*OverTheCounterChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalOverTheCounter) SetChannelCode(v OverTheCounterChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *InternalOverTheCounter) GetChannelProperties() OverTheCounterChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalOverTheCounter) GetChannelPropertiesOk() (*OverTheCounterChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalOverTheCounter) SetChannelProperties(v OverTheCounterChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetPaymentCodeId

`func (o *InternalOverTheCounter) GetPaymentCodeId() string`

GetPaymentCodeId returns the PaymentCodeId field if non-nil, zero value otherwise.

### GetPaymentCodeIdOk

`func (o *InternalOverTheCounter) GetPaymentCodeIdOk() (*string, bool)`

GetPaymentCodeIdOk returns a tuple with the PaymentCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCodeId

`func (o *InternalOverTheCounter) SetPaymentCodeId(v string)`

SetPaymentCodeId sets PaymentCodeId field to given value.

### HasPaymentCodeId

`func (o *InternalOverTheCounter) HasPaymentCodeId() bool`

HasPaymentCodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


