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

// checks if the Payout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payout{}

// Payout struct for Payout
type Payout struct {
	Assets Assets `json:"assets"`
	// A reference to a transaction output with a transaction ID and index.
	PayoutId string `json:"payoutId"`
	Role string `json:"role"`
}

// NewPayout instantiates a new Payout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayout(assets Assets, payoutId string, role string) *Payout {
	this := Payout{}
	this.Assets = assets
	this.PayoutId = payoutId
	this.Role = role
	return &this
}

// NewPayoutWithDefaults instantiates a new Payout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutWithDefaults() *Payout {
	this := Payout{}
	return &this
}

// GetAssets returns the Assets field value
func (o *Payout) GetAssets() Assets {
	if o == nil {
		var ret Assets
		return ret
	}

	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value
// and a boolean to check if the value has been set.
func (o *Payout) GetAssetsOk() (*Assets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assets, true
}

// SetAssets sets field value
func (o *Payout) SetAssets(v Assets) {
	o.Assets = v
}

// GetPayoutId returns the PayoutId field value
func (o *Payout) GetPayoutId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayoutId
}

// GetPayoutIdOk returns a tuple with the PayoutId field value
// and a boolean to check if the value has been set.
func (o *Payout) GetPayoutIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayoutId, true
}

// SetPayoutId sets field value
func (o *Payout) SetPayoutId(v string) {
	o.PayoutId = v
}

// GetRole returns the Role field value
func (o *Payout) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Payout) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Payout) SetRole(v string) {
	o.Role = v
}

func (o Payout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assets"] = o.Assets
	toSerialize["payoutId"] = o.PayoutId
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullablePayout struct {
	value *Payout
	isSet bool
}

func (v NullablePayout) Get() *Payout {
	return v.value
}

func (v *NullablePayout) Set(val *Payout) {
	v.value = val
	v.isSet = true
}

func (v NullablePayout) IsSet() bool {
	return v.isSet
}

func (v *NullablePayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayout(val *Payout) *NullablePayout {
	return &NullablePayout{value: val, isSet: true}
}

func (v NullablePayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


