# ModelsTBACreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | The chain type. The types include &#x60;conflux&#x60;&#x60; and &#x60;conflux_test&#x60;&#x60; | 
**CollectionToAdd** | Pointer to **string** | the collection to add after the tba is created. Will create the collection if it is not created yet. | [optional] 
**Salt** | Pointer to **string** | salt parameter to adjust tba address | [optional] 
**TokenContract** | **string** | address of the nft contract | 
**TokenId** | **string** | token id of nft | 

## Methods

### NewModelsTBACreateDto

`func NewModelsTBACreateDto(chain string, tokenContract string, tokenId string, ) *ModelsTBACreateDto`

NewModelsTBACreateDto instantiates a new ModelsTBACreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTBACreateDtoWithDefaults

`func NewModelsTBACreateDtoWithDefaults() *ModelsTBACreateDto`

NewModelsTBACreateDtoWithDefaults instantiates a new ModelsTBACreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ModelsTBACreateDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ModelsTBACreateDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ModelsTBACreateDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetCollectionToAdd

`func (o *ModelsTBACreateDto) GetCollectionToAdd() string`

GetCollectionToAdd returns the CollectionToAdd field if non-nil, zero value otherwise.

### GetCollectionToAddOk

`func (o *ModelsTBACreateDto) GetCollectionToAddOk() (*string, bool)`

GetCollectionToAddOk returns a tuple with the CollectionToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionToAdd

`func (o *ModelsTBACreateDto) SetCollectionToAdd(v string)`

SetCollectionToAdd sets CollectionToAdd field to given value.

### HasCollectionToAdd

`func (o *ModelsTBACreateDto) HasCollectionToAdd() bool`

HasCollectionToAdd returns a boolean if a field has been set.

### GetSalt

`func (o *ModelsTBACreateDto) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *ModelsTBACreateDto) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *ModelsTBACreateDto) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *ModelsTBACreateDto) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetTokenContract

`func (o *ModelsTBACreateDto) GetTokenContract() string`

GetTokenContract returns the TokenContract field if non-nil, zero value otherwise.

### GetTokenContractOk

`func (o *ModelsTBACreateDto) GetTokenContractOk() (*string, bool)`

GetTokenContractOk returns a tuple with the TokenContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContract

`func (o *ModelsTBACreateDto) SetTokenContract(v string)`

SetTokenContract sets TokenContract field to given value.


### GetTokenId

`func (o *ModelsTBACreateDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ModelsTBACreateDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ModelsTBACreateDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


