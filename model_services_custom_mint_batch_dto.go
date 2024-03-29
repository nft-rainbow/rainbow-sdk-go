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

// checks if the ServicesCustomMintBatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesCustomMintBatchDto{}

// ServicesCustomMintBatchDto struct for ServicesCustomMintBatchDto
type ServicesCustomMintBatchDto struct {
	Chain string `json:"chain"`
	ContractAddress string `json:"contract_address"`
	MintItems []ServicesMintItemDto `json:"mint_items"`
	AdditionalProperties map[string]interface{}
}

type _ServicesCustomMintBatchDto ServicesCustomMintBatchDto

// NewServicesCustomMintBatchDto instantiates a new ServicesCustomMintBatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesCustomMintBatchDto(chain string, contractAddress string, mintItems []ServicesMintItemDto) *ServicesCustomMintBatchDto {
	this := ServicesCustomMintBatchDto{}
	this.Chain = chain
	this.ContractAddress = contractAddress
	this.MintItems = mintItems
	return &this
}

// NewServicesCustomMintBatchDtoWithDefaults instantiates a new ServicesCustomMintBatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesCustomMintBatchDtoWithDefaults() *ServicesCustomMintBatchDto {
	this := ServicesCustomMintBatchDto{}
	return &this
}

// GetChain returns the Chain field value
func (o *ServicesCustomMintBatchDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesCustomMintBatchDto) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesCustomMintBatchDto) SetChain(v string) {
	o.Chain = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ServicesCustomMintBatchDto) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesCustomMintBatchDto) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ServicesCustomMintBatchDto) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetMintItems returns the MintItems field value
func (o *ServicesCustomMintBatchDto) GetMintItems() []ServicesMintItemDto {
	if o == nil {
		var ret []ServicesMintItemDto
		return ret
	}

	return o.MintItems
}

// GetMintItemsOk returns a tuple with the MintItems field value
// and a boolean to check if the value has been set.
func (o *ServicesCustomMintBatchDto) GetMintItemsOk() ([]ServicesMintItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.MintItems, true
}

// SetMintItems sets field value
func (o *ServicesCustomMintBatchDto) SetMintItems(v []ServicesMintItemDto) {
	o.MintItems = v
}

func (o ServicesCustomMintBatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesCustomMintBatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chain"] = o.Chain
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["mint_items"] = o.MintItems

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesCustomMintBatchDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chain",
		"contract_address",
		"mint_items",
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

	varServicesCustomMintBatchDto := _ServicesCustomMintBatchDto{}

	err = json.Unmarshal(data, &varServicesCustomMintBatchDto)

	if err != nil {
		return err
	}

	*o = ServicesCustomMintBatchDto(varServicesCustomMintBatchDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "mint_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesCustomMintBatchDto struct {
	value *ServicesCustomMintBatchDto
	isSet bool
}

func (v NullableServicesCustomMintBatchDto) Get() *ServicesCustomMintBatchDto {
	return v.value
}

func (v *NullableServicesCustomMintBatchDto) Set(val *ServicesCustomMintBatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesCustomMintBatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesCustomMintBatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesCustomMintBatchDto(val *ServicesCustomMintBatchDto) *NullableServicesCustomMintBatchDto {
	return &NullableServicesCustomMintBatchDto{value: val, isSet: true}
}

func (v NullableServicesCustomMintBatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesCustomMintBatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


