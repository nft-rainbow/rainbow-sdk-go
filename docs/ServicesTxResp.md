# ServicesTxResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**ErrorMsg** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**IsSuccess** | Pointer to **bool** |  | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**StateCode** | Pointer to [**ModelsTxState**](ModelsTxState.md) |  | [optional] 
**StateMsg** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

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

### GetData

`func (o *ServicesTxResp) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesTxResp) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesTxResp) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ServicesTxResp) HasData() bool`

HasData returns a boolean if a field has been set.

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

### GetFrom

`func (o *ServicesTxResp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServicesTxResp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServicesTxResp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ServicesTxResp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

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

### GetNonce

`func (o *ServicesTxResp) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ServicesTxResp) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ServicesTxResp) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ServicesTxResp) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

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

### GetTo

`func (o *ServicesTxResp) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServicesTxResp) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServicesTxResp) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ServicesTxResp) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *ServicesTxResp) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServicesTxResp) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServicesTxResp) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServicesTxResp) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


