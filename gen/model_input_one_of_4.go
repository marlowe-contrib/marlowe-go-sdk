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

// checks if the InputOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputOneOf4{}

// InputOneOf4 Deposit funds into an account in a contract
type InputOneOf4 struct {
	InputFromParty Party `json:"input_from_party"`
	IntoAccount Party `json:"into_account"`
	OfToken Token `json:"of_token"`
	ThatDeposits int32 `json:"that_deposits"`
}

// NewInputOneOf4 instantiates a new InputOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOneOf4(inputFromParty Party, intoAccount Party, ofToken Token, thatDeposits int32) *InputOneOf4 {
	this := InputOneOf4{}
	this.InputFromParty = inputFromParty
	this.IntoAccount = intoAccount
	this.OfToken = ofToken
	this.ThatDeposits = thatDeposits
	return &this
}

// NewInputOneOf4WithDefaults instantiates a new InputOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOneOf4WithDefaults() *InputOneOf4 {
	this := InputOneOf4{}
	return &this
}

// GetInputFromParty returns the InputFromParty field value
func (o *InputOneOf4) GetInputFromParty() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.InputFromParty
}

// GetInputFromPartyOk returns a tuple with the InputFromParty field value
// and a boolean to check if the value has been set.
func (o *InputOneOf4) GetInputFromPartyOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputFromParty, true
}

// SetInputFromParty sets field value
func (o *InputOneOf4) SetInputFromParty(v Party) {
	o.InputFromParty = v
}

// GetIntoAccount returns the IntoAccount field value
func (o *InputOneOf4) GetIntoAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.IntoAccount
}

// GetIntoAccountOk returns a tuple with the IntoAccount field value
// and a boolean to check if the value has been set.
func (o *InputOneOf4) GetIntoAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntoAccount, true
}

// SetIntoAccount sets field value
func (o *InputOneOf4) SetIntoAccount(v Party) {
	o.IntoAccount = v
}

// GetOfToken returns the OfToken field value
func (o *InputOneOf4) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *InputOneOf4) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *InputOneOf4) SetOfToken(v Token) {
	o.OfToken = v
}

// GetThatDeposits returns the ThatDeposits field value
func (o *InputOneOf4) GetThatDeposits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ThatDeposits
}

// GetThatDepositsOk returns a tuple with the ThatDeposits field value
// and a boolean to check if the value has been set.
func (o *InputOneOf4) GetThatDepositsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThatDeposits, true
}

// SetThatDeposits sets field value
func (o *InputOneOf4) SetThatDeposits(v int32) {
	o.ThatDeposits = v
}

func (o InputOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input_from_party"] = o.InputFromParty
	toSerialize["into_account"] = o.IntoAccount
	toSerialize["of_token"] = o.OfToken
	toSerialize["that_deposits"] = o.ThatDeposits
	return toSerialize, nil
}

type NullableInputOneOf4 struct {
	value *InputOneOf4
	isSet bool
}

func (v NullableInputOneOf4) Get() *InputOneOf4 {
	return v.value
}

func (v *NullableInputOneOf4) Set(val *InputOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOneOf4(val *InputOneOf4) *NullableInputOneOf4 {
	return &NullableInputOneOf4{value: val, isSet: true}
}

func (v NullableInputOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


