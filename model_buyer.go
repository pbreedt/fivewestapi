/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"encoding/json"
	"time"
)

// checks if the Buyer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Buyer{}

// Buyer struct for Buyer
type Buyer struct {
	ReferenceBuyerId NullableString `json:"reference_buyer_id,omitempty"`
	BuyerName NullableName `json:"buyer_name,omitempty"`
	BuyerPhoneCountryCode NullableString `json:"buyer_phone_country_code,omitempty"`
	BuyerPhoneNo NullableString `json:"buyer_phone_no,omitempty"`
	BuyerEmail NullableString `json:"buyer_email,omitempty"`
	BuyerRegistrationTime NullableTime `json:"buyer_registration_time,omitempty"`
	BuyerBrowserLanguage NullableString `json:"buyer_browser_language,omitempty"`
}

// NewBuyer instantiates a new Buyer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyer() *Buyer {
	this := Buyer{}
	return &this
}

// NewBuyerWithDefaults instantiates a new Buyer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerWithDefaults() *Buyer {
	this := Buyer{}
	return &this
}

// GetReferenceBuyerId returns the ReferenceBuyerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetReferenceBuyerId() string {
	if o == nil || IsNil(o.ReferenceBuyerId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceBuyerId.Get()
}

// GetReferenceBuyerIdOk returns a tuple with the ReferenceBuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetReferenceBuyerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceBuyerId.Get(), o.ReferenceBuyerId.IsSet()
}

// HasReferenceBuyerId returns a boolean if a field has been set.
func (o *Buyer) HasReferenceBuyerId() bool {
	if o != nil && o.ReferenceBuyerId.IsSet() {
		return true
	}

	return false
}

// SetReferenceBuyerId gets a reference to the given NullableString and assigns it to the ReferenceBuyerId field.
func (o *Buyer) SetReferenceBuyerId(v string) {
	o.ReferenceBuyerId.Set(&v)
}
// SetReferenceBuyerIdNil sets the value for ReferenceBuyerId to be an explicit nil
func (o *Buyer) SetReferenceBuyerIdNil() {
	o.ReferenceBuyerId.Set(nil)
}

// UnsetReferenceBuyerId ensures that no value is present for ReferenceBuyerId, not even an explicit nil
func (o *Buyer) UnsetReferenceBuyerId() {
	o.ReferenceBuyerId.Unset()
}

// GetBuyerName returns the BuyerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerName() Name {
	if o == nil || IsNil(o.BuyerName.Get()) {
		var ret Name
		return ret
	}
	return *o.BuyerName.Get()
}

// GetBuyerNameOk returns a tuple with the BuyerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerName.Get(), o.BuyerName.IsSet()
}

// HasBuyerName returns a boolean if a field has been set.
func (o *Buyer) HasBuyerName() bool {
	if o != nil && o.BuyerName.IsSet() {
		return true
	}

	return false
}

// SetBuyerName gets a reference to the given NullableName and assigns it to the BuyerName field.
func (o *Buyer) SetBuyerName(v Name) {
	o.BuyerName.Set(&v)
}
// SetBuyerNameNil sets the value for BuyerName to be an explicit nil
func (o *Buyer) SetBuyerNameNil() {
	o.BuyerName.Set(nil)
}

// UnsetBuyerName ensures that no value is present for BuyerName, not even an explicit nil
func (o *Buyer) UnsetBuyerName() {
	o.BuyerName.Unset()
}

// GetBuyerPhoneCountryCode returns the BuyerPhoneCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerPhoneCountryCode() string {
	if o == nil || IsNil(o.BuyerPhoneCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerPhoneCountryCode.Get()
}

// GetBuyerPhoneCountryCodeOk returns a tuple with the BuyerPhoneCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerPhoneCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerPhoneCountryCode.Get(), o.BuyerPhoneCountryCode.IsSet()
}

// HasBuyerPhoneCountryCode returns a boolean if a field has been set.
func (o *Buyer) HasBuyerPhoneCountryCode() bool {
	if o != nil && o.BuyerPhoneCountryCode.IsSet() {
		return true
	}

	return false
}

// SetBuyerPhoneCountryCode gets a reference to the given NullableString and assigns it to the BuyerPhoneCountryCode field.
func (o *Buyer) SetBuyerPhoneCountryCode(v string) {
	o.BuyerPhoneCountryCode.Set(&v)
}
// SetBuyerPhoneCountryCodeNil sets the value for BuyerPhoneCountryCode to be an explicit nil
func (o *Buyer) SetBuyerPhoneCountryCodeNil() {
	o.BuyerPhoneCountryCode.Set(nil)
}

// UnsetBuyerPhoneCountryCode ensures that no value is present for BuyerPhoneCountryCode, not even an explicit nil
func (o *Buyer) UnsetBuyerPhoneCountryCode() {
	o.BuyerPhoneCountryCode.Unset()
}

// GetBuyerPhoneNo returns the BuyerPhoneNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerPhoneNo() string {
	if o == nil || IsNil(o.BuyerPhoneNo.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerPhoneNo.Get()
}

// GetBuyerPhoneNoOk returns a tuple with the BuyerPhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerPhoneNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerPhoneNo.Get(), o.BuyerPhoneNo.IsSet()
}

// HasBuyerPhoneNo returns a boolean if a field has been set.
func (o *Buyer) HasBuyerPhoneNo() bool {
	if o != nil && o.BuyerPhoneNo.IsSet() {
		return true
	}

	return false
}

