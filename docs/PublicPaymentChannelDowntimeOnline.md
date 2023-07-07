# PublicPaymentChannelDowntimeOnline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Payment channel health event | [optional] 
**BusinessId** | Pointer to **string** | Business ID registered to payment channel downtime or online callback URL | [optional] 
**Created** | Pointer to **string** | Date time the event is triggered | [optional] 
**Data** | Pointer to [**PublicPaymentChannelDowntimeOnlineData**](PublicPaymentChannelDowntimeOnlineData.md) |  | [optional] 

## Methods

### NewPublicPaymentChannelDowntimeOnline

`func NewPublicPaymentChannelDowntimeOnline() *PublicPaymentChannelDowntimeOnline`

NewPublicPaymentChannelDowntimeOnline instantiates a new PublicPaymentChannelDowntimeOnline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentChannelDowntimeOnlineWithDefaults

`func NewPublicPaymentChannelDowntimeOnlineWithDefaults() *PublicPaymentChannelDowntimeOnline`

NewPublicPaymentChannelDowntimeOnlineWithDefaults instantiates a new PublicPaymentChannelDowntimeOnline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PublicPaymentChannelDowntimeOnline) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PublicPaymentChannelDowntimeOnline) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PublicPaymentChannelDowntimeOnline) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PublicPaymentChannelDowntimeOnline) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetBusinessId

`func (o *PublicPaymentChannelDowntimeOnline) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PublicPaymentChannelDowntimeOnline) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PublicPaymentChannelDowntimeOnline) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *PublicPaymentChannelDowntimeOnline) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreated

`func (o *PublicPaymentChannelDowntimeOnline) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicPaymentChannelDowntimeOnline) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicPaymentChannelDowntimeOnline) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicPaymentChannelDowntimeOnline) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *PublicPaymentChannelDowntimeOnline) GetData() PublicPaymentChannelDowntimeOnlineData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicPaymentChannelDowntimeOnline) GetDataOk() (*PublicPaymentChannelDowntimeOnlineData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicPaymentChannelDowntimeOnline) SetData(v PublicPaymentChannelDowntimeOnlineData)`

SetData sets Data field to given value.

### HasData

`func (o *PublicPaymentChannelDowntimeOnline) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


