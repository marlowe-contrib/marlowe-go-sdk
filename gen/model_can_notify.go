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

// checks if the CanNotify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CanNotify{}

// CanNotify Notify Input tha can be applied for a given contract
type CanNotify struct {
	// Index of a \"Case Action\" in a \"When\"
	CaseIndex int32 `json:"case_index"`
	// Indicates if a given contract continuation is merkleized
	IsMerkleizedContinuation bool `json:"is_merkleized_continuation"`
}

// NewCanNotify instantiates a new CanNotify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanNotify(caseIndex int32, isMerkleizedContinuation bool) *CanNotify {
	this := CanNotify{}
	this.CaseIndex = caseIndex
	this.IsMerkleizedContinuation = isMerkleizedContinuation
	return &this
}

// NewCanNotifyWithDefaults instantiates a new CanNotify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanNotifyWithDefaults() *CanNotify {
	this := CanNotify{}
	return &this
}

// GetCaseIndex returns the CaseIndex field value
func (o *CanNotify) GetCaseIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CaseIndex
}

// GetCaseIndexOk returns a tuple with the CaseIndex field value
// and a boolean to check if the value has been set.
func (o *CanNotify) GetCaseIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseIndex, true
}

// SetCaseIndex sets field value
func (o *CanNotify) SetCaseIndex(v int32) {
	o.CaseIndex = v
}

// GetIsMerkleizedContinuation returns the IsMerkleizedContinuation field value
func (o *CanNotify) GetIsMerkleizedContinuation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMerkleizedContinuation
}

// GetIsMerkleizedContinuationOk returns a tuple with the IsMerkleizedContinuation field value
// and a boolean to check if the value has been set.
func (o *CanNotify) GetIsMerkleizedContinuationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMerkleizedContinuation, true
}

// SetIsMerkleizedContinuation sets field value
func (o *CanNotify) SetIsMerkleizedContinuation(v bool) {
	o.IsMerkleizedContinuation = v
}

func (o CanNotify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CanNotify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["case_index"] = o.CaseIndex
	toSerialize["is_merkleized_continuation"] = o.IsMerkleizedContinuation
	return toSerialize, nil
}

type NullableCanNotify struct {
	value *CanNotify
	isSet bool
}

func (v NullableCanNotify) Get() *CanNotify {
	return v.value
}

func (v *NullableCanNotify) Set(val *CanNotify) {
	v.value = val
	v.isSet = true
}

func (v NullableCanNotify) IsSet() bool {
	return v.isSet
}

func (v *NullableCanNotify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanNotify(val *CanNotify) *NullableCanNotify {
	return &NullableCanNotify{value: val, isSet: true}
}

func (v NullableCanNotify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanNotify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


