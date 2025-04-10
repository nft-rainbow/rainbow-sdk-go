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

// checks if the ModelsAutoSponsorContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAutoSponsorContract{}

// ModelsAutoSponsorContract struct for ModelsAutoSponsorContract
type ModelsAutoSponsorContract struct {
	Address *string `json:"address,omitempty"`
	AutoSponsor *bool `json:"auto_sponsor,omitempty"`
	// useless
	ContractId *int32 `json:"contract_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Id *int32 `json:"id,omitempty"`
	StorageRechargeAmount *int32 `json:"storage_recharge_amount,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsAutoSponsorContract ModelsAutoSponsorContract

// NewModelsAutoSponsorContract instantiates a new ModelsAutoSponsorContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAutoSponsorContract() *ModelsAutoSponsorContract {
	this := ModelsAutoSponsorContract{}
	return &this
}

// NewModelsAutoSponsorContractWithDefaults instantiates a new ModelsAutoSponsorContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAutoSponsorContractWithDefaults() *ModelsAutoSponsorContract {
	this := ModelsAutoSponsorContract{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ModelsAutoSponsorContract) SetAddress(v string) {
	o.Address = &v
}

// GetAutoSponsor returns the AutoSponsor field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetAutoSponsor() bool {
	if o == nil || IsNil(o.AutoSponsor) {
		var ret bool
		return ret
	}
	return *o.AutoSponsor
}

// GetAutoSponsorOk returns a tuple with the AutoSponsor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetAutoSponsorOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSponsor) {
		return nil, false
	}
	return o.AutoSponsor, true
}

// HasAutoSponsor returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasAutoSponsor() bool {
	if o != nil && !IsNil(o.AutoSponsor) {
		return true
	}

	return false
}

// SetAutoSponsor gets a reference to the given bool and assigns it to the AutoSponsor field.
func (o *ModelsAutoSponsorContract) SetAutoSponsor(v bool) {
	o.AutoSponsor = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetContractId() int32 {
	if o == nil || IsNil(o.ContractId) {
		var ret int32
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetContractIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given int32 and assigns it to the ContractId field.
func (o *ModelsAutoSponsorContract) SetContractId(v int32) {
	o.ContractId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsAutoSponsorContract) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetDeletedAt() GormDeletedAt {
	if o == nil || IsNil(o.DeletedAt) {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsAutoSponsorContract) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsAutoSponsorContract) SetId(v int32) {
	o.Id = &v
}

// GetStorageRechargeAmount returns the StorageRechargeAmount field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetStorageRechargeAmount() int32 {
	if o == nil || IsNil(o.StorageRechargeAmount) {
		var ret int32
		return ret
	}
	return *o.StorageRechargeAmount
}

// GetStorageRechargeAmountOk returns a tuple with the StorageRechargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetStorageRechargeAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageRechargeAmount) {
		return nil, false
	}
	return o.StorageRechargeAmount, true
}

// HasStorageRechargeAmount returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasStorageRechargeAmount() bool {
	if o != nil && !IsNil(o.StorageRechargeAmount) {
		return true
	}

	return false
}

// SetStorageRechargeAmount gets a reference to the given int32 and assigns it to the StorageRechargeAmount field.
func (o *ModelsAutoSponsorContract) SetStorageRechargeAmount(v int32) {
	o.StorageRechargeAmount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsAutoSponsorContract) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ModelsAutoSponsorContract) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAutoSponsorContract) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ModelsAutoSponsorContract) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ModelsAutoSponsorContract) SetUserId(v int32) {
	o.UserId = &v
}

func (o ModelsAutoSponsorContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAutoSponsorContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AutoSponsor) {
		toSerialize["auto_sponsor"] = o.AutoSponsor
	}
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.StorageRechargeAmount) {
		toSerialize["storage_recharge_amount"] = o.StorageRechargeAmount
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsAutoSponsorContract) UnmarshalJSON(data []byte) (err error) {
	varModelsAutoSponsorContract := _ModelsAutoSponsorContract{}

	err = json.Unmarshal(data, &varModelsAutoSponsorContract)

	if err != nil {
		return err
	}

	*o = ModelsAutoSponsorContract(varModelsAutoSponsorContract)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "auto_sponsor")
		delete(additionalProperties, "contract_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "storage_recharge_amount")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsAutoSponsorContract struct {
	value *ModelsAutoSponsorContract
	isSet bool
}

func (v NullableModelsAutoSponsorContract) Get() *ModelsAutoSponsorContract {
	return v.value
}

func (v *NullableModelsAutoSponsorContract) Set(val *ModelsAutoSponsorContract) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAutoSponsorContract) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAutoSponsorContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAutoSponsorContract(val *ModelsAutoSponsorContract) *NullableModelsAutoSponsorContract {
	return &NullableModelsAutoSponsorContract{value: val, isSet: true}
}

func (v NullableModelsAutoSponsorContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAutoSponsorContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


