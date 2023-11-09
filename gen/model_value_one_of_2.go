/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueOneOf2{}

// ValueOneOf2 struct for ValueOneOf2
type ValueOneOf2 struct {
	Add Value `json:"add"`
	And Value `json:"and"`
}

// NewValueOneOf2 instantiates a new ValueOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueOneOf2(add Value, and Value) *ValueOneOf2 {
	this := ValueOneOf2{}
	this.Add = add
	this.And = and
	return &this
}

// NewValueOneOf2WithDefaults instantiates a new ValueOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueOneOf2WithDefaults() *ValueOneOf2 {
	this := ValueOneOf2{}
	return &this
}

// GetAdd returns the Add field value
func (o *ValueOneOf2) GetAdd() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Add
}

// GetAddOk returns a tuple with the Add field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf2) GetAddOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Add, true
}

// SetAdd sets field value
func (o *ValueOneOf2) SetAdd(v Value) {
	o.Add = v
}

// GetAnd returns the And field value
func (o *ValueOneOf2) GetAnd() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.And
}

// GetAndOk returns a tuple with the And field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf2) GetAndOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.And, true
}

// SetAnd sets field value
func (o *ValueOneOf2) SetAnd(v Value) {
	o.And = v
}

func (o ValueOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add"] = o.Add
	toSerialize["and"] = o.And
	return toSerialize, nil
}

type NullableValueOneOf2 struct {
	value *ValueOneOf2
	isSet bool
}

func (v NullableValueOneOf2) Get() *ValueOneOf2 {
	return v.value
}

func (v *NullableValueOneOf2) Set(val *ValueOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableValueOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableValueOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueOneOf2(val *ValueOneOf2) *NullableValueOneOf2 {
	return &NullableValueOneOf2{value: val, isSet: true}
}

func (v NullableValueOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

