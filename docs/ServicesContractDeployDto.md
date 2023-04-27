# ServicesContractDeployDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSponsor** | Pointer to **bool** | default: true | [optional] 
**BaseUri** | Pointer to **string** |  | [optional] 
**Chain** | **string** |  | 
**IsSponsorForAllUser** | Pointer to **bool** | default: true | [optional] 
**Name** | **string** |  | 
**OwnerAddress** | Pointer to **string** |  | [optional] 
**RoyaltiesAddress** | Pointer to **string** |  | [optional] 
**RoyaltiesBps** | Pointer to **int32** |  | [optional] 
**Symbol** | **string** |  | 
**TokensTransferableByAdmin** | Pointer to **bool** | default: true | [optional] 
**TokensTransferableByUser** | Pointer to **bool** | default: true | [optional] 
**TransferCooldownTime** | Pointer to **int32** | default: 0 | [optional] 
**Type** | **string** |  | 

## Methods

### NewServicesContractDeployDto

`func NewServicesContractDeployDto(chain string, name string, symbol string, type_ string, ) *ServicesContractDeployDto`

NewServicesContractDeployDto instantiates a new ServicesContractDeployDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesContractDeployDtoWithDefaults

`func NewServicesContractDeployDtoWithDefaults() *ServicesContractDeployDto`

NewServicesContractDeployDtoWithDefaults instantiates a new ServicesContractDeployDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSponsor

`func (o *ServicesContractDeployDto) GetAutoSponsor() bool`

GetAutoSponsor returns the AutoSponsor field if non-nil, zero value otherwise.

### GetAutoSponsorOk

`func (o *ServicesContractDeployDto) GetAutoSponsorOk() (*bool, bool)`

GetAutoSponsorOk returns a tuple with the AutoSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSponsor

`func (o *ServicesContractDeployDto) SetAutoSponsor(v bool)`

SetAutoSponsor sets AutoSponsor field to given value.

### HasAutoSponsor

`func (o *ServicesContractDeployDto) HasAutoSponsor() bool`

HasAutoSponsor returns a boolean if a field has been set.

### GetBaseUri

`func (o *ServicesContractDeployDto) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *ServicesContractDeployDto) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *ServicesContractDeployDto) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *ServicesContractDeployDto) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.

### GetChain

`func (o *ServicesContractDeployDto) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesContractDeployDto) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesContractDeployDto) SetChain(v string)`

SetChain sets Chain field to given value.


### GetIsSponsorForAllUser

`func (o *ServicesContractDeployDto) GetIsSponsorForAllUser() bool`

GetIsSponsorForAllUser returns the IsSponsorForAllUser field if non-nil, zero value otherwise.

### GetIsSponsorForAllUserOk

`func (o *ServicesContractDeployDto) GetIsSponsorForAllUserOk() (*bool, bool)`

GetIsSponsorForAllUserOk returns a tuple with the IsSponsorForAllUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSponsorForAllUser

`func (o *ServicesContractDeployDto) SetIsSponsorForAllUser(v bool)`

SetIsSponsorForAllUser sets IsSponsorForAllUser field to given value.

### HasIsSponsorForAllUser

`func (o *ServicesContractDeployDto) HasIsSponsorForAllUser() bool`

HasIsSponsorForAllUser returns a boolean if a field has been set.

### GetName

`func (o *ServicesContractDeployDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesContractDeployDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesContractDeployDto) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerAddress

