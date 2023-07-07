# ChannelWithCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to [**PaymentMethodCountry**](PaymentMethodCountry.md) |  | [optional] 

## Methods

### NewChannelWithCountry

`func NewChannelWithCountry() *ChannelWithCountry`

NewChannelWithCountry instantiates a new ChannelWithCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithCountryWithDefaults

`func NewChannelWithCountryWithDefaults() *ChannelWithCountry`

NewChannelWithCountryWithDefaults instantiates a new ChannelWithCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *ChannelWithCountry) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *ChannelWithCountry) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *ChannelWithCountry) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *ChannelWithCountry) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *ChannelWithCountry) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *ChannelWithCountry) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetCountry

`func (o *ChannelWithCountry) GetCountry() PaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ChannelWithCountry) GetCountryOk() (*PaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ChannelWithCountry) SetCountry(v PaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ChannelWithCountry) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


