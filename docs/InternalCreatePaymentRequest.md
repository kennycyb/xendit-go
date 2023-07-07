# InternalCreatePaymentRequest

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

### NewInternalCreatePaymentRequest

`func NewInternalCreatePaymentRequest(currency PaymentRequestCurrency, ) *InternalCreatePaymentRequest`

NewInternalCreatePaymentRequest instantiates a new InternalCreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalCreatePaymentRequestWithDefaults

`func NewInternalCreatePaymentRequestWithDefaults() *InternalCreatePaymentRequest`

NewInternalCreatePaymentRequestWithDefaults instantiates a new InternalCreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *InternalCreatePaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalCreatePaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalCreatePaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalCreatePaymentRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAmount

`func (o *InternalCreatePaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalCreatePaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalCreatePaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalCreatePaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalCreatePaymentRequest) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalCreatePaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalCreatePaymentRequest) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *InternalCreatePaymentRequest) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InternalCreatePaymentRequest) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InternalCreatePaymentRequest) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *InternalCreatePaymentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetDescription

`func (o *InternalCreatePaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalCreatePaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalCreatePaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalCreatePaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternalCreatePaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternalCreatePaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaptureMethod

`func (o *InternalCreatePaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *InternalCreatePaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *InternalCreatePaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *InternalCreatePaymentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### SetCaptureMethodNil

`func (o *InternalCreatePaymentRequest) SetCaptureMethodNil(b bool)`

 SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil

### UnsetCaptureMethod
`func (o *InternalCreatePaymentRequest) UnsetCaptureMethod()`

UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
### GetInitiator

`func (o *InternalCreatePaymentRequest) GetInitiator() PaymentRequestInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InternalCreatePaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InternalCreatePaymentRequest) SetInitiator(v PaymentRequestInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InternalCreatePaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *InternalCreatePaymentRequest) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *InternalCreatePaymentRequest) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetPaymentMethodId

`func (o *InternalCreatePaymentRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *InternalCreatePaymentRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *InternalCreatePaymentRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *InternalCreatePaymentRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetChannelProperties

`func (o *InternalCreatePaymentRequest) GetChannelProperties() PublicChannelPropertiesOverwrite`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalCreatePaymentRequest) GetChannelPropertiesOk() (*PublicChannelPropertiesOverwrite, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalCreatePaymentRequest) SetChannelProperties(v PublicChannelPropertiesOverwrite)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *InternalCreatePaymentRequest) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *InternalCreatePaymentRequest) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *InternalCreatePaymentRequest) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetShippingInformation

`func (o *InternalCreatePaymentRequest) GetShippingInformation() PaymentRequestShippingInformation`

GetShippingInformation returns the ShippingInformation field if non-nil, zero value otherwise.

### GetShippingInformationOk

`func (o *InternalCreatePaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool)`

GetShippingInformationOk returns a tuple with the ShippingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInformation

`func (o *InternalCreatePaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation)`

SetShippingInformation sets ShippingInformation field to given value.

### HasShippingInformation

`func (o *InternalCreatePaymentRequest) HasShippingInformation() bool`

HasShippingInformation returns a boolean if a field has been set.

### SetShippingInformationNil

`func (o *InternalCreatePaymentRequest) SetShippingInformationNil(b bool)`

 SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil

### UnsetShippingInformation
`func (o *InternalCreatePaymentRequest) UnsetShippingInformation()`

UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
### GetItems

`func (o *InternalCreatePaymentRequest) GetItems() []PaymentRequestBasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalCreatePaymentRequest) GetItemsOk() (*[]PaymentRequestBasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalCreatePaymentRequest) SetItems(v []PaymentRequestBasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *InternalCreatePaymentRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *InternalCreatePaymentRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *InternalCreatePaymentRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetCustomerId

`func (o *InternalCreatePaymentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternalCreatePaymentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternalCreatePaymentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InternalCreatePaymentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InternalCreatePaymentRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InternalCreatePaymentRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *InternalCreatePaymentRequest) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalCreatePaymentRequest) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalCreatePaymentRequest) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalCreatePaymentRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *InternalCreatePaymentRequest) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *InternalCreatePaymentRequest) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetMetadata

`func (o *InternalCreatePaymentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalCreatePaymentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalCreatePaymentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalCreatePaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalCreatePaymentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalCreatePaymentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetInternalMetadata

`func (o *InternalCreatePaymentRequest) GetInternalMetadata() map[string]interface{}`

GetInternalMetadata returns the InternalMetadata field if non-nil, zero value otherwise.

### GetInternalMetadataOk

`func (o *InternalCreatePaymentRequest) GetInternalMetadataOk() (*map[string]interface{}, bool)`

GetInternalMetadataOk returns a tuple with the InternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadata

`func (o *InternalCreatePaymentRequest) SetInternalMetadata(v map[string]interface{})`

SetInternalMetadata sets InternalMetadata field to given value.

### HasInternalMetadata

`func (o *InternalCreatePaymentRequest) HasInternalMetadata() bool`

HasInternalMetadata returns a boolean if a field has been set.

### SetInternalMetadataNil

`func (o *InternalCreatePaymentRequest) SetInternalMetadataNil(b bool)`

 SetInternalMetadataNil sets the value for InternalMetadata to be an explicit nil

### UnsetInternalMetadata
`func (o *InternalCreatePaymentRequest) UnsetInternalMetadata()`

UnsetInternalMetadata ensures that no value is present for InternalMetadata, not even an explicit nil
### GetInvoice

`func (o *InternalCreatePaymentRequest) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InternalCreatePaymentRequest) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InternalCreatePaymentRequest) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InternalCreatePaymentRequest) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### SetInvoiceNil

`func (o *InternalCreatePaymentRequest) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *InternalCreatePaymentRequest) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


