# InternalDirectBankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | [**DirectBankTransferChannelCode**](DirectBankTransferChannelCode.md) |  | 
**ChannelProperties** | [**DirectBankTransferChannelProperties**](DirectBankTransferChannelProperties.md) |  | 
**LinkedTokenId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInternalDirectBankTransfer

`func NewInternalDirectBankTransfer(channelCode DirectBankTransferChannelCode, channelProperties DirectBankTransferChannelProperties, ) *InternalDirectBankTransfer`

NewInternalDirectBankTransfer instantiates a new InternalDirectBankTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDirectBankTransferWithDefaults

`func NewInternalDirectBankTransferWithDefaults() *InternalDirectBankTransfer`

NewInternalDirectBankTransferWithDefaults instantiates a new InternalDirectBankTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *InternalDirectBankTransfer) GetChannelCode() DirectBankTransferChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *InternalDirectBankTransfer) GetChannelCodeOk() (*DirectBankTransferChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *InternalDirectBankTransfer) SetChannelCode(v DirectBankTransferChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *InternalDirectBankTransfer) GetChannelProperties() DirectBankTransferChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *InternalDirectBankTransfer) GetChannelPropertiesOk() (*DirectBankTransferChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *InternalDirectBankTransfer) SetChannelProperties(v DirectBankTransferChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetLinkedTokenId

`func (o *InternalDirectBankTransfer) GetLinkedTokenId() string`

GetLinkedTokenId returns the LinkedTokenId field if non-nil, zero value otherwise.

### GetLinkedTokenIdOk

`func (o *InternalDirectBankTransfer) GetLinkedTokenIdOk() (*string, bool)`

GetLinkedTokenIdOk returns a tuple with the LinkedTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTokenId

`func (o *InternalDirectBankTransfer) SetLinkedTokenId(v string)`

SetLinkedTokenId sets LinkedTokenId field to given value.

### HasLinkedTokenId

`func (o *InternalDirectBankTransfer) HasLinkedTokenId() bool`

HasLinkedTokenId returns a boolean if a field has been set.

### SetLinkedTokenIdNil

`func (o *InternalDirectBankTransfer) SetLinkedTokenIdNil(b bool)`

 SetLinkedTokenIdNil sets the value for LinkedTokenId to be an explicit nil

### UnsetLinkedTokenId
`func (o *InternalDirectBankTransfer) UnsetLinkedTokenId()`

UnsetLinkedTokenId ensures that no value is present for LinkedTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


