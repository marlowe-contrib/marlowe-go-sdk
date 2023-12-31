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

// checks if the IfValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IfValue{}

// IfValue struct for IfValue
type IfValue struct {
	Else Value `json:"else"`
	If Observation `json:"if"`
	Then Value `json:"then"`
}

// NewIfValue instantiates a new IfValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIfValue(else_ Value, if_ Observation, then Value) *IfValue {
	this := IfValue{}
	this.Else = else_
	this.If = if_
	this.Then = then
	return &this
}

// NewIfValueWithDefaults instantiates a new IfValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIfValueWithDefaults() *IfValue {
	this := IfValue{}
	return &this
}

// GetElse returns the Else field value
func (o *IfValue) GetElse() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Else
}

// GetElseOk returns a tuple with the Else field value
// and a boolean to check if the value has been set.
func (o *IfValue) GetElseOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Else, true
}

// SetElse sets field value
func (o *IfValue) SetElse(v Value) {
	o.Else = v
}

// GetIf returns the If field value
func (o *IfValue) GetIf() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *IfValue) GetIfOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *IfValue) SetIf(v Observation) {
	o.If = v
}

// GetThen returns the Then field value
func (o *IfValue) GetThen() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Then
}

// GetThenOk returns a tuple with the Then field value
// and a boolean to check if the value has been set.
func (o *IfValue) GetThenOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Then, true
}

// SetThen sets field value
func (o *IfValue) SetThen(v Value) {
	o.Then = v
}

func (o IfValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IfValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["else"] = o.Else
	toSerialize["if"] = o.If
	toSerialize["then"] = o.Then
	return toSerialize, nil
}

type NullableIfValue struct {
	value *IfValue
	isSet bool
}

func (v NullableIfValue) Get() *IfValue {
	return v.value
}

func (v *NullableIfValue) Set(val *IfValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIfValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIfValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIfValue(val *IfValue) *NullableIfValue {
	return &NullableIfValue{value: val, isSet: true}
}

func (v NullableIfValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIfValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


