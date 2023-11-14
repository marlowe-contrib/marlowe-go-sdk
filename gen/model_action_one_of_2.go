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

// checks if the ActionOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionOneOf2{}

// ActionOneOf2 struct for ActionOneOf2
type ActionOneOf2 struct {
	NotifyIf Observation `json:"notify_if"`
}

// NewActionOneOf2 instantiates a new ActionOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionOneOf2(notifyIf Observation) *ActionOneOf2 {
	this := ActionOneOf2{}
	this.NotifyIf = notifyIf
	return &this
}

// NewActionOneOf2WithDefaults instantiates a new ActionOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionOneOf2WithDefaults() *ActionOneOf2 {
	this := ActionOneOf2{}
	return &this
}

// GetNotifyIf returns the NotifyIf field value
func (o *ActionOneOf2) GetNotifyIf() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.NotifyIf
}

// GetNotifyIfOk returns a tuple with the NotifyIf field value
// and a boolean to check if the value has been set.
func (o *ActionOneOf2) GetNotifyIfOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyIf, true
}

// SetNotifyIf sets field value
func (o *ActionOneOf2) SetNotifyIf(v Observation) {
	o.NotifyIf = v
}

func (o ActionOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notify_if"] = o.NotifyIf
	return toSerialize, nil
}

type NullableActionOneOf2 struct {
	value *ActionOneOf2
	isSet bool
}

func (v NullableActionOneOf2) Get() *ActionOneOf2 {
	return v.value
}

func (v *NullableActionOneOf2) Set(val *ActionOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableActionOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableActionOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionOneOf2(val *ActionOneOf2) *NullableActionOneOf2 {
	return &NullableActionOneOf2{value: val, isSet: true}
}

func (v NullableActionOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


