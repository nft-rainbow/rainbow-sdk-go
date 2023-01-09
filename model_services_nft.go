/*
Rainbow-API

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rainbowsdk

import (
	"encoding/json"
)

// ServicesNFT struct for ServicesNFT
type ServicesNFT struct {
	ContractAddress *string `json:"contract_address,omitempty"`
	Owner *string `json:"owner,omitempty"`
	TokenId *string `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesNFT ServicesNFT

// NewServicesNFT instantiates a new ServicesNFT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesNFT() *ServicesNFT {
	this := ServicesNFT{}
	return &this
}

// NewServicesNFTWithDefaults instantiates a new ServicesNFT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesNFTWithDefaults() *ServicesNFT {
	this := ServicesNFT{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ServicesNFT) GetContractAddress() string {
	if o == nil || isNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesNFT) GetContractAddressOk() (*string, bool) {
	if o == nil || isNil(o.ContractAddress) {
    return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ServicesNFT) HasContractAddress() bool {
	if o != nil && !isNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ServicesNFT) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServicesNFT) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesNFT) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServicesNFT) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServicesNFT) SetOwner(v string) {
	o.Owner = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ServicesNFT) GetTokenId() string {
	if o == nil || isNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesNFT) GetTokenIdOk() (*string, bool) {
	if o == nil || isNil(o.TokenId) {
    return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ServicesNFT) HasTokenId() bool {
	if o != nil && !isNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ServicesNFT) SetTokenId(v string) {
	o.TokenId = &v
}

func (o ServicesNFT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesNFT) UnmarshalJSON(bytes []byte) (err error) {
	varServicesNFT := _ServicesNFT{}

	if err = json.Unmarshal(bytes, &varServicesNFT); err == nil {
		*o = ServicesNFT(varServicesNFT)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesNFT struct {
	value *ServicesNFT
	isSet bool
}

func (v NullableServicesNFT) Get() *ServicesNFT {
	return v.value
}

func (v *NullableServicesNFT) Set(val *ServicesNFT) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesNFT) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesNFT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesNFT(val *ServicesNFT) *NullableServicesNFT {
	return &NullableServicesNFT{value: val, isSet: true}
}

func (v NullableServicesNFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesNFT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


