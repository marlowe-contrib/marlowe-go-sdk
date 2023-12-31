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

// checks if the GreaterOrEqual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GreaterOrEqual{}

// GreaterOrEqual struct for GreaterOrEqual
type GreaterOrEqual struct {
	GeThan Value `json:"ge_than"`
	Value Value `json:"value"`
}

// NewGreaterOrEqual instantiates a new GreaterOrEqual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterOrEqual(geThan Value, value Value) *GreaterOrEqual {
	this := GreaterOrEqual{}
	this.GeThan = geThan
	this.Value = value
	return &this
}

// NewGreaterOrEqualWithDefaults instantiates a new GreaterOrEqual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterOrEqualWithDefaults() *GreaterOrEqual {
	this := GreaterOrEqual{}
	return &this
}

// GetGeThan returns the GeThan field value
func (o *GreaterOrEqual) GetGeThan() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.GeThan
}

// GetGeThanOk returns a tuple with the GeThan field value
// and a boolean to check if the value has been set.
func (o *GreaterOrEqual) GetGeThanOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeThan, true
}

// SetGeThan sets field value
func (o *GreaterOrEqual) SetGeThan(v Value) {
	o.GeThan = v
}

// GetValue returns the Value field value
func (o *GreaterOrEqual) GetValue() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GreaterOrEqual) GetValueOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GreaterOrEqual) SetValue(v Value) {
	o.Value = v
}

func (o GreaterOrEqual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreaterOrEqual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ge_than"] = o.GeThan
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableGreaterOrEqual struct {
	value *GreaterOrEqual
	isSet bool
}

func (v NullableGreaterOrEqual) Get() *GreaterOrEqual {
	return v.value
}

func (v *NullableGreaterOrEqual) Set(val *GreaterOrEqual) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterOrEqual) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterOrEqual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterOrEqual(val *GreaterOrEqual) *NullableGreaterOrEqual {
	return &NullableGreaterOrEqual{value: val, isSet: true}
}

func (v NullableGreaterOrEqual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterOrEqual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


