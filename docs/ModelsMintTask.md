# ModelsMintTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
**BlockReason** | Pointer to [**EnumsTransactionBlockReason**](EnumsTransactionBlockReason.md) | 没有发送到tx engine的原因，比如余额不足或网络错误等 | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**ChainType** | Pointer to **int32** |  | [optional] 
**Contract** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**MintTo** | Pointer to **string** |  | [optional] 
**MintType** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failed | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**TokenUri** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsMintTask

`func NewModelsMintTask() *ModelsMintTask`

NewModelsMintTask instantiates a new ModelsMintTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMintTaskWithDefaults

`func NewModelsMintTaskWithDefaults() *ModelsMintTask`

NewModelsMintTaskWithDefaults instantiates a new ModelsMintTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsMintTask) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsMintTask) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsMintTask) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsMintTask) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppId

`func (o *ModelsMintTask) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsMintTask) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsMintTask) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsMintTask) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBlockReason

`func (o *ModelsMintTask) GetBlockReason() EnumsTransactionBlockReason`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *ModelsMintTask) GetBlockReasonOk() (*EnumsTransactionBlockReason, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *ModelsMintTask) SetBlockReason(v EnumsTransactionBlockReason)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *ModelsMintTask) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsMintTask) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsMintTask) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsMintTask) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsMintTask) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsMintTask) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsMintTask) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsMintTask) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsMintTask) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetContract

`func (o *ModelsMintTask) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ModelsMintTask) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ModelsMintTask) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ModelsMintTask) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetContractType

`func (o *ModelsMintTask) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ModelsMintTask) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ModelsMintTask) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ModelsMintTask) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsMintTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsMintTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsMintTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsMintTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsMintTask) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsMintTask) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsMintTask) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsMintTask) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetError

`func (o *ModelsMintTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsMintTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsMintTask) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModelsMintTask) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHash

`func (o *ModelsMintTask) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsMintTask) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsMintTask) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsMintTask) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsMintTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMintTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMintTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMintTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMintTo

`func (o *ModelsMintTask) GetMintTo() string`

GetMintTo returns the MintTo field if non-nil, zero value otherwise.

### GetMintToOk

`func (o *ModelsMintTask) GetMintToOk() (*string, bool)`

GetMintToOk returns a tuple with the MintTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTo

`func (o *ModelsMintTask) SetMintTo(v string)`

SetMintTo sets MintTo field to given value.

### HasMintTo

`func (o *ModelsMintTask) HasMintTo() bool`

HasMintTo returns a boolean if a field has been set.

### GetMintType

`func (o *ModelsMintTask) GetMintType() int32`

GetMintType returns the MintType field if non-nil, zero value otherwise.

### GetMintTypeOk

`func (o *ModelsMintTask) GetMintTypeOk() (*int32, bool)`

GetMintTypeOk returns a tuple with the MintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintType

`func (o *ModelsMintTask) SetMintType(v int32)`

SetMintType sets MintType field to given value.

### HasMintType

`func (o *ModelsMintTask) HasMintType() bool`

HasMintType returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsMintTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsMintTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsMintTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsMintTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenId

`func (o *ModelsMintTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsMintTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsMintTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ModelsMintTask) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenUri

`func (o *ModelsMintTask) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *ModelsMintTask) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *ModelsMintTask) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *ModelsMintTask) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetTxId

`func (o *ModelsMintTask) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ModelsMintTask) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ModelsMintTask) SetTxId(v int32)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ModelsMintTask) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsMintTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsMintTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsMintTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsMintTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


