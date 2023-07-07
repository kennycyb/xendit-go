# DirectBankTransferChannelProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization is successful | [optional] 
**FailureReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization failed | [optional] 

## Methods

### NewDirectBankTransferChannelProperties

`func NewDirectBankTransferChannelProperties() *DirectBankTransferChannelProperties`

NewDirectBankTransferChannelProperties instantiates a new DirectBankTransferChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectBankTransferChannelPropertiesWithDefaults

`func NewDirectBankTransferChannelPropertiesWithDefaults() *DirectBankTransferChannelProperties`

NewDirectBankTransferChannelPropertiesWithDefaults instantiates a new DirectBankTransferChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedReturnUrl

`func (o *DirectBankTransferChannelProperties) GetAuthorizedReturnUrl() string`

GetAuthorizedReturnUrl returns the AuthorizedReturnUrl field if non-nil, zero value otherwise.

### GetAuthorizedReturnUrlOk

`func (o *DirectBankTransferChannelProperties) GetAuthorizedReturnUrlOk() (*string, bool)`

GetAuthorizedReturnUrlOk returns a tuple with the AuthorizedReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedReturnUrl

`func (o *DirectBankTransferChannelProperties) SetAuthorizedReturnUrl(v string)`

SetAuthorizedReturnUrl sets AuthorizedReturnUrl field to given value.

### HasAuthorizedReturnUrl

`func (o *DirectBankTransferChannelProperties) HasAuthorizedReturnUrl() bool`

HasAuthorizedReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *DirectBankTransferChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *DirectBankTransferChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *DirectBankTransferChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *DirectBankTransferChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


