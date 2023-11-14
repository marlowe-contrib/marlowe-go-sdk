/*
Marlowe Runtime REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PayoutsPayoutIdGet200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutsPayoutIdGet200ResponseLinks{}

// PayoutsPayoutIdGet200ResponseLinks struct for PayoutsPayoutIdGet200ResponseLinks
type PayoutsPayoutIdGet200ResponseLinks struct {
	Contract *string `json:"contract,omitempty"`
	Transaction *string `json:"transaction,omitempty"`
	Withdrawal *string `json:"withdrawal,omitempty"`
}

// NewPayoutsPayoutIdGet200ResponseLinks instantiates a new PayoutsPayoutIdGet200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutsPayoutIdGet200ResponseLinks() *PayoutsPayoutIdGet200ResponseLinks {
	this := PayoutsPayoutIdGet200ResponseLinks{}
	return &this
}

// NewPayoutsPayoutIdGet200ResponseLinksWithDefaults instantiates a new PayoutsPayoutIdGet200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutsPayoutIdGet200ResponseLinksWithDefaults() *PayoutsPayoutIdGet200ResponseLinks {
	this := PayoutsPayoutIdGet200ResponseLinks{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetContract() string {
	if o == nil || IsNil(o.Contract) {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetContractOk() (*string, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *PayoutsPayoutIdGet200ResponseLinks) SetContract(v string) {
	o.Contract = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetTransaction() string {
	if o == nil || IsNil(o.Transaction) {
		var ret string
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetTransactionOk() (*string, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given string and assigns it to the Transaction field.
func (o *PayoutsPayoutIdGet200ResponseLinks) SetTransaction(v string) {
	o.Transaction = &v
}

// GetWithdrawal returns the Withdrawal field value if set, zero value otherwise.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetWithdrawal() string {
	if o == nil || IsNil(o.Withdrawal) {
		var ret string
		return ret
	}
	return *o.Withdrawal
}

// GetWithdrawalOk returns a tuple with the Withdrawal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) GetWithdrawalOk() (*string, bool) {
	if o == nil || IsNil(o.Withdrawal) {
		return nil, false
	}
	return o.Withdrawal, true
}

// HasWithdrawal returns a boolean if a field has been set.
func (o *PayoutsPayoutIdGet200ResponseLinks) HasWithdrawal() bool {
	if o != nil && !IsNil(o.Withdrawal) {
		return true
	}

	return false
}

// SetWithdrawal gets a reference to the given string and assigns it to the Withdrawal field.
func (o *PayoutsPayoutIdGet200ResponseLinks) SetWithdrawal(v string) {
	o.Withdrawal = &v
}

func (o PayoutsPayoutIdGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutsPayoutIdGet200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.Withdrawal) {
		toSerialize["withdrawal"] = o.Withdrawal
	}
	return toSerialize, nil
}

type NullablePayoutsPayoutIdGet200ResponseLinks struct {
	value *PayoutsPayoutIdGet200ResponseLinks
	isSet bool
}

func (v NullablePayoutsPayoutIdGet200ResponseLinks) Get() *PayoutsPayoutIdGet200ResponseLinks {
	return v.value
}

func (v *NullablePayoutsPayoutIdGet200ResponseLinks) Set(val *PayoutsPayoutIdGet200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutsPayoutIdGet200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutsPayoutIdGet200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutsPayoutIdGet200ResponseLinks(val *PayoutsPayoutIdGet200ResponseLinks) *NullablePayoutsPayoutIdGet200ResponseLinks {
	return &NullablePayoutsPayoutIdGet200ResponseLinks{value: val, isSet: true}
}

func (v NullablePayoutsPayoutIdGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutsPayoutIdGet200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


