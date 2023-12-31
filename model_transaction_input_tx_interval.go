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

// checks if the TransactionInputTxInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInputTxInterval{}

// TransactionInputTxInterval Time interval.
type TransactionInputTxInterval struct {
	From int32 `json:"from"`
	To int32 `json:"to"`
}

// NewTransactionInputTxInterval instantiates a new TransactionInputTxInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInputTxInterval(from int32, to int32) *TransactionInputTxInterval {
	this := TransactionInputTxInterval{}
	this.From = from
	this.To = to
	return &this
}

// NewTransactionInputTxIntervalWithDefaults instantiates a new TransactionInputTxInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInputTxIntervalWithDefaults() *TransactionInputTxInterval {
	this := TransactionInputTxInterval{}
	return &this
}

// GetFrom returns the From field value
func (o *TransactionInputTxInterval) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *TransactionInputTxInterval) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *TransactionInputTxInterval) SetFrom(v int32) {
	o.From = v
}

// GetTo returns the To field value
func (o *TransactionInputTxInterval) GetTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *TransactionInputTxInterval) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *TransactionInputTxInterval) SetTo(v int32) {
	o.To = v
}

func (o TransactionInputTxInterval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInputTxInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return toSerialize, nil
}

type NullableTransactionInputTxInterval struct {
	value *TransactionInputTxInterval
	isSet bool
}

func (v NullableTransactionInputTxInterval) Get() *TransactionInputTxInterval {
	return v.value
}

func (v *NullableTransactionInputTxInterval) Set(val *TransactionInputTxInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInputTxInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInputTxInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInputTxInterval(val *TransactionInputTxInterval) *NullableTransactionInputTxInterval {
	return &NullableTransactionInputTxInterval{value: val, isSet: true}
}

func (v NullableTransactionInputTxInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInputTxInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


