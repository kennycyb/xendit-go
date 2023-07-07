# CardToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PaymentMethodId** | **string** |  | 
**ReferenceId** | **string** |  | 
**Status** | **string** |  | 
**Currency** | **string** |  | 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**MaskedCardNumber** | Pointer to **NullableString** |  | [optional] 
**ExpiryMonth** | Pointer to **NullableString** |  | [optional] 
**ExpiryYear** | Pointer to **NullableString** |  | [optional] 
**CardInfo** | Pointer to [**NullableCardTokenInfo**](CardTokenInfo.md) |  | [optional] 
**BillingDetails** | Pointer to [**NullableCardBillingDetails**](CardBillingDetails.md) |  | [optional] 

## Methods

### NewCardToken

`func NewCardToken(id string, paymentMethodId string, referenceId string, status string, currency string, ) *CardToken`

NewCardToken instantiates a new CardToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenWithDefaults

`func NewCardTokenWithDefaults() *CardToken`

NewCardTokenWithDefaults instantiates a new CardToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardToken) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentMethodId

`func (o *CardToken) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CardToken) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CardToken) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetReferenceId

`func (o *CardToken) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CardToken) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CardToken) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetStatus

`func (o *CardToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardToken) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *CardToken) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CardToken) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CardToken) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRedirectUrl

`func (o *CardToken) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CardToken) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CardToken) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CardToken) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CardToken) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CardToken) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetMaskedCardNumber

`func (o *CardToken) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *CardToken) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *CardToken) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.

### HasMaskedCardNumber

`func (o *CardToken) HasMaskedCardNumber() bool`

HasMaskedCardNumber returns a boolean if a field has been set.

### SetMaskedCardNumberNil

`func (o *CardToken) SetMaskedCardNumberNil(b bool)`

 SetMaskedCardNumberNil sets the value for MaskedCardNumber to be an explicit nil

### UnsetMaskedCardNumber
`func (o *CardToken) UnsetMaskedCardNumber()`

UnsetMaskedCardNumber ensures that no value is present for MaskedCardNumber, not even an explicit nil
### GetExpiryMonth

`func (o *CardToken) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CardToken) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CardToken) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *CardToken) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### SetExpiryMonthNil

`func (o *CardToken) SetExpiryMonthNil(b bool)`

 SetExpiryMonthNil sets the value for ExpiryMonth to be an explicit nil

### UnsetExpiryMonth
`func (o *CardToken) UnsetExpiryMonth()`

UnsetExpiryMonth ensures that no value is present for ExpiryMonth, not even an explicit nil
### GetExpiryYear

`func (o *CardToken) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CardToken) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CardToken) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *CardToken) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### SetExpiryYearNil

`func (o *CardToken) SetExpiryYearNil(b bool)`

 SetExpiryYearNil sets the value for ExpiryYear to be an explicit nil

### UnsetExpiryYear
`func (o *CardToken) UnsetExpiryYear()`

UnsetExpiryYear ensures that no value is present for ExpiryYear, not even an explicit nil
### GetCardInfo

`func (o *CardToken) GetCardInfo() CardTokenInfo`

GetCardInfo returns the CardInfo field if non-nil, zero value otherwise.

### GetCardInfoOk

`func (o *CardToken) GetCardInfoOk() (*CardTokenInfo, bool)`

GetCardInfoOk returns a tuple with the CardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInfo

`func (o *CardToken) SetCardInfo(v CardTokenInfo)`

SetCardInfo sets CardInfo field to given value.

### HasCardInfo

`func (o *CardToken) HasCardInfo() bool`

HasCardInfo returns a boolean if a field has been set.

### SetCardInfoNil

`func (o *CardToken) SetCardInfoNil(b bool)`

 SetCardInfoNil sets the value for CardInfo to be an explicit nil

### UnsetCardInfo
`func (o *CardToken) UnsetCardInfo()`

UnsetCardInfo ensures that no value is present for CardInfo, not even an explicit nil
### GetBillingDetails

`func (o *CardToken) GetBillingDetails() CardBillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *CardToken) GetBillingDetailsOk() (*CardBillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *CardToken) SetBillingDetails(v CardBillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *CardToken) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *CardToken) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *CardToken) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


