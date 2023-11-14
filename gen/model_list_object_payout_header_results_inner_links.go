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

// checks if the ListObjectPayoutHeaderResultsInnerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListObjectPayoutHeaderResultsInnerLinks{}

// ListObjectPayoutHeaderResultsInnerLinks struct for ListObjectPayoutHeaderResultsInnerLinks
type ListObjectPayoutHeaderResultsInnerLinks struct {
	Payout *string `json:"payout,omitempty"`
}

// NewListObjectPayoutHeaderResultsInnerLinks instantiates a new ListObjectPayoutHeaderResultsInnerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectPayoutHeaderResultsInnerLinks() *ListObjectPayoutHeaderResultsInnerLinks {
	this := ListObjectPayoutHeaderResultsInnerLinks{}
	return &this
}

// NewListObjectPayoutHeaderResultsInnerLinksWithDefaults instantiates a new ListObjectPayoutHeaderResultsInnerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectPayoutHeaderResultsInnerLinksWithDefaults() *ListObjectPayoutHeaderResultsInnerLinks {
	this := ListObjectPayoutHeaderResultsInnerLinks{}
	return &this
}

// GetPayout returns the Payout field value if set, zero value otherwise.
func (o *ListObjectPayoutHeaderResultsInnerLinks) GetPayout() string {
	if o == nil || IsNil(o.Payout) {
		var ret string
		return ret
	}
	return *o.Payout
}

// GetPayoutOk returns a tuple with the Payout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectPayoutHeaderResultsInnerLinks) GetPayoutOk() (*string, bool) {
	if o == nil || IsNil(o.Payout) {
		return nil, false
	}
	return o.Payout, true
}

// HasPayout returns a boolean if a field has been set.
func (o *ListObjectPayoutHeaderResultsInnerLinks) HasPayout() bool {
	if o != nil && !IsNil(o.Payout) {
		return true
	}

	return false
}

// SetPayout gets a reference to the given string and assigns it to the Payout field.
func (o *ListObjectPayoutHeaderResultsInnerLinks) SetPayout(v string) {
	o.Payout = &v
}

func (o ListObjectPayoutHeaderResultsInnerLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListObjectPayoutHeaderResultsInnerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payout) {
		toSerialize["payout"] = o.Payout
	}
	return toSerialize, nil
}

type NullableListObjectPayoutHeaderResultsInnerLinks struct {
	value *ListObjectPayoutHeaderResultsInnerLinks
	isSet bool
}

func (v NullableListObjectPayoutHeaderResultsInnerLinks) Get() *ListObjectPayoutHeaderResultsInnerLinks {
	return v.value
}

func (v *NullableListObjectPayoutHeaderResultsInnerLinks) Set(val *ListObjectPayoutHeaderResultsInnerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectPayoutHeaderResultsInnerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectPayoutHeaderResultsInnerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectPayoutHeaderResultsInnerLinks(val *ListObjectPayoutHeaderResultsInnerLinks) *NullableListObjectPayoutHeaderResultsInnerLinks {
	return &NullableListObjectPayoutHeaderResultsInnerLinks{value: val, isSet: true}
}

func (v NullableListObjectPayoutHeaderResultsInnerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectPayoutHeaderResultsInnerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


