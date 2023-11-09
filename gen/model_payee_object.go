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

// PayeeObject - A recipient of a payment
type PayeeObject struct {
	PayeeObjectOneOf *PayeeObjectOneOf
	PayeeObjectOneOf1 *PayeeObjectOneOf1
}

// PayeeObjectOneOfAsPayeeObject is a convenience function that returns PayeeObjectOneOf wrapped in PayeeObject
func PayeeObjectOneOfAsPayeeObject(v *PayeeObjectOneOf) PayeeObject {
	return PayeeObject{
		PayeeObjectOneOf: v,
	}
}

// PayeeObjectOneOf1AsPayeeObject is a convenience function that returns PayeeObjectOneOf1 wrapped in PayeeObject
func PayeeObjectOneOf1AsPayeeObject(v *PayeeObjectOneOf1) PayeeObject {
	return PayeeObject{
		PayeeObjectOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PayeeObject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PayeeObjectOneOf
	err = newStrictDecoder(data).Decode(&dst.PayeeObjectOneOf)
	if err == nil {
		jsonPayeeObjectOneOf, _ := json.Marshal(dst.PayeeObjectOneOf)
		if string(jsonPayeeObjectOneOf) == "{}" { // empty struct
			dst.PayeeObjectOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PayeeObjectOneOf = nil
	}

	// try to unmarshal data into PayeeObjectOneOf1
	err = newStrictDecoder(data).Decode(&dst.PayeeObjectOneOf1)
	if err == nil {
		jsonPayeeObjectOneOf1, _ := json.Marshal(dst.PayeeObjectOneOf1)
		if string(jsonPayeeObjectOneOf1) == "{}" { // empty struct
			dst.PayeeObjectOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.PayeeObjectOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PayeeObjectOneOf = nil
		dst.PayeeObjectOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PayeeObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PayeeObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PayeeObject) MarshalJSON() ([]byte, error) {
	if src.PayeeObjectOneOf != nil {
		return json.Marshal(&src.PayeeObjectOneOf)
	}

	if src.PayeeObjectOneOf1 != nil {
		return json.Marshal(&src.PayeeObjectOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PayeeObject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PayeeObjectOneOf != nil {
		return obj.PayeeObjectOneOf
	}

	if obj.PayeeObjectOneOf1 != nil {
		return obj.PayeeObjectOneOf1
	}

	// all schemas are nil
	return nil
}

type NullablePayeeObject struct {
	value *PayeeObject
	isSet bool
}

func (v NullablePayeeObject) Get() *PayeeObject {
	return v.value
}

func (v *NullablePayeeObject) Set(val *PayeeObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePayeeObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePayeeObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayeeObject(val *PayeeObject) *NullablePayeeObject {
	return &NullablePayeeObject{value: val, isSet: true}
}

func (v NullablePayeeObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayeeObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

