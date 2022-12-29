# ServicesEasyMintFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**Description** | **string** |  | 
**File** | [**MultipartFileHeader**](MultipartFileHeader.md) |  | 
**MintToAddress** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewServicesEasyMintFileDto

`func NewServicesEasyMintFileDto(chain string, description string, file MultipartFileHeader, mintToAddress string, name string, ) *ServicesEasyMintFileDto`

NewServicesEasyMintFileDto instantiates a new ServicesEasyMintFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesEasyMintFileDtoWithDefaults

`func NewServicesEasyMintFileDtoWithDefaults() *ServicesEasyMintFileDto`

NewServicesEasyMintFileDtoWithDefaults instantiates a new ServicesEasyMintFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesEasyMintFileDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesEasyMintFileDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesEasyMintFileDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetDescription

`func (o *ServicesEasyMintFileDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesEasyMintFileDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesEasyMintFileDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFile

`func (o *ServicesEasyMintFileDto) GetFile() MultipartFileHeader`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ServicesEasyMintFileDto) GetFileOk() (*MultipartFileHeader, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ServicesEasyMintFileDto) SetFile(v MultipartFileHeader)`

SetFile sets File field to given value.


### GetMintToAddress

`func (o *ServicesEasyMintFileDto) GetMintToAddress() string`

GetMintToAddress returns the MintToAddress field if non-nil, zero value otherwise.

### GetMintToAddressOk

`func (o *ServicesEasyMintFileDto) GetMintToAddressOk() (*string, bool)`

GetMintToAddressOk returns a tuple with the MintToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintToAddress

`func (o *ServicesEasyMintFileDto) SetMintToAddress(v string)`

SetMintToAddress sets MintToAddress field to given value.


### GetName

`func (o *ServicesEasyMintFileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesEasyMintFileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesEasyMintFileDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


