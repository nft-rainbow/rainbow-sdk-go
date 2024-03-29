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

// checks if the ModelsContractTaskQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsContractTaskQueryResult{}

// ModelsContractTaskQueryResult struct for ModelsContractTaskQueryResult
type ModelsContractTaskQueryResult struct {
	Count *int32 `json:"count,omitempty"`
	Items []ModelsContract `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsContractTaskQueryResult ModelsContractTaskQueryResult

// NewModelsContractTaskQueryResult instantiates a new ModelsContractTaskQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsContractTaskQueryResult() *ModelsContractTaskQueryResult {
	this := ModelsContractTaskQueryResult{}
	return &this
}

// NewModelsContractTaskQueryResultWithDefaults instantiates a new ModelsContractTaskQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsContractTaskQueryResultWithDefaults() *ModelsContractTaskQueryResult {
	this := ModelsContractTaskQueryResult{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelsContractTaskQueryResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContractTaskQueryResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelsContractTaskQueryResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelsContractTaskQueryResult) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelsContractTaskQueryResult) GetItems() []ModelsContract {
	if o == nil || IsNil(o.Items) {
		var ret []ModelsContract
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsContractTaskQueryResult) GetItemsOk() ([]ModelsContract, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelsContractTaskQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelsContract and assigns it to the Items field.
func (o *ModelsContractTaskQueryResult) SetItems(v []ModelsContract) {
	o.Items = v
}

func (o ModelsContractTaskQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsContractTaskQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsContractTaskQueryResult) UnmarshalJSON(data []byte) (err error) {
	varModelsContractTaskQueryResult := _ModelsContractTaskQueryResult{}

	err = json.Unmarshal(data, &varModelsContractTaskQueryResult)

	if err != nil {
		return err
	}

	*o = ModelsContractTaskQueryResult(varModelsContractTaskQueryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsContractTaskQueryResult struct {
	value *ModelsContractTaskQueryResult
	isSet bool
}

func (v NullableModelsContractTaskQueryResult) Get() *ModelsContractTaskQueryResult {
	return v.value
}

func (v *NullableModelsContractTaskQueryResult) Set(val *ModelsContractTaskQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsContractTaskQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsContractTaskQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsContractTaskQueryResult(val *ModelsContractTaskQueryResult) *NullableModelsContractTaskQueryResult {
	return &NullableModelsContractTaskQueryResult{value: val, isSet: true}
}

func (v NullableModelsContractTaskQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsContractTaskQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


