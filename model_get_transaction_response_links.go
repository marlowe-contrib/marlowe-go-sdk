/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marloweruntime

import (
	"encoding/json"
)

// checks if the GetTransactionResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionResponseLinks{}

// GetTransactionResponseLinks struct for GetTransactionResponseLinks
type GetTransactionResponseLinks struct {
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

// NewGetTransactionResponseLinks instantiates a new GetTransactionResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionResponseLinks() *GetTransactionResponseLinks {
	this := GetTransactionResponseLinks{}
	return &this
}

// NewGetTransactionResponseLinksWithDefaults instantiates a new GetTransactionResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionResponseLinksWithDefaults() *GetTransactionResponseLinks {
	this := GetTransactionResponseLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *GetTransactionResponseLinks) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *GetTransactionResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *GetTransactionResponseLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *GetTransactionResponseLinks) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseLinks) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *GetTransactionResponseLinks) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *GetTransactionResponseLinks) SetPrevious(v string) {
	o.Previous = &v
}

func (o GetTransactionResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	return toSerialize, nil
}

type NullableGetTransactionResponseLinks struct {
	value *GetTransactionResponseLinks
	isSet bool
}

func (v NullableGetTransactionResponseLinks) Get() *GetTransactionResponseLinks {
	return v.value
}

func (v *NullableGetTransactionResponseLinks) Set(val *GetTransactionResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionResponseLinks(val *GetTransactionResponseLinks) *NullableGetTransactionResponseLinks {
	return &NullableGetTransactionResponseLinks{value: val, isSet: true}
}

func (v NullableGetTransactionResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


