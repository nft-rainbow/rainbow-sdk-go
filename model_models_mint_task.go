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

// ModelsMintTask struct for ModelsMintTask
type ModelsMintTask struct {
	Amount *int32 `json:"amount,omitempty"`
	AppId *int32 `json:"app_id,omitempty"`
	ChainId *int32 `json:"chain_id,omitempty"`
	ChainType *int32 `json:"chain_type,omitempty"`
	Contract *string `json:"contract,omitempty"`
	ContractType *int32 `json:"contract_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Error *string `json:"error,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Id *int32 `json:"id,omitempty"`
	MintTo *string `json:"mint_to,omitempty"`
	MintType *int32 `json:"mint_type,omitempty"`
	// 0-pending, 1-success, 2-failed
	Status *int32 `json:"status,omitempty"`
	TokenId *string `json:"token_id,omitempty"`
	TokenUri *string `json:"token_uri,omitempty"`
	TxId *int32 `json:"tx_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsMintTask ModelsMintTask

// NewModelsMintTask instantiates a new ModelsMintTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsMintTask() *ModelsMintTask {
	this := ModelsMintTask{}
	return &this
}

// NewModelsMintTaskWithDefaults instantiates a new ModelsMintTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsMintTaskWithDefaults() *ModelsMintTask {
	this := ModelsMintTask{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ModelsMintTask) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelsMintTask) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ModelsMintTask) SetAmount(v int32) {
	o.Amount = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ModelsMintTask) GetAppId() int32 {
	if o == nil || o.AppId == nil {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetAppIdOk() (*int32, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ModelsMintTask) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *ModelsMintTask) SetAppId(v int32) {
	o.AppId = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *ModelsMintTask) GetChainId() int32 {
	if o == nil || o.ChainId == nil {
		var ret int32
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetChainIdOk() (*int32, bool) {
	if o == nil || o.ChainId == nil {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *ModelsMintTask) HasChainId() bool {
	if o != nil && o.ChainId != nil {
		return true
	}

	return false
}

// SetChainId gets a reference to the given int32 and assigns it to the ChainId field.
func (o *ModelsMintTask) SetChainId(v int32) {
	o.ChainId = &v
}

// GetChainType returns the ChainType field value if set, zero value otherwise.
func (o *ModelsMintTask) GetChainType() int32 {
	if o == nil || o.ChainType == nil {
		var ret int32
		return ret
	}
	return *o.ChainType
}

// GetChainTypeOk returns a tuple with the ChainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetChainTypeOk() (*int32, bool) {
	if o == nil || o.ChainType == nil {
		return nil, false
	}
	return o.ChainType, true
}

// HasChainType returns a boolean if a field has been set.
func (o *ModelsMintTask) HasChainType() bool {
	if o != nil && o.ChainType != nil {
		return true
	}

	return false
}

// SetChainType gets a reference to the given int32 and assigns it to the ChainType field.
func (o *ModelsMintTask) SetChainType(v int32) {
	o.ChainType = &v
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *ModelsMintTask) GetContract() string {
	if o == nil || o.Contract == nil {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetContractOk() (*string, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *ModelsMintTask) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *ModelsMintTask) SetContract(v string) {
	o.Contract = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *ModelsMintTask) GetContractType() int32 {
	if o == nil || o.ContractType == nil {
		var ret int32
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetContractTypeOk() (*int32, bool) {
	if o == nil || o.ContractType == nil {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *ModelsMintTask) HasContractType() bool {
	if o != nil && o.ContractType != nil {
		return true
	}

	return false
}

// SetContractType gets a reference to the given int32 and assigns it to the ContractType field.
func (o *ModelsMintTask) SetContractType(v int32) {
	o.ContractType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsMintTask) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsMintTask) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsMintTask) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsMintTask) GetDeletedAt() GormDeletedAt {
	if o == nil || o.DeletedAt == nil {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsMintTask) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsMintTask) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsMintTask) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsMintTask) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelsMintTask) SetError(v string) {
	o.Error = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ModelsMintTask) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ModelsMintTask) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ModelsMintTask) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsMintTask) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsMintTask) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsMintTask) SetId(v int32) {
	o.Id = &v
}

// GetMintTo returns the MintTo field value if set, zero value otherwise.
func (o *ModelsMintTask) GetMintTo() string {
	if o == nil || o.MintTo == nil {
		var ret string
		return ret
	}
	return *o.MintTo
}

// GetMintToOk returns a tuple with the MintTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetMintToOk() (*string, bool) {
	if o == nil || o.MintTo == nil {
		return nil, false
	}
	return o.MintTo, true
}

// HasMintTo returns a boolean if a field has been set.
func (o *ModelsMintTask) HasMintTo() bool {
	if o != nil && o.MintTo != nil {
		return true
	}

	return false
}

// SetMintTo gets a reference to the given string and assigns it to the MintTo field.
func (o *ModelsMintTask) SetMintTo(v string) {
	o.MintTo = &v
}

// GetMintType returns the MintType field value if set, zero value otherwise.
func (o *ModelsMintTask) GetMintType() int32 {
	if o == nil || o.MintType == nil {
		var ret int32
		return ret
	}
	return *o.MintType
}

// GetMintTypeOk returns a tuple with the MintType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetMintTypeOk() (*int32, bool) {
	if o == nil || o.MintType == nil {
		return nil, false
	}
	return o.MintType, true
}

// HasMintType returns a boolean if a field has been set.
func (o *ModelsMintTask) HasMintType() bool {
	if o != nil && o.MintType != nil {
		return true
	}

	return false
}

// SetMintType gets a reference to the given int32 and assigns it to the MintType field.
func (o *ModelsMintTask) SetMintType(v int32) {
	o.MintType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsMintTask) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsMintTask) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ModelsMintTask) SetStatus(v int32) {
	o.Status = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ModelsMintTask) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ModelsMintTask) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ModelsMintTask) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenUri returns the TokenUri field value if set, zero value otherwise.
func (o *ModelsMintTask) GetTokenUri() string {
	if o == nil || o.TokenUri == nil {
		var ret string
		return ret
	}
	return *o.TokenUri
}

// GetTokenUriOk returns a tuple with the TokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetTokenUriOk() (*string, bool) {
	if o == nil || o.TokenUri == nil {
		return nil, false
	}
	return o.TokenUri, true
}

// HasTokenUri returns a boolean if a field has been set.
func (o *ModelsMintTask) HasTokenUri() bool {
	if o != nil && o.TokenUri != nil {
		return true
	}

	return false
}

// SetTokenUri gets a reference to the given string and assigns it to the TokenUri field.
func (o *ModelsMintTask) SetTokenUri(v string) {
	o.TokenUri = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *ModelsMintTask) GetTxId() int32 {
	if o == nil || o.TxId == nil {
		var ret int32
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetTxIdOk() (*int32, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *ModelsMintTask) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int32 and assigns it to the TxId field.
func (o *ModelsMintTask) SetTxId(v int32) {
	o.TxId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsMintTask) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTask) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsMintTask) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsMintTask) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsMintTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.ChainId != nil {
		toSerialize["chain_id"] = o.ChainId
	}
	if o.ChainType != nil {
		toSerialize["chain_type"] = o.ChainType
	}
	if o.Contract != nil {
		toSerialize["contract"] = o.Contract
	}
	if o.ContractType != nil {
		toSerialize["contract_type"] = o.ContractType
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MintTo != nil {
		toSerialize["mint_to"] = o.MintTo
	}
	if o.MintType != nil {
		toSerialize["mint_type"] = o.MintType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TokenId != nil {
		toSerialize["token_id"] = o.TokenId
	}
	if o.TokenUri != nil {
		toSerialize["token_uri"] = o.TokenUri
	}
	if o.TxId != nil {
		toSerialize["tx_id"] = o.TxId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsMintTask) UnmarshalJSON(bytes []byte) (err error) {
	varModelsMintTask := _ModelsMintTask{}

	if err = json.Unmarshal(bytes, &varModelsMintTask); err == nil {
		*o = ModelsMintTask(varModelsMintTask)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_type")
		delete(additionalProperties, "contract")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "error")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mint_to")
		delete(additionalProperties, "mint_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "token_uri")
		delete(additionalProperties, "tx_id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsMintTask struct {
	value *ModelsMintTask
	isSet bool
}

func (v NullableModelsMintTask) Get() *ModelsMintTask {
	return v.value
}

func (v *NullableModelsMintTask) Set(val *ModelsMintTask) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsMintTask) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsMintTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsMintTask(val *ModelsMintTask) *NullableModelsMintTask {
	return &NullableModelsMintTask{value: val, isSet: true}
}

func (v NullableModelsMintTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsMintTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

