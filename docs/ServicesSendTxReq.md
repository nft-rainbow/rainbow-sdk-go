# ServicesSendTxReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**Data** | Pointer to **string** |  | [optional] 
**From** | **string** | AccountId uint   &#x60;json:\&quot;account_id\&quot; form:\&quot;account_id\&quot; binding:\&quot;required\&quot;&#x60; | 
**To** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesSendTxReq

`func NewServicesSendTxReq(chain string, from string, to string, ) *ServicesSendTxReq`

NewServicesSendTxReq instantiates a new ServicesSendTxReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesSendTxReqWithDefaults

`func NewServicesSendTxReqWithDefaults() *ServicesSendTxReq`

NewServicesSendTxReqWithDefaults instantiates a new ServicesSendTxReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ServicesSendTxReq) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ServicesSendTxReq) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ServicesSendTxReq) SetChain(v string)`

SetChain sets Chain field to given value.


### GetData

`func (o *ServicesSendTxReq) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesSendTxReq) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesSendTxReq) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ServicesSendTxReq) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFrom

`func (o *ServicesSendTxReq) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServicesSendTxReq) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServicesSendTxReq) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ServicesSendTxReq) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServicesSendTxReq) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServicesSendTxReq) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *ServicesSendTxReq) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServicesSendTxReq) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServicesSendTxReq) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServicesSendTxReq) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


