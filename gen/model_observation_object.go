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

// ObservationObject - A time-varying expression that evaluates to an integer
type ObservationObject struct {
	ActionObjectOneOf *ActionObjectOneOf
	ObservationObjectOneOf *ObservationObjectOneOf
	ObservationObjectOneOf1 *ObservationObjectOneOf1
	ObservationObjectOneOf2 *ObservationObjectOneOf2
	ObservationObjectOneOf3 *ObservationObjectOneOf3
	ObservationObjectOneOf4 *ObservationObjectOneOf4
	ObservationObjectOneOf5 *ObservationObjectOneOf5
	ObservationObjectOneOf6 *ObservationObjectOneOf6
	ObservationObjectOneOf7 *ObservationObjectOneOf7
	ObservationObjectOneOf8 *ObservationObjectOneOf8
	Bool *bool
}

// ActionObjectOneOfAsObservationObject is a convenience function that returns ActionObjectOneOf wrapped in ObservationObject
func ActionObjectOneOfAsObservationObject(v *ActionObjectOneOf) ObservationObject {
	return ObservationObject{
		ActionObjectOneOf: v,
	}
}

// ObservationObjectOneOfAsObservationObject is a convenience function that returns ObservationObjectOneOf wrapped in ObservationObject
func ObservationObjectOneOfAsObservationObject(v *ObservationObjectOneOf) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf: v,
	}
}

// ObservationObjectOneOf1AsObservationObject is a convenience function that returns ObservationObjectOneOf1 wrapped in ObservationObject
func ObservationObjectOneOf1AsObservationObject(v *ObservationObjectOneOf1) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf1: v,
	}
}

// ObservationObjectOneOf2AsObservationObject is a convenience function that returns ObservationObjectOneOf2 wrapped in ObservationObject
func ObservationObjectOneOf2AsObservationObject(v *ObservationObjectOneOf2) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf2: v,
	}
}

// ObservationObjectOneOf3AsObservationObject is a convenience function that returns ObservationObjectOneOf3 wrapped in ObservationObject
func ObservationObjectOneOf3AsObservationObject(v *ObservationObjectOneOf3) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf3: v,
	}
}

// ObservationObjectOneOf4AsObservationObject is a convenience function that returns ObservationObjectOneOf4 wrapped in ObservationObject
func ObservationObjectOneOf4AsObservationObject(v *ObservationObjectOneOf4) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf4: v,
	}
}

// ObservationObjectOneOf5AsObservationObject is a convenience function that returns ObservationObjectOneOf5 wrapped in ObservationObject
func ObservationObjectOneOf5AsObservationObject(v *ObservationObjectOneOf5) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf5: v,
	}
}

// ObservationObjectOneOf6AsObservationObject is a convenience function that returns ObservationObjectOneOf6 wrapped in ObservationObject
func ObservationObjectOneOf6AsObservationObject(v *ObservationObjectOneOf6) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf6: v,
	}
}

// ObservationObjectOneOf7AsObservationObject is a convenience function that returns ObservationObjectOneOf7 wrapped in ObservationObject
func ObservationObjectOneOf7AsObservationObject(v *ObservationObjectOneOf7) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf7: v,
	}
}

// ObservationObjectOneOf8AsObservationObject is a convenience function that returns ObservationObjectOneOf8 wrapped in ObservationObject
func ObservationObjectOneOf8AsObservationObject(v *ObservationObjectOneOf8) ObservationObject {
	return ObservationObject{
		ObservationObjectOneOf8: v,
	}
}

