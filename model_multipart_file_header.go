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

// checks if the MultipartFileHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipartFileHeader{}

// MultipartFileHeader struct for MultipartFileHeader
type MultipartFileHeader struct {
	Filename *string `json:"filename,omitempty"`
	Header *map[string][]string `json:"header,omitempty"`
	Size *int32 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultipartFileHeader MultipartFileHeader

// NewMultipartFileHeader instantiates a new MultipartFileHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipartFileHeader() *MultipartFileHeader {
	this := MultipartFileHeader{}
	return &this
}

// NewMultipartFileHeaderWithDefaults instantiates a new MultipartFileHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipartFileHeaderWithDefaults() *MultipartFileHeader {
	this := MultipartFileHeader{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *MultipartFileHeader) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartFileHeader) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *MultipartFileHeader) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *MultipartFileHeader) SetFilename(v string) {
	o.Filename = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *MultipartFileHeader) GetHeader() map[string][]string {
	if o == nil || IsNil(o.Header) {
		var ret map[string][]string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartFileHeader) GetHeaderOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *MultipartFileHeader) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given map[string][]string and assigns it to the Header field.
func (o *MultipartFileHeader) SetHeader(v map[string][]string) {
	o.Header = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MultipartFileHeader) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartFileHeader) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MultipartFileHeader) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MultipartFileHeader) SetSize(v int32) {
	o.Size = &v
}

func (o MultipartFileHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipartFileHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MultipartFileHeader) UnmarshalJSON(data []byte) (err error) {
	varMultipartFileHeader := _MultipartFileHeader{}

	err = json.Unmarshal(data, &varMultipartFileHeader)

	if err != nil {
		return err
	}

	*o = MultipartFileHeader(varMultipartFileHeader)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filename")
		delete(additionalProperties, "header")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultipartFileHeader struct {
	value *MultipartFileHeader
	isSet bool
}

func (v NullableMultipartFileHeader) Get() *MultipartFileHeader {
	return v.value
}

func (v *NullableMultipartFileHeader) Set(val *MultipartFileHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipartFileHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipartFileHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipartFileHeader(val *MultipartFileHeader) *NullableMultipartFileHeader {
	return &NullableMultipartFileHeader{value: val, isSet: true}
}

func (v NullableMultipartFileHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipartFileHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


