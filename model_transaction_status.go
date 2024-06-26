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

// TransactionStatus the model 'TransactionStatus'
type TransactionStatus string

// List of TransactionStatus
const (
	TRANSACTIONSTATUS_COMPLETE TransactionStatus = "complete"
	TRANSACTIONSTATUS_PENDING TransactionStatus = "pending"
	TRANSACTIONSTATUS_UNDERPAID TransactionStatus = "underpaid"
	TRANSACTIONSTATUS_OVERPAID TransactionStatus = "overpaid"
	TRANSACTIONSTATUS_EXPIRED TransactionStatus = "expired"
)

// All allowed values of TransactionStatus enum
var AllowedTransactionStatusEnumValues = []TransactionStatus{
	"complete",
	"pending",
	"underpaid",
	"overpaid",
	"expired",
}

func (v *TransactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatus(value)
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionStatus", value)
}

// NewTransactionStatusFromValue returns a pointer to a valid TransactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusFromValue(v string) (*TransactionStatus, error) {
	ev := TransactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatus: valid values are %v", v, AllowedTransactionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatus) IsValid() bool {
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatus value
func (v TransactionStatus) Ptr() *TransactionStatus {
	return &v
}

type NullableTransactionStatus struct {
	value *TransactionStatus
	isSet bool
}

func (v NullableTransactionStatus) Get() *TransactionStatus {
	return v.value
}

func (v *NullableTransactionStatus) Set(val *TransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatus(val *TransactionStatus) *NullableTransactionStatus {
	return &NullableTransactionStatus{value: val, isSet: true}
}

func (v NullableTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

