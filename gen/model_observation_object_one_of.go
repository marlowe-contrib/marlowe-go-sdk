/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationObjectOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationObjectOneOf{}

// ObservationObjectOneOf struct for ObservationObjectOneOf
type ObservationObjectOneOf struct {
	And ObservationObject `json:"and"`
	Both ObservationObject `json:"both"`
}

// NewObservationObjectOneOf instantiates a new ObservationObjectOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationObjectOneOf(and ObservationObject, both ObservationObject) *ObservationObjectOneOf {
	this := ObservationObjectOneOf{}
	this.And = and
	this.Both = both
	return &this
}

// NewObservationObjectOneOfWithDefaults instantiates a new ObservationObjectOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationObjectOneOfWithDefaults() *ObservationObjectOneOf {
	this := ObservationObjectOneOf{}
	return &this
}

// GetAnd returns the And field value
func (o *ObservationObjectOneOf) GetAnd() ObservationObject {
	if o == nil {
		var ret ObservationObject
		return ret
	}

	return o.And
}

// GetAndOk returns a tuple with the And field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf) GetAndOk() (*ObservationObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.And, true
}

// SetAnd sets field value
func (o *ObservationObjectOneOf) SetAnd(v ObservationObject) {
	o.And = v
}

// GetBoth returns the Both field value
func (o *ObservationObjectOneOf) GetBoth() ObservationObject {
	if o == nil {
		var ret ObservationObject
		return ret
	}

	return o.Both
}

// GetBothOk returns a tuple with the Both field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf) GetBothOk() (*ObservationObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Both, true
}

// SetBoth sets field value
func (o *ObservationObjectOneOf) SetBoth(v ObservationObject) {
	o.Both = v
}

func (o ObservationObjectOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationObjectOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["and"] = o.And
	toSerialize["both"] = o.Both
	return toSerialize, nil
}

type NullableObservationObjectOneOf struct {
	value *ObservationObjectOneOf
	isSet bool
}

func (v NullableObservationObjectOneOf) Get() *ObservationObjectOneOf {
	return v.value
}

func (v *NullableObservationObjectOneOf) Set(val *ObservationObjectOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObjectOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObjectOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObjectOneOf(val *ObservationObjectOneOf) *NullableObservationObjectOneOf {
	return &NullableObservationObjectOneOf{value: val, isSet: true}
}

func (v NullableObservationObjectOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObjectOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


