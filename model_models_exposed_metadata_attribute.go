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

// checks if the ModelsExposedMetadataAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsExposedMetadataAttribute{}

// ModelsExposedMetadataAttribute struct for ModelsExposedMetadataAttribute
type ModelsExposedMetadataAttribute struct {
	AttributeName *string `json:"attribute_name,omitempty"`
	DisplayType *string `json:"display_type,omitempty"`
	TraitType *string `json:"trait_type,omitempty"`
	// TODO support number
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsExposedMetadataAttribute ModelsExposedMetadataAttribute

// NewModelsExposedMetadataAttribute instantiates a new ModelsExposedMetadataAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsExposedMetadataAttribute() *ModelsExposedMetadataAttribute {
	this := ModelsExposedMetadataAttribute{}
	return &this
}

// NewModelsExposedMetadataAttributeWithDefaults instantiates a new ModelsExposedMetadataAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsExposedMetadataAttributeWithDefaults() *ModelsExposedMetadataAttribute {
	this := ModelsExposedMetadataAttribute{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *ModelsExposedMetadataAttribute) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataAttribute) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *ModelsExposedMetadataAttribute) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *ModelsExposedMetadataAttribute) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *ModelsExposedMetadataAttribute) GetDisplayType() string {
	if o == nil || IsNil(o.DisplayType) {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataAttribute) GetDisplayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayType) {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *ModelsExposedMetadataAttribute) HasDisplayType() bool {
	if o != nil && !IsNil(o.DisplayType) {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *ModelsExposedMetadataAttribute) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetTraitType returns the TraitType field value if set, zero value otherwise.
func (o *ModelsExposedMetadataAttribute) GetTraitType() string {
	if o == nil || IsNil(o.TraitType) {
		var ret string
		return ret
	}
	return *o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataAttribute) GetTraitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TraitType) {
		return nil, false
	}
	return o.TraitType, true
}

// HasTraitType returns a boolean if a field has been set.
func (o *ModelsExposedMetadataAttribute) HasTraitType() bool {
	if o != nil && !IsNil(o.TraitType) {
		return true
	}

	return false
}

// SetTraitType gets a reference to the given string and assigns it to the TraitType field.
func (o *ModelsExposedMetadataAttribute) SetTraitType(v string) {
	o.TraitType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModelsExposedMetadataAttribute) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsExposedMetadataAttribute) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModelsExposedMetadataAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModelsExposedMetadataAttribute) SetValue(v string) {
	o.Value = &v
}

func (o ModelsExposedMetadataAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsExposedMetadataAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attribute_name"] = o.AttributeName
	}
	if !IsNil(o.DisplayType) {
		toSerialize["display_type"] = o.DisplayType
	}
	if !IsNil(o.TraitType) {
		toSerialize["trait_type"] = o.TraitType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsExposedMetadataAttribute) UnmarshalJSON(data []byte) (err error) {
	varModelsExposedMetadataAttribute := _ModelsExposedMetadataAttribute{}

	err = json.Unmarshal(data, &varModelsExposedMetadataAttribute)

	if err != nil {
		return err
	}

	*o = ModelsExposedMetadataAttribute(varModelsExposedMetadataAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute_name")
		delete(additionalProperties, "display_type")
		delete(additionalProperties, "trait_type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsExposedMetadataAttribute struct {
	value *ModelsExposedMetadataAttribute
	isSet bool
}

func (v NullableModelsExposedMetadataAttribute) Get() *ModelsExposedMetadataAttribute {
	return v.value
}

func (v *NullableModelsExposedMetadataAttribute) Set(val *ModelsExposedMetadataAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsExposedMetadataAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsExposedMetadataAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsExposedMetadataAttribute(val *ModelsExposedMetadataAttribute) *NullableModelsExposedMetadataAttribute {
	return &NullableModelsExposedMetadataAttribute{value: val, isSet: true}
}

func (v NullableModelsExposedMetadataAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsExposedMetadataAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


