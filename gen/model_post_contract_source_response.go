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

// checks if the PostContractSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContractSourceResponse{}

// PostContractSourceResponse struct for PostContractSourceResponse
type PostContractSourceResponse struct {
	// The hex-encoded identifier of a Marlowe contract source
	ContractSourceId string `json:"contractSourceId"`
	IntermediateIds map[string]string `json:"intermediateIds"`
}

// NewPostContractSourceResponse instantiates a new PostContractSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContractSourceResponse(contractSourceId string, intermediateIds map[string]string) *PostContractSourceResponse {
	this := PostContractSourceResponse{}
	this.ContractSourceId = contractSourceId
	this.IntermediateIds = intermediateIds
	return &this
}

// NewPostContractSourceResponseWithDefaults instantiates a new PostContractSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContractSourceResponseWithDefaults() *PostContractSourceResponse {
	this := PostContractSourceResponse{}
	return &this
}

// GetContractSourceId returns the ContractSourceId field value
func (o *PostContractSourceResponse) GetContractSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractSourceId
}

// GetContractSourceIdOk returns a tuple with the ContractSourceId field value
// and a boolean to check if the value has been set.
func (o *PostContractSourceResponse) GetContractSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractSourceId, true
}

// SetContractSourceId sets field value
func (o *PostContractSourceResponse) SetContractSourceId(v string) {
	o.ContractSourceId = v
}

// GetIntermediateIds returns the IntermediateIds field value
func (o *PostContractSourceResponse) GetIntermediateIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.IntermediateIds
}

// GetIntermediateIdsOk returns a tuple with the IntermediateIds field value
// and a boolean to check if the value has been set.
func (o *PostContractSourceResponse) GetIntermediateIdsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntermediateIds, true
}

// SetIntermediateIds sets field value
func (o *PostContractSourceResponse) SetIntermediateIds(v map[string]string) {
	o.IntermediateIds = v
}

func (o PostContractSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostContractSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractSourceId"] = o.ContractSourceId
	toSerialize["intermediateIds"] = o.IntermediateIds
	return toSerialize, nil
}

type NullablePostContractSourceResponse struct {
	value *PostContractSourceResponse
	isSet bool
}

func (v NullablePostContractSourceResponse) Get() *PostContractSourceResponse {
	return v.value
}

func (v *NullablePostContractSourceResponse) Set(val *PostContractSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContractSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContractSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContractSourceResponse(val *PostContractSourceResponse) *NullablePostContractSourceResponse {
	return &NullablePostContractSourceResponse{value: val, isSet: true}
}

func (v NullablePostContractSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostContractSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


