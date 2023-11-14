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

// Observation - A time-varying expression that evaluates to an integer
type Observation struct {
	ObservationOneOf *ObservationOneOf
	ObservationOneOf1 *ObservationOneOf1
	ObservationOneOf2 *ObservationOneOf2
	ObservationOneOf3 *ObservationOneOf3
	ObservationOneOf4 *ObservationOneOf4
	ObservationOneOf5 *ObservationOneOf5
	ObservationOneOf6 *ObservationOneOf6
	ObservationOneOf7 *ObservationOneOf7
	ObservationOneOf8 *ObservationOneOf8
	Bool *bool
}

// ObservationOneOfAsObservation is a convenience function that returns ObservationOneOf wrapped in Observation
func ObservationOneOfAsObservation(v *ObservationOneOf) Observation {
	return Observation{
		ObservationOneOf: v,
	}
}

// ObservationOneOf1AsObservation is a convenience function that returns ObservationOneOf1 wrapped in Observation
func ObservationOneOf1AsObservation(v *ObservationOneOf1) Observation {
	return Observation{
		ObservationOneOf1: v,
	}
}

// ObservationOneOf2AsObservation is a convenience function that returns ObservationOneOf2 wrapped in Observation
func ObservationOneOf2AsObservation(v *ObservationOneOf2) Observation {
	return Observation{
		ObservationOneOf2: v,
	}
}

// ObservationOneOf3AsObservation is a convenience function that returns ObservationOneOf3 wrapped in Observation
func ObservationOneOf3AsObservation(v *ObservationOneOf3) Observation {
	return Observation{
		ObservationOneOf3: v,
	}
}

// ObservationOneOf4AsObservation is a convenience function that returns ObservationOneOf4 wrapped in Observation
func ObservationOneOf4AsObservation(v *ObservationOneOf4) Observation {
	return Observation{
		ObservationOneOf4: v,
	}
}

// ObservationOneOf5AsObservation is a convenience function that returns ObservationOneOf5 wrapped in Observation
func ObservationOneOf5AsObservation(v *ObservationOneOf5) Observation {
	return Observation{
		ObservationOneOf5: v,
	}
}

// ObservationOneOf6AsObservation is a convenience function that returns ObservationOneOf6 wrapped in Observation
func ObservationOneOf6AsObservation(v *ObservationOneOf6) Observation {
	return Observation{
		ObservationOneOf6: v,
	}
}

// ObservationOneOf7AsObservation is a convenience function that returns ObservationOneOf7 wrapped in Observation
func ObservationOneOf7AsObservation(v *ObservationOneOf7) Observation {
	return Observation{
		ObservationOneOf7: v,
	}
}

// ObservationOneOf8AsObservation is a convenience function that returns ObservationOneOf8 wrapped in Observation
func ObservationOneOf8AsObservation(v *ObservationOneOf8) Observation {
	return Observation{
		ObservationOneOf8: v,
	}
}

