# InternalLATLinked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**ChannelCode** | **string** |  | 
**Id** | **string** |  | 
**Type** | **string** |  | 
**Accounts** | [**[]InternalLATLinkedAccountsInner**](InternalLATLinkedAccountsInner.md) |  | 

## Methods

### NewInternalLATLinked

`func NewInternalLATLinked(event string, timestamp time.Time, channelCode string, id string, type_ string, accounts []InternalLATLinkedAccountsInner, ) *InternalLATLinked`

NewInternalLATLinked instantiates a new InternalLATLinked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalLATLinkedWithDefaults

`func NewInternalLATLinkedWithDefaults() *InternalLATLinked`

NewInternalLATLinkedWithDefaults instantiates a new InternalLATLinked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *InternalLATLinked) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *InternalLATLinked) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *InternalLATLinked) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTimestamp

`func (o *InternalLATLinked) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InternalLATLinked) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InternalLATLinked) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetChannelCode

`func (o *InternalLATLinked) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalLATLinked) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalLATLinked) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetId

`func (o *InternalLATLinked) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalLATLinked) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalLATLinked) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InternalLATLinked) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalLATLinked) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalLATLinked) SetType(v string)`

SetType sets Type field to given value.


### GetAccounts

`func (o *InternalLATLinked) GetAccounts() []InternalLATLinkedAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *InternalLATLinked) GetAccountsOk() (*[]InternalLATLinkedAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *InternalLATLinked) SetAccounts(v []InternalLATLinkedAccountsInner)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


