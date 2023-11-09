/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InputOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputOneOf2{}

// InputOneOf2 Make a choice in a contract
type InputOneOf2 struct {
	ForChoiceId ChoiceId `json:"for_choice_id"`
	InputThatChoosesNum int32 `json:"input_that_chooses_num"`
}

// NewInputOneOf2 instantiates a new InputOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOneOf2(forChoiceId ChoiceId, inputThatChoosesNum int32) *InputOneOf2 {
	this := InputOneOf2{}
	this.ForChoiceId = forChoiceId
	this.InputThatChoosesNum = inputThatChoosesNum
	return &this
}

// NewInputOneOf2WithDefaults instantiates a new InputOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOneOf2WithDefaults() *InputOneOf2 {
	this := InputOneOf2{}
	return &this
}

// GetForChoiceId returns the ForChoiceId field value
func (o *InputOneOf2) GetForChoiceId() ChoiceId {
	if o == nil {
		var ret ChoiceId
		return ret
	}

	return o.ForChoiceId
}

// GetForChoiceIdOk returns a tuple with the ForChoiceId field value
// and a boolean to check if the value has been set.
func (o *InputOneOf2) GetForChoiceIdOk() (*ChoiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForChoiceId, true
}

// SetForChoiceId sets field value
func (o *InputOneOf2) SetForChoiceId(v ChoiceId) {
	o.ForChoiceId = v
}

// GetInputThatChoosesNum returns the InputThatChoosesNum field value
func (o *InputOneOf2) GetInputThatChoosesNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputThatChoosesNum
}

// GetInputThatChoosesNumOk returns a tuple with the InputThatChoosesNum field value
// and a boolean to check if the value has been set.
func (o *InputOneOf2) GetInputThatChoosesNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputThatChoosesNum, true
}

// SetInputThatChoosesNum sets field value
func (o *InputOneOf2) SetInputThatChoosesNum(v int32) {
	o.InputThatChoosesNum = v
}

func (o InputOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["for_choice_id"] = o.ForChoiceId
	toSerialize["input_that_chooses_num"] = o.InputThatChoosesNum
	return toSerialize, nil
}

type NullableInputOneOf2 struct {
	value *InputOneOf2
	isSet bool
}

func (v NullableInputOneOf2) Get() *InputOneOf2 {
	return v.value
}

func (v *NullableInputOneOf2) Set(val *InputOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOneOf2(val *InputOneOf2) *NullableInputOneOf2 {
	return &NullableInputOneOf2{value: val, isSet: true}
}

func (v NullableInputOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

