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

// checks if the ServicesTransferBatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesTransferBatchDto{}

// ServicesTransferBatchDto struct for ServicesTransferBatchDto
type ServicesTransferBatchDto struct {
	Chain string `json:"chain"`
	ContractAddress string `json:"contract_address"`
	ContractType string `json:"contract_type"`
	Items []ServicesTransferItemDto `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _ServicesTransferBatchDto ServicesTransferBatchDto

// NewServicesTransferBatchDto instantiates a new ServicesTransferBatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesTransferBatchDto(chain string, contractAddress string, contractType string, items []ServicesTransferItemDto) *ServicesTransferBatchDto {
	this := ServicesTransferBatchDto{}
	this.Chain = chain
	this.ContractAddress = contractAddress
	this.ContractType = contractType
	this.Items = items
	return &this
}

// NewServicesTransferBatchDtoWithDefaults instantiates a new ServicesTransferBatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesTransferBatchDtoWithDefaults() *ServicesTransferBatchDto {
	this := ServicesTransferBatchDto{}
	return &this
}

// GetChain returns the Chain field value
func (o *ServicesTransferBatchDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferBatchDto) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesTransferBatchDto) SetChain(v string) {
	o.Chain = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ServicesTransferBatchDto) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferBatchDto) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ServicesTransferBatchDto) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetContractType returns the ContractType field value
func (o *ServicesTransferBatchDto) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferBatchDto) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *ServicesTransferBatchDto) SetContractType(v string) {
	o.ContractType = v
}

// GetItems returns the Items field value
func (o *ServicesTransferBatchDto) GetItems() []ServicesTransferItemDto {
	if o == nil {
		var ret []ServicesTransferItemDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferBatchDto) GetItemsOk() ([]ServicesTransferItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ServicesTransferBatchDto) SetItems(v []ServicesTransferItemDto) {
	o.Items = v
}

func (o ServicesTransferBatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesTransferBatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chain"] = o.Chain
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["contract_type"] = o.ContractType
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesTransferBatchDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chain",
		"contract_address",
		"contract_type",
		"items",
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

	varServicesTransferBatchDto := _ServicesTransferBatchDto{}

	err = json.Unmarshal(data, &varServicesTransferBatchDto)

	if err != nil {
		return err
	}

	*o = ServicesTransferBatchDto(varServicesTransferBatchDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesTransferBatchDto struct {
	value *ServicesTransferBatchDto
	isSet bool
}

func (v NullableServicesTransferBatchDto) Get() *ServicesTransferBatchDto {
	return v.value
}

func (v *NullableServicesTransferBatchDto) Set(val *ServicesTransferBatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesTransferBatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesTransferBatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesTransferBatchDto(val *ServicesTransferBatchDto) *NullableServicesTransferBatchDto {
	return &NullableServicesTransferBatchDto{value: val, isSet: true}
}

func (v NullableServicesTransferBatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesTransferBatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


