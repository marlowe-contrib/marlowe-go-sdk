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

// checks if the ValueOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueOneOf3{}

// ValueOneOf3 struct for ValueOneOf3
type ValueOneOf3 struct {
	Minus Value `json:"minus"`
	Value Value `json:"value"`
}

// NewValueOneOf3 instantiates a new ValueOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueOneOf3(minus Value, value Value) *ValueOneOf3 {
	this := ValueOneOf3{}
	this.Minus = minus
	this.Value = value
	return &this
}

// NewValueOneOf3WithDefaults instantiates a new ValueOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueOneOf3WithDefaults() *ValueOneOf3 {
	this := ValueOneOf3{}
	return &this
}

// GetMinus returns the Minus field value
func (o *ValueOneOf3) GetMinus() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Minus
}

// GetMinusOk returns a tuple with the Minus field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf3) GetMinusOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minus, true
}

// SetMinus sets field value
func (o *ValueOneOf3) SetMinus(v Value) {
	o.Minus = v
}

// GetValue returns the Value field value
func (o *ValueOneOf3) GetValue() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf3) GetValueOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValueOneOf3) SetValue(v Value) {
	o.Value = v
}

func (o ValueOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minus"] = o.Minus
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableValueOneOf3 struct {
	value *ValueOneOf3
	isSet bool
}

func (v NullableValueOneOf3) Get() *ValueOneOf3 {
	return v.value
}

func (v *NullableValueOneOf3) Set(val *ValueOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableValueOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableValueOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueOneOf3(val *ValueOneOf3) *NullableValueOneOf3 {
	return &NullableValueOneOf3{value: val, isSet: true}
}

func (v NullableValueOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


