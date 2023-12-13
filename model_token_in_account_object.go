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

// checks if the TokenInAccountObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenInAccountObject{}

// TokenInAccountObject struct for TokenInAccountObject
type TokenInAccountObject struct {
	AmountOfToken TokenObject `json:"amount_of_token"`
	InAccount PartyObject `json:"in_account"`
}

// NewTokenInAccountObject instantiates a new TokenInAccountObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenInAccountObject(amountOfToken TokenObject, inAccount PartyObject) *TokenInAccountObject {
	this := TokenInAccountObject{}
	this.AmountOfToken = amountOfToken
	this.InAccount = inAccount
	return &this
}

// NewTokenInAccountObjectWithDefaults instantiates a new TokenInAccountObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenInAccountObjectWithDefaults() *TokenInAccountObject {
	this := TokenInAccountObject{}
	return &this
}

// GetAmountOfToken returns the AmountOfToken field value
func (o *TokenInAccountObject) GetAmountOfToken() TokenObject {
	if o == nil {
		var ret TokenObject
		return ret
	}

	return o.AmountOfToken
}

// GetAmountOfTokenOk returns a tuple with the AmountOfToken field value
// and a boolean to check if the value has been set.
func (o *TokenInAccountObject) GetAmountOfTokenOk() (*TokenObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountOfToken, true
}

// SetAmountOfToken sets field value
func (o *TokenInAccountObject) SetAmountOfToken(v TokenObject) {
	o.AmountOfToken = v
}

// GetInAccount returns the InAccount field value
func (o *TokenInAccountObject) GetInAccount() PartyObject {
	if o == nil {
		var ret PartyObject
		return ret
	}

	return o.InAccount
}

// GetInAccountOk returns a tuple with the InAccount field value
// and a boolean to check if the value has been set.
func (o *TokenInAccountObject) GetInAccountOk() (*PartyObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAccount, true
}

// SetInAccount sets field value
func (o *TokenInAccountObject) SetInAccount(v PartyObject) {
	o.InAccount = v
}

func (o TokenInAccountObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenInAccountObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_of_token"] = o.AmountOfToken
	toSerialize["in_account"] = o.InAccount
	return toSerialize, nil
}

type NullableTokenInAccountObject struct {
	value *TokenInAccountObject
	isSet bool
}

func (v NullableTokenInAccountObject) Get() *TokenInAccountObject {
	return v.value
}

func (v *NullableTokenInAccountObject) Set(val *TokenInAccountObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenInAccountObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenInAccountObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenInAccountObject(val *TokenInAccountObject) *NullableTokenInAccountObject {
	return &NullableTokenInAccountObject{value: val, isSet: true}
}

func (v NullableTokenInAccountObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenInAccountObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

