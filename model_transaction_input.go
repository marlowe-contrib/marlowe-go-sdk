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

// checks if the TransactionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInput{}

// TransactionInput Marlowe transaction input.
type TransactionInput struct {
	TxInputs []Input `json:"tx_inputs"`
	TxInterval TransactionInputTxInterval `json:"tx_interval"`
}

// NewTransactionInput instantiates a new TransactionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInput(txInputs []Input, txInterval TransactionInputTxInterval) *TransactionInput {
	this := TransactionInput{}
	this.TxInputs = txInputs
	this.TxInterval = txInterval
	return &this
}

// NewTransactionInputWithDefaults instantiates a new TransactionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInputWithDefaults() *TransactionInput {
	this := TransactionInput{}
	return &this
}

// GetTxInputs returns the TxInputs field value
func (o *TransactionInput) GetTxInputs() []Input {
	if o == nil {
		var ret []Input
		return ret
	}

	return o.TxInputs
}

// GetTxInputsOk returns a tuple with the TxInputs field value
// and a boolean to check if the value has been set.
func (o *TransactionInput) GetTxInputsOk() ([]Input, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxInputs, true
}

// SetTxInputs sets field value
func (o *TransactionInput) SetTxInputs(v []Input) {
	o.TxInputs = v
}

// GetTxInterval returns the TxInterval field value
func (o *TransactionInput) GetTxInterval() TransactionInputTxInterval {
	if o == nil {
		var ret TransactionInputTxInterval
		return ret
	}

	return o.TxInterval
}

// GetTxIntervalOk returns a tuple with the TxInterval field value
// and a boolean to check if the value has been set.
func (o *TransactionInput) GetTxIntervalOk() (*TransactionInputTxInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxInterval, true
}

// SetTxInterval sets field value
func (o *TransactionInput) SetTxInterval(v TransactionInputTxInterval) {
	o.TxInterval = v
}

func (o TransactionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_inputs"] = o.TxInputs
	toSerialize["tx_interval"] = o.TxInterval
	return toSerialize, nil
}

type NullableTransactionInput struct {
	value *TransactionInput
	isSet bool
}

func (v NullableTransactionInput) Get() *TransactionInput {
	return v.value
}

func (v *NullableTransactionInput) Set(val *TransactionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInput(val *TransactionInput) *NullableTransactionInput {
	return &NullableTransactionInput{value: val, isSet: true}
}

func (v NullableTransactionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


