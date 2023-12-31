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

// checks if the TextEnvelope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextEnvelope{}

// TextEnvelope struct for TextEnvelope
type TextEnvelope struct {
	CborHex string `json:"cborHex"`
	Description string `json:"description"`
	// What type of data is encoded in the CBOR Hex. Valid values include \"Tx <era>\", \"TxBody <era>\", and \"ShelleyTxWitness <era>\" where <era> is one of \"BabbageEra\", \"ConwayEra\".
	Type string `json:"type"`
}

// NewTextEnvelope instantiates a new TextEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEnvelope(cborHex string, description string, type_ string) *TextEnvelope {
	this := TextEnvelope{}
	this.CborHex = cborHex
	this.Description = description
	this.Type = type_
	return &this
}

// NewTextEnvelopeWithDefaults instantiates a new TextEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEnvelopeWithDefaults() *TextEnvelope {
	this := TextEnvelope{}
	return &this
}

// GetCborHex returns the CborHex field value
func (o *TextEnvelope) GetCborHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CborHex
}

// GetCborHexOk returns a tuple with the CborHex field value
// and a boolean to check if the value has been set.
func (o *TextEnvelope) GetCborHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CborHex, true
}

// SetCborHex sets field value
func (o *TextEnvelope) SetCborHex(v string) {
	o.CborHex = v
}

// GetDescription returns the Description field value
func (o *TextEnvelope) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TextEnvelope) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TextEnvelope) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *TextEnvelope) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextEnvelope) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextEnvelope) SetType(v string) {
	o.Type = v
}

func (o TextEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextEnvelope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cborHex"] = o.CborHex
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTextEnvelope struct {
	value *TextEnvelope
	isSet bool
}

func (v NullableTextEnvelope) Get() *TextEnvelope {
	return v.value
}

func (v *NullableTextEnvelope) Set(val *TextEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEnvelope(val *TextEnvelope) *NullableTextEnvelope {
	return &NullableTextEnvelope{value: val, isSet: true}
}

func (v NullableTextEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


