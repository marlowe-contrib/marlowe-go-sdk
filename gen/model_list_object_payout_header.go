/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListObjectPayoutHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListObjectPayoutHeader{}

// ListObjectPayoutHeader struct for ListObjectPayoutHeader
type ListObjectPayoutHeader struct {
	Results []ListObjectPayoutHeaderResultsInner `json:"results"`
}

// NewListObjectPayoutHeader instantiates a new ListObjectPayoutHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectPayoutHeader(results []ListObjectPayoutHeaderResultsInner) *ListObjectPayoutHeader {
	this := ListObjectPayoutHeader{}
	this.Results = results
	return &this
}

// NewListObjectPayoutHeaderWithDefaults instantiates a new ListObjectPayoutHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectPayoutHeaderWithDefaults() *ListObjectPayoutHeader {
	this := ListObjectPayoutHeader{}
	return &this
}

// GetResults returns the Results field value
func (o *ListObjectPayoutHeader) GetResults() []ListObjectPayoutHeaderResultsInner {
	if o == nil {
		var ret []ListObjectPayoutHeaderResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ListObjectPayoutHeader) GetResultsOk() ([]ListObjectPayoutHeaderResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ListObjectPayoutHeader) SetResults(v []ListObjectPayoutHeaderResultsInner) {
	o.Results = v
}

func (o ListObjectPayoutHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListObjectPayoutHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableListObjectPayoutHeader struct {
	value *ListObjectPayoutHeader
	isSet bool
}

func (v NullableListObjectPayoutHeader) Get() *ListObjectPayoutHeader {
	return v.value
}

func (v *NullableListObjectPayoutHeader) Set(val *ListObjectPayoutHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectPayoutHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectPayoutHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectPayoutHeader(val *ListObjectPayoutHeader) *NullableListObjectPayoutHeader {
	return &NullableListObjectPayoutHeader{value: val, isSet: true}
}

func (v NullableListObjectPayoutHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectPayoutHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

