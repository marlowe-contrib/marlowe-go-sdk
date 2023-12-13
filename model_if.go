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

// checks if the If type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &If{}

// If If an observation is true, the first contract applies, otherwise the second contract applies.
type If struct {
	Else Contract `json:"else"`
	If Observation `json:"if"`
	Then Contract `json:"then"`
}

// NewIf instantiates a new If object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIf(else_ Contract, if_ Observation, then Contract) *If {
	this := If{}
	this.Else = else_
	this.If = if_
	this.Then = then
	return &this
}

// NewIfWithDefaults instantiates a new If object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIfWithDefaults() *If {
	this := If{}
	return &this
}

// GetElse returns the Else field value
func (o *If) GetElse() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Else
}

// GetElseOk returns a tuple with the Else field value
// and a boolean to check if the value has been set.
func (o *If) GetElseOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Else, true
}

// SetElse sets field value
func (o *If) SetElse(v Contract) {
	o.Else = v
}

// GetIf returns the If field value
func (o *If) GetIf() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *If) GetIfOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *If) SetIf(v Observation) {
	o.If = v
}

// GetThen returns the Then field value
func (o *If) GetThen() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Then
}

// GetThenOk returns a tuple with the Then field value
// and a boolean to check if the value has been set.
func (o *If) GetThenOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Then, true
}

// SetThen sets field value
func (o *If) SetThen(v Contract) {
	o.Then = v
}

func (o If) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o If) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["else"] = o.Else
	toSerialize["if"] = o.If
	toSerialize["then"] = o.Then
	return toSerialize, nil
}

type NullableIf struct {
	value *If
	isSet bool
}

func (v NullableIf) Get() *If {
	return v.value
}

func (v *NullableIf) Set(val *If) {
	v.value = val
	v.isSet = true
}

func (v NullableIf) IsSet() bool {
	return v.isSet
}

func (v *NullableIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIf(val *If) *NullableIf {
	return &NullableIf{value: val, isSet: true}
}

func (v NullableIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

