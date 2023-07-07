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

// checks if the InternalDirectDebit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalDirectDebit{}

// InternalDirectDebit Direct Debit Payment Method Details
type InternalDirectDebit struct {
	ChannelCode DirectDebitChannelCode `json:"channel_code"`
	ChannelProperties NullableDirectDebitChannelProperties `json:"channel_properties"`
	Type DirectDebitType `json:"type"`
	BankAccount NullableDirectDebitBankAccount `json:"bank_account,omitempty"`
	DebitCard NullableDirectDebitDebitCard `json:"debit_card,omitempty"`
	LinkedAccountTokenId NullableString `json:"linked_account_token_id,omitempty"`
	LinkedAccountId NullableString `json:"linked_account_id,omitempty"`
	AccessToken NullableString `json:"access_token,omitempty"`
}

// NewInternalDirectDebit instantiates a new InternalDirectDebit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDirectDebit(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, type_ DirectDebitType) *InternalDirectDebit {
	this := InternalDirectDebit{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	this.Type = type_
	return &this
}

// NewInternalDirectDebitWithDefaults instantiates a new InternalDirectDebit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDirectDebitWithDefaults() *InternalDirectDebit {
	this := InternalDirectDebit{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *InternalDirectDebit) GetChannelCode() DirectDebitChannelCode {
	if o == nil {
		var ret DirectDebitChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *InternalDirectDebit) GetChannelCodeOk() (*DirectDebitChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *InternalDirectDebit) SetChannelCode(v DirectDebitChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for DirectDebitChannelProperties will be returned
func (o *InternalDirectDebit) GetChannelProperties() DirectDebitChannelProperties {
	if o == nil || o.ChannelProperties.Get() == nil {
		var ret DirectDebitChannelProperties
		return ret
	}

	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// SetChannelProperties sets field value
func (o *InternalDirectDebit) SetChannelProperties(v DirectDebitChannelProperties) {
	o.ChannelProperties.Set(&v)
}

// GetType returns the Type field value
func (o *InternalDirectDebit) GetType() DirectDebitType {
	if o == nil {
		var ret DirectDebitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InternalDirectDebit) GetTypeOk() (*DirectDebitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InternalDirectDebit) SetType(v DirectDebitType) {
	o.Type = v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectDebit) GetBankAccount() DirectDebitBankAccount {
	if o == nil || utils.IsNil(o.BankAccount.Get()) {
		var ret DirectDebitBankAccount
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetBankAccountOk() (*DirectDebitBankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *InternalDirectDebit) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableDirectDebitBankAccount and assigns it to the BankAccount field.
func (o *InternalDirectDebit) SetBankAccount(v DirectDebitBankAccount) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *InternalDirectDebit) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *InternalDirectDebit) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetDebitCard returns the DebitCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectDebit) GetDebitCard() DirectDebitDebitCard {
	if o == nil || utils.IsNil(o.DebitCard.Get()) {
		var ret DirectDebitDebitCard
		return ret
	}
	return *o.DebitCard.Get()
}

// GetDebitCardOk returns a tuple with the DebitCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetDebitCardOk() (*DirectDebitDebitCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitCard.Get(), o.DebitCard.IsSet()
}

// HasDebitCard returns a boolean if a field has been set.
func (o *InternalDirectDebit) HasDebitCard() bool {
	if o != nil && o.DebitCard.IsSet() {
		return true
	}

	return false
}

// SetDebitCard gets a reference to the given NullableDirectDebitDebitCard and assigns it to the DebitCard field.
func (o *InternalDirectDebit) SetDebitCard(v DirectDebitDebitCard) {
	o.DebitCard.Set(&v)
}
// SetDebitCardNil sets the value for DebitCard to be an explicit nil
func (o *InternalDirectDebit) SetDebitCardNil() {
	o.DebitCard.Set(nil)
}

// UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil
func (o *InternalDirectDebit) UnsetDebitCard() {
	o.DebitCard.Unset()
}

// GetLinkedAccountTokenId returns the LinkedAccountTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectDebit) GetLinkedAccountTokenId() string {
	if o == nil || utils.IsNil(o.LinkedAccountTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedAccountTokenId.Get()
}

// GetLinkedAccountTokenIdOk returns a tuple with the LinkedAccountTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetLinkedAccountTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedAccountTokenId.Get(), o.LinkedAccountTokenId.IsSet()
}

// HasLinkedAccountTokenId returns a boolean if a field has been set.
func (o *InternalDirectDebit) HasLinkedAccountTokenId() bool {
	if o != nil && o.LinkedAccountTokenId.IsSet() {
		return true
	}

	return false
}

// SetLinkedAccountTokenId gets a reference to the given NullableString and assigns it to the LinkedAccountTokenId field.
func (o *InternalDirectDebit) SetLinkedAccountTokenId(v string) {
	o.LinkedAccountTokenId.Set(&v)
}
// SetLinkedAccountTokenIdNil sets the value for LinkedAccountTokenId to be an explicit nil
func (o *InternalDirectDebit) SetLinkedAccountTokenIdNil() {
	o.LinkedAccountTokenId.Set(nil)
}

// UnsetLinkedAccountTokenId ensures that no value is present for LinkedAccountTokenId, not even an explicit nil
func (o *InternalDirectDebit) UnsetLinkedAccountTokenId() {
	o.LinkedAccountTokenId.Unset()
}

// GetLinkedAccountId returns the LinkedAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectDebit) GetLinkedAccountId() string {
	if o == nil || utils.IsNil(o.LinkedAccountId.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedAccountId.Get()
}

// GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetLinkedAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedAccountId.Get(), o.LinkedAccountId.IsSet()
}

