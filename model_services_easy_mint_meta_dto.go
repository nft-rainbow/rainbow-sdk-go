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

// ServicesEasyMintMetaDto struct for ServicesEasyMintMetaDto
type ServicesEasyMintMetaDto struct {
	Chain string `json:"chain"`
	Description string `json:"description"`
	FileUrl string `json:"file_url"`
	MintToAddress string `json:"mint_to_address"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ServicesEasyMintMetaDto ServicesEasyMintMetaDto

// NewServicesEasyMintMetaDto instantiates a new ServicesEasyMintMetaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesEasyMintMetaDto(chain string, description string, fileUrl string, mintToAddress string, name string) *ServicesEasyMintMetaDto {
	this := ServicesEasyMintMetaDto{}
	this.Chain = chain
	this.Description = description
	this.FileUrl = fileUrl
	this.MintToAddress = mintToAddress
	this.Name = name
	return &this
}

// NewServicesEasyMintMetaDtoWithDefaults instantiates a new ServicesEasyMintMetaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesEasyMintMetaDtoWithDefaults() *ServicesEasyMintMetaDto {
	this := ServicesEasyMintMetaDto{}
	return &this
}

// GetChain returns the Chain field value
func (o *ServicesEasyMintMetaDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetChainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesEasyMintMetaDto) SetChain(v string) {
	o.Chain = v
}

// GetDescription returns the Description field value
func (o *ServicesEasyMintMetaDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServicesEasyMintMetaDto) SetDescription(v string) {
	o.Description = v
}

// GetFileUrl returns the FileUrl field value
func (o *ServicesEasyMintMetaDto) GetFileUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetFileUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileUrl, true
}

// SetFileUrl sets field value
func (o *ServicesEasyMintMetaDto) SetFileUrl(v string) {
	o.FileUrl = v
}

// GetMintToAddress returns the MintToAddress field value
func (o *ServicesEasyMintMetaDto) GetMintToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MintToAddress
}

// GetMintToAddressOk returns a tuple with the MintToAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetMintToAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MintToAddress, true
}

// SetMintToAddress sets field value
func (o *ServicesEasyMintMetaDto) SetMintToAddress(v string) {
	o.MintToAddress = v
}

// GetName returns the Name field value
func (o *ServicesEasyMintMetaDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServicesEasyMintMetaDto) SetName(v string) {
	o.Name = v
}

func (o ServicesEasyMintMetaDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["file_url"] = o.FileUrl
	}
	if true {
		toSerialize["mint_to_address"] = o.MintToAddress
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesEasyMintMetaDto) UnmarshalJSON(bytes []byte) (err error) {
	varServicesEasyMintMetaDto := _ServicesEasyMintMetaDto{}

	if err = json.Unmarshal(bytes, &varServicesEasyMintMetaDto); err == nil {
		*o = ServicesEasyMintMetaDto(varServicesEasyMintMetaDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "chain")
		delete(additionalProperties, "description")
		delete(additionalProperties, "file_url")
		delete(additionalProperties, "mint_to_address")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesEasyMintMetaDto struct {
	value *ServicesEasyMintMetaDto
	isSet bool
}

func (v NullableServicesEasyMintMetaDto) Get() *ServicesEasyMintMetaDto {
	return v.value
}

func (v *NullableServicesEasyMintMetaDto) Set(val *ServicesEasyMintMetaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesEasyMintMetaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesEasyMintMetaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesEasyMintMetaDto(val *ServicesEasyMintMetaDto) *NullableServicesEasyMintMetaDto {
	return &NullableServicesEasyMintMetaDto{value: val, isSet: true}
}

func (v NullableServicesEasyMintMetaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesEasyMintMetaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


