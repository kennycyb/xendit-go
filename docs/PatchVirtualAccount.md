# PatchVirtualAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**MinAmount** | Pointer to **NullableFloat64** |  | [optional] 
**MaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**ChannelProperties** | Pointer to [**VirtualAccountChannelPropertiesPatch**](VirtualAccountChannelPropertiesPatch.md) |  | [optional] 
**AlternativeDisplayTypes** | Pointer to **[]string** | Alternative display requested for the virtual account | [optional] 

## Methods

### NewPatchVirtualAccount

`func NewPatchVirtualAccount() *PatchVirtualAccount`

NewPatchVirtualAccount instantiates a new PatchVirtualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchVirtualAccountWithDefaults

`func NewPatchVirtualAccountWithDefaults() *PatchVirtualAccount`

NewPatchVirtualAccountWithDefaults instantiates a new PatchVirtualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PatchVirtualAccount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PatchVirtualAccount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PatchVirtualAccount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PatchVirtualAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PatchVirtualAccount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PatchVirtualAccount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetMinAmount

`func (o *PatchVirtualAccount) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *PatchVirtualAccount) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *PatchVirtualAccount) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *PatchVirtualAccount) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *PatchVirtualAccount) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *PatchVirtualAccount) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *PatchVirtualAccount) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PatchVirtualAccount) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PatchVirtualAccount) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *PatchVirtualAccount) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *PatchVirtualAccount) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *PatchVirtualAccount) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetChannelProperties

`func (o *PatchVirtualAccount) GetChannelProperties() VirtualAccountChannelPropertiesPatch`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PatchVirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelPropertiesPatch, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PatchVirtualAccount) SetChannelProperties(v VirtualAccountChannelPropertiesPatch)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PatchVirtualAccount) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAlternativeDisplayTypes

`func (o *PatchVirtualAccount) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *PatchVirtualAccount) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *PatchVirtualAccount) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *PatchVirtualAccount) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


