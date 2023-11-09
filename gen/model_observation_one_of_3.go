/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationOneOf3{}

// ObservationOneOf3 struct for ObservationOneOf3
type ObservationOneOf3 struct {
	ChoseSomethingFor ChoiceId `json:"chose_something_for"`
}

// NewObservationOneOf3 instantiates a new ObservationOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationOneOf3(choseSomethingFor ChoiceId) *ObservationOneOf3 {
	this := ObservationOneOf3{}
	this.ChoseSomethingFor = choseSomethingFor
	return &this
}

// NewObservationOneOf3WithDefaults instantiates a new ObservationOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationOneOf3WithDefaults() *ObservationOneOf3 {
	this := ObservationOneOf3{}
	return &this
}

// GetChoseSomethingFor returns the ChoseSomethingFor field value
func (o *ObservationOneOf3) GetChoseSomethingFor() ChoiceId {
	if o == nil {
		var ret ChoiceId
		return ret
	}

	return o.ChoseSomethingFor
}

// GetChoseSomethingForOk returns a tuple with the ChoseSomethingFor field value
// and a boolean to check if the value has been set.
func (o *ObservationOneOf3) GetChoseSomethingForOk() (*ChoiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoseSomethingFor, true
}

// SetChoseSomethingFor sets field value
func (o *ObservationOneOf3) SetChoseSomethingFor(v ChoiceId) {
	o.ChoseSomethingFor = v
}

func (o ObservationOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chose_something_for"] = o.ChoseSomethingFor
	return toSerialize, nil
}

type NullableObservationOneOf3 struct {
	value *ObservationOneOf3
	isSet bool
}

func (v NullableObservationOneOf3) Get() *ObservationOneOf3 {
	return v.value
}

func (v *NullableObservationOneOf3) Set(val *ObservationOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationOneOf3(val *ObservationOneOf3) *NullableObservationOneOf3 {
	return &NullableObservationOneOf3{value: val, isSet: true}
}

func (v NullableObservationOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

