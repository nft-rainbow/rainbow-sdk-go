# ModelsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to [**[]ModelsMetadataAttribute**](ModelsMetadataAttribute.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Description** | **string** |  | 
**ExternalLink** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Image** | **string** |  | 
**MetadataId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsMetadata

`func NewModelsMetadata(description string, image string, name string, ) *ModelsMetadata`

NewModelsMetadata instantiates a new ModelsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMetadataWithDefaults

`func NewModelsMetadataWithDefaults() *ModelsMetadata`

NewModelsMetadataWithDefaults instantiates a new ModelsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ModelsMetadata) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsMetadata) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsMetadata) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsMetadata) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAttributes

`func (o *ModelsMetadata) GetAttributes() []ModelsMetadataAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ModelsMetadata) GetAttributesOk() (*[]ModelsMetadataAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ModelsMetadata) SetAttributes(v []ModelsMetadataAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ModelsMetadata) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsMetadata) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsMetadata) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsMetadata) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsMetadata) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsMetadata) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsMetadata) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalLink

`func (o *ModelsMetadata) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *ModelsMetadata) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *ModelsMetadata) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *ModelsMetadata) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetId

`func (o *ModelsMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ModelsMetadata) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModelsMetadata) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModelsMetadata) SetImage(v string)`

SetImage sets Image field to given value.


### GetMetadataId

`func (o *ModelsMetadata) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *ModelsMetadata) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *ModelsMetadata) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *ModelsMetadata) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetName

`func (o *ModelsMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *ModelsMetadata) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsMetadata) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsMetadata) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUri

`func (o *ModelsMetadata) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ModelsMetadata) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ModelsMetadata) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ModelsMetadata) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


