# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerMetadata** | Pointer to [**InvoicePartnerMetadata**](InvoicePartnerMetadata.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerMetadata

`func (o *Invoice) GetPartnerMetadata() InvoicePartnerMetadata`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *Invoice) GetPartnerMetadataOk() (*InvoicePartnerMetadata, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *Invoice) SetPartnerMetadata(v InvoicePartnerMetadata)`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *Invoice) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


