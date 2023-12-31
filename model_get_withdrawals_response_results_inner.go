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

// checks if the GetWithdrawalsResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWithdrawalsResponseResultsInner{}

// GetWithdrawalsResponseResultsInner struct for GetWithdrawalsResponseResultsInner
type GetWithdrawalsResponseResultsInner struct {
	Links GetWithdrawalsResponseResultsInnerLinks `json:"links"`
	Resource WithdrawalHeader `json:"resource"`
}

// NewGetWithdrawalsResponseResultsInner instantiates a new GetWithdrawalsResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWithdrawalsResponseResultsInner(links GetWithdrawalsResponseResultsInnerLinks, resource WithdrawalHeader) *GetWithdrawalsResponseResultsInner {
	this := GetWithdrawalsResponseResultsInner{}
	this.Links = links
	this.Resource = resource
	return &this
}

// NewGetWithdrawalsResponseResultsInnerWithDefaults instantiates a new GetWithdrawalsResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWithdrawalsResponseResultsInnerWithDefaults() *GetWithdrawalsResponseResultsInner {
	this := GetWithdrawalsResponseResultsInner{}
	return &this
}

// GetLinks returns the Links field value
func (o *GetWithdrawalsResponseResultsInner) GetLinks() GetWithdrawalsResponseResultsInnerLinks {
	if o == nil {
		var ret GetWithdrawalsResponseResultsInnerLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetWithdrawalsResponseResultsInner) GetLinksOk() (*GetWithdrawalsResponseResultsInnerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetWithdrawalsResponseResultsInner) SetLinks(v GetWithdrawalsResponseResultsInnerLinks) {
	o.Links = v
}

// GetResource returns the Resource field value
func (o *GetWithdrawalsResponseResultsInner) GetResource() WithdrawalHeader {
	if o == nil {
		var ret WithdrawalHeader
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *GetWithdrawalsResponseResultsInner) GetResourceOk() (*WithdrawalHeader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *GetWithdrawalsResponseResultsInner) SetResource(v WithdrawalHeader) {
	o.Resource = v
}

func (o GetWithdrawalsResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWithdrawalsResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableGetWithdrawalsResponseResultsInner struct {
	value *GetWithdrawalsResponseResultsInner
	isSet bool
}

func (v NullableGetWithdrawalsResponseResultsInner) Get() *GetWithdrawalsResponseResultsInner {
	return v.value
}

func (v *NullableGetWithdrawalsResponseResultsInner) Set(val *GetWithdrawalsResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWithdrawalsResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWithdrawalsResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWithdrawalsResponseResultsInner(val *GetWithdrawalsResponseResultsInner) *NullableGetWithdrawalsResponseResultsInner {
	return &NullableGetWithdrawalsResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetWithdrawalsResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWithdrawalsResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


