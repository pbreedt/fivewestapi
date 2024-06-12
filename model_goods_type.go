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

// GoodsType the model 'GoodsType'
type GoodsType string

// List of GoodsType
const (
	GOODSTYPE__01 GoodsType = "01"
	GOODSTYPE__02 GoodsType = "02"
)

// All allowed values of GoodsType enum
var AllowedGoodsTypeEnumValues = []GoodsType{
	"01",
	"02",
}

func (v *GoodsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GoodsType(value)
	for _, existing := range AllowedGoodsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GoodsType", value)
}

// NewGoodsTypeFromValue returns a pointer to a valid GoodsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGoodsTypeFromValue(v string) (*GoodsType, error) {
	ev := GoodsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GoodsType: valid values are %v", v, AllowedGoodsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GoodsType) IsValid() bool {
	for _, existing := range AllowedGoodsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoodsType value
func (v GoodsType) Ptr() *GoodsType {
	return &v
}

type NullableGoodsType struct {
	value *GoodsType
	isSet bool
}

func (v NullableGoodsType) Get() *GoodsType {
	return v.value
}

func (v *NullableGoodsType) Set(val *GoodsType) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsType) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsType(val *GoodsType) *NullableGoodsType {
	return &NullableGoodsType{value: val, isSet: true}
}

func (v NullableGoodsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
