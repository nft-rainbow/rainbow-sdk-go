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

// ServicesContractDeployDto struct for ServicesContractDeployDto
type ServicesContractDeployDto struct {
	// default: true
	AutoSponsor *bool `json:"auto_sponsor,omitempty"`
	BaseUri *string `json:"base_uri,omitempty"`
	Chain string `json:"chain"`
	// default: true
	IsSponsorForAllUser *bool `json:"is_sponsor_for_all_user,omitempty"`
	Name string `json:"name"`
	OwnerAddress *string `json:"owner_address,omitempty"`
	RoyaltiesAddress *string `json:"royalties_address,omitempty"`
	RoyaltiesBps *int32 `json:"royalties_bps,omitempty"`
	Symbol string `json:"symbol"`
	// default: true
	TokensBurnable *bool `json:"tokens_burnable,omitempty"`
	// default: true
	TokensTransferableByAdmin *bool `json:"tokens_transferable_by_admin,omitempty"`
	// default: true
	TokensTransferableByUser *bool `json:"tokens_transferable_by_user,omitempty"`
	// default: 0
	TransferCooldownTime *int32 `json:"transfer_cooldown_time,omitempty"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ServicesContractDeployDto ServicesContractDeployDto

// NewServicesContractDeployDto instantiates a new ServicesContractDeployDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesContractDeployDto(chain string, name string, symbol string, type_ string) *ServicesContractDeployDto {
	this := ServicesContractDeployDto{}
	this.Chain = chain
	this.Name = name
	this.Symbol = symbol
	this.Type = type_
	return &this
}

// NewServicesContractDeployDtoWithDefaults instantiates a new ServicesContractDeployDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesContractDeployDtoWithDefaults() *ServicesContractDeployDto {
	this := ServicesContractDeployDto{}
	return &this
}

// GetAutoSponsor returns the AutoSponsor field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetAutoSponsor() bool {
	if o == nil || isNil(o.AutoSponsor) {
		var ret bool
		return ret
	}
	return *o.AutoSponsor
}

// GetAutoSponsorOk returns a tuple with the AutoSponsor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetAutoSponsorOk() (*bool, bool) {
	if o == nil || isNil(o.AutoSponsor) {
    return nil, false
	}
	return o.AutoSponsor, true
}

// HasAutoSponsor returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasAutoSponsor() bool {
	if o != nil && !isNil(o.AutoSponsor) {
		return true
	}

	return false
}

// SetAutoSponsor gets a reference to the given bool and assigns it to the AutoSponsor field.
func (o *ServicesContractDeployDto) SetAutoSponsor(v bool) {
	o.AutoSponsor = &v
}

// GetBaseUri returns the BaseUri field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetBaseUri() string {
	if o == nil || isNil(o.BaseUri) {
		var ret string
		return ret
	}
	return *o.BaseUri
}

// GetBaseUriOk returns a tuple with the BaseUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetBaseUriOk() (*string, bool) {
	if o == nil || isNil(o.BaseUri) {
    return nil, false
	}
	return o.BaseUri, true
}

// HasBaseUri returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasBaseUri() bool {
	if o != nil && !isNil(o.BaseUri) {
		return true
	}

	return false
}

// SetBaseUri gets a reference to the given string and assigns it to the BaseUri field.
func (o *ServicesContractDeployDto) SetBaseUri(v string) {
	o.BaseUri = &v
}

// GetChain returns the Chain field value
func (o *ServicesContractDeployDto) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetChainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *ServicesContractDeployDto) SetChain(v string) {
	o.Chain = v
}

// GetIsSponsorForAllUser returns the IsSponsorForAllUser field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetIsSponsorForAllUser() bool {
	if o == nil || isNil(o.IsSponsorForAllUser) {
		var ret bool
		return ret
	}
	return *o.IsSponsorForAllUser
}

// GetIsSponsorForAllUserOk returns a tuple with the IsSponsorForAllUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetIsSponsorForAllUserOk() (*bool, bool) {
	if o == nil || isNil(o.IsSponsorForAllUser) {
    return nil, false
	}
	return o.IsSponsorForAllUser, true
}

// HasIsSponsorForAllUser returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasIsSponsorForAllUser() bool {
	if o != nil && !isNil(o.IsSponsorForAllUser) {
		return true
	}

	return false
}

// SetIsSponsorForAllUser gets a reference to the given bool and assigns it to the IsSponsorForAllUser field.
func (o *ServicesContractDeployDto) SetIsSponsorForAllUser(v bool) {
	o.IsSponsorForAllUser = &v
}

// GetName returns the Name field value
func (o *ServicesContractDeployDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServicesContractDeployDto) SetName(v string) {
	o.Name = v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetOwnerAddress() string {
	if o == nil || isNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetOwnerAddressOk() (*string, bool) {
	if o == nil || isNil(o.OwnerAddress) {
    return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasOwnerAddress() bool {
	if o != nil && !isNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ServicesContractDeployDto) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetRoyaltiesAddress returns the RoyaltiesAddress field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetRoyaltiesAddress() string {
	if o == nil || isNil(o.RoyaltiesAddress) {
		var ret string
		return ret
	}
	return *o.RoyaltiesAddress
}

// GetRoyaltiesAddressOk returns a tuple with the RoyaltiesAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetRoyaltiesAddressOk() (*string, bool) {
	if o == nil || isNil(o.RoyaltiesAddress) {
    return nil, false
	}
	return o.RoyaltiesAddress, true
}

// HasRoyaltiesAddress returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasRoyaltiesAddress() bool {
	if o != nil && !isNil(o.RoyaltiesAddress) {
		return true
	}

	return false
}

// SetRoyaltiesAddress gets a reference to the given string and assigns it to the RoyaltiesAddress field.
func (o *ServicesContractDeployDto) SetRoyaltiesAddress(v string) {
	o.RoyaltiesAddress = &v
}

// GetRoyaltiesBps returns the RoyaltiesBps field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetRoyaltiesBps() int32 {
	if o == nil || isNil(o.RoyaltiesBps) {
		var ret int32
		return ret
	}
	return *o.RoyaltiesBps
}

// GetRoyaltiesBpsOk returns a tuple with the RoyaltiesBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetRoyaltiesBpsOk() (*int32, bool) {
	if o == nil || isNil(o.RoyaltiesBps) {
    return nil, false
	}
	return o.RoyaltiesBps, true
}

// HasRoyaltiesBps returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasRoyaltiesBps() bool {
	if o != nil && !isNil(o.RoyaltiesBps) {
		return true
	}

	return false
}

// SetRoyaltiesBps gets a reference to the given int32 and assigns it to the RoyaltiesBps field.
func (o *ServicesContractDeployDto) SetRoyaltiesBps(v int32) {
	o.RoyaltiesBps = &v
}

// GetSymbol returns the Symbol field value
func (o *ServicesContractDeployDto) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetSymbolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ServicesContractDeployDto) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokensBurnable returns the TokensBurnable field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetTokensBurnable() bool {
	if o == nil || isNil(o.TokensBurnable) {
		var ret bool
		return ret
	}
	return *o.TokensBurnable
}

// GetTokensBurnableOk returns a tuple with the TokensBurnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetTokensBurnableOk() (*bool, bool) {
	if o == nil || isNil(o.TokensBurnable) {
    return nil, false
	}
	return o.TokensBurnable, true
}

// HasTokensBurnable returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasTokensBurnable() bool {
	if o != nil && !isNil(o.TokensBurnable) {
		return true
	}

	return false
}

// SetTokensBurnable gets a reference to the given bool and assigns it to the TokensBurnable field.
func (o *ServicesContractDeployDto) SetTokensBurnable(v bool) {
	o.TokensBurnable = &v
}

// GetTokensTransferableByAdmin returns the TokensTransferableByAdmin field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetTokensTransferableByAdmin() bool {
	if o == nil || isNil(o.TokensTransferableByAdmin) {
		var ret bool
		return ret
	}
	return *o.TokensTransferableByAdmin
}

// GetTokensTransferableByAdminOk returns a tuple with the TokensTransferableByAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetTokensTransferableByAdminOk() (*bool, bool) {
	if o == nil || isNil(o.TokensTransferableByAdmin) {
    return nil, false
	}
	return o.TokensTransferableByAdmin, true
}

// HasTokensTransferableByAdmin returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasTokensTransferableByAdmin() bool {
	if o != nil && !isNil(o.TokensTransferableByAdmin) {
		return true
	}

	return false
}

// SetTokensTransferableByAdmin gets a reference to the given bool and assigns it to the TokensTransferableByAdmin field.
func (o *ServicesContractDeployDto) SetTokensTransferableByAdmin(v bool) {
	o.TokensTransferableByAdmin = &v
}

// GetTokensTransferableByUser returns the TokensTransferableByUser field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetTokensTransferableByUser() bool {
	if o == nil || isNil(o.TokensTransferableByUser) {
		var ret bool
		return ret
	}
	return *o.TokensTransferableByUser
}

// GetTokensTransferableByUserOk returns a tuple with the TokensTransferableByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetTokensTransferableByUserOk() (*bool, bool) {
	if o == nil || isNil(o.TokensTransferableByUser) {
    return nil, false
	}
	return o.TokensTransferableByUser, true
}

// HasTokensTransferableByUser returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasTokensTransferableByUser() bool {
	if o != nil && !isNil(o.TokensTransferableByUser) {
		return true
	}

	return false
}

// SetTokensTransferableByUser gets a reference to the given bool and assigns it to the TokensTransferableByUser field.
func (o *ServicesContractDeployDto) SetTokensTransferableByUser(v bool) {
	o.TokensTransferableByUser = &v
}

// GetTransferCooldownTime returns the TransferCooldownTime field value if set, zero value otherwise.
func (o *ServicesContractDeployDto) GetTransferCooldownTime() int32 {
	if o == nil || isNil(o.TransferCooldownTime) {
		var ret int32
		return ret
	}
	return *o.TransferCooldownTime
}

// GetTransferCooldownTimeOk returns a tuple with the TransferCooldownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetTransferCooldownTimeOk() (*int32, bool) {
	if o == nil || isNil(o.TransferCooldownTime) {
    return nil, false
	}
	return o.TransferCooldownTime, true
}

// HasTransferCooldownTime returns a boolean if a field has been set.
func (o *ServicesContractDeployDto) HasTransferCooldownTime() bool {
	if o != nil && !isNil(o.TransferCooldownTime) {
		return true
	}

	return false
}

// SetTransferCooldownTime gets a reference to the given int32 and assigns it to the TransferCooldownTime field.
func (o *ServicesContractDeployDto) SetTransferCooldownTime(v int32) {
	o.TransferCooldownTime = &v
}

// GetType returns the Type field value
func (o *ServicesContractDeployDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServicesContractDeployDto) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServicesContractDeployDto) SetType(v string) {
	o.Type = v
}

func (o ServicesContractDeployDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutoSponsor) {
		toSerialize["auto_sponsor"] = o.AutoSponsor
	}
	if !isNil(o.BaseUri) {
		toSerialize["base_uri"] = o.BaseUri
	}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if !isNil(o.IsSponsorForAllUser) {
		toSerialize["is_sponsor_for_all_user"] = o.IsSponsorForAllUser
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OwnerAddress) {
		toSerialize["owner_address"] = o.OwnerAddress
	}
	if !isNil(o.RoyaltiesAddress) {
		toSerialize["royalties_address"] = o.RoyaltiesAddress
	}
	if !isNil(o.RoyaltiesBps) {
		toSerialize["royalties_bps"] = o.RoyaltiesBps
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if !isNil(o.TokensBurnable) {
		toSerialize["tokens_burnable"] = o.TokensBurnable
	}
	if !isNil(o.TokensTransferableByAdmin) {
		toSerialize["tokens_transferable_by_admin"] = o.TokensTransferableByAdmin
	}
	if !isNil(o.TokensTransferableByUser) {
		toSerialize["tokens_transferable_by_user"] = o.TokensTransferableByUser
	}
	if !isNil(o.TransferCooldownTime) {
		toSerialize["transfer_cooldown_time"] = o.TransferCooldownTime
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesContractDeployDto) UnmarshalJSON(bytes []byte) (err error) {
	varServicesContractDeployDto := _ServicesContractDeployDto{}

	if err = json.Unmarshal(bytes, &varServicesContractDeployDto); err == nil {
		*o = ServicesContractDeployDto(varServicesContractDeployDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_sponsor")
		delete(additionalProperties, "base_uri")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "is_sponsor_for_all_user")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_address")
		delete(additionalProperties, "royalties_address")
		delete(additionalProperties, "royalties_bps")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "tokens_burnable")
		delete(additionalProperties, "tokens_transferable_by_admin")
		delete(additionalProperties, "tokens_transferable_by_user")
		delete(additionalProperties, "transfer_cooldown_time")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesContractDeployDto struct {
	value *ServicesContractDeployDto
	isSet bool
}

func (v NullableServicesContractDeployDto) Get() *ServicesContractDeployDto {
	return v.value
}

func (v *NullableServicesContractDeployDto) Set(val *ServicesContractDeployDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesContractDeployDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesContractDeployDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesContractDeployDto(val *ServicesContractDeployDto) *NullableServicesContractDeployDto {
	return &NullableServicesContractDeployDto{value: val, isSet: true}
}

func (v NullableServicesContractDeployDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesContractDeployDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


