# ServicesTransferItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**TokenId** | **string** |  | 
**TransferFromAddress** | **string** |  | 
**TransferToAddress** | **string** |  | 

## Methods

### NewServicesTransferItemDto

`func NewServicesTransferItemDto(tokenId string, transferFromAddress string, transferToAddress string, ) *ServicesTransferItemDto`

NewServicesTransferItemDto instantiates a new ServicesTransferItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesTransferItemDtoWithDefaults

`func NewServicesTransferItemDtoWithDefaults() *ServicesTransferItemDto`

NewServicesTransferItemDtoWithDefaults instantiates a new ServicesTransferItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesTransferItemDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesTransferItemDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesTransferItemDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesTransferItemDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *ServicesTransferItemDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ServicesTransferItemDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ServicesTransferItemDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTransferFromAddress

`func (o *ServicesTransferItemDto) GetTransferFromAddress() string`

GetTransferFromAddress returns the TransferFromAddress field if non-nil, zero value otherwise.

### GetTransferFromAddressOk

`func (o *ServicesTransferItemDto) GetTransferFromAddressOk() (*string, bool)`

GetTransferFromAddressOk returns a tuple with the TransferFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFromAddress

`func (o *ServicesTransferItemDto) SetTransferFromAddress(v string)`

SetTransferFromAddress sets TransferFromAddress field to given value.


### GetTransferToAddress

`func (o *ServicesTransferItemDto) GetTransferToAddress() string`

GetTransferToAddress returns the TransferToAddress field if non-nil, zero value otherwise.

### GetTransferToAddressOk

`func (o *ServicesTransferItemDto) GetTransferToAddressOk() (*string, bool)`

GetTransferToAddressOk returns a tuple with the TransferToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferToAddress

`func (o *ServicesTransferItemDto) SetTransferToAddress(v string)`

SetTransferToAddress sets TransferToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


