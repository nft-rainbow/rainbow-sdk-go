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

// checks if the ModelsAccountDisplayPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAccountDisplayPart{}

// ModelsAccountDisplayPart struct for ModelsAccountDisplayPart
type ModelsAccountDisplayPart struct {
	ConfluxMainnetAddress *string `json:"conflux_mainnet_address,omitempty"`
	ConfluxTestnetAddress *string `json:"conflux_testnet_address,omitempty"`
	Phone *string `json:"phone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsAccountDisplayPart ModelsAccountDisplayPart

// NewModelsAccountDisplayPart instantiates a new ModelsAccountDisplayPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAccountDisplayPart() *ModelsAccountDisplayPart {
	this := ModelsAccountDisplayPart{}
	return &this
}

// NewModelsAccountDisplayPartWithDefaults instantiates a new ModelsAccountDisplayPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAccountDisplayPartWithDefaults() *ModelsAccountDisplayPart {
	this := ModelsAccountDisplayPart{}
	return &this
}

// GetConfluxMainnetAddress returns the ConfluxMainnetAddress field value if set, zero value otherwise.
func (o *ModelsAccountDisplayPart) GetConfluxMainnetAddress() string {
	if o == nil || IsNil(o.ConfluxMainnetAddress) {
		var ret string
		return ret
	}
	return *o.ConfluxMainnetAddress
}

// GetConfluxMainnetAddressOk returns a tuple with the ConfluxMainnetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAccountDisplayPart) GetConfluxMainnetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ConfluxMainnetAddress) {
		return nil, false
	}
	return o.ConfluxMainnetAddress, true
}

// HasConfluxMainnetAddress returns a boolean if a field has been set.
func (o *ModelsAccountDisplayPart) HasConfluxMainnetAddress() bool {
	if o != nil && !IsNil(o.ConfluxMainnetAddress) {
		return true
	}

	return false
}

// SetConfluxMainnetAddress gets a reference to the given string and assigns it to the ConfluxMainnetAddress field.
func (o *ModelsAccountDisplayPart) SetConfluxMainnetAddress(v string) {
	o.ConfluxMainnetAddress = &v
}

// GetConfluxTestnetAddress returns the ConfluxTestnetAddress field value if set, zero value otherwise.
func (o *ModelsAccountDisplayPart) GetConfluxTestnetAddress() string {
	if o == nil || IsNil(o.ConfluxTestnetAddress) {
		var ret string
		return ret
	}
	return *o.ConfluxTestnetAddress
}

// GetConfluxTestnetAddressOk returns a tuple with the ConfluxTestnetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAccountDisplayPart) GetConfluxTestnetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ConfluxTestnetAddress) {
		return nil, false
	}
	return o.ConfluxTestnetAddress, true
}

// HasConfluxTestnetAddress returns a boolean if a field has been set.
func (o *ModelsAccountDisplayPart) HasConfluxTestnetAddress() bool {
	if o != nil && !IsNil(o.ConfluxTestnetAddress) {
		return true
	}

	return false
}

// SetConfluxTestnetAddress gets a reference to the given string and assigns it to the ConfluxTestnetAddress field.
func (o *ModelsAccountDisplayPart) SetConfluxTestnetAddress(v string) {
	o.ConfluxTestnetAddress = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ModelsAccountDisplayPart) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAccountDisplayPart) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ModelsAccountDisplayPart) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ModelsAccountDisplayPart) SetPhone(v string) {
	o.Phone = &v
}

func (o ModelsAccountDisplayPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAccountDisplayPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfluxMainnetAddress) {
		toSerialize["conflux_mainnet_address"] = o.ConfluxMainnetAddress
	}
	if !IsNil(o.ConfluxTestnetAddress) {
		toSerialize["conflux_testnet_address"] = o.ConfluxTestnetAddress
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsAccountDisplayPart) UnmarshalJSON(data []byte) (err error) {
	varModelsAccountDisplayPart := _ModelsAccountDisplayPart{}

	err = json.Unmarshal(data, &varModelsAccountDisplayPart)

	if err != nil {
		return err
	}

	*o = ModelsAccountDisplayPart(varModelsAccountDisplayPart)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conflux_mainnet_address")
		delete(additionalProperties, "conflux_testnet_address")
		delete(additionalProperties, "phone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsAccountDisplayPart struct {
	value *ModelsAccountDisplayPart
	isSet bool
}

func (v NullableModelsAccountDisplayPart) Get() *ModelsAccountDisplayPart {
	return v.value
}

func (v *NullableModelsAccountDisplayPart) Set(val *ModelsAccountDisplayPart) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAccountDisplayPart) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAccountDisplayPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAccountDisplayPart(val *ModelsAccountDisplayPart) *NullableModelsAccountDisplayPart {
	return &NullableModelsAccountDisplayPart{value: val, isSet: true}
}

func (v NullableModelsAccountDisplayPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAccountDisplayPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