// SetBuyerPhoneNo gets a reference to the given NullableString and assigns it to the BuyerPhoneNo field.
func (o *Buyer) SetBuyerPhoneNo(v string) {
	o.BuyerPhoneNo.Set(&v)
}
// SetBuyerPhoneNoNil sets the value for BuyerPhoneNo to be an explicit nil
func (o *Buyer) SetBuyerPhoneNoNil() {
	o.BuyerPhoneNo.Set(nil)
}

// UnsetBuyerPhoneNo ensures that no value is present for BuyerPhoneNo, not even an explicit nil
func (o *Buyer) UnsetBuyerPhoneNo() {
	o.BuyerPhoneNo.Unset()
}

// GetBuyerEmail returns the BuyerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerEmail() string {
	if o == nil || IsNil(o.BuyerEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerEmail.Get()
}

// GetBuyerEmailOk returns a tuple with the BuyerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerEmail.Get(), o.BuyerEmail.IsSet()
}

// HasBuyerEmail returns a boolean if a field has been set.
func (o *Buyer) HasBuyerEmail() bool {
	if o != nil && o.BuyerEmail.IsSet() {
		return true
	}

	return false
}

// SetBuyerEmail gets a reference to the given NullableString and assigns it to the BuyerEmail field.
func (o *Buyer) SetBuyerEmail(v string) {
	o.BuyerEmail.Set(&v)
}
// SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil
func (o *Buyer) SetBuyerEmailNil() {
	o.BuyerEmail.Set(nil)
}

// UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
func (o *Buyer) UnsetBuyerEmail() {
	o.BuyerEmail.Unset()
}

// GetBuyerRegistrationTime returns the BuyerRegistrationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerRegistrationTime() time.Time {
	if o == nil || IsNil(o.BuyerRegistrationTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BuyerRegistrationTime.Get()
}

// GetBuyerRegistrationTimeOk returns a tuple with the BuyerRegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerRegistrationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerRegistrationTime.Get(), o.BuyerRegistrationTime.IsSet()
}

// HasBuyerRegistrationTime returns a boolean if a field has been set.
func (o *Buyer) HasBuyerRegistrationTime() bool {
	if o != nil && o.BuyerRegistrationTime.IsSet() {
		return true
	}

	return false
}

// SetBuyerRegistrationTime gets a reference to the given NullableTime and assigns it to the BuyerRegistrationTime field.
func (o *Buyer) SetBuyerRegistrationTime(v time.Time) {
	o.BuyerRegistrationTime.Set(&v)
}
// SetBuyerRegistrationTimeNil sets the value for BuyerRegistrationTime to be an explicit nil
func (o *Buyer) SetBuyerRegistrationTimeNil() {
	o.BuyerRegistrationTime.Set(nil)
}

// UnsetBuyerRegistrationTime ensures that no value is present for BuyerRegistrationTime, not even an explicit nil
func (o *Buyer) UnsetBuyerRegistrationTime() {
	o.BuyerRegistrationTime.Unset()
}

// GetBuyerBrowserLanguage returns the BuyerBrowserLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBuyerBrowserLanguage() string {
	if o == nil || IsNil(o.BuyerBrowserLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerBrowserLanguage.Get()
}

// GetBuyerBrowserLanguageOk returns a tuple with the BuyerBrowserLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBuyerBrowserLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerBrowserLanguage.Get(), o.BuyerBrowserLanguage.IsSet()
}

// HasBuyerBrowserLanguage returns a boolean if a field has been set.
func (o *Buyer) HasBuyerBrowserLanguage() bool {
	if o != nil && o.BuyerBrowserLanguage.IsSet() {
		return true
	}

	return false
}

// SetBuyerBrowserLanguage gets a reference to the given NullableString and assigns it to the BuyerBrowserLanguage field.
func (o *Buyer) SetBuyerBrowserLanguage(v string) {
	o.BuyerBrowserLanguage.Set(&v)
}
// SetBuyerBrowserLanguageNil sets the value for BuyerBrowserLanguage to be an explicit nil
func (o *Buyer) SetBuyerBrowserLanguageNil() {
	o.BuyerBrowserLanguage.Set(nil)
}

// UnsetBuyerBrowserLanguage ensures that no value is present for BuyerBrowserLanguage, not even an explicit nil
func (o *Buyer) UnsetBuyerBrowserLanguage() {
	o.BuyerBrowserLanguage.Unset()
}

func (o Buyer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Buyer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceBuyerId.IsSet() {
		toSerialize["reference_buyer_id"] = o.ReferenceBuyerId.Get()
	}
	if o.BuyerName.IsSet() {
		toSerialize["buyer_name"] = o.BuyerName.Get()
	}
	if o.BuyerPhoneCountryCode.IsSet() {
		toSerialize["buyer_phone_country_code"] = o.BuyerPhoneCountryCode.Get()
	}
	if o.BuyerPhoneNo.IsSet() {
		toSerialize["buyer_phone_no"] = o.BuyerPhoneNo.Get()
	}
	if o.BuyerEmail.IsSet() {
		toSerialize["buyer_email"] = o.BuyerEmail.Get()
	}
	if o.BuyerRegistrationTime.IsSet() {
		toSerialize["buyer_registration_time"] = o.BuyerRegistrationTime.Get()
	}
	if o.BuyerBrowserLanguage.IsSet() {
		toSerialize["buyer_browser_language"] = o.BuyerBrowserLanguage.Get()
	}
	return toSerialize, nil
}

type NullableBuyer struct {
	value *Buyer
	isSet bool
}

func (v NullableBuyer) Get() *Buyer {
	return v.value
}

func (v *NullableBuyer) Set(val *Buyer) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyer) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyer(val *Buyer) *NullableBuyer {
	return &NullableBuyer{value: val, isSet: true}
}

func (v NullableBuyer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

