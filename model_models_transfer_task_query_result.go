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

// ModelsTransferTaskQueryResult struct for ModelsTransferTaskQueryResult
type ModelsTransferTaskQueryResult struct {
	Count *int32 `json:"count,omitempty"`
	Items []ModelsTransferTask `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsTransferTaskQueryResult ModelsTransferTaskQueryResult

// NewModelsTransferTaskQueryResult instantiates a new ModelsTransferTaskQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTransferTaskQueryResult() *ModelsTransferTaskQueryResult {
	this := ModelsTransferTaskQueryResult{}
	return &this
}

// NewModelsTransferTaskQueryResultWithDefaults instantiates a new ModelsTransferTaskQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTransferTaskQueryResultWithDefaults() *ModelsTransferTaskQueryResult {
	this := ModelsTransferTaskQueryResult{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelsTransferTaskQueryResult) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTaskQueryResult) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelsTransferTaskQueryResult) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelsTransferTaskQueryResult) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelsTransferTaskQueryResult) GetItems() []ModelsTransferTask {
	if o == nil || isNil(o.Items) {
		var ret []ModelsTransferTask
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTransferTaskQueryResult) GetItemsOk() ([]ModelsTransferTask, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelsTransferTaskQueryResult) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelsTransferTask and assigns it to the Items field.
func (o *ModelsTransferTaskQueryResult) SetItems(v []ModelsTransferTask) {
	o.Items = v
}

func (o ModelsTransferTaskQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsTransferTaskQueryResult) UnmarshalJSON(bytes []byte) (err error) {
	varModelsTransferTaskQueryResult := _ModelsTransferTaskQueryResult{}

	if err = json.Unmarshal(bytes, &varModelsTransferTaskQueryResult); err == nil {
		*o = ModelsTransferTaskQueryResult(varModelsTransferTaskQueryResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsTransferTaskQueryResult struct {
	value *ModelsTransferTaskQueryResult
	isSet bool
}

func (v NullableModelsTransferTaskQueryResult) Get() *ModelsTransferTaskQueryResult {
	return v.value
}

func (v *NullableModelsTransferTaskQueryResult) Set(val *ModelsTransferTaskQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTransferTaskQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTransferTaskQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTransferTaskQueryResult(val *ModelsTransferTaskQueryResult) *NullableModelsTransferTaskQueryResult {
	return &NullableModelsTransferTaskQueryResult{value: val, isSet: true}
}

func (v NullableModelsTransferTaskQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTransferTaskQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


