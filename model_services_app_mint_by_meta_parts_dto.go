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

// checks if the ServicesAppMintByMetaPartsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesAppMintByMetaPartsDto{}

// ServicesAppMintByMetaPartsDto struct for ServicesAppMintByMetaPartsDto
type ServicesAppMintByMetaPartsDto struct {
	// amount on same token id, only erc1155 contract could set large than 1, others set null or 1
	Amount *int32 `json:"amount,omitempty"`
	AnimationUrl *string `json:"animation_url,omitempty"`
	Attributes []ModelsExposedMetadataAttribute `json:"attributes,omitempty"`
	Chain string `json:"chain"`
	ContractAddress *string `json:"contract_address,omitempty"`
	Description *string `json:"description,omitempty"`
	FileUrl string `json:"file_url"`
	MintToAddress string `json:"mint_to_address"`
	Name string `json:"name"`
	// mint number, everyone with different token id (not conflict with AmountOnSameTokenID)
	Number *int32 `json:"number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesAppMintByMetaPartsDto ServicesAppMintByMetaPartsDto

// NewServicesAppMintByMetaPartsDto instantiates a new ServicesAppMintByMetaPartsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesAppMintByMetaPartsDto(chain string, fileUrl string, mintToAddress string, name string) *ServicesAppMintByMetaPartsDto {
	this := ServicesAppMintByMetaPartsDto{}
	this.Chain = chain
	this.FileUrl = fileUrl
	this.MintToAddress = mintToAddress
	this.Name = name
	return &this
}

// NewServicesAppMintByMetaPartsDtoWithDefaults instantiates a new ServicesAppMintByMetaPartsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesAppMintByMetaPartsDtoWithDefaults() *ServicesAppMintByMetaPartsDto {
	this := ServicesAppMintByMetaPartsDto{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetAmount() int32 {
	if o == nil || isNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetAmountOk() (*int32, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ServicesAppMintByMetaPartsDto) SetAmount(v int32) {
	o.Amount = &v
}

// GetAnimationUrl returns the AnimationUrl field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetAnimationUrl() string {
	if o == nil || isNil(o.AnimationUrl) {
		var ret string
		return ret
	}
	return *o.AnimationUrl
}

// GetAnimationUrlOk returns a tuple with the AnimationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetAnimationUrlOk() (*string, bool) {
	if o == nil || isNil(o.AnimationUrl) {
		return nil, false
	}
	return o.AnimationUrl, true
}

// HasAnimationUrl returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasAnimationUrl() bool {
	if o != nil && !isNil(o.AnimationUrl) {
		return true
	}

	return false
}

// SetAnimationUrl gets a reference to the given string and assigns it to the AnimationUrl field.
func (o *ServicesAppMintByMetaPartsDto) SetAnimationUrl(v string) {
	o.AnimationUrl = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetAttributes() []ModelsExposedMetadataAttribute {
	if o == nil || isNil(o.Attributes) {
		var ret []ModelsExposedMetadataAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetAttributesOk() ([]ModelsExposedMetadataAttribute, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []ModelsExposedMetadataAttribute and assigns it to the Attributes field.
func (o *ServicesAppMintByMetaPartsDto) SetAttributes(v []ModelsExposedMetadataAttribute) {
	o.Attributes = v
}

// GetChain returns the Chain field value
func (o *ServicesAppMintByMetaPartsDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesAppMintByMetaPartsDto) SetChain(v string) {
	o.Chain = v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetContractAddress() string {
	if o == nil || isNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetContractAddressOk() (*string, bool) {
	if o == nil || isNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasContractAddress() bool {
	if o != nil && !isNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ServicesAppMintByMetaPartsDto) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicesAppMintByMetaPartsDto) SetDescription(v string) {
	o.Description = &v
}

// GetFileUrl returns the FileUrl field value
func (o *ServicesAppMintByMetaPartsDto) GetFileUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUrl, true
}

// SetFileUrl sets field value
func (o *ServicesAppMintByMetaPartsDto) SetFileUrl(v string) {
	o.FileUrl = v
}

// GetMintToAddress returns the MintToAddress field value
func (o *ServicesAppMintByMetaPartsDto) GetMintToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MintToAddress
}

// GetMintToAddressOk returns a tuple with the MintToAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetMintToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MintToAddress, true
}

// SetMintToAddress sets field value
func (o *ServicesAppMintByMetaPartsDto) SetMintToAddress(v string) {
	o.MintToAddress = v
}

// GetName returns the Name field value
func (o *ServicesAppMintByMetaPartsDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServicesAppMintByMetaPartsDto) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ServicesAppMintByMetaPartsDto) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesAppMintByMetaPartsDto) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ServicesAppMintByMetaPartsDto) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *ServicesAppMintByMetaPartsDto) SetNumber(v int32) {
	o.Number = &v
}

func (o ServicesAppMintByMetaPartsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesAppMintByMetaPartsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.AnimationUrl) {
		toSerialize["animation_url"] = o.AnimationUrl
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["chain"] = o.Chain
	if !isNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["file_url"] = o.FileUrl
	toSerialize["mint_to_address"] = o.MintToAddress
	toSerialize["name"] = o.Name
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesAppMintByMetaPartsDto) UnmarshalJSON(bytes []byte) (err error) {
	varServicesAppMintByMetaPartsDto := _ServicesAppMintByMetaPartsDto{}

	if err = json.Unmarshal(bytes, &varServicesAppMintByMetaPartsDto); err == nil {
		*o = ServicesAppMintByMetaPartsDto(varServicesAppMintByMetaPartsDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "animation_url")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "description")
		delete(additionalProperties, "file_url")
		delete(additionalProperties, "mint_to_address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesAppMintByMetaPartsDto struct {
	value *ServicesAppMintByMetaPartsDto
	isSet bool
}

func (v NullableServicesAppMintByMetaPartsDto) Get() *ServicesAppMintByMetaPartsDto {
	return v.value
}

func (v *NullableServicesAppMintByMetaPartsDto) Set(val *ServicesAppMintByMetaPartsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesAppMintByMetaPartsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesAppMintByMetaPartsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesAppMintByMetaPartsDto(val *ServicesAppMintByMetaPartsDto) *NullableServicesAppMintByMetaPartsDto {
	return &NullableServicesAppMintByMetaPartsDto{value: val, isSet: true}
}

func (v NullableServicesAppMintByMetaPartsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesAppMintByMetaPartsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


