# ModelsTBACreationTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **int32** |  | [optional] 
**BlockReason** | Pointer to [**EnumsTransactionBlockReason**](EnumsTransactionBlockReason.md) |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**ChainType** | Pointer to **int32** |  | [optional] 
**CollectionToAdd** | Pointer to **string** | the collection to add after the tba is created | [optional] 
**Contract** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Implementation** | Pointer to **string** | tba implementation address | [optional] 
**PrecomputedAddress** | Pointer to **string** | the precomputed tba address | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failed | [optional] 
**TokenContract** | Pointer to **string** | the nft token contract address | [optional] 
**TokenContractChainId** | Pointer to **int32** | the chain id of the token contract | [optional] 
**TokenId** | Pointer to **string** | the token id of the nft | [optional] 
**TxId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** | the user created the task | [optional] 

## Methods

### NewModelsTBACreationTask

`func NewModelsTBACreationTask() *ModelsTBACreationTask`

NewModelsTBACreationTask instantiates a new ModelsTBACreationTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTBACreationTaskWithDefaults

`func NewModelsTBACreationTaskWithDefaults() *ModelsTBACreationTask`

NewModelsTBACreationTaskWithDefaults instantiates a new ModelsTBACreationTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ModelsTBACreationTask) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsTBACreationTask) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsTBACreationTask) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsTBACreationTask) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBlockReason

`func (o *ModelsTBACreationTask) GetBlockReason() EnumsTransactionBlockReason`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *ModelsTBACreationTask) GetBlockReasonOk() (*EnumsTransactionBlockReason, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *ModelsTBACreationTask) SetBlockReason(v EnumsTransactionBlockReason)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *ModelsTBACreationTask) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsTBACreationTask) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsTBACreationTask) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsTBACreationTask) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsTBACreationTask) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsTBACreationTask) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsTBACreationTask) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsTBACreationTask) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsTBACreationTask) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetCollectionToAdd

`func (o *ModelsTBACreationTask) GetCollectionToAdd() string`

GetCollectionToAdd returns the CollectionToAdd field if non-nil, zero value otherwise.

### GetCollectionToAddOk

`func (o *ModelsTBACreationTask) GetCollectionToAddOk() (*string, bool)`

GetCollectionToAddOk returns a tuple with the CollectionToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionToAdd

`func (o *ModelsTBACreationTask) SetCollectionToAdd(v string)`

SetCollectionToAdd sets CollectionToAdd field to given value.

### HasCollectionToAdd

`func (o *ModelsTBACreationTask) HasCollectionToAdd() bool`

HasCollectionToAdd returns a boolean if a field has been set.

### GetContract

`func (o *ModelsTBACreationTask) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ModelsTBACreationTask) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ModelsTBACreationTask) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ModelsTBACreationTask) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetContractType

`func (o *ModelsTBACreationTask) GetContractType() int32`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ModelsTBACreationTask) GetContractTypeOk() (*int32, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ModelsTBACreationTask) SetContractType(v int32)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ModelsTBACreationTask) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsTBACreationTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsTBACreationTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsTBACreationTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsTBACreationTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsTBACreationTask) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsTBACreationTask) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsTBACreationTask) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsTBACreationTask) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetError

`func (o *ModelsTBACreationTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsTBACreationTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsTBACreationTask) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModelsTBACreationTask) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHash

`func (o *ModelsTBACreationTask) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsTBACreationTask) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsTBACreationTask) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsTBACreationTask) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsTBACreationTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTBACreationTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTBACreationTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTBACreationTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImplementation

`func (o *ModelsTBACreationTask) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *ModelsTBACreationTask) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *ModelsTBACreationTask) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *ModelsTBACreationTask) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### GetPrecomputedAddress

`func (o *ModelsTBACreationTask) GetPrecomputedAddress() string`

GetPrecomputedAddress returns the PrecomputedAddress field if non-nil, zero value otherwise.

### GetPrecomputedAddressOk

`func (o *ModelsTBACreationTask) GetPrecomputedAddressOk() (*string, bool)`

GetPrecomputedAddressOk returns a tuple with the PrecomputedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecomputedAddress

`func (o *ModelsTBACreationTask) SetPrecomputedAddress(v string)`

SetPrecomputedAddress sets PrecomputedAddress field to given value.

### HasPrecomputedAddress

`func (o *ModelsTBACreationTask) HasPrecomputedAddress() bool`

HasPrecomputedAddress returns a boolean if a field has been set.

### GetSalt

`func (o *ModelsTBACreationTask) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *ModelsTBACreationTask) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *ModelsTBACreationTask) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *ModelsTBACreationTask) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsTBACreationTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsTBACreationTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsTBACreationTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsTBACreationTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenContract

`func (o *ModelsTBACreationTask) GetTokenContract() string`

GetTokenContract returns the TokenContract field if non-nil, zero value otherwise.

### GetTokenContractOk

`func (o *ModelsTBACreationTask) GetTokenContractOk() (*string, bool)`

GetTokenContractOk returns a tuple with the TokenContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContract

`func (o *ModelsTBACreationTask) SetTokenContract(v string)`

SetTokenContract sets TokenContract field to given value.

### HasTokenContract

`func (o *ModelsTBACreationTask) HasTokenContract() bool`

HasTokenContract returns a boolean if a field has been set.

### GetTokenContractChainId

`func (o *ModelsTBACreationTask) GetTokenContractChainId() int32`

GetTokenContractChainId returns the TokenContractChainId field if non-nil, zero value otherwise.

### GetTokenContractChainIdOk

`func (o *ModelsTBACreationTask) GetTokenContractChainIdOk() (*int32, bool)`

GetTokenContractChainIdOk returns a tuple with the TokenContractChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContractChainId

`func (o *ModelsTBACreationTask) SetTokenContractChainId(v int32)`

SetTokenContractChainId sets TokenContractChainId field to given value.

### HasTokenContractChainId

`func (o *ModelsTBACreationTask) HasTokenContractChainId() bool`

HasTokenContractChainId returns a boolean if a field has been set.

### GetTokenId

`func (o *ModelsTBACreationTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsTBACreationTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsTBACreationTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ModelsTBACreationTask) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTxId

`func (o *ModelsTBACreationTask) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ModelsTBACreationTask) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ModelsTBACreationTask) SetTxId(v int32)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ModelsTBACreationTask) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsTBACreationTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTBACreationTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTBACreationTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsTBACreationTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ModelsTBACreationTask) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelsTBACreationTask) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelsTBACreationTask) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelsTBACreationTask) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


