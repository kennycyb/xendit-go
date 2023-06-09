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

// checks if the EwalletAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EwalletAccount{}

// EwalletAccount Ewallet Account Properties
type EwalletAccount struct {
	// Name of the eWallet account holder. The value is null if unavailableName of the eWallet account holder. The value is null if unavailable
	Name NullableString `json:"name,omitempty"`
	// Identifier from eWallet provider e.g. phone number. The value is null if unavailable
	AccountDetails NullableString `json:"account_details,omitempty"`
	// The main balance amount on eWallet account provided from eWallet provider. The value is null if unavailable
	Balance NullableFloat64 `json:"balance,omitempty"`
	// The point balance amount on eWallet account. Applicable only on some eWallet provider that has point system. The value is null if unavailabl
	PointBalance NullableFloat64 `json:"point_balance,omitempty"`
}

// NewEwalletAccount instantiates a new EwalletAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEwalletAccount() *EwalletAccount {
	this := EwalletAccount{}
	return &this
}

// NewEwalletAccountWithDefaults instantiates a new EwalletAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEwalletAccountWithDefaults() *EwalletAccount {
	this := EwalletAccount{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EwalletAccount) GetName() string {
	if o == nil || utils.IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EwalletAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EwalletAccount) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EwalletAccount) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EwalletAccount) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EwalletAccount) UnsetName() {
	o.Name.Unset()
}

// GetAccountDetails returns the AccountDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EwalletAccount) GetAccountDetails() string {
	if o == nil || utils.IsNil(o.AccountDetails.Get()) {
		var ret string
		return ret
	}
	return *o.AccountDetails.Get()
}

// GetAccountDetailsOk returns a tuple with the AccountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EwalletAccount) GetAccountDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountDetails.Get(), o.AccountDetails.IsSet()
}

// HasAccountDetails returns a boolean if a field has been set.
func (o *EwalletAccount) HasAccountDetails() bool {
	if o != nil && o.AccountDetails.IsSet() {
		return true
	}

	return false
}

// SetAccountDetails gets a reference to the given NullableString and assigns it to the AccountDetails field.
func (o *EwalletAccount) SetAccountDetails(v string) {
	o.AccountDetails.Set(&v)
}
// SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil
func (o *EwalletAccount) SetAccountDetailsNil() {
	o.AccountDetails.Set(nil)
}

// UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
func (o *EwalletAccount) UnsetAccountDetails() {
	o.AccountDetails.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EwalletAccount) GetBalance() float64 {
	if o == nil || utils.IsNil(o.Balance.Get()) {
		var ret float64
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EwalletAccount) GetBalanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *EwalletAccount) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableFloat64 and assigns it to the Balance field.
func (o *EwalletAccount) SetBalance(v float64) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *EwalletAccount) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *EwalletAccount) UnsetBalance() {
	o.Balance.Unset()
}

// GetPointBalance returns the PointBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EwalletAccount) GetPointBalance() float64 {
	if o == nil || utils.IsNil(o.PointBalance.Get()) {
		var ret float64
		return ret
	}
	return *o.PointBalance.Get()
}

// GetPointBalanceOk returns a tuple with the PointBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EwalletAccount) GetPointBalanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointBalance.Get(), o.PointBalance.IsSet()
}

// HasPointBalance returns a boolean if a field has been set.
func (o *EwalletAccount) HasPointBalance() bool {
	if o != nil && o.PointBalance.IsSet() {
		return true
	}

	return false
}

// SetPointBalance gets a reference to the given NullableFloat64 and assigns it to the PointBalance field.
func (o *EwalletAccount) SetPointBalance(v float64) {
	o.PointBalance.Set(&v)
}
// SetPointBalanceNil sets the value for PointBalance to be an explicit nil
func (o *EwalletAccount) SetPointBalanceNil() {
	o.PointBalance.Set(nil)
}

// UnsetPointBalance ensures that no value is present for PointBalance, not even an explicit nil
func (o *EwalletAccount) UnsetPointBalance() {
	o.PointBalance.Unset()
}

func (o EwalletAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EwalletAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AccountDetails.IsSet() {
		toSerialize["account_details"] = o.AccountDetails.Get()
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.PointBalance.IsSet() {
		toSerialize["point_balance"] = o.PointBalance.Get()
	}
	return toSerialize, nil
}

type NullableEwalletAccount struct {
	value *EwalletAccount
	isSet bool
}

func (v NullableEwalletAccount) Get() *EwalletAccount {
	return v.value
}

func (v *NullableEwalletAccount) Set(val *EwalletAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableEwalletAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableEwalletAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEwalletAccount(val *EwalletAccount) *NullableEwalletAccount {
	return &NullableEwalletAccount{value: val, isSet: true}
}

func (v NullableEwalletAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEwalletAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


