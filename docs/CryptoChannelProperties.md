# CryptoChannelProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessReturnUrl** | Pointer to **NullableString** | URL where the end-customer is redirected if the authorization is successful | [optional] 
**FailureReturnUrl** | Pointer to **NullableString** | URL where the end-customer is redirected if the authorization failed | [optional] 
**CancelReturnUrl** | Pointer to **NullableString** | URL where the end-customer is redirected if the authorization cancelled | [optional] 

## Methods

### NewCryptoChannelProperties

`func NewCryptoChannelProperties() *CryptoChannelProperties`

NewCryptoChannelProperties instantiates a new CryptoChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoChannelPropertiesWithDefaults

`func NewCryptoChannelPropertiesWithDefaults() *CryptoChannelProperties`

NewCryptoChannelPropertiesWithDefaults instantiates a new CryptoChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *CryptoChannelProperties) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *CryptoChannelProperties) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *CryptoChannelProperties) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *CryptoChannelProperties) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### SetSuccessReturnUrlNil

`func (o *CryptoChannelProperties) SetSuccessReturnUrlNil(b bool)`

 SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil

### UnsetSuccessReturnUrl
`func (o *CryptoChannelProperties) UnsetSuccessReturnUrl()`

UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
### GetFailureReturnUrl

`func (o *CryptoChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *CryptoChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *CryptoChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *CryptoChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### SetFailureReturnUrlNil

`func (o *CryptoChannelProperties) SetFailureReturnUrlNil(b bool)`

 SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil

### UnsetFailureReturnUrl
`func (o *CryptoChannelProperties) UnsetFailureReturnUrl()`

UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
### GetCancelReturnUrl

`func (o *CryptoChannelProperties) GetCancelReturnUrl() string`

GetCancelReturnUrl returns the CancelReturnUrl field if non-nil, zero value otherwise.

### GetCancelReturnUrlOk

`func (o *CryptoChannelProperties) GetCancelReturnUrlOk() (*string, bool)`

GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReturnUrl

`func (o *CryptoChannelProperties) SetCancelReturnUrl(v string)`

SetCancelReturnUrl sets CancelReturnUrl field to given value.

### HasCancelReturnUrl

`func (o *CryptoChannelProperties) HasCancelReturnUrl() bool`

HasCancelReturnUrl returns a boolean if a field has been set.

### SetCancelReturnUrlNil

`func (o *CryptoChannelProperties) SetCancelReturnUrlNil(b bool)`

 SetCancelReturnUrlNil sets the value for CancelReturnUrl to be an explicit nil

### UnsetCancelReturnUrl
`func (o *CryptoChannelProperties) UnsetCancelReturnUrl()`

UnsetCancelReturnUrl ensures that no value is present for CancelReturnUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


