/*
Marlowe Runtime REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TransactionError - A Marlowe transaction error.
type TransactionError struct {
	TransactionErrorOneOf *TransactionErrorOneOf
	String *string
}

// TransactionErrorOneOfAsTransactionError is a convenience function that returns TransactionErrorOneOf wrapped in TransactionError
func TransactionErrorOneOfAsTransactionError(v *TransactionErrorOneOf) TransactionError {
	return TransactionError{
		TransactionErrorOneOf: v,
	}
}

// stringAsTransactionError is a convenience function that returns string wrapped in TransactionError
func StringAsTransactionError(v *string) TransactionError {
	return TransactionError{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TransactionErrorOneOf
	err = newStrictDecoder(data).Decode(&dst.TransactionErrorOneOf)
	if err == nil {
		jsonTransactionErrorOneOf, _ := json.Marshal(dst.TransactionErrorOneOf)
		if string(jsonTransactionErrorOneOf) == "{}" { // empty struct
			dst.TransactionErrorOneOf = nil
		} else {
			match++
		}
	} else {
		dst.TransactionErrorOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TransactionErrorOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransactionError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransactionError)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionError) MarshalJSON() ([]byte, error) {
	if src.TransactionErrorOneOf != nil {
		return json.Marshal(&src.TransactionErrorOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionError) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionErrorOneOf != nil {
		return obj.TransactionErrorOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTransactionError struct {
	value *TransactionError
	isSet bool
}

func (v NullableTransactionError) Get() *TransactionError {
	return v.value
}

func (v *NullableTransactionError) Set(val *TransactionError) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionError) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionError(val *TransactionError) *NullableTransactionError {
	return &NullableTransactionError{value: val, isSet: true}
}

func (v NullableTransactionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


