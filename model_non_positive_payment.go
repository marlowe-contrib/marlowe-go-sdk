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

// checks if the NonPositivePayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonPositivePayment{}

// NonPositivePayment A warning for a non-positive payment.
type NonPositivePayment struct {
	Account Party `json:"account"`
	AskedToPay int32 `json:"asked_to_pay"`
	OfToken Token `json:"of_token"`
	ToPayee Payee `json:"to_payee"`
}

// NewNonPositivePayment instantiates a new NonPositivePayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonPositivePayment(account Party, askedToPay int32, ofToken Token, toPayee Payee) *NonPositivePayment {
	this := NonPositivePayment{}
	this.Account = account
	this.AskedToPay = askedToPay
	this.OfToken = ofToken
	this.ToPayee = toPayee
	return &this
}

// NewNonPositivePaymentWithDefaults instantiates a new NonPositivePayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonPositivePaymentWithDefaults() *NonPositivePayment {
	this := NonPositivePayment{}
	return &this
}

// GetAccount returns the Account field value
func (o *NonPositivePayment) GetAccount() Party {
	if o == nil {
		var ret Party
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NonPositivePayment) GetAccountOk() (*Party, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NonPositivePayment) SetAccount(v Party) {
	o.Account = v
}

// GetAskedToPay returns the AskedToPay field value
func (o *NonPositivePayment) GetAskedToPay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AskedToPay
}

// GetAskedToPayOk returns a tuple with the AskedToPay field value
// and a boolean to check if the value has been set.
func (o *NonPositivePayment) GetAskedToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AskedToPay, true
}

// SetAskedToPay sets field value
func (o *NonPositivePayment) SetAskedToPay(v int32) {
	o.AskedToPay = v
}

// GetOfToken returns the OfToken field value
func (o *NonPositivePayment) GetOfToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.OfToken
}

// GetOfTokenOk returns a tuple with the OfToken field value
// and a boolean to check if the value has been set.
func (o *NonPositivePayment) GetOfTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfToken, true
}

// SetOfToken sets field value
func (o *NonPositivePayment) SetOfToken(v Token) {
	o.OfToken = v
}

// GetToPayee returns the ToPayee field value
func (o *NonPositivePayment) GetToPayee() Payee {
	if o == nil {
		var ret Payee
		return ret
	}

	return o.ToPayee
}

// GetToPayeeOk returns a tuple with the ToPayee field value
// and a boolean to check if the value has been set.
func (o *NonPositivePayment) GetToPayeeOk() (*Payee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToPayee, true
}

// SetToPayee sets field value
func (o *NonPositivePayment) SetToPayee(v Payee) {
	o.ToPayee = v
}

func (o NonPositivePayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonPositivePayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["asked_to_pay"] = o.AskedToPay
	toSerialize["of_token"] = o.OfToken
	toSerialize["to_payee"] = o.ToPayee
	return toSerialize, nil
}

type NullableNonPositivePayment struct {
	value *NonPositivePayment
	isSet bool
}

func (v NullableNonPositivePayment) Get() *NonPositivePayment {
	return v.value
}

func (v *NullableNonPositivePayment) Set(val *NonPositivePayment) {
	v.value = val
	v.isSet = true
}

func (v NullableNonPositivePayment) IsSet() bool {
	return v.isSet
}

func (v *NullableNonPositivePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonPositivePayment(val *NonPositivePayment) *NullableNonPositivePayment {
	return &NullableNonPositivePayment{value: val, isSet: true}
}

func (v NullableNonPositivePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonPositivePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


