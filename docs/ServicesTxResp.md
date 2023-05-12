# ServicesTxResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMsg** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**IsSuccess** | Pointer to **bool** |  | [optional] 
**StateCode** | Pointer to [**ModelsTxState**](ModelsTxState.md) |  | [optional] 
**StateMsg** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesTxResp

`func NewServicesTxResp() *ServicesTxResp`

NewServicesTxResp instantiates a new ServicesTxResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesTxRespWithDefaults

`func NewServicesTxRespWithDefaults() *ServicesTxResp`

NewServicesTxRespWithDefaults instantiates a new ServicesTxResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMsg

`func (o *ServicesTxResp) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *ServicesTxResp) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *ServicesTxResp) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *ServicesTxResp) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetHash

`func (o *ServicesTxResp) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ServicesTxResp) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ServicesTxResp) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ServicesTxResp) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsFinalized

`func (o *ServicesTxResp) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *ServicesTxResp) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *ServicesTxResp) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *ServicesTxResp) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetIsSuccess

`func (o *ServicesTxResp) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *ServicesTxResp) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *ServicesTxResp) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *ServicesTxResp) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetStateCode

`func (o *ServicesTxResp) GetStateCode() ModelsTxState`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *ServicesTxResp) GetStateCodeOk() (*ModelsTxState, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *ServicesTxResp) SetStateCode(v ModelsTxState)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *ServicesTxResp) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetStateMsg

`func (o *ServicesTxResp) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *ServicesTxResp) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *ServicesTxResp) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *ServicesTxResp) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


