/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationObjectOneOf8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationObjectOneOf8{}

// ObservationObjectOneOf8 struct for ObservationObjectOneOf8
type ObservationObjectOneOf8 struct {
	EqualTo ValueObject `json:"equal_to"`
	Value ValueObject `json:"value"`
}

// NewObservationObjectOneOf8 instantiates a new ObservationObjectOneOf8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationObjectOneOf8(equalTo ValueObject, value ValueObject) *ObservationObjectOneOf8 {
	this := ObservationObjectOneOf8{}
	this.EqualTo = equalTo
	this.Value = value
	return &this
}

// NewObservationObjectOneOf8WithDefaults instantiates a new ObservationObjectOneOf8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationObjectOneOf8WithDefaults() *ObservationObjectOneOf8 {
	this := ObservationObjectOneOf8{}
	return &this
}

// GetEqualTo returns the EqualTo field value
func (o *ObservationObjectOneOf8) GetEqualTo() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.EqualTo
}

// GetEqualToOk returns a tuple with the EqualTo field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf8) GetEqualToOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EqualTo, true
}

// SetEqualTo sets field value
func (o *ObservationObjectOneOf8) SetEqualTo(v ValueObject) {
	o.EqualTo = v
}

// GetValue returns the Value field value
func (o *ObservationObjectOneOf8) GetValue() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf8) GetValueOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ObservationObjectOneOf8) SetValue(v ValueObject) {
	o.Value = v
}

func (o ObservationObjectOneOf8) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationObjectOneOf8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["equal_to"] = o.EqualTo
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableObservationObjectOneOf8 struct {
	value *ObservationObjectOneOf8
	isSet bool
}

func (v NullableObservationObjectOneOf8) Get() *ObservationObjectOneOf8 {
	return v.value
}

func (v *NullableObservationObjectOneOf8) Set(val *ObservationObjectOneOf8) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObjectOneOf8) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObjectOneOf8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObjectOneOf8(val *ObservationObjectOneOf8) *NullableObservationObjectOneOf8 {
	return &NullableObservationObjectOneOf8{value: val, isSet: true}
}

func (v NullableObservationObjectOneOf8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObjectOneOf8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

