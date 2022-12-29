# ServicesMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnimationUrl** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]ModelsExposedMetadataAttribute**](ModelsExposedMetadataAttribute.md) |  | [optional] 
**Description** | **string** |  | 
**ExternalLink** | Pointer to **string** |  | [optional] 
**Image** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewServicesMetadataDto

`func NewServicesMetadataDto(description string, image string, name string, ) *ServicesMetadataDto`

NewServicesMetadataDto instantiates a new ServicesMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesMetadataDtoWithDefaults

`func NewServicesMetadataDtoWithDefaults() *ServicesMetadataDto`

NewServicesMetadataDtoWithDefaults instantiates a new ServicesMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnimationUrl

`func (o *ServicesMetadataDto) GetAnimationUrl() string`

GetAnimationUrl returns the AnimationUrl field if non-nil, zero value otherwise.

### GetAnimationUrlOk

`func (o *ServicesMetadataDto) GetAnimationUrlOk() (*string, bool)`

GetAnimationUrlOk returns a tuple with the AnimationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationUrl

`func (o *ServicesMetadataDto) SetAnimationUrl(v string)`

SetAnimationUrl sets AnimationUrl field to given value.

### HasAnimationUrl

`func (o *ServicesMetadataDto) HasAnimationUrl() bool`

HasAnimationUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *ServicesMetadataDto) GetAttributes() []ModelsExposedMetadataAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServicesMetadataDto) GetAttributesOk() (*[]ModelsExposedMetadataAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServicesMetadataDto) SetAttributes(v []ModelsExposedMetadataAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServicesMetadataDto) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *ServicesMetadataDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesMetadataDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesMetadataDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalLink

`func (o *ServicesMetadataDto) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *ServicesMetadataDto) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *ServicesMetadataDto) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *ServicesMetadataDto) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetImage

`func (o *ServicesMetadataDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServicesMetadataDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServicesMetadataDto) SetImage(v string)`

SetImage sets Image field to given value.


### GetName

`func (o *ServicesMetadataDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesMetadataDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesMetadataDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


