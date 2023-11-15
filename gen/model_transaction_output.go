/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TransactionOutput - Marlowe transaction output.
type TransactionOutput struct {
	TransactionOutputOneOf *TransactionOutputOneOf
	TransactionOutputOneOf1 *TransactionOutputOneOf1
}

// TransactionOutputOneOfAsTransactionOutput is a convenience function that returns TransactionOutputOneOf wrapped in TransactionOutput
func TransactionOutputOneOfAsTransactionOutput(v *TransactionOutputOneOf) TransactionOutput {
	return TransactionOutput{
		TransactionOutputOneOf: v,
	}
}

// TransactionOutputOneOf1AsTransactionOutput is a convenience function that returns TransactionOutputOneOf1 wrapped in TransactionOutput
func TransactionOutputOneOf1AsTransactionOutput(v *TransactionOutputOneOf1) TransactionOutput {
	return TransactionOutput{
		TransactionOutputOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionOutput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TransactionOutputOneOf
	err = newStrictDecoder(data).Decode(&dst.TransactionOutputOneOf)
	if err == nil {
		jsonTransactionOutputOneOf, _ := json.Marshal(dst.TransactionOutputOneOf)
		if string(jsonTransactionOutputOneOf) == "{}" { // empty struct
			dst.TransactionOutputOneOf = nil
		} else {
			match++
		}
	} else {
		dst.TransactionOutputOneOf = nil
	}

	// try to unmarshal data into TransactionOutputOneOf1
	err = newStrictDecoder(data).Decode(&dst.TransactionOutputOneOf1)
	if err == nil {
		jsonTransactionOutputOneOf1, _ := json.Marshal(dst.TransactionOutputOneOf1)
		if string(jsonTransactionOutputOneOf1) == "{}" { // empty struct
			dst.TransactionOutputOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.TransactionOutputOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TransactionOutputOneOf = nil
		dst.TransactionOutputOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransactionOutput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransactionOutput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionOutput) MarshalJSON() ([]byte, error) {
	if src.TransactionOutputOneOf != nil {
		return json.Marshal(&src.TransactionOutputOneOf)
	}

	if src.TransactionOutputOneOf1 != nil {
		return json.Marshal(&src.TransactionOutputOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionOutput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionOutputOneOf != nil {
		return obj.TransactionOutputOneOf
	}

	if obj.TransactionOutputOneOf1 != nil {
		return obj.TransactionOutputOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableTransactionOutput struct {
	value *TransactionOutput
	isSet bool
}

func (v NullableTransactionOutput) Get() *TransactionOutput {
	return v.value
}

func (v *NullableTransactionOutput) Set(val *TransactionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOutput(val *TransactionOutput) *NullableTransactionOutput {
	return &NullableTransactionOutput{value: val, isSet: true}
}

func (v NullableTransactionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


