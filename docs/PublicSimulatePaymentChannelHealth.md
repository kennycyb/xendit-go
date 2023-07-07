# PublicSimulatePaymentChannelHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Payment channel health event | [optional] 
**BusinessId** | Pointer to **string** | Business ID registered to payment channel downtime or online callback URL | [optional] 
**Created** | Pointer to **string** | Date time the event is triggered | [optional] 
**Data** | Pointer to [**PublicPaymentChannelDowntimeOnlineData**](PublicPaymentChannelDowntimeOnlineData.md) |  | [optional] 

## Methods

### NewPublicSimulatePaymentChannelHealth

`func NewPublicSimulatePaymentChannelHealth() *PublicSimulatePaymentChannelHealth`

NewPublicSimulatePaymentChannelHealth instantiates a new PublicSimulatePaymentChannelHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSimulatePaymentChannelHealthWithDefaults

`func NewPublicSimulatePaymentChannelHealthWithDefaults() *PublicSimulatePaymentChannelHealth`

NewPublicSimulatePaymentChannelHealthWithDefaults instantiates a new PublicSimulatePaymentChannelHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PublicSimulatePaymentChannelHealth) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PublicSimulatePaymentChannelHealth) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PublicSimulatePaymentChannelHealth) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PublicSimulatePaymentChannelHealth) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetBusinessId

`func (o *PublicSimulatePaymentChannelHealth) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PublicSimulatePaymentChannelHealth) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PublicSimulatePaymentChannelHealth) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *PublicSimulatePaymentChannelHealth) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreated

`func (o *PublicSimulatePaymentChannelHealth) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicSimulatePaymentChannelHealth) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicSimulatePaymentChannelHealth) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicSimulatePaymentChannelHealth) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *PublicSimulatePaymentChannelHealth) GetData() PublicPaymentChannelDowntimeOnlineData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicSimulatePaymentChannelHealth) GetDataOk() (*PublicPaymentChannelDowntimeOnlineData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicSimulatePaymentChannelHealth) SetData(v PublicPaymentChannelDowntimeOnlineData)`

SetData sets Data field to given value.

### HasData

`func (o *PublicSimulatePaymentChannelHealth) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


