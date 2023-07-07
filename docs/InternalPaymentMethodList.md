# InternalPaymentMethodList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InternalPaymentMethod**](InternalPaymentMethod.md) |  | 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewInternalPaymentMethodList

`func NewInternalPaymentMethodList(data []InternalPaymentMethod, ) *InternalPaymentMethodList`

NewInternalPaymentMethodList instantiates a new InternalPaymentMethodList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalPaymentMethodListWithDefaults

`func NewInternalPaymentMethodListWithDefaults() *InternalPaymentMethodList`

NewInternalPaymentMethodListWithDefaults instantiates a new InternalPaymentMethodList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InternalPaymentMethodList) GetData() []InternalPaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalPaymentMethodList) GetDataOk() (*[]InternalPaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalPaymentMethodList) SetData(v []InternalPaymentMethod)`

SetData sets Data field to given value.


### GetHasMore

`func (o *InternalPaymentMethodList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InternalPaymentMethodList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InternalPaymentMethodList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InternalPaymentMethodList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


