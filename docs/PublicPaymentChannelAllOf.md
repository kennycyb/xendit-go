# PublicPaymentChannelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to **string** | The specific Xendit code used to identify the partner channel | [optional] 
**Type** | Pointer to **string** | The payment method type | [optional] 
**Country** | Pointer to **string** | The country where the channel operates  in ISO 3166-2 country code | [optional] 
**ChannelName** | Pointer to **string** | Official parter name of the payment channel | [optional] 
**ChannelProperties** | Pointer to [**[]ChannelProperty**](ChannelProperty.md) |  | [optional] 
**LogoUrl** | Pointer to **string** | If available, this contains a URL to the logo of the partner channel | [optional] 
**AmountLimits** | Pointer to [**[]ChannelAmountLimits**](ChannelAmountLimits.md) |  | [optional] 

## Methods

### NewPublicPaymentChannelAllOf

`func NewPublicPaymentChannelAllOf() *PublicPaymentChannelAllOf`

NewPublicPaymentChannelAllOf instantiates a new PublicPaymentChannelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentChannelAllOfWithDefaults

`func NewPublicPaymentChannelAllOfWithDefaults() *PublicPaymentChannelAllOf`

NewPublicPaymentChannelAllOfWithDefaults instantiates a new PublicPaymentChannelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PublicPaymentChannelAllOf) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PublicPaymentChannelAllOf) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PublicPaymentChannelAllOf) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *PublicPaymentChannelAllOf) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetType

`func (o *PublicPaymentChannelAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicPaymentChannelAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicPaymentChannelAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicPaymentChannelAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *PublicPaymentChannelAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicPaymentChannelAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicPaymentChannelAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicPaymentChannelAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetChannelName

`func (o *PublicPaymentChannelAllOf) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *PublicPaymentChannelAllOf) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *PublicPaymentChannelAllOf) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *PublicPaymentChannelAllOf) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetChannelProperties

`func (o *PublicPaymentChannelAllOf) GetChannelProperties() []ChannelProperty`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicPaymentChannelAllOf) GetChannelPropertiesOk() (*[]ChannelProperty, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicPaymentChannelAllOf) SetChannelProperties(v []ChannelProperty)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PublicPaymentChannelAllOf) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetLogoUrl

`func (o *PublicPaymentChannelAllOf) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *PublicPaymentChannelAllOf) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *PublicPaymentChannelAllOf) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *PublicPaymentChannelAllOf) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetAmountLimits

`func (o *PublicPaymentChannelAllOf) GetAmountLimits() []ChannelAmountLimits`

GetAmountLimits returns the AmountLimits field if non-nil, zero value otherwise.

### GetAmountLimitsOk

`func (o *PublicPaymentChannelAllOf) GetAmountLimitsOk() (*[]ChannelAmountLimits, bool)`

GetAmountLimitsOk returns a tuple with the AmountLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimits

`func (o *PublicPaymentChannelAllOf) SetAmountLimits(v []ChannelAmountLimits)`

SetAmountLimits sets AmountLimits field to given value.

### HasAmountLimits

`func (o *PublicPaymentChannelAllOf) HasAmountLimits() bool`

HasAmountLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


