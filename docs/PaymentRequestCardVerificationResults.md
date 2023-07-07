# PaymentRequestCardVerificationResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSecure** | [**PaymentRequestCardVerificationResultsThreeDeeSecure**](PaymentRequestCardVerificationResultsThreeDeeSecure.md) |  | 
**CvvResult** | Pointer to **string** |  | [optional] 
**AddressVerificationResult** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentRequestCardVerificationResults

`func NewPaymentRequestCardVerificationResults(threeDSecure PaymentRequestCardVerificationResultsThreeDeeSecure, ) *PaymentRequestCardVerificationResults`

NewPaymentRequestCardVerificationResults instantiates a new PaymentRequestCardVerificationResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestCardVerificationResultsWithDefaults

`func NewPaymentRequestCardVerificationResultsWithDefaults() *PaymentRequestCardVerificationResults`

NewPaymentRequestCardVerificationResultsWithDefaults instantiates a new PaymentRequestCardVerificationResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSecure

`func (o *PaymentRequestCardVerificationResults) GetThreeDSecure() PaymentRequestCardVerificationResultsThreeDeeSecure`

GetThreeDSecure returns the ThreeDSecure field if non-nil, zero value otherwise.

### GetThreeDSecureOk

`func (o *PaymentRequestCardVerificationResults) GetThreeDSecureOk() (*PaymentRequestCardVerificationResultsThreeDeeSecure, bool)`

GetThreeDSecureOk returns a tuple with the ThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecure

`func (o *PaymentRequestCardVerificationResults) SetThreeDSecure(v PaymentRequestCardVerificationResultsThreeDeeSecure)`

SetThreeDSecure sets ThreeDSecure field to given value.


### GetCvvResult

`func (o *PaymentRequestCardVerificationResults) GetCvvResult() string`

GetCvvResult returns the CvvResult field if non-nil, zero value otherwise.

### GetCvvResultOk

`func (o *PaymentRequestCardVerificationResults) GetCvvResultOk() (*string, bool)`

GetCvvResultOk returns a tuple with the CvvResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResult

`func (o *PaymentRequestCardVerificationResults) SetCvvResult(v string)`

SetCvvResult sets CvvResult field to given value.

### HasCvvResult

`func (o *PaymentRequestCardVerificationResults) HasCvvResult() bool`

HasCvvResult returns a boolean if a field has been set.

### GetAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.

### HasAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) HasAddressVerificationResult() bool`

HasAddressVerificationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


