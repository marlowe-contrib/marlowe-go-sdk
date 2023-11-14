/*
Marlowe Runtime REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContractsContractIdTransactionsPost201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractsContractIdTransactionsPost201Response{}

// ContractsContractIdTransactionsPost201Response struct for ContractsContractIdTransactionsPost201Response
type ContractsContractIdTransactionsPost201Response struct {
	Links ContractsContractIdTransactionsPost201ResponseLinks `json:"links"`
	Resource ApplyInputsTxEnvelope `json:"resource"`
}

// NewContractsContractIdTransactionsPost201Response instantiates a new ContractsContractIdTransactionsPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractsContractIdTransactionsPost201Response(links ContractsContractIdTransactionsPost201ResponseLinks, resource ApplyInputsTxEnvelope) *ContractsContractIdTransactionsPost201Response {
	this := ContractsContractIdTransactionsPost201Response{}
	this.Links = links
	this.Resource = resource
	return &this
}

// NewContractsContractIdTransactionsPost201ResponseWithDefaults instantiates a new ContractsContractIdTransactionsPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractsContractIdTransactionsPost201ResponseWithDefaults() *ContractsContractIdTransactionsPost201Response {
	this := ContractsContractIdTransactionsPost201Response{}
	return &this
}

// GetLinks returns the Links field value
func (o *ContractsContractIdTransactionsPost201Response) GetLinks() ContractsContractIdTransactionsPost201ResponseLinks {
	if o == nil {
		var ret ContractsContractIdTransactionsPost201ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ContractsContractIdTransactionsPost201Response) GetLinksOk() (*ContractsContractIdTransactionsPost201ResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ContractsContractIdTransactionsPost201Response) SetLinks(v ContractsContractIdTransactionsPost201ResponseLinks) {
	o.Links = v
}

// GetResource returns the Resource field value
func (o *ContractsContractIdTransactionsPost201Response) GetResource() ApplyInputsTxEnvelope {
	if o == nil {
		var ret ApplyInputsTxEnvelope
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ContractsContractIdTransactionsPost201Response) GetResourceOk() (*ApplyInputsTxEnvelope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ContractsContractIdTransactionsPost201Response) SetResource(v ApplyInputsTxEnvelope) {
	o.Resource = v
}

func (o ContractsContractIdTransactionsPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractsContractIdTransactionsPost201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableContractsContractIdTransactionsPost201Response struct {
	value *ContractsContractIdTransactionsPost201Response
	isSet bool
}

func (v NullableContractsContractIdTransactionsPost201Response) Get() *ContractsContractIdTransactionsPost201Response {
	return v.value
}

func (v *NullableContractsContractIdTransactionsPost201Response) Set(val *ContractsContractIdTransactionsPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableContractsContractIdTransactionsPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableContractsContractIdTransactionsPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractsContractIdTransactionsPost201Response(val *ContractsContractIdTransactionsPost201Response) *NullableContractsContractIdTransactionsPost201Response {
	return &NullableContractsContractIdTransactionsPost201Response{value: val, isSet: true}
}

func (v NullableContractsContractIdTransactionsPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractsContractIdTransactionsPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


