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

// MarloweStateBoundValuesInnerInner - struct for MarloweStateBoundValuesInnerInner
type MarloweStateBoundValuesInnerInner struct {
	Int32 *int32
	String *string
}

// int32AsMarloweStateBoundValuesInnerInner is a convenience function that returns int32 wrapped in MarloweStateBoundValuesInnerInner
func Int32AsMarloweStateBoundValuesInnerInner(v *int32) MarloweStateBoundValuesInnerInner {
	return MarloweStateBoundValuesInnerInner{
		Int32: v,
	}
}

// stringAsMarloweStateBoundValuesInnerInner is a convenience function that returns string wrapped in MarloweStateBoundValuesInnerInner
func StringAsMarloweStateBoundValuesInnerInner(v *string) MarloweStateBoundValuesInnerInner {
	return MarloweStateBoundValuesInnerInner{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarloweStateBoundValuesInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MarloweStateBoundValuesInnerInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarloweStateBoundValuesInnerInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarloweStateBoundValuesInnerInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarloweStateBoundValuesInnerInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableMarloweStateBoundValuesInnerInner struct {
	value *MarloweStateBoundValuesInnerInner
	isSet bool
}

func (v NullableMarloweStateBoundValuesInnerInner) Get() *MarloweStateBoundValuesInnerInner {
	return v.value
}

func (v *NullableMarloweStateBoundValuesInnerInner) Set(val *MarloweStateBoundValuesInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarloweStateBoundValuesInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarloweStateBoundValuesInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarloweStateBoundValuesInnerInner(val *MarloweStateBoundValuesInnerInner) *NullableMarloweStateBoundValuesInnerInner {
	return &NullableMarloweStateBoundValuesInnerInner{value: val, isSet: true}
}

func (v NullableMarloweStateBoundValuesInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarloweStateBoundValuesInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

