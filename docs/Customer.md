# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CustomerData**](CustomerData.md) |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Customer) GetData() []CustomerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Customer) GetDataOk() (*[]CustomerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Customer) SetData(v []CustomerData)`

SetData sets Data field to given value.

### HasData

`func (o *Customer) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


