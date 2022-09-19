# ServicesContractAdminUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminAddress** | **string** | Chain           string &#x60;form:\&quot;chain\&quot; json:\&quot;chain\&quot; binding:\&quot;required,oneof&#x3D;conflux conflux_test\&quot;&#x60; ContractAddress string &#x60;form:\&quot;contract_address\&quot; json:\&quot;contract_address\&quot; binding:\&quot;required\&quot;&#x60; | 

## Methods

### NewServicesContractAdminUpdateDto

`func NewServicesContractAdminUpdateDto(adminAddress string, ) *ServicesContractAdminUpdateDto`

NewServicesContractAdminUpdateDto instantiates a new ServicesContractAdminUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesContractAdminUpdateDtoWithDefaults

`func NewServicesContractAdminUpdateDtoWithDefaults() *ServicesContractAdminUpdateDto`

NewServicesContractAdminUpdateDtoWithDefaults instantiates a new ServicesContractAdminUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminAddress

`func (o *ServicesContractAdminUpdateDto) GetAdminAddress() string`

GetAdminAddress returns the AdminAddress field if non-nil, zero value otherwise.

### GetAdminAddressOk

`func (o *ServicesContractAdminUpdateDto) GetAdminAddressOk() (*string, bool)`

GetAdminAddressOk returns a tuple with the AdminAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAddress

`func (o *ServicesContractAdminUpdateDto) SetAdminAddress(v string)`

SetAdminAddress sets AdminAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


