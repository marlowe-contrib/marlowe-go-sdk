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

// Party - A participant in a contract
type Party struct {
	PartyOneOf *PartyOneOf
	PartyOneOf1 *PartyOneOf1
}

// PartyOneOfAsParty is a convenience function that returns PartyOneOf wrapped in Party
func PartyOneOfAsParty(v *PartyOneOf) Party {
	return Party{
		PartyOneOf: v,
	}
}

// PartyOneOf1AsParty is a convenience function that returns PartyOneOf1 wrapped in Party
func PartyOneOf1AsParty(v *PartyOneOf1) Party {
	return Party{
		PartyOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Party) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PartyOneOf
	err = newStrictDecoder(data).Decode(&dst.PartyOneOf)
	if err == nil {
		jsonPartyOneOf, _ := json.Marshal(dst.PartyOneOf)
		if string(jsonPartyOneOf) == "{}" { // empty struct
			dst.PartyOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PartyOneOf = nil
	}

	// try to unmarshal data into PartyOneOf1
	err = newStrictDecoder(data).Decode(&dst.PartyOneOf1)
	if err == nil {
		jsonPartyOneOf1, _ := json.Marshal(dst.PartyOneOf1)
		if string(jsonPartyOneOf1) == "{}" { // empty struct
			dst.PartyOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.PartyOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PartyOneOf = nil
		dst.PartyOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Party)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Party)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Party) MarshalJSON() ([]byte, error) {
	if src.PartyOneOf != nil {
		return json.Marshal(&src.PartyOneOf)
	}

	if src.PartyOneOf1 != nil {
		return json.Marshal(&src.PartyOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Party) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PartyOneOf != nil {
		return obj.PartyOneOf
	}

	if obj.PartyOneOf1 != nil {
		return obj.PartyOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableParty struct {
	value *Party
	isSet bool
}

func (v NullableParty) Get() *Party {
	return v.value
}

func (v *NullableParty) Set(val *Party) {
	v.value = val
	v.isSet = true
}

func (v NullableParty) IsSet() bool {
	return v.isSet
}

func (v *NullableParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParty(val *Party) *NullableParty {
	return &NullableParty{value: val, isSet: true}
}

func (v NullableParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

