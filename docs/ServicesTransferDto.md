# ServicesTransferDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**ContractType** | **string** |  | 
**TokenId** | **string** |  | 
**TransferFromAddress** | **string** |  | 
**TransferToAddress** | **string** |  | 

## Methods

### NewServicesTransferDto

`func NewServicesTransferDto(chain string, contractAddress string, contractType string, tokenId string, transferFromAddress string, transferToAddress string, ) *ServicesTransferDto`

NewServicesTransferDto instantiates a new ServicesTransferDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesTransferDtoWithDefaults

`func NewServicesTransferDtoWithDefaults() *ServicesTransferDto`

NewServicesTransferDtoWithDefaults instantiates a new ServicesTransferDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesTransferDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesTransferDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesTransferDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesTransferDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChain

`func (o *ServicesTransferDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesTransferDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesTransferDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesTransferDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesTransferDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesTransferDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetContractType

`func (o *ServicesTransferDto) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ServicesTransferDto) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ServicesTransferDto) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetTokenId

`func (o *ServicesTransferDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ServicesTransferDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ServicesTransferDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTransferFromAddress

`func (o *ServicesTransferDto) GetTransferFromAddress() string`

GetTransferFromAddress returns the TransferFromAddress field if non-nil, zero value otherwise.

### GetTransferFromAddressOk

`func (o *ServicesTransferDto) GetTransferFromAddressOk() (*string, bool)`

GetTransferFromAddressOk returns a tuple with the TransferFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFromAddress

`func (o *ServicesTransferDto) SetTransferFromAddress(v string)`

SetTransferFromAddress sets TransferFromAddress field to given value.


### GetTransferToAddress

`func (o *ServicesTransferDto) GetTransferToAddress() string`

GetTransferToAddress returns the TransferToAddress field if non-nil, zero value otherwise.

### GetTransferToAddressOk

`func (o *ServicesTransferDto) GetTransferToAddressOk() (*string, bool)`

GetTransferToAddressOk returns a tuple with the TransferToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferToAddress

`func (o *ServicesTransferDto) SetTransferToAddress(v string)`

SetTransferToAddress sets TransferToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


