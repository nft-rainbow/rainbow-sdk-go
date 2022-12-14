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

// RainbowErrorsRainbowErrorDetailInfo struct for RainbowErrorsRainbowErrorDetailInfo
type RainbowErrorsRainbowErrorDetailInfo struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RainbowErrorsRainbowErrorDetailInfo RainbowErrorsRainbowErrorDetailInfo

// NewRainbowErrorsRainbowErrorDetailInfo instantiates a new RainbowErrorsRainbowErrorDetailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRainbowErrorsRainbowErrorDetailInfo() *RainbowErrorsRainbowErrorDetailInfo {
	this := RainbowErrorsRainbowErrorDetailInfo{}
	return &this
}

// NewRainbowErrorsRainbowErrorDetailInfoWithDefaults instantiates a new RainbowErrorsRainbowErrorDetailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRainbowErrorsRainbowErrorDetailInfoWithDefaults() *RainbowErrorsRainbowErrorDetailInfo {
	this := RainbowErrorsRainbowErrorDetailInfo{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RainbowErrorsRainbowErrorDetailInfo) GetCode() int32 {
	if o == nil || isNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RainbowErrorsRainbowErrorDetailInfo) GetCodeOk() (*int32, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RainbowErrorsRainbowErrorDetailInfo) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *RainbowErrorsRainbowErrorDetailInfo) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RainbowErrorsRainbowErrorDetailInfo) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RainbowErrorsRainbowErrorDetailInfo) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RainbowErrorsRainbowErrorDetailInfo) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RainbowErrorsRainbowErrorDetailInfo) SetMessage(v string) {
	o.Message = &v
}

func (o RainbowErrorsRainbowErrorDetailInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RainbowErrorsRainbowErrorDetailInfo) UnmarshalJSON(bytes []byte) (err error) {
	varRainbowErrorsRainbowErrorDetailInfo := _RainbowErrorsRainbowErrorDetailInfo{}

	if err = json.Unmarshal(bytes, &varRainbowErrorsRainbowErrorDetailInfo); err == nil {
		*o = RainbowErrorsRainbowErrorDetailInfo(varRainbowErrorsRainbowErrorDetailInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRainbowErrorsRainbowErrorDetailInfo struct {
	value *RainbowErrorsRainbowErrorDetailInfo
	isSet bool
}

func (v NullableRainbowErrorsRainbowErrorDetailInfo) Get() *RainbowErrorsRainbowErrorDetailInfo {
	return v.value
}

func (v *NullableRainbowErrorsRainbowErrorDetailInfo) Set(val *RainbowErrorsRainbowErrorDetailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRainbowErrorsRainbowErrorDetailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRainbowErrorsRainbowErrorDetailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRainbowErrorsRainbowErrorDetailInfo(val *RainbowErrorsRainbowErrorDetailInfo) *NullableRainbowErrorsRainbowErrorDetailInfo {
	return &NullableRainbowErrorsRainbowErrorDetailInfo{value: val, isSet: true}
}

func (v NullableRainbowErrorsRainbowErrorDetailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRainbowErrorsRainbowErrorDetailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


