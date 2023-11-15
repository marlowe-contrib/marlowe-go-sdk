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

// checks if the GetTransactionsResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionsResponseResultsInner{}

// GetTransactionsResponseResultsInner struct for GetTransactionsResponseResultsInner
type GetTransactionsResponseResultsInner struct {
	Links ApplyInputsResponseLinks `json:"links"`
	Resource TxHeader `json:"resource"`
}

// NewGetTransactionsResponseResultsInner instantiates a new GetTransactionsResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionsResponseResultsInner(links ApplyInputsResponseLinks, resource TxHeader) *GetTransactionsResponseResultsInner {
	this := GetTransactionsResponseResultsInner{}
	this.Links = links
	this.Resource = resource
	return &this
}

// NewGetTransactionsResponseResultsInnerWithDefaults instantiates a new GetTransactionsResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionsResponseResultsInnerWithDefaults() *GetTransactionsResponseResultsInner {
	this := GetTransactionsResponseResultsInner{}
	return &this
}

// GetLinks returns the Links field value
func (o *GetTransactionsResponseResultsInner) GetLinks() ApplyInputsResponseLinks {
	if o == nil {
		var ret ApplyInputsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetTransactionsResponseResultsInner) GetLinksOk() (*ApplyInputsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetTransactionsResponseResultsInner) SetLinks(v ApplyInputsResponseLinks) {
	o.Links = v
}

// GetResource returns the Resource field value
func (o *GetTransactionsResponseResultsInner) GetResource() TxHeader {
	if o == nil {
		var ret TxHeader
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *GetTransactionsResponseResultsInner) GetResourceOk() (*TxHeader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *GetTransactionsResponseResultsInner) SetResource(v TxHeader) {
	o.Resource = v
}

func (o GetTransactionsResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionsResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableGetTransactionsResponseResultsInner struct {
	value *GetTransactionsResponseResultsInner
	isSet bool
}

func (v NullableGetTransactionsResponseResultsInner) Get() *GetTransactionsResponseResultsInner {
	return v.value
}

func (v *NullableGetTransactionsResponseResultsInner) Set(val *GetTransactionsResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionsResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionsResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionsResponseResultsInner(val *GetTransactionsResponseResultsInner) *NullableGetTransactionsResponseResultsInner {
	return &NullableGetTransactionsResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetTransactionsResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionsResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


