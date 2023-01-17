# ServicesBurnBatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**ContractType** | **string** |  | 
**Items** | [**[]ServicesBurnItemDto**](ServicesBurnItemDto.md) |  | 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesBurnBatchDto

`func NewServicesBurnBatchDto(chain string, contractAddress string, contractType string, items []ServicesBurnItemDto, ) *ServicesBurnBatchDto`

NewServicesBurnBatchDto instantiates a new ServicesBurnBatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesBurnBatchDtoWithDefaults

`func NewServicesBurnBatchDtoWithDefaults() *ServicesBurnBatchDto`

NewServicesBurnBatchDtoWithDefaults instantiates a new ServicesBurnBatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesBurnBatchDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesBurnBatchDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesBurnBatchDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesBurnBatchDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesBurnBatchDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesBurnBatchDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetContractType

`func (o *ServicesBurnBatchDto) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ServicesBurnBatchDto) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ServicesBurnBatchDto) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetItems

`func (o *ServicesBurnBatchDto) GetItems() []ServicesBurnItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServicesBurnBatchDto) GetItemsOk() (*[]ServicesBurnItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServicesBurnBatchDto) SetItems(v []ServicesBurnItemDto)`

SetItems sets Items field to given value.


### GetUser

`func (o *ServicesBurnBatchDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServicesBurnBatchDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServicesBurnBatchDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ServicesBurnBatchDto) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


