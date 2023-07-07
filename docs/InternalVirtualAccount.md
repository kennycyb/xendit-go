# InternalVirtualAccount

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
**VirtualAccountId** | Pointer to **string** |  | [optional] 
**AlternativeDisplays** | Pointer to [**[]VirtualAccountAlternativeDisplay**](VirtualAccountAlternativeDisplay.md) |  | [optional] 

## Methods

### NewInternalVirtualAccount

`func NewInternalVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties, ) *InternalVirtualAccount`

NewInternalVirtualAccount instantiates a new InternalVirtualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalVirtualAccountWithDefaults

`func NewInternalVirtualAccountWithDefaults() *InternalVirtualAccount`

NewInternalVirtualAccountWithDefaults instantiates a new InternalVirtualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalVirtualAccount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalVirtualAccount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalVirtualAccount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalVirtualAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InternalVirtualAccount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InternalVirtualAccount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetMinAmount

`func (o *InternalVirtualAccount) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *InternalVirtualAccount) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *InternalVirtualAccount) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *InternalVirtualAccount) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *InternalVirtualAccount) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *InternalVirtualAccount) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *InternalVirtualAccount) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *InternalVirtualAccount) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *InternalVirtualAccount) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *InternalVirtualAccount) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *InternalVirtualAccount) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *InternalVirtualAccount) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetCurrency

`func (o *InternalVirtualAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalVirtualAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalVirtualAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InternalVirtualAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *InternalVirtualAccount) GetChannelCode() VirtualAccountChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalVirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalVirtualAccount) SetChannelCode(v VirtualAccountChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *InternalVirtualAccount) GetChannelProperties() VirtualAccountChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalVirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalVirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAlternativeDisplayTypes

`func (o *InternalVirtualAccount) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *InternalVirtualAccount) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *InternalVirtualAccount) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *InternalVirtualAccount) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.

### GetVirtualAccountId

`func (o *InternalVirtualAccount) GetVirtualAccountId() string`

GetVirtualAccountId returns the VirtualAccountId field if non-nil, zero value otherwise.

### GetVirtualAccountIdOk

`func (o *InternalVirtualAccount) GetVirtualAccountIdOk() (*string, bool)`

GetVirtualAccountIdOk returns a tuple with the VirtualAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountId

`func (o *InternalVirtualAccount) SetVirtualAccountId(v string)`

SetVirtualAccountId sets VirtualAccountId field to given value.

### HasVirtualAccountId

`func (o *InternalVirtualAccount) HasVirtualAccountId() bool`

HasVirtualAccountId returns a boolean if a field has been set.

### GetAlternativeDisplays

`func (o *InternalVirtualAccount) GetAlternativeDisplays() []VirtualAccountAlternativeDisplay`

GetAlternativeDisplays returns the AlternativeDisplays field if non-nil, zero value otherwise.

### GetAlternativeDisplaysOk

`func (o *InternalVirtualAccount) GetAlternativeDisplaysOk() (*[]VirtualAccountAlternativeDisplay, bool)`

GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplays

`func (o *InternalVirtualAccount) SetAlternativeDisplays(v []VirtualAccountAlternativeDisplay)`

SetAlternativeDisplays sets AlternativeDisplays field to given value.

### HasAlternativeDisplays

`func (o *InternalVirtualAccount) HasAlternativeDisplays() bool`

HasAlternativeDisplays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


