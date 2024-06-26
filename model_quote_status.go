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

// QuoteStatus An enumeration.
type QuoteStatus string

// List of QuoteStatus
const (
	QUOTESTATUS_ACCEPTED QuoteStatus = "accepted"
	QUOTESTATUS_UNACCEPTED QuoteStatus = "unaccepted"
)

// All allowed values of QuoteStatus enum
var AllowedQuoteStatusEnumValues = []QuoteStatus{
	"accepted",
	"unaccepted",
}

func (v *QuoteStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuoteStatus(value)
	for _, existing := range AllowedQuoteStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuoteStatus", value)
}

// NewQuoteStatusFromValue returns a pointer to a valid QuoteStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuoteStatusFromValue(v string) (*QuoteStatus, error) {
	ev := QuoteStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuoteStatus: valid values are %v", v, AllowedQuoteStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuoteStatus) IsValid() bool {
	for _, existing := range AllowedQuoteStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuoteStatus value
func (v QuoteStatus) Ptr() *QuoteStatus {
	return &v
}

type NullableQuoteStatus struct {
	value *QuoteStatus
	isSet bool
}

func (v NullableQuoteStatus) Get() *QuoteStatus {
	return v.value
}

func (v *NullableQuoteStatus) Set(val *QuoteStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteStatus(val *QuoteStatus) *NullableQuoteStatus {
	return &NullableQuoteStatus{value: val, isSet: true}
}

func (v NullableQuoteStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

