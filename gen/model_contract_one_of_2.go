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

// checks if the ContractOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractOneOf2{}

// ContractOneOf2 Wait for an action to be performed and apply the matching contract when it does. Apply the timeout contract if no actions have been performed in the timeout period.
type ContractOneOf2 struct {
	Timeout int32 `json:"timeout"`
	TimeoutContinuation Contract `json:"timeout_continuation"`
	When []Case `json:"when"`
}

// NewContractOneOf2 instantiates a new ContractOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractOneOf2(timeout int32, timeoutContinuation Contract, when []Case) *ContractOneOf2 {
	this := ContractOneOf2{}
	this.Timeout = timeout
	this.TimeoutContinuation = timeoutContinuation
	this.When = when
	return &this
}

// NewContractOneOf2WithDefaults instantiates a new ContractOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractOneOf2WithDefaults() *ContractOneOf2 {
	this := ContractOneOf2{}
	return &this
}

// GetTimeout returns the Timeout field value
func (o *ContractOneOf2) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *ContractOneOf2) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *ContractOneOf2) SetTimeout(v int32) {
	o.Timeout = v
}

// GetTimeoutContinuation returns the TimeoutContinuation field value
func (o *ContractOneOf2) GetTimeoutContinuation() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.TimeoutContinuation
}

// GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field value
// and a boolean to check if the value has been set.
func (o *ContractOneOf2) GetTimeoutContinuationOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutContinuation, true
}

// SetTimeoutContinuation sets field value
func (o *ContractOneOf2) SetTimeoutContinuation(v Contract) {
	o.TimeoutContinuation = v
}

// GetWhen returns the When field value
func (o *ContractOneOf2) GetWhen() []Case {
	if o == nil {
		var ret []Case
		return ret
	}

	return o.When
}

// GetWhenOk returns a tuple with the When field value
// and a boolean to check if the value has been set.
func (o *ContractOneOf2) GetWhenOk() ([]Case, bool) {
	if o == nil {
		return nil, false
	}
	return o.When, true
}

// SetWhen sets field value
func (o *ContractOneOf2) SetWhen(v []Case) {
	o.When = v
}

func (o ContractOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeout"] = o.Timeout
	toSerialize["timeout_continuation"] = o.TimeoutContinuation
	toSerialize["when"] = o.When
	return toSerialize, nil
}

type NullableContractOneOf2 struct {
	value *ContractOneOf2
	isSet bool
}

func (v NullableContractOneOf2) Get() *ContractOneOf2 {
	return v.value
}

func (v *NullableContractOneOf2) Set(val *ContractOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableContractOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableContractOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractOneOf2(val *ContractOneOf2) *NullableContractOneOf2 {
	return &NullableContractOneOf2{value: val, isSet: true}
}

func (v NullableContractOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


