# CustomerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IndividualDetails** | Pointer to [**IndividualDetails**](IndividualDetails.md) |  | [optional] 
**BusinessDetails** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**MobileNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**HashedPhoneNumber** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IdentityAccounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**KycDocuments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerData

`func NewCustomerData() *CustomerData`

NewCustomerData instantiates a new CustomerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDataWithDefaults

`func NewCustomerDataWithDefaults() *CustomerData`

NewCustomerDataWithDefaults instantiates a new CustomerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceId

`func (o *CustomerData) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CustomerData) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CustomerData) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CustomerData) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetType

`func (o *CustomerData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndividualDetails

`func (o *CustomerData) GetIndividualDetails() IndividualDetails`

GetIndividualDetails returns the IndividualDetails field if non-nil, zero value otherwise.

### GetIndividualDetailsOk

`func (o *CustomerData) GetIndividualDetailsOk() (*IndividualDetails, bool)`

GetIndividualDetailsOk returns a tuple with the IndividualDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDetails

`func (o *CustomerData) SetIndividualDetails(v IndividualDetails)`

SetIndividualDetails sets IndividualDetails field to given value.

### HasIndividualDetails

`func (o *CustomerData) HasIndividualDetails() bool`

HasIndividualDetails returns a boolean if a field has been set.

### GetBusinessDetails

`func (o *CustomerData) GetBusinessDetails() string`

GetBusinessDetails returns the BusinessDetails field if non-nil, zero value otherwise.

### GetBusinessDetailsOk

`func (o *CustomerData) GetBusinessDetailsOk() (*string, bool)`

GetBusinessDetailsOk returns a tuple with the BusinessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDetails

`func (o *CustomerData) SetBusinessDetails(v string)`

SetBusinessDetails sets BusinessDetails field to given value.

### HasBusinessDetails

`func (o *CustomerData) HasBusinessDetails() bool`

HasBusinessDetails returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobileNumber

`func (o *CustomerData) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CustomerData) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *CustomerData) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *CustomerData) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetHashedPhoneNumber

`func (o *CustomerData) GetHashedPhoneNumber() string`

GetHashedPhoneNumber returns the HashedPhoneNumber field if non-nil, zero value otherwise.

### GetHashedPhoneNumberOk

`func (o *CustomerData) GetHashedPhoneNumberOk() (*string, bool)`

GetHashedPhoneNumberOk returns a tuple with the HashedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedPhoneNumber

`func (o *CustomerData) SetHashedPhoneNumber(v string)`

SetHashedPhoneNumber sets HashedPhoneNumber field to given value.

### HasHashedPhoneNumber

`func (o *CustomerData) HasHashedPhoneNumber() bool`

HasHashedPhoneNumber returns a boolean if a field has been set.

### GetAddresses

`func (o *CustomerData) GetAddresses() []map[string]interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CustomerData) GetAddressesOk() (*[]map[string]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CustomerData) SetAddresses(v []map[string]interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CustomerData) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetIdentityAccounts

`func (o *CustomerData) GetIdentityAccounts() []map[string]interface{}`

GetIdentityAccounts returns the IdentityAccounts field if non-nil, zero value otherwise.

### GetIdentityAccountsOk

`func (o *CustomerData) GetIdentityAccountsOk() (*[]map[string]interface{}, bool)`

GetIdentityAccountsOk returns a tuple with the IdentityAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccounts

`func (o *CustomerData) SetIdentityAccounts(v []map[string]interface{})`

SetIdentityAccounts sets IdentityAccounts field to given value.

### HasIdentityAccounts

`func (o *CustomerData) HasIdentityAccounts() bool`

HasIdentityAccounts returns a boolean if a field has been set.

### GetKycDocuments

`func (o *CustomerData) GetKycDocuments() []map[string]interface{}`

GetKycDocuments returns the KycDocuments field if non-nil, zero value otherwise.

### GetKycDocumentsOk

`func (o *CustomerData) GetKycDocumentsOk() (*[]map[string]interface{}, bool)`

GetKycDocumentsOk returns a tuple with the KycDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDocuments

`func (o *CustomerData) SetKycDocuments(v []map[string]interface{})`

SetKycDocuments sets KycDocuments field to given value.

### HasKycDocuments

`func (o *CustomerData) HasKycDocuments() bool`

HasKycDocuments returns a boolean if a field has been set.

### GetDescription

`func (o *CustomerData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreated

`func (o *CustomerData) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerData) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerData) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CustomerData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *CustomerData) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomerData) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CustomerData) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CustomerData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


