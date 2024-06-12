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

// TradeSide An enumeration.
type TradeSide string

// List of TradeSide
const (
	TRADESIDE_BUY TradeSide = "buy"
	TRADESIDE_SELL TradeSide = "sell"
)

// All allowed values of TradeSide enum
var AllowedTradeSideEnumValues = []TradeSide{
	"buy",
	"sell",
}

func (v *TradeSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TradeSide(value)
	for _, existing := range AllowedTradeSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TradeSide", value)
}

// NewTradeSideFromValue returns a pointer to a valid TradeSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeSideFromValue(v string) (*TradeSide, error) {
	ev := TradeSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TradeSide: valid values are %v", v, AllowedTradeSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradeSide) IsValid() bool {
	for _, existing := range AllowedTradeSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TradeSide value
func (v TradeSide) Ptr() *TradeSide {
	return &v
}

type NullableTradeSide struct {
	value *TradeSide
	isSet bool
}

func (v NullableTradeSide) Get() *TradeSide {
	return v.value
}

func (v *NullableTradeSide) Set(val *TradeSide) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeSide) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeSide(val *TradeSide) *NullableTradeSide {
	return &NullableTradeSide{value: val, isSet: true}
}

func (v NullableTradeSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
