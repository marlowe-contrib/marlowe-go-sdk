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

// checks if the Not type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Not{}

// Not struct for Not
type Not struct {
	Not Observation `json:"not"`
}

// NewNot instantiates a new Not object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNot(not Observation) *Not {
	this := Not{}
	this.Not = not
	return &this
}

// NewNotWithDefaults instantiates a new Not object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotWithDefaults() *Not {
	this := Not{}
	return &this
}

// GetNot returns the Not field value
func (o *Not) GetNot() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.Not
}

// GetNotOk returns a tuple with the Not field value
// and a boolean to check if the value has been set.
func (o *Not) GetNotOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Not, true
}

// SetNot sets field value
func (o *Not) SetNot(v Observation) {
	o.Not = v
}

func (o Not) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Not) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["not"] = o.Not
	return toSerialize, nil
}

type NullableNot struct {
	value *Not
	isSet bool
}

func (v NullableNot) Get() *Not {
	return v.value
}

func (v *NullableNot) Set(val *Not) {
	v.value = val
	v.isSet = true
}

func (v NullableNot) IsSet() bool {
	return v.isSet
}

func (v *NullableNot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNot(val *Not) *NullableNot {
	return &NullableNot{value: val, isSet: true}
}

func (v NullableNot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


