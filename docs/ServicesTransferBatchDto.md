# ServicesTransferBatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**ContractType** | **string** |  | 
**Items** | [**[]ServicesTransferItemDto**](ServicesTransferItemDto.md) |  | 

## Methods

### NewServicesTransferBatchDto

`func NewServicesTransferBatchDto(chain string, contractAddress string, contractType string, items []ServicesTransferItemDto, ) *ServicesTransferBatchDto`

NewServicesTransferBatchDto instantiates a new ServicesTransferBatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesTransferBatchDtoWithDefaults

`func NewServicesTransferBatchDtoWithDefaults() *ServicesTransferBatchDto`

NewServicesTransferBatchDtoWithDefaults instantiates a new ServicesTransferBatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesTransferBatchDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesTransferBatchDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesTransferBatchDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesTransferBatchDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesTransferBatchDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesTransferBatchDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetContractType

`func (o *ServicesTransferBatchDto) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ServicesTransferBatchDto) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ServicesTransferBatchDto) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetItems

`func (o *ServicesTransferBatchDto) GetItems() []ServicesTransferItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServicesTransferBatchDto) GetItemsOk() (*[]ServicesTransferItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServicesTransferBatchDto) SetItems(v []ServicesTransferItemDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


