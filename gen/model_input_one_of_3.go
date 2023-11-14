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

// checks if the InputOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputOneOf3{}

// InputOneOf3 Deposit funds into an account in a contract and provide the continuation of the contract
type InputOneOf3 struct {
	ContinuationHash string `json:"continuation_hash"`
	InputFromParty Party `json:"input_from_party"`
	IntoAccount Party `json:"into_account"`
	MerkleizedContinuation Contract `json:"merkleized_continuation"`
	OfToken Token `json:"of_token"`
	ThatDeposits int32 `json:"that_deposits"`
}

// NewInputOneOf3 instantiates a new InputOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOneOf3(continuationHash string, inputFromParty Party, intoAccount Party, merkleizedContinuation Contract, ofToken Token, thatDeposits int32) *InputOneOf3 {
	this := InputOneOf3{}
	this.ContinuationHash = continuationHash
	this.InputFromParty = inputFromParty
	this.IntoAccount = intoAccount
	this.MerkleizedContinuation = merkleizedContinuation
	this.OfToken = ofToken
	this.ThatDeposits = thatDeposits
	return &this
}

// NewInputOneOf3WithDefaults instantiates a new InputOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOneOf3WithDefaults() *InputOneOf3 {
	this := InputOneOf3{}
	return &this
}

// GetContinuationHash returns the ContinuationHash field value
func (o *InputOneOf3) GetContinuationHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinuationHash
}

// GetContinuationHashOk returns a tuple with the ContinuationHash field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetContinuationHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinuationHash, true
}

// SetContinuationHash sets field value
func (o *InputOneOf3) SetContinuationHash(v string) {
	o.ContinuationHash = v
}

// GetInputFromParty returns the InputFromParty field value
func (o *InputOneOf3) GetInputFromParty() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.InputFromParty
}

// GetInputFromPartyOk returns a tuple with the InputFromParty field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetInputFromPartyOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputFromParty, true
}

// SetInputFromParty sets field value
func (o *InputOneOf3) SetInputFromParty(v Party) {
	o.InputFromParty = v
}

// GetIntoAccount returns the IntoAccount field value
func (o *InputOneOf3) GetIntoAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.IntoAccount
}

// GetIntoAccountOk returns a tuple with the IntoAccount field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetIntoAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntoAccount, true
}

// SetIntoAccount sets field value
func (o *InputOneOf3) SetIntoAccount(v Party) {
	o.IntoAccount = v
}

// GetMerkleizedContinuation returns the MerkleizedContinuation field value
func (o *InputOneOf3) GetMerkleizedContinuation() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.MerkleizedContinuation
}

// GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetMerkleizedContinuationOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleizedContinuation, true
}

// SetMerkleizedContinuation sets field value
func (o *InputOneOf3) SetMerkleizedContinuation(v Contract) {
	o.MerkleizedContinuation = v
}

// GetOfToken returns the OfToken field value
func (o *InputOneOf3) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *InputOneOf3) SetOfToken(v Token) {
	o.OfToken = v
}

// GetThatDeposits returns the ThatDeposits field value
func (o *InputOneOf3) GetThatDeposits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ThatDeposits
}

// GetThatDepositsOk returns a tuple with the ThatDeposits field value
// and a boolean to check if the value has been set.
func (o *InputOneOf3) GetThatDepositsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThatDeposits, true
}

// SetThatDeposits sets field value
func (o *InputOneOf3) SetThatDeposits(v int32) {
	o.ThatDeposits = v
}

func (o InputOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["continuation_hash"] = o.ContinuationHash
	toSerialize["input_from_party"] = o.InputFromParty
	toSerialize["into_account"] = o.IntoAccount
	toSerialize["merkleized_continuation"] = o.MerkleizedContinuation
	toSerialize["of_token"] = o.OfToken
	toSerialize["that_deposits"] = o.ThatDeposits
	return toSerialize, nil
}

type NullableInputOneOf3 struct {
	value *InputOneOf3
	isSet bool
}

func (v NullableInputOneOf3) Get() *InputOneOf3 {
	return v.value
}

func (v *NullableInputOneOf3) Set(val *InputOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOneOf3(val *InputOneOf3) *NullableInputOneOf3 {
	return &NullableInputOneOf3{value: val, isSet: true}
}

func (v NullableInputOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