// boolAsObservationObject is a convenience function that returns bool wrapped in ObservationObject
func BoolAsObservationObject(v *bool) ObservationObject {
	return ObservationObject{
		Bool: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ObservationObject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionObjectOneOf
	err = newStrictDecoder(data).Decode(&dst.ActionObjectOneOf)
	if err == nil {
		jsonActionObjectOneOf, _ := json.Marshal(dst.ActionObjectOneOf)
		if string(jsonActionObjectOneOf) == "{}" { // empty struct
			dst.ActionObjectOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ActionObjectOneOf = nil
	}

	// try to unmarshal data into ObservationObjectOneOf
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf)
	if err == nil {
		jsonObservationObjectOneOf, _ := json.Marshal(dst.ObservationObjectOneOf)
		if string(jsonObservationObjectOneOf) == "{}" { // empty struct
			dst.ObservationObjectOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf = nil
	}

	// try to unmarshal data into ObservationObjectOneOf1
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf1)
	if err == nil {
		jsonObservationObjectOneOf1, _ := json.Marshal(dst.ObservationObjectOneOf1)
		if string(jsonObservationObjectOneOf1) == "{}" { // empty struct
			dst.ObservationObjectOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf1 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf2
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf2)
	if err == nil {
		jsonObservationObjectOneOf2, _ := json.Marshal(dst.ObservationObjectOneOf2)
		if string(jsonObservationObjectOneOf2) == "{}" { // empty struct
			dst.ObservationObjectOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf2 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf3
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf3)
	if err == nil {
		jsonObservationObjectOneOf3, _ := json.Marshal(dst.ObservationObjectOneOf3)
		if string(jsonObservationObjectOneOf3) == "{}" { // empty struct
			dst.ObservationObjectOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf3 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf4
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf4)
	if err == nil {
		jsonObservationObjectOneOf4, _ := json.Marshal(dst.ObservationObjectOneOf4)
		if string(jsonObservationObjectOneOf4) == "{}" { // empty struct
			dst.ObservationObjectOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf4 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf5
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf5)
	if err == nil {
		jsonObservationObjectOneOf5, _ := json.Marshal(dst.ObservationObjectOneOf5)
		if string(jsonObservationObjectOneOf5) == "{}" { // empty struct
			dst.ObservationObjectOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf5 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf6
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf6)
	if err == nil {
		jsonObservationObjectOneOf6, _ := json.Marshal(dst.ObservationObjectOneOf6)
		if string(jsonObservationObjectOneOf6) == "{}" { // empty struct
			dst.ObservationObjectOneOf6 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf6 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf7
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf7)
	if err == nil {
		jsonObservationObjectOneOf7, _ := json.Marshal(dst.ObservationObjectOneOf7)
		if string(jsonObservationObjectOneOf7) == "{}" { // empty struct
			dst.ObservationObjectOneOf7 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf7 = nil
	}

	// try to unmarshal data into ObservationObjectOneOf8
	err = newStrictDecoder(data).Decode(&dst.ObservationObjectOneOf8)
	if err == nil {
		jsonObservationObjectOneOf8, _ := json.Marshal(dst.ObservationObjectOneOf8)
		if string(jsonObservationObjectOneOf8) == "{}" { // empty struct
			dst.ObservationObjectOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationObjectOneOf8 = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionObjectOneOf = nil
		dst.ObservationObjectOneOf = nil
		dst.ObservationObjectOneOf1 = nil
		dst.ObservationObjectOneOf2 = nil
		dst.ObservationObjectOneOf3 = nil
		dst.ObservationObjectOneOf4 = nil
		dst.ObservationObjectOneOf5 = nil
		dst.ObservationObjectOneOf6 = nil
		dst.ObservationObjectOneOf7 = nil
		dst.ObservationObjectOneOf8 = nil
		dst.Bool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ObservationObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ObservationObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ObservationObject) MarshalJSON() ([]byte, error) {
	if src.ActionObjectOneOf != nil {
		return json.Marshal(&src.ActionObjectOneOf)
	}

	if src.ObservationObjectOneOf != nil {
		return json.Marshal(&src.ObservationObjectOneOf)
	}

	if src.ObservationObjectOneOf1 != nil {
		return json.Marshal(&src.ObservationObjectOneOf1)
	}

	if src.ObservationObjectOneOf2 != nil {
		return json.Marshal(&src.ObservationObjectOneOf2)
	}

	if src.ObservationObjectOneOf3 != nil {
		return json.Marshal(&src.ObservationObjectOneOf3)
	}

	if src.ObservationObjectOneOf4 != nil {
		return json.Marshal(&src.ObservationObjectOneOf4)
	}

	if src.ObservationObjectOneOf5 != nil {
		return json.Marshal(&src.ObservationObjectOneOf5)
	}

	if src.ObservationObjectOneOf6 != nil {
		return json.Marshal(&src.ObservationObjectOneOf6)
	}

	if src.ObservationObjectOneOf7 != nil {
		return json.Marshal(&src.ObservationObjectOneOf7)
	}

	if src.ObservationObjectOneOf8 != nil {
		return json.Marshal(&src.ObservationObjectOneOf8)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ObservationObject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActionObjectOneOf != nil {
		return obj.ActionObjectOneOf
	}

	if obj.ObservationObjectOneOf != nil {
		return obj.ObservationObjectOneOf
	}

	if obj.ObservationObjectOneOf1 != nil {
		return obj.ObservationObjectOneOf1
	}

	if obj.ObservationObjectOneOf2 != nil {
		return obj.ObservationObjectOneOf2
	}

	if obj.ObservationObjectOneOf3 != nil {
		return obj.ObservationObjectOneOf3
	}

	if obj.ObservationObjectOneOf4 != nil {
		return obj.ObservationObjectOneOf4
	}

	if obj.ObservationObjectOneOf5 != nil {
		return obj.ObservationObjectOneOf5
	}

	if obj.ObservationObjectOneOf6 != nil {
		return obj.ObservationObjectOneOf6
	}

	if obj.ObservationObjectOneOf7 != nil {
		return obj.ObservationObjectOneOf7
	}

	if obj.ObservationObjectOneOf8 != nil {
		return obj.ObservationObjectOneOf8
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableObservationObject struct {
	value *ObservationObject
	isSet bool
}

func (v NullableObservationObject) Get() *ObservationObject {
	return v.value
}

func (v *NullableObservationObject) Set(val *ObservationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationObject(val *ObservationObject) *NullableObservationObject {
	return &NullableObservationObject{value: val, isSet: true}
}

func (v NullableObservationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


