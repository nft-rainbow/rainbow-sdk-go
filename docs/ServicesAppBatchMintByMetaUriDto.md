# ServicesAppBatchMintByMetaUriDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**ContractAddress** | **string** |  | 
**MintItems** | [**[]ServicesMintItemDto**](ServicesMintItemDto.md) |  | 

## Methods

### NewServicesAppBatchMintByMetaUriDto

`func NewServicesAppBatchMintByMetaUriDto(chain string, contractAddress string, mintItems []ServicesMintItemDto, ) *ServicesAppBatchMintByMetaUriDto`

NewServicesAppBatchMintByMetaUriDto instantiates a new ServicesAppBatchMintByMetaUriDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesAppBatchMintByMetaUriDtoWithDefaults

`func NewServicesAppBatchMintByMetaUriDtoWithDefaults() *ServicesAppBatchMintByMetaUriDto`

NewServicesAppBatchMintByMetaUriDtoWithDefaults instantiates a new ServicesAppBatchMintByMetaUriDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesAppBatchMintByMetaUriDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesAppBatchMintByMetaUriDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesAppBatchMintByMetaUriDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetContractAddress

`func (o *ServicesAppBatchMintByMetaUriDto) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ServicesAppBatchMintByMetaUriDto) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ServicesAppBatchMintByMetaUriDto) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetMintItems

`func (o *ServicesAppBatchMintByMetaUriDto) GetMintItems() []ServicesMintItemDto`

GetMintItems returns the MintItems field if non-nil, zero value otherwise.

### GetMintItemsOk

`func (o *ServicesAppBatchMintByMetaUriDto) GetMintItemsOk() (*[]ServicesMintItemDto, bool)`

GetMintItemsOk returns a tuple with the MintItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintItems

`func (o *ServicesAppBatchMintByMetaUriDto) SetMintItems(v []ServicesMintItemDto)`

SetMintItems sets MintItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


