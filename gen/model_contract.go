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

// Contract - Contract terms specified in Marlowe
type Contract struct {
	ContractOneOf *ContractOneOf
	ContractOneOf1 *ContractOneOf1
	ContractOneOf2 *ContractOneOf2
	ContractOneOf3 *ContractOneOf3
	ContractOneOf4 *ContractOneOf4
	String *string
}

// ContractOneOfAsContract is a convenience function that returns ContractOneOf wrapped in Contract
func ContractOneOfAsContract(v *ContractOneOf) Contract {
	return Contract{
		ContractOneOf: v,
	}
}

// ContractOneOf1AsContract is a convenience function that returns ContractOneOf1 wrapped in Contract
func ContractOneOf1AsContract(v *ContractOneOf1) Contract {
	return Contract{
		ContractOneOf1: v,
	}
}

// ContractOneOf2AsContract is a convenience function that returns ContractOneOf2 wrapped in Contract
func ContractOneOf2AsContract(v *ContractOneOf2) Contract {
	return Contract{
		ContractOneOf2: v,
	}
}

// ContractOneOf3AsContract is a convenience function that returns ContractOneOf3 wrapped in Contract
func ContractOneOf3AsContract(v *ContractOneOf3) Contract {
	return Contract{
		ContractOneOf3: v,
	}
}

// ContractOneOf4AsContract is a convenience function that returns ContractOneOf4 wrapped in Contract
func ContractOneOf4AsContract(v *ContractOneOf4) Contract {
	return Contract{
		ContractOneOf4: v,
	}
}

// stringAsContract is a convenience function that returns string wrapped in Contract
func StringAsContract(v *string) Contract {
	return Contract{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Contract) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContractOneOf
	err = newStrictDecoder(data).Decode(&dst.ContractOneOf)
	if err == nil {
		jsonContractOneOf, _ := json.Marshal(dst.ContractOneOf)
		if string(jsonContractOneOf) == "{}" { // empty struct
			dst.ContractOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ContractOneOf = nil
	}

	// try to unmarshal data into ContractOneOf1
	err = newStrictDecoder(data).Decode(&dst.ContractOneOf1)
	if err == nil {
		jsonContractOneOf1, _ := json.Marshal(dst.ContractOneOf1)
		if string(jsonContractOneOf1) == "{}" { // empty struct
			dst.ContractOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ContractOneOf1 = nil
	}

	// try to unmarshal data into ContractOneOf2
	err = newStrictDecoder(data).Decode(&dst.ContractOneOf2)
	if err == nil {
		jsonContractOneOf2, _ := json.Marshal(dst.ContractOneOf2)
		if string(jsonContractOneOf2) == "{}" { // empty struct
			dst.ContractOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.ContractOneOf2 = nil
	}

	// try to unmarshal data into ContractOneOf3
	err = newStrictDecoder(data).Decode(&dst.ContractOneOf3)
	if err == nil {
		jsonContractOneOf3, _ := json.Marshal(dst.ContractOneOf3)
		if string(jsonContractOneOf3) == "{}" { // empty struct
			dst.ContractOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.ContractOneOf3 = nil
	}

	// try to unmarshal data into ContractOneOf4
	err = newStrictDecoder(data).Decode(&dst.ContractOneOf4)
	if err == nil {
		jsonContractOneOf4, _ := json.Marshal(dst.ContractOneOf4)
		if string(jsonContractOneOf4) == "{}" { // empty struct
			dst.ContractOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.ContractOneOf4 = nil
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
		dst.ContractOneOf = nil
		dst.ContractOneOf1 = nil
		dst.ContractOneOf2 = nil
		dst.ContractOneOf3 = nil
		dst.ContractOneOf4 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Contract)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Contract)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Contract) MarshalJSON() ([]byte, error) {
	if src.ContractOneOf != nil {
		return json.Marshal(&src.ContractOneOf)
	}

	if src.ContractOneOf1 != nil {
		return json.Marshal(&src.ContractOneOf1)
	}

	if src.ContractOneOf2 != nil {
		return json.Marshal(&src.ContractOneOf2)
	}

	if src.ContractOneOf3 != nil {
		return json.Marshal(&src.ContractOneOf3)
	}

	if src.ContractOneOf4 != nil {
		return json.Marshal(&src.ContractOneOf4)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Contract) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContractOneOf != nil {
		return obj.ContractOneOf
	}

	if obj.ContractOneOf1 != nil {
		return obj.ContractOneOf1
	}

	if obj.ContractOneOf2 != nil {
		return obj.ContractOneOf2
	}

	if obj.ContractOneOf3 != nil {
		return obj.ContractOneOf3
	}

	if obj.ContractOneOf4 != nil {
		return obj.ContractOneOf4
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


