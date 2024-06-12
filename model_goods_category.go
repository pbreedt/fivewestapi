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

// GoodsCategory the model 'GoodsCategory'
type GoodsCategory string

// List of GoodsCategory
const (
	GOODSCATEGORY__0000 GoodsCategory = "0000"
	GOODSCATEGORY__1000 GoodsCategory = "1000"
	GOODSCATEGORY__2000 GoodsCategory = "2000"
	GOODSCATEGORY__3000 GoodsCategory = "3000"
	GOODSCATEGORY__4000 GoodsCategory = "4000"
	GOODSCATEGORY__5000 GoodsCategory = "5000"
	GOODSCATEGORY__6000 GoodsCategory = "6000"
	GOODSCATEGORY__7000 GoodsCategory = "7000"
	GOODSCATEGORY__8000 GoodsCategory = "8000"
	GOODSCATEGORY__9000 GoodsCategory = "9000"
	GOODSCATEGORY_A000 GoodsCategory = "A000"
	GOODSCATEGORY_B000 GoodsCategory = "B000"
	GOODSCATEGORY_C000 GoodsCategory = "C000"
	GOODSCATEGORY_D000 GoodsCategory = "D000"
	GOODSCATEGORY_E000 GoodsCategory = "E000"
	GOODSCATEGORY_F000 GoodsCategory = "F000"
	GOODSCATEGORY_Z000 GoodsCategory = "Z000"
)

// All allowed values of GoodsCategory enum
var AllowedGoodsCategoryEnumValues = []GoodsCategory{
	"0000",
	"1000",
	"2000",
	"3000",
	"4000",
	"5000",
	"6000",
	"7000",
	"8000",
	"9000",
	"A000",
	"B000",
	"C000",
	"D000",
	"E000",
	"F000",
	"Z000",
}

func (v *GoodsCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GoodsCategory(value)
	for _, existing := range AllowedGoodsCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GoodsCategory", value)
}

// NewGoodsCategoryFromValue returns a pointer to a valid GoodsCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGoodsCategoryFromValue(v string) (*GoodsCategory, error) {
	ev := GoodsCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GoodsCategory: valid values are %v", v, AllowedGoodsCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GoodsCategory) IsValid() bool {
	for _, existing := range AllowedGoodsCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoodsCategory value
func (v GoodsCategory) Ptr() *GoodsCategory {
	return &v
}

type NullableGoodsCategory struct {
	value *GoodsCategory
	isSet bool
}

func (v NullableGoodsCategory) Get() *GoodsCategory {
	return v.value
}

func (v *NullableGoodsCategory) Set(val *GoodsCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsCategory(val *GoodsCategory) *NullableGoodsCategory {
	return &NullableGoodsCategory{value: val, isSet: true}
}

func (v NullableGoodsCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
