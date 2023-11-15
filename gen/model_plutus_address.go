/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PlutusAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlutusAddress{}

// PlutusAddress A Plutus address.
type PlutusAddress struct {
	AddressCredential PlutusCredential `json:"addressCredential"`
	AddressStakingCredential *PlutusStakingCredential `json:"addressStakingCredential,omitempty"`
}

// NewPlutusAddress instantiates a new PlutusAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlutusAddress(addressCredential PlutusCredential) *PlutusAddress {
	this := PlutusAddress{}
	this.AddressCredential = addressCredential
	return &this
}

// NewPlutusAddressWithDefaults instantiates a new PlutusAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlutusAddressWithDefaults() *PlutusAddress {
	this := PlutusAddress{}
	return &this
}

// GetAddressCredential returns the AddressCredential field value
func (o *PlutusAddress) GetAddressCredential() PlutusCredential {
	if o == nil {
		var ret PlutusCredential
		return ret
	}

	return o.AddressCredential
}

// GetAddressCredentialOk returns a tuple with the AddressCredential field value
// and a boolean to check if the value has been set.
func (o *PlutusAddress) GetAddressCredentialOk() (*PlutusCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressCredential, true
}

// SetAddressCredential sets field value
func (o *PlutusAddress) SetAddressCredential(v PlutusCredential) {
	o.AddressCredential = v
}

// GetAddressStakingCredential returns the AddressStakingCredential field value if set, zero value otherwise.
func (o *PlutusAddress) GetAddressStakingCredential() PlutusStakingCredential {
	if o == nil || IsNil(o.AddressStakingCredential) {
		var ret PlutusStakingCredential
		return ret
	}
	return *o.AddressStakingCredential
}

// GetAddressStakingCredentialOk returns a tuple with the AddressStakingCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlutusAddress) GetAddressStakingCredentialOk() (*PlutusStakingCredential, bool) {
	if o == nil || IsNil(o.AddressStakingCredential) {
		return nil, false
	}
	return o.AddressStakingCredential, true
}

// HasAddressStakingCredential returns a boolean if a field has been set.
func (o *PlutusAddress) HasAddressStakingCredential() bool {
	if o != nil && !IsNil(o.AddressStakingCredential) {
		return true
	}

	return false
}

// SetAddressStakingCredential gets a reference to the given PlutusStakingCredential and assigns it to the AddressStakingCredential field.
func (o *PlutusAddress) SetAddressStakingCredential(v PlutusStakingCredential) {
	o.AddressStakingCredential = &v
}

func (o PlutusAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlutusAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressCredential"] = o.AddressCredential
	if !IsNil(o.AddressStakingCredential) {
		toSerialize["addressStakingCredential"] = o.AddressStakingCredential
	}
	return toSerialize, nil
}

type NullablePlutusAddress struct {
	value *PlutusAddress
	isSet bool
}

func (v NullablePlutusAddress) Get() *PlutusAddress {
	return v.value
}

func (v *NullablePlutusAddress) Set(val *PlutusAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePlutusAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePlutusAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlutusAddress(val *PlutusAddress) *NullablePlutusAddress {
	return &NullablePlutusAddress{value: val, isSet: true}
}

func (v NullablePlutusAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlutusAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


