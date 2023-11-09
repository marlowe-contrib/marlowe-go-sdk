/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationObjectOneOf7 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationObjectOneOf7{}

// ObservationObjectOneOf7 struct for ObservationObjectOneOf7
type ObservationObjectOneOf7 struct {
	LeThan ValueObject `json:"le_than"`
	Value ValueObject `json:"value"`
}

// NewObservationObjectOneOf7 instantiates a new ObservationObjectOneOf7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationObjectOneOf7(leThan ValueObject, value ValueObject) *ObservationObjectOneOf7 {
	this := ObservationObjectOneOf7{}
	this.LeThan = leThan
	this.Value = value
	return &this
}

// NewObservationObjectOneOf7WithDefaults instantiates a new ObservationObjectOneOf7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationObjectOneOf7WithDefaults() *ObservationObjectOneOf7 {
	this := ObservationObjectOneOf7{}
	return &this
}

// GetLeThan returns the LeThan field value
func (o *ObservationObjectOneOf7) GetLeThan() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.LeThan
}

// GetLeThanOk returns a tuple with the LeThan field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf7) GetLeThanOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeThan, true
}

// SetLeThan sets field value
func (o *ObservationObjectOneOf7) SetLeThan(v ValueObject) {
	o.LeThan = v
}

// GetValue returns the Value field value
func (o *ObservationObjectOneOf7) GetValue() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf7) GetValueOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ObservationObjectOneOf7) SetValue(v ValueObject) {
	o.Value = v
}

func (o ObservationObjectOneOf7) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationObjectOneOf7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["le_than"] = o.LeThan
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableObservationObjectOneOf7 struct {
	value *ObservationObjectOneOf7
	isSet bool
}

func (v NullableObservationObjectOneOf7) Get() *ObservationObjectOneOf7 {
	return v.value
}

func (v *NullableObservationObjectOneOf7) Set(val *ObservationObjectOneOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObjectOneOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObjectOneOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObjectOneOf7(val *ObservationObjectOneOf7) *NullableObservationObjectOneOf7 {
	return &NullableObservationObjectOneOf7{value: val, isSet: true}
}

func (v NullableObservationObjectOneOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObjectOneOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


