# UpdatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PaymentRequestStatus**](PaymentRequestStatus.md) |  | [optional] 
**Actions** | Pointer to [**[]PaymentRequestAction**](PaymentRequestAction.md) |  | [optional] 
**InstrumentId** | Pointer to **string** |  | [optional] 
**Initiator** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePaymentRequest

`func NewUpdatePaymentRequest() *UpdatePaymentRequest`

NewUpdatePaymentRequest instantiates a new UpdatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentRequestWithDefaults

`func NewUpdatePaymentRequestWithDefaults() *UpdatePaymentRequest`

NewUpdatePaymentRequestWithDefaults instantiates a new UpdatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdatePaymentRequest) GetStatus() PaymentRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePaymentRequest) SetStatus(v PaymentRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdatePaymentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *UpdatePaymentRequest) GetActions() []PaymentRequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdatePaymentRequest) GetActionsOk() (*[]PaymentRequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdatePaymentRequest) SetActions(v []PaymentRequestAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdatePaymentRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetInstrumentId

`func (o *UpdatePaymentRequest) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *UpdatePaymentRequest) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *UpdatePaymentRequest) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *UpdatePaymentRequest) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetInitiator

`func (o *UpdatePaymentRequest) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *UpdatePaymentRequest) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *UpdatePaymentRequest) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *UpdatePaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


