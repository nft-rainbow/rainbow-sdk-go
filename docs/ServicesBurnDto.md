# ServicesBurnDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**ContractType** | **string** |  | 
**TokenId** | **string** |  | 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesBurnDto

`func NewServicesBurnDto(chain string, contractAddress string, contractType string, tokenId string, ) *ServicesBurnDto`

NewServicesBurnDto instantiates a new ServicesBurnDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesBurnDtoWithDefaults

`func NewServicesBurnDtoWithDefaults() *ServicesBurnDto`

NewServicesBurnDtoWithDefaults instantiates a new ServicesBurnDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesBurnDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesBurnDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesBurnDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesBurnDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChain

`func (o *ServicesBurnDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesBurnDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesBurnDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesBurnDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesBurnDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesBurnDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetContractType

`func (o *ServicesBurnDto) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ServicesBurnDto) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ServicesBurnDto) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetTokenId

`func (o *ServicesBurnDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ServicesBurnDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ServicesBurnDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetUser

`func (o *ServicesBurnDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServicesBurnDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServicesBurnDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ServicesBurnDto) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


