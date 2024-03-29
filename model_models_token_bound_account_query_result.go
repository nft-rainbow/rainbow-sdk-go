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

// checks if the ModelsTokenBoundAccountQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsTokenBoundAccountQueryResult{}

// ModelsTokenBoundAccountQueryResult struct for ModelsTokenBoundAccountQueryResult
type ModelsTokenBoundAccountQueryResult struct {
	Count *int32 `json:"count,omitempty"`
	Items []ModelsTokenBoundAccount `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsTokenBoundAccountQueryResult ModelsTokenBoundAccountQueryResult

// NewModelsTokenBoundAccountQueryResult instantiates a new ModelsTokenBoundAccountQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTokenBoundAccountQueryResult() *ModelsTokenBoundAccountQueryResult {
	this := ModelsTokenBoundAccountQueryResult{}
	return &this
}

// NewModelsTokenBoundAccountQueryResultWithDefaults instantiates a new ModelsTokenBoundAccountQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTokenBoundAccountQueryResultWithDefaults() *ModelsTokenBoundAccountQueryResult {
	this := ModelsTokenBoundAccountQueryResult{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelsTokenBoundAccountQueryResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTokenBoundAccountQueryResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelsTokenBoundAccountQueryResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelsTokenBoundAccountQueryResult) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelsTokenBoundAccountQueryResult) GetItems() []ModelsTokenBoundAccount {
	if o == nil || IsNil(o.Items) {
		var ret []ModelsTokenBoundAccount
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTokenBoundAccountQueryResult) GetItemsOk() ([]ModelsTokenBoundAccount, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelsTokenBoundAccountQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelsTokenBoundAccount and assigns it to the Items field.
func (o *ModelsTokenBoundAccountQueryResult) SetItems(v []ModelsTokenBoundAccount) {
	o.Items = v
}

func (o ModelsTokenBoundAccountQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsTokenBoundAccountQueryResult) ToMap() (map[string]interface{}, error) {
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

func (o *ModelsTokenBoundAccountQueryResult) UnmarshalJSON(data []byte) (err error) {
	varModelsTokenBoundAccountQueryResult := _ModelsTokenBoundAccountQueryResult{}

	err = json.Unmarshal(data, &varModelsTokenBoundAccountQueryResult)

	if err != nil {
		return err
	}

	*o = ModelsTokenBoundAccountQueryResult(varModelsTokenBoundAccountQueryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsTokenBoundAccountQueryResult struct {
	value *ModelsTokenBoundAccountQueryResult
	isSet bool
}

func (v NullableModelsTokenBoundAccountQueryResult) Get() *ModelsTokenBoundAccountQueryResult {
	return v.value
}

func (v *NullableModelsTokenBoundAccountQueryResult) Set(val *ModelsTokenBoundAccountQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTokenBoundAccountQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTokenBoundAccountQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTokenBoundAccountQueryResult(val *ModelsTokenBoundAccountQueryResult) *NullableModelsTokenBoundAccountQueryResult {
	return &NullableModelsTokenBoundAccountQueryResult{value: val, isSet: true}
}

func (v NullableModelsTokenBoundAccountQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTokenBoundAccountQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


