/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransactionWarningOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionWarningOneOf2{}

// TransactionWarningOneOf2 A warning for partial payment.
type TransactionWarningOneOf2 struct {
	Account Party `json:"account"`
	AskedToPay int32 `json:"asked_to_pay"`
	ButOnlyPaid int32 `json:"but_only_paid"`
	OfToken Token `json:"of_token"`
	ToPayee Payee `json:"to_payee"`
}

// NewTransactionWarningOneOf2 instantiates a new TransactionWarningOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionWarningOneOf2(account Party, askedToPay int32, butOnlyPaid int32, ofToken Token, toPayee Payee) *TransactionWarningOneOf2 {
	this := TransactionWarningOneOf2{}
	this.Account = account
	this.AskedToPay = askedToPay
	this.ButOnlyPaid = butOnlyPaid
	this.OfToken = ofToken
	this.ToPayee = toPayee
	return &this
}

// NewTransactionWarningOneOf2WithDefaults instantiates a new TransactionWarningOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWarningOneOf2WithDefaults() *TransactionWarningOneOf2 {
	this := TransactionWarningOneOf2{}
	return &this
}

// GetAccount returns the Account field value
func (o *TransactionWarningOneOf2) GetAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf2) GetAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *TransactionWarningOneOf2) SetAccount(v Party) {
	o.Account = v
}

// GetAskedToPay returns the AskedToPay field value
func (o *TransactionWarningOneOf2) GetAskedToPay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AskedToPay
}

// GetAskedToPayOk returns a tuple with the AskedToPay field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf2) GetAskedToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AskedToPay, true
}

// SetAskedToPay sets field value
func (o *TransactionWarningOneOf2) SetAskedToPay(v int32) {
	o.AskedToPay = v
}

// GetButOnlyPaid returns the ButOnlyPaid field value
func (o *TransactionWarningOneOf2) GetButOnlyPaid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ButOnlyPaid
}

// GetButOnlyPaidOk returns a tuple with the ButOnlyPaid field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf2) GetButOnlyPaidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButOnlyPaid, true
}

// SetButOnlyPaid sets field value
func (o *TransactionWarningOneOf2) SetButOnlyPaid(v int32) {
	o.ButOnlyPaid = v
}

// GetOfToken returns the OfToken field value
func (o *TransactionWarningOneOf2) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf2) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *TransactionWarningOneOf2) SetOfToken(v Token) {
	o.OfToken = v
}

// GetToPayee returns the ToPayee field value
func (o *TransactionWarningOneOf2) GetToPayee() Payee {
	if o == nil {
		var ret Payee
		return ret
	}

	return o.ToPayee
}

// GetToPayeeOk returns a tuple with the ToPayee field value
// and a boolean to check if the value has been set.
func (o *TransactionWarningOneOf2) GetToPayeeOk() (*Payee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToPayee, true
}

// SetToPayee sets field value
func (o *TransactionWarningOneOf2) SetToPayee(v Payee) {
	o.ToPayee = v
}

func (o TransactionWarningOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionWarningOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["asked_to_pay"] = o.AskedToPay
	toSerialize["but_only_paid"] = o.ButOnlyPaid
	toSerialize["of_token"] = o.OfToken
	toSerialize["to_payee"] = o.ToPayee
	return toSerialize, nil
}

type NullableTransactionWarningOneOf2 struct {
	value *TransactionWarningOneOf2
	isSet bool
}

func (v NullableTransactionWarningOneOf2) Get() *TransactionWarningOneOf2 {
	return v.value
}

func (v *NullableTransactionWarningOneOf2) Set(val *TransactionWarningOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionWarningOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionWarningOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionWarningOneOf2(val *TransactionWarningOneOf2) *NullableTransactionWarningOneOf2 {
	return &NullableTransactionWarningOneOf2{value: val, isSet: true}
}

func (v NullableTransactionWarningOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionWarningOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

