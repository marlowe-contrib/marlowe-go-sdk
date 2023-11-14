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

// checks if the ListObjectPayoutHeaderResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListObjectPayoutHeaderResultsInner{}

// ListObjectPayoutHeaderResultsInner struct for ListObjectPayoutHeaderResultsInner
type ListObjectPayoutHeaderResultsInner struct {
	Links ListObjectPayoutHeaderResultsInnerLinks `json:"links"`
	Resource PayoutHeader `json:"resource"`
}

// NewListObjectPayoutHeaderResultsInner instantiates a new ListObjectPayoutHeaderResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectPayoutHeaderResultsInner(links ListObjectPayoutHeaderResultsInnerLinks, resource PayoutHeader) *ListObjectPayoutHeaderResultsInner {
	this := ListObjectPayoutHeaderResultsInner{}
	this.Links = links
	this.Resource = resource
	return &this
}

// NewListObjectPayoutHeaderResultsInnerWithDefaults instantiates a new ListObjectPayoutHeaderResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectPayoutHeaderResultsInnerWithDefaults() *ListObjectPayoutHeaderResultsInner {
	this := ListObjectPayoutHeaderResultsInner{}
	return &this
}

// GetLinks returns the Links field value
func (o *ListObjectPayoutHeaderResultsInner) GetLinks() ListObjectPayoutHeaderResultsInnerLinks {
	if o == nil {
		var ret ListObjectPayoutHeaderResultsInnerLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListObjectPayoutHeaderResultsInner) GetLinksOk() (*ListObjectPayoutHeaderResultsInnerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListObjectPayoutHeaderResultsInner) SetLinks(v ListObjectPayoutHeaderResultsInnerLinks) {
	o.Links = v
}

// GetResource returns the Resource field value
func (o *ListObjectPayoutHeaderResultsInner) GetResource() PayoutHeader {
	if o == nil {
		var ret PayoutHeader
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ListObjectPayoutHeaderResultsInner) GetResourceOk() (*PayoutHeader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ListObjectPayoutHeaderResultsInner) SetResource(v PayoutHeader) {
	o.Resource = v
}

func (o ListObjectPayoutHeaderResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListObjectPayoutHeaderResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableListObjectPayoutHeaderResultsInner struct {
	value *ListObjectPayoutHeaderResultsInner
	isSet bool
}

func (v NullableListObjectPayoutHeaderResultsInner) Get() *ListObjectPayoutHeaderResultsInner {
	return v.value
}

func (v *NullableListObjectPayoutHeaderResultsInner) Set(val *ListObjectPayoutHeaderResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectPayoutHeaderResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectPayoutHeaderResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectPayoutHeaderResultsInner(val *ListObjectPayoutHeaderResultsInner) *NullableListObjectPayoutHeaderResultsInner {
	return &NullableListObjectPayoutHeaderResultsInner{value: val, isSet: true}
}

func (v NullableListObjectPayoutHeaderResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectPayoutHeaderResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


