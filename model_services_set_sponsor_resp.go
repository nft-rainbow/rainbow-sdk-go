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

// checks if the ServicesSetSponsorResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesSetSponsorResp{}

// ServicesSetSponsorResp struct for ServicesSetSponsorResp
type ServicesSetSponsorResp struct {
	SponsorCollateralTxId *int32 `json:"sponsor_collateral_tx_id,omitempty"`
	SponsorGasTxId *int32 `json:"sponsor_gas_tx_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesSetSponsorResp ServicesSetSponsorResp

// NewServicesSetSponsorResp instantiates a new ServicesSetSponsorResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesSetSponsorResp() *ServicesSetSponsorResp {
	this := ServicesSetSponsorResp{}
	return &this
}

// NewServicesSetSponsorRespWithDefaults instantiates a new ServicesSetSponsorResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesSetSponsorRespWithDefaults() *ServicesSetSponsorResp {
	this := ServicesSetSponsorResp{}
	return &this
}

// GetSponsorCollateralTxId returns the SponsorCollateralTxId field value if set, zero value otherwise.
func (o *ServicesSetSponsorResp) GetSponsorCollateralTxId() int32 {
	if o == nil || IsNil(o.SponsorCollateralTxId) {
		var ret int32
		return ret
	}
	return *o.SponsorCollateralTxId
}

// GetSponsorCollateralTxIdOk returns a tuple with the SponsorCollateralTxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSetSponsorResp) GetSponsorCollateralTxIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SponsorCollateralTxId) {
		return nil, false
	}
	return o.SponsorCollateralTxId, true
}

// HasSponsorCollateralTxId returns a boolean if a field has been set.
func (o *ServicesSetSponsorResp) HasSponsorCollateralTxId() bool {
	if o != nil && !IsNil(o.SponsorCollateralTxId) {
		return true
	}

	return false
}

// SetSponsorCollateralTxId gets a reference to the given int32 and assigns it to the SponsorCollateralTxId field.
func (o *ServicesSetSponsorResp) SetSponsorCollateralTxId(v int32) {
	o.SponsorCollateralTxId = &v
}

// GetSponsorGasTxId returns the SponsorGasTxId field value if set, zero value otherwise.
func (o *ServicesSetSponsorResp) GetSponsorGasTxId() int32 {
	if o == nil || IsNil(o.SponsorGasTxId) {
		var ret int32
		return ret
	}
	return *o.SponsorGasTxId
}

// GetSponsorGasTxIdOk returns a tuple with the SponsorGasTxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSetSponsorResp) GetSponsorGasTxIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SponsorGasTxId) {
		return nil, false
	}
	return o.SponsorGasTxId, true
}

// HasSponsorGasTxId returns a boolean if a field has been set.
func (o *ServicesSetSponsorResp) HasSponsorGasTxId() bool {
	if o != nil && !IsNil(o.SponsorGasTxId) {
		return true
	}

	return false
}

// SetSponsorGasTxId gets a reference to the given int32 and assigns it to the SponsorGasTxId field.
func (o *ServicesSetSponsorResp) SetSponsorGasTxId(v int32) {
	o.SponsorGasTxId = &v
}

func (o ServicesSetSponsorResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesSetSponsorResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SponsorCollateralTxId) {
		toSerialize["sponsor_collateral_tx_id"] = o.SponsorCollateralTxId
	}
	if !IsNil(o.SponsorGasTxId) {
		toSerialize["sponsor_gas_tx_id"] = o.SponsorGasTxId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesSetSponsorResp) UnmarshalJSON(data []byte) (err error) {
	varServicesSetSponsorResp := _ServicesSetSponsorResp{}

	err = json.Unmarshal(data, &varServicesSetSponsorResp)

	if err != nil {
		return err
	}

	*o = ServicesSetSponsorResp(varServicesSetSponsorResp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sponsor_collateral_tx_id")
		delete(additionalProperties, "sponsor_gas_tx_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesSetSponsorResp struct {
	value *ServicesSetSponsorResp
	isSet bool
}

func (v NullableServicesSetSponsorResp) Get() *ServicesSetSponsorResp {
	return v.value
}

func (v *NullableServicesSetSponsorResp) Set(val *ServicesSetSponsorResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesSetSponsorResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesSetSponsorResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesSetSponsorResp(val *ServicesSetSponsorResp) *NullableServicesSetSponsorResp {
	return &NullableServicesSetSponsorResp{value: val, isSet: true}
}

func (v NullableServicesSetSponsorResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesSetSponsorResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


