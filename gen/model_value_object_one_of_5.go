/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueObjectOneOf5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueObjectOneOf5{}

// ValueObjectOneOf5 struct for ValueObjectOneOf5
type ValueObjectOneOf5 struct {
	By ValueObject `json:"by"`
	Divide ValueObject `json:"divide"`
}

// NewValueObjectOneOf5 instantiates a new ValueObjectOneOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueObjectOneOf5(by ValueObject, divide ValueObject) *ValueObjectOneOf5 {
	this := ValueObjectOneOf5{}
	this.By = by
	this.Divide = divide
	return &this
}

// NewValueObjectOneOf5WithDefaults instantiates a new ValueObjectOneOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueObjectOneOf5WithDefaults() *ValueObjectOneOf5 {
	this := ValueObjectOneOf5{}
	return &this
}

// GetBy returns the By field value
func (o *ValueObjectOneOf5) GetBy() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.By
}

// GetByOk returns a tuple with the By field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf5) GetByOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.By, true
}

// SetBy sets field value
func (o *ValueObjectOneOf5) SetBy(v ValueObject) {
	o.By = v
}

// GetDivide returns the Divide field value
func (o *ValueObjectOneOf5) GetDivide() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Divide
}

// GetDivideOk returns a tuple with the Divide field value
// and a boolean to check if the value has been set.
func (o *ValueObjectOneOf5) GetDivideOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Divide, true
}

// SetDivide sets field value
func (o *ValueObjectOneOf5) SetDivide(v ValueObject) {
	o.Divide = v
}

func (o ValueObjectOneOf5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueObjectOneOf5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["by"] = o.By
	toSerialize["divide"] = o.Divide
	return toSerialize, nil
}

type NullableValueObjectOneOf5 struct {
	value *ValueObjectOneOf5
	isSet bool
}

func (v NullableValueObjectOneOf5) Get() *ValueObjectOneOf5 {
	return v.value
}

func (v *NullableValueObjectOneOf5) Set(val *ValueObjectOneOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableValueObjectOneOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableValueObjectOneOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueObjectOneOf5(val *ValueObjectOneOf5) *NullableValueObjectOneOf5 {
	return &NullableValueObjectOneOf5{value: val, isSet: true}
}

func (v NullableValueObjectOneOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueObjectOneOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


