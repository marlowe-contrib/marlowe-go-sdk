/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListObjectContractSourceId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListObjectContractSourceId{}

// ListObjectContractSourceId struct for ListObjectContractSourceId
type ListObjectContractSourceId struct {
	Results []string `json:"results"`
}

// NewListObjectContractSourceId instantiates a new ListObjectContractSourceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectContractSourceId(results []string) *ListObjectContractSourceId {
	this := ListObjectContractSourceId{}
	this.Results = results
	return &this
}

// NewListObjectContractSourceIdWithDefaults instantiates a new ListObjectContractSourceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectContractSourceIdWithDefaults() *ListObjectContractSourceId {
	this := ListObjectContractSourceId{}
	return &this
}

// GetResults returns the Results field value
func (o *ListObjectContractSourceId) GetResults() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ListObjectContractSourceId) GetResultsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ListObjectContractSourceId) SetResults(v []string) {
	o.Results = v
}

func (o ListObjectContractSourceId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListObjectContractSourceId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableListObjectContractSourceId struct {
	value *ListObjectContractSourceId
	isSet bool
}

func (v NullableListObjectContractSourceId) Get() *ListObjectContractSourceId {
	return v.value
}

func (v *NullableListObjectContractSourceId) Set(val *ListObjectContractSourceId) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectContractSourceId) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectContractSourceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectContractSourceId(val *ListObjectContractSourceId) *NullableListObjectContractSourceId {
	return &NullableListObjectContractSourceId{value: val, isSet: true}
}

func (v NullableListObjectContractSourceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectContractSourceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


