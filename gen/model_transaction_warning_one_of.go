/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransactionWarningOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionWarningOneOf{}

// TransactionWarningOneOf A warning for a non-positive deposit.
type TransactionWarningOneOf struct {
	AskedToDeposit int32 `json:"asked_to_deposit"`
	InAccount Party `json:"in_account"`
	OfToken Token `json:"of_token"`
	Party Party `json:"party"`
}

// NewTransactionWarningOneOf instantiates a new TransactionWarningOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionWarningOneOf(askedToDeposit int32, inAccount Party, ofToken Token, party Party) *TransactionWarningOneOf {
	this := TransactionWarningOneOf{}
	this.AskedToDeposit = askedToDeposit
	this.InAccount = inAccount
	this.OfToken = ofToken
	this.Party = party
	return &this
}

// NewTransactionWarningOneOfWithDefaults instantiates a new TransactionWarningOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWarningOneOfWithDefaults() *TransactionWarningOneOf {
	this := TransactionWarningOneOf{}
	return &this
}

// GetAskedToDeposit returns the AskedToDeposit field value
func (o *TransactionWarningOneOf) GetAskedToDeposit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AskedToDeposit
}

// GetAskedToDepositOk returns a tuple with the AskedToDeposit field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf) GetAskedToDepositOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AskedToDeposit, true
}

// SetAskedToDeposit sets field value
func (o *TransactionWarningOneOf) SetAskedToDeposit(v int32) {
	o.AskedToDeposit = v
}

// GetInAccount returns the InAccount field value
func (o *TransactionWarningOneOf) GetInAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.InAccount
}

// GetInAccountOk returns a tuple with the InAccount field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf) GetInAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAccount, true
}

// SetInAccount sets field value
func (o *TransactionWarningOneOf) SetInAccount(v Party) {
	o.InAccount = v
}

// GetOfToken returns the OfToken field value
func (o *TransactionWarningOneOf) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *TransactionWarningOneOf) SetOfToken(v Token) {
	o.OfToken = v
}

// GetParty returns the Party field value
func (o *TransactionWarningOneOf) GetParty() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.Party
}

// GetPartyOk returns a tuple with the Party field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf) GetPartyOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Party, true
}

// SetParty sets field value
func (o *TransactionWarningOneOf) SetParty(v Party) {
	o.Party = v
}

func (o TransactionWarningOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionWarningOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asked_to_deposit"] = o.AskedToDeposit
	toSerialize["in_account"] = o.InAccount
	toSerialize["of_token"] = o.OfToken
	toSerialize["party"] = o.Party
	return toSerialize, nil
}

type NullableTransactionWarningOneOf struct {
	value *TransactionWarningOneOf
	isSet bool
}

func (v NullableTransactionWarningOneOf) Get() *TransactionWarningOneOf {
	return v.value
}

func (v *NullableTransactionWarningOneOf) Set(val *TransactionWarningOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionWarningOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionWarningOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionWarningOneOf(val *TransactionWarningOneOf) *NullableTransactionWarningOneOf {
	return &NullableTransactionWarningOneOf{value: val, isSet: true}
}

func (v NullableTransactionWarningOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionWarningOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

