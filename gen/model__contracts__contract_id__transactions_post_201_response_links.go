/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContractsContractIdTransactionsPost201ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractsContractIdTransactionsPost201ResponseLinks{}

// ContractsContractIdTransactionsPost201ResponseLinks struct for ContractsContractIdTransactionsPost201ResponseLinks
type ContractsContractIdTransactionsPost201ResponseLinks struct {
	Transaction *string `json:"transaction,omitempty"`
}

// NewContractsContractIdTransactionsPost201ResponseLinks instantiates a new ContractsContractIdTransactionsPost201ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractsContractIdTransactionsPost201ResponseLinks() *ContractsContractIdTransactionsPost201ResponseLinks {
	this := ContractsContractIdTransactionsPost201ResponseLinks{}
	return &this
}

// NewContractsContractIdTransactionsPost201ResponseLinksWithDefaults instantiates a new ContractsContractIdTransactionsPost201ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractsContractIdTransactionsPost201ResponseLinksWithDefaults() *ContractsContractIdTransactionsPost201ResponseLinks {
	this := ContractsContractIdTransactionsPost201ResponseLinks{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *ContractsContractIdTransactionsPost201ResponseLinks) GetTransaction() string {
	if o == nil || IsNil(o.Transaction) {
		var ret string
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractsContractIdTransactionsPost201ResponseLinks) GetTransactionOk() (*string, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *ContractsContractIdTransactionsPost201ResponseLinks) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given string and assigns it to the Transaction field.
func (o *ContractsContractIdTransactionsPost201ResponseLinks) SetTransaction(v string) {
	o.Transaction = &v
}

func (o ContractsContractIdTransactionsPost201ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractsContractIdTransactionsPost201ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	return toSerialize, nil
}

type NullableContractsContractIdTransactionsPost201ResponseLinks struct {
	value *ContractsContractIdTransactionsPost201ResponseLinks
	isSet bool
}

func (v NullableContractsContractIdTransactionsPost201ResponseLinks) Get() *ContractsContractIdTransactionsPost201ResponseLinks {
	return v.value
}

func (v *NullableContractsContractIdTransactionsPost201ResponseLinks) Set(val *ContractsContractIdTransactionsPost201ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableContractsContractIdTransactionsPost201ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableContractsContractIdTransactionsPost201ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractsContractIdTransactionsPost201ResponseLinks(val *ContractsContractIdTransactionsPost201ResponseLinks) *NullableContractsContractIdTransactionsPost201ResponseLinks {
	return &NullableContractsContractIdTransactionsPost201ResponseLinks{value: val, isSet: true}
}

func (v NullableContractsContractIdTransactionsPost201ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractsContractIdTransactionsPost201ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

