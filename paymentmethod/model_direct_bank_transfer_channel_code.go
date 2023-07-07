/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.79.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package paymentmethod

import (
	"encoding/json"
	
	"fmt"
)

// DirectBankTransferChannelCode Direct Bank Transfer Channel Code
type DirectBankTransferChannelCode string

// List of DirectBankTransferChannelCode
const (
	DIRECTBANKTRANSFERCHANNELCODE_BDO DirectBankTransferChannelCode = "DP_BDO"
	DIRECTBANKTRANSFERCHANNELCODE_MBT DirectBankTransferChannelCode = "DP_MBT"
)

// All allowed values of DirectBankTransferChannelCode enum
var AllowedDirectBankTransferChannelCodeEnumValues = []DirectBankTransferChannelCode{
	"DP_BDO",
	"DP_MBT",
}

func (v *DirectBankTransferChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectBankTransferChannelCode(value)
	for _, existing := range AllowedDirectBankTransferChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DirectBankTransferChannelCode", value)
}

// NewDirectBankTransferChannelCodeFromValue returns a pointer to a valid DirectBankTransferChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectBankTransferChannelCodeFromValue(v string) (*DirectBankTransferChannelCode, error) {
	ev := DirectBankTransferChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DirectBankTransferChannelCode: valid values are %v", v, AllowedDirectBankTransferChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DirectBankTransferChannelCode) IsValid() bool {
	for _, existing := range AllowedDirectBankTransferChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DirectBankTransferChannelCode value
func (v DirectBankTransferChannelCode) Ptr() *DirectBankTransferChannelCode {
	return &v
}

type NullableDirectBankTransferChannelCode struct {
	value *DirectBankTransferChannelCode
	isSet bool
}

func (v NullableDirectBankTransferChannelCode) Get() *DirectBankTransferChannelCode {
	return v.value
}

func (v *NullableDirectBankTransferChannelCode) Set(val *DirectBankTransferChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectBankTransferChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectBankTransferChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectBankTransferChannelCode(val *DirectBankTransferChannelCode) *NullableDirectBankTransferChannelCode {
	return &NullableDirectBankTransferChannelCode{value: val, isSet: true}
}

func (v NullableDirectBankTransferChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectBankTransferChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

