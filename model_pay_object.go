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

// checks if the PayObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayObject{}

// PayObject A payment will be sent from an account to a payee.
type PayObject struct {
	FromAccount PartyObject `json:"from_account"`
	Pay ValueObject `json:"pay"`
	Then ContractObject `json:"then"`
	To PayeeObject `json:"to"`
	Token TokenObject `json:"token"`
}

// NewPayObject instantiates a new PayObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayObject(fromAccount PartyObject, pay ValueObject, then ContractObject, to PayeeObject, token TokenObject) *PayObject {
	this := PayObject{}
	this.FromAccount = fromAccount
	this.Pay = pay
	this.Then = then
	this.To = to
	this.Token = token
	return &this
}

// NewPayObjectWithDefaults instantiates a new PayObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayObjectWithDefaults() *PayObject {
	this := PayObject{}
	return &this
}

// GetFromAccount returns the FromAccount field value
func (o *PayObject) GetFromAccount() PartyObject {
	if o == nil {
		var ret PartyObject
		return ret
	}

	return o.FromAccount
}

// GetFromAccountOk returns a tuple with the FromAccount field value
// and a boolean to check if the value has been set.
func (o *PayObject) GetFromAccountOk() (*PartyObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAccount, true
}

// SetFromAccount sets field value
func (o *PayObject) SetFromAccount(v PartyObject) {
	o.FromAccount = v
}

// GetPay returns the Pay field value
func (o *PayObject) GetPay() ValueObject {
	if o == nil {
		var ret ValueObject
		return ret
	}

	return o.Pay
}

// GetPayOk returns a tuple with the Pay field value
// and a boolean to check if the value has been set.
func (o *PayObject) GetPayOk() (*ValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pay, true
}

// SetPay sets field value
func (o *PayObject) SetPay(v ValueObject) {
	o.Pay = v
}

// GetThen returns the Then field value
func (o *PayObject) GetThen() ContractObject {
	if o == nil {
		var ret ContractObject
		return ret
	}

	return o.Then
}

// GetThenOk returns a tuple with the Then field value
// and a boolean to check if the value has been set.
func (o *PayObject) GetThenOk() (*ContractObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Then, true
}

// SetThen sets field value
func (o *PayObject) SetThen(v ContractObject) {
	o.Then = v
}

// GetTo returns the To field value
func (o *PayObject) GetTo() PayeeObject {
	if o == nil {
		var ret PayeeObject
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PayObject) GetToOk() (*PayeeObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PayObject) SetTo(v PayeeObject) {
	o.To = v
}

// GetToken returns the Token field value
func (o *PayObject) GetToken() TokenObject {
	if o == nil {
		var ret TokenObject
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PayObject) GetTokenOk() (*TokenObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PayObject) SetToken(v TokenObject) {
	o.Token = v
}

func (o PayObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_account"] = o.FromAccount
	toSerialize["pay"] = o.Pay
	toSerialize["then"] = o.Then
	toSerialize["to"] = o.To
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullablePayObject struct {
	value *PayObject
	isSet bool
}

func (v NullablePayObject) Get() *PayObject {
	return v.value
}

func (v *NullablePayObject) Set(val *PayObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePayObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePayObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayObject(val *PayObject) *NullablePayObject {
	return &NullablePayObject{value: val, isSet: true}
}

func (v NullablePayObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

