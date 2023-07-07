# PublicCreatePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PaymentMethodType**](PaymentMethodType.md) |  | 
**Country** | Pointer to [**NullableOptionalPaymentMethodCountry**](OptionalPaymentMethodCountry.md) |  | [optional] 
**Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cryptocurrency** | Pointer to [**MutableCrypto**](MutableCrypto.md) |  | [optional] 
**DirectBankTransfer** | Pointer to [**MutableDirectBankTransfer**](MutableDirectBankTransfer.md) |  | [optional] 
**DirectDebit** | Pointer to [**MutableDirectDebit**](MutableDirectDebit.md) |  | [optional] 
**Ewallet** | Pointer to [**MutableEwallet**](MutableEwallet.md) |  | [optional] 
**OverTheCounter** | Pointer to [**MutableOverTheCounter**](MutableOverTheCounter.md) |  | [optional] 
**VirtualAccount** | Pointer to [**MutableVirtualAccount**](MutableVirtualAccount.md) |  | [optional] 
**QrCode** | Pointer to [**MutableQRCode**](MutableQRCode.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BillingInformation** | Pointer to [**NullableBillingInformation**](BillingInformation.md) |  | [optional] 

## Methods

### NewPublicCreatePaymentMethod

`func NewPublicCreatePaymentMethod(type_ PaymentMethodType, reusability PaymentMethodReusability, ) *PublicCreatePaymentMethod`

NewPublicCreatePaymentMethod instantiates a new PublicCreatePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCreatePaymentMethodWithDefaults

`func NewPublicCreatePaymentMethodWithDefaults() *PublicCreatePaymentMethod`

NewPublicCreatePaymentMethodWithDefaults instantiates a new PublicCreatePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicCreatePaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicCreatePaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicCreatePaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetCountry

`func (o *PublicCreatePaymentMethod) GetCountry() OptionalPaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicCreatePaymentMethod) GetCountryOk() (*OptionalPaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicCreatePaymentMethod) SetCountry(v OptionalPaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicCreatePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PublicCreatePaymentMethod) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PublicCreatePaymentMethod) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetReusability

`func (o *PublicCreatePaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PublicCreatePaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PublicCreatePaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetCustomerId

`func (o *PublicCreatePaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PublicCreatePaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PublicCreatePaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PublicCreatePaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PublicCreatePaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PublicCreatePaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetReferenceId

`func (o *PublicCreatePaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PublicCreatePaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PublicCreatePaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PublicCreatePaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *PublicCreatePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicCreatePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicCreatePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicCreatePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PublicCreatePaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PublicCreatePaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCryptocurrency

`func (o *PublicCreatePaymentMethod) GetCryptocurrency() MutableCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *PublicCreatePaymentMethod) GetCryptocurrencyOk() (*MutableCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *PublicCreatePaymentMethod) SetCryptocurrency(v MutableCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *PublicCreatePaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### GetDirectBankTransfer

`func (o *PublicCreatePaymentMethod) GetDirectBankTransfer() MutableDirectBankTransfer`

GetDirectBankTransfer returns the DirectBankTransfer field if non-nil, zero value otherwise.

### GetDirectBankTransferOk

`func (o *PublicCreatePaymentMethod) GetDirectBankTransferOk() (*MutableDirectBankTransfer, bool)`

GetDirectBankTransferOk returns a tuple with the DirectBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectBankTransfer

`func (o *PublicCreatePaymentMethod) SetDirectBankTransfer(v MutableDirectBankTransfer)`

SetDirectBankTransfer sets DirectBankTransfer field to given value.

### HasDirectBankTransfer

`func (o *PublicCreatePaymentMethod) HasDirectBankTransfer() bool`

HasDirectBankTransfer returns a boolean if a field has been set.

### GetDirectDebit

`func (o *PublicCreatePaymentMethod) GetDirectDebit() MutableDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PublicCreatePaymentMethod) GetDirectDebitOk() (*MutableDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PublicCreatePaymentMethod) SetDirectDebit(v MutableDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PublicCreatePaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### GetEwallet

`func (o *PublicCreatePaymentMethod) GetEwallet() MutableEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PublicCreatePaymentMethod) GetEwalletOk() (*MutableEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PublicCreatePaymentMethod) SetEwallet(v MutableEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PublicCreatePaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *PublicCreatePaymentMethod) GetOverTheCounter() MutableOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PublicCreatePaymentMethod) GetOverTheCounterOk() (*MutableOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PublicCreatePaymentMethod) SetOverTheCounter(v MutableOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PublicCreatePaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *PublicCreatePaymentMethod) GetVirtualAccount() MutableVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PublicCreatePaymentMethod) GetVirtualAccountOk() (*MutableVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PublicCreatePaymentMethod) SetVirtualAccount(v MutableVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PublicCreatePaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetQrCode

`func (o *PublicCreatePaymentMethod) GetQrCode() MutableQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PublicCreatePaymentMethod) GetQrCodeOk() (*MutableQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PublicCreatePaymentMethod) SetQrCode(v MutableQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PublicCreatePaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetMetadata

`func (o *PublicCreatePaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicCreatePaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicCreatePaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicCreatePaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PublicCreatePaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PublicCreatePaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBillingInformation

`func (o *PublicCreatePaymentMethod) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *PublicCreatePaymentMethod) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *PublicCreatePaymentMethod) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *PublicCreatePaymentMethod) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *PublicCreatePaymentMethod) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *PublicCreatePaymentMethod) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


