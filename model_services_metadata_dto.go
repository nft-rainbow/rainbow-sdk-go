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

// ServicesMetadataDto struct for ServicesMetadataDto
type ServicesMetadataDto struct {
	Attributes []ModelsExposedMetadataAttribute `json:"attributes,omitempty"`
	Description string `json:"description"`
	ExternalLink *string `json:"external_link,omitempty"`
	Image string `json:"image"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMetadataDto ServicesMetadataDto

// NewServicesMetadataDto instantiates a new ServicesMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMetadataDto(description string, image string, name string) *ServicesMetadataDto {
	this := ServicesMetadataDto{}
	this.Description = description
	this.Image = image
	this.Name = name
	return &this
}

// NewServicesMetadataDtoWithDefaults instantiates a new ServicesMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMetadataDtoWithDefaults() *ServicesMetadataDto {
	this := ServicesMetadataDto{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServicesMetadataDto) GetAttributes() []ModelsExposedMetadataAttribute {
	if o == nil || o.Attributes == nil {
		var ret []ModelsExposedMetadataAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMetadataDto) GetAttributesOk() ([]ModelsExposedMetadataAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServicesMetadataDto) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []ModelsExposedMetadataAttribute and assigns it to the Attributes field.
func (o *ServicesMetadataDto) SetAttributes(v []ModelsExposedMetadataAttribute) {
	o.Attributes = v
}

// GetDescription returns the Description field value
func (o *ServicesMetadataDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServicesMetadataDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServicesMetadataDto) SetDescription(v string) {
	o.Description = v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *ServicesMetadataDto) GetExternalLink() string {
	if o == nil || o.ExternalLink == nil {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMetadataDto) GetExternalLinkOk() (*string, bool) {
	if o == nil || o.ExternalLink == nil {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *ServicesMetadataDto) HasExternalLink() bool {
	if o != nil && o.ExternalLink != nil {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *ServicesMetadataDto) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetImage returns the Image field value
func (o *ServicesMetadataDto) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ServicesMetadataDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ServicesMetadataDto) SetImage(v string) {
	o.Image = v
}

// GetName returns the Name field value
func (o *ServicesMetadataDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServicesMetadataDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServicesMetadataDto) SetName(v string) {
	o.Name = v
}

func (o ServicesMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.ExternalLink != nil {
		toSerialize["external_link"] = o.ExternalLink
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesMetadataDto) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMetadataDto := _ServicesMetadataDto{}

	if err = json.Unmarshal(bytes, &varServicesMetadataDto); err == nil {
		*o = ServicesMetadataDto(varServicesMetadataDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "description")
		delete(additionalProperties, "external_link")
		delete(additionalProperties, "image")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMetadataDto struct {
	value *ServicesMetadataDto
	isSet bool
}

func (v NullableServicesMetadataDto) Get() *ServicesMetadataDto {
	return v.value
}

func (v *NullableServicesMetadataDto) Set(val *ServicesMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMetadataDto(val *ServicesMetadataDto) *NullableServicesMetadataDto {
	return &NullableServicesMetadataDto{value: val, isSet: true}
}

func (v NullableServicesMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


