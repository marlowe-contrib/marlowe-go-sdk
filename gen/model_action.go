/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Action - A contract which becomes active when an action occurs.
type Action struct {
	ActionOneOf *ActionOneOf
	ActionOneOf1 *ActionOneOf1
	ActionOneOf2 *ActionOneOf2
}

// ActionOneOfAsAction is a convenience function that returns ActionOneOf wrapped in Action
func ActionOneOfAsAction(v *ActionOneOf) Action {
	return Action{
		ActionOneOf: v,
	}
}

// ActionOneOf1AsAction is a convenience function that returns ActionOneOf1 wrapped in Action
func ActionOneOf1AsAction(v *ActionOneOf1) Action {
	return Action{
		ActionOneOf1: v,
	}
}

// ActionOneOf2AsAction is a convenience function that returns ActionOneOf2 wrapped in Action
func ActionOneOf2AsAction(v *ActionOneOf2) Action {
	return Action{
		ActionOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Action) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionOneOf
	err = newStrictDecoder(data).Decode(&dst.ActionOneOf)
	if err == nil {
		jsonActionOneOf, _ := json.Marshal(dst.ActionOneOf)
		if string(jsonActionOneOf) == "{}" { // empty struct
			dst.ActionOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ActionOneOf = nil
	}

	// try to unmarshal data into ActionOneOf1
	err = newStrictDecoder(data).Decode(&dst.ActionOneOf1)
	if err == nil {
		jsonActionOneOf1, _ := json.Marshal(dst.ActionOneOf1)
		if string(jsonActionOneOf1) == "{}" { // empty struct
			dst.ActionOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ActionOneOf1 = nil
	}

	// try to unmarshal data into ActionOneOf2
	err = newStrictDecoder(data).Decode(&dst.ActionOneOf2)
	if err == nil {
		jsonActionOneOf2, _ := json.Marshal(dst.ActionOneOf2)
		if string(jsonActionOneOf2) == "{}" { // empty struct
			dst.ActionOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.ActionOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionOneOf = nil
		dst.ActionOneOf1 = nil
		dst.ActionOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Action)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Action)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Action) MarshalJSON() ([]byte, error) {
	if src.ActionOneOf != nil {
		return json.Marshal(&src.ActionOneOf)
	}

	if src.ActionOneOf1 != nil {
		return json.Marshal(&src.ActionOneOf1)
	}

	if src.ActionOneOf2 != nil {
		return json.Marshal(&src.ActionOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Action) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActionOneOf != nil {
		return obj.ActionOneOf
	}

	if obj.ActionOneOf1 != nil {
		return obj.ActionOneOf1
	}

	if obj.ActionOneOf2 != nil {
		return obj.ActionOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

