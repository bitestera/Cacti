/*
Hyperledger Cactus Plugin - Consortium Web Service

Manage a Cactus consortium through the APIs. Needs administrative privileges.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-consortium-manual

import (
	"encoding/json"
)

// checks if the JWSGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JWSGeneral{}

// JWSGeneral struct for JWSGeneral
type JWSGeneral struct {
	Payload string `json:"payload"`
	Signatures []JWSRecipient `json:"signatures"`
}

// NewJWSGeneral instantiates a new JWSGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWSGeneral(payload string, signatures []JWSRecipient) *JWSGeneral {
	this := JWSGeneral{}
	this.Payload = payload
	this.Signatures = signatures
	return &this
}

// NewJWSGeneralWithDefaults instantiates a new JWSGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWSGeneralWithDefaults() *JWSGeneral {
	this := JWSGeneral{}
	return &this
}

// GetPayload returns the Payload field value
func (o *JWSGeneral) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *JWSGeneral) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *JWSGeneral) SetPayload(v string) {
	o.Payload = v
}

// GetSignatures returns the Signatures field value
func (o *JWSGeneral) GetSignatures() []JWSRecipient {
	if o == nil {
		var ret []JWSRecipient
		return ret
	}

	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value
// and a boolean to check if the value has been set.
func (o *JWSGeneral) GetSignaturesOk() ([]JWSRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures, true
}

// SetSignatures sets field value
func (o *JWSGeneral) SetSignatures(v []JWSRecipient) {
	o.Signatures = v
}

func (o JWSGeneral) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JWSGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	toSerialize["signatures"] = o.Signatures
	return toSerialize, nil
}

type NullableJWSGeneral struct {
	value *JWSGeneral
	isSet bool
}

func (v NullableJWSGeneral) Get() *JWSGeneral {
	return v.value
}

func (v *NullableJWSGeneral) Set(val *JWSGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableJWSGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableJWSGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWSGeneral(val *JWSGeneral) *NullableJWSGeneral {
	return &NullableJWSGeneral{value: val, isSet: true}
}

func (v NullableJWSGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWSGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


