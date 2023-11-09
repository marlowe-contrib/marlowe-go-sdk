/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransactionOutputOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionOutputOneOf{}

// TransactionOutputOneOf Marlowe transaction output information.
type TransactionOutputOneOf struct {
	Contract Contract `json:"contract"`
	Payments []Payment `json:"payments"`
	State MarloweState `json:"state"`
	Warnings []TransactionWarning `json:"warnings"`
}

// NewTransactionOutputOneOf instantiates a new TransactionOutputOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionOutputOneOf(contract Contract, payments []Payment, state MarloweState, warnings []TransactionWarning) *TransactionOutputOneOf {
	this := TransactionOutputOneOf{}
	this.Contract = contract
	this.Payments = payments
	this.State = state
	this.Warnings = warnings
	return &this
}

// NewTransactionOutputOneOfWithDefaults instantiates a new TransactionOutputOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionOutputOneOfWithDefaults() *TransactionOutputOneOf {
	this := TransactionOutputOneOf{}
	return &this
}

// GetContract returns the Contract field value
func (o *TransactionOutputOneOf) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *TransactionOutputOneOf) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *TransactionOutputOneOf) SetContract(v Contract) {
	o.Contract = v
}

// GetPayments returns the Payments field value
func (o *TransactionOutputOneOf) GetPayments() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *TransactionOutputOneOf) GetPaymentsOk() ([]Payment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *TransactionOutputOneOf) SetPayments(v []Payment) {
	o.Payments = v
}

// GetState returns the State field value
func (o *TransactionOutputOneOf) GetState() MarloweState {
	if o == nil {
		var ret MarloweState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TransactionOutputOneOf) GetStateOk() (*MarloweState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TransactionOutputOneOf) SetState(v MarloweState) {
	o.State = v
}

// GetWarnings returns the Warnings field value
func (o *TransactionOutputOneOf) GetWarnings() []TransactionWarning {
	if o == nil {
		var ret []TransactionWarning
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *TransactionOutputOneOf) GetWarningsOk() ([]TransactionWarning, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *TransactionOutputOneOf) SetWarnings(v []TransactionWarning) {
	o.Warnings = v
}

func (o TransactionOutputOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionOutputOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract"] = o.Contract
	toSerialize["payments"] = o.Payments
	toSerialize["state"] = o.State
	toSerialize["warnings"] = o.Warnings
	return toSerialize, nil
}

type NullableTransactionOutputOneOf struct {
	value *TransactionOutputOneOf
	isSet bool
}

func (v NullableTransactionOutputOneOf) Get() *TransactionOutputOneOf {
	return v.value
}

func (v *NullableTransactionOutputOneOf) Set(val *TransactionOutputOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOutputOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOutputOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOutputOneOf(val *TransactionOutputOneOf) *NullableTransactionOutputOneOf {
	return &NullableTransactionOutputOneOf{value: val, isSet: true}
}

func (v NullableTransactionOutputOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOutputOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

