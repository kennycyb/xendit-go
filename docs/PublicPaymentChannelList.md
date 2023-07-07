# PublicPaymentChannelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PublicPaymentChannel**](PublicPaymentChannel.md) | Array of resources that match the provided filters | [optional] 
**Links** | Pointer to [**[]PublicPaymentChannelListLinksInner**](PublicPaymentChannelListLinksInner.md) | Array of objects that can be used to assist on navigating through the data | [optional] 
**HasMore** | Pointer to **bool** | Indicates whether there are more items in the list | [optional] 

## Methods

### NewPublicPaymentChannelList

`func NewPublicPaymentChannelList() *PublicPaymentChannelList`

NewPublicPaymentChannelList instantiates a new PublicPaymentChannelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentChannelListWithDefaults

`func NewPublicPaymentChannelListWithDefaults() *PublicPaymentChannelList`

NewPublicPaymentChannelListWithDefaults instantiates a new PublicPaymentChannelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PublicPaymentChannelList) GetData() []PublicPaymentChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicPaymentChannelList) GetDataOk() (*[]PublicPaymentChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicPaymentChannelList) SetData(v []PublicPaymentChannel)`

SetData sets Data field to given value.

### HasData

`func (o *PublicPaymentChannelList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *PublicPaymentChannelList) GetLinks() []PublicPaymentChannelListLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PublicPaymentChannelList) GetLinksOk() (*[]PublicPaymentChannelListLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PublicPaymentChannelList) SetLinks(v []PublicPaymentChannelListLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PublicPaymentChannelList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetHasMore

`func (o *PublicPaymentChannelList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PublicPaymentChannelList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PublicPaymentChannelList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *PublicPaymentChannelList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


