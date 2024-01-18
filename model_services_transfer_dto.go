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

// checks if the ServicesTransferDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesTransferDto{}

// ServicesTransferDto struct for ServicesTransferDto
type ServicesTransferDto struct {
	Amount *int32 `json:"amount,omitempty"`
	Chain string `json:"chain"`
	ContractAddress string `json:"contract_address"`
	ContractType string `json:"contract_type"`
	TokenId string `json:"token_id"`
	TransferFromAddress string `json:"transfer_from_address"`
	TransferToAddress string `json:"transfer_to_address"`
	AdditionalProperties map[string]interface{}
}

type _ServicesTransferDto ServicesTransferDto

// NewServicesTransferDto instantiates a new ServicesTransferDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesTransferDto(chain string, contractAddress string, contractType string, tokenId string, transferFromAddress string, transferToAddress string) *ServicesTransferDto {
	this := ServicesTransferDto{}
	this.Chain = chain
	this.ContractAddress = contractAddress
	this.ContractType = contractType
	this.TokenId = tokenId
	this.TransferFromAddress = transferFromAddress
	this.TransferToAddress = transferToAddress
	return &this
}

// NewServicesTransferDtoWithDefaults instantiates a new ServicesTransferDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesTransferDtoWithDefaults() *ServicesTransferDto {
	this := ServicesTransferDto{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ServicesTransferDto) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ServicesTransferDto) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ServicesTransferDto) SetAmount(v int32) {
	o.Amount = &v
}

// GetChain returns the Chain field value
func (o *ServicesTransferDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesTransferDto) SetChain(v string) {
	o.Chain = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ServicesTransferDto) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ServicesTransferDto) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetContractType returns the ContractType field value
func (o *ServicesTransferDto) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *ServicesTransferDto) SetContractType(v string) {
	o.ContractType = v
}

// GetTokenId returns the TokenId field value
func (o *ServicesTransferDto) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *ServicesTransferDto) SetTokenId(v string) {
	o.TokenId = v
}

// GetTransferFromAddress returns the TransferFromAddress field value
func (o *ServicesTransferDto) GetTransferFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferFromAddress
}

// GetTransferFromAddressOk returns a tuple with the TransferFromAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetTransferFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferFromAddress, true
}

// SetTransferFromAddress sets field value
func (o *ServicesTransferDto) SetTransferFromAddress(v string) {
	o.TransferFromAddress = v
}

// GetTransferToAddress returns the TransferToAddress field value
func (o *ServicesTransferDto) GetTransferToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferToAddress
}

// GetTransferToAddressOk returns a tuple with the TransferToAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesTransferDto) GetTransferToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferToAddress, true
}

// SetTransferToAddress sets field value
func (o *ServicesTransferDto) SetTransferToAddress(v string) {
	o.TransferToAddress = v
}

func (o ServicesTransferDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesTransferDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["chain"] = o.Chain
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["contract_type"] = o.ContractType
	toSerialize["token_id"] = o.TokenId
	toSerialize["transfer_from_address"] = o.TransferFromAddress
	toSerialize["transfer_to_address"] = o.TransferToAddress

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesTransferDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chain",
		"contract_address",
		"contract_type",
		"token_id",
		"transfer_from_address",
		"transfer_to_address",
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

	varServicesTransferDto := _ServicesTransferDto{}

	err = json.Unmarshal(data, &varServicesTransferDto)

	if err != nil {
		return err
	}

	*o = ServicesTransferDto(varServicesTransferDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "transfer_from_address")
		delete(additionalProperties, "transfer_to_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesTransferDto struct {
	value *ServicesTransferDto
	isSet bool
}

func (v NullableServicesTransferDto) Get() *ServicesTransferDto {
	return v.value
}

func (v *NullableServicesTransferDto) Set(val *ServicesTransferDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesTransferDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesTransferDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesTransferDto(val *ServicesTransferDto) *NullableServicesTransferDto {
	return &NullableServicesTransferDto{value: val, isSet: true}
}

func (v NullableServicesTransferDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesTransferDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


