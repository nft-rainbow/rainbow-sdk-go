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
	AnimationUrl *string `json:"animation_url,omitempty"`
	Attributes []ModelsExposedMetadataAttribute `json:"attributes,omitempty"`
	Chain string `json:"chain"`
	ContractAddress *string `json:"contract_address,omitempty"`
	Description *string `json:"description,omitempty"`
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
func NewServicesEasyMintMetaDto(chain string, fileUrl string, mintToAddress string, name string) *ServicesEasyMintMetaDto {
	this := ServicesEasyMintMetaDto{}
	this.Chain = chain
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

// GetAnimationUrl returns the AnimationUrl field value if set, zero value otherwise.
func (o *ServicesEasyMintMetaDto) GetAnimationUrl() string {
	if o == nil || IsNil(o.AnimationUrl) {
		var ret string
		return ret
	}
	return *o.AnimationUrl
}

// GetAnimationUrlOk returns a tuple with the AnimationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetAnimationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AnimationUrl) {
    return nil, false
	}
	return o.AnimationUrl, true
}

// HasAnimationUrl returns a boolean if a field has been set.
func (o *ServicesEasyMintMetaDto) HasAnimationUrl() bool {
	if o != nil && !IsNil(o.AnimationUrl) {
		return true
	}

	return false
}

// SetAnimationUrl gets a reference to the given string and assigns it to the AnimationUrl field.
func (o *ServicesEasyMintMetaDto) SetAnimationUrl(v string) {
	o.AnimationUrl = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServicesEasyMintMetaDto) GetAttributes() []ModelsExposedMetadataAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []ModelsExposedMetadataAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetAttributesOk() ([]ModelsExposedMetadataAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServicesEasyMintMetaDto) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []ModelsExposedMetadataAttribute and assigns it to the Attributes field.
func (o *ServicesEasyMintMetaDto) SetAttributes(v []ModelsExposedMetadataAttribute) {
	o.Attributes = v
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

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ServicesEasyMintMetaDto) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
    return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ServicesEasyMintMetaDto) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ServicesEasyMintMetaDto) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicesEasyMintMetaDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesEasyMintMetaDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicesEasyMintMetaDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicesEasyMintMetaDto) SetDescription(v string) {
	o.Description = &v
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
	if !IsNil(o.AnimationUrl) {
		toSerialize["animation_url"] = o.AnimationUrl
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !IsNil(o.Description) {
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
		delete(additionalProperties, "animation_url")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
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


