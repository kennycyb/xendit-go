# InternalQRCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ChannelCode** | Pointer to [**NullableQRCodeChannelCode**](QRCodeChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**NullableQRCodeChannelProperties**](QRCodeChannelProperties.md) |  | [optional] 
**QrCodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalQRCode

`func NewInternalQRCode() *InternalQRCode`

NewInternalQRCode instantiates a new InternalQRCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalQRCodeWithDefaults

`func NewInternalQRCodeWithDefaults() *InternalQRCode`

NewInternalQRCodeWithDefaults instantiates a new InternalQRCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalQRCode) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalQRCode) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalQRCode) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalQRCode) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InternalQRCode) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InternalQRCode) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *InternalQRCode) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalQRCode) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalQRCode) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InternalQRCode) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *InternalQRCode) GetChannelCode() QRCodeChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalQRCode) GetChannelCodeOk() (*QRCodeChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalQRCode) SetChannelCode(v QRCodeChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *InternalQRCode) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *InternalQRCode) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *InternalQRCode) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *InternalQRCode) GetChannelProperties() QRCodeChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalQRCode) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalQRCode) SetChannelProperties(v QRCodeChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *InternalQRCode) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *InternalQRCode) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *InternalQRCode) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetQrCodeId

`func (o *InternalQRCode) GetQrCodeId() string`

GetQrCodeId returns the QrCodeId field if non-nil, zero value otherwise.

### GetQrCodeIdOk

`func (o *InternalQRCode) GetQrCodeIdOk() (*string, bool)`

GetQrCodeIdOk returns a tuple with the QrCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeId

`func (o *InternalQRCode) SetQrCodeId(v string)`

SetQrCodeId sets QrCodeId field to given value.

### HasQrCodeId

`func (o *InternalQRCode) HasQrCodeId() bool`

HasQrCodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


