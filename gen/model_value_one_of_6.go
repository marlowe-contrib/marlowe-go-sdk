/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueOneOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueOneOf6{}

// ValueOneOf6 struct for ValueOneOf6
type ValueOneOf6 struct {
	ValueOfChoice ChoiceId `json:"value_of_choice"`
}

// NewValueOneOf6 instantiates a new ValueOneOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueOneOf6(valueOfChoice ChoiceId) *ValueOneOf6 {
	this := ValueOneOf6{}
	this.ValueOfChoice = valueOfChoice
	return &this
}

// NewValueOneOf6WithDefaults instantiates a new ValueOneOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueOneOf6WithDefaults() *ValueOneOf6 {
	this := ValueOneOf6{}
	return &this
}

// GetValueOfChoice returns the ValueOfChoice field value
func (o *ValueOneOf6) GetValueOfChoice() ChoiceId {
	if o == nil {
		var ret ChoiceId
		return ret
	}

	return o.ValueOfChoice
}

// GetValueOfChoiceOk returns a tuple with the ValueOfChoice field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf6) GetValueOfChoiceOk() (*ChoiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueOfChoice, true
}

// SetValueOfChoice sets field value
func (o *ValueOneOf6) SetValueOfChoice(v ChoiceId) {
	o.ValueOfChoice = v
}

func (o ValueOneOf6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueOneOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value_of_choice"] = o.ValueOfChoice
	return toSerialize, nil
}

type NullableValueOneOf6 struct {
	value *ValueOneOf6
	isSet bool
}

func (v NullableValueOneOf6) Get() *ValueOneOf6 {
	return v.value
}

func (v *NullableValueOneOf6) Set(val *ValueOneOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableValueOneOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableValueOneOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueOneOf6(val *ValueOneOf6) *NullableValueOneOf6 {
	return &NullableValueOneOf6{value: val, isSet: true}
}

func (v NullableValueOneOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueOneOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


