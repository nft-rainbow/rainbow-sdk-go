# ModelsContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
**BaseUri** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**ChainType** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerAddress** | Pointer to **string** |  | [optional] 
**RoyaltiesAddress** | Pointer to **string** |  | [optional] 
**RoyaltiesBps** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failed | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TokensBurnable** | Pointer to **bool** |  | [optional] 
**TokensTransferable** | Pointer to **bool** |  | [optional] 
**TransferCooldownTime** | Pointer to **int32** |  | [optional] 
**TxId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** | 1-ERC721, 2-ERC1155 | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsContract

`func NewModelsContract() *ModelsContract`

NewModelsContract instantiates a new ModelsContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsContractWithDefaults

`func NewModelsContractWithDefaults() *ModelsContract`

NewModelsContractWithDefaults instantiates a new ModelsContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ModelsContract) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelsContract) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelsContract) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelsContract) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAppId

`func (o *ModelsContract) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsContract) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsContract) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsContract) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBaseUri

`func (o *ModelsContract) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *ModelsContract) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *ModelsContract) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *ModelsContract) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsContract) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsContract) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsContract) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsContract) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsContract) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsContract) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsContract) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsContract) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsContract) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsContract) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsContract) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsContract) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsContract) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsContract) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsContract) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsContract) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHash

`func (o *ModelsContract) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsContract) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsContract) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsContract) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsContract) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsContract) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsContract) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsContract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsContract) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsContract) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *ModelsContract) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *ModelsContract) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *ModelsContract) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *ModelsContract) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetRoyaltiesAddress

`func (o *ModelsContract) GetRoyaltiesAddress() string`

GetRoyaltiesAddress returns the RoyaltiesAddress field if non-nil, zero value otherwise.

### GetRoyaltiesAddressOk

`func (o *ModelsContract) GetRoyaltiesAddressOk() (*string, bool)`

GetRoyaltiesAddressOk returns a tuple with the RoyaltiesAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltiesAddress

`func (o *ModelsContract) SetRoyaltiesAddress(v string)`

SetRoyaltiesAddress sets RoyaltiesAddress field to given value.

### HasRoyaltiesAddress

`func (o *ModelsContract) HasRoyaltiesAddress() bool`

HasRoyaltiesAddress returns a boolean if a field has been set.

### GetRoyaltiesBps

`func (o *ModelsContract) GetRoyaltiesBps() int32`

GetRoyaltiesBps returns the RoyaltiesBps field if non-nil, zero value otherwise.

### GetRoyaltiesBpsOk

`func (o *ModelsContract) GetRoyaltiesBpsOk() (*int32, bool)`

GetRoyaltiesBpsOk returns a tuple with the RoyaltiesBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltiesBps

`func (o *ModelsContract) SetRoyaltiesBps(v int32)`

SetRoyaltiesBps sets RoyaltiesBps field to given value.

### HasRoyaltiesBps

`func (o *ModelsContract) HasRoyaltiesBps() bool`

HasRoyaltiesBps returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsContract) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsContract) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsContract) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsContract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *ModelsContract) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ModelsContract) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ModelsContract) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ModelsContract) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTokensBurnable

`func (o *ModelsContract) GetTokensBurnable() bool`

GetTokensBurnable returns the TokensBurnable field if non-nil, zero value otherwise.

### GetTokensBurnableOk

`func (o *ModelsContract) GetTokensBurnableOk() (*bool, bool)`

GetTokensBurnableOk returns a tuple with the TokensBurnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensBurnable

`func (o *ModelsContract) SetTokensBurnable(v bool)`

SetTokensBurnable sets TokensBurnable field to given value.

### HasTokensBurnable

`func (o *ModelsContract) HasTokensBurnable() bool`

HasTokensBurnable returns a boolean if a field has been set.

### GetTokensTransferable

`func (o *ModelsContract) GetTokensTransferable() bool`

GetTokensTransferable returns the TokensTransferable field if non-nil, zero value otherwise.

### GetTokensTransferableOk

`func (o *ModelsContract) GetTokensTransferableOk() (*bool, bool)`

GetTokensTransferableOk returns a tuple with the TokensTransferable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensTransferable

`func (o *ModelsContract) SetTokensTransferable(v bool)`

SetTokensTransferable sets TokensTransferable field to given value.

### HasTokensTransferable

`func (o *ModelsContract) HasTokensTransferable() bool`

HasTokensTransferable returns a boolean if a field has been set.

### GetTransferCooldownTime

`func (o *ModelsContract) GetTransferCooldownTime() int32`

GetTransferCooldownTime returns the TransferCooldownTime field if non-nil, zero value otherwise.

### GetTransferCooldownTimeOk

`func (o *ModelsContract) GetTransferCooldownTimeOk() (*int32, bool)`

GetTransferCooldownTimeOk returns a tuple with the TransferCooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferCooldownTime

`func (o *ModelsContract) SetTransferCooldownTime(v int32)`

SetTransferCooldownTime sets TransferCooldownTime field to given value.

### HasTransferCooldownTime

`func (o *ModelsContract) HasTransferCooldownTime() bool`

HasTransferCooldownTime returns a boolean if a field has been set.

### GetTxId

`func (o *ModelsContract) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ModelsContract) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ModelsContract) SetTxId(v int32)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ModelsContract) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetType

`func (o *ModelsContract) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsContract) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsContract) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsContract) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsContract) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsContract) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsContract) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsContract) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


