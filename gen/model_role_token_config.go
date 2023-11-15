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

// RoleTokenConfig - struct for RoleTokenConfig
type RoleTokenConfig struct {
	RoleTokenConfigOneOf *RoleTokenConfigOneOf
	RoleTokenConfigOneOf1 *RoleTokenConfigOneOf1
	String *string
}

// RoleTokenConfigOneOfAsRoleTokenConfig is a convenience function that returns RoleTokenConfigOneOf wrapped in RoleTokenConfig
func RoleTokenConfigOneOfAsRoleTokenConfig(v *RoleTokenConfigOneOf) RoleTokenConfig {
	return RoleTokenConfig{
		RoleTokenConfigOneOf: v,
	}
}

// RoleTokenConfigOneOf1AsRoleTokenConfig is a convenience function that returns RoleTokenConfigOneOf1 wrapped in RoleTokenConfig
func RoleTokenConfigOneOf1AsRoleTokenConfig(v *RoleTokenConfigOneOf1) RoleTokenConfig {
	return RoleTokenConfig{
		RoleTokenConfigOneOf1: v,
	}
}

// stringAsRoleTokenConfig is a convenience function that returns string wrapped in RoleTokenConfig
func StringAsRoleTokenConfig(v *string) RoleTokenConfig {
	return RoleTokenConfig{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RoleTokenConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RoleTokenConfigOneOf
	err = newStrictDecoder(data).Decode(&dst.RoleTokenConfigOneOf)
	if err == nil {
		jsonRoleTokenConfigOneOf, _ := json.Marshal(dst.RoleTokenConfigOneOf)
		if string(jsonRoleTokenConfigOneOf) == "{}" { // empty struct
			dst.RoleTokenConfigOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RoleTokenConfigOneOf = nil
	}

	// try to unmarshal data into RoleTokenConfigOneOf1
	err = newStrictDecoder(data).Decode(&dst.RoleTokenConfigOneOf1)
	if err == nil {
		jsonRoleTokenConfigOneOf1, _ := json.Marshal(dst.RoleTokenConfigOneOf1)
		if string(jsonRoleTokenConfigOneOf1) == "{}" { // empty struct
			dst.RoleTokenConfigOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.RoleTokenConfigOneOf1 = nil
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
		dst.RoleTokenConfigOneOf = nil
		dst.RoleTokenConfigOneOf1 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RoleTokenConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RoleTokenConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RoleTokenConfig) MarshalJSON() ([]byte, error) {
	if src.RoleTokenConfigOneOf != nil {
		return json.Marshal(&src.RoleTokenConfigOneOf)
	}

	if src.RoleTokenConfigOneOf1 != nil {
		return json.Marshal(&src.RoleTokenConfigOneOf1)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RoleTokenConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RoleTokenConfigOneOf != nil {
		return obj.RoleTokenConfigOneOf
	}

	if obj.RoleTokenConfigOneOf1 != nil {
		return obj.RoleTokenConfigOneOf1
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRoleTokenConfig struct {
	value *RoleTokenConfig
	isSet bool
}

func (v NullableRoleTokenConfig) Get() *RoleTokenConfig {
	return v.value
}

func (v *NullableRoleTokenConfig) Set(val *RoleTokenConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleTokenConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleTokenConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleTokenConfig(val *RoleTokenConfig) *NullableRoleTokenConfig {
	return &NullableRoleTokenConfig{value: val, isSet: true}
}

func (v NullableRoleTokenConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleTokenConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


