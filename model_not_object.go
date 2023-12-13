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

// checks if the NotObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotObject{}

// NotObject struct for NotObject
type NotObject struct {
	Not ObservationObject `json:"not"`
}

// NewNotObject instantiates a new NotObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotObject(not ObservationObject) *NotObject {
	this := NotObject{}
	this.Not = not
	return &this
}

// NewNotObjectWithDefaults instantiates a new NotObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotObjectWithDefaults() *NotObject {
	this := NotObject{}
	return &this
}

// GetNot returns the Not field value
func (o *NotObject) GetNot() ObservationObject {
	if o == nil {
		var ret ObservationObject
		return ret
	}

	return o.Not
}

// GetNotOk returns a tuple with the Not field value
// and a boolean to check if the value has been set.
func (o *NotObject) GetNotOk() (*ObservationObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Not, true
}

// SetNot sets field value
func (o *NotObject) SetNot(v ObservationObject) {
	o.Not = v
}

func (o NotObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["not"] = o.Not
	return toSerialize, nil
}

type NullableNotObject struct {
	value *NotObject
	isSet bool
}

func (v NullableNotObject) Get() *NotObject {
	return v.value
}

func (v *NullableNotObject) Set(val *NotObject) {
	v.value = val
	v.isSet = true
}

func (v NullableNotObject) IsSet() bool {
	return v.isSet
}

func (v *NullableNotObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotObject(val *NotObject) *NullableNotObject {
	return &NullableNotObject{value: val, isSet: true}
}

func (v NullableNotObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

