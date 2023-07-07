# InternalV1PatchPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInternalV1PatchPaymentMethod

`func NewInternalV1PatchPaymentMethod() *InternalV1PatchPaymentMethod`

NewInternalV1PatchPaymentMethod instantiates a new InternalV1PatchPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalV1PatchPaymentMethodWithDefaults

`func NewInternalV1PatchPaymentMethodWithDefaults() *InternalV1PatchPaymentMethod`

NewInternalV1PatchPaymentMethodWithDefaults instantiates a new InternalV1PatchPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InternalV1PatchPaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalV1PatchPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalV1PatchPaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalV1PatchPaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalV1PatchPaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalV1PatchPaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalV1PatchPaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalV1PatchPaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InternalV1PatchPaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InternalV1PatchPaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetUpdated

`func (o *InternalV1PatchPaymentMethod) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InternalV1PatchPaymentMethod) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InternalV1PatchPaymentMethod) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *InternalV1PatchPaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


