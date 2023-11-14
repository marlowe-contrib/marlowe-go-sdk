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

// checks if the PayeeObjectOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayeeObjectOneOf{}

// PayeeObjectOneOf Pays funds into a party's account in the contract.
type PayeeObjectOneOf struct {
	Account PartyObject `json:"account"`
}

// NewPayeeObjectOneOf instantiates a new PayeeObjectOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayeeObjectOneOf(account PartyObject) *PayeeObjectOneOf {
	this := PayeeObjectOneOf{}
	this.Account = account
	return &this
}

// NewPayeeObjectOneOfWithDefaults instantiates a new PayeeObjectOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayeeObjectOneOfWithDefaults() *PayeeObjectOneOf {
	this := PayeeObjectOneOf{}
	return &this
}

// GetAccount returns the Account field value
func (o *PayeeObjectOneOf) GetAccount() PartyObject {
	if o == nil {
		var ret PartyObject
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *PayeeObjectOneOf) GetAccountOk() (*PartyObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *PayeeObjectOneOf) SetAccount(v PartyObject) {
	o.Account = v
}

func (o PayeeObjectOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayeeObjectOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	return toSerialize, nil
}

type NullablePayeeObjectOneOf struct {
	value *PayeeObjectOneOf
	isSet bool
}

func (v NullablePayeeObjectOneOf) Get() *PayeeObjectOneOf {
	return v.value
}

func (v *NullablePayeeObjectOneOf) Set(val *PayeeObjectOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePayeeObjectOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePayeeObjectOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayeeObjectOneOf(val *PayeeObjectOneOf) *NullablePayeeObjectOneOf {
	return &NullablePayeeObjectOneOf{value: val, isSet: true}
}

func (v NullablePayeeObjectOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayeeObjectOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


