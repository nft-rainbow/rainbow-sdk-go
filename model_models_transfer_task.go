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

// checks if the ModelsTransferTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsTransferTask{}

// ModelsTransferTask struct for ModelsTransferTask
type ModelsTransferTask struct {
	Amount *int32 `json:"amount,omitempty"`
	AppId *int32 `json:"app_id,omitempty"`
	BlockReason *EnumsTransactionBlockReason `json:"block_reason,omitempty"`
	ChainId *int32 `json:"chain_id,omitempty"`
	ChainType *int32 `json:"chain_type,omitempty"`
	Contract *string `json:"contract,omitempty"`
	ContractType *int32 `json:"contract_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Error *string `json:"error,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Id *int32 `json:"id,omitempty"`
	// 0-pending, 1-success, 2-failed
	Status *int32 `json:"status,omitempty"`
	TokenId *string `json:"token_id,omitempty"`
	TransferFrom *string `json:"transfer_from,omitempty"`
	TransferTo *string `json:"transfer_to,omitempty"`
	TxId *int32 `json:"tx_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsTransferTask ModelsTransferTask

// NewModelsTransferTask instantiates a new ModelsTransferTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTransferTask() *ModelsTransferTask {
	this := ModelsTransferTask{}
	return &this
}

// NewModelsTransferTaskWithDefaults instantiates a new ModelsTransferTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTransferTaskWithDefaults() *ModelsTransferTask {
	this := ModelsTransferTask{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ModelsTransferTask) SetAmount(v int32) {
	o.Amount = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetAppId() int32 {
	if o == nil || IsNil(o.AppId) {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *ModelsTransferTask) SetAppId(v int32) {
	o.AppId = &v
}

// GetBlockReason returns the BlockReason field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetBlockReason() EnumsTransactionBlockReason {
	if o == nil || IsNil(o.BlockReason) {
		var ret EnumsTransactionBlockReason
		return ret
	}
	return *o.BlockReason
}

// GetBlockReasonOk returns a tuple with the BlockReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetBlockReasonOk() (*EnumsTransactionBlockReason, bool) {
	if o == nil || IsNil(o.BlockReason) {
		return nil, false
	}
	return o.BlockReason, true
}

// HasBlockReason returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasBlockReason() bool {
	if o != nil && !IsNil(o.BlockReason) {
		return true
	}

	return false
}

// SetBlockReason gets a reference to the given EnumsTransactionBlockReason and assigns it to the BlockReason field.
func (o *ModelsTransferTask) SetBlockReason(v EnumsTransactionBlockReason) {
	o.BlockReason = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetChainId() int32 {
	if o == nil || IsNil(o.ChainId) {
		var ret int32
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetChainIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given int32 and assigns it to the ChainId field.
func (o *ModelsTransferTask) SetChainId(v int32) {
	o.ChainId = &v
}

// GetChainType returns the ChainType field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetChainType() int32 {
	if o == nil || IsNil(o.ChainType) {
		var ret int32
		return ret
	}
	return *o.ChainType
}

// GetChainTypeOk returns a tuple with the ChainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetChainTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ChainType) {
		return nil, false
	}
	return o.ChainType, true
}

// HasChainType returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasChainType() bool {
	if o != nil && !IsNil(o.ChainType) {
		return true
	}

	return false
}

// SetChainType gets a reference to the given int32 and assigns it to the ChainType field.
func (o *ModelsTransferTask) SetChainType(v int32) {
	o.ChainType = &v
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetContract() string {
	if o == nil || IsNil(o.Contract) {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetContractOk() (*string, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *ModelsTransferTask) SetContract(v string) {
	o.Contract = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetContractType() int32 {
	if o == nil || IsNil(o.ContractType) {
		var ret int32
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetContractTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given int32 and assigns it to the ContractType field.
func (o *ModelsTransferTask) SetContractType(v int32) {
	o.ContractType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsTransferTask) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetDeletedAt() GormDeletedAt {
	if o == nil || IsNil(o.DeletedAt) {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsTransferTask) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelsTransferTask) SetError(v string) {
	o.Error = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ModelsTransferTask) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsTransferTask) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ModelsTransferTask) SetStatus(v int32) {
	o.Status = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ModelsTransferTask) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTransferFrom returns the TransferFrom field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetTransferFrom() string {
	if o == nil || IsNil(o.TransferFrom) {
		var ret string
		return ret
	}
	return *o.TransferFrom
}

// GetTransferFromOk returns a tuple with the TransferFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetTransferFromOk() (*string, bool) {
	if o == nil || IsNil(o.TransferFrom) {
		return nil, false
	}
	return o.TransferFrom, true
}

// HasTransferFrom returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasTransferFrom() bool {
	if o != nil && !IsNil(o.TransferFrom) {
		return true
	}

	return false
}

// SetTransferFrom gets a reference to the given string and assigns it to the TransferFrom field.
func (o *ModelsTransferTask) SetTransferFrom(v string) {
	o.TransferFrom = &v
}

// GetTransferTo returns the TransferTo field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetTransferTo() string {
	if o == nil || IsNil(o.TransferTo) {
		var ret string
		return ret
	}
	return *o.TransferTo
}

// GetTransferToOk returns a tuple with the TransferTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetTransferToOk() (*string, bool) {
	if o == nil || IsNil(o.TransferTo) {
		return nil, false
	}
	return o.TransferTo, true
}

// HasTransferTo returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasTransferTo() bool {
	if o != nil && !IsNil(o.TransferTo) {
		return true
	}

	return false
}

// SetTransferTo gets a reference to the given string and assigns it to the TransferTo field.
func (o *ModelsTransferTask) SetTransferTo(v string) {
	o.TransferTo = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetTxId() int32 {
	if o == nil || IsNil(o.TxId) {
		var ret int32
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetTxIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int32 and assigns it to the TxId field.
func (o *ModelsTransferTask) SetTxId(v int32) {
	o.TxId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsTransferTask) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTask) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsTransferTask) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsTransferTask) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsTransferTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsTransferTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !IsNil(o.BlockReason) {
		toSerialize["block_reason"] = o.BlockReason
	}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.ChainType) {
		toSerialize["chain_type"] = o.ChainType
	}
	if !IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !IsNil(o.ContractType) {
		toSerialize["contract_type"] = o.ContractType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.TransferFrom) {
		toSerialize["transfer_from"] = o.TransferFrom
	}
	if !IsNil(o.TransferTo) {
		toSerialize["transfer_to"] = o.TransferTo
	}
	if !IsNil(o.TxId) {
		toSerialize["tx_id"] = o.TxId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsTransferTask) UnmarshalJSON(data []byte) (err error) {
	varModelsTransferTask := _ModelsTransferTask{}

	err = json.Unmarshal(data, &varModelsTransferTask)

	if err != nil {
		return err
	}

	*o = ModelsTransferTask(varModelsTransferTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "block_reason")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_type")
		delete(additionalProperties, "contract")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "error")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "transfer_from")
		delete(additionalProperties, "transfer_to")
		delete(additionalProperties, "tx_id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsTransferTask struct {
	value *ModelsTransferTask
	isSet bool
}

func (v NullableModelsTransferTask) Get() *ModelsTransferTask {
	return v.value
}

func (v *NullableModelsTransferTask) Set(val *ModelsTransferTask) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTransferTask) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTransferTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTransferTask(val *ModelsTransferTask) *NullableModelsTransferTask {
	return &NullableModelsTransferTask{value: val, isSet: true}
}

func (v NullableModelsTransferTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTransferTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


