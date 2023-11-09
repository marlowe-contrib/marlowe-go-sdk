/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ActionOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionOneOf1{}

// ActionOneOf1 struct for ActionOneOf1
type ActionOneOf1 struct {
	ChooseBetween []Bound `json:"choose_between"`
	ForChoice ChoiceId `json:"for_choice"`
}

// NewActionOneOf1 instantiates a new ActionOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionOneOf1(chooseBetween []Bound, forChoice ChoiceId) *ActionOneOf1 {
	this := ActionOneOf1{}
	this.ChooseBetween = chooseBetween
	this.ForChoice = forChoice
	return &this
}

// NewActionOneOf1WithDefaults instantiates a new ActionOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionOneOf1WithDefaults() *ActionOneOf1 {
	this := ActionOneOf1{}
	return &this
}

// GetChooseBetween returns the ChooseBetween field value
func (o *ActionOneOf1) GetChooseBetween() []Bound {
	if o == nil {
		var ret []Bound
		return ret
	}

	return o.ChooseBetween
}

// GetChooseBetweenOk returns a tuple with the ChooseBetween field value
// and a boolean to check if the value has been set.
func (o *ActionOneOf1) GetChooseBetweenOk() ([]Bound, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChooseBetween, true
}

// SetChooseBetween sets field value
func (o *ActionOneOf1) SetChooseBetween(v []Bound) {
	o.ChooseBetween = v
}

// GetForChoice returns the ForChoice field value
func (o *ActionOneOf1) GetForChoice() ChoiceId {
	if o == nil {
		var ret ChoiceId
		return ret
	}

	return o.ForChoice
}

// GetForChoiceOk returns a tuple with the ForChoice field value
// and a boolean to check if the value has been set.
func (o *ActionOneOf1) GetForChoiceOk() (*ChoiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForChoice, true
}

// SetForChoice sets field value
func (o *ActionOneOf1) SetForChoice(v ChoiceId) {
	o.ForChoice = v
}

func (o ActionOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["choose_between"] = o.ChooseBetween
	toSerialize["for_choice"] = o.ForChoice
	return toSerialize, nil
}

type NullableActionOneOf1 struct {
	value *ActionOneOf1
	isSet bool
}

func (v NullableActionOneOf1) Get() *ActionOneOf1 {
	return v.value
}

func (v *NullableActionOneOf1) Set(val *ActionOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableActionOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableActionOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionOneOf1(val *ActionOneOf1) *NullableActionOneOf1 {
	return &NullableActionOneOf1{value: val, isSet: true}
}

func (v NullableActionOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

