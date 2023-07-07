# PublicPaymentChannelDowntimeOnlineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Payment channel health event | [optional] 
**BusinessId** | Pointer to **string** | Business ID registered to payment channel downtime or online callback URL | [optional] 
**Created** | Pointer to **string** | Date time the event is triggered | [optional] 
**Data** | Pointer to [**PublicPaymentChannelDowntimeOnlineData**](PublicPaymentChannelDowntimeOnlineData.md) |  | [optional] 

## Methods

### NewPublicPaymentChannelDowntimeOnlineAllOf

`func NewPublicPaymentChannelDowntimeOnlineAllOf() *PublicPaymentChannelDowntimeOnlineAllOf`

NewPublicPaymentChannelDowntimeOnlineAllOf instantiates a new PublicPaymentChannelDowntimeOnlineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentChannelDowntimeOnlineAllOfWithDefaults

`func NewPublicPaymentChannelDowntimeOnlineAllOfWithDefaults() *PublicPaymentChannelDowntimeOnlineAllOf`

NewPublicPaymentChannelDowntimeOnlineAllOfWithDefaults instantiates a new PublicPaymentChannelDowntimeOnlineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetBusinessId

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreated

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetData() PublicPaymentChannelDowntimeOnlineData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) GetDataOk() (*PublicPaymentChannelDowntimeOnlineData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) SetData(v PublicPaymentChannelDowntimeOnlineData)`

SetData sets Data field to given value.

### HasData

`func (o *PublicPaymentChannelDowntimeOnlineAllOf) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


