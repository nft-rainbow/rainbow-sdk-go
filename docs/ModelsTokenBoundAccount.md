# ModelsTokenBoundAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the tba | [optional] 
**ChainId** | Pointer to **int32** | The chainId of the TBA | [optional] 
**ChainType** | Pointer to **int32** | no need add to unique index, because conflux core testnet and eth mainnet has different address | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Implementation** | Pointer to **string** | Account logic implementation | [optional] 
**Salt** | Pointer to **string** | Salt parameter that adjusts the tba address | [optional] 
**TokenContract** | Pointer to **string** | The address of the nft contract | [optional] 
**TokenContractChainId** | Pointer to **int32** | NOTE: this is not the chain id of contract but the erc6551 defined chainid | [optional] 
**TokenId** | Pointer to **string** | Token id of the nft | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsTokenBoundAccount

`func NewModelsTokenBoundAccount() *ModelsTokenBoundAccount`

NewModelsTokenBoundAccount instantiates a new ModelsTokenBoundAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTokenBoundAccountWithDefaults

`func NewModelsTokenBoundAccountWithDefaults() *ModelsTokenBoundAccount`

NewModelsTokenBoundAccountWithDefaults instantiates a new ModelsTokenBoundAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ModelsTokenBoundAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelsTokenBoundAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelsTokenBoundAccount) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelsTokenBoundAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChainId

`func (o *ModelsTokenBoundAccount) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ModelsTokenBoundAccount) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ModelsTokenBoundAccount) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ModelsTokenBoundAccount) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainType

`func (o *ModelsTokenBoundAccount) GetChainType() int32`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ModelsTokenBoundAccount) GetChainTypeOk() (*int32, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ModelsTokenBoundAccount) SetChainType(v int32)`

SetChainType sets ChainType field to given value.

### HasChainType

`func (o *ModelsTokenBoundAccount) HasChainType() bool`

HasChainType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsTokenBoundAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsTokenBoundAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsTokenBoundAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsTokenBoundAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsTokenBoundAccount) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsTokenBoundAccount) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsTokenBoundAccount) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsTokenBoundAccount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsTokenBoundAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTokenBoundAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTokenBoundAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTokenBoundAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImplementation

`func (o *ModelsTokenBoundAccount) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *ModelsTokenBoundAccount) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *ModelsTokenBoundAccount) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *ModelsTokenBoundAccount) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### GetSalt

`func (o *ModelsTokenBoundAccount) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *ModelsTokenBoundAccount) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *ModelsTokenBoundAccount) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *ModelsTokenBoundAccount) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetTokenContract

`func (o *ModelsTokenBoundAccount) GetTokenContract() string`

GetTokenContract returns the TokenContract field if non-nil, zero value otherwise.

### GetTokenContractOk

`func (o *ModelsTokenBoundAccount) GetTokenContractOk() (*string, bool)`

GetTokenContractOk returns a tuple with the TokenContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContract

`func (o *ModelsTokenBoundAccount) SetTokenContract(v string)`

SetTokenContract sets TokenContract field to given value.

### HasTokenContract

`func (o *ModelsTokenBoundAccount) HasTokenContract() bool`

HasTokenContract returns a boolean if a field has been set.

### GetTokenContractChainId

`func (o *ModelsTokenBoundAccount) GetTokenContractChainId() int32`

GetTokenContractChainId returns the TokenContractChainId field if non-nil, zero value otherwise.

### GetTokenContractChainIdOk

`func (o *ModelsTokenBoundAccount) GetTokenContractChainIdOk() (*int32, bool)`

GetTokenContractChainIdOk returns a tuple with the TokenContractChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContractChainId

`func (o *ModelsTokenBoundAccount) SetTokenContractChainId(v int32)`

SetTokenContractChainId sets TokenContractChainId field to given value.

### HasTokenContractChainId

`func (o *ModelsTokenBoundAccount) HasTokenContractChainId() bool`

HasTokenContractChainId returns a boolean if a field has been set.

### GetTokenId

`func (o *ModelsTokenBoundAccount) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsTokenBoundAccount) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsTokenBoundAccount) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ModelsTokenBoundAccount) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsTokenBoundAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTokenBoundAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTokenBoundAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsTokenBoundAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


