/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueObjectOneOf7 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueObjectOneOf7{}

// ValueObjectOneOf7 struct for ValueObjectOneOf7
type ValueObjectOneOf7 struct {
	Else ValueObject `json:"else"`
	If ObservationObject `json:"if"`
	Then ValueObject `json:"then"`
}

// NewValueObjectOneOf7 instantiates a new ValueObjectOneOf7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueObjectOneOf7(else_ ValueObject, if_ ObservationObject, then ValueObject) *ValueObjectOneOf7 {
	this := ValueObjectOneOf7{}
	this.Else = else_
	this.If = if_
	this.Then = then
	return &this
}

// NewValueObjectOneOf7WithDefaults instantiates a new ValueObjectOneOf7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueObjectOneOf7WithDefaults() *ValueObjectOneOf7 {
	this := ValueObjectOneOf7{}
	return &this
}

// GetElse returns the Else field value
func (o *ValueObjectOneOf7) GetElse() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Else
}

// GetElseOk returns a tuple with the Else field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf7) GetElseOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Else, true
}

// SetElse sets field value
func (o *ValueObjectOneOf7) SetElse(v ValueObject) {
	o.Else = v
}

// GetIf returns the If field value
func (o *ValueObjectOneOf7) GetIf() ObservationObject {
	if o == nil {
		var ret ObservationObject
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf7) GetIfOk() (*ObservationObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *ValueObjectOneOf7) SetIf(v ObservationObject) {
	o.If = v
}

// GetThen returns the Then field value
func (o *ValueObjectOneOf7) GetThen() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Then
}

// GetThenOk returns a tuple with the Then field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf7) GetThenOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Then, true
}

// SetThen sets field value
func (o *ValueObjectOneOf7) SetThen(v ValueObject) {
	o.Then = v
}

func (o ValueObjectOneOf7) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueObjectOneOf7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["else"] = o.Else
	toSerialize["if"] = o.If
	toSerialize["then"] = o.Then
	return toSerialize, nil
}

type NullableValueObjectOneOf7 struct {
	value *ValueObjectOneOf7
	isSet bool
}

func (v NullableValueObjectOneOf7) Get() *ValueObjectOneOf7 {
	return v.value
}

func (v *NullableValueObjectOneOf7) Set(val *ValueObjectOneOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableValueObjectOneOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableValueObjectOneOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueObjectOneOf7(val *ValueObjectOneOf7) *NullableValueObjectOneOf7 {
	return &NullableValueObjectOneOf7{value: val, isSet: true}
}

func (v NullableValueObjectOneOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueObjectOneOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


