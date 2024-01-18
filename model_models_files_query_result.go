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

// checks if the ModelsFilesQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsFilesQueryResult{}

// ModelsFilesQueryResult struct for ModelsFilesQueryResult
type ModelsFilesQueryResult struct {
	Count *int32 `json:"count,omitempty"`
	Items []ModelsExposedFile `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsFilesQueryResult ModelsFilesQueryResult

// NewModelsFilesQueryResult instantiates a new ModelsFilesQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsFilesQueryResult() *ModelsFilesQueryResult {
	this := ModelsFilesQueryResult{}
	return &this
}

// NewModelsFilesQueryResultWithDefaults instantiates a new ModelsFilesQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsFilesQueryResultWithDefaults() *ModelsFilesQueryResult {
	this := ModelsFilesQueryResult{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelsFilesQueryResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsFilesQueryResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelsFilesQueryResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelsFilesQueryResult) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelsFilesQueryResult) GetItems() []ModelsExposedFile {
	if o == nil || IsNil(o.Items) {
		var ret []ModelsExposedFile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsFilesQueryResult) GetItemsOk() ([]ModelsExposedFile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelsFilesQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelsExposedFile and assigns it to the Items field.
func (o *ModelsFilesQueryResult) SetItems(v []ModelsExposedFile) {
	o.Items = v
}

func (o ModelsFilesQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsFilesQueryResult) ToMap() (map[string]interface{}, error) {
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

func (o *ModelsFilesQueryResult) UnmarshalJSON(data []byte) (err error) {
	varModelsFilesQueryResult := _ModelsFilesQueryResult{}

	err = json.Unmarshal(data, &varModelsFilesQueryResult)

	if err != nil {
		return err
	}

	*o = ModelsFilesQueryResult(varModelsFilesQueryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsFilesQueryResult struct {
	value *ModelsFilesQueryResult
	isSet bool
}

func (v NullableModelsFilesQueryResult) Get() *ModelsFilesQueryResult {
	return v.value
}

func (v *NullableModelsFilesQueryResult) Set(val *ModelsFilesQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsFilesQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsFilesQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsFilesQueryResult(val *ModelsFilesQueryResult) *NullableModelsFilesQueryResult {
	return &NullableModelsFilesQueryResult{value: val, isSet: true}
}

func (v NullableModelsFilesQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsFilesQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


