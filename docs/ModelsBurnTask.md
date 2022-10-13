# ModelsBurnTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
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
**TxId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsBurnTask

`func NewModelsBurnTask() *ModelsBurnTask`

NewModelsBurnTask instantiates a new ModelsBurnTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBurnTaskWithDefaults

`func NewModelsBurnTaskWithDefaults() *ModelsBurnTask`

NewModelsBurnTaskWithDefaults instantiates a new ModelsBurnTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsBurnTask) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsBurnTask) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsBurnTask) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsBurnTask) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppId

`func (o *ModelsBurnTask) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelsBurnTask) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelsBurnTask) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelsBurnTask) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsBurnTask) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsBurnTask) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsBurnTask) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsBurnTask) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsBurnTask) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsBurnTask) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsBurnTask) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsBurnTask) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetContract

`func (o *ModelsBurnTask) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ModelsBurnTask) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ModelsBurnTask) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ModelsBurnTask) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetContractType

`func (o *ModelsBurnTask) GetContractType() int32`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ModelsBurnTask) GetContractTypeOk() (*int32, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ModelsBurnTask) SetContractType(v int32)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ModelsBurnTask) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsBurnTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsBurnTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsBurnTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsBurnTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsBurnTask) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsBurnTask) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsBurnTask) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsBurnTask) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetError

`func (o *ModelsBurnTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsBurnTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsBurnTask) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModelsBurnTask) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHash

`func (o *ModelsBurnTask) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsBurnTask) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsBurnTask) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsBurnTask) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsBurnTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsBurnTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsBurnTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsBurnTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsBurnTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsBurnTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsBurnTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsBurnTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenId

`func (o *ModelsBurnTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsBurnTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsBurnTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ModelsBurnTask) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTxId

`func (o *ModelsBurnTask) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ModelsBurnTask) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ModelsBurnTask) SetTxId(v int32)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ModelsBurnTask) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsBurnTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsBurnTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsBurnTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsBurnTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *ModelsBurnTask) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelsBurnTask) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelsBurnTask) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelsBurnTask) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


