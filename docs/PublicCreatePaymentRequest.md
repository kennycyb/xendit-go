# PublicCreatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Currency** | [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  | 
**PaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CaptureMethod** | Pointer to [**NullablePaymentRequestCaptureMethod**](PaymentRequestCaptureMethod.md) |  | [optional] 
**Initiator** | Pointer to [**NullablePaymentRequestInitiator**](PaymentRequestInitiator.md) |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**ChannelProperties** | Pointer to [**NullablePublicChannelPropertiesOverwrite**](PublicChannelPropertiesOverwrite.md) |  | [optional] 
**ShippingInformation** | Pointer to [**NullablePaymentRequestShippingInformation**](PaymentRequestShippingInformation.md) |  | [optional] 
**Items** | Pointer to [**[]PaymentRequestBasketItem**](PaymentRequestBasketItem.md) |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**InternalMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Invoice** | Pointer to [**NullableInvoice**](Invoice.md) |  | [optional] 

## Methods

### NewPublicCreatePaymentRequest

`func NewPublicCreatePaymentRequest(currency PaymentRequestCurrency, ) *PublicCreatePaymentRequest`

NewPublicCreatePaymentRequest instantiates a new PublicCreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCreatePaymentRequestWithDefaults

`func NewPublicCreatePaymentRequestWithDefaults() *PublicCreatePaymentRequest`

NewPublicCreatePaymentRequestWithDefaults instantiates a new PublicCreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *PublicCreatePaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PublicCreatePaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PublicCreatePaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PublicCreatePaymentRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAmount

`func (o *PublicCreatePaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicCreatePaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicCreatePaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PublicCreatePaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PublicCreatePaymentRequest) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PublicCreatePaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PublicCreatePaymentRequest) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *PublicCreatePaymentRequest) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PublicCreatePaymentRequest) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PublicCreatePaymentRequest) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PublicCreatePaymentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetDescription

`func (o *PublicCreatePaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicCreatePaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicCreatePaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicCreatePaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PublicCreatePaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PublicCreatePaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaptureMethod

`func (o *PublicCreatePaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *PublicCreatePaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *PublicCreatePaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *PublicCreatePaymentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### SetCaptureMethodNil

`func (o *PublicCreatePaymentRequest) SetCaptureMethodNil(b bool)`

 SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil

### UnsetCaptureMethod
`func (o *PublicCreatePaymentRequest) UnsetCaptureMethod()`

UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
### GetInitiator

`func (o *PublicCreatePaymentRequest) GetInitiator() PaymentRequestInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *PublicCreatePaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *PublicCreatePaymentRequest) SetInitiator(v PaymentRequestInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *PublicCreatePaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *PublicCreatePaymentRequest) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *PublicCreatePaymentRequest) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetPaymentMethodId

`func (o *PublicCreatePaymentRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PublicCreatePaymentRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PublicCreatePaymentRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PublicCreatePaymentRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetChannelProperties

`func (o *PublicCreatePaymentRequest) GetChannelProperties() PublicChannelPropertiesOverwrite`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PublicCreatePaymentRequest) GetChannelPropertiesOk() (*PublicChannelPropertiesOverwrite, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PublicCreatePaymentRequest) SetChannelProperties(v PublicChannelPropertiesOverwrite)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PublicCreatePaymentRequest) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *PublicCreatePaymentRequest) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *PublicCreatePaymentRequest) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetShippingInformation

`func (o *PublicCreatePaymentRequest) GetShippingInformation() PaymentRequestShippingInformation`

GetShippingInformation returns the ShippingInformation field if non-nil, zero value otherwise.

### GetShippingInformationOk

`func (o *PublicCreatePaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool)`

GetShippingInformationOk returns a tuple with the ShippingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInformation

`func (o *PublicCreatePaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation)`

SetShippingInformation sets ShippingInformation field to given value.

### HasShippingInformation

`func (o *PublicCreatePaymentRequest) HasShippingInformation() bool`

HasShippingInformation returns a boolean if a field has been set.

### SetShippingInformationNil

`func (o *PublicCreatePaymentRequest) SetShippingInformationNil(b bool)`

 SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil

### UnsetShippingInformation
`func (o *PublicCreatePaymentRequest) UnsetShippingInformation()`

UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
### GetItems

`func (o *PublicCreatePaymentRequest) GetItems() []PaymentRequestBasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PublicCreatePaymentRequest) GetItemsOk() (*[]PaymentRequestBasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PublicCreatePaymentRequest) SetItems(v []PaymentRequestBasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PublicCreatePaymentRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *PublicCreatePaymentRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *PublicCreatePaymentRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetCustomerId

`func (o *PublicCreatePaymentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PublicCreatePaymentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PublicCreatePaymentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PublicCreatePaymentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PublicCreatePaymentRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PublicCreatePaymentRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *PublicCreatePaymentRequest) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *PublicCreatePaymentRequest) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *PublicCreatePaymentRequest) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *PublicCreatePaymentRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *PublicCreatePaymentRequest) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *PublicCreatePaymentRequest) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetMetadata

`func (o *PublicCreatePaymentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicCreatePaymentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicCreatePaymentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicCreatePaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PublicCreatePaymentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PublicCreatePaymentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetInternalMetadata

`func (o *PublicCreatePaymentRequest) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *PublicCreatePaymentRequest) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *PublicCreatePaymentRequest) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *PublicCreatePaymentRequest) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *PublicCreatePaymentRequest) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *PublicCreatePaymentRequest) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil
### GetInvoice

`func (o *PublicCreatePaymentRequest) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *PublicCreatePaymentRequest) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *PublicCreatePaymentRequest) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *PublicCreatePaymentRequest) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### SetInvoiceNil

`func (o *PublicCreatePaymentRequest) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *PublicCreatePaymentRequest) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