// boolAsObservation is a convenience function that returns bool wrapped in Observation
func BoolAsObservation(v *bool) Observation {
	return Observation{
		Bool: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Observation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservationOneOf
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf)
	if err == nil {
		jsonObservationOneOf, _ := json.Marshal(dst.ObservationOneOf)
		if string(jsonObservationOneOf) == "{}" { // empty struct
			dst.ObservationOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf = nil
	}

	// try to unmarshal data into ObservationOneOf1
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf1)
	if err == nil {
		jsonObservationOneOf1, _ := json.Marshal(dst.ObservationOneOf1)
		if string(jsonObservationOneOf1) == "{}" { // empty struct
			dst.ObservationOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf1 = nil
	}

	// try to unmarshal data into ObservationOneOf2
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf2)
	if err == nil {
		jsonObservationOneOf2, _ := json.Marshal(dst.ObservationOneOf2)
		if string(jsonObservationOneOf2) == "{}" { // empty struct
			dst.ObservationOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf2 = nil
	}

	// try to unmarshal data into ObservationOneOf3
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf3)
	if err == nil {
		jsonObservationOneOf3, _ := json.Marshal(dst.ObservationOneOf3)
		if string(jsonObservationOneOf3) == "{}" { // empty struct
			dst.ObservationOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf3 = nil
	}

	// try to unmarshal data into ObservationOneOf4
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf4)
	if err == nil {
		jsonObservationOneOf4, _ := json.Marshal(dst.ObservationOneOf4)
		if string(jsonObservationOneOf4) == "{}" { // empty struct
			dst.ObservationOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf4 = nil
	}

	// try to unmarshal data into ObservationOneOf5
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf5)
	if err == nil {
		jsonObservationOneOf5, _ := json.Marshal(dst.ObservationOneOf5)
		if string(jsonObservationOneOf5) == "{}" { // empty struct
			dst.ObservationOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf5 = nil
	}

	// try to unmarshal data into ObservationOneOf6
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf6)
	if err == nil {
		jsonObservationOneOf6, _ := json.Marshal(dst.ObservationOneOf6)
		if string(jsonObservationOneOf6) == "{}" { // empty struct
			dst.ObservationOneOf6 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf6 = nil
	}

	// try to unmarshal data into ObservationOneOf7
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf7)
	if err == nil {
		jsonObservationOneOf7, _ := json.Marshal(dst.ObservationOneOf7)
		if string(jsonObservationOneOf7) == "{}" { // empty struct
			dst.ObservationOneOf7 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf7 = nil
	}

	// try to unmarshal data into ObservationOneOf8
	err = newStrictDecoder(data).Decode(&dst.ObservationOneOf8)
	if err == nil {
		jsonObservationOneOf8, _ := json.Marshal(dst.ObservationOneOf8)
		if string(jsonObservationOneOf8) == "{}" { // empty struct
			dst.ObservationOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.ObservationOneOf8 = nil
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
		dst.ObservationOneOf = nil
		dst.ObservationOneOf1 = nil
		dst.ObservationOneOf2 = nil
		dst.ObservationOneOf3 = nil
		dst.ObservationOneOf4 = nil
		dst.ObservationOneOf5 = nil
		dst.ObservationOneOf6 = nil
		dst.ObservationOneOf7 = nil
		dst.ObservationOneOf8 = nil
		dst.Bool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Observation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Observation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Observation) MarshalJSON() ([]byte, error) {
	if src.ObservationOneOf != nil {
		return json.Marshal(&src.ObservationOneOf)
	}

	if src.ObservationOneOf1 != nil {
		return json.Marshal(&src.ObservationOneOf1)
	}

	if src.ObservationOneOf2 != nil {
		return json.Marshal(&src.ObservationOneOf2)
	}

	if src.ObservationOneOf3 != nil {
		return json.Marshal(&src.ObservationOneOf3)
	}

	if src.ObservationOneOf4 != nil {
		return json.Marshal(&src.ObservationOneOf4)
	}

	if src.ObservationOneOf5 != nil {
		return json.Marshal(&src.ObservationOneOf5)
	}

	if src.ObservationOneOf6 != nil {
		return json.Marshal(&src.ObservationOneOf6)
	}

	if src.ObservationOneOf7 != nil {
		return json.Marshal(&src.ObservationOneOf7)
	}

	if src.ObservationOneOf8 != nil {
		return json.Marshal(&src.ObservationOneOf8)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Observation) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ObservationOneOf != nil {
		return obj.ObservationOneOf
	}

	if obj.ObservationOneOf1 != nil {
		return obj.ObservationOneOf1
	}

	if obj.ObservationOneOf2 != nil {
		return obj.ObservationOneOf2
	}

	if obj.ObservationOneOf3 != nil {
		return obj.ObservationOneOf3
	}

	if obj.ObservationOneOf4 != nil {
		return obj.ObservationOneOf4
	}

	if obj.ObservationOneOf5 != nil {
		return obj.ObservationOneOf5
	}

	if obj.ObservationOneOf6 != nil {
		return obj.ObservationOneOf6
	}

	if obj.ObservationOneOf7 != nil {
		return obj.ObservationOneOf7
	}

	if obj.ObservationOneOf8 != nil {
		return obj.ObservationOneOf8
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableObservation struct {
	value *Observation
	isSet bool
}

func (v NullableObservation) Get() *Observation {
	return v.value
}

func (v *NullableObservation) Set(val *Observation) {
	v.value = val
	v.isSet = true
}

func (v NullableObservation) IsSet() bool {
	return v.isSet
}

func (v *NullableObservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservation(val *Observation) *NullableObservation {
	return &NullableObservation{value: val, isSet: true}
}

func (v NullableObservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


