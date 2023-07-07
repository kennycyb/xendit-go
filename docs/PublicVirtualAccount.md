# PublicVirtualAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  | [optional] 
**ChannelCode** | [**VirtualAccountChannelCode**](VirtualAccountChannelCode.md) |  | 
**ChannelProperties** | [**VirtualAccountChannelProperties**](VirtualAccountChannelProperties.md) |  | 
**AlternativeDisplayTypes** | Pointer to **[]string** | Alternative display requested for the virtual account | [optional] 
**AlternativeDisplays** | Pointer to [**[]VirtualAccountAlternativeDisplay**](VirtualAccountAlternativeDisplay.md) |  | [optional] 

## Methods

### NewPublicVirtualAccount

`func NewPublicVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties, ) *PublicVirtualAccount`

NewPublicVirtualAccount instantiates a new PublicVirtualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicVirtualAccountWithDefaults

`func NewPublicVirtualAccountWithDefaults() *PublicVirtualAccount`

NewPublicVirtualAccountWithDefaults instantiates a new PublicVirtualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PublicVirtualAccount) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicVirtualAccount) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicVirtualAccount) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PublicVirtualAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PublicVirtualAccount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PublicVirtualAccount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *PublicVirtualAccount) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PublicVirtualAccount) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PublicVirtualAccount) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PublicVirtualAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *PublicVirtualAccount) GetChannelCode() VirtualAccountChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicVirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicVirtualAccount) SetChannelCode(v VirtualAccountChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *PublicVirtualAccount) GetChannelProperties() VirtualAccountChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicVirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicVirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAlternativeDisplayTypes

`func (o *PublicVirtualAccount) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *PublicVirtualAccount) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *PublicVirtualAccount) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *PublicVirtualAccount) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.

### GetAlternativeDisplays

`func (o *PublicVirtualAccount) GetAlternativeDisplays() []VirtualAccountAlternativeDisplay`

GetAlternativeDisplays returns the AlternativeDisplays field if non-nil, zero value otherwise.

### GetAlternativeDisplaysOk

`func (o *PublicVirtualAccount) GetAlternativeDisplaysOk() (*[]VirtualAccountAlternativeDisplay, bool)`

GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplays

`func (o *PublicVirtualAccount) SetAlternativeDisplays(v []VirtualAccountAlternativeDisplay)`

SetAlternativeDisplays sets AlternativeDisplays field to given value.

### HasAlternativeDisplays

`func (o *PublicVirtualAccount) HasAlternativeDisplays() bool`

HasAlternativeDisplays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


