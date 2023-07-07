# PaymentRequestBasketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Category** | **string** |  | 
**SubCategory** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**Quantity** | **float32** |  | 
**Price** | **float32** |  | 
**PayerChargedCurrency** | Pointer to **string** |  | [optional] 
**PayerChargedPrice** | Pointer to **float32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaymentRequestBasketItem

`func NewPaymentRequestBasketItem(name string, category string, currency string, quantity float32, price float32, ) *PaymentRequestBasketItem`

NewPaymentRequestBasketItem instantiates a new PaymentRequestBasketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestBasketItemWithDefaults

`func NewPaymentRequestBasketItemWithDefaults() *PaymentRequestBasketItem`

NewPaymentRequestBasketItemWithDefaults instantiates a new PaymentRequestBasketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *PaymentRequestBasketItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentRequestBasketItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentRequestBasketItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentRequestBasketItem) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetName

`func (o *PaymentRequestBasketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentRequestBasketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentRequestBasketItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PaymentRequestBasketItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestBasketItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestBasketItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequestBasketItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PaymentRequestBasketItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentRequestBasketItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentRequestBasketItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentRequestBasketItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *PaymentRequestBasketItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PaymentRequestBasketItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PaymentRequestBasketItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSubCategory

`func (o *PaymentRequestBasketItem) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *PaymentRequestBasketItem) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *PaymentRequestBasketItem) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *PaymentRequestBasketItem) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentRequestBasketItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentRequestBasketItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentRequestBasketItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetQuantity

`func (o *PaymentRequestBasketItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentRequestBasketItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentRequestBasketItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *PaymentRequestBasketItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PaymentRequestBasketItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PaymentRequestBasketItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPayerChargedCurrency

`func (o *PaymentRequestBasketItem) GetPayerChargedCurrency() string`

GetPayerChargedCurrency returns the PayerChargedCurrency field if non-nil, zero value otherwise.

### GetPayerChargedCurrencyOk

`func (o *PaymentRequestBasketItem) GetPayerChargedCurrencyOk() (*string, bool)`

GetPayerChargedCurrencyOk returns a tuple with the PayerChargedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerChargedCurrency

`func (o *PaymentRequestBasketItem) SetPayerChargedCurrency(v string)`

SetPayerChargedCurrency sets PayerChargedCurrency field to given value.

### HasPayerChargedCurrency

`func (o *PaymentRequestBasketItem) HasPayerChargedCurrency() bool`

HasPayerChargedCurrency returns a boolean if a field has been set.

### GetPayerChargedPrice

`func (o *PaymentRequestBasketItem) GetPayerChargedPrice() float32`

GetPayerChargedPrice returns the PayerChargedPrice field if non-nil, zero value otherwise.

### GetPayerChargedPriceOk

`func (o *PaymentRequestBasketItem) GetPayerChargedPriceOk() (*float32, bool)`

GetPayerChargedPriceOk returns a tuple with the PayerChargedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerChargedPrice

`func (o *PaymentRequestBasketItem) SetPayerChargedPrice(v float32)`

SetPayerChargedPrice sets PayerChargedPrice field to given value.

### HasPayerChargedPrice

`func (o *PaymentRequestBasketItem) HasPayerChargedPrice() bool`

HasPayerChargedPrice returns a boolean if a field has been set.

### GetUrl

`func (o *PaymentRequestBasketItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentRequestBasketItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentRequestBasketItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaymentRequestBasketItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentRequestBasketItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentRequestBasketItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentRequestBasketItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentRequestBasketItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


