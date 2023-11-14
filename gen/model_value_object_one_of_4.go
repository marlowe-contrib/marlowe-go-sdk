/*
Marlowe Runtime REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueObjectOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueObjectOneOf4{}

// ValueObjectOneOf4 struct for ValueObjectOneOf4
type ValueObjectOneOf4 struct {
	Multiply ValueObject `json:"multiply"`
	Times ValueObject `json:"times"`
}

// NewValueObjectOneOf4 instantiates a new ValueObjectOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueObjectOneOf4(multiply ValueObject, times ValueObject) *ValueObjectOneOf4 {
	this := ValueObjectOneOf4{}
	this.Multiply = multiply
	this.Times = times
	return &this
}

// NewValueObjectOneOf4WithDefaults instantiates a new ValueObjectOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueObjectOneOf4WithDefaults() *ValueObjectOneOf4 {
	this := ValueObjectOneOf4{}
	return &this
}

// GetMultiply returns the Multiply field value
func (o *ValueObjectOneOf4) GetMultiply() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Multiply
}

// GetMultiplyOk returns a tuple with the Multiply field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf4) GetMultiplyOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiply, true
}

// SetMultiply sets field value
func (o *ValueObjectOneOf4) SetMultiply(v ValueObject) {
	o.Multiply = v
}

// GetTimes returns the Times field value
func (o *ValueObjectOneOf4) GetTimes() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf4) GetTimesOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Times, true
}

// SetTimes sets field value
func (o *ValueObjectOneOf4) SetTimes(v ValueObject) {
	o.Times = v
}

func (o ValueObjectOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueObjectOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["multiply"] = o.Multiply
	toSerialize["times"] = o.Times
	return toSerialize, nil
}

type NullableValueObjectOneOf4 struct {
	value *ValueObjectOneOf4
	isSet bool
}

func (v NullableValueObjectOneOf4) Get() *ValueObjectOneOf4 {
	return v.value
}

func (v *NullableValueObjectOneOf4) Set(val *ValueObjectOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableValueObjectOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableValueObjectOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueObjectOneOf4(val *ValueObjectOneOf4) *NullableValueObjectOneOf4 {
	return &NullableValueObjectOneOf4{value: val, isSet: true}
}

func (v NullableValueObjectOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueObjectOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


