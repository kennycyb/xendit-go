# PatchOverTheCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**ChannelProperties** | Pointer to [**OverTheCounterChannelPropertiesPatch**](OverTheCounterChannelPropertiesPatch.md) |  | [optional] 

## Methods

### NewPatchOverTheCounter

`func NewPatchOverTheCounter() *PatchOverTheCounter`

NewPatchOverTheCounter instantiates a new PatchOverTheCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOverTheCounterWithDefaults

`func NewPatchOverTheCounterWithDefaults() *PatchOverTheCounter`

NewPatchOverTheCounterWithDefaults instantiates a new PatchOverTheCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PatchOverTheCounter) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PatchOverTheCounter) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PatchOverTheCounter) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PatchOverTheCounter) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PatchOverTheCounter) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PatchOverTheCounter) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetChannelProperties

`func (o *PatchOverTheCounter) GetChannelProperties() OverTheCounterChannelPropertiesPatch`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PatchOverTheCounter) GetChannelPropertiesOk() (*OverTheCounterChannelPropertiesPatch, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PatchOverTheCounter) SetChannelProperties(v OverTheCounterChannelPropertiesPatch)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PatchOverTheCounter) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


