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

// ModelsExposedMetadataQueryResult struct for ModelsExposedMetadataQueryResult
type ModelsExposedMetadataQueryResult struct {
	Count *int32 `json:"count,omitempty"`
	Items []ModelsExposedMetadata `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsExposedMetadataQueryResult ModelsExposedMetadataQueryResult

// NewModelsExposedMetadataQueryResult instantiates a new ModelsExposedMetadataQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsExposedMetadataQueryResult() *ModelsExposedMetadataQueryResult {
	this := ModelsExposedMetadataQueryResult{}
	return &this
}

// NewModelsExposedMetadataQueryResultWithDefaults instantiates a new ModelsExposedMetadataQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsExposedMetadataQueryResultWithDefaults() *ModelsExposedMetadataQueryResult {
	this := ModelsExposedMetadataQueryResult{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelsExposedMetadataQueryResult) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataQueryResult) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelsExposedMetadataQueryResult) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelsExposedMetadataQueryResult) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelsExposedMetadataQueryResult) GetItems() []ModelsExposedMetadata {
	if o == nil || o.Items == nil {
		var ret []ModelsExposedMetadata
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataQueryResult) GetItemsOk() ([]ModelsExposedMetadata, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelsExposedMetadataQueryResult) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelsExposedMetadata and assigns it to the Items field.
func (o *ModelsExposedMetadataQueryResult) SetItems(v []ModelsExposedMetadata) {
	o.Items = v
}

func (o ModelsExposedMetadataQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsExposedMetadataQueryResult) UnmarshalJSON(bytes []byte) (err error) {
	varModelsExposedMetadataQueryResult := _ModelsExposedMetadataQueryResult{}

	if err = json.Unmarshal(bytes, &varModelsExposedMetadataQueryResult); err == nil {
		*o = ModelsExposedMetadataQueryResult(varModelsExposedMetadataQueryResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsExposedMetadataQueryResult struct {
	value *ModelsExposedMetadataQueryResult
	isSet bool
}

func (v NullableModelsExposedMetadataQueryResult) Get() *ModelsExposedMetadataQueryResult {
	return v.value
}

func (v *NullableModelsExposedMetadataQueryResult) Set(val *ModelsExposedMetadataQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsExposedMetadataQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsExposedMetadataQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsExposedMetadataQueryResult(val *ModelsExposedMetadataQueryResult) *NullableModelsExposedMetadataQueryResult {
	return &NullableModelsExposedMetadataQueryResult{value: val, isSet: true}
}

func (v NullableModelsExposedMetadataQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsExposedMetadataQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


