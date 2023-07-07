# MutableDirectDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**DirectDebitChannelCode**](DirectDebitChannelCode.md) |  | 
**ChannelProperties** | [**NullableDirectDebitChannelProperties**](DirectDebitChannelProperties.md) |  | 

## Methods

### NewMutableDirectDebit

`func NewMutableDirectDebit(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, ) *MutableDirectDebit`

NewMutableDirectDebit instantiates a new MutableDirectDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableDirectDebitWithDefaults

`func NewMutableDirectDebitWithDefaults() *MutableDirectDebit`

NewMutableDirectDebitWithDefaults instantiates a new MutableDirectDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *MutableDirectDebit) GetChannelCode() DirectDebitChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *MutableDirectDebit) GetChannelCodeOk() (*DirectDebitChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *MutableDirectDebit) SetChannelCode(v DirectDebitChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *MutableDirectDebit) GetChannelProperties() DirectDebitChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *MutableDirectDebit) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *MutableDirectDebit) SetChannelProperties(v DirectDebitChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *MutableDirectDebit) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *MutableDirectDebit) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


