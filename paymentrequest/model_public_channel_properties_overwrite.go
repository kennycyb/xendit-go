/*
Charge

This API is used for Charge

API version: 1.31.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package paymentrequest

import (
	"encoding/json"
	
	
	utils "github.com/kennycyb/xendit-go/utils"
	
)

// checks if the PublicChannelPropertiesOverwrite type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PublicChannelPropertiesOverwrite{}

// PublicChannelPropertiesOverwrite struct for PublicChannelPropertiesOverwrite
type PublicChannelPropertiesOverwrite struct {
	// URL where the end-customer is redirected if the authorization is successful
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization failed
	FailureReturnUrl *string `json:"failure_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization cancelled
	CancelReturnUrl *string `json:"cancel_return_url,omitempty"`
	// REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only.
	RedeemPoints *string `json:"redeem_points,omitempty"`
	// Toggle used to require end-customer to input undergo OTP validation before completing a payment. OTP will always be required for transactions greater than 1,000,000 IDR. For BRI tokenized payment use only.
	RequireAuth *bool `json:"require_auth,omitempty"`
	// Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use
	MerchantIdTag *string `json:"merchant_id_tag,omitempty"`
	// Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging.
	CardonfileType *string `json:"cardonfile_type,omitempty"`
}

// NewPublicChannelPropertiesOverwrite instantiates a new PublicChannelPropertiesOverwrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicChannelPropertiesOverwrite() *PublicChannelPropertiesOverwrite {
	this := PublicChannelPropertiesOverwrite{}
	return &this
}

// NewPublicChannelPropertiesOverwriteWithDefaults instantiates a new PublicChannelPropertiesOverwrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicChannelPropertiesOverwriteWithDefaults() *PublicChannelPropertiesOverwrite {
	this := PublicChannelPropertiesOverwrite{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *PublicChannelPropertiesOverwrite) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		return nil, false
	}
	return o.FailureReturnUrl, true
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasFailureReturnUrl() bool {
	if o != nil && !utils.IsNil(o.FailureReturnUrl) {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given string and assigns it to the FailureReturnUrl field.
func (o *PublicChannelPropertiesOverwrite) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl = &v
}

// GetCancelReturnUrl returns the CancelReturnUrl field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetCancelReturnUrl() string {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		var ret string
		return ret
	}
	return *o.CancelReturnUrl
}

// GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetCancelReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		return nil, false
	}
	return o.CancelReturnUrl, true
}

// HasCancelReturnUrl returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasCancelReturnUrl() bool {
	if o != nil && !utils.IsNil(o.CancelReturnUrl) {
		return true
	}

	return false
}

// SetCancelReturnUrl gets a reference to the given string and assigns it to the CancelReturnUrl field.
func (o *PublicChannelPropertiesOverwrite) SetCancelReturnUrl(v string) {
	o.CancelReturnUrl = &v
}

// GetRedeemPoints returns the RedeemPoints field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetRedeemPoints() string {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		var ret string
		return ret
	}
	return *o.RedeemPoints
}

// GetRedeemPointsOk returns a tuple with the RedeemPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetRedeemPointsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		return nil, false
	}
	return o.RedeemPoints, true
}

// HasRedeemPoints returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasRedeemPoints() bool {
	if o != nil && !utils.IsNil(o.RedeemPoints) {
		return true
	}

	return false
}

// SetRedeemPoints gets a reference to the given string and assigns it to the RedeemPoints field.
func (o *PublicChannelPropertiesOverwrite) SetRedeemPoints(v string) {
	o.RedeemPoints = &v
}

// GetRequireAuth returns the RequireAuth field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetRequireAuth() bool {
	if o == nil || utils.IsNil(o.RequireAuth) {
		var ret bool
		return ret
	}
	return *o.RequireAuth
}

// GetRequireAuthOk returns a tuple with the RequireAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetRequireAuthOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RequireAuth) {
		return nil, false
	}
	return o.RequireAuth, true
}

// HasRequireAuth returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasRequireAuth() bool {
	if o != nil && !utils.IsNil(o.RequireAuth) {
		return true
	}

	return false
}

// SetRequireAuth gets a reference to the given bool and assigns it to the RequireAuth field.
func (o *PublicChannelPropertiesOverwrite) SetRequireAuth(v bool) {
	o.RequireAuth = &v
}

// GetMerchantIdTag returns the MerchantIdTag field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetMerchantIdTag() string {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		var ret string
		return ret
	}
	return *o.MerchantIdTag
}

// GetMerchantIdTagOk returns a tuple with the MerchantIdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetMerchantIdTagOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		return nil, false
	}
	return o.MerchantIdTag, true
}

// HasMerchantIdTag returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasMerchantIdTag() bool {
	if o != nil && !utils.IsNil(o.MerchantIdTag) {
		return true
	}

	return false
}

// SetMerchantIdTag gets a reference to the given string and assigns it to the MerchantIdTag field.
func (o *PublicChannelPropertiesOverwrite) SetMerchantIdTag(v string) {
	o.MerchantIdTag = &v
}

// GetCardonfileType returns the CardonfileType field value if set, zero value otherwise.
func (o *PublicChannelPropertiesOverwrite) GetCardonfileType() string {
	if o == nil || utils.IsNil(o.CardonfileType) {
		var ret string
		return ret
	}
	return *o.CardonfileType
}

// GetCardonfileTypeOk returns a tuple with the CardonfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicChannelPropertiesOverwrite) GetCardonfileTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardonfileType) {
		return nil, false
	}
	return o.CardonfileType, true
}

// HasCardonfileType returns a boolean if a field has been set.
func (o *PublicChannelPropertiesOverwrite) HasCardonfileType() bool {
	if o != nil && !utils.IsNil(o.CardonfileType) {
		return true
	}

	return false
}

// SetCardonfileType gets a reference to the given string and assigns it to the CardonfileType field.
func (o *PublicChannelPropertiesOverwrite) SetCardonfileType(v string) {
	o.CardonfileType = &v
}

func (o PublicChannelPropertiesOverwrite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicChannelPropertiesOverwrite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SuccessReturnUrl) {
		toSerialize["success_return_url"] = o.SuccessReturnUrl
	}
	if !utils.IsNil(o.FailureReturnUrl) {
		toSerialize["failure_return_url"] = o.FailureReturnUrl
	}
	if !utils.IsNil(o.CancelReturnUrl) {
		toSerialize["cancel_return_url"] = o.CancelReturnUrl
	}
	if !utils.IsNil(o.RedeemPoints) {
		toSerialize["redeem_points"] = o.RedeemPoints
	}
	if !utils.IsNil(o.RequireAuth) {
		toSerialize["require_auth"] = o.RequireAuth
	}
	if !utils.IsNil(o.MerchantIdTag) {
		toSerialize["merchant_id_tag"] = o.MerchantIdTag
	}
	if !utils.IsNil(o.CardonfileType) {
		toSerialize["cardonfile_type"] = o.CardonfileType
	}
	return toSerialize, nil
}

type NullablePublicChannelPropertiesOverwrite struct {
	value *PublicChannelPropertiesOverwrite
	isSet bool
}

func (v NullablePublicChannelPropertiesOverwrite) Get() *PublicChannelPropertiesOverwrite {
	return v.value
}

func (v *NullablePublicChannelPropertiesOverwrite) Set(val *PublicChannelPropertiesOverwrite) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicChannelPropertiesOverwrite) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicChannelPropertiesOverwrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicChannelPropertiesOverwrite(val *PublicChannelPropertiesOverwrite) *NullablePublicChannelPropertiesOverwrite {
	return &NullablePublicChannelPropertiesOverwrite{value: val, isSet: true}
}

func (v NullablePublicChannelPropertiesOverwrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicChannelPropertiesOverwrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

