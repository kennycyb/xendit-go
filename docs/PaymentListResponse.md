# PaymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Payment**](Payment.md) |  | 
**HasMore** | **bool** |  | 
**Links** | [**[]Link**](Link.md) |  | 

## Methods

### NewPaymentListResponse

`func NewPaymentListResponse(data []Payment, hasMore bool, links []Link, ) *PaymentListResponse`

NewPaymentListResponse instantiates a new PaymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentListResponseWithDefaults

`func NewPaymentListResponseWithDefaults() *PaymentListResponse`

NewPaymentListResponseWithDefaults instantiates a new PaymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentListResponse) GetData() []Payment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentListResponse) GetDataOk() (*[]Payment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentListResponse) SetData(v []Payment)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PaymentListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PaymentListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PaymentListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLinks

`func (o *PaymentListResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentListResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentListResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


