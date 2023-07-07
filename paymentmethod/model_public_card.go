/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.79.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package paymentmethod

import (
	"encoding/json"
	
	
	utils "github.com/kennycyb/xendit-go/utils"
	
)

// checks if the PublicCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicCard{}

// PublicCard Card Payment Method Details
type PublicCard struct {
	Currency NullableString `json:"currency"`
	ChannelProperties NullableCardChannelProperties `json:"channel_properties"`
	CardInformation *TokenizedCardInformation `json:"card_information,omitempty"`
	CardVerificationResults NullableCardVerificationResults `json:"card_verification_results,omitempty"`
}

// NewPublicCard instantiates a new PublicCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicCard(currency NullableString, channelProperties NullableCardChannelProperties) *PublicCard {
	this := PublicCard{}
	this.Currency = currency
	this.ChannelProperties = channelProperties
	return &this
}

// NewPublicCardWithDefaults instantiates a new PublicCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicCardWithDefaults() *PublicCard {
	this := PublicCard{}
	return &this
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PublicCard) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicCard) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *PublicCard) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for CardChannelProperties will be returned
func (o *PublicCard) GetChannelProperties() CardChannelProperties {
	if o == nil || o.ChannelProperties.Get() == nil {
		var ret CardChannelProperties
		return ret
	}

	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicCard) GetChannelPropertiesOk() (*CardChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// SetChannelProperties sets field value
func (o *PublicCard) SetChannelProperties(v CardChannelProperties) {
	o.ChannelProperties.Set(&v)
}

// GetCardInformation returns the CardInformation field value if set, zero value otherwise.
func (o *PublicCard) GetCardInformation() TokenizedCardInformation {
	if o == nil || utils.IsNil(o.CardInformation) {
		var ret TokenizedCardInformation
		return ret
	}
	return *o.CardInformation
}

// GetCardInformationOk returns a tuple with the CardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCard) GetCardInformationOk() (*TokenizedCardInformation, bool) {
	if o == nil || utils.IsNil(o.CardInformation) {
		return nil, false
	}
	return o.CardInformation, true
}

// HasCardInformation returns a boolean if a field has been set.
func (o *PublicCard) HasCardInformation() bool {
	if o != nil && !utils.IsNil(o.CardInformation) {
		return true
	}

	return false
}

// SetCardInformation gets a reference to the given TokenizedCardInformation and assigns it to the CardInformation field.
func (o *PublicCard) SetCardInformation(v TokenizedCardInformation) {
	o.CardInformation = &v
}

// GetCardVerificationResults returns the CardVerificationResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicCard) GetCardVerificationResults() CardVerificationResults {
	if o == nil || utils.IsNil(o.CardVerificationResults.Get()) {
		var ret CardVerificationResults
		return ret
	}
	return *o.CardVerificationResults.Get()
}

// GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicCard) GetCardVerificationResultsOk() (*CardVerificationResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardVerificationResults.Get(), o.CardVerificationResults.IsSet()
}

// HasCardVerificationResults returns a boolean if a field has been set.
func (o *PublicCard) HasCardVerificationResults() bool {
	if o != nil && o.CardVerificationResults.IsSet() {
		return true
	}

	return false
}

// SetCardVerificationResults gets a reference to the given NullableCardVerificationResults and assigns it to the CardVerificationResults field.
func (o *PublicCard) SetCardVerificationResults(v CardVerificationResults) {
	o.CardVerificationResults.Set(&v)
}
// SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil
func (o *PublicCard) SetCardVerificationResultsNil() {
	o.CardVerificationResults.Set(nil)
}

// UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
func (o *PublicCard) UnsetCardVerificationResults() {
	o.CardVerificationResults.Unset()
}

func (o PublicCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["channel_properties"] = o.ChannelProperties.Get()
	if !utils.IsNil(o.CardInformation) {
		toSerialize["card_information"] = o.CardInformation
	}
	if o.CardVerificationResults.IsSet() {
		toSerialize["card_verification_results"] = o.CardVerificationResults.Get()
	}
	return toSerialize, nil
}

type NullablePublicCard struct {
	value *PublicCard
	isSet bool
}

func (v NullablePublicCard) Get() *PublicCard {
	return v.value
}

func (v *NullablePublicCard) Set(val *PublicCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicCard(val *PublicCard) *NullablePublicCard {
	return &NullablePublicCard{value: val, isSet: true}
}

func (v NullablePublicCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


