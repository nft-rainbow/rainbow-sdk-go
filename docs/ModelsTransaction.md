# ModelsTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** |  | [optional] 
**ChainType** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**EpochHeight** | Pointer to **float32** |  | [optional] 
**EpochNumber** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Gas** | Pointer to **float32** |  | [optional] 
**GasCoveredBySponsor** | Pointer to **bool** |  | [optional] 
**GasFee** | Pointer to **float32** |  | [optional] 
**GasPrice** | Pointer to **float32** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**PendingReason** | Pointer to **string** |  | [optional] 
**PreHashs** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StorageCollateralized** | Pointer to **float32** |  | [optional] 
**StorageCoveredBySponsor** | Pointer to **bool** |  | [optional] 
**StorageLimit** | Pointer to **float32** |  | [optional] 
**TaskType** | Pointer to [**ModelsTaskType**](ModelsTaskType.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TryReceiptCnt** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewModelsTransaction

`func NewModelsTransaction() *ModelsTransaction`

NewModelsTransaction instantiates a new ModelsTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTransactionWithDefaults

`func NewModelsTransactionWithDefaults() *ModelsTransaction`

NewModelsTransactionWithDefaults instantiates a new ModelsTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ModelsTransaction) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsTransaction) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsTransaction) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsTransaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsTransaction) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsTransaction) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsTransaction) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsTransaction) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *ModelsTransaction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelsTransaction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelsTransaction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ModelsTransaction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsTransaction) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsTransaction) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsTransaction) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsTransaction) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEpochHeight

`func (o *ModelsTransaction) GetEpochHeight() float32`

GetEpochHeight returns the EpochHeight field if non-nil, zero value otherwise.

### GetEpochHeightOk

`func (o *ModelsTransaction) GetEpochHeightOk() (*float32, bool)`

GetEpochHeightOk returns a tuple with the EpochHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochHeight

`func (o *ModelsTransaction) SetEpochHeight(v float32)`

SetEpochHeight sets EpochHeight field to given value.

### HasEpochHeight

`func (o *ModelsTransaction) HasEpochHeight() bool`

HasEpochHeight returns a boolean if a field has been set.

### GetEpochNumber

`func (o *ModelsTransaction) GetEpochNumber() int32`

GetEpochNumber returns the EpochNumber field if non-nil, zero value otherwise.

### GetEpochNumberOk

`func (o *ModelsTransaction) GetEpochNumberOk() (*int32, bool)`

GetEpochNumberOk returns a tuple with the EpochNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNumber

`func (o *ModelsTransaction) SetEpochNumber(v int32)`

SetEpochNumber sets EpochNumber field to given value.

### HasEpochNumber

`func (o *ModelsTransaction) HasEpochNumber() bool`

HasEpochNumber returns a boolean if a field has been set.

### GetError

`func (o *ModelsTransaction) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsTransaction) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsTransaction) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModelsTransaction) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFrom

`func (o *ModelsTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ModelsTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ModelsTransaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ModelsTransaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *ModelsTransaction) GetGas() float32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *ModelsTransaction) GetGasOk() (*float32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *ModelsTransaction) SetGas(v float32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *ModelsTransaction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasCoveredBySponsor

`func (o *ModelsTransaction) GetGasCoveredBySponsor() bool`

GetGasCoveredBySponsor returns the GasCoveredBySponsor field if non-nil, zero value otherwise.

### GetGasCoveredBySponsorOk

`func (o *ModelsTransaction) GetGasCoveredBySponsorOk() (*bool, bool)`

GetGasCoveredBySponsorOk returns a tuple with the GasCoveredBySponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCoveredBySponsor

`func (o *ModelsTransaction) SetGasCoveredBySponsor(v bool)`

SetGasCoveredBySponsor sets GasCoveredBySponsor field to given value.

### HasGasCoveredBySponsor

`func (o *ModelsTransaction) HasGasCoveredBySponsor() bool`

HasGasCoveredBySponsor returns a boolean if a field has been set.

### GetGasFee

`func (o *ModelsTransaction) GetGasFee() float32`

GetGasFee returns the GasFee field if non-nil, zero value otherwise.

### GetGasFeeOk

`func (o *ModelsTransaction) GetGasFeeOk() (*float32, bool)`

GetGasFeeOk returns a tuple with the GasFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFee

`func (o *ModelsTransaction) SetGasFee(v float32)`

SetGasFee sets GasFee field to given value.

### HasGasFee

`func (o *ModelsTransaction) HasGasFee() bool`

HasGasFee returns a boolean if a field has been set.

### GetGasPrice

`func (o *ModelsTransaction) GetGasPrice() float32`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ModelsTransaction) GetGasPriceOk() (*float32, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ModelsTransaction) SetGasPrice(v float32)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ModelsTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetHash

