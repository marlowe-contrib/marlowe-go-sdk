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

// checks if the DepositActionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositActionObject{}

// DepositActionObject struct for DepositActionObject
type DepositActionObject struct {
	Deposits Value `json:"deposits"`
	IntoAccount Party `json:"into_account"`
	OfToken Token `json:"of_token"`
	Party Party `json:"party"`
}

// NewDepositActionObject instantiates a new DepositActionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositActionObject(deposits Value, intoAccount Party, ofToken Token, party Party) *DepositActionObject {
	this := DepositActionObject{}
	this.Deposits = deposits
	this.IntoAccount = intoAccount
	this.OfToken = ofToken
	this.Party = party
	return &this
}

// NewDepositActionObjectWithDefaults instantiates a new DepositActionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositActionObjectWithDefaults() *DepositActionObject {
	this := DepositActionObject{}
	return &this
}

// GetDeposits returns the Deposits field value
func (o *DepositActionObject) GetDeposits() Value {
	if o == nil {
		var ret Value
		return ret
	}

	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value
// and a boolean to check if the value has been set.
func (o *DepositActionObject) GetDepositsOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deposits, true
}

// SetDeposits sets field value
func (o *DepositActionObject) SetDeposits(v Value) {
	o.Deposits = v
}

// GetIntoAccount returns the IntoAccount field value
func (o *DepositActionObject) GetIntoAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.IntoAccount
}

// GetIntoAccountOk returns a tuple with the IntoAccount field value
// and a boolean to check if the value has been set.
func (o *DepositActionObject) GetIntoAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntoAccount, true
}

// SetIntoAccount sets field value
func (o *DepositActionObject) SetIntoAccount(v Party) {
	o.IntoAccount = v
}

// GetOfToken returns the OfToken field value
func (o *DepositActionObject) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *DepositActionObject) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *DepositActionObject) SetOfToken(v Token) {
	o.OfToken = v
}

// GetParty returns the Party field value
func (o *DepositActionObject) GetParty() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.Party
}

// GetPartyOk returns a tuple with the Party field value
// and a boolean to check if the value has been set.
func (o *DepositActionObject) GetPartyOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Party, true
}

// SetParty sets field value
func (o *DepositActionObject) SetParty(v Party) {
	o.Party = v
}

func (o DepositActionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositActionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deposits"] = o.Deposits
	toSerialize["into_account"] = o.IntoAccount
	toSerialize["of_token"] = o.OfToken
	toSerialize["party"] = o.Party
	return toSerialize, nil
}

type NullableDepositActionObject struct {
	value *DepositActionObject
	isSet bool
}

func (v NullableDepositActionObject) Get() *DepositActionObject {
	return v.value
}

func (v *NullableDepositActionObject) Set(val *DepositActionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositActionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositActionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositActionObject(val *DepositActionObject) *NullableDepositActionObject {
	return &NullableDepositActionObject{value: val, isSet: true}
}

func (v NullableDepositActionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositActionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


