# ServicesMintItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**MetadataUri** | Pointer to **string** |  | [optional] 
**MintToAddress** | **string** |  | 
**TokenId** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesMintItemDto

`func NewServicesMintItemDto(mintToAddress string, ) *ServicesMintItemDto`

NewServicesMintItemDto instantiates a new ServicesMintItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesMintItemDtoWithDefaults

`func NewServicesMintItemDtoWithDefaults() *ServicesMintItemDto`

NewServicesMintItemDtoWithDefaults instantiates a new ServicesMintItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesMintItemDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesMintItemDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesMintItemDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesMintItemDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadataUri

`func (o *ServicesMintItemDto) GetMetadataUri() string`

GetMetadataUri returns the MetadataUri field if non-nil, zero value otherwise.

### GetMetadataUriOk

`func (o *ServicesMintItemDto) GetMetadataUriOk() (*string, bool)`

GetMetadataUriOk returns a tuple with the MetadataUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUri

`func (o *ServicesMintItemDto) SetMetadataUri(v string)`

SetMetadataUri sets MetadataUri field to given value.

### HasMetadataUri

`func (o *ServicesMintItemDto) HasMetadataUri() bool`

HasMetadataUri returns a boolean if a field has been set.

### GetMintToAddress

`func (o *ServicesMintItemDto) GetMintToAddress() string`

GetMintToAddress returns the MintToAddress field if non-nil, zero value otherwise.

### GetMintToAddressOk

`func (o *ServicesMintItemDto) GetMintToAddressOk() (*string, bool)`

GetMintToAddressOk returns a tuple with the MintToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintToAddress

`func (o *ServicesMintItemDto) SetMintToAddress(v string)`

SetMintToAddress sets MintToAddress field to given value.


### GetTokenId

`func (o *ServicesMintItemDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ServicesMintItemDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ServicesMintItemDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ServicesMintItemDto) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


