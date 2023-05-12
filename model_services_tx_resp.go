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

// checks if the ServicesTxResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesTxResp{}

// ServicesTxResp struct for ServicesTxResp
type ServicesTxResp struct {
	ErrorMsg *string `json:"error_msg,omitempty"`
	Hash *string `json:"hash,omitempty"`
	IsFinalized *bool `json:"is_finalized,omitempty"`
	IsSuccess *bool `json:"is_success,omitempty"`
	StateCode *ModelsTxState `json:"state_code,omitempty"`
	StateMsg *string `json:"state_msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesTxResp ServicesTxResp

// NewServicesTxResp instantiates a new ServicesTxResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesTxResp() *ServicesTxResp {
	this := ServicesTxResp{}
	return &this
}

// NewServicesTxRespWithDefaults instantiates a new ServicesTxResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesTxRespWithDefaults() *ServicesTxResp {
	this := ServicesTxResp{}
	return &this
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *ServicesTxResp) GetErrorMsg() string {
	if o == nil || isNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetErrorMsgOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *ServicesTxResp) HasErrorMsg() bool {
	if o != nil && !isNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *ServicesTxResp) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ServicesTxResp) GetHash() string {
	if o == nil || isNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetHashOk() (*string, bool) {
	if o == nil || isNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ServicesTxResp) HasHash() bool {
	if o != nil && !isNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ServicesTxResp) SetHash(v string) {
	o.Hash = &v
}

// GetIsFinalized returns the IsFinalized field value if set, zero value otherwise.
func (o *ServicesTxResp) GetIsFinalized() bool {
	if o == nil || isNil(o.IsFinalized) {
		var ret bool
		return ret
	}
	return *o.IsFinalized
}

// GetIsFinalizedOk returns a tuple with the IsFinalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetIsFinalizedOk() (*bool, bool) {
	if o == nil || isNil(o.IsFinalized) {
		return nil, false
	}
	return o.IsFinalized, true
}

// HasIsFinalized returns a boolean if a field has been set.
func (o *ServicesTxResp) HasIsFinalized() bool {
	if o != nil && !isNil(o.IsFinalized) {
		return true
	}

	return false
}

// SetIsFinalized gets a reference to the given bool and assigns it to the IsFinalized field.
func (o *ServicesTxResp) SetIsFinalized(v bool) {
	o.IsFinalized = &v
}

// GetIsSuccess returns the IsSuccess field value if set, zero value otherwise.
func (o *ServicesTxResp) GetIsSuccess() bool {
	if o == nil || isNil(o.IsSuccess) {
		var ret bool
		return ret
	}
	return *o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetIsSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.IsSuccess) {
		return nil, false
	}
	return o.IsSuccess, true
}

// HasIsSuccess returns a boolean if a field has been set.
func (o *ServicesTxResp) HasIsSuccess() bool {
	if o != nil && !isNil(o.IsSuccess) {
		return true
	}

	return false
}

// SetIsSuccess gets a reference to the given bool and assigns it to the IsSuccess field.
func (o *ServicesTxResp) SetIsSuccess(v bool) {
	o.IsSuccess = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *ServicesTxResp) GetStateCode() ModelsTxState {
	if o == nil || isNil(o.StateCode) {
		var ret ModelsTxState
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetStateCodeOk() (*ModelsTxState, bool) {
	if o == nil || isNil(o.StateCode) {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *ServicesTxResp) HasStateCode() bool {
	if o != nil && !isNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given ModelsTxState and assigns it to the StateCode field.
func (o *ServicesTxResp) SetStateCode(v ModelsTxState) {
	o.StateCode = &v
}

// GetStateMsg returns the StateMsg field value if set, zero value otherwise.
func (o *ServicesTxResp) GetStateMsg() string {
	if o == nil || isNil(o.StateMsg) {
		var ret string
		return ret
	}
	return *o.StateMsg
}

// GetStateMsgOk returns a tuple with the StateMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesTxResp) GetStateMsgOk() (*string, bool) {
	if o == nil || isNil(o.StateMsg) {
		return nil, false
	}
	return o.StateMsg, true
}

// HasStateMsg returns a boolean if a field has been set.
func (o *ServicesTxResp) HasStateMsg() bool {
	if o != nil && !isNil(o.StateMsg) {
		return true
	}

	return false
}

// SetStateMsg gets a reference to the given string and assigns it to the StateMsg field.
func (o *ServicesTxResp) SetStateMsg(v string) {
	o.StateMsg = &v
}

func (o ServicesTxResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesTxResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ErrorMsg) {
		toSerialize["error_msg"] = o.ErrorMsg
	}
	if !isNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !isNil(o.IsFinalized) {
		toSerialize["is_finalized"] = o.IsFinalized
	}
	if !isNil(o.IsSuccess) {
		toSerialize["is_success"] = o.IsSuccess
	}
	if !isNil(o.StateCode) {
		toSerialize["state_code"] = o.StateCode
	}
	if !isNil(o.StateMsg) {
		toSerialize["state_msg"] = o.StateMsg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesTxResp) UnmarshalJSON(bytes []byte) (err error) {
	varServicesTxResp := _ServicesTxResp{}

	if err = json.Unmarshal(bytes, &varServicesTxResp); err == nil {
		*o = ServicesTxResp(varServicesTxResp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error_msg")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "is_finalized")
		delete(additionalProperties, "is_success")
		delete(additionalProperties, "state_code")
		delete(additionalProperties, "state_msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesTxResp struct {
	value *ServicesTxResp
	isSet bool
}

func (v NullableServicesTxResp) Get() *ServicesTxResp {
	return v.value
}

func (v *NullableServicesTxResp) Set(val *ServicesTxResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesTxResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesTxResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesTxResp(val *ServicesTxResp) *NullableServicesTxResp {
	return &NullableServicesTxResp{value: val, isSet: true}
}

func (v NullableServicesTxResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesTxResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


