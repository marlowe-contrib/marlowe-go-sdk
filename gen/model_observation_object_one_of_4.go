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

// checks if the ObservationObjectOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationObjectOneOf4{}

// ObservationObjectOneOf4 struct for ObservationObjectOneOf4
type ObservationObjectOneOf4 struct {
	GeThan ValueObject `json:"ge_than"`
	Value ValueObject `json:"value"`
}

// NewObservationObjectOneOf4 instantiates a new ObservationObjectOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationObjectOneOf4(geThan ValueObject, value ValueObject) *ObservationObjectOneOf4 {
	this := ObservationObjectOneOf4{}
	this.GeThan = geThan
	this.Value = value
	return &this
}

// NewObservationObjectOneOf4WithDefaults instantiates a new ObservationObjectOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationObjectOneOf4WithDefaults() *ObservationObjectOneOf4 {
	this := ObservationObjectOneOf4{}
	return &this
}

// GetGeThan returns the GeThan field value
func (o *ObservationObjectOneOf4) GetGeThan() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.GeThan
}

// GetGeThanOk returns a tuple with the GeThan field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf4) GetGeThanOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeThan, true
}

// SetGeThan sets field value
func (o *ObservationObjectOneOf4) SetGeThan(v ValueObject) {
	o.GeThan = v
}

// GetValue returns the Value field value
func (o *ObservationObjectOneOf4) GetValue() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf4) GetValueOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ObservationObjectOneOf4) SetValue(v ValueObject) {
	o.Value = v
}

func (o ObservationObjectOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationObjectOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ge_than"] = o.GeThan
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableObservationObjectOneOf4 struct {
	value *ObservationObjectOneOf4
	isSet bool
}

func (v NullableObservationObjectOneOf4) Get() *ObservationObjectOneOf4 {
	return v.value
}

func (v *NullableObservationObjectOneOf4) Set(val *ObservationObjectOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObjectOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObjectOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObjectOneOf4(val *ObservationObjectOneOf4) *NullableObservationObjectOneOf4 {
	return &NullableObservationObjectOneOf4{value: val, isSet: true}
}

func (v NullableObservationObjectOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObjectOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


