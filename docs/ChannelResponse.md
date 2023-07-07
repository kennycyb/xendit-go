# ChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]ChannelWithCountry**](ChannelWithCountry.md) |  | [optional] 
**Errors** | Pointer to [**[]ChannelResponseErrorsInner**](ChannelResponseErrorsInner.md) |  | [optional] 

## Methods

### NewChannelResponse

`func NewChannelResponse() *ChannelResponse`

NewChannelResponse instantiates a new ChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelResponseWithDefaults

`func NewChannelResponseWithDefaults() *ChannelResponse`

NewChannelResponseWithDefaults instantiates a new ChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ChannelResponse) GetChannels() []ChannelWithCountry`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ChannelResponse) GetChannelsOk() (*[]ChannelWithCountry, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ChannelResponse) SetChannels(v []ChannelWithCountry)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetErrors

`func (o *ChannelResponse) GetErrors() []ChannelResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ChannelResponse) GetErrorsOk() (*[]ChannelResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ChannelResponse) SetErrors(v []ChannelResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ChannelResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ChannelResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ChannelResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


