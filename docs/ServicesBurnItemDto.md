# ServicesBurnItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**TokenId** | **string** |  | 

## Methods

### NewServicesBurnItemDto

`func NewServicesBurnItemDto(tokenId string, ) *ServicesBurnItemDto`

NewServicesBurnItemDto instantiates a new ServicesBurnItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesBurnItemDtoWithDefaults

`func NewServicesBurnItemDtoWithDefaults() *ServicesBurnItemDto`

NewServicesBurnItemDtoWithDefaults instantiates a new ServicesBurnItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesBurnItemDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesBurnItemDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesBurnItemDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicesBurnItemDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *ServicesBurnItemDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ServicesBurnItemDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ServicesBurnItemDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


