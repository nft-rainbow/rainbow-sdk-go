# ModelsTBACollectionItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **[]string** | Addresses of the specified tbas. Will throw an error if the tba is not created yet | 

## Methods

### NewModelsTBACollectionItemsDto

`func NewModelsTBACollectionItemsDto(accounts []string, ) *ModelsTBACollectionItemsDto`

NewModelsTBACollectionItemsDto instantiates a new ModelsTBACollectionItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTBACollectionItemsDtoWithDefaults

`func NewModelsTBACollectionItemsDtoWithDefaults() *ModelsTBACollectionItemsDto`

NewModelsTBACollectionItemsDtoWithDefaults instantiates a new ModelsTBACollectionItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ModelsTBACollectionItemsDto) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ModelsTBACollectionItemsDto) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ModelsTBACollectionItemsDto) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


