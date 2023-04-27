# ModelsMintTaskOrError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**MintTask** | Pointer to [**ModelsMintTask**](ModelsMintTask.md) |  | [optional] 

## Methods

### NewModelsMintTaskOrError

`func NewModelsMintTaskOrError() *ModelsMintTaskOrError`

NewModelsMintTaskOrError instantiates a new ModelsMintTaskOrError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMintTaskOrErrorWithDefaults

`func NewModelsMintTaskOrErrorWithDefaults() *ModelsMintTaskOrError`

NewModelsMintTaskOrErrorWithDefaults instantiates a new ModelsMintTaskOrError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ModelsMintTaskOrError) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelsMintTaskOrError) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelsMintTaskOrError) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ModelsMintTaskOrError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMintTask

`func (o *ModelsMintTaskOrError) GetMintTask() ModelsMintTask`

GetMintTask returns the MintTask field if non-nil, zero value otherwise.

### GetMintTaskOk

`func (o *ModelsMintTaskOrError) GetMintTaskOk() (*ModelsMintTask, bool)`

GetMintTaskOk returns a tuple with the MintTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTask

`func (o *ModelsMintTaskOrError) SetMintTask(v ModelsMintTask)`

SetMintTask sets MintTask field to given value.

### HasMintTask

`func (o *ModelsMintTaskOrError) HasMintTask() bool`

HasMintTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


