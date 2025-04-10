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

// ModelsMintTaskOrError struct for ModelsMintTaskOrError
type ModelsMintTaskOrError struct {
	Error map[string]interface{} `json:"error,omitempty"`
	MintTask *ModelsMintTask `json:"mint_task,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsMintTaskOrError ModelsMintTaskOrError

// NewModelsMintTaskOrError instantiates a new ModelsMintTaskOrError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsMintTaskOrError() *ModelsMintTaskOrError {
	this := ModelsMintTaskOrError{}
	return &this
}

// NewModelsMintTaskOrErrorWithDefaults instantiates a new ModelsMintTaskOrError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsMintTaskOrErrorWithDefaults() *ModelsMintTaskOrError {
	this := ModelsMintTaskOrError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsMintTaskOrError) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTaskOrError) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
    return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsMintTaskOrError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *ModelsMintTaskOrError) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetMintTask returns the MintTask field value if set, zero value otherwise.
func (o *ModelsMintTaskOrError) GetMintTask() ModelsMintTask {
	if o == nil || IsNil(o.MintTask) {
		var ret ModelsMintTask
		return ret
	}
	return *o.MintTask
}

// GetMintTaskOk returns a tuple with the MintTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMintTaskOrError) GetMintTaskOk() (*ModelsMintTask, bool) {
	if o == nil || IsNil(o.MintTask) {
    return nil, false
	}
	return o.MintTask, true
}

// HasMintTask returns a boolean if a field has been set.
func (o *ModelsMintTaskOrError) HasMintTask() bool {
	if o != nil && !IsNil(o.MintTask) {
		return true
	}

	return false
}

// SetMintTask gets a reference to the given ModelsMintTask and assigns it to the MintTask field.
func (o *ModelsMintTaskOrError) SetMintTask(v ModelsMintTask) {
	o.MintTask = &v
}

func (o ModelsMintTaskOrError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.MintTask) {
		toSerialize["mint_task"] = o.MintTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsMintTaskOrError) UnmarshalJSON(bytes []byte) (err error) {
	varModelsMintTaskOrError := _ModelsMintTaskOrError{}

	if err = json.Unmarshal(bytes, &varModelsMintTaskOrError); err == nil {
		*o = ModelsMintTaskOrError(varModelsMintTaskOrError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "mint_task")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsMintTaskOrError struct {
	value *ModelsMintTaskOrError
	isSet bool
}

func (v NullableModelsMintTaskOrError) Get() *ModelsMintTaskOrError {
	return v.value
}

func (v *NullableModelsMintTaskOrError) Set(val *ModelsMintTaskOrError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsMintTaskOrError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsMintTaskOrError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsMintTaskOrError(val *ModelsMintTaskOrError) *NullableModelsMintTaskOrError {
	return &NullableModelsMintTaskOrError{value: val, isSet: true}
}

func (v NullableModelsMintTaskOrError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsMintTaskOrError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


