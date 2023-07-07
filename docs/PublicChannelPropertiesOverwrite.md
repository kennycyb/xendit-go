# PublicChannelPropertiesOverwrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization is successful | [optional] 
**FailureReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization failed | [optional] 
**CancelReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization cancelled | [optional] 
**RedeemPoints** | Pointer to **string** | REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only. | [optional] 
**RequireAuth** | Pointer to **bool** | Toggle used to require end-customer to input undergo OTP validation before completing a payment. OTP will always be required for transactions greater than 1,000,000 IDR. For BRI tokenized payment use only. | [optional] 
**MerchantIdTag** | Pointer to **string** | Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use | [optional] 
**CardonfileType** | Pointer to **string** | Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging. | [optional] 

## Methods

### NewPublicChannelPropertiesOverwrite

`func NewPublicChannelPropertiesOverwrite() *PublicChannelPropertiesOverwrite`

NewPublicChannelPropertiesOverwrite instantiates a new PublicChannelPropertiesOverwrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicChannelPropertiesOverwriteWithDefaults

`func NewPublicChannelPropertiesOverwriteWithDefaults() *PublicChannelPropertiesOverwrite`

NewPublicChannelPropertiesOverwriteWithDefaults instantiates a new PublicChannelPropertiesOverwrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *PublicChannelPropertiesOverwrite) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *PublicChannelPropertiesOverwrite) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *PublicChannelPropertiesOverwrite) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *PublicChannelPropertiesOverwrite) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *PublicChannelPropertiesOverwrite) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *PublicChannelPropertiesOverwrite) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *PublicChannelPropertiesOverwrite) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *PublicChannelPropertiesOverwrite) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### GetCancelReturnUrl

`func (o *PublicChannelPropertiesOverwrite) GetCancelReturnUrl() string`

GetCancelReturnUrl returns the CancelReturnUrl field if non-nil, zero value otherwise.

### GetCancelReturnUrlOk

`func (o *PublicChannelPropertiesOverwrite) GetCancelReturnUrlOk() (*string, bool)`

GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReturnUrl

`func (o *PublicChannelPropertiesOverwrite) SetCancelReturnUrl(v string)`

SetCancelReturnUrl sets CancelReturnUrl field to given value.

### HasCancelReturnUrl

`func (o *PublicChannelPropertiesOverwrite) HasCancelReturnUrl() bool`

HasCancelReturnUrl returns a boolean if a field has been set.

### GetRedeemPoints

`func (o *PublicChannelPropertiesOverwrite) GetRedeemPoints() string`

GetRedeemPoints returns the RedeemPoints field if non-nil, zero value otherwise.

### GetRedeemPointsOk

`func (o *PublicChannelPropertiesOverwrite) GetRedeemPointsOk() (*string, bool)`

GetRedeemPointsOk returns a tuple with the RedeemPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPoints

`func (o *PublicChannelPropertiesOverwrite) SetRedeemPoints(v string)`

SetRedeemPoints sets RedeemPoints field to given value.

### HasRedeemPoints

`func (o *PublicChannelPropertiesOverwrite) HasRedeemPoints() bool`

HasRedeemPoints returns a boolean if a field has been set.

### GetRequireAuth

`func (o *PublicChannelPropertiesOverwrite) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *PublicChannelPropertiesOverwrite) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *PublicChannelPropertiesOverwrite) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *PublicChannelPropertiesOverwrite) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### GetMerchantIdTag

`func (o *PublicChannelPropertiesOverwrite) GetMerchantIdTag() string`

GetMerchantIdTag returns the MerchantIdTag field if non-nil, zero value otherwise.

### GetMerchantIdTagOk

`func (o *PublicChannelPropertiesOverwrite) GetMerchantIdTagOk() (*string, bool)`

GetMerchantIdTagOk returns a tuple with the MerchantIdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdTag

`func (o *PublicChannelPropertiesOverwrite) SetMerchantIdTag(v string)`

SetMerchantIdTag sets MerchantIdTag field to given value.

### HasMerchantIdTag

`func (o *PublicChannelPropertiesOverwrite) HasMerchantIdTag() bool`

HasMerchantIdTag returns a boolean if a field has been set.

### GetCardonfileType

`func (o *PublicChannelPropertiesOverwrite) GetCardonfileType() string`

GetCardonfileType returns the CardonfileType field if non-nil, zero value otherwise.

### GetCardonfileTypeOk

`func (o *PublicChannelPropertiesOverwrite) GetCardonfileTypeOk() (*string, bool)`

GetCardonfileTypeOk returns a tuple with the CardonfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardonfileType

`func (o *PublicChannelPropertiesOverwrite) SetCardonfileType(v string)`

SetCardonfileType sets CardonfileType field to given value.

### HasCardonfileType

`func (o *PublicChannelPropertiesOverwrite) HasCardonfileType() bool`

HasCardonfileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


