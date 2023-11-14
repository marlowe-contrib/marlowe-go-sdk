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

// checks if the ObservationOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationOneOf2{}

// ObservationOneOf2 struct for ObservationOneOf2
type ObservationOneOf2 struct {
	Not Observation `json:"not"`
}

// NewObservationOneOf2 instantiates a new ObservationOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationOneOf2(not Observation) *ObservationOneOf2 {
	this := ObservationOneOf2{}
	this.Not = not
	return &this
}

// NewObservationOneOf2WithDefaults instantiates a new ObservationOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationOneOf2WithDefaults() *ObservationOneOf2 {
	this := ObservationOneOf2{}
	return &this
}

// GetNot returns the Not field value
func (o *ObservationOneOf2) GetNot() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.Not
}

// GetNotOk returns a tuple with the Not field value
// and a boolean to check if the value has been set.
func (o *ObservationOneOf2) GetNotOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Not, true
}

// SetNot sets field value
func (o *ObservationOneOf2) SetNot(v Observation) {
	o.Not = v
}

func (o ObservationOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["not"] = o.Not
	return toSerialize, nil
}

type NullableObservationOneOf2 struct {
	value *ObservationOneOf2
	isSet bool
}

func (v NullableObservationOneOf2) Get() *ObservationOneOf2 {
	return v.value
}

func (v *NullableObservationOneOf2) Set(val *ObservationOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationOneOf2(val *ObservationOneOf2) *NullableObservationOneOf2 {
	return &NullableObservationOneOf2{value: val, isSet: true}
}

func (v NullableObservationOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


