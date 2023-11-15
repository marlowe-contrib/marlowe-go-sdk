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

// checks if the ContractObjectOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractObjectOneOf2{}

// ContractObjectOneOf2 Wait for an action to be performed and apply the matching contract when it does. Apply the timeout contract if no actions have been performed in the timeout period.
type ContractObjectOneOf2 struct {
	Timeout int32 `json:"timeout"`
	TimeoutContinuation ContractObject `json:"timeout_continuation"`
	When []CaseObject `json:"when"`
}

// NewContractObjectOneOf2 instantiates a new ContractObjectOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractObjectOneOf2(timeout int32, timeoutContinuation ContractObject, when []CaseObject) *ContractObjectOneOf2 {
	this := ContractObjectOneOf2{}
	this.Timeout = timeout
	this.TimeoutContinuation = timeoutContinuation
	this.When = when
	return &this
}

// NewContractObjectOneOf2WithDefaults instantiates a new ContractObjectOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractObjectOneOf2WithDefaults() *ContractObjectOneOf2 {
	this := ContractObjectOneOf2{}
	return &this
}

// GetTimeout returns the Timeout field value
func (o *ContractObjectOneOf2) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *ContractObjectOneOf2) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *ContractObjectOneOf2) SetTimeout(v int32) {
	o.Timeout = v
}

// GetTimeoutContinuation returns the TimeoutContinuation field value
func (o *ContractObjectOneOf2) GetTimeoutContinuation() ContractObject {
	if o == nil {
		var ret ContractObject
		return ret
	}

	return o.TimeoutContinuation
}

// GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field value
// and a boolean to check if the value has been set.
func (o *ContractObjectOneOf2) GetTimeoutContinuationOk() (*ContractObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutContinuation, true
}

// SetTimeoutContinuation sets field value
func (o *ContractObjectOneOf2) SetTimeoutContinuation(v ContractObject) {
	o.TimeoutContinuation = v
}

// GetWhen returns the When field value
func (o *ContractObjectOneOf2) GetWhen() []CaseObject {
	if o == nil {
		var ret []CaseObject
		return ret
	}

	return o.When
}

// GetWhenOk returns a tuple with the When field value
// and a boolean to check if the value has been set.
func (o *ContractObjectOneOf2) GetWhenOk() ([]CaseObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.When, true
}

// SetWhen sets field value
func (o *ContractObjectOneOf2) SetWhen(v []CaseObject) {
	o.When = v
}

func (o ContractObjectOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractObjectOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeout"] = o.Timeout
	toSerialize["timeout_continuation"] = o.TimeoutContinuation
	toSerialize["when"] = o.When
	return toSerialize, nil
}

type NullableContractObjectOneOf2 struct {
	value *ContractObjectOneOf2
	isSet bool
}

func (v NullableContractObjectOneOf2) Get() *ContractObjectOneOf2 {
	return v.value
}

func (v *NullableContractObjectOneOf2) Set(val *ContractObjectOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableContractObjectOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableContractObjectOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractObjectOneOf2(val *ContractObjectOneOf2) *NullableContractObjectOneOf2 {
	return &NullableContractObjectOneOf2{value: val, isSet: true}
}

func (v NullableContractObjectOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractObjectOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


