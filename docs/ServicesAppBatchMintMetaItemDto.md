# ServicesAppBatchMintMetaItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | amount on same token id, only erc1155 contract could set large than 1, others set null or 1 | [optional] 
**AnimationUrl** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]ModelsExposedMetadataAttribute**](ModelsExposedMetadataAttribute.md) |  | [optional] 
**Chain** | **string** |  | 
**ContractAddress** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FileUrl** | **string** |  | 
**MintToAddress** | **string** |  | 
**Name** | **string** |  | 
**Number** | Pointer to **int32** | mint number, everyone with different token id (not conflict with AmountOnSameTokenID) | [optional] 

## Methods

### NewServicesAppBatchMintMetaItemDto

`func NewServicesAppBatchMintMetaItemDto(chain string, fileUrl string, mintToAddress string, name string, ) *ServicesAppBatchMintMetaItemDto`

NewServicesAppBatchMintMetaItemDto instantiates a new ServicesAppBatchMintMetaItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesAppBatchMintMetaItemDtoWithDefaults

`func NewServicesAppBatchMintMetaItemDtoWithDefaults() *ServicesAppBatchMintMetaItemDto`

NewServicesAppBatchMintMetaItemDtoWithDefaults instantiates a new ServicesAppBatchMintMetaItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesAppBatchMintMetaItemDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesAppBatchMintMetaItemDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesAppBatchMintMetaItemDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesAppBatchMintMetaItemDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAnimationUrl

`func (o *ServicesAppBatchMintMetaItemDto) GetAnimationUrl() string`

GetAnimationUrl returns the AnimationUrl field if non-nil, zero value otherwise.

### GetAnimationUrlOk

`func (o *ServicesAppBatchMintMetaItemDto) GetAnimationUrlOk() (*string, bool)`

GetAnimationUrlOk returns a tuple with the AnimationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationUrl

`func (o *ServicesAppBatchMintMetaItemDto) SetAnimationUrl(v string)`

SetAnimationUrl sets AnimationUrl field to given value.

### HasAnimationUrl

`func (o *ServicesAppBatchMintMetaItemDto) HasAnimationUrl() bool`

HasAnimationUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *ServicesAppBatchMintMetaItemDto) GetAttributes() []ModelsExposedMetadataAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServicesAppBatchMintMetaItemDto) GetAttributesOk() (*[]ModelsExposedMetadataAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServicesAppBatchMintMetaItemDto) SetAttributes(v []ModelsExposedMetadataAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServicesAppBatchMintMetaItemDto) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetChain

`func (o *ServicesAppBatchMintMetaItemDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesAppBatchMintMetaItemDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesAppBatchMintMetaItemDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesAppBatchMintMetaItemDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesAppBatchMintMetaItemDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesAppBatchMintMetaItemDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *ServicesAppBatchMintMetaItemDto) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetDescription

`func (o *ServicesAppBatchMintMetaItemDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesAppBatchMintMetaItemDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesAppBatchMintMetaItemDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicesAppBatchMintMetaItemDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileUrl

`func (o *ServicesAppBatchMintMetaItemDto) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ServicesAppBatchMintMetaItemDto) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ServicesAppBatchMintMetaItemDto) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetMintToAddress

`func (o *ServicesAppBatchMintMetaItemDto) GetMintToAddress() string`

GetMintToAddress returns the MintToAddress field if non-nil, zero value otherwise.

### GetMintToAddressOk

`func (o *ServicesAppBatchMintMetaItemDto) GetMintToAddressOk() (*string, bool)`

GetMintToAddressOk returns a tuple with the MintToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintToAddress

`func (o *ServicesAppBatchMintMetaItemDto) SetMintToAddress(v string)`

SetMintToAddress sets MintToAddress field to given value.


### GetName

`func (o *ServicesAppBatchMintMetaItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesAppBatchMintMetaItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesAppBatchMintMetaItemDto) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *ServicesAppBatchMintMetaItemDto) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServicesAppBatchMintMetaItemDto) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServicesAppBatchMintMetaItemDto) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServicesAppBatchMintMetaItemDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


