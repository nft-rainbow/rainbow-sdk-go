# ModelsTokenBoundAccountCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** | The chain id of the tba collection. All tbas in the collection are of same chain id. | [optional] 
**ChainType** | Pointer to **int32** | the type of the chain. 1-CFX, 2-ETH. All tbas in the collection are of same chain type | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Description** | Pointer to **string** | Collection description | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]ModelsTokenBoundAccount**](ModelsTokenBoundAccount.md) |  | [optional] 
**Name** | Pointer to **string** | The name of tba collection. The name is unique for the user. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** | The user created the collection | [optional] 

## Methods

### NewModelsTokenBoundAccountCollection

`func NewModelsTokenBoundAccountCollection() *ModelsTokenBoundAccountCollection`

NewModelsTokenBoundAccountCollection instantiates a new ModelsTokenBoundAccountCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTokenBoundAccountCollectionWithDefaults

`func NewModelsTokenBoundAccountCollectionWithDefaults() *ModelsTokenBoundAccountCollection`

NewModelsTokenBoundAccountCollectionWithDefaults instantiates a new ModelsTokenBoundAccountCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ModelsTokenBoundAccountCollection) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsTokenBoundAccountCollection) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsTokenBoundAccountCollection) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsTokenBoundAccountCollection) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsTokenBoundAccountCollection) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsTokenBoundAccountCollection) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsTokenBoundAccountCollection) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsTokenBoundAccountCollection) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsTokenBoundAccountCollection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsTokenBoundAccountCollection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsTokenBoundAccountCollection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsTokenBoundAccountCollection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsTokenBoundAccountCollection) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsTokenBoundAccountCollection) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsTokenBoundAccountCollection) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsTokenBoundAccountCollection) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsTokenBoundAccountCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsTokenBoundAccountCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsTokenBoundAccountCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsTokenBoundAccountCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModelsTokenBoundAccountCollection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTokenBoundAccountCollection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTokenBoundAccountCollection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTokenBoundAccountCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ModelsTokenBoundAccountCollection) GetItems() []ModelsTokenBoundAccount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelsTokenBoundAccountCollection) GetItemsOk() (*[]ModelsTokenBoundAccount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelsTokenBoundAccountCollection) SetItems(v []ModelsTokenBoundAccount)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelsTokenBoundAccountCollection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *ModelsTokenBoundAccountCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTokenBoundAccountCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTokenBoundAccountCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsTokenBoundAccountCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsTokenBoundAccountCollection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTokenBoundAccountCollection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTokenBoundAccountCollection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsTokenBoundAccountCollection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ModelsTokenBoundAccountCollection) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelsTokenBoundAccountCollection) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelsTokenBoundAccountCollection) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelsTokenBoundAccountCollection) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


