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

// PlutusCredential - A Plutus credential.
type PlutusCredential struct {
	PlutusCredentialOneOf *PlutusCredentialOneOf
	PlutusCredentialOneOf1 *PlutusCredentialOneOf1
}

// PlutusCredentialOneOfAsPlutusCredential is a convenience function that returns PlutusCredentialOneOf wrapped in PlutusCredential
func PlutusCredentialOneOfAsPlutusCredential(v *PlutusCredentialOneOf) PlutusCredential {
	return PlutusCredential{
		PlutusCredentialOneOf: v,
	}
}

// PlutusCredentialOneOf1AsPlutusCredential is a convenience function that returns PlutusCredentialOneOf1 wrapped in PlutusCredential
func PlutusCredentialOneOf1AsPlutusCredential(v *PlutusCredentialOneOf1) PlutusCredential {
	return PlutusCredential{
		PlutusCredentialOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PlutusCredential) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PlutusCredentialOneOf
	err = newStrictDecoder(data).Decode(&dst.PlutusCredentialOneOf)
	if err == nil {
		jsonPlutusCredentialOneOf, _ := json.Marshal(dst.PlutusCredentialOneOf)
		if string(jsonPlutusCredentialOneOf) == "{}" { // empty struct
			dst.PlutusCredentialOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PlutusCredentialOneOf = nil
	}

	// try to unmarshal data into PlutusCredentialOneOf1
	err = newStrictDecoder(data).Decode(&dst.PlutusCredentialOneOf1)
	if err == nil {
		jsonPlutusCredentialOneOf1, _ := json.Marshal(dst.PlutusCredentialOneOf1)
		if string(jsonPlutusCredentialOneOf1) == "{}" { // empty struct
			dst.PlutusCredentialOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.PlutusCredentialOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PlutusCredentialOneOf = nil
		dst.PlutusCredentialOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PlutusCredential)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PlutusCredential)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PlutusCredential) MarshalJSON() ([]byte, error) {
	if src.PlutusCredentialOneOf != nil {
		return json.Marshal(&src.PlutusCredentialOneOf)
	}

	if src.PlutusCredentialOneOf1 != nil {
		return json.Marshal(&src.PlutusCredentialOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PlutusCredential) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PlutusCredentialOneOf != nil {
		return obj.PlutusCredentialOneOf
	}

	if obj.PlutusCredentialOneOf1 != nil {
		return obj.PlutusCredentialOneOf1
	}

	// all schemas are nil
	return nil
}

type NullablePlutusCredential struct {
	value *PlutusCredential
	isSet bool
}

func (v NullablePlutusCredential) Get() *PlutusCredential {
	return v.value
}

func (v *NullablePlutusCredential) Set(val *PlutusCredential) {
	v.value = val
	v.isSet = true
}

func (v NullablePlutusCredential) IsSet() bool {
	return v.isSet
}

func (v *NullablePlutusCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlutusCredential(val *PlutusCredential) *NullablePlutusCredential {
	return &NullablePlutusCredential{value: val, isSet: true}
}

func (v NullablePlutusCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlutusCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


