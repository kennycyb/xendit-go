# CardBillingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**AddressLine1** | Pointer to **NullableString** |  | [optional] 
**AddressLine2** | Pointer to **NullableString** |  | [optional] 
**AddressCity** | Pointer to **NullableString** |  | [optional] 
**AddressCountry** | Pointer to **NullableString** |  | [optional] 
**AddressPostalCode** | Pointer to **NullableString** |  | [optional] 
**AddressState** | Pointer to **NullableString** |  | [optional] 
**GivenNames** | Pointer to **NullableString** |  | [optional] 
**MiddleName** | Pointer to **NullableString** |  | [optional] 
**Surname** | Pointer to **NullableString** |  | [optional] 
**MobileNumber** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**NullableCardBillingDetailsAddress**](CardBillingDetailsAddress.md) |  | [optional] 

## Methods

### NewCardBillingDetails

`func NewCardBillingDetails() *CardBillingDetails`

NewCardBillingDetails instantiates a new CardBillingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardBillingDetailsWithDefaults

`func NewCardBillingDetailsWithDefaults() *CardBillingDetails`

NewCardBillingDetailsWithDefaults instantiates a new CardBillingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CardBillingDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardBillingDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardBillingDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CardBillingDetails) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CardBillingDetails) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CardBillingDetails) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CardBillingDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardBillingDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardBillingDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CardBillingDetails) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CardBillingDetails) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CardBillingDetails) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *CardBillingDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CardBillingDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CardBillingDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CardBillingDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CardBillingDetails) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CardBillingDetails) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *CardBillingDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardBillingDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardBillingDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardBillingDetails) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CardBillingDetails) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CardBillingDetails) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAddressLine1

`func (o *CardBillingDetails) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CardBillingDetails) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CardBillingDetails) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CardBillingDetails) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *CardBillingDetails) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *CardBillingDetails) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *CardBillingDetails) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CardBillingDetails) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CardBillingDetails) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CardBillingDetails) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *CardBillingDetails) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *CardBillingDetails) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAddressCity

`func (o *CardBillingDetails) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *CardBillingDetails) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *CardBillingDetails) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *CardBillingDetails) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *CardBillingDetails) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *CardBillingDetails) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressCountry

`func (o *CardBillingDetails) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *CardBillingDetails) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *CardBillingDetails) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *CardBillingDetails) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### SetAddressCountryNil

`func (o *CardBillingDetails) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *CardBillingDetails) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetAddressPostalCode

`func (o *CardBillingDetails) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *CardBillingDetails) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *CardBillingDetails) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *CardBillingDetails) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### SetAddressPostalCodeNil

`func (o *CardBillingDetails) SetAddressPostalCodeNil(b bool)`

 SetAddressPostalCodeNil sets the value for AddressPostalCode to be an explicit nil

### UnsetAddressPostalCode
`func (o *CardBillingDetails) UnsetAddressPostalCode()`

UnsetAddressPostalCode ensures that no value is present for AddressPostalCode, not even an explicit nil
### GetAddressState

`func (o *CardBillingDetails) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *CardBillingDetails) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *CardBillingDetails) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *CardBillingDetails) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### SetAddressStateNil

`func (o *CardBillingDetails) SetAddressStateNil(b bool)`

 SetAddressStateNil sets the value for AddressState to be an explicit nil

### UnsetAddressState
`func (o *CardBillingDetails) UnsetAddressState()`

UnsetAddressState ensures that no value is present for AddressState, not even an explicit nil
### GetGivenNames

`func (o *CardBillingDetails) GetGivenNames() string`

GetGivenNames returns the GivenNames field if non-nil, zero value otherwise.

### GetGivenNamesOk

`func (o *CardBillingDetails) GetGivenNamesOk() (*string, bool)`

GetGivenNamesOk returns a tuple with the GivenNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNames

`func (o *CardBillingDetails) SetGivenNames(v string)`

SetGivenNames sets GivenNames field to given value.

### HasGivenNames

`func (o *CardBillingDetails) HasGivenNames() bool`

HasGivenNames returns a boolean if a field has been set.

### SetGivenNamesNil

`func (o *CardBillingDetails) SetGivenNamesNil(b bool)`

 SetGivenNamesNil sets the value for GivenNames to be an explicit nil

### UnsetGivenNames
`func (o *CardBillingDetails) UnsetGivenNames()`

UnsetGivenNames ensures that no value is present for GivenNames, not even an explicit nil
### GetMiddleName

`func (o *CardBillingDetails) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CardBillingDetails) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CardBillingDetails) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CardBillingDetails) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *CardBillingDetails) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *CardBillingDetails) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetSurname

`func (o *CardBillingDetails) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CardBillingDetails) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CardBillingDetails) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *CardBillingDetails) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *CardBillingDetails) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *CardBillingDetails) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetMobileNumber

`func (o *CardBillingDetails) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CardBillingDetails) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *CardBillingDetails) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *CardBillingDetails) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *CardBillingDetails) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *CardBillingDetails) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetPhoneNumber

`func (o *CardBillingDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CardBillingDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CardBillingDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CardBillingDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *CardBillingDetails) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *CardBillingDetails) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetAddress

`func (o *CardBillingDetails) GetAddress() CardBillingDetailsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CardBillingDetails) GetAddressOk() (*CardBillingDetailsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CardBillingDetails) SetAddress(v CardBillingDetailsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CardBillingDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CardBillingDetails) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CardBillingDetails) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


