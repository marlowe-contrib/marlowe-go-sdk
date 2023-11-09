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

// MarloweStateChoicesInnerInner - struct for MarloweStateChoicesInnerInner
type MarloweStateChoicesInnerInner struct {
	ChoiceId *ChoiceId
	Int32 *int32
}

// ChoiceIdAsMarloweStateChoicesInnerInner is a convenience function that returns ChoiceId wrapped in MarloweStateChoicesInnerInner
func ChoiceIdAsMarloweStateChoicesInnerInner(v *ChoiceId) MarloweStateChoicesInnerInner {
	return MarloweStateChoicesInnerInner{
		ChoiceId: v,
	}
}

// int32AsMarloweStateChoicesInnerInner is a convenience function that returns int32 wrapped in MarloweStateChoicesInnerInner
func Int32AsMarloweStateChoicesInnerInner(v *int32) MarloweStateChoicesInnerInner {
	return MarloweStateChoicesInnerInner{
		Int32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarloweStateChoicesInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChoiceId
	err = newStrictDecoder(data).Decode(&dst.ChoiceId)
	if err == nil {
		jsonChoiceId, _ := json.Marshal(dst.ChoiceId)
		if string(jsonChoiceId) == "{}" { // empty struct
			dst.ChoiceId = nil
		} else {
			match++
		}
	} else {
		dst.ChoiceId = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChoiceId = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MarloweStateChoicesInnerInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarloweStateChoicesInnerInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarloweStateChoicesInnerInner) MarshalJSON() ([]byte, error) {
	if src.ChoiceId != nil {
		return json.Marshal(&src.ChoiceId)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarloweStateChoicesInnerInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChoiceId != nil {
		return obj.ChoiceId
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableMarloweStateChoicesInnerInner struct {
	value *MarloweStateChoicesInnerInner
	isSet bool
}

func (v NullableMarloweStateChoicesInnerInner) Get() *MarloweStateChoicesInnerInner {
	return v.value
}

func (v *NullableMarloweStateChoicesInnerInner) Set(val *MarloweStateChoicesInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarloweStateChoicesInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarloweStateChoicesInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarloweStateChoicesInnerInner(val *MarloweStateChoicesInnerInner) *NullableMarloweStateChoicesInnerInner {
	return &NullableMarloweStateChoicesInnerInner{value: val, isSet: true}
}

func (v NullableMarloweStateChoicesInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarloweStateChoicesInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

