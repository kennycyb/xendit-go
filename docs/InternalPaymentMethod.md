# InternalPaymentMethod

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
**Ewallet** | Pointer to [**NullableInternalEwallet**](InternalEwallet.md) |  | [optional] 
**Card** | Pointer to [**NullableInternalCard**](InternalCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**NullableInternalCrypto**](InternalCrypto.md) |  | [optional] 
**DirectBankTransfer** | Pointer to [**NullableInternalDirectBankTransfer**](InternalDirectBankTransfer.md) |  | [optional] 
**DirectDebit** | Pointer to [**NullableInternalDirectDebit**](InternalDirectDebit.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullableInternalOverTheCounter**](InternalOverTheCounter.md) |  | [optional] 
**VirtualAccount** | Pointer to [**NullableInternalVirtualAccount**](InternalVirtualAccount.md) |  | [optional] 
**QrCode** | Pointer to [**NullableInternalQRCode**](InternalQRCode.md) |  | [optional] 
**XCallbackUrl** | Pointer to **NullableString** |  | [optional] 
**ClientType** | Pointer to **NullableString** |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInternalPaymentMethod

`func NewInternalPaymentMethod(id string, ) *InternalPaymentMethod`

NewInternalPaymentMethod instantiates a new InternalPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalPaymentMethodWithDefaults

`func NewInternalPaymentMethodWithDefaults() *InternalPaymentMethod`

NewInternalPaymentMethodWithDefaults instantiates a new InternalPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalPaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalPaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalPaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetBusinessId

`func (o *InternalPaymentMethod) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *InternalPaymentMethod) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *InternalPaymentMethod) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *InternalPaymentMethod) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetType

`func (o *InternalPaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalPaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalPaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *InternalPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *InternalPaymentMethod) GetCountry() PaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalPaymentMethod) GetCountryOk() (*PaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalPaymentMethod) SetCountry(v PaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalPaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomerId

`func (o *InternalPaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternalPaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternalPaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InternalPaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InternalPaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InternalPaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *InternalPaymentMethod) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalPaymentMethod) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalPaymentMethod) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalPaymentMethod) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *InternalPaymentMethod) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *InternalPaymentMethod) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetReferenceId

`func (o *InternalPaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalPaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalPaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalPaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *InternalPaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalPaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalPaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalPaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternalPaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternalPaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *InternalPaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalPaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalPaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReusability

`func (o *InternalPaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *InternalPaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *InternalPaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *InternalPaymentMethod) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetActions

`func (o *InternalPaymentMethod) GetActions() []PaymentMethodAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InternalPaymentMethod) GetActionsOk() (*[]PaymentMethodAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InternalPaymentMethod) SetActions(v []PaymentMethodAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *InternalPaymentMethod) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalPaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalPaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalPaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalPaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalPaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalPaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBillingInformation

`func (o *InternalPaymentMethod) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *InternalPaymentMethod) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *InternalPaymentMethod) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *InternalPaymentMethod) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *InternalPaymentMethod) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *InternalPaymentMethod) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
### GetFailureCode

`func (o *InternalPaymentMethod) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *InternalPaymentMethod) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *InternalPaymentMethod) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *InternalPaymentMethod) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *InternalPaymentMethod) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *InternalPaymentMethod) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCreated

`func (o *InternalPaymentMethod) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InternalPaymentMethod) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InternalPaymentMethod) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InternalPaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *InternalPaymentMethod) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InternalPaymentMethod) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InternalPaymentMethod) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *InternalPaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetEwallet

`func (o *InternalPaymentMethod) GetEwallet() InternalEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *InternalPaymentMethod) GetEwalletOk() (*InternalEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *InternalPaymentMethod) SetEwallet(v InternalEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *InternalPaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *InternalPaymentMethod) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *InternalPaymentMethod) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetCard

