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

// checks if the TxOutputError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TxOutputError{}

// TxOutputError Marlowe transaction error.
type TxOutputError struct {
	TransactionError TransactionError `json:"transaction_error"`
}

// NewTxOutputError instantiates a new TxOutputError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxOutputError(transactionError TransactionError) *TxOutputError {
	this := TxOutputError{}
	this.TransactionError = transactionError
	return &this
}

// NewTxOutputErrorWithDefaults instantiates a new TxOutputError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxOutputErrorWithDefaults() *TxOutputError {
	this := TxOutputError{}
	return &this
}

// GetTransactionError returns the TransactionError field value
func (o *TxOutputError) GetTransactionError() TransactionError {
	if o == nil {
		var ret TransactionError
		return ret
	}

	return o.TransactionError
}

// GetTransactionErrorOk returns a tuple with the TransactionError field value
// and a boolean to check if the value has been set.
func (o *TxOutputError) GetTransactionErrorOk() (*TransactionError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionError, true
}

// SetTransactionError sets field value
func (o *TxOutputError) SetTransactionError(v TransactionError) {
	o.TransactionError = v
}

func (o TxOutputError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxOutputError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_error"] = o.TransactionError
	return toSerialize, nil
}

type NullableTxOutputError struct {
	value *TxOutputError
	isSet bool
}

func (v NullableTxOutputError) Get() *TxOutputError {
	return v.value
}

func (v *NullableTxOutputError) Set(val *TxOutputError) {
	v.value = val
	v.isSet = true
}

func (v NullableTxOutputError) IsSet() bool {
	return v.isSet
}

func (v *NullableTxOutputError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxOutputError(val *TxOutputError) *NullableTxOutputError {
	return &NullableTxOutputError{value: val, isSet: true}
}

func (v NullableTxOutputError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxOutputError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


