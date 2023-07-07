# PublicQRCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ChannelCode** | Pointer to [**NullableQRCodeChannelCode**](QRCodeChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**NullableQRCodeChannelProperties**](QRCodeChannelProperties.md) |  | [optional] 

## Methods

### NewPublicQRCode

`func NewPublicQRCode() *PublicQRCode`

NewPublicQRCode instantiates a new PublicQRCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicQRCodeWithDefaults

`func NewPublicQRCodeWithDefaults() *PublicQRCode`

NewPublicQRCodeWithDefaults instantiates a new PublicQRCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PublicQRCode) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicQRCode) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicQRCode) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PublicQRCode) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PublicQRCode) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PublicQRCode) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *PublicQRCode) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PublicQRCode) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PublicQRCode) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PublicQRCode) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *PublicQRCode) GetChannelCode() QRCodeChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicQRCode) GetChannelCodeOk() (*QRCodeChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicQRCode) SetChannelCode(v QRCodeChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *PublicQRCode) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *PublicQRCode) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *PublicQRCode) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *PublicQRCode) GetChannelProperties() QRCodeChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicQRCode) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicQRCode) SetChannelProperties(v QRCodeChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PublicQRCode) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *PublicQRCode) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *PublicQRCode) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


