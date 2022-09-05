# ModelsFilesQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]ModelsExposedFile**](ModelsExposedFile.md) |  | [optional] 

## Methods

### NewModelsFilesQueryResult

`func NewModelsFilesQueryResult() *ModelsFilesQueryResult`

NewModelsFilesQueryResult instantiates a new ModelsFilesQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsFilesQueryResultWithDefaults

`func NewModelsFilesQueryResultWithDefaults() *ModelsFilesQueryResult`

NewModelsFilesQueryResultWithDefaults instantiates a new ModelsFilesQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ModelsFilesQueryResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelsFilesQueryResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelsFilesQueryResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ModelsFilesQueryResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ModelsFilesQueryResult) GetItems() []ModelsExposedFile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelsFilesQueryResult) GetItemsOk() (*[]ModelsExposedFile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelsFilesQueryResult) SetItems(v []ModelsExposedFile)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelsFilesQueryResult) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


