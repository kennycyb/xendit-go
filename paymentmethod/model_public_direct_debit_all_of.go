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

// checks if the PublicDirectDebitAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicDirectDebitAllOf{}

// PublicDirectDebitAllOf struct for PublicDirectDebitAllOf
type PublicDirectDebitAllOf struct {
	Type DirectDebitType `json:"type"`
	BankAccount NullableDirectDebitBankAccount `json:"bank_account,omitempty"`
	DebitCard NullableDirectDebitDebitCard `json:"debit_card,omitempty"`
}

// NewPublicDirectDebitAllOf instantiates a new PublicDirectDebitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDirectDebitAllOf(type_ DirectDebitType) *PublicDirectDebitAllOf {
	this := PublicDirectDebitAllOf{}
	this.Type = type_
	return &this
}

// NewPublicDirectDebitAllOfWithDefaults instantiates a new PublicDirectDebitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDirectDebitAllOfWithDefaults() *PublicDirectDebitAllOf {
	this := PublicDirectDebitAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *PublicDirectDebitAllOf) GetType() DirectDebitType {
	if o == nil {
		var ret DirectDebitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PublicDirectDebitAllOf) GetTypeOk() (*DirectDebitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PublicDirectDebitAllOf) SetType(v DirectDebitType) {
	o.Type = v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicDirectDebitAllOf) GetBankAccount() DirectDebitBankAccount {
	if o == nil || utils.IsNil(o.BankAccount.Get()) {
		var ret DirectDebitBankAccount
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicDirectDebitAllOf) GetBankAccountOk() (*DirectDebitBankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *PublicDirectDebitAllOf) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableDirectDebitBankAccount and assigns it to the BankAccount field.
func (o *PublicDirectDebitAllOf) SetBankAccount(v DirectDebitBankAccount) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *PublicDirectDebitAllOf) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *PublicDirectDebitAllOf) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetDebitCard returns the DebitCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicDirectDebitAllOf) GetDebitCard() DirectDebitDebitCard {
	if o == nil || utils.IsNil(o.DebitCard.Get()) {
		var ret DirectDebitDebitCard
		return ret
	}
	return *o.DebitCard.Get()
}

// GetDebitCardOk returns a tuple with the DebitCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicDirectDebitAllOf) GetDebitCardOk() (*DirectDebitDebitCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitCard.Get(), o.DebitCard.IsSet()
}

// HasDebitCard returns a boolean if a field has been set.
func (o *PublicDirectDebitAllOf) HasDebitCard() bool {
	if o != nil && o.DebitCard.IsSet() {
		return true
	}

	return false
}

// SetDebitCard gets a reference to the given NullableDirectDebitDebitCard and assigns it to the DebitCard field.
func (o *PublicDirectDebitAllOf) SetDebitCard(v DirectDebitDebitCard) {
	o.DebitCard.Set(&v)
}
// SetDebitCardNil sets the value for DebitCard to be an explicit nil
func (o *PublicDirectDebitAllOf) SetDebitCardNil() {
	o.DebitCard.Set(nil)
}

// UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil
func (o *PublicDirectDebitAllOf) UnsetDebitCard() {
	o.DebitCard.Unset()
}

func (o PublicDirectDebitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDirectDebitAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.BankAccount.IsSet() {
		toSerialize["bank_account"] = o.BankAccount.Get()
	}
	if o.DebitCard.IsSet() {
		toSerialize["debit_card"] = o.DebitCard.Get()
	}
	return toSerialize, nil
}

type NullablePublicDirectDebitAllOf struct {
	value *PublicDirectDebitAllOf
	isSet bool
}

func (v NullablePublicDirectDebitAllOf) Get() *PublicDirectDebitAllOf {
	return v.value
}

func (v *NullablePublicDirectDebitAllOf) Set(val *PublicDirectDebitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDirectDebitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDirectDebitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDirectDebitAllOf(val *PublicDirectDebitAllOf) *NullablePublicDirectDebitAllOf {
	return &NullablePublicDirectDebitAllOf{value: val, isSet: true}
}

func (v NullablePublicDirectDebitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDirectDebitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

