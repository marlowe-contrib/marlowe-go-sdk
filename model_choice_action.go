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

// checks if the ChoiceAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChoiceAction{}

// ChoiceAction struct for ChoiceAction
type ChoiceAction struct {
	ChooseBetween []Bound `json:"choose_between"`
	ForChoice ChoiceId `json:"for_choice"`
}

// NewChoiceAction instantiates a new ChoiceAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChoiceAction(chooseBetween []Bound, forChoice ChoiceId) *ChoiceAction {
	this := ChoiceAction{}
	this.ChooseBetween = chooseBetween
	this.ForChoice = forChoice
	return &this
}

// NewChoiceActionWithDefaults instantiates a new ChoiceAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChoiceActionWithDefaults() *ChoiceAction {
	this := ChoiceAction{}
	return &this
}

// GetChooseBetween returns the ChooseBetween field value
func (o *ChoiceAction) GetChooseBetween() []Bound {
	if o == nil {
		var ret []Bound
		return ret
	}

	return o.ChooseBetween
}

// GetChooseBetweenOk returns a tuple with the ChooseBetween field value
// and a boolean to check if the value has been set.
func (o *ChoiceAction) GetChooseBetweenOk() ([]Bound, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChooseBetween, true
}

// SetChooseBetween sets field value
func (o *ChoiceAction) SetChooseBetween(v []Bound) {
	o.ChooseBetween = v
}

// GetForChoice returns the ForChoice field value
func (o *ChoiceAction) GetForChoice() ChoiceId {
	if o == nil {
		var ret ChoiceId
		return ret
	}

	return o.ForChoice
}

// GetForChoiceOk returns a tuple with the ForChoice field value
// and a boolean to check if the value has been set.
func (o *ChoiceAction) GetForChoiceOk() (*ChoiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForChoice, true
}

// SetForChoice sets field value
func (o *ChoiceAction) SetForChoice(v ChoiceId) {
	o.ForChoice = v
}

func (o ChoiceAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChoiceAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["choose_between"] = o.ChooseBetween
	toSerialize["for_choice"] = o.ForChoice
	return toSerialize, nil
}

type NullableChoiceAction struct {
	value *ChoiceAction
	isSet bool
}

func (v NullableChoiceAction) Get() *ChoiceAction {
	return v.value
}

func (v *NullableChoiceAction) Set(val *ChoiceAction) {
	v.value = val
	v.isSet = true
}

func (v NullableChoiceAction) IsSet() bool {
	return v.isSet
}

func (v *NullableChoiceAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChoiceAction(val *ChoiceAction) *NullableChoiceAction {
	return &NullableChoiceAction{value: val, isSet: true}
}

func (v NullableChoiceAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChoiceAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

