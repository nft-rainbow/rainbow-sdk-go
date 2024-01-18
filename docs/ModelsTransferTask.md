# ModelsTransferTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
**BlockReason** | Pointer to [**EnumsTransactionBlockReason**](EnumsTransactionBlockReason.md) |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**ChainType** | Pointer to **int32** |  | [optional] 
**Contract** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failed | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**TransferFrom** | Pointer to **string** |  | [optional] 
**TransferTo** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsTransferTask

`func NewModelsTransferTask() *ModelsTransferTask`

NewModelsTransferTask instantiates a new ModelsTransferTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTransferTaskWithDefaults

`func NewModelsTransferTaskWithDefaults() *ModelsTransferTask`

NewModelsTransferTaskWithDefaults instantiates a new ModelsTransferTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsTransferTask) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsTransferTask) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsTransferTask) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsTransferTask) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppId

`func (o *ModelsTransferTask) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsTransferTask) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsTransferTask) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsTransferTask) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBlockReason

`func (o *ModelsTransferTask) GetBlockReason() EnumsTransactionBlockReason`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *ModelsTransferTask) GetBlockReasonOk() (*EnumsTransactionBlockReason, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *ModelsTransferTask) SetBlockReason(v EnumsTransactionBlockReason)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *ModelsTransferTask) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsTransferTask) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsTransferTask) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsTransferTask) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsTransferTask) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsTransferTask) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsTransferTask) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsTransferTask) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsTransferTask) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetContract

`func (o *ModelsTransferTask) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ModelsTransferTask) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ModelsTransferTask) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ModelsTransferTask) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetContractType

`func (o *ModelsTransferTask) GetContractType() int32`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ModelsTransferTask) GetContractTypeOk() (*int32, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ModelsTransferTask) SetContractType(v int32)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ModelsTransferTask) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsTransferTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsTransferTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsTransferTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsTransferTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsTransferTask) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsTransferTask) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsTransferTask) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsTransferTask) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetError

`func (o *ModelsTransferTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsTransferTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsTransferTask) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModelsTransferTask) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHash

`func (o *ModelsTransferTask) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsTransferTask) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsTransferTask) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsTransferTask) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsTransferTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTransferTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTransferTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTransferTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsTransferTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsTransferTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsTransferTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsTransferTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenId

`func (o *ModelsTransferTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsTransferTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsTransferTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ModelsTransferTask) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTransferFrom

`func (o *ModelsTransferTask) GetTransferFrom() string`

GetTransferFrom returns the TransferFrom field if non-nil, zero value otherwise.

### GetTransferFromOk

`func (o *ModelsTransferTask) GetTransferFromOk() (*string, bool)`

GetTransferFromOk returns a tuple with the TransferFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFrom

`func (o *ModelsTransferTask) SetTransferFrom(v string)`

SetTransferFrom sets TransferFrom field to given value.

### HasTransferFrom

`func (o *ModelsTransferTask) HasTransferFrom() bool`

HasTransferFrom returns a boolean if a field has been set.

### GetTransferTo

`func (o *ModelsTransferTask) GetTransferTo() string`

GetTransferTo returns the TransferTo field if non-nil, zero value otherwise.

### GetTransferToOk

`func (o *ModelsTransferTask) GetTransferToOk() (*string, bool)`

GetTransferToOk returns a tuple with the TransferTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTo

`func (o *ModelsTransferTask) SetTransferTo(v string)`

SetTransferTo sets TransferTo field to given value.

### HasTransferTo

`func (o *ModelsTransferTask) HasTransferTo() bool`

HasTransferTo returns a boolean if a field has been set.

### GetTxId

`func (o *ModelsTransferTask) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ModelsTransferTask) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ModelsTransferTask) SetTxId(v int32)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ModelsTransferTask) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsTransferTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTransferTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTransferTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsTransferTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


