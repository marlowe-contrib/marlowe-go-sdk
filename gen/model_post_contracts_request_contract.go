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

// PostContractsRequestContract - struct for PostContractsRequestContract
type PostContractsRequestContract struct {
	Contract *Contract
	String *string
}

// ContractAsPostContractsRequestContract is a convenience function that returns Contract wrapped in PostContractsRequestContract
func ContractAsPostContractsRequestContract(v *Contract) PostContractsRequestContract {
	return PostContractsRequestContract{
		Contract: v,
	}
}

// stringAsPostContractsRequestContract is a convenience function that returns string wrapped in PostContractsRequestContract
func StringAsPostContractsRequestContract(v *string) PostContractsRequestContract {
	return PostContractsRequestContract{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostContractsRequestContract) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Contract
	err = newStrictDecoder(data).Decode(&dst.Contract)
	if err == nil {
		jsonContract, _ := json.Marshal(dst.Contract)
		if string(jsonContract) == "{}" { // empty struct
			dst.Contract = nil
		} else {
			match++
		}
	} else {
		dst.Contract = nil
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
		dst.Contract = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostContractsRequestContract)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostContractsRequestContract)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostContractsRequestContract) MarshalJSON() ([]byte, error) {
	if src.Contract != nil {
		return json.Marshal(&src.Contract)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostContractsRequestContract) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Contract != nil {
		return obj.Contract
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePostContractsRequestContract struct {
	value *PostContractsRequestContract
	isSet bool
}

func (v NullablePostContractsRequestContract) Get() *PostContractsRequestContract {
	return v.value
}

func (v *NullablePostContractsRequestContract) Set(val *PostContractsRequestContract) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContractsRequestContract) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContractsRequestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContractsRequestContract(val *PostContractsRequestContract) *NullablePostContractsRequestContract {
	return &NullablePostContractsRequestContract{value: val, isSet: true}
}

func (v NullablePostContractsRequestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostContractsRequestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


