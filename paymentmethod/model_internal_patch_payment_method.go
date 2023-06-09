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

// checks if the InternalPatchPaymentMethod type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalPatchPaymentMethod{}

// InternalPatchPaymentMethod struct for InternalPatchPaymentMethod
type InternalPatchPaymentMethod struct {
	Status *PaymentMethodStatus `json:"status,omitempty"`
	Cryptocurrency *PatchCrypto `json:"cryptocurrency,omitempty"`
	OverTheCounter *PatchOverTheCounter `json:"over_the_counter,omitempty"`
}

// NewInternalPatchPaymentMethod instantiates a new InternalPatchPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalPatchPaymentMethod() *InternalPatchPaymentMethod {
	this := InternalPatchPaymentMethod{}
	return &this
}

// NewInternalPatchPaymentMethodWithDefaults instantiates a new InternalPatchPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalPatchPaymentMethodWithDefaults() *InternalPatchPaymentMethod {
	this := InternalPatchPaymentMethod{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalPatchPaymentMethod) GetStatus() PaymentMethodStatus {
	if o == nil || utils.IsNil(o.Status) {
		var ret PaymentMethodStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPatchPaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalPatchPaymentMethod) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentMethodStatus and assigns it to the Status field.
func (o *InternalPatchPaymentMethod) SetStatus(v PaymentMethodStatus) {
	o.Status = &v
}

// GetCryptocurrency returns the Cryptocurrency field value if set, zero value otherwise.
func (o *InternalPatchPaymentMethod) GetCryptocurrency() PatchCrypto {
	if o == nil || utils.IsNil(o.Cryptocurrency) {
		var ret PatchCrypto
		return ret
	}
	return *o.Cryptocurrency
}

// GetCryptocurrencyOk returns a tuple with the Cryptocurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPatchPaymentMethod) GetCryptocurrencyOk() (*PatchCrypto, bool) {
	if o == nil || utils.IsNil(o.Cryptocurrency) {
		return nil, false
	}
	return o.Cryptocurrency, true
}

// HasCryptocurrency returns a boolean if a field has been set.
func (o *InternalPatchPaymentMethod) HasCryptocurrency() bool {
	if o != nil && !utils.IsNil(o.Cryptocurrency) {
		return true
	}

	return false
}

// SetCryptocurrency gets a reference to the given PatchCrypto and assigns it to the Cryptocurrency field.
func (o *InternalPatchPaymentMethod) SetCryptocurrency(v PatchCrypto) {
	o.Cryptocurrency = &v
}

// GetOverTheCounter returns the OverTheCounter field value if set, zero value otherwise.
func (o *InternalPatchPaymentMethod) GetOverTheCounter() PatchOverTheCounter {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		var ret PatchOverTheCounter
		return ret
	}
	return *o.OverTheCounter
}

// GetOverTheCounterOk returns a tuple with the OverTheCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalPatchPaymentMethod) GetOverTheCounterOk() (*PatchOverTheCounter, bool) {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		return nil, false
	}
	return o.OverTheCounter, true
}

// HasOverTheCounter returns a boolean if a field has been set.
func (o *InternalPatchPaymentMethod) HasOverTheCounter() bool {
	if o != nil && !utils.IsNil(o.OverTheCounter) {
		return true
	}

	return false
}

// SetOverTheCounter gets a reference to the given PatchOverTheCounter and assigns it to the OverTheCounter field.
func (o *InternalPatchPaymentMethod) SetOverTheCounter(v PatchOverTheCounter) {
	o.OverTheCounter = &v
}

func (o InternalPatchPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalPatchPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Cryptocurrency) {
		toSerialize["cryptocurrency"] = o.Cryptocurrency
	}
	if !utils.IsNil(o.OverTheCounter) {
		toSerialize["over_the_counter"] = o.OverTheCounter
	}
	return toSerialize, nil
}

type NullableInternalPatchPaymentMethod struct {
	value *InternalPatchPaymentMethod
	isSet bool
}

func (v NullableInternalPatchPaymentMethod) Get() *InternalPatchPaymentMethod {
	return v.value
}

func (v *NullableInternalPatchPaymentMethod) Set(val *InternalPatchPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalPatchPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalPatchPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalPatchPaymentMethod(val *InternalPatchPaymentMethod) *NullableInternalPatchPaymentMethod {
	return &NullableInternalPatchPaymentMethod{value: val, isSet: true}
}

func (v NullableInternalPatchPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalPatchPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


