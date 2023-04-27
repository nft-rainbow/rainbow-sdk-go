# ServicesAppMintByMetaPartsDto

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

### NewServicesAppMintByMetaPartsDto

`func NewServicesAppMintByMetaPartsDto(chain string, fileUrl string, mintToAddress string, name string, ) *ServicesAppMintByMetaPartsDto`

NewServicesAppMintByMetaPartsDto instantiates a new ServicesAppMintByMetaPartsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesAppMintByMetaPartsDtoWithDefaults

`func NewServicesAppMintByMetaPartsDtoWithDefaults() *ServicesAppMintByMetaPartsDto`

NewServicesAppMintByMetaPartsDtoWithDefaults instantiates a new ServicesAppMintByMetaPartsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesAppMintByMetaPartsDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesAppMintByMetaPartsDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesAppMintByMetaPartsDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesAppMintByMetaPartsDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAnimationUrl

`func (o *ServicesAppMintByMetaPartsDto) GetAnimationUrl() string`

GetAnimationUrl returns the AnimationUrl field if non-nil, zero value otherwise.

### GetAnimationUrlOk

`func (o *ServicesAppMintByMetaPartsDto) GetAnimationUrlOk() (*string, bool)`

GetAnimationUrlOk returns a tuple with the AnimationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationUrl

`func (o *ServicesAppMintByMetaPartsDto) SetAnimationUrl(v string)`

SetAnimationUrl sets AnimationUrl field to given value.

### HasAnimationUrl

`func (o *ServicesAppMintByMetaPartsDto) HasAnimationUrl() bool`

HasAnimationUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *ServicesAppMintByMetaPartsDto) GetAttributes() []ModelsExposedMetadataAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServicesAppMintByMetaPartsDto) GetAttributesOk() (*[]ModelsExposedMetadataAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServicesAppMintByMetaPartsDto) SetAttributes(v []ModelsExposedMetadataAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServicesAppMintByMetaPartsDto) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetChain

`func (o *ServicesAppMintByMetaPartsDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesAppMintByMetaPartsDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesAppMintByMetaPartsDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesAppMintByMetaPartsDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesAppMintByMetaPartsDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesAppMintByMetaPartsDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *ServicesAppMintByMetaPartsDto) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetDescription

`func (o *ServicesAppMintByMetaPartsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesAppMintByMetaPartsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesAppMintByMetaPartsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicesAppMintByMetaPartsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileUrl

`func (o *ServicesAppMintByMetaPartsDto) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ServicesAppMintByMetaPartsDto) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ServicesAppMintByMetaPartsDto) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetMintToAddress

`func (o *ServicesAppMintByMetaPartsDto) GetMintToAddress() string`

GetMintToAddress returns the MintToAddress field if non-nil, zero value otherwise.

### GetMintToAddressOk

`func (o *ServicesAppMintByMetaPartsDto) GetMintToAddressOk() (*string, bool)`

GetMintToAddressOk returns a tuple with the MintToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintToAddress

`func (o *ServicesAppMintByMetaPartsDto) SetMintToAddress(v string)`

SetMintToAddress sets MintToAddress field to given value.


### GetName

`func (o *ServicesAppMintByMetaPartsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesAppMintByMetaPartsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesAppMintByMetaPartsDto) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *ServicesAppMintByMetaPartsDto) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServicesAppMintByMetaPartsDto) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServicesAppMintByMetaPartsDto) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServicesAppMintByMetaPartsDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