`func (o *InternalPaymentMethod) GetCard() InternalCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *InternalPaymentMethod) GetCardOk() (*InternalCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *InternalPaymentMethod) SetCard(v InternalCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *InternalPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *InternalPaymentMethod) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *InternalPaymentMethod) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCryptocurrency

`func (o *InternalPaymentMethod) GetCryptocurrency() InternalCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *InternalPaymentMethod) GetCryptocurrencyOk() (*InternalCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *InternalPaymentMethod) SetCryptocurrency(v InternalCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *InternalPaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### SetCryptocurrencyNil

`func (o *InternalPaymentMethod) SetCryptocurrencyNil(b bool)`

 SetCryptocurrencyNil sets the value for Cryptocurrency to be an explicit nil

### UnsetCryptocurrency
`func (o *InternalPaymentMethod) UnsetCryptocurrency()`

UnsetCryptocurrency ensures that no value is present for Cryptocurrency, not even an explicit nil
### GetDirectBankTransfer

`func (o *InternalPaymentMethod) GetDirectBankTransfer() InternalDirectBankTransfer`

GetDirectBankTransfer returns the DirectBankTransfer field if non-nil, zero value otherwise.

### GetDirectBankTransferOk

`func (o *InternalPaymentMethod) GetDirectBankTransferOk() (*InternalDirectBankTransfer, bool)`

GetDirectBankTransferOk returns a tuple with the DirectBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectBankTransfer

`func (o *InternalPaymentMethod) SetDirectBankTransfer(v InternalDirectBankTransfer)`

SetDirectBankTransfer sets DirectBankTransfer field to given value.

### HasDirectBankTransfer

`func (o *InternalPaymentMethod) HasDirectBankTransfer() bool`

HasDirectBankTransfer returns a boolean if a field has been set.

### SetDirectBankTransferNil

`func (o *InternalPaymentMethod) SetDirectBankTransferNil(b bool)`

 SetDirectBankTransferNil sets the value for DirectBankTransfer to be an explicit nil

### UnsetDirectBankTransfer
`func (o *InternalPaymentMethod) UnsetDirectBankTransfer()`

UnsetDirectBankTransfer ensures that no value is present for DirectBankTransfer, not even an explicit nil
### GetDirectDebit

`func (o *InternalPaymentMethod) GetDirectDebit() InternalDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *InternalPaymentMethod) GetDirectDebitOk() (*InternalDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *InternalPaymentMethod) SetDirectDebit(v InternalDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *InternalPaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *InternalPaymentMethod) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *InternalPaymentMethod) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetOverTheCounter

`func (o *InternalPaymentMethod) GetOverTheCounter() InternalOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *InternalPaymentMethod) GetOverTheCounterOk() (*InternalOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *InternalPaymentMethod) SetOverTheCounter(v InternalOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *InternalPaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *InternalPaymentMethod) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *InternalPaymentMethod) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetVirtualAccount

`func (o *InternalPaymentMethod) GetVirtualAccount() InternalVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *InternalPaymentMethod) GetVirtualAccountOk() (*InternalVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *InternalPaymentMethod) SetVirtualAccount(v InternalVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *InternalPaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### SetVirtualAccountNil

`func (o *InternalPaymentMethod) SetVirtualAccountNil(b bool)`

 SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil

### UnsetVirtualAccount
`func (o *InternalPaymentMethod) UnsetVirtualAccount()`

UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
### GetQrCode

`func (o *InternalPaymentMethod) GetQrCode() InternalQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *InternalPaymentMethod) GetQrCodeOk() (*InternalQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *InternalPaymentMethod) SetQrCode(v InternalQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *InternalPaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *InternalPaymentMethod) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *InternalPaymentMethod) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
### GetXCallbackUrl

`func (o *InternalPaymentMethod) GetXCallbackUrl() string`

GetXCallbackUrl returns the XCallbackUrl field if non-nil, zero value otherwise.

### GetXCallbackUrlOk

`func (o *InternalPaymentMethod) GetXCallbackUrlOk() (*string, bool)`

GetXCallbackUrlOk returns a tuple with the XCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCallbackUrl

`func (o *InternalPaymentMethod) SetXCallbackUrl(v string)`

SetXCallbackUrl sets XCallbackUrl field to given value.

### HasXCallbackUrl

`func (o *InternalPaymentMethod) HasXCallbackUrl() bool`

HasXCallbackUrl returns a boolean if a field has been set.

### SetXCallbackUrlNil

`func (o *InternalPaymentMethod) SetXCallbackUrlNil(b bool)`

 SetXCallbackUrlNil sets the value for XCallbackUrl to be an explicit nil

### UnsetXCallbackUrl
`func (o *InternalPaymentMethod) UnsetXCallbackUrl()`

UnsetXCallbackUrl ensures that no value is present for XCallbackUrl, not even an explicit nil
### GetClientType

`func (o *InternalPaymentMethod) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *InternalPaymentMethod) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *InternalPaymentMethod) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *InternalPaymentMethod) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### SetClientTypeNil

`func (o *InternalPaymentMethod) SetClientTypeNil(b bool)`

 SetClientTypeNil sets the value for ClientType to be an explicit nil

### UnsetClientType
`func (o *InternalPaymentMethod) UnsetClientType()`

UnsetClientType ensures that no value is present for ClientType, not even an explicit nil
### GetInternalMetadata

`func (o *InternalPaymentMethod) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalPaymentMethod) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalPaymentMethod) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalPaymentMethod) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalPaymentMethod) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalPaymentMethod) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


