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

// ServicesMintItemDto struct for ServicesMintItemDto
type ServicesMintItemDto struct {
	Amount *int32 `json:"amount,omitempty"`
	MetadataUri string `json:"metadata_uri"`
	MintToAddress string `json:"mint_to_address"`
	TokenId *string `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMintItemDto ServicesMintItemDto

// NewServicesMintItemDto instantiates a new ServicesMintItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMintItemDto(metadataUri string, mintToAddress string) *ServicesMintItemDto {
	this := ServicesMintItemDto{}
	this.MetadataUri = metadataUri
	this.MintToAddress = mintToAddress
	return &this
}

// NewServicesMintItemDtoWithDefaults instantiates a new ServicesMintItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMintItemDtoWithDefaults() *ServicesMintItemDto {
	this := ServicesMintItemDto{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ServicesMintItemDto) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMintItemDto) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ServicesMintItemDto) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ServicesMintItemDto) SetAmount(v int32) {
	o.Amount = &v
}

// GetMetadataUri returns the MetadataUri field value
func (o *ServicesMintItemDto) GetMetadataUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataUri
}

// GetMetadataUriOk returns a tuple with the MetadataUri field value
// and a boolean to check if the value has been set.
func (o *ServicesMintItemDto) GetMetadataUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataUri, true
}

// SetMetadataUri sets field value
func (o *ServicesMintItemDto) SetMetadataUri(v string) {
	o.MetadataUri = v
}

// GetMintToAddress returns the MintToAddress field value
func (o *ServicesMintItemDto) GetMintToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MintToAddress
}

// GetMintToAddressOk returns a tuple with the MintToAddress field value
// and a boolean to check if the value has been set.
func (o *ServicesMintItemDto) GetMintToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MintToAddress, true
}

// SetMintToAddress sets field value
func (o *ServicesMintItemDto) SetMintToAddress(v string) {
	o.MintToAddress = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ServicesMintItemDto) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMintItemDto) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ServicesMintItemDto) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ServicesMintItemDto) SetTokenId(v string) {
	o.TokenId = &v
}

func (o ServicesMintItemDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["metadata_uri"] = o.MetadataUri
	}
	if true {
		toSerialize["mint_to_address"] = o.MintToAddress
	}
	if o.TokenId != nil {
		toSerialize["token_id"] = o.TokenId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesMintItemDto) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMintItemDto := _ServicesMintItemDto{}

	if err = json.Unmarshal(bytes, &varServicesMintItemDto); err == nil {
		*o = ServicesMintItemDto(varServicesMintItemDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "metadata_uri")
		delete(additionalProperties, "mint_to_address")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMintItemDto struct {
	value *ServicesMintItemDto
	isSet bool
}

func (v NullableServicesMintItemDto) Get() *ServicesMintItemDto {
	return v.value
}

func (v *NullableServicesMintItemDto) Set(val *ServicesMintItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMintItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMintItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMintItemDto(val *ServicesMintItemDto) *NullableServicesMintItemDto {
	return &NullableServicesMintItemDto{value: val, isSet: true}
}

func (v NullableServicesMintItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMintItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


