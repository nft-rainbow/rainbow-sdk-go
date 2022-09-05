# ServicesCustomMintBatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**ContractType** | **string** |  | 
**MintItems** | [**[]ServicesMintItemDto**](ServicesMintItemDto.md) |  | 

## Methods

### NewServicesCustomMintBatchDto

`func NewServicesCustomMintBatchDto(chain string, contractAddress string, contractType string, mintItems []ServicesMintItemDto, ) *ServicesCustomMintBatchDto`

NewServicesCustomMintBatchDto instantiates a new ServicesCustomMintBatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesCustomMintBatchDtoWithDefaults

`func NewServicesCustomMintBatchDtoWithDefaults() *ServicesCustomMintBatchDto`

NewServicesCustomMintBatchDtoWithDefaults instantiates a new ServicesCustomMintBatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesCustomMintBatchDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesCustomMintBatchDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesCustomMintBatchDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesCustomMintBatchDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesCustomMintBatchDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesCustomMintBatchDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetContractType

`func (o *ServicesCustomMintBatchDto) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ServicesCustomMintBatchDto) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ServicesCustomMintBatchDto) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetMintItems

`func (o *ServicesCustomMintBatchDto) GetMintItems() []ServicesMintItemDto`

GetMintItems returns the MintItems field if non-nil, zero value otherwise.

### GetMintItemsOk

`func (o *ServicesCustomMintBatchDto) GetMintItemsOk() (*[]ServicesMintItemDto, bool)`

GetMintItemsOk returns a tuple with the MintItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintItems

`func (o *ServicesCustomMintBatchDto) SetMintItems(v []ServicesMintItemDto)`

SetMintItems sets MintItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


