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

// MarloweStateAccountsInnerInnerOneOfInner - struct for MarloweStateAccountsInnerInnerOneOfInner
type MarloweStateAccountsInnerInnerOneOfInner struct {
	Party *Party
	Token *Token
}

// PartyAsMarloweStateAccountsInnerInnerOneOfInner is a convenience function that returns Party wrapped in MarloweStateAccountsInnerInnerOneOfInner
func PartyAsMarloweStateAccountsInnerInnerOneOfInner(v *Party) MarloweStateAccountsInnerInnerOneOfInner {
	return MarloweStateAccountsInnerInnerOneOfInner{
		Party: v,
	}
}

// TokenAsMarloweStateAccountsInnerInnerOneOfInner is a convenience function that returns Token wrapped in MarloweStateAccountsInnerInnerOneOfInner
func TokenAsMarloweStateAccountsInnerInnerOneOfInner(v *Token) MarloweStateAccountsInnerInnerOneOfInner {
	return MarloweStateAccountsInnerInnerOneOfInner{
		Token: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarloweStateAccountsInnerInnerOneOfInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Party
	err = newStrictDecoder(data).Decode(&dst.Party)
	if err == nil {
		jsonParty, _ := json.Marshal(dst.Party)
		if string(jsonParty) == "{}" { // empty struct
			dst.Party = nil
		} else {
			match++
		}
	} else {
		dst.Party = nil
	}

	// try to unmarshal data into Token
	err = newStrictDecoder(data).Decode(&dst.Token)
	if err == nil {
		jsonToken, _ := json.Marshal(dst.Token)
		if string(jsonToken) == "{}" { // empty struct
			dst.Token = nil
		} else {
			match++
		}
	} else {
		dst.Token = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Party = nil
		dst.Token = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MarloweStateAccountsInnerInnerOneOfInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarloweStateAccountsInnerInnerOneOfInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarloweStateAccountsInnerInnerOneOfInner) MarshalJSON() ([]byte, error) {
	if src.Party != nil {
		return json.Marshal(&src.Party)
	}

	if src.Token != nil {
		return json.Marshal(&src.Token)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarloweStateAccountsInnerInnerOneOfInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Party != nil {
		return obj.Party
	}

	if obj.Token != nil {
		return obj.Token
	}

	// all schemas are nil
	return nil
}

type NullableMarloweStateAccountsInnerInnerOneOfInner struct {
	value *MarloweStateAccountsInnerInnerOneOfInner
	isSet bool
}

func (v NullableMarloweStateAccountsInnerInnerOneOfInner) Get() *MarloweStateAccountsInnerInnerOneOfInner {
	return v.value
}

func (v *NullableMarloweStateAccountsInnerInnerOneOfInner) Set(val *MarloweStateAccountsInnerInnerOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarloweStateAccountsInnerInnerOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarloweStateAccountsInnerInnerOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarloweStateAccountsInnerInnerOneOfInner(val *MarloweStateAccountsInnerInnerOneOfInner) *NullableMarloweStateAccountsInnerInnerOneOfInner {
	return &NullableMarloweStateAccountsInnerInnerOneOfInner{value: val, isSet: true}
}

func (v NullableMarloweStateAccountsInnerInnerOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarloweStateAccountsInnerInnerOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


