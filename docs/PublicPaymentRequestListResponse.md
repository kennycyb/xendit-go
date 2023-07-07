# PublicPaymentRequestListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PublicPaymentRequest**](PublicPaymentRequest.md) |  | 
**HasMore** | **bool** |  | 

## Methods

### NewPublicPaymentRequestListResponse

`func NewPublicPaymentRequestListResponse(data []PublicPaymentRequest, hasMore bool, ) *PublicPaymentRequestListResponse`

NewPublicPaymentRequestListResponse instantiates a new PublicPaymentRequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentRequestListResponseWithDefaults

`func NewPublicPaymentRequestListResponseWithDefaults() *PublicPaymentRequestListResponse`

NewPublicPaymentRequestListResponseWithDefaults instantiates a new PublicPaymentRequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PublicPaymentRequestListResponse) GetData() []PublicPaymentRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicPaymentRequestListResponse) GetDataOk() (*[]PublicPaymentRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicPaymentRequestListResponse) SetData(v []PublicPaymentRequest)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PublicPaymentRequestListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PublicPaymentRequestListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PublicPaymentRequestListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


