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

// checks if the CardTokenInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardTokenInfo{}

// CardTokenInfo Card Token Info
type CardTokenInfo struct {
	Bank *string `json:"bank,omitempty"`
	Country *string `json:"country,omitempty"`
	Type *string `json:"type,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CardArtUrl *string `json:"card_art_url,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// NewCardTokenInfo instantiates a new CardTokenInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTokenInfo() *CardTokenInfo {
	this := CardTokenInfo{}
	return &this
}

// NewCardTokenInfoWithDefaults instantiates a new CardTokenInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTokenInfoWithDefaults() *CardTokenInfo {
	this := CardTokenInfo{}
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *CardTokenInfo) GetBank() string {
	if o == nil || utils.IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetBankOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *CardTokenInfo) HasBank() bool {
	if o != nil && !utils.IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *CardTokenInfo) SetBank(v string) {
	o.Bank = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CardTokenInfo) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CardTokenInfo) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CardTokenInfo) SetCountry(v string) {
	o.Country = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CardTokenInfo) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CardTokenInfo) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CardTokenInfo) SetType(v string) {
	o.Type = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *CardTokenInfo) GetBrand() string {
	if o == nil || utils.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetBrandOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *CardTokenInfo) HasBrand() bool {
	if o != nil && !utils.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *CardTokenInfo) SetBrand(v string) {
	o.Brand = &v
}

// GetCardArtUrl returns the CardArtUrl field value if set, zero value otherwise.
func (o *CardTokenInfo) GetCardArtUrl() string {
	if o == nil || utils.IsNil(o.CardArtUrl) {
		var ret string
		return ret
	}
	return *o.CardArtUrl
}

// GetCardArtUrlOk returns a tuple with the CardArtUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetCardArtUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardArtUrl) {
		return nil, false
	}
	return o.CardArtUrl, true
}

// HasCardArtUrl returns a boolean if a field has been set.
func (o *CardTokenInfo) HasCardArtUrl() bool {
	if o != nil && !utils.IsNil(o.CardArtUrl) {
		return true
	}

	return false
}

// SetCardArtUrl gets a reference to the given string and assigns it to the CardArtUrl field.
func (o *CardTokenInfo) SetCardArtUrl(v string) {
	o.CardArtUrl = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CardTokenInfo) GetFingerprint() string {
	if o == nil || utils.IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenInfo) GetFingerprintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CardTokenInfo) HasFingerprint() bool {
	if o != nil && !utils.IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CardTokenInfo) SetFingerprint(v string) {
	o.Fingerprint = &v
}

func (o CardTokenInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardTokenInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !utils.IsNil(o.CardArtUrl) {
		toSerialize["card_art_url"] = o.CardArtUrl
	}
	if !utils.IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	return toSerialize, nil
}

type NullableCardTokenInfo struct {
	value *CardTokenInfo
	isSet bool
}

func (v NullableCardTokenInfo) Get() *CardTokenInfo {
	return v.value
}

func (v *NullableCardTokenInfo) Set(val *CardTokenInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTokenInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTokenInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTokenInfo(val *CardTokenInfo) *NullableCardTokenInfo {
	return &NullableCardTokenInfo{value: val, isSet: true}
}

func (v NullableCardTokenInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTokenInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


