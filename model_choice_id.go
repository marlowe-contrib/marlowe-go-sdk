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

// checks if the ChoiceId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChoiceId{}

// ChoiceId Refers to a party by role name.
type ChoiceId struct {
	ChoiceName string `json:"choice_name"`
	ChoiceOwner Party `json:"choice_owner"`
}

// NewChoiceId instantiates a new ChoiceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChoiceId(choiceName string, choiceOwner Party) *ChoiceId {
	this := ChoiceId{}
	this.ChoiceName = choiceName
	this.ChoiceOwner = choiceOwner
	return &this
}

// NewChoiceIdWithDefaults instantiates a new ChoiceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChoiceIdWithDefaults() *ChoiceId {
	this := ChoiceId{}
	return &this
}

// GetChoiceName returns the ChoiceName field value
func (o *ChoiceId) GetChoiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChoiceName
}

// GetChoiceNameOk returns a tuple with the ChoiceName field value
// and a boolean to check if the value has been set.
func (o *ChoiceId) GetChoiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoiceName, true
}

// SetChoiceName sets field value
func (o *ChoiceId) SetChoiceName(v string) {
	o.ChoiceName = v
}

// GetChoiceOwner returns the ChoiceOwner field value
func (o *ChoiceId) GetChoiceOwner() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.ChoiceOwner
}

// GetChoiceOwnerOk returns a tuple with the ChoiceOwner field value
// and a boolean to check if the value has been set.
func (o *ChoiceId) GetChoiceOwnerOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoiceOwner, true
}

// SetChoiceOwner sets field value
func (o *ChoiceId) SetChoiceOwner(v Party) {
	o.ChoiceOwner = v
}

func (o ChoiceId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChoiceId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["choice_name"] = o.ChoiceName
	toSerialize["choice_owner"] = o.ChoiceOwner
	return toSerialize, nil
}

type NullableChoiceId struct {
	value *ChoiceId
	isSet bool
}

func (v NullableChoiceId) Get() *ChoiceId {
	return v.value
}

func (v *NullableChoiceId) Set(val *ChoiceId) {
	v.value = val
	v.isSet = true
}

func (v NullableChoiceId) IsSet() bool {
	return v.isSet
}

func (v *NullableChoiceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChoiceId(val *ChoiceId) *NullableChoiceId {
	return &NullableChoiceId{value: val, isSet: true}
}

func (v NullableChoiceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChoiceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


