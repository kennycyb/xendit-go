# PublicPaymentChannelListLinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Target URI that should contain a target to Internationalized Resource Identifiers (IRI) | [optional] 
**Rel** | Pointer to **string** | The link relation type described how the current context (source) is related to target resource | [optional] 
**Method** | Pointer to **string** | The HTTP method to be used to call &#x60;href&#x60; | [optional] 

## Methods

### NewPublicPaymentChannelListLinksInner

`func NewPublicPaymentChannelListLinksInner() *PublicPaymentChannelListLinksInner`

NewPublicPaymentChannelListLinksInner instantiates a new PublicPaymentChannelListLinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPaymentChannelListLinksInnerWithDefaults

`func NewPublicPaymentChannelListLinksInnerWithDefaults() *PublicPaymentChannelListLinksInner`

NewPublicPaymentChannelListLinksInnerWithDefaults instantiates a new PublicPaymentChannelListLinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PublicPaymentChannelListLinksInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PublicPaymentChannelListLinksInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PublicPaymentChannelListLinksInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PublicPaymentChannelListLinksInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetRel

`func (o *PublicPaymentChannelListLinksInner) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *PublicPaymentChannelListLinksInner) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *PublicPaymentChannelListLinksInner) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *PublicPaymentChannelListLinksInner) HasRel() bool`

HasRel returns a boolean if a field has been set.

### GetMethod

`func (o *PublicPaymentChannelListLinksInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PublicPaymentChannelListLinksInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PublicPaymentChannelListLinksInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PublicPaymentChannelListLinksInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


