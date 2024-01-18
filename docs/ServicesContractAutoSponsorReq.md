# ServicesContractAutoSponsorReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSponsor** | Pointer to **bool** | Address                  string &#x60;gorm:\&quot;type:varchar(256);index\&quot; json:\&quot;address\&quot;&#x60; | [optional] 
**StorageRechargeAmount** | Pointer to **int32** |  | [optional] 
**StorageRechargeThreshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesContractAutoSponsorReq

`func NewServicesContractAutoSponsorReq() *ServicesContractAutoSponsorReq`

NewServicesContractAutoSponsorReq instantiates a new ServicesContractAutoSponsorReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesContractAutoSponsorReqWithDefaults

`func NewServicesContractAutoSponsorReqWithDefaults() *ServicesContractAutoSponsorReq`

NewServicesContractAutoSponsorReqWithDefaults instantiates a new ServicesContractAutoSponsorReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSponsor

`func (o *ServicesContractAutoSponsorReq) GetAutoSponsor() bool`

GetAutoSponsor returns the AutoSponsor field if non-nil, zero value otherwise.

### GetAutoSponsorOk

`func (o *ServicesContractAutoSponsorReq) GetAutoSponsorOk() (*bool, bool)`

GetAutoSponsorOk returns a tuple with the AutoSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSponsor

`func (o *ServicesContractAutoSponsorReq) SetAutoSponsor(v bool)`

SetAutoSponsor sets AutoSponsor field to given value.

### HasAutoSponsor

`func (o *ServicesContractAutoSponsorReq) HasAutoSponsor() bool`

HasAutoSponsor returns a boolean if a field has been set.

### GetStorageRechargeAmount

`func (o *ServicesContractAutoSponsorReq) GetStorageRechargeAmount() int32`

GetStorageRechargeAmount returns the StorageRechargeAmount field if non-nil, zero value otherwise.

### GetStorageRechargeAmountOk

`func (o *ServicesContractAutoSponsorReq) GetStorageRechargeAmountOk() (*int32, bool)`

GetStorageRechargeAmountOk returns a tuple with the StorageRechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRechargeAmount

`func (o *ServicesContractAutoSponsorReq) SetStorageRechargeAmount(v int32)`

SetStorageRechargeAmount sets StorageRechargeAmount field to given value.

### HasStorageRechargeAmount

`func (o *ServicesContractAutoSponsorReq) HasStorageRechargeAmount() bool`

HasStorageRechargeAmount returns a boolean if a field has been set.

### GetStorageRechargeThreshold

`func (o *ServicesContractAutoSponsorReq) GetStorageRechargeThreshold() int32`

GetStorageRechargeThreshold returns the StorageRechargeThreshold field if non-nil, zero value otherwise.

### GetStorageRechargeThresholdOk

`func (o *ServicesContractAutoSponsorReq) GetStorageRechargeThresholdOk() (*int32, bool)`

GetStorageRechargeThresholdOk returns a tuple with the StorageRechargeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRechargeThreshold

`func (o *ServicesContractAutoSponsorReq) SetStorageRechargeThreshold(v int32)`

SetStorageRechargeThreshold sets StorageRechargeThreshold field to given value.

### HasStorageRechargeThreshold

`func (o *ServicesContractAutoSponsorReq) HasStorageRechargeThreshold() bool`

HasStorageRechargeThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


