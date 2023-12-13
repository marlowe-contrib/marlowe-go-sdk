/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marloweruntime

import (
	"encoding/json"
)

// checks if the Or type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Or{}

// Or struct for Or
type Or struct {
	Either Observation `json:"either"`
	Or Observation `json:"or"`
}

// NewOr instantiates a new Or object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOr(either Observation, or Observation) *Or {
	this := Or{}
	this.Either = either
	this.Or = or
	return &this
}

// NewOrWithDefaults instantiates a new Or object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrWithDefaults() *Or {
	this := Or{}
	return &this
}

// GetEither returns the Either field value
func (o *Or) GetEither() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.Either
}

// GetEitherOk returns a tuple with the Either field value
// and a boolean to check if the value has been set.
func (o *Or) GetEitherOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Either, true
}

// SetEither sets field value
func (o *Or) SetEither(v Observation) {
	o.Either = v
}

// GetOr returns the Or field value
func (o *Or) GetOr() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.Or
}

// GetOrOk returns a tuple with the Or field value
// and a boolean to check if the value has been set.
func (o *Or) GetOrOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Or, true
}

// SetOr sets field value
func (o *Or) SetOr(v Observation) {
	o.Or = v
}

func (o Or) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Or) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["either"] = o.Either
	toSerialize["or"] = o.Or
	return toSerialize, nil
}

type NullableOr struct {
	value *Or
	isSet bool
}

func (v NullableOr) Get() *Or {
	return v.value
}

func (v *NullableOr) Set(val *Or) {
	v.value = val
	v.isSet = true
}

func (v NullableOr) IsSet() bool {
	return v.isSet
}

func (v *NullableOr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOr(val *Or) *NullableOr {
	return &NullableOr{value: val, isSet: true}
}

func (v NullableOr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

