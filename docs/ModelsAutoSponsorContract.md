# ModelsAutoSponsorContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AutoSponsor** | Pointer to **bool** |  | [optional] 
**ContractId** | Pointer to **int32** | useless | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**StorageRechargeAmount** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsAutoSponsorContract

`func NewModelsAutoSponsorContract() *ModelsAutoSponsorContract`

NewModelsAutoSponsorContract instantiates a new ModelsAutoSponsorContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAutoSponsorContractWithDefaults

`func NewModelsAutoSponsorContractWithDefaults() *ModelsAutoSponsorContract`

NewModelsAutoSponsorContractWithDefaults instantiates a new ModelsAutoSponsorContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ModelsAutoSponsorContract) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelsAutoSponsorContract) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelsAutoSponsorContract) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelsAutoSponsorContract) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAutoSponsor

`func (o *ModelsAutoSponsorContract) GetAutoSponsor() bool`

GetAutoSponsor returns the AutoSponsor field if non-nil, zero value otherwise.

### GetAutoSponsorOk

`func (o *ModelsAutoSponsorContract) GetAutoSponsorOk() (*bool, bool)`

GetAutoSponsorOk returns a tuple with the AutoSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSponsor

`func (o *ModelsAutoSponsorContract) SetAutoSponsor(v bool)`

SetAutoSponsor sets AutoSponsor field to given value.

### HasAutoSponsor

`func (o *ModelsAutoSponsorContract) HasAutoSponsor() bool`

HasAutoSponsor returns a boolean if a field has been set.

### GetContractId

`func (o *ModelsAutoSponsorContract) GetContractId() int32`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ModelsAutoSponsorContract) GetContractIdOk() (*int32, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ModelsAutoSponsorContract) SetContractId(v int32)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ModelsAutoSponsorContract) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsAutoSponsorContract) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAutoSponsorContract) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAutoSponsorContract) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsAutoSponsorContract) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsAutoSponsorContract) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsAutoSponsorContract) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsAutoSponsorContract) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsAutoSponsorContract) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsAutoSponsorContract) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAutoSponsorContract) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAutoSponsorContract) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAutoSponsorContract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStorageRechargeAmount

`func (o *ModelsAutoSponsorContract) GetStorageRechargeAmount() int32`

GetStorageRechargeAmount returns the StorageRechargeAmount field if non-nil, zero value otherwise.

### GetStorageRechargeAmountOk

`func (o *ModelsAutoSponsorContract) GetStorageRechargeAmountOk() (*int32, bool)`

GetStorageRechargeAmountOk returns a tuple with the StorageRechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRechargeAmount

`func (o *ModelsAutoSponsorContract) SetStorageRechargeAmount(v int32)`

SetStorageRechargeAmount sets StorageRechargeAmount field to given value.

### HasStorageRechargeAmount

`func (o *ModelsAutoSponsorContract) HasStorageRechargeAmount() bool`

HasStorageRechargeAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsAutoSponsorContract) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsAutoSponsorContract) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsAutoSponsorContract) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsAutoSponsorContract) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ModelsAutoSponsorContract) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelsAutoSponsorContract) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelsAutoSponsorContract) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelsAutoSponsorContract) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


