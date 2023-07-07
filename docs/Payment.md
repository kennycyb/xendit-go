# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**BusinessId** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PaymentRequestId** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ChannelCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**PaymentRequestCountry**](PaymentRequestCountry.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Currency** | Pointer to [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  | [optional] 
**PaymentDetail** | Pointer to **map[string]interface{}** |  | [optional] 
**PaymentMethod** | Pointer to [**PublicPaymentMethod**](PublicPaymentMethod.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBusinessId

`func (o *Payment) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Payment) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Payment) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Payment) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *Payment) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Payment) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Payment) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Payment) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetType

`func (o *Payment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Payment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPaymentRequestId

`func (o *Payment) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *Payment) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *Payment) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.

### HasPaymentRequestId

`func (o *Payment) HasPaymentRequestId() bool`

HasPaymentRequestId returns a boolean if a field has been set.

### GetReferenceId

`func (o *Payment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Payment) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Payment) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Payment) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetChannelCode

`func (o *Payment) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Payment) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Payment) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *Payment) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetCountry

`func (o *Payment) GetCountry() PaymentRequestCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Payment) GetCountryOk() (*PaymentRequestCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Payment) SetCountry(v PaymentRequestCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Payment) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmount

`func (o *Payment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Payment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Payment) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Payment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentDetail

`func (o *Payment) GetPaymentDetail() map[string]interface{}`

GetPaymentDetail returns the PaymentDetail field if non-nil, zero value otherwise.

### GetPaymentDetailOk

`func (o *Payment) GetPaymentDetailOk() (*map[string]interface{}, bool)`

GetPaymentDetailOk returns a tuple with the PaymentDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetail

`func (o *Payment) SetPaymentDetail(v map[string]interface{})`

SetPaymentDetail sets PaymentDetail field to given value.

### HasPaymentDetail

`func (o *Payment) HasPaymentDetail() bool`

HasPaymentDetail returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Payment) GetPaymentMethod() PublicPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Payment) GetPaymentMethodOk() (*PublicPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Payment) SetPaymentMethod(v PublicPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Payment) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetMetadata

`func (o *Payment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Payment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Payment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Payment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Payment) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Payment) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreated

`func (o *Payment) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Payment) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Payment) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Payment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Payment) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Payment) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Payment) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Payment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


