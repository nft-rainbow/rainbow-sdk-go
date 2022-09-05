# ModelsMetadataAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**DisplayType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**MetadataId** | Pointer to **int32** |  | [optional] 
**TraitType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsMetadataAttribute

`func NewModelsMetadataAttribute() *ModelsMetadataAttribute`

NewModelsMetadataAttribute instantiates a new ModelsMetadataAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMetadataAttributeWithDefaults

`func NewModelsMetadataAttributeWithDefaults() *ModelsMetadataAttribute`

NewModelsMetadataAttributeWithDefaults instantiates a new ModelsMetadataAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *ModelsMetadataAttribute) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ModelsMetadataAttribute) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ModelsMetadataAttribute) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *ModelsMetadataAttribute) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsMetadataAttribute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsMetadataAttribute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsMetadataAttribute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsMetadataAttribute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsMetadataAttribute) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsMetadataAttribute) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsMetadataAttribute) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsMetadataAttribute) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDisplayType

`func (o *ModelsMetadataAttribute) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *ModelsMetadataAttribute) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *ModelsMetadataAttribute) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *ModelsMetadataAttribute) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetId

`func (o *ModelsMetadataAttribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMetadataAttribute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMetadataAttribute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMetadataAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataId

`func (o *ModelsMetadataAttribute) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *ModelsMetadataAttribute) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *ModelsMetadataAttribute) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *ModelsMetadataAttribute) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetTraitType

`func (o *ModelsMetadataAttribute) GetTraitType() string`

GetTraitType returns the TraitType field if non-nil, zero value otherwise.

### GetTraitTypeOk

`func (o *ModelsMetadataAttribute) GetTraitTypeOk() (*string, bool)`

GetTraitTypeOk returns a tuple with the TraitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitType

`func (o *ModelsMetadataAttribute) SetTraitType(v string)`

SetTraitType sets TraitType field to given value.

### HasTraitType

`func (o *ModelsMetadataAttribute) HasTraitType() bool`

HasTraitType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsMetadataAttribute) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsMetadataAttribute) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsMetadataAttribute) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsMetadataAttribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *ModelsMetadataAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelsMetadataAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelsMetadataAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelsMetadataAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


