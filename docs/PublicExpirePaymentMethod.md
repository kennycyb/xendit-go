# PublicExpirePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessReturnUrl** | Pointer to **NullableString** | URL where the end customer is redirected if the unlinking authorization is successful. | [optional] 
**FailureReturnUrl** | Pointer to **NullableString** | URL where the end customer is redirected if the unlinking authorization is failed. | [optional] 

## Methods

### NewPublicExpirePaymentMethod

`func NewPublicExpirePaymentMethod() *PublicExpirePaymentMethod`

NewPublicExpirePaymentMethod instantiates a new PublicExpirePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExpirePaymentMethodWithDefaults

`func NewPublicExpirePaymentMethodWithDefaults() *PublicExpirePaymentMethod`

NewPublicExpirePaymentMethodWithDefaults instantiates a new PublicExpirePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *PublicExpirePaymentMethod) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *PublicExpirePaymentMethod) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *PublicExpirePaymentMethod) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *PublicExpirePaymentMethod) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### SetSuccessReturnUrlNil

`func (o *PublicExpirePaymentMethod) SetSuccessReturnUrlNil(b bool)`

 SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil

### UnsetSuccessReturnUrl
`func (o *PublicExpirePaymentMethod) UnsetSuccessReturnUrl()`

UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
### GetFailureReturnUrl

`func (o *PublicExpirePaymentMethod) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *PublicExpirePaymentMethod) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *PublicExpirePaymentMethod) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *PublicExpirePaymentMethod) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### SetFailureReturnUrlNil

`func (o *PublicExpirePaymentMethod) SetFailureReturnUrlNil(b bool)`

 SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil

### UnsetFailureReturnUrl
`func (o *PublicExpirePaymentMethod) UnsetFailureReturnUrl()`

UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