// HasLinkedAccountId returns a boolean if a field has been set.
func (o *InternalDirectDebit) HasLinkedAccountId() bool {
	if o != nil && o.LinkedAccountId.IsSet() {
		return true
	}

	return false
}

// SetLinkedAccountId gets a reference to the given NullableString and assigns it to the LinkedAccountId field.
func (o *InternalDirectDebit) SetLinkedAccountId(v string) {
	o.LinkedAccountId.Set(&v)
}
// SetLinkedAccountIdNil sets the value for LinkedAccountId to be an explicit nil
func (o *InternalDirectDebit) SetLinkedAccountIdNil() {
	o.LinkedAccountId.Set(nil)
}

// UnsetLinkedAccountId ensures that no value is present for LinkedAccountId, not even an explicit nil
func (o *InternalDirectDebit) UnsetLinkedAccountId() {
	o.LinkedAccountId.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalDirectDebit) GetAccessToken() string {
	if o == nil || utils.IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalDirectDebit) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *InternalDirectDebit) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *InternalDirectDebit) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *InternalDirectDebit) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *InternalDirectDebit) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o InternalDirectDebit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDirectDebit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties.Get()
	toSerialize["type"] = o.Type
	if o.BankAccount.IsSet() {
		toSerialize["bank_account"] = o.BankAccount.Get()
	}
	if o.DebitCard.IsSet() {
		toSerialize["debit_card"] = o.DebitCard.Get()
	}
	if o.LinkedAccountTokenId.IsSet() {
		toSerialize["linked_account_token_id"] = o.LinkedAccountTokenId.Get()
	}
	if o.LinkedAccountId.IsSet() {
		toSerialize["linked_account_id"] = o.LinkedAccountId.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	return toSerialize, nil
}

type NullableInternalDirectDebit struct {
	value *InternalDirectDebit
	isSet bool
}

func (v NullableInternalDirectDebit) Get() *InternalDirectDebit {
	return v.value
}

func (v *NullableInternalDirectDebit) Set(val *InternalDirectDebit) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDirectDebit) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDirectDebit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDirectDebit(val *InternalDirectDebit) *NullableInternalDirectDebit {
	return &NullableInternalDirectDebit{value: val, isSet: true}
}

func (v NullableInternalDirectDebit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDirectDebit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

