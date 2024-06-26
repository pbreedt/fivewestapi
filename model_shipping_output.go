/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"encoding/json"
)

// checks if the ShippingOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingOutput{}

// ShippingOutput struct for ShippingOutput
type ShippingOutput struct {
	ShippingName NullableName `json:"shipping_name,omitempty"`
	ShippingAddress NullableShippingAddress `json:"shipping_address,omitempty"`
	ShippingPhoneNo NullableString `json:"shipping_phone_no,omitempty"`
}

// NewShippingOutput instantiates a new ShippingOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingOutput() *ShippingOutput {
	this := ShippingOutput{}
	return &this
}

// NewShippingOutputWithDefaults instantiates a new ShippingOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingOutputWithDefaults() *ShippingOutput {
	this := ShippingOutput{}
	return &this
}

// GetShippingName returns the ShippingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingOutput) GetShippingName() Name {
	if o == nil || IsNil(o.ShippingName.Get()) {
		var ret Name
		return ret
	}
	return *o.ShippingName.Get()
}

// GetShippingNameOk returns a tuple with the ShippingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingOutput) GetShippingNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingName.Get(), o.ShippingName.IsSet()
}

// HasShippingName returns a boolean if a field has been set.
func (o *ShippingOutput) HasShippingName() bool {
	if o != nil && o.ShippingName.IsSet() {
		return true
	}

	return false
}

// SetShippingName gets a reference to the given NullableName and assigns it to the ShippingName field.
func (o *ShippingOutput) SetShippingName(v Name) {
	o.ShippingName.Set(&v)
}
// SetShippingNameNil sets the value for ShippingName to be an explicit nil
func (o *ShippingOutput) SetShippingNameNil() {
	o.ShippingName.Set(nil)
}

// UnsetShippingName ensures that no value is present for ShippingName, not even an explicit nil
func (o *ShippingOutput) UnsetShippingName() {
	o.ShippingName.Unset()
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingOutput) GetShippingAddress() ShippingAddress {
	if o == nil || IsNil(o.ShippingAddress.Get()) {
		var ret ShippingAddress
		return ret
	}
	return *o.ShippingAddress.Get()
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingOutput) GetShippingAddressOk() (*ShippingAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingAddress.Get(), o.ShippingAddress.IsSet()
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ShippingOutput) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress.IsSet() {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given NullableShippingAddress and assigns it to the ShippingAddress field.
func (o *ShippingOutput) SetShippingAddress(v ShippingAddress) {
	o.ShippingAddress.Set(&v)
}
// SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil
func (o *ShippingOutput) SetShippingAddressNil() {
	o.ShippingAddress.Set(nil)
}

// UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil
func (o *ShippingOutput) UnsetShippingAddress() {
	o.ShippingAddress.Unset()
}

// GetShippingPhoneNo returns the ShippingPhoneNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingOutput) GetShippingPhoneNo() string {
	if o == nil || IsNil(o.ShippingPhoneNo.Get()) {
		var ret string
		return ret
	}
	return *o.ShippingPhoneNo.Get()
}

// GetShippingPhoneNoOk returns a tuple with the ShippingPhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingOutput) GetShippingPhoneNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingPhoneNo.Get(), o.ShippingPhoneNo.IsSet()
}

// HasShippingPhoneNo returns a boolean if a field has been set.
func (o *ShippingOutput) HasShippingPhoneNo() bool {
	if o != nil && o.ShippingPhoneNo.IsSet() {
		return true
	}

	return false
}

// SetShippingPhoneNo gets a reference to the given NullableString and assigns it to the ShippingPhoneNo field.
func (o *ShippingOutput) SetShippingPhoneNo(v string) {
	o.ShippingPhoneNo.Set(&v)
}
// SetShippingPhoneNoNil sets the value for ShippingPhoneNo to be an explicit nil
func (o *ShippingOutput) SetShippingPhoneNoNil() {
	o.ShippingPhoneNo.Set(nil)
}

// UnsetShippingPhoneNo ensures that no value is present for ShippingPhoneNo, not even an explicit nil
func (o *ShippingOutput) UnsetShippingPhoneNo() {
	o.ShippingPhoneNo.Unset()
}

func (o ShippingOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ShippingName.IsSet() {
		toSerialize["shipping_name"] = o.ShippingName.Get()
	}
	if o.ShippingAddress.IsSet() {
		toSerialize["shipping_address"] = o.ShippingAddress.Get()
	}
	if o.ShippingPhoneNo.IsSet() {
		toSerialize["shipping_phone_no"] = o.ShippingPhoneNo.Get()
	}
	return toSerialize, nil
}

type NullableShippingOutput struct {
	value *ShippingOutput
	isSet bool
}

func (v NullableShippingOutput) Get() *ShippingOutput {
	return v.value
}

func (v *NullableShippingOutput) Set(val *ShippingOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingOutput(val *ShippingOutput) *NullableShippingOutput {
	return &NullableShippingOutput{value: val, isSet: true}
}

func (v NullableShippingOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


