/*
Rainbow-API

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rainbowsdk

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutersUpdateTokenUriReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutersUpdateTokenUriReq{}

// RoutersUpdateTokenUriReq struct for RoutersUpdateTokenUriReq
type RoutersUpdateTokenUriReq struct {
	Chain string `json:"chain"`
	ContractType string `json:"contract_type"`
	TokenUri *string `json:"token_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutersUpdateTokenUriReq RoutersUpdateTokenUriReq

// NewRoutersUpdateTokenUriReq instantiates a new RoutersUpdateTokenUriReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutersUpdateTokenUriReq(chain string, contractType string) *RoutersUpdateTokenUriReq {
	this := RoutersUpdateTokenUriReq{}
	this.Chain = chain
	this.ContractType = contractType
	return &this
}

// NewRoutersUpdateTokenUriReqWithDefaults instantiates a new RoutersUpdateTokenUriReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutersUpdateTokenUriReqWithDefaults() *RoutersUpdateTokenUriReq {
	this := RoutersUpdateTokenUriReq{}
	return &this
}

// GetChain returns the Chain field value
func (o *RoutersUpdateTokenUriReq) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *RoutersUpdateTokenUriReq) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *RoutersUpdateTokenUriReq) SetChain(v string) {
	o.Chain = v
}

// GetContractType returns the ContractType field value
func (o *RoutersUpdateTokenUriReq) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *RoutersUpdateTokenUriReq) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *RoutersUpdateTokenUriReq) SetContractType(v string) {
	o.ContractType = v
}

// GetTokenUri returns the TokenUri field value if set, zero value otherwise.
func (o *RoutersUpdateTokenUriReq) GetTokenUri() string {
	if o == nil || IsNil(o.TokenUri) {
		var ret string
		return ret
	}
	return *o.TokenUri
}

// GetTokenUriOk returns a tuple with the TokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutersUpdateTokenUriReq) GetTokenUriOk() (*string, bool) {
	if o == nil || IsNil(o.TokenUri) {
		return nil, false
	}
	return o.TokenUri, true
}

// HasTokenUri returns a boolean if a field has been set.
func (o *RoutersUpdateTokenUriReq) HasTokenUri() bool {
	if o != nil && !IsNil(o.TokenUri) {
		return true
	}

	return false
}

// SetTokenUri gets a reference to the given string and assigns it to the TokenUri field.
func (o *RoutersUpdateTokenUriReq) SetTokenUri(v string) {
	o.TokenUri = &v
}

func (o RoutersUpdateTokenUriReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutersUpdateTokenUriReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chain"] = o.Chain
	toSerialize["contract_type"] = o.ContractType
	if !IsNil(o.TokenUri) {
		toSerialize["token_uri"] = o.TokenUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutersUpdateTokenUriReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chain",
		"contract_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoutersUpdateTokenUriReq := _RoutersUpdateTokenUriReq{}

	err = json.Unmarshal(data, &varRoutersUpdateTokenUriReq)

	if err != nil {
		return err
	}

	*o = RoutersUpdateTokenUriReq(varRoutersUpdateTokenUriReq)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "token_uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutersUpdateTokenUriReq struct {
	value *RoutersUpdateTokenUriReq
	isSet bool
}

func (v NullableRoutersUpdateTokenUriReq) Get() *RoutersUpdateTokenUriReq {
	return v.value
}

func (v *NullableRoutersUpdateTokenUriReq) Set(val *RoutersUpdateTokenUriReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutersUpdateTokenUriReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutersUpdateTokenUriReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutersUpdateTokenUriReq(val *RoutersUpdateTokenUriReq) *NullableRoutersUpdateTokenUriReq {
	return &NullableRoutersUpdateTokenUriReq{value: val, isSet: true}
}

func (v NullableRoutersUpdateTokenUriReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutersUpdateTokenUriReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


