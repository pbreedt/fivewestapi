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

// checks if the RfqRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RfqRequest{}

// RfqRequest struct for RfqRequest
type RfqRequest struct {
	FromAmount *float32 `json:"from_amount,omitempty"`
	ToAmount *float32 `json:"to_amount,omitempty"`
	FromCurrency string `json:"from_currency"`
	ToCurrency string `json:"to_currency"`
}

type _RfqRequest RfqRequest

// NewRfqRequest instantiates a new RfqRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRfqRequest(fromCurrency string, toCurrency string) *RfqRequest {
	this := RfqRequest{}
	this.FromCurrency = fromCurrency
	this.ToCurrency = toCurrency
	return &this
}

// NewRfqRequestWithDefaults instantiates a new RfqRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRfqRequestWithDefaults() *RfqRequest {
	this := RfqRequest{}
	return &this
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *RfqRequest) GetFromAmount() float32 {
	if o == nil || IsNil(o.FromAmount) {
		var ret float32
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfqRequest) GetFromAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *RfqRequest) HasFromAmount() bool {
	if o != nil && !IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given float32 and assigns it to the FromAmount field.
func (o *RfqRequest) SetFromAmount(v float32) {
	o.FromAmount = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *RfqRequest) GetToAmount() float32 {
	if o == nil || IsNil(o.ToAmount) {
		var ret float32
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfqRequest) GetToAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *RfqRequest) HasToAmount() bool {
	if o != nil && !IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given float32 and assigns it to the ToAmount field.
func (o *RfqRequest) SetToAmount(v float32) {
	o.ToAmount = &v
}

// GetFromCurrency returns the FromCurrency field value
func (o *RfqRequest) GetFromCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromCurrency
}

// GetFromCurrencyOk returns a tuple with the FromCurrency field value
// and a boolean to check if the value has been set.
func (o *RfqRequest) GetFromCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromCurrency, true
}

// SetFromCurrency sets field value
func (o *RfqRequest) SetFromCurrency(v string) {
	o.FromCurrency = v
}

// GetToCurrency returns the ToCurrency field value
func (o *RfqRequest) GetToCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToCurrency
}

// GetToCurrencyOk returns a tuple with the ToCurrency field value
// and a boolean to check if the value has been set.
func (o *RfqRequest) GetToCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToCurrency, true
}

// SetToCurrency sets field value
func (o *RfqRequest) SetToCurrency(v string) {
	o.ToCurrency = v
}

func (o RfqRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RfqRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromAmount) {
		toSerialize["from_amount"] = o.FromAmount
	}
	if !IsNil(o.ToAmount) {
		toSerialize["to_amount"] = o.ToAmount
	}
	toSerialize["from_currency"] = o.FromCurrency
	toSerialize["to_currency"] = o.ToCurrency
	return toSerialize, nil
}

func (o *RfqRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_currency",
		"to_currency",
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

	varRfqRequest := _RfqRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRfqRequest)

	if err != nil {
		return err
	}

	*o = RfqRequest(varRfqRequest)

	return err
}

type NullableRfqRequest struct {
	value *RfqRequest
	isSet bool
}

func (v NullableRfqRequest) Get() *RfqRequest {
	return v.value
}

func (v *NullableRfqRequest) Set(val *RfqRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRfqRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRfqRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfqRequest(val *RfqRequest) *NullableRfqRequest {
	return &NullableRfqRequest{value: val, isSet: true}
}

func (v NullableRfqRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfqRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

