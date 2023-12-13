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

// TransactionWarning - A transaction semantics warning.
type TransactionWarning struct {
	AssertFail *AssertFail
	NonPositiveDeposit *NonPositiveDeposit
	NonPositivePayment *NonPositivePayment
	PartialPayment *PartialPayment
	VariableShadowing *VariableShadowing
}

// AssertFailAsTransactionWarning is a convenience function that returns AssertFail wrapped in TransactionWarning
func AssertFailAsTransactionWarning(v *AssertFail) TransactionWarning {
	return TransactionWarning{
		AssertFail: v,
	}
}

// NonPositiveDepositAsTransactionWarning is a convenience function that returns NonPositiveDeposit wrapped in TransactionWarning
func NonPositiveDepositAsTransactionWarning(v *NonPositiveDeposit) TransactionWarning {
	return TransactionWarning{
		NonPositiveDeposit: v,
	}
}

// NonPositivePaymentAsTransactionWarning is a convenience function that returns NonPositivePayment wrapped in TransactionWarning
func NonPositivePaymentAsTransactionWarning(v *NonPositivePayment) TransactionWarning {
	return TransactionWarning{
		NonPositivePayment: v,
	}
}

// PartialPaymentAsTransactionWarning is a convenience function that returns PartialPayment wrapped in TransactionWarning
func PartialPaymentAsTransactionWarning(v *PartialPayment) TransactionWarning {
	return TransactionWarning{
		PartialPayment: v,
	}
}

// VariableShadowingAsTransactionWarning is a convenience function that returns VariableShadowing wrapped in TransactionWarning
func VariableShadowingAsTransactionWarning(v *VariableShadowing) TransactionWarning {
	return TransactionWarning{
		VariableShadowing: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionWarning) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AssertFail
	err = newStrictDecoder(data).Decode(&dst.AssertFail)
	if err == nil {
		jsonAssertFail, _ := json.Marshal(dst.AssertFail)
		if string(jsonAssertFail) == "{}" { // empty struct
			dst.AssertFail = nil
		} else {
			match++
		}
	} else {
		dst.AssertFail = nil
	}

	// try to unmarshal data into NonPositiveDeposit
	err = newStrictDecoder(data).Decode(&dst.NonPositiveDeposit)
	if err == nil {
		jsonNonPositiveDeposit, _ := json.Marshal(dst.NonPositiveDeposit)
		if string(jsonNonPositiveDeposit) == "{}" { // empty struct
			dst.NonPositiveDeposit = nil
		} else {
			match++
		}
	} else {
		dst.NonPositiveDeposit = nil
	}

	// try to unmarshal data into NonPositivePayment
	err = newStrictDecoder(data).Decode(&dst.NonPositivePayment)
	if err == nil {
		jsonNonPositivePayment, _ := json.Marshal(dst.NonPositivePayment)
		if string(jsonNonPositivePayment) == "{}" { // empty struct
			dst.NonPositivePayment = nil
		} else {
			match++
		}
	} else {
		dst.NonPositivePayment = nil
	}

	// try to unmarshal data into PartialPayment
	err = newStrictDecoder(data).Decode(&dst.PartialPayment)
	if err == nil {
		jsonPartialPayment, _ := json.Marshal(dst.PartialPayment)
		if string(jsonPartialPayment) == "{}" { // empty struct
			dst.PartialPayment = nil
		} else {
			match++
		}
	} else {
		dst.PartialPayment = nil
	}

	// try to unmarshal data into VariableShadowing
	err = newStrictDecoder(data).Decode(&dst.VariableShadowing)
	if err == nil {
		jsonVariableShadowing, _ := json.Marshal(dst.VariableShadowing)
		if string(jsonVariableShadowing) == "{}" { // empty struct
			dst.VariableShadowing = nil
		} else {
			match++
		}
	} else {
		dst.VariableShadowing = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssertFail = nil
		dst.NonPositiveDeposit = nil
		dst.NonPositivePayment = nil
		dst.PartialPayment = nil
		dst.VariableShadowing = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransactionWarning)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransactionWarning)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionWarning) MarshalJSON() ([]byte, error) {
	if src.AssertFail != nil {
		return json.Marshal(&src.AssertFail)
	}

	if src.NonPositiveDeposit != nil {
		return json.Marshal(&src.NonPositiveDeposit)
	}

	if src.NonPositivePayment != nil {
		return json.Marshal(&src.NonPositivePayment)
	}

	if src.PartialPayment != nil {
		return json.Marshal(&src.PartialPayment)
	}

	if src.VariableShadowing != nil {
		return json.Marshal(&src.VariableShadowing)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionWarning) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AssertFail != nil {
		return obj.AssertFail
	}

	if obj.NonPositiveDeposit != nil {
		return obj.NonPositiveDeposit
	}

	if obj.NonPositivePayment != nil {
		return obj.NonPositivePayment
	}

	if obj.PartialPayment != nil {
		return obj.PartialPayment
	}

	if obj.VariableShadowing != nil {
		return obj.VariableShadowing
	}

	// all schemas are nil
	return nil
}

type NullableTransactionWarning struct {
	value *TransactionWarning
	isSet bool
}

func (v NullableTransactionWarning) Get() *TransactionWarning {
	return v.value
}

func (v *NullableTransactionWarning) Set(val *TransactionWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionWarning(val *TransactionWarning) *NullableTransactionWarning {
	return &NullableTransactionWarning{value: val, isSet: true}
}

func (v NullableTransactionWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


