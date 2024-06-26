/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"encoding/json"
	"fmt"
)

// Id struct for Id
type Id struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Id) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Id)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Id) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableId struct {
	value *Id
	isSet bool
}

func (v NullableId) Get() *Id {
	return v.value
}

func (v *NullableId) Set(val *Id) {
	v.value = val
	v.isSet = true
}

func (v NullableId) IsSet() bool {
	return v.isSet
}

func (v *NullableId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableId(val *Id) *NullableId {
	return &NullableId{value: val, isSet: true}
}

func (v NullableId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


