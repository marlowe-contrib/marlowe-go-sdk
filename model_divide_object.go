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

// checks if the DivideObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DivideObject{}

// DivideObject struct for DivideObject
type DivideObject struct {
	By ValueObject `json:"by"`
	Divide ValueObject `json:"divide"`
}

// NewDivideObject instantiates a new DivideObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDivideObject(by ValueObject, divide ValueObject) *DivideObject {
	this := DivideObject{}
	this.By = by
	this.Divide = divide
	return &this
}

// NewDivideObjectWithDefaults instantiates a new DivideObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDivideObjectWithDefaults() *DivideObject {
	this := DivideObject{}
	return &this
}

// GetBy returns the By field value
func (o *DivideObject) GetBy() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.By
}

// GetByOk returns a tuple with the By field value
// and a boolean to check if the value has been set.
func (o *DivideObject) GetByOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.By, true
}

// SetBy sets field value
func (o *DivideObject) SetBy(v ValueObject) {
	o.By = v
}

// GetDivide returns the Divide field value
func (o *DivideObject) GetDivide() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Divide
}

// GetDivideOk returns a tuple with the Divide field value
// and a boolean to check if the value has been set.
func (o *DivideObject) GetDivideOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Divide, true
}

// SetDivide sets field value
func (o *DivideObject) SetDivide(v ValueObject) {
	o.Divide = v
}

func (o DivideObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DivideObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["by"] = o.By
	toSerialize["divide"] = o.Divide
	return toSerialize, nil
}

type NullableDivideObject struct {
	value *DivideObject
	isSet bool
}

func (v NullableDivideObject) Get() *DivideObject {
	return v.value
}

func (v *NullableDivideObject) Set(val *DivideObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDivideObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDivideObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDivideObject(val *DivideObject) *NullableDivideObject {
	return &NullableDivideObject{value: val, isSet: true}
}

func (v NullableDivideObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDivideObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


