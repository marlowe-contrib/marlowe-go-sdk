/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationOneOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationOneOf6{}

// ObservationOneOf6 struct for ObservationOneOf6
type ObservationOneOf6 struct {
	Lt Value `json:"lt"`
	Value Value `json:"value"`
}

// NewObservationOneOf6 instantiates a new ObservationOneOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationOneOf6(lt Value, value Value) *ObservationOneOf6 {
	this := ObservationOneOf6{}
	this.Lt = lt
	this.Value = value
	return &this
}

// NewObservationOneOf6WithDefaults instantiates a new ObservationOneOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationOneOf6WithDefaults() *ObservationOneOf6 {
	this := ObservationOneOf6{}
	return &this
}

// GetLt returns the Lt field value
func (o *ObservationOneOf6) GetLt() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Lt
}

// GetLtOk returns a tuple with the Lt field value
// and a boolean to check if the value has been set.
func (o *ObservationOneOf6) GetLtOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lt, true
}

// SetLt sets field value
func (o *ObservationOneOf6) SetLt(v Value) {
	o.Lt = v
}

// GetValue returns the Value field value
func (o *ObservationOneOf6) GetValue() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservationOneOf6) GetValueOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ObservationOneOf6) SetValue(v Value) {
	o.Value = v
}

func (o ObservationOneOf6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationOneOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lt"] = o.Lt
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableObservationOneOf6 struct {
	value *ObservationOneOf6
	isSet bool
}

func (v NullableObservationOneOf6) Get() *ObservationOneOf6 {
	return v.value
}

func (v *NullableObservationOneOf6) Set(val *ObservationOneOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationOneOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationOneOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationOneOf6(val *ObservationOneOf6) *NullableObservationOneOf6 {
	return &NullableObservationOneOf6{value: val, isSet: true}
}

func (v NullableObservationOneOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationOneOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

