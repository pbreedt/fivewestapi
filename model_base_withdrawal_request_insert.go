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

// checks if the BaseWithdrawalRequestInsert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseWithdrawalRequestInsert{}

// BaseWithdrawalRequestInsert struct for BaseWithdrawalRequestInsert
type BaseWithdrawalRequestInsert struct {
	Amount float32 `json:"amount"`
	Destination Destination `json:"destination"`
	CurrencyId int32 `json:"currency_id"`
}

type _BaseWithdrawalRequestInsert BaseWithdrawalRequestInsert

// NewBaseWithdrawalRequestInsert instantiates a new BaseWithdrawalRequestInsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseWithdrawalRequestInsert(amount float32, destination Destination, currencyId int32) *BaseWithdrawalRequestInsert {
	this := BaseWithdrawalRequestInsert{}
	this.Amount = amount
	this.Destination = destination
	this.CurrencyId = currencyId
	return &this
}

// NewBaseWithdrawalRequestInsertWithDefaults instantiates a new BaseWithdrawalRequestInsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseWithdrawalRequestInsertWithDefaults() *BaseWithdrawalRequestInsert {
	this := BaseWithdrawalRequestInsert{}
	return &this
}

// GetAmount returns the Amount field value
func (o *BaseWithdrawalRequestInsert) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BaseWithdrawalRequestInsert) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BaseWithdrawalRequestInsert) SetAmount(v float32) {
	o.Amount = v
}

// GetDestination returns the Destination field value
func (o *BaseWithdrawalRequestInsert) GetDestination() Destination {
	if o == nil {
		var ret Destination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *BaseWithdrawalRequestInsert) GetDestinationOk() (*Destination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *BaseWithdrawalRequestInsert) SetDestination(v Destination) {
	o.Destination = v
}

// GetCurrencyId returns the CurrencyId field value
func (o *BaseWithdrawalRequestInsert) GetCurrencyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value
// and a boolean to check if the value has been set.
func (o *BaseWithdrawalRequestInsert) GetCurrencyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyId, true
}

// SetCurrencyId sets field value
func (o *BaseWithdrawalRequestInsert) SetCurrencyId(v int32) {
	o.CurrencyId = v
}

func (o BaseWithdrawalRequestInsert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseWithdrawalRequestInsert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["destination"] = o.Destination
	toSerialize["currency_id"] = o.CurrencyId
	return toSerialize, nil
}

func (o *BaseWithdrawalRequestInsert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"destination",
		"currency_id",
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

	varBaseWithdrawalRequestInsert := _BaseWithdrawalRequestInsert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseWithdrawalRequestInsert)

	if err != nil {
		return err
	}

	*o = BaseWithdrawalRequestInsert(varBaseWithdrawalRequestInsert)

	return err
}

type NullableBaseWithdrawalRequestInsert struct {
	value *BaseWithdrawalRequestInsert
	isSet bool
}

func (v NullableBaseWithdrawalRequestInsert) Get() *BaseWithdrawalRequestInsert {
	return v.value
}

func (v *NullableBaseWithdrawalRequestInsert) Set(val *BaseWithdrawalRequestInsert) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseWithdrawalRequestInsert) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseWithdrawalRequestInsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseWithdrawalRequestInsert(val *BaseWithdrawalRequestInsert) *NullableBaseWithdrawalRequestInsert {
	return &NullableBaseWithdrawalRequestInsert{value: val, isSet: true}
}

func (v NullableBaseWithdrawalRequestInsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseWithdrawalRequestInsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

