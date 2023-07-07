# MutableVirtualAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**MinAmount** | Pointer to **NullableFloat64** |  | [optional] 
**MaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ChannelCode** | [**VirtualAccountChannelCode**](VirtualAccountChannelCode.md) |  | 
**ChannelProperties** | [**VirtualAccountChannelProperties**](VirtualAccountChannelProperties.md) |  | 
**AlternativeDisplayTypes** | Pointer to **[]string** | Alternative display requested for the virtual account | [optional] 

## Methods

### NewMutableVirtualAccount

`func NewMutableVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties, ) *MutableVirtualAccount`

NewMutableVirtualAccount instantiates a new MutableVirtualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableVirtualAccountWithDefaults

`func NewMutableVirtualAccountWithDefaults() *MutableVirtualAccount`

NewMutableVirtualAccountWithDefaults instantiates a new MutableVirtualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MutableVirtualAccount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MutableVirtualAccount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MutableVirtualAccount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MutableVirtualAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *MutableVirtualAccount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *MutableVirtualAccount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetMinAmount

`func (o *MutableVirtualAccount) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *MutableVirtualAccount) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *MutableVirtualAccount) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *MutableVirtualAccount) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *MutableVirtualAccount) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *MutableVirtualAccount) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *MutableVirtualAccount) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *MutableVirtualAccount) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *MutableVirtualAccount) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *MutableVirtualAccount) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *MutableVirtualAccount) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *MutableVirtualAccount) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetCurrency

`func (o *MutableVirtualAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MutableVirtualAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MutableVirtualAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MutableVirtualAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *MutableVirtualAccount) GetChannelCode() VirtualAccountChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *MutableVirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *MutableVirtualAccount) SetChannelCode(v VirtualAccountChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *MutableVirtualAccount) GetChannelProperties() VirtualAccountChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *MutableVirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *MutableVirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAlternativeDisplayTypes

`func (o *MutableVirtualAccount) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *MutableVirtualAccount) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *MutableVirtualAccount) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *MutableVirtualAccount) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


