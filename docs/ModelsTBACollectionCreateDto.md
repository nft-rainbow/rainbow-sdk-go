# ModelsTBACollectionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | The chain type. The types include &#x60;conflux&#x60;&#x60; and &#x60;conflux_test&#x60;&#x60; | 
**Description** | Pointer to **string** | The collection description | [optional] 
**Name** | **string** | The name of the collection. Should be unique for the user. | 

## Methods

### NewModelsTBACollectionCreateDto

`func NewModelsTBACollectionCreateDto(chain string, name string, ) *ModelsTBACollectionCreateDto`

NewModelsTBACollectionCreateDto instantiates a new ModelsTBACollectionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTBACollectionCreateDtoWithDefaults

`func NewModelsTBACollectionCreateDtoWithDefaults() *ModelsTBACollectionCreateDto`

NewModelsTBACollectionCreateDtoWithDefaults instantiates a new ModelsTBACollectionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ModelsTBACollectionCreateDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ModelsTBACollectionCreateDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ModelsTBACollectionCreateDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetDescription

`func (o *ModelsTBACollectionCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsTBACollectionCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsTBACollectionCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsTBACollectionCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ModelsTBACollectionCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTBACollectionCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTBACollectionCreateDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


