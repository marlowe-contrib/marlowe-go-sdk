/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationObjectOneOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationObjectOneOf6{}

// ObservationObjectOneOf6 struct for ObservationObjectOneOf6
type ObservationObjectOneOf6 struct {
	Lt ValueObject `json:"lt"`
	Value ValueObject `json:"value"`
}

// NewObservationObjectOneOf6 instantiates a new ObservationObjectOneOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationObjectOneOf6(lt ValueObject, value ValueObject) *ObservationObjectOneOf6 {
	this := ObservationObjectOneOf6{}
	this.Lt = lt
	this.Value = value
	return &this
}

// NewObservationObjectOneOf6WithDefaults instantiates a new ObservationObjectOneOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationObjectOneOf6WithDefaults() *ObservationObjectOneOf6 {
	this := ObservationObjectOneOf6{}
	return &this
}

// GetLt returns the Lt field value
func (o *ObservationObjectOneOf6) GetLt() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Lt
}

// GetLtOk returns a tuple with the Lt field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf6) GetLtOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lt, true
}

// SetLt sets field value
func (o *ObservationObjectOneOf6) SetLt(v ValueObject) {
	o.Lt = v
}

// GetValue returns the Value field value
func (o *ObservationObjectOneOf6) GetValue() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservationObjectOneOf6) GetValueOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ObservationObjectOneOf6) SetValue(v ValueObject) {
	o.Value = v
}

func (o ObservationObjectOneOf6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationObjectOneOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lt"] = o.Lt
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableObservationObjectOneOf6 struct {
	value *ObservationObjectOneOf6
	isSet bool
}

func (v NullableObservationObjectOneOf6) Get() *ObservationObjectOneOf6 {
	return v.value
}

func (v *NullableObservationObjectOneOf6) Set(val *ObservationObjectOneOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObjectOneOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObjectOneOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObjectOneOf6(val *ObservationObjectOneOf6) *NullableObservationObjectOneOf6 {
	return &NullableObservationObjectOneOf6{value: val, isSet: true}
}

func (v NullableObservationObjectOneOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObjectOneOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


