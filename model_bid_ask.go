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

// checks if the BidAsk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidAsk{}

// BidAsk struct for BidAsk
type BidAsk struct {
	BaseCurrency string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
	Bid float32 `json:"bid"`
	Ask float32 `json:"ask"`
}

type _BidAsk BidAsk

// NewBidAsk instantiates a new BidAsk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidAsk(baseCurrency string, quoteCurrency string, bid float32, ask float32) *BidAsk {
	this := BidAsk{}
	this.BaseCurrency = baseCurrency
	this.QuoteCurrency = quoteCurrency
	this.Bid = bid
	this.Ask = ask
	return &this
}

// NewBidAskWithDefaults instantiates a new BidAsk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidAskWithDefaults() *BidAsk {
	this := BidAsk{}
	return &this
}

// GetBaseCurrency returns the BaseCurrency field value
func (o *BidAsk) GetBaseCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value
// and a boolean to check if the value has been set.
func (o *BidAsk) GetBaseCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrency, true
}

// SetBaseCurrency sets field value
func (o *BidAsk) SetBaseCurrency(v string) {
	o.BaseCurrency = v
}

// GetQuoteCurrency returns the QuoteCurrency field value
func (o *BidAsk) GetQuoteCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteCurrency
}

// GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field value
// and a boolean to check if the value has been set.
func (o *BidAsk) GetQuoteCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteCurrency, true
}

// SetQuoteCurrency sets field value
func (o *BidAsk) SetQuoteCurrency(v string) {
	o.QuoteCurrency = v
}

// GetBid returns the Bid field value
func (o *BidAsk) GetBid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *BidAsk) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *BidAsk) SetBid(v float32) {
	o.Bid = v
}

// GetAsk returns the Ask field value
func (o *BidAsk) GetAsk() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ask
}

// GetAskOk returns a tuple with the Ask field value
// and a boolean to check if the value has been set.
func (o *BidAsk) GetAskOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ask, true
}

// SetAsk sets field value
func (o *BidAsk) SetAsk(v float32) {
	o.Ask = v
}

func (o BidAsk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BidAsk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_currency"] = o.BaseCurrency
	toSerialize["quote_currency"] = o.QuoteCurrency
	toSerialize["bid"] = o.Bid
	toSerialize["ask"] = o.Ask
	return toSerialize, nil
}

func (o *BidAsk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_currency",
		"quote_currency",
		"bid",
		"ask",
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

	varBidAsk := _BidAsk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBidAsk)

	if err != nil {
		return err
	}

	*o = BidAsk(varBidAsk)

	return err
}

type NullableBidAsk struct {
	value *BidAsk
	isSet bool
}

func (v NullableBidAsk) Get() *BidAsk {
	return v.value
}

func (v *NullableBidAsk) Set(val *BidAsk) {
	v.value = val
	v.isSet = true
}

func (v NullableBidAsk) IsSet() bool {
	return v.isSet
}

func (v *NullableBidAsk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidAsk(val *BidAsk) *NullableBidAsk {
	return &NullableBidAsk{value: val, isSet: true}
}

func (v NullableBidAsk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBidAsk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

