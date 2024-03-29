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

// checks if the ServicesInsertAccountReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesInsertAccountReq{}

// ServicesInsertAccountReq struct for ServicesInsertAccountReq
type ServicesInsertAccountReq struct {
	CreateIfExists *bool `json:"create_if_exists,omitempty"`
	Phone *string `json:"phone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesInsertAccountReq ServicesInsertAccountReq

// NewServicesInsertAccountReq instantiates a new ServicesInsertAccountReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesInsertAccountReq() *ServicesInsertAccountReq {
	this := ServicesInsertAccountReq{}
	return &this
}

// NewServicesInsertAccountReqWithDefaults instantiates a new ServicesInsertAccountReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesInsertAccountReqWithDefaults() *ServicesInsertAccountReq {
	this := ServicesInsertAccountReq{}
	return &this
}

// GetCreateIfExists returns the CreateIfExists field value if set, zero value otherwise.
func (o *ServicesInsertAccountReq) GetCreateIfExists() bool {
	if o == nil || IsNil(o.CreateIfExists) {
		var ret bool
		return ret
	}
	return *o.CreateIfExists
}

// GetCreateIfExistsOk returns a tuple with the CreateIfExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesInsertAccountReq) GetCreateIfExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateIfExists) {
		return nil, false
	}
	return o.CreateIfExists, true
}

// HasCreateIfExists returns a boolean if a field has been set.
func (o *ServicesInsertAccountReq) HasCreateIfExists() bool {
	if o != nil && !IsNil(o.CreateIfExists) {
		return true
	}

	return false
}

// SetCreateIfExists gets a reference to the given bool and assigns it to the CreateIfExists field.
func (o *ServicesInsertAccountReq) SetCreateIfExists(v bool) {
	o.CreateIfExists = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ServicesInsertAccountReq) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesInsertAccountReq) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ServicesInsertAccountReq) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ServicesInsertAccountReq) SetPhone(v string) {
	o.Phone = &v
}

func (o ServicesInsertAccountReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesInsertAccountReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateIfExists) {
		toSerialize["create_if_exists"] = o.CreateIfExists
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesInsertAccountReq) UnmarshalJSON(data []byte) (err error) {
	varServicesInsertAccountReq := _ServicesInsertAccountReq{}

	err = json.Unmarshal(data, &varServicesInsertAccountReq)

	if err != nil {
		return err
	}

	*o = ServicesInsertAccountReq(varServicesInsertAccountReq)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "create_if_exists")
		delete(additionalProperties, "phone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesInsertAccountReq struct {
	value *ServicesInsertAccountReq
	isSet bool
}

func (v NullableServicesInsertAccountReq) Get() *ServicesInsertAccountReq {
	return v.value
}

func (v *NullableServicesInsertAccountReq) Set(val *ServicesInsertAccountReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesInsertAccountReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesInsertAccountReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesInsertAccountReq(val *ServicesInsertAccountReq) *NullableServicesInsertAccountReq {
	return &NullableServicesInsertAccountReq{value: val, isSet: true}
}

func (v NullableServicesInsertAccountReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesInsertAccountReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


