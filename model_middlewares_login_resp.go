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

// checks if the MiddlewaresLoginResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewaresLoginResp{}

// MiddlewaresLoginResp struct for MiddlewaresLoginResp
type MiddlewaresLoginResp struct {
	Expire *string `json:"expire,omitempty"`
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiddlewaresLoginResp MiddlewaresLoginResp

// NewMiddlewaresLoginResp instantiates a new MiddlewaresLoginResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewaresLoginResp() *MiddlewaresLoginResp {
	this := MiddlewaresLoginResp{}
	return &this
}

// NewMiddlewaresLoginRespWithDefaults instantiates a new MiddlewaresLoginResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewaresLoginRespWithDefaults() *MiddlewaresLoginResp {
	this := MiddlewaresLoginResp{}
	return &this
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *MiddlewaresLoginResp) GetExpire() string {
	if o == nil || isNil(o.Expire) {
		var ret string
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewaresLoginResp) GetExpireOk() (*string, bool) {
	if o == nil || isNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *MiddlewaresLoginResp) HasExpire() bool {
	if o != nil && !isNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given string and assigns it to the Expire field.
func (o *MiddlewaresLoginResp) SetExpire(v string) {
	o.Expire = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MiddlewaresLoginResp) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewaresLoginResp) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MiddlewaresLoginResp) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MiddlewaresLoginResp) SetToken(v string) {
	o.Token = &v
}

func (o MiddlewaresLoginResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewaresLoginResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MiddlewaresLoginResp) UnmarshalJSON(bytes []byte) (err error) {
	varMiddlewaresLoginResp := _MiddlewaresLoginResp{}

	if err = json.Unmarshal(bytes, &varMiddlewaresLoginResp); err == nil {
		*o = MiddlewaresLoginResp(varMiddlewaresLoginResp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expire")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiddlewaresLoginResp struct {
	value *MiddlewaresLoginResp
	isSet bool
}

func (v NullableMiddlewaresLoginResp) Get() *MiddlewaresLoginResp {
	return v.value
}

func (v *NullableMiddlewaresLoginResp) Set(val *MiddlewaresLoginResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewaresLoginResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewaresLoginResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewaresLoginResp(val *MiddlewaresLoginResp) *NullableMiddlewaresLoginResp {
	return &NullableMiddlewaresLoginResp{value: val, isSet: true}
}

func (v NullableMiddlewaresLoginResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewaresLoginResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


