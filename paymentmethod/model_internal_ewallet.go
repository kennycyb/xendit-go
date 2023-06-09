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

// checks if the InternalEwallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalEwallet{}

// InternalEwallet Ewallet Payment Method Details
type InternalEwallet struct {
	ChannelCode EwalletChannelCode `json:"channel_code"`
	ChannelProperties *EwalletChannelProperties `json:"channel_properties,omitempty"`
	Account *EwalletAccount `json:"account,omitempty"`
	LinkedAccountTokenId NullableString `json:"linked_account_token_id,omitempty"`
	LinkedAccountId NullableString `json:"linked_account_id,omitempty"`
	AccessToken NullableString `json:"access_token,omitempty"`
}

// NewInternalEwallet instantiates a new InternalEwallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalEwallet(channelCode EwalletChannelCode) *InternalEwallet {
	this := InternalEwallet{}
	this.ChannelCode = channelCode
	return &this
}

// NewInternalEwalletWithDefaults instantiates a new InternalEwallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalEwalletWithDefaults() *InternalEwallet {
	this := InternalEwallet{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *InternalEwallet) GetChannelCode() EwalletChannelCode {
	if o == nil {
		var ret EwalletChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *InternalEwallet) GetChannelCodeOk() (*EwalletChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *InternalEwallet) SetChannelCode(v EwalletChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *InternalEwallet) GetChannelProperties() EwalletChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret EwalletChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalEwallet) GetChannelPropertiesOk() (*EwalletChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *InternalEwallet) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given EwalletChannelProperties and assigns it to the ChannelProperties field.
func (o *InternalEwallet) SetChannelProperties(v EwalletChannelProperties) {
	o.ChannelProperties = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *InternalEwallet) GetAccount() EwalletAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret EwalletAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalEwallet) GetAccountOk() (*EwalletAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *InternalEwallet) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given EwalletAccount and assigns it to the Account field.
func (o *InternalEwallet) SetAccount(v EwalletAccount) {
	o.Account = &v
}

// GetLinkedAccountTokenId returns the LinkedAccountTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalEwallet) GetLinkedAccountTokenId() string {
	if o == nil || utils.IsNil(o.LinkedAccountTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedAccountTokenId.Get()
}

// GetLinkedAccountTokenIdOk returns a tuple with the LinkedAccountTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalEwallet) GetLinkedAccountTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedAccountTokenId.Get(), o.LinkedAccountTokenId.IsSet()
}

// HasLinkedAccountTokenId returns a boolean if a field has been set.
func (o *InternalEwallet) HasLinkedAccountTokenId() bool {
	if o != nil && o.LinkedAccountTokenId.IsSet() {
		return true
	}

	return false
}

// SetLinkedAccountTokenId gets a reference to the given NullableString and assigns it to the LinkedAccountTokenId field.
func (o *InternalEwallet) SetLinkedAccountTokenId(v string) {
	o.LinkedAccountTokenId.Set(&v)
}
// SetLinkedAccountTokenIdNil sets the value for LinkedAccountTokenId to be an explicit nil
func (o *InternalEwallet) SetLinkedAccountTokenIdNil() {
	o.LinkedAccountTokenId.Set(nil)
}

// UnsetLinkedAccountTokenId ensures that no value is present for LinkedAccountTokenId, not even an explicit nil
func (o *InternalEwallet) UnsetLinkedAccountTokenId() {
	o.LinkedAccountTokenId.Unset()
}

// GetLinkedAccountId returns the LinkedAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalEwallet) GetLinkedAccountId() string {
	if o == nil || utils.IsNil(o.LinkedAccountId.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedAccountId.Get()
}

// GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalEwallet) GetLinkedAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedAccountId.Get(), o.LinkedAccountId.IsSet()
}

// HasLinkedAccountId returns a boolean if a field has been set.
func (o *InternalEwallet) HasLinkedAccountId() bool {
	if o != nil && o.LinkedAccountId.IsSet() {
		return true
	}

	return false
}

// SetLinkedAccountId gets a reference to the given NullableString and assigns it to the LinkedAccountId field.
func (o *InternalEwallet) SetLinkedAccountId(v string) {
	o.LinkedAccountId.Set(&v)
}
// SetLinkedAccountIdNil sets the value for LinkedAccountId to be an explicit nil
func (o *InternalEwallet) SetLinkedAccountIdNil() {
	o.LinkedAccountId.Set(nil)
}

// UnsetLinkedAccountId ensures that no value is present for LinkedAccountId, not even an explicit nil
func (o *InternalEwallet) UnsetLinkedAccountId() {
	o.LinkedAccountId.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalEwallet) GetAccessToken() string {
	if o == nil || utils.IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalEwallet) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *InternalEwallet) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *InternalEwallet) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *InternalEwallet) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *InternalEwallet) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o InternalEwallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalEwallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
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

type NullableInternalEwallet struct {
	value *InternalEwallet
	isSet bool
}

func (v NullableInternalEwallet) Get() *InternalEwallet {
	return v.value
}

func (v *NullableInternalEwallet) Set(val *InternalEwallet) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalEwallet) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalEwallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalEwallet(val *InternalEwallet) *NullableInternalEwallet {
	return &NullableInternalEwallet{value: val, isSet: true}
}

func (v NullableInternalEwallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalEwallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


