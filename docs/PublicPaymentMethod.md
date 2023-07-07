# PublicPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BusinessId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PaymentMethodType**](PaymentMethodType.md) |  | [optional] 
**Country** | Pointer to [**PaymentMethodCountry**](PaymentMethodCountry.md) |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**Reusability** | Pointer to [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | [optional] 
**Actions** | Pointer to [**[]PaymentMethodAction**](PaymentMethodAction.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BillingInformation** | Pointer to [**NullableBillingInformation**](BillingInformation.md) |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Ewallet** | Pointer to [**NullablePublicEwallet**](PublicEwallet.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullablePublicDirectDebit**](PublicDirectDebit.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullablePublicOverTheCounter**](PublicOverTheCounter.md) |  | [optional] 
**Card** | Pointer to [**NullablePublicCard**](PublicCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**NullablePublicCrypto**](PublicCrypto.md) |  | [optional] 
**QrCode** | Pointer to [**NullablePublicQRCode**](PublicQRCode.md) |  | [optional] 

## Methods

### NewPublicPaymentMethod

`func NewPublicPaymentMethod(id string, ) *PublicPaymentMethod`

NewPublicPaymentMethod instantiates a new PublicPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentMethodWithDefaults

`func NewPublicPaymentMethodWithDefaults() *PublicPaymentMethod`

NewPublicPaymentMethodWithDefaults instantiates a new PublicPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicPaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicPaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicPaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetBusinessId

`func (o *PublicPaymentMethod) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PublicPaymentMethod) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PublicPaymentMethod) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *PublicPaymentMethod) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetType

`func (o *PublicPaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicPaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicPaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *PublicPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *PublicPaymentMethod) GetCountry() PaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicPaymentMethod) GetCountryOk() (*PaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicPaymentMethod) SetCountry(v PaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicPaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomerId

`func (o *PublicPaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PublicPaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PublicPaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PublicPaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PublicPaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PublicPaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *PublicPaymentMethod) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *PublicPaymentMethod) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *PublicPaymentMethod) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *PublicPaymentMethod) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *PublicPaymentMethod) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *PublicPaymentMethod) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetReferenceId

`func (o *PublicPaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PublicPaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PublicPaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PublicPaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *PublicPaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicPaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicPaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicPaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PublicPaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PublicPaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *PublicPaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicPaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicPaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReusability

`func (o *PublicPaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PublicPaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PublicPaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *PublicPaymentMethod) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetActions

`func (o *PublicPaymentMethod) GetActions() []PaymentMethodAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PublicPaymentMethod) GetActionsOk() (*[]PaymentMethodAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PublicPaymentMethod) SetActions(v []PaymentMethodAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PublicPaymentMethod) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *PublicPaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicPaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicPaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicPaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PublicPaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PublicPaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBillingInformation

`func (o *PublicPaymentMethod) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *PublicPaymentMethod) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *PublicPaymentMethod) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *PublicPaymentMethod) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *PublicPaymentMethod) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *PublicPaymentMethod) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
### GetFailureCode

`func (o *PublicPaymentMethod) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PublicPaymentMethod) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PublicPaymentMethod) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PublicPaymentMethod) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PublicPaymentMethod) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PublicPaymentMethod) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCreated

`func (o *PublicPaymentMethod) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicPaymentMethod) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicPaymentMethod) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicPaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PublicPaymentMethod) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PublicPaymentMethod) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PublicPaymentMethod) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PublicPaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetEwallet

`func (o *PublicPaymentMethod) GetEwallet() PublicEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PublicPaymentMethod) GetEwalletOk() (*PublicEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PublicPaymentMethod) SetEwallet(v PublicEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PublicPaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *PublicPaymentMethod) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *PublicPaymentMethod) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetDirectDebit

`func (o *PublicPaymentMethod) GetDirectDebit() PublicDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PublicPaymentMethod) GetDirectDebitOk() (*PublicDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PublicPaymentMethod) SetDirectDebit(v PublicDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PublicPaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *PublicPaymentMethod) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *PublicPaymentMethod) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetOverTheCounter

`func (o *PublicPaymentMethod) GetOverTheCounter() PublicOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PublicPaymentMethod) GetOverTheCounterOk() (*PublicOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PublicPaymentMethod) SetOverTheCounter(v PublicOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PublicPaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *PublicPaymentMethod) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *PublicPaymentMethod) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetCard

`func (o *PublicPaymentMethod) GetCard() PublicCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PublicPaymentMethod) GetCardOk() (*PublicCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PublicPaymentMethod) SetCard(v PublicCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *PublicPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *PublicPaymentMethod) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *PublicPaymentMethod) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCryptocurrency

`func (o *PublicPaymentMethod) GetCryptocurrency() PublicCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *PublicPaymentMethod) GetCryptocurrencyOk() (*PublicCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *PublicPaymentMethod) SetCryptocurrency(v PublicCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *PublicPaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### SetCryptocurrencyNil

`func (o *PublicPaymentMethod) SetCryptocurrencyNil(b bool)`

 SetCryptocurrencyNil sets the value for Cryptocurrency to be an explicit nil

### UnsetCryptocurrency
`func (o *PublicPaymentMethod) UnsetCryptocurrency()`

UnsetCryptocurrency ensures that no value is present for Cryptocurrency, not even an explicit nil
### GetQrCode

`func (o *PublicPaymentMethod) GetQrCode() PublicQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PublicPaymentMethod) GetQrCodeOk() (*PublicQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PublicPaymentMethod) SetQrCode(v PublicQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PublicPaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *PublicPaymentMethod) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PublicPaymentMethod) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


