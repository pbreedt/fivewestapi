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

// Mode An enumeration.
type Mode string

// List of Mode
const (
	MODE_OTC Mode = "otc"
	MODE_ARB Mode = "arb"
	MODE_RFQ Mode = "rfq"
	MODE_EXCHANGE Mode = "exchange"
	MODE_PAYMENT Mode = "payment"
	MODE_BOT Mode = "bot"
)

// All allowed values of Mode enum
var AllowedModeEnumValues = []Mode{
	"otc",
	"arb",
	"rfq",
	"exchange",
	"payment",
	"bot",
}

func (v *Mode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Mode(value)
	for _, existing := range AllowedModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Mode", value)
}

// NewModeFromValue returns a pointer to a valid Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModeFromValue(v string) (*Mode, error) {
	ev := Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Mode: valid values are %v", v, AllowedModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Mode) IsValid() bool {
	for _, existing := range AllowedModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Mode value
func (v Mode) Ptr() *Mode {
	return &v
}

type NullableMode struct {
	value *Mode
	isSet bool
}

func (v NullableMode) Get() *Mode {
	return v.value
}

func (v *NullableMode) Set(val *Mode) {
	v.value = val
	v.isSet = true
}

func (v NullableMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMode(val *Mode) *NullableMode {
	return &NullableMode{value: val, isSet: true}
}

func (v NullableMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
