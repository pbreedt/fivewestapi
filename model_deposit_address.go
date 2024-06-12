/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DepositAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositAddress{}

// DepositAddress struct for DepositAddress
type DepositAddress struct {
	Address string `json:"address"`
	Memo NullableString `json:"memo,omitempty"`
}

type _DepositAddress DepositAddress

// NewDepositAddress instantiates a new DepositAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositAddress(address string) *DepositAddress {
	this := DepositAddress{}
	this.Address = address
	return &this
}

// NewDepositAddressWithDefaults instantiates a new DepositAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositAddressWithDefaults() *DepositAddress {
	this := DepositAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *DepositAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DepositAddress) SetAddress(v string) {
	o.Address = v
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositAddress) GetMemo() string {
	if o == nil || IsNil(o.Memo.Get()) {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositAddress) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *DepositAddress) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *DepositAddress) SetMemo(v string) {
	o.Memo.Set(&v)
}
// SetMemoNil sets the value for Memo to be an explicit nil
func (o *DepositAddress) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *DepositAddress) UnsetMemo() {
	o.Memo.Unset()
}

func (o DepositAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if o.Memo.IsSet() {
		toSerialize["memo"] = o.Memo.Get()
	}
	return toSerialize, nil
}

func (o *DepositAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDepositAddress := _DepositAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepositAddress)

	if err != nil {
		return err
	}

	*o = DepositAddress(varDepositAddress)

	return err
}

type NullableDepositAddress struct {
	value *DepositAddress
	isSet bool
}

func (v NullableDepositAddress) Get() *DepositAddress {
	return v.value
}

func (v *NullableDepositAddress) Set(val *DepositAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositAddress(val *DepositAddress) *NullableDepositAddress {
	return &NullableDepositAddress{value: val, isSet: true}
}

func (v NullableDepositAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

