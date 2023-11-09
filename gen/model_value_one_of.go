/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValueOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueOneOf{}

// ValueOneOf struct for ValueOneOf
type ValueOneOf struct {
	AmountOfToken Token `json:"amount_of_token"`
	InAccount Party `json:"in_account"`
}

// NewValueOneOf instantiates a new ValueOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueOneOf(amountOfToken Token, inAccount Party) *ValueOneOf {
	this := ValueOneOf{}
	this.AmountOfToken = amountOfToken
	this.InAccount = inAccount
	return &this
}

// NewValueOneOfWithDefaults instantiates a new ValueOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueOneOfWithDefaults() *ValueOneOf {
	this := ValueOneOf{}
	return &this
}

// GetAmountOfToken returns the AmountOfToken field value
func (o *ValueOneOf) GetAmountOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.AmountOfToken
}

// GetAmountOfTokenOk returns a tuple with the AmountOfToken field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf) GetAmountOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountOfToken, true
}

// SetAmountOfToken sets field value
func (o *ValueOneOf) SetAmountOfToken(v Token) {
	o.AmountOfToken = v
}

// GetInAccount returns the InAccount field value
func (o *ValueOneOf) GetInAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.InAccount
}

// GetInAccountOk returns a tuple with the InAccount field value
// and a boolean to check if the value has been set.
func (o *ValueOneOf) GetInAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAccount, true
}

// SetInAccount sets field value
func (o *ValueOneOf) SetInAccount(v Party) {
	o.InAccount = v
}

func (o ValueOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_of_token"] = o.AmountOfToken
	toSerialize["in_account"] = o.InAccount
	return toSerialize, nil
}

type NullableValueOneOf struct {
	value *ValueOneOf
	isSet bool
}

func (v NullableValueOneOf) Get() *ValueOneOf {
	return v.value
}

func (v *NullableValueOneOf) Set(val *ValueOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableValueOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableValueOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueOneOf(val *ValueOneOf) *NullableValueOneOf {
	return &NullableValueOneOf{value: val, isSet: true}
}

func (v NullableValueOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


