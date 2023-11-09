/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PayeeOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayeeOneOf{}

// PayeeOneOf Pays funds into a party's account in the contract.
type PayeeOneOf struct {
	Account Party `json:"account"`
}

// NewPayeeOneOf instantiates a new PayeeOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayeeOneOf(account Party) *PayeeOneOf {
	this := PayeeOneOf{}
	this.Account = account
	return &this
}

// NewPayeeOneOfWithDefaults instantiates a new PayeeOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayeeOneOfWithDefaults() *PayeeOneOf {
	this := PayeeOneOf{}
	return &this
}

// GetAccount returns the Account field value
func (o *PayeeOneOf) GetAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *PayeeOneOf) GetAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *PayeeOneOf) SetAccount(v Party) {
	o.Account = v
}

func (o PayeeOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayeeOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	return toSerialize, nil
}

type NullablePayeeOneOf struct {
	value *PayeeOneOf
	isSet bool
}

func (v NullablePayeeOneOf) Get() *PayeeOneOf {
	return v.value
}

func (v *NullablePayeeOneOf) Set(val *PayeeOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePayeeOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePayeeOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayeeOneOf(val *PayeeOneOf) *NullablePayeeOneOf {
	return &NullablePayeeOneOf{value: val, isSet: true}
}

func (v NullablePayeeOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayeeOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


