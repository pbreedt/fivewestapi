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

// checks if the GoodsUnitAmountInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoodsUnitAmountInput{}

// GoodsUnitAmountInput struct for GoodsUnitAmountInput
type GoodsUnitAmountInput struct {
	Currency string `json:"currency"`
	Amount Amount `json:"amount"`
}

type _GoodsUnitAmountInput GoodsUnitAmountInput

// NewGoodsUnitAmountInput instantiates a new GoodsUnitAmountInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoodsUnitAmountInput(currency string, amount Amount) *GoodsUnitAmountInput {
	this := GoodsUnitAmountInput{}
	this.Currency = currency
	this.Amount = amount
	return &this
}

// NewGoodsUnitAmountInputWithDefaults instantiates a new GoodsUnitAmountInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoodsUnitAmountInputWithDefaults() *GoodsUnitAmountInput {
	this := GoodsUnitAmountInput{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *GoodsUnitAmountInput) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GoodsUnitAmountInput) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GoodsUnitAmountInput) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *GoodsUnitAmountInput) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GoodsUnitAmountInput) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GoodsUnitAmountInput) SetAmount(v Amount) {
	o.Amount = v
}

func (o GoodsUnitAmountInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoodsUnitAmountInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *GoodsUnitAmountInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"amount",
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

	varGoodsUnitAmountInput := _GoodsUnitAmountInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGoodsUnitAmountInput)

	if err != nil {
		return err
	}

	*o = GoodsUnitAmountInput(varGoodsUnitAmountInput)

	return err
}

type NullableGoodsUnitAmountInput struct {
	value *GoodsUnitAmountInput
	isSet bool
}

func (v NullableGoodsUnitAmountInput) Get() *GoodsUnitAmountInput {
	return v.value
}

func (v *NullableGoodsUnitAmountInput) Set(val *GoodsUnitAmountInput) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsUnitAmountInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsUnitAmountInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsUnitAmountInput(val *GoodsUnitAmountInput) *NullableGoodsUnitAmountInput {
	return &NullableGoodsUnitAmountInput{value: val, isSet: true}
}

func (v NullableGoodsUnitAmountInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsUnitAmountInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

