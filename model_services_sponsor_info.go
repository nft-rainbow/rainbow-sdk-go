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

// checks if the ServicesSponsorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesSponsorInfo{}

// ServicesSponsorInfo struct for ServicesSponsorInfo
type ServicesSponsorInfo struct {
	CollateralSponsor *string `json:"collateral_sponsor,omitempty"`
	CollateralSponsorBalance *string `json:"collateral_sponsor_balance,omitempty"`
	GasSponsor *string `json:"gas_sponsor,omitempty"`
	GasSponsorBalance *string `json:"gas_sponsor_balance,omitempty"`
	GasUpperBound *string `json:"gas_upper_bound,omitempty"`
	IsAllWhiteListed *bool `json:"is_all_white_listed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesSponsorInfo ServicesSponsorInfo

// NewServicesSponsorInfo instantiates a new ServicesSponsorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesSponsorInfo() *ServicesSponsorInfo {
	this := ServicesSponsorInfo{}
	return &this
}

// NewServicesSponsorInfoWithDefaults instantiates a new ServicesSponsorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesSponsorInfoWithDefaults() *ServicesSponsorInfo {
	this := ServicesSponsorInfo{}
	return &this
}

// GetCollateralSponsor returns the CollateralSponsor field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetCollateralSponsor() string {
	if o == nil || isNil(o.CollateralSponsor) {
		var ret string
		return ret
	}
	return *o.CollateralSponsor
}

// GetCollateralSponsorOk returns a tuple with the CollateralSponsor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetCollateralSponsorOk() (*string, bool) {
	if o == nil || isNil(o.CollateralSponsor) {
		return nil, false
	}
	return o.CollateralSponsor, true
}

// HasCollateralSponsor returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasCollateralSponsor() bool {
	if o != nil && !isNil(o.CollateralSponsor) {
		return true
	}

	return false
}

// SetCollateralSponsor gets a reference to the given string and assigns it to the CollateralSponsor field.
func (o *ServicesSponsorInfo) SetCollateralSponsor(v string) {
	o.CollateralSponsor = &v
}

// GetCollateralSponsorBalance returns the CollateralSponsorBalance field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetCollateralSponsorBalance() string {
	if o == nil || isNil(o.CollateralSponsorBalance) {
		var ret string
		return ret
	}
	return *o.CollateralSponsorBalance
}

// GetCollateralSponsorBalanceOk returns a tuple with the CollateralSponsorBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetCollateralSponsorBalanceOk() (*string, bool) {
	if o == nil || isNil(o.CollateralSponsorBalance) {
		return nil, false
	}
	return o.CollateralSponsorBalance, true
}

// HasCollateralSponsorBalance returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasCollateralSponsorBalance() bool {
	if o != nil && !isNil(o.CollateralSponsorBalance) {
		return true
	}

	return false
}

// SetCollateralSponsorBalance gets a reference to the given string and assigns it to the CollateralSponsorBalance field.
func (o *ServicesSponsorInfo) SetCollateralSponsorBalance(v string) {
	o.CollateralSponsorBalance = &v
}

// GetGasSponsor returns the GasSponsor field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetGasSponsor() string {
	if o == nil || isNil(o.GasSponsor) {
		var ret string
		return ret
	}
	return *o.GasSponsor
}

// GetGasSponsorOk returns a tuple with the GasSponsor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetGasSponsorOk() (*string, bool) {
	if o == nil || isNil(o.GasSponsor) {
		return nil, false
	}
	return o.GasSponsor, true
}

// HasGasSponsor returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasGasSponsor() bool {
	if o != nil && !isNil(o.GasSponsor) {
		return true
	}

	return false
}

// SetGasSponsor gets a reference to the given string and assigns it to the GasSponsor field.
func (o *ServicesSponsorInfo) SetGasSponsor(v string) {
	o.GasSponsor = &v
}

// GetGasSponsorBalance returns the GasSponsorBalance field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetGasSponsorBalance() string {
	if o == nil || isNil(o.GasSponsorBalance) {
		var ret string
		return ret
	}
	return *o.GasSponsorBalance
}

// GetGasSponsorBalanceOk returns a tuple with the GasSponsorBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetGasSponsorBalanceOk() (*string, bool) {
	if o == nil || isNil(o.GasSponsorBalance) {
		return nil, false
	}
	return o.GasSponsorBalance, true
}

// HasGasSponsorBalance returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasGasSponsorBalance() bool {
	if o != nil && !isNil(o.GasSponsorBalance) {
		return true
	}

	return false
}

// SetGasSponsorBalance gets a reference to the given string and assigns it to the GasSponsorBalance field.
func (o *ServicesSponsorInfo) SetGasSponsorBalance(v string) {
	o.GasSponsorBalance = &v
}

// GetGasUpperBound returns the GasUpperBound field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetGasUpperBound() string {
	if o == nil || isNil(o.GasUpperBound) {
		var ret string
		return ret
	}
	return *o.GasUpperBound
}

// GetGasUpperBoundOk returns a tuple with the GasUpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetGasUpperBoundOk() (*string, bool) {
	if o == nil || isNil(o.GasUpperBound) {
		return nil, false
	}
	return o.GasUpperBound, true
}

// HasGasUpperBound returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasGasUpperBound() bool {
	if o != nil && !isNil(o.GasUpperBound) {
		return true
	}

	return false
}

// SetGasUpperBound gets a reference to the given string and assigns it to the GasUpperBound field.
func (o *ServicesSponsorInfo) SetGasUpperBound(v string) {
	o.GasUpperBound = &v
}

// GetIsAllWhiteListed returns the IsAllWhiteListed field value if set, zero value otherwise.
func (o *ServicesSponsorInfo) GetIsAllWhiteListed() bool {
	if o == nil || isNil(o.IsAllWhiteListed) {
		var ret bool
		return ret
	}
	return *o.IsAllWhiteListed
}

// GetIsAllWhiteListedOk returns a tuple with the IsAllWhiteListed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesSponsorInfo) GetIsAllWhiteListedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAllWhiteListed) {
		return nil, false
	}
	return o.IsAllWhiteListed, true
}

// HasIsAllWhiteListed returns a boolean if a field has been set.
func (o *ServicesSponsorInfo) HasIsAllWhiteListed() bool {
	if o != nil && !isNil(o.IsAllWhiteListed) {
		return true
	}

	return false
}

// SetIsAllWhiteListed gets a reference to the given bool and assigns it to the IsAllWhiteListed field.
func (o *ServicesSponsorInfo) SetIsAllWhiteListed(v bool) {
	o.IsAllWhiteListed = &v
}

func (o ServicesSponsorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesSponsorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CollateralSponsor) {
		toSerialize["collateral_sponsor"] = o.CollateralSponsor
	}
	if !isNil(o.CollateralSponsorBalance) {
		toSerialize["collateral_sponsor_balance"] = o.CollateralSponsorBalance
	}
	if !isNil(o.GasSponsor) {
		toSerialize["gas_sponsor"] = o.GasSponsor
	}
	if !isNil(o.GasSponsorBalance) {
		toSerialize["gas_sponsor_balance"] = o.GasSponsorBalance
	}
	if !isNil(o.GasUpperBound) {
		toSerialize["gas_upper_bound"] = o.GasUpperBound
	}
	if !isNil(o.IsAllWhiteListed) {
		toSerialize["is_all_white_listed"] = o.IsAllWhiteListed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesSponsorInfo) UnmarshalJSON(bytes []byte) (err error) {
	varServicesSponsorInfo := _ServicesSponsorInfo{}

	if err = json.Unmarshal(bytes, &varServicesSponsorInfo); err == nil {
		*o = ServicesSponsorInfo(varServicesSponsorInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "collateral_sponsor")
		delete(additionalProperties, "collateral_sponsor_balance")
		delete(additionalProperties, "gas_sponsor")
		delete(additionalProperties, "gas_sponsor_balance")
		delete(additionalProperties, "gas_upper_bound")
		delete(additionalProperties, "is_all_white_listed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesSponsorInfo struct {
	value *ServicesSponsorInfo
	isSet bool
}

func (v NullableServicesSponsorInfo) Get() *ServicesSponsorInfo {
	return v.value
}

func (v *NullableServicesSponsorInfo) Set(val *ServicesSponsorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesSponsorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesSponsorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesSponsorInfo(val *ServicesSponsorInfo) *NullableServicesSponsorInfo {
	return &NullableServicesSponsorInfo{value: val, isSet: true}
}

func (v NullableServicesSponsorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesSponsorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


