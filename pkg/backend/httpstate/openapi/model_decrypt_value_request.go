/*
Pulumi Service API

The Pulumi Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DecryptValueRequest struct for DecryptValueRequest
type DecryptValueRequest struct {
	Ciphertext string `json:"ciphertext"`
	AdditionalProperties map[string]interface{}
}

type _DecryptValueRequest DecryptValueRequest

// NewDecryptValueRequest instantiates a new DecryptValueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptValueRequest(ciphertext string) *DecryptValueRequest {
	this := DecryptValueRequest{}
	this.Ciphertext = ciphertext
	return &this
}

// NewDecryptValueRequestWithDefaults instantiates a new DecryptValueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptValueRequestWithDefaults() *DecryptValueRequest {
	this := DecryptValueRequest{}
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *DecryptValueRequest) GetCiphertext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *DecryptValueRequest) GetCiphertextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *DecryptValueRequest) SetCiphertext(v string) {
	o.Ciphertext = v
}

func (o DecryptValueRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ciphertext"] = o.Ciphertext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DecryptValueRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDecryptValueRequest := _DecryptValueRequest{}

	if err = json.Unmarshal(bytes, &varDecryptValueRequest); err == nil {
		*o = DecryptValueRequest(varDecryptValueRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ciphertext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDecryptValueRequest struct {
	value *DecryptValueRequest
	isSet bool
}

func (v NullableDecryptValueRequest) Get() *DecryptValueRequest {
	return v.value
}

func (v *NullableDecryptValueRequest) Set(val *DecryptValueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptValueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptValueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptValueRequest(val *DecryptValueRequest) *NullableDecryptValueRequest {
	return &NullableDecryptValueRequest{value: val, isSet: true}
}

func (v NullableDecryptValueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptValueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


