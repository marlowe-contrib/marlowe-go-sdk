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

// ActionObject - A contract which becomes active when an action occurs.
type ActionObject struct {
	ActionOneOf *ActionOneOf
	ActionOneOf1 *ActionOneOf1
	ActionOneOf2 *ActionOneOf2
	ValueObjectOneOf8 *ValueObjectOneOf8
}

// ActionOneOfAsActionObject is a convenience function that returns ActionOneOf wrapped in ActionObject
func ActionOneOfAsActionObject(v *ActionOneOf) ActionObject {
	return ActionObject{
		ActionOneOf: v,
	}
}

// ActionOneOf1AsActionObject is a convenience function that returns ActionOneOf1 wrapped in ActionObject
func ActionOneOf1AsActionObject(v *ActionOneOf1) ActionObject {
	return ActionObject{
		ActionOneOf1: v,
	}
}

// ActionOneOf2AsActionObject is a convenience function that returns ActionOneOf2 wrapped in ActionObject
func ActionOneOf2AsActionObject(v *ActionOneOf2) ActionObject {
	return ActionObject{
		ActionOneOf2: v,
	}
}

// ValueObjectOneOf8AsActionObject is a convenience function that returns ValueObjectOneOf8 wrapped in ActionObject
func ValueObjectOneOf8AsActionObject(v *ValueObjectOneOf8) ActionObject {
	return ActionObject{
		ValueObjectOneOf8: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ActionObject) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ValueObjectOneOf8
	err = newStrictDecoder(data).Decode(&dst.ValueObjectOneOf8)
	if err == nil {
		jsonValueObjectOneOf8, _ := json.Marshal(dst.ValueObjectOneOf8)
		if string(jsonValueObjectOneOf8) == "{}" { // empty struct
			dst.ValueObjectOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.ValueObjectOneOf8 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionOneOf = nil
		dst.ActionOneOf1 = nil
		dst.ActionOneOf2 = nil
		dst.ValueObjectOneOf8 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ActionObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ActionObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActionObject) MarshalJSON() ([]byte, error) {
	if src.ActionOneOf != nil {
		return json.Marshal(&src.ActionOneOf)
	}

	if src.ActionOneOf1 != nil {
		return json.Marshal(&src.ActionOneOf1)
	}

	if src.ActionOneOf2 != nil {
		return json.Marshal(&src.ActionOneOf2)
	}

	if src.ValueObjectOneOf8 != nil {
		return json.Marshal(&src.ValueObjectOneOf8)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActionObject) GetActualInstance() (interface{}) {
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

	if obj.ValueObjectOneOf8 != nil {
		return obj.ValueObjectOneOf8
	}

	// all schemas are nil
	return nil
}

type NullableActionObject struct {
	value *ActionObject
	isSet bool
}

func (v NullableActionObject) Get() *ActionObject {
	return v.value
}

func (v *NullableActionObject) Set(val *ActionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableActionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableActionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionObject(val *ActionObject) *NullableActionObject {
	return &NullableActionObject{value: val, isSet: true}
}

func (v NullableActionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