`func (o *ServicesContractDeployDto) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *ServicesContractDeployDto) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *ServicesContractDeployDto) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *ServicesContractDeployDto) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetRoyaltiesAddress

`func (o *ServicesContractDeployDto) GetRoyaltiesAddress() string`

GetRoyaltiesAddress returns the RoyaltiesAddress field if non-nil, zero value otherwise.

### GetRoyaltiesAddressOk

`func (o *ServicesContractDeployDto) GetRoyaltiesAddressOk() (*string, bool)`

GetRoyaltiesAddressOk returns a tuple with the RoyaltiesAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltiesAddress

`func (o *ServicesContractDeployDto) SetRoyaltiesAddress(v string)`

SetRoyaltiesAddress sets RoyaltiesAddress field to given value.

### HasRoyaltiesAddress

`func (o *ServicesContractDeployDto) HasRoyaltiesAddress() bool`

HasRoyaltiesAddress returns a boolean if a field has been set.

### GetRoyaltiesBps

`func (o *ServicesContractDeployDto) GetRoyaltiesBps() int32`

GetRoyaltiesBps returns the RoyaltiesBps field if non-nil, zero value otherwise.

### GetRoyaltiesBpsOk

`func (o *ServicesContractDeployDto) GetRoyaltiesBpsOk() (*int32, bool)`

GetRoyaltiesBpsOk returns a tuple with the RoyaltiesBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltiesBps

`func (o *ServicesContractDeployDto) SetRoyaltiesBps(v int32)`

SetRoyaltiesBps sets RoyaltiesBps field to given value.

### HasRoyaltiesBps

`func (o *ServicesContractDeployDto) HasRoyaltiesBps() bool`

HasRoyaltiesBps returns a boolean if a field has been set.

### GetSymbol

`func (o *ServicesContractDeployDto) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ServicesContractDeployDto) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ServicesContractDeployDto) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokensTransferableByAdmin

`func (o *ServicesContractDeployDto) GetTokensTransferableByAdmin() bool`

GetTokensTransferableByAdmin returns the TokensTransferableByAdmin field if non-nil, zero value otherwise.

### GetTokensTransferableByAdminOk

`func (o *ServicesContractDeployDto) GetTokensTransferableByAdminOk() (*bool, bool)`

GetTokensTransferableByAdminOk returns a tuple with the TokensTransferableByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensTransferableByAdmin

`func (o *ServicesContractDeployDto) SetTokensTransferableByAdmin(v bool)`

SetTokensTransferableByAdmin sets TokensTransferableByAdmin field to given value.

### HasTokensTransferableByAdmin

`func (o *ServicesContractDeployDto) HasTokensTransferableByAdmin() bool`

HasTokensTransferableByAdmin returns a boolean if a field has been set.

### GetTokensTransferableByUser

`func (o *ServicesContractDeployDto) GetTokensTransferableByUser() bool`

GetTokensTransferableByUser returns the TokensTransferableByUser field if non-nil, zero value otherwise.

### GetTokensTransferableByUserOk

`func (o *ServicesContractDeployDto) GetTokensTransferableByUserOk() (*bool, bool)`

GetTokensTransferableByUserOk returns a tuple with the TokensTransferableByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensTransferableByUser

`func (o *ServicesContractDeployDto) SetTokensTransferableByUser(v bool)`

SetTokensTransferableByUser sets TokensTransferableByUser field to given value.

### HasTokensTransferableByUser

`func (o *ServicesContractDeployDto) HasTokensTransferableByUser() bool`

HasTokensTransferableByUser returns a boolean if a field has been set.

### GetTransferCooldownTime

`func (o *ServicesContractDeployDto) GetTransferCooldownTime() int32`

GetTransferCooldownTime returns the TransferCooldownTime field if non-nil, zero value otherwise.

### GetTransferCooldownTimeOk

`func (o *ServicesContractDeployDto) GetTransferCooldownTimeOk() (*int32, bool)`

GetTransferCooldownTimeOk returns a tuple with the TransferCooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferCooldownTime

`func (o *ServicesContractDeployDto) SetTransferCooldownTime(v int32)`

SetTransferCooldownTime sets TransferCooldownTime field to given value.

### HasTransferCooldownTime

`func (o *ServicesContractDeployDto) HasTransferCooldownTime() bool`

HasTransferCooldownTime returns a boolean if a field has been set.

### GetType

`func (o *ServicesContractDeployDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicesContractDeployDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicesContractDeployDto) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


