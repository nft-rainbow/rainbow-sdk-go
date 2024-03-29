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

// checks if the ServicesBurnBatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesBurnBatchDto{}

// ServicesBurnBatchDto struct for ServicesBurnBatchDto
type ServicesBurnBatchDto struct {
	Chain string `json:"chain"`
	ContractAddress string `json:"contract_address"`
	ContractType string `json:"contract_type"`
	Items []ServicesBurnItemDto `json:"items"`
	User *string `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesBurnBatchDto ServicesBurnBatchDto

// NewServicesBurnBatchDto instantiates a new ServicesBurnBatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesBurnBatchDto(chain string, contractAddress string, contractType string, items []ServicesBurnItemDto) *ServicesBurnBatchDto {
	this := ServicesBurnBatchDto{}
	this.Chain = chain
	this.ContractAddress = contractAddress
	this.ContractType = contractType
	this.Items = items
	return &this
}

// NewServicesBurnBatchDtoWithDefaults instantiates a new ServicesBurnBatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesBurnBatchDtoWithDefaults() *ServicesBurnBatchDto {
	this := ServicesBurnBatchDto{}
	return &this
}

// GetChain returns the Chain field value
func (o *ServicesBurnBatchDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesBurnBatchDto) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesBurnBatchDto) SetChain(v string) {
	o.Chain = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ServicesBurnBatchDto) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesBurnBatchDto) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ServicesBurnBatchDto) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetContractType returns the ContractType field value
func (o *ServicesBurnBatchDto) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *ServicesBurnBatchDto) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *ServicesBurnBatchDto) SetContractType(v string) {
	o.ContractType = v
}

// GetItems returns the Items field value
func (o *ServicesBurnBatchDto) GetItems() []ServicesBurnItemDto {
	if o == nil {
		var ret []ServicesBurnItemDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServicesBurnBatchDto) GetItemsOk() ([]ServicesBurnItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ServicesBurnBatchDto) SetItems(v []ServicesBurnItemDto) {
	o.Items = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ServicesBurnBatchDto) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesBurnBatchDto) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ServicesBurnBatchDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ServicesBurnBatchDto) SetUser(v string) {
	o.User = &v
}

func (o ServicesBurnBatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesBurnBatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chain"] = o.Chain
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["contract_type"] = o.ContractType
	toSerialize["items"] = o.Items
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesBurnBatchDto) UnmarshalJSON(data []byte) (err error) {
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

	varServicesBurnBatchDto := _ServicesBurnBatchDto{}

	err = json.Unmarshal(data, &varServicesBurnBatchDto)

	if err != nil {
		return err
	}

	*o = ServicesBurnBatchDto(varServicesBurnBatchDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "items")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesBurnBatchDto struct {
	value *ServicesBurnBatchDto
	isSet bool
}

func (v NullableServicesBurnBatchDto) Get() *ServicesBurnBatchDto {
	return v.value
}

func (v *NullableServicesBurnBatchDto) Set(val *ServicesBurnBatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesBurnBatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesBurnBatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesBurnBatchDto(val *ServicesBurnBatchDto) *NullableServicesBurnBatchDto {
	return &NullableServicesBurnBatchDto{value: val, isSet: true}
}

func (v NullableServicesBurnBatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesBurnBatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


