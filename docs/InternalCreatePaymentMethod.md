# InternalCreatePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PaymentMethodType**](PaymentMethodType.md) |  | 
**Country** | Pointer to [**NullableOptionalPaymentMethodCountry**](OptionalPaymentMethodCountry.md) |  | [optional] 
**Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Card** | Pointer to [**MutableCard**](MutableCard.md) |  | [optional] 
**Cryptocurrency** | Pointer to [**MutableCrypto**](MutableCrypto.md) |  | [optional] 
**DirectBankTransfer** | Pointer to [**MutableDirectBankTransfer**](MutableDirectBankTransfer.md) |  | [optional] 
**DirectDebit** | Pointer to [**MutableDirectDebit**](MutableDirectDebit.md) |  | [optional] 
**Ewallet** | Pointer to [**MutableEwallet**](MutableEwallet.md) |  | [optional] 
**OverTheCounter** | Pointer to [**MutableOverTheCounter**](MutableOverTheCounter.md) |  | [optional] 
**QrCode** | Pointer to [**MutableQRCode**](MutableQRCode.md) |  | [optional] 
**VirtualAccount** | Pointer to [**MutableVirtualAccount**](MutableVirtualAccount.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Items** | Pointer to [**[]BasketItem**](BasketItem.md) |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BillingInformation** | Pointer to [**NullableBillingInformation**](BillingInformation.md) |  | [optional] 

## Methods

### NewInternalCreatePaymentMethod

`func NewInternalCreatePaymentMethod(type_ PaymentMethodType, reusability PaymentMethodReusability, ) *InternalCreatePaymentMethod`

NewInternalCreatePaymentMethod instantiates a new InternalCreatePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalCreatePaymentMethodWithDefaults

`func NewInternalCreatePaymentMethodWithDefaults() *InternalCreatePaymentMethod`

NewInternalCreatePaymentMethodWithDefaults instantiates a new InternalCreatePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternalCreatePaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalCreatePaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalCreatePaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetCountry

`func (o *InternalCreatePaymentMethod) GetCountry() OptionalPaymentMethodCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalCreatePaymentMethod) GetCountryOk() (*OptionalPaymentMethodCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalCreatePaymentMethod) SetCountry(v OptionalPaymentMethodCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalCreatePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *InternalCreatePaymentMethod) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *InternalCreatePaymentMethod) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetReusability

`func (o *InternalCreatePaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *InternalCreatePaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *InternalCreatePaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetCustomerId

`func (o *InternalCreatePaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternalCreatePaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternalCreatePaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InternalCreatePaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InternalCreatePaymentMethod) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InternalCreatePaymentMethod) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *InternalCreatePaymentMethod) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalCreatePaymentMethod) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalCreatePaymentMethod) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalCreatePaymentMethod) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *InternalCreatePaymentMethod) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *InternalCreatePaymentMethod) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetReferenceId

`func (o *InternalCreatePaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalCreatePaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalCreatePaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalCreatePaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *InternalCreatePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalCreatePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalCreatePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalCreatePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternalCreatePaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternalCreatePaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCard

`func (o *InternalCreatePaymentMethod) GetCard() MutableCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *InternalCreatePaymentMethod) GetCardOk() (*MutableCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *InternalCreatePaymentMethod) SetCard(v MutableCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *InternalCreatePaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCryptocurrency

`func (o *InternalCreatePaymentMethod) GetCryptocurrency() MutableCrypto`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *InternalCreatePaymentMethod) GetCryptocurrencyOk() (*MutableCrypto, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *InternalCreatePaymentMethod) SetCryptocurrency(v MutableCrypto)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *InternalCreatePaymentMethod) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### GetDirectBankTransfer

`func (o *InternalCreatePaymentMethod) GetDirectBankTransfer() MutableDirectBankTransfer`

GetDirectBankTransfer returns the DirectBankTransfer field if non-nil, zero value otherwise.

### GetDirectBankTransferOk

`func (o *InternalCreatePaymentMethod) GetDirectBankTransferOk() (*MutableDirectBankTransfer, bool)`

GetDirectBankTransferOk returns a tuple with the DirectBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectBankTransfer

`func (o *InternalCreatePaymentMethod) SetDirectBankTransfer(v MutableDirectBankTransfer)`

SetDirectBankTransfer sets DirectBankTransfer field to given value.

### HasDirectBankTransfer

`func (o *InternalCreatePaymentMethod) HasDirectBankTransfer() bool`

HasDirectBankTransfer returns a boolean if a field has been set.

### GetDirectDebit

`func (o *InternalCreatePaymentMethod) GetDirectDebit() MutableDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *InternalCreatePaymentMethod) GetDirectDebitOk() (*MutableDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *InternalCreatePaymentMethod) SetDirectDebit(v MutableDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *InternalCreatePaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### GetEwallet

`func (o *InternalCreatePaymentMethod) GetEwallet() MutableEwallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *InternalCreatePaymentMethod) GetEwalletOk() (*MutableEwallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *InternalCreatePaymentMethod) SetEwallet(v MutableEwallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *InternalCreatePaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *InternalCreatePaymentMethod) GetOverTheCounter() MutableOverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *InternalCreatePaymentMethod) GetOverTheCounterOk() (*MutableOverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *InternalCreatePaymentMethod) SetOverTheCounter(v MutableOverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *InternalCreatePaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### GetQrCode

`func (o *InternalCreatePaymentMethod) GetQrCode() MutableQRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *InternalCreatePaymentMethod) GetQrCodeOk() (*MutableQRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *InternalCreatePaymentMethod) SetQrCode(v MutableQRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *InternalCreatePaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *InternalCreatePaymentMethod) GetVirtualAccount() MutableVirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *InternalCreatePaymentMethod) GetVirtualAccountOk() (*MutableVirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *InternalCreatePaymentMethod) SetVirtualAccount(v MutableVirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *InternalCreatePaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalCreatePaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalCreatePaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalCreatePaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalCreatePaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalCreatePaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalCreatePaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetItems

`func (o *InternalCreatePaymentMethod) GetItems() []BasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalCreatePaymentMethod) GetItemsOk() (*[]BasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalCreatePaymentMethod) SetItems(v []BasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *InternalCreatePaymentMethod) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *InternalCreatePaymentMethod) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *InternalCreatePaymentMethod) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetInternalMetadata

`func (o *InternalCreatePaymentMethod) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalCreatePaymentMethod) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalCreatePaymentMethod) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalCreatePaymentMethod) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalCreatePaymentMethod) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalCreatePaymentMethod) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil
### GetBillingInformation

`func (o *InternalCreatePaymentMethod) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *InternalCreatePaymentMethod) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *InternalCreatePaymentMethod) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *InternalCreatePaymentMethod) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *InternalCreatePaymentMethod) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *InternalCreatePaymentMethod) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


