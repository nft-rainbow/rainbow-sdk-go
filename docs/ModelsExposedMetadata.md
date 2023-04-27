# ModelsExposedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnimationUrl** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to [**[]ModelsExposedMetadataAttribute**](ModelsExposedMetadataAttribute.md) |  | [optional] 
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

### NewModelsExposedMetadata

`func NewModelsExposedMetadata(description string, image string, name string, ) *ModelsExposedMetadata`

NewModelsExposedMetadata instantiates a new ModelsExposedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsExposedMetadataWithDefaults

`func NewModelsExposedMetadataWithDefaults() *ModelsExposedMetadata`

NewModelsExposedMetadataWithDefaults instantiates a new ModelsExposedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnimationUrl

`func (o *ModelsExposedMetadata) GetAnimationUrl() string`

GetAnimationUrl returns the AnimationUrl field if non-nil, zero value otherwise.

### GetAnimationUrlOk

`func (o *ModelsExposedMetadata) GetAnimationUrlOk() (*string, bool)`

GetAnimationUrlOk returns a tuple with the AnimationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationUrl

`func (o *ModelsExposedMetadata) SetAnimationUrl(v string)`

SetAnimationUrl sets AnimationUrl field to given value.

### HasAnimationUrl

`func (o *ModelsExposedMetadata) HasAnimationUrl() bool`

HasAnimationUrl returns a boolean if a field has been set.

### GetAppId

`func (o *ModelsExposedMetadata) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsExposedMetadata) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsExposedMetadata) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsExposedMetadata) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAttributes

`func (o *ModelsExposedMetadata) GetAttributes() []ModelsExposedMetadataAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ModelsExposedMetadata) GetAttributesOk() (*[]ModelsExposedMetadataAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ModelsExposedMetadata) SetAttributes(v []ModelsExposedMetadataAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ModelsExposedMetadata) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsExposedMetadata) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsExposedMetadata) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsExposedMetadata) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsExposedMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsExposedMetadata) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsExposedMetadata) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsExposedMetadata) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsExposedMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsExposedMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsExposedMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsExposedMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalLink

`func (o *ModelsExposedMetadata) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *ModelsExposedMetadata) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *ModelsExposedMetadata) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *ModelsExposedMetadata) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetId

`func (o *ModelsExposedMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsExposedMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsExposedMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsExposedMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ModelsExposedMetadata) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModelsExposedMetadata) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModelsExposedMetadata) SetImage(v string)`

SetImage sets Image field to given value.


### GetMetadataId

`func (o *ModelsExposedMetadata) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *ModelsExposedMetadata) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *ModelsExposedMetadata) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *ModelsExposedMetadata) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetName

`func (o *ModelsExposedMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsExposedMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsExposedMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *ModelsExposedMetadata) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsExposedMetadata) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsExposedMetadata) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsExposedMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUri

`func (o *ModelsExposedMetadata) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ModelsExposedMetadata) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ModelsExposedMetadata) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ModelsExposedMetadata) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


