# MultipartFileHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **map[string][]string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewMultipartFileHeader

`func NewMultipartFileHeader() *MultipartFileHeader`

NewMultipartFileHeader instantiates a new MultipartFileHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipartFileHeaderWithDefaults

`func NewMultipartFileHeaderWithDefaults() *MultipartFileHeader`

NewMultipartFileHeaderWithDefaults instantiates a new MultipartFileHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *MultipartFileHeader) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MultipartFileHeader) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MultipartFileHeader) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MultipartFileHeader) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetHeader

`func (o *MultipartFileHeader) GetHeader() map[string][]string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *MultipartFileHeader) GetHeaderOk() (*map[string][]string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *MultipartFileHeader) SetHeader(v map[string][]string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *MultipartFileHeader) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSize

`func (o *MultipartFileHeader) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MultipartFileHeader) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MultipartFileHeader) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MultipartFileHeader) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


