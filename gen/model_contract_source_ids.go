/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContractSourceIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractSourceIds{}

// ContractSourceIds struct for ContractSourceIds
type ContractSourceIds struct {
	Results []string `json:"results"`
}

// NewContractSourceIds instantiates a new ContractSourceIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractSourceIds(results []string) *ContractSourceIds {
	this := ContractSourceIds{}
	this.Results = results
	return &this
}

// NewContractSourceIdsWithDefaults instantiates a new ContractSourceIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractSourceIdsWithDefaults() *ContractSourceIds {
	this := ContractSourceIds{}
	return &this
}

// GetResults returns the Results field value
func (o *ContractSourceIds) GetResults() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ContractSourceIds) GetResultsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ContractSourceIds) SetResults(v []string) {
	o.Results = v
}

func (o ContractSourceIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractSourceIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableContractSourceIds struct {
	value *ContractSourceIds
	isSet bool
}

func (v NullableContractSourceIds) Get() *ContractSourceIds {
	return v.value
}

func (v *NullableContractSourceIds) Set(val *ContractSourceIds) {
	v.value = val
	v.isSet = true
}

func (v NullableContractSourceIds) IsSet() bool {
	return v.isSet
}

func (v *NullableContractSourceIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractSourceIds(val *ContractSourceIds) *NullableContractSourceIds {
	return &NullableContractSourceIds{value: val, isSet: true}
}

func (v NullableContractSourceIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractSourceIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


