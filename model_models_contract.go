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

// ModelsContract struct for ModelsContract
type ModelsContract struct {
	Address *string `json:"address,omitempty"`
	AppId *int32 `json:"app_id,omitempty"`
	BaseUri *string `json:"base_uri,omitempty"`
	ChainId *int32 `json:"chain_id,omitempty"`
	ChainType *int32 `json:"chain_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerAddress *string `json:"owner_address,omitempty"`
	RoyaltiesAddress *string `json:"royalties_address,omitempty"`
	RoyaltiesBps *int32 `json:"royalties_bps,omitempty"`
	// 0-pending, 1-success, 2-failed
	Status *int32 `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TokensBurnable *bool `json:"tokens_burnable,omitempty"`
	TokensTransferable *bool `json:"tokens_transferable,omitempty"`
	TransferCooldownTime *int32 `json:"transfer_cooldown_time,omitempty"`
	TxId *int32 `json:"tx_id,omitempty"`
	// 1-ERC721, 2-ERC1155
	Type *int32 `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsContract ModelsContract

// NewModelsContract instantiates a new ModelsContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsContract() *ModelsContract {
	this := ModelsContract{}
	return &this
}

// NewModelsContractWithDefaults instantiates a new ModelsContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsContractWithDefaults() *ModelsContract {
	this := ModelsContract{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ModelsContract) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelsContract) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ModelsContract) SetAddress(v string) {
	o.Address = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ModelsContract) GetAppId() int32 {
	if o == nil || o.AppId == nil {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetAppIdOk() (*int32, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ModelsContract) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *ModelsContract) SetAppId(v int32) {
	o.AppId = &v
}

// GetBaseUri returns the BaseUri field value if set, zero value otherwise.
func (o *ModelsContract) GetBaseUri() string {
	if o == nil || o.BaseUri == nil {
		var ret string
		return ret
	}
	return *o.BaseUri
}

// GetBaseUriOk returns a tuple with the BaseUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetBaseUriOk() (*string, bool) {
	if o == nil || o.BaseUri == nil {
		return nil, false
	}
	return o.BaseUri, true
}

// HasBaseUri returns a boolean if a field has been set.
func (o *ModelsContract) HasBaseUri() bool {
	if o != nil && o.BaseUri != nil {
		return true
	}

	return false
}

// SetBaseUri gets a reference to the given string and assigns it to the BaseUri field.
func (o *ModelsContract) SetBaseUri(v string) {
	o.BaseUri = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *ModelsContract) GetChainId() int32 {
	if o == nil || o.ChainId == nil {
		var ret int32
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetChainIdOk() (*int32, bool) {
	if o == nil || o.ChainId == nil {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *ModelsContract) HasChainId() bool {
	if o != nil && o.ChainId != nil {
		return true
	}

	return false
}

// SetChainId gets a reference to the given int32 and assigns it to the ChainId field.
func (o *ModelsContract) SetChainId(v int32) {
	o.ChainId = &v
}

// GetChainType returns the ChainType field value if set, zero value otherwise.
func (o *ModelsContract) GetChainType() int32 {
	if o == nil || o.ChainType == nil {
		var ret int32
		return ret
	}
	return *o.ChainType
}

// GetChainTypeOk returns a tuple with the ChainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetChainTypeOk() (*int32, bool) {
	if o == nil || o.ChainType == nil {
		return nil, false
	}
	return o.ChainType, true
}

// HasChainType returns a boolean if a field has been set.
func (o *ModelsContract) HasChainType() bool {
	if o != nil && o.ChainType != nil {
		return true
	}

	return false
}

// SetChainType gets a reference to the given int32 and assigns it to the ChainType field.
func (o *ModelsContract) SetChainType(v int32) {
	o.ChainType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsContract) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsContract) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsContract) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsContract) GetDeletedAt() GormDeletedAt {
	if o == nil || o.DeletedAt == nil {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsContract) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsContract) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ModelsContract) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ModelsContract) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ModelsContract) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsContract) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsContract) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsContract) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsContract) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsContract) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsContract) SetName(v string) {
	o.Name = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ModelsContract) GetOwnerAddress() string {
	if o == nil || o.OwnerAddress == nil {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetOwnerAddressOk() (*string, bool) {
	if o == nil || o.OwnerAddress == nil {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ModelsContract) HasOwnerAddress() bool {
	if o != nil && o.OwnerAddress != nil {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ModelsContract) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetRoyaltiesAddress returns the RoyaltiesAddress field value if set, zero value otherwise.
func (o *ModelsContract) GetRoyaltiesAddress() string {
	if o == nil || o.RoyaltiesAddress == nil {
		var ret string
		return ret
	}
	return *o.RoyaltiesAddress
}

// GetRoyaltiesAddressOk returns a tuple with the RoyaltiesAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetRoyaltiesAddressOk() (*string, bool) {
	if o == nil || o.RoyaltiesAddress == nil {
		return nil, false
	}
	return o.RoyaltiesAddress, true
}

// HasRoyaltiesAddress returns a boolean if a field has been set.
func (o *ModelsContract) HasRoyaltiesAddress() bool {
	if o != nil && o.RoyaltiesAddress != nil {
		return true
	}

	return false
}

// SetRoyaltiesAddress gets a reference to the given string and assigns it to the RoyaltiesAddress field.
func (o *ModelsContract) SetRoyaltiesAddress(v string) {
	o.RoyaltiesAddress = &v
}

// GetRoyaltiesBps returns the RoyaltiesBps field value if set, zero value otherwise.
func (o *ModelsContract) GetRoyaltiesBps() int32 {
	if o == nil || o.RoyaltiesBps == nil {
		var ret int32
		return ret
	}
	return *o.RoyaltiesBps
}

// GetRoyaltiesBpsOk returns a tuple with the RoyaltiesBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetRoyaltiesBpsOk() (*int32, bool) {
	if o == nil || o.RoyaltiesBps == nil {
		return nil, false
	}
	return o.RoyaltiesBps, true
}

// HasRoyaltiesBps returns a boolean if a field has been set.
func (o *ModelsContract) HasRoyaltiesBps() bool {
	if o != nil && o.RoyaltiesBps != nil {
		return true
	}

	return false
}

// SetRoyaltiesBps gets a reference to the given int32 and assigns it to the RoyaltiesBps field.
func (o *ModelsContract) SetRoyaltiesBps(v int32) {
	o.RoyaltiesBps = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsContract) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsContract) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ModelsContract) SetStatus(v int32) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ModelsContract) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ModelsContract) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ModelsContract) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTokensBurnable returns the TokensBurnable field value if set, zero value otherwise.
func (o *ModelsContract) GetTokensBurnable() bool {
	if o == nil || o.TokensBurnable == nil {
		var ret bool
		return ret
	}
	return *o.TokensBurnable
}

// GetTokensBurnableOk returns a tuple with the TokensBurnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetTokensBurnableOk() (*bool, bool) {
	if o == nil || o.TokensBurnable == nil {
		return nil, false
	}
	return o.TokensBurnable, true
}

// HasTokensBurnable returns a boolean if a field has been set.
func (o *ModelsContract) HasTokensBurnable() bool {
	if o != nil && o.TokensBurnable != nil {
		return true
	}

	return false
}

// SetTokensBurnable gets a reference to the given bool and assigns it to the TokensBurnable field.
func (o *ModelsContract) SetTokensBurnable(v bool) {
	o.TokensBurnable = &v
}

// GetTokensTransferable returns the TokensTransferable field value if set, zero value otherwise.
func (o *ModelsContract) GetTokensTransferable() bool {
	if o == nil || o.TokensTransferable == nil {
		var ret bool
		return ret
	}
	return *o.TokensTransferable
}

// GetTokensTransferableOk returns a tuple with the TokensTransferable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetTokensTransferableOk() (*bool, bool) {
	if o == nil || o.TokensTransferable == nil {
		return nil, false
	}
	return o.TokensTransferable, true
}

// HasTokensTransferable returns a boolean if a field has been set.
func (o *ModelsContract) HasTokensTransferable() bool {
	if o != nil && o.TokensTransferable != nil {
		return true
	}

	return false
}

// SetTokensTransferable gets a reference to the given bool and assigns it to the TokensTransferable field.
func (o *ModelsContract) SetTokensTransferable(v bool) {
	o.TokensTransferable = &v
}

// GetTransferCooldownTime returns the TransferCooldownTime field value if set, zero value otherwise.
func (o *ModelsContract) GetTransferCooldownTime() int32 {
	if o == nil || o.TransferCooldownTime == nil {
		var ret int32
		return ret
	}
	return *o.TransferCooldownTime
}

// GetTransferCooldownTimeOk returns a tuple with the TransferCooldownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetTransferCooldownTimeOk() (*int32, bool) {
	if o == nil || o.TransferCooldownTime == nil {
		return nil, false
	}
	return o.TransferCooldownTime, true
}

// HasTransferCooldownTime returns a boolean if a field has been set.
func (o *ModelsContract) HasTransferCooldownTime() bool {
	if o != nil && o.TransferCooldownTime != nil {
		return true
	}

	return false
}

// SetTransferCooldownTime gets a reference to the given int32 and assigns it to the TransferCooldownTime field.
func (o *ModelsContract) SetTransferCooldownTime(v int32) {
	o.TransferCooldownTime = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *ModelsContract) GetTxId() int32 {
	if o == nil || o.TxId == nil {
		var ret int32
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetTxIdOk() (*int32, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *ModelsContract) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int32 and assigns it to the TxId field.
func (o *ModelsContract) SetTxId(v int32) {
	o.TxId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelsContract) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelsContract) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *ModelsContract) SetType(v int32) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsContract) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContract) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsContract) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsContract) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsContract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.BaseUri != nil {
		toSerialize["base_uri"] = o.BaseUri
	}
	if o.ChainId != nil {
		toSerialize["chain_id"] = o.ChainId
	}
	if o.ChainType != nil {
		toSerialize["chain_type"] = o.ChainType
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerAddress != nil {
		toSerialize["owner_address"] = o.OwnerAddress
	}
	if o.RoyaltiesAddress != nil {
		toSerialize["royalties_address"] = o.RoyaltiesAddress
	}
	if o.RoyaltiesBps != nil {
		toSerialize["royalties_bps"] = o.RoyaltiesBps
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.TokensBurnable != nil {
		toSerialize["tokens_burnable"] = o.TokensBurnable
	}
	if o.TokensTransferable != nil {
		toSerialize["tokens_transferable"] = o.TokensTransferable
	}
	if o.TransferCooldownTime != nil {
		toSerialize["transfer_cooldown_time"] = o.TransferCooldownTime
	}
	if o.TxId != nil {
		toSerialize["tx_id"] = o.TxId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsContract) UnmarshalJSON(bytes []byte) (err error) {
	varModelsContract := _ModelsContract{}

	if err = json.Unmarshal(bytes, &varModelsContract); err == nil {
		*o = ModelsContract(varModelsContract)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "base_uri")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_address")
		delete(additionalProperties, "royalties_address")
		delete(additionalProperties, "royalties_bps")
		delete(additionalProperties, "status")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "tokens_burnable")
		delete(additionalProperties, "tokens_transferable")
		delete(additionalProperties, "transfer_cooldown_time")
		delete(additionalProperties, "tx_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsContract struct {
	value *ModelsContract
	isSet bool
}

func (v NullableModelsContract) Get() *ModelsContract {
	return v.value
}

func (v *NullableModelsContract) Set(val *ModelsContract) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsContract) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsContract(val *ModelsContract) *NullableModelsContract {
	return &NullableModelsContract{value: val, isSet: true}
}

func (v NullableModelsContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