`func (o *ModelsTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsTransaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsTransaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTransaction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNonce

`func (o *ModelsTransaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ModelsTransaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ModelsTransaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ModelsTransaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPendingReason

`func (o *ModelsTransaction) GetPendingReason() string`

GetPendingReason returns the PendingReason field if non-nil, zero value otherwise.

### GetPendingReasonOk

`func (o *ModelsTransaction) GetPendingReasonOk() (*string, bool)`

GetPendingReasonOk returns a tuple with the PendingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReason

`func (o *ModelsTransaction) SetPendingReason(v string)`

SetPendingReason sets PendingReason field to given value.

### HasPendingReason

`func (o *ModelsTransaction) HasPendingReason() bool`

HasPendingReason returns a boolean if a field has been set.

### GetPreHashs

`func (o *ModelsTransaction) GetPreHashs() string`

GetPreHashs returns the PreHashs field if non-nil, zero value otherwise.

### GetPreHashsOk

`func (o *ModelsTransaction) GetPreHashsOk() (*string, bool)`

GetPreHashsOk returns a tuple with the PreHashs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreHashs

`func (o *ModelsTransaction) SetPreHashs(v string)`

SetPreHashs sets PreHashs field to given value.

### HasPreHashs

`func (o *ModelsTransaction) HasPreHashs() bool`

HasPreHashs returns a boolean if a field has been set.

### GetStage

`func (o *ModelsTransaction) GetStage() int32`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ModelsTransaction) GetStageOk() (*int32, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ModelsTransaction) SetStage(v int32)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ModelsTransaction) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *ModelsTransaction) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsTransaction) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsTransaction) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *ModelsTransaction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsTransaction) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsTransaction) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsTransaction) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageCollateralized

`func (o *ModelsTransaction) GetStorageCollateralized() float32`

GetStorageCollateralized returns the StorageCollateralized field if non-nil, zero value otherwise.

### GetStorageCollateralizedOk

`func (o *ModelsTransaction) GetStorageCollateralizedOk() (*float32, bool)`

GetStorageCollateralizedOk returns a tuple with the StorageCollateralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCollateralized

`func (o *ModelsTransaction) SetStorageCollateralized(v float32)`

SetStorageCollateralized sets StorageCollateralized field to given value.

### HasStorageCollateralized

`func (o *ModelsTransaction) HasStorageCollateralized() bool`

HasStorageCollateralized returns a boolean if a field has been set.

### GetStorageCoveredBySponsor

`func (o *ModelsTransaction) GetStorageCoveredBySponsor() bool`

GetStorageCoveredBySponsor returns the StorageCoveredBySponsor field if non-nil, zero value otherwise.

### GetStorageCoveredBySponsorOk

`func (o *ModelsTransaction) GetStorageCoveredBySponsorOk() (*bool, bool)`

GetStorageCoveredBySponsorOk returns a tuple with the StorageCoveredBySponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCoveredBySponsor

`func (o *ModelsTransaction) SetStorageCoveredBySponsor(v bool)`

SetStorageCoveredBySponsor sets StorageCoveredBySponsor field to given value.

### HasStorageCoveredBySponsor

`func (o *ModelsTransaction) HasStorageCoveredBySponsor() bool`

HasStorageCoveredBySponsor returns a boolean if a field has been set.

### GetStorageLimit

`func (o *ModelsTransaction) GetStorageLimit() float32`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *ModelsTransaction) GetStorageLimitOk() (*float32, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *ModelsTransaction) SetStorageLimit(v float32)`

SetStorageLimit sets StorageLimit field to given value.

### HasStorageLimit

`func (o *ModelsTransaction) HasStorageLimit() bool`

HasStorageLimit returns a boolean if a field has been set.

### GetTaskType

`func (o *ModelsTransaction) GetTaskType() ModelsTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ModelsTransaction) GetTaskTypeOk() (*ModelsTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ModelsTransaction) SetTaskType(v ModelsTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *ModelsTransaction) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetTo

`func (o *ModelsTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ModelsTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ModelsTransaction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ModelsTransaction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTryReceiptCnt

`func (o *ModelsTransaction) GetTryReceiptCnt() int32`

GetTryReceiptCnt returns the TryReceiptCnt field if non-nil, zero value otherwise.

### GetTryReceiptCntOk

`func (o *ModelsTransaction) GetTryReceiptCntOk() (*int32, bool)`

GetTryReceiptCntOk returns a tuple with the TryReceiptCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryReceiptCnt

`func (o *ModelsTransaction) SetTryReceiptCnt(v int32)`

SetTryReceiptCnt sets TryReceiptCnt field to given value.

### HasTryReceiptCnt

`func (o *ModelsTransaction) HasTryReceiptCnt() bool`

HasTryReceiptCnt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsTransaction) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTransaction) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTransaction) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsTransaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *ModelsTransaction) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelsTransaction) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelsTransaction) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelsTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


