/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marloweruntime

import (
	"encoding/json"
	"fmt"
)

// PartyObject - A participant in a contract
type PartyObject struct {
	LabelRef *LabelRef
	PartyAddress *PartyAddress
	PartyRoleName *PartyRoleName
}

// LabelRefAsPartyObject is a convenience function that returns LabelRef wrapped in PartyObject
func LabelRefAsPartyObject(v *LabelRef) PartyObject {
	return PartyObject{
		LabelRef: v,
	}
}

// PartyAddressAsPartyObject is a convenience function that returns PartyAddress wrapped in PartyObject
func PartyAddressAsPartyObject(v *PartyAddress) PartyObject {
	return PartyObject{
		PartyAddress: v,
	}
}

// PartyRoleNameAsPartyObject is a convenience function that returns PartyRoleName wrapped in PartyObject
func PartyRoleNameAsPartyObject(v *PartyRoleName) PartyObject {
	return PartyObject{
		PartyRoleName: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PartyObject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LabelRef
	err = newStrictDecoder(data).Decode(&dst.LabelRef)
	if err == nil {
		jsonLabelRef, _ := json.Marshal(dst.LabelRef)
		if string(jsonLabelRef) == "{}" { // empty struct
			dst.LabelRef = nil
		} else {
			match++
		}
	} else {
		dst.LabelRef = nil
	}

	// try to unmarshal data into PartyAddress
	err = newStrictDecoder(data).Decode(&dst.PartyAddress)
	if err == nil {
		jsonPartyAddress, _ := json.Marshal(dst.PartyAddress)
		if string(jsonPartyAddress) == "{}" { // empty struct
			dst.PartyAddress = nil
		} else {
			match++
		}
	} else {
		dst.PartyAddress = nil
	}

	// try to unmarshal data into PartyRoleName
	err = newStrictDecoder(data).Decode(&dst.PartyRoleName)
	if err == nil {
		jsonPartyRoleName, _ := json.Marshal(dst.PartyRoleName)
		if string(jsonPartyRoleName) == "{}" { // empty struct
			dst.PartyRoleName = nil
		} else {
			match++
		}
	} else {
		dst.PartyRoleName = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LabelRef = nil
		dst.PartyAddress = nil
		dst.PartyRoleName = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PartyObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PartyObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PartyObject) MarshalJSON() ([]byte, error) {
	if src.LabelRef != nil {
		return json.Marshal(&src.LabelRef)
	}

	if src.PartyAddress != nil {
		return json.Marshal(&src.PartyAddress)
	}

	if src.PartyRoleName != nil {
		return json.Marshal(&src.PartyRoleName)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PartyObject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LabelRef != nil {
		return obj.LabelRef
	}

	if obj.PartyAddress != nil {
		return obj.PartyAddress
	}

	if obj.PartyRoleName != nil {
		return obj.PartyRoleName
	}

	// all schemas are nil
	return nil
}

type NullablePartyObject struct {
	value *PartyObject
	isSet bool
}

func (v NullablePartyObject) Get() *PartyObject {
	return v.value
}

func (v *NullablePartyObject) Set(val *PartyObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyObject(val *PartyObject) *NullablePartyObject {
	return &NullablePartyObject{value: val, isSet: true}
}

func (v NullablePartyObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

