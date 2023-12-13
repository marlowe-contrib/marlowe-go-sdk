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

// Contract - Contract terms specified in Marlowe
type Contract struct {
	Assert *Assert
	Close *Close
	If *If
	Let *Let
	Pay *Pay
	When *When
}

// AssertAsContract is a convenience function that returns Assert wrapped in Contract
func AssertAsContract(v *Assert) Contract {
	return Contract{
		Assert: v,
	}
}

// CloseAsContract is a convenience function that returns Close wrapped in Contract
func CloseAsContract(v *Close) Contract {
	return Contract{
		Close: v,
	}
}

// IfAsContract is a convenience function that returns If wrapped in Contract
func IfAsContract(v *If) Contract {
	return Contract{
		If: v,
	}
}

// LetAsContract is a convenience function that returns Let wrapped in Contract
func LetAsContract(v *Let) Contract {
	return Contract{
		Let: v,
	}
}

// PayAsContract is a convenience function that returns Pay wrapped in Contract
func PayAsContract(v *Pay) Contract {
	return Contract{
		Pay: v,
	}
}

// WhenAsContract is a convenience function that returns When wrapped in Contract
func WhenAsContract(v *When) Contract {
	return Contract{
		When: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Contract) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Assert
	err = newStrictDecoder(data).Decode(&dst.Assert)
	if err == nil {
		jsonAssert, _ := json.Marshal(dst.Assert)
		if string(jsonAssert) == "{}" { // empty struct
			dst.Assert = nil
		} else {
			match++
		}
	} else {
		dst.Assert = nil
	}

	// try to unmarshal data into Close
	err = newStrictDecoder(data).Decode(&dst.Close)
	if err == nil {
		jsonClose, _ := json.Marshal(dst.Close)
		if string(jsonClose) == "{}" { // empty struct
			dst.Close = nil
		} else {
			match++
		}
	} else {
		dst.Close = nil
	}

	// try to unmarshal data into If
	err = newStrictDecoder(data).Decode(&dst.If)
	if err == nil {
		jsonIf, _ := json.Marshal(dst.If)
		if string(jsonIf) == "{}" { // empty struct
			dst.If = nil
		} else {
			match++
		}
	} else {
		dst.If = nil
	}

	// try to unmarshal data into Let
	err = newStrictDecoder(data).Decode(&dst.Let)
	if err == nil {
		jsonLet, _ := json.Marshal(dst.Let)
		if string(jsonLet) == "{}" { // empty struct
			dst.Let = nil
		} else {
			match++
		}
	} else {
		dst.Let = nil
	}

	// try to unmarshal data into Pay
	err = newStrictDecoder(data).Decode(&dst.Pay)
	if err == nil {
		jsonPay, _ := json.Marshal(dst.Pay)
		if string(jsonPay) == "{}" { // empty struct
			dst.Pay = nil
		} else {
			match++
		}
	} else {
		dst.Pay = nil
	}

	// try to unmarshal data into When
	err = newStrictDecoder(data).Decode(&dst.When)
	if err == nil {
		jsonWhen, _ := json.Marshal(dst.When)
		if string(jsonWhen) == "{}" { // empty struct
			dst.When = nil
		} else {
			match++
		}
	} else {
		dst.When = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Assert = nil
		dst.Close = nil
		dst.If = nil
		dst.Let = nil
		dst.Pay = nil
		dst.When = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Contract)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Contract)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Contract) MarshalJSON() ([]byte, error) {
	if src.Assert != nil {
		return json.Marshal(&src.Assert)
	}

	if src.Close != nil {
		return json.Marshal(&src.Close)
	}

	if src.If != nil {
		return json.Marshal(&src.If)
	}

	if src.Let != nil {
		return json.Marshal(&src.Let)
	}

	if src.Pay != nil {
		return json.Marshal(&src.Pay)
	}

	if src.When != nil {
		return json.Marshal(&src.When)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Contract) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Assert != nil {
		return obj.Assert
	}

	if obj.Close != nil {
		return obj.Close
	}

	if obj.If != nil {
		return obj.If
	}

	if obj.Let != nil {
		return obj.Let
	}

	if obj.Pay != nil {
		return obj.Pay
	}

	if obj.When != nil {
		return obj.When
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


