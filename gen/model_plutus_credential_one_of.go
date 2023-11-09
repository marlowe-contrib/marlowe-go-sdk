/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PlutusCredentialOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlutusCredentialOneOf{}

// PlutusCredentialOneOf A Plutus public key credential.
type PlutusCredentialOneOf struct {
	PubKeyCredential string `json:"pubKeyCredential"`
}

// NewPlutusCredentialOneOf instantiates a new PlutusCredentialOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlutusCredentialOneOf(pubKeyCredential string) *PlutusCredentialOneOf {
	this := PlutusCredentialOneOf{}
	this.PubKeyCredential = pubKeyCredential
	return &this
}

// NewPlutusCredentialOneOfWithDefaults instantiates a new PlutusCredentialOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlutusCredentialOneOfWithDefaults() *PlutusCredentialOneOf {
	this := PlutusCredentialOneOf{}
	return &this
}

// GetPubKeyCredential returns the PubKeyCredential field value
func (o *PlutusCredentialOneOf) GetPubKeyCredential() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubKeyCredential
}

// GetPubKeyCredentialOk returns a tuple with the PubKeyCredential field value
// and a boolean to check if the value has been set.
func (o *PlutusCredentialOneOf) GetPubKeyCredentialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubKeyCredential, true
}

// SetPubKeyCredential sets field value
func (o *PlutusCredentialOneOf) SetPubKeyCredential(v string) {
	o.PubKeyCredential = v
}

func (o PlutusCredentialOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlutusCredentialOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pubKeyCredential"] = o.PubKeyCredential
	return toSerialize, nil
}

type NullablePlutusCredentialOneOf struct {
	value *PlutusCredentialOneOf
	isSet bool
}

func (v NullablePlutusCredentialOneOf) Get() *PlutusCredentialOneOf {
	return v.value
}

func (v *NullablePlutusCredentialOneOf) Set(val *PlutusCredentialOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlutusCredentialOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlutusCredentialOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlutusCredentialOneOf(val *PlutusCredentialOneOf) *NullablePlutusCredentialOneOf {
	return &NullablePlutusCredentialOneOf{value: val, isSet: true}
}

func (v NullablePlutusCredentialOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlutusCredentialOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


