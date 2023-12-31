/*
Marlowe Runtime REST API

REST API for Marlowe Runtime

API version: 0.0.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marloweruntime

import (
	"encoding/json"
)

// checks if the GetPayoutResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPayoutResponseLinks{}

// GetPayoutResponseLinks struct for GetPayoutResponseLinks
type GetPayoutResponseLinks struct {
	Contract *string `json:"contract,omitempty"`
	Transaction *string `json:"transaction,omitempty"`
	Withdrawal *string `json:"withdrawal,omitempty"`
}

// NewGetPayoutResponseLinks instantiates a new GetPayoutResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayoutResponseLinks() *GetPayoutResponseLinks {
	this := GetPayoutResponseLinks{}
	return &this
}

// NewGetPayoutResponseLinksWithDefaults instantiates a new GetPayoutResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayoutResponseLinksWithDefaults() *GetPayoutResponseLinks {
	this := GetPayoutResponseLinks{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *GetPayoutResponseLinks) GetContract() string {
	if o == nil || IsNil(o.Contract) {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayoutResponseLinks) GetContractOk() (*string, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *GetPayoutResponseLinks) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *GetPayoutResponseLinks) SetContract(v string) {
	o.Contract = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *GetPayoutResponseLinks) GetTransaction() string {
	if o == nil || IsNil(o.Transaction) {
		var ret string
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayoutResponseLinks) GetTransactionOk() (*string, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *GetPayoutResponseLinks) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given string and assigns it to the Transaction field.
func (o *GetPayoutResponseLinks) SetTransaction(v string) {
	o.Transaction = &v
}

// GetWithdrawal returns the Withdrawal field value if set, zero value otherwise.
func (o *GetPayoutResponseLinks) GetWithdrawal() string {
	if o == nil || IsNil(o.Withdrawal) {
		var ret string
		return ret
	}
	return *o.Withdrawal
}

// GetWithdrawalOk returns a tuple with the Withdrawal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayoutResponseLinks) GetWithdrawalOk() (*string, bool) {
	if o == nil || IsNil(o.Withdrawal) {
		return nil, false
	}
	return o.Withdrawal, true
}

// HasWithdrawal returns a boolean if a field has been set.
func (o *GetPayoutResponseLinks) HasWithdrawal() bool {
	if o != nil && !IsNil(o.Withdrawal) {
		return true
	}

	return false
}

// SetWithdrawal gets a reference to the given string and assigns it to the Withdrawal field.
func (o *GetPayoutResponseLinks) SetWithdrawal(v string) {
	o.Withdrawal = &v
}

func (o GetPayoutResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayoutResponseLinks) ToMap() (map[string]interface{}, error) {
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

type NullableGetPayoutResponseLinks struct {
	value *GetPayoutResponseLinks
	isSet bool
}

func (v NullableGetPayoutResponseLinks) Get() *GetPayoutResponseLinks {
	return v.value
}

func (v *NullableGetPayoutResponseLinks) Set(val *GetPayoutResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayoutResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayoutResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayoutResponseLinks(val *GetPayoutResponseLinks) *NullableGetPayoutResponseLinks {
	return &NullableGetPayoutResponseLinks{value: val, isSet: true}
}

func (v NullableGetPayoutResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayoutResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


