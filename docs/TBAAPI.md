# \TBAAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TBACollectionEndpoint**](TBAAPI.md#TBACollectionEndpoint) | **Post** /v1/tba/collections | Create a new TBA collection
[**TBACollectionMetaInfoModificationEndpoint**](TBAAPI.md#TBACollectionMetaInfoModificationEndpoint) | **Put** /v1/tba/collections/{collection_name} | Modify meta information of a TBA collection
[**TBACollectionNewItemsEndpoint**](TBAAPI.md#TBACollectionNewItemsEndpoint) | **Post** /v1/tba/collections/{collection_name}/items | Add new items to a TBA collection
[**TBACollectionQueryEndpoint**](TBAAPI.md#TBACollectionQueryEndpoint) | **Get** /v1/tba/collections | Query token bound account (TBA) collections
[**TBACollectionRemoveItemsEndpoint**](TBAAPI.md#TBACollectionRemoveItemsEndpoint) | **Delete** /v1/tba/collections/{collection_name}/items | Remove items from a TBA collection
[**TBACreationEndpoint**](TBAAPI.md#TBACreationEndpoint) | **Post** /v1/tba/accounts | Create a new token bound account (TBA)
[**TBAQueryEndpoint**](TBAAPI.md#TBAQueryEndpoint) | **Get** /v1/tba/accounts | Get all token bound accounts on chain



## TBACollectionEndpoint

> ModelsTokenBoundAccountCollectionQueryResult TBACollectionEndpoint(ctx).Authorization(authorization).CreateDto(createDto).Execute()

Create a new TBA collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	createDto := *openapiclient.NewModelsTBACollectionCreateDto("Chain_example", "Name_example") // ModelsTBACollectionCreateDto | TBA Collection creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACollectionEndpoint(context.Background()).Authorization(authorization).CreateDto(createDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACollectionEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACollectionEndpoint`: ModelsTokenBoundAccountCollectionQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACollectionEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTBACollectionEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **createDto** | [**ModelsTBACollectionCreateDto**](ModelsTBACollectionCreateDto.md) | TBA Collection creation details | 

### Return type

[**ModelsTokenBoundAccountCollectionQueryResult**](ModelsTokenBoundAccountCollectionQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBACollectionMetaInfoModificationEndpoint

> map[string]map[string]interface{} TBACollectionMetaInfoModificationEndpoint(ctx, collectionName).Authorization(authorization).ModifyDto(modifyDto).Execute()

Modify meta information of a TBA collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	collectionName := "collectionName_example" // string | Name of the TBA collection to be modified
	modifyDto := *openapiclient.NewModelsTBACollectionModifyDto() // ModelsTBACollectionModifyDto | Modification details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACollectionMetaInfoModificationEndpoint(context.Background(), collectionName).Authorization(authorization).ModifyDto(modifyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACollectionMetaInfoModificationEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACollectionMetaInfoModificationEndpoint`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACollectionMetaInfoModificationEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** | Name of the TBA collection to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiTBACollectionMetaInfoModificationEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **modifyDto** | [**ModelsTBACollectionModifyDto**](ModelsTBACollectionModifyDto.md) | Modification details | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBACollectionNewItemsEndpoint

> map[string]map[string]interface{} TBACollectionNewItemsEndpoint(ctx, collectionName).Authorization(authorization).ItemsDto(itemsDto).Execute()

Add new items to a TBA collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	collectionName := "collectionName_example" // string | Name of the TBA collection to add items to
	itemsDto := *openapiclient.NewModelsTBACollectionItemsDto([]string{"Accounts_example"}) // ModelsTBACollectionItemsDto | Details of the accounts to be added

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACollectionNewItemsEndpoint(context.Background(), collectionName).Authorization(authorization).ItemsDto(itemsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACollectionNewItemsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACollectionNewItemsEndpoint`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACollectionNewItemsEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** | Name of the TBA collection to add items to | 

### Other Parameters

Other parameters are passed through a pointer to a apiTBACollectionNewItemsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **itemsDto** | [**ModelsTBACollectionItemsDto**](ModelsTBACollectionItemsDto.md) | Details of the accounts to be added | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBACollectionQueryEndpoint

> ModelsTokenBoundAccountCollectionQueryResult TBACollectionQueryEndpoint(ctx).Authorization(authorization).Chain(chain).Name(name).Page(page).Limit(limit).Execute()

Query token bound account (TBA) collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	chain := "chain_example" // string | The chain type. The types include `conflux` and `conflux_test`
	name := "name_example" // string | Collection name (optional)
	page := int32(56) // int32 | Page number (optional)
	limit := int32(56) // int32 | Limit per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACollectionQueryEndpoint(context.Background()).Authorization(authorization).Chain(chain).Name(name).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACollectionQueryEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACollectionQueryEndpoint`: ModelsTokenBoundAccountCollectionQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACollectionQueryEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTBACollectionQueryEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **chain** | **string** | The chain type. The types include &#x60;conflux&#x60; and &#x60;conflux_test&#x60; | 
 **name** | **string** | Collection name | 
 **page** | **int32** | Page number | 
 **limit** | **int32** | Limit per page | 

### Return type

[**ModelsTokenBoundAccountCollectionQueryResult**](ModelsTokenBoundAccountCollectionQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBACollectionRemoveItemsEndpoint

> map[string]map[string]interface{} TBACollectionRemoveItemsEndpoint(ctx, collectionName).Authorization(authorization).ItemsDto(itemsDto).Execute()

Remove items from a TBA collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	collectionName := "collectionName_example" // string | Name of the TBA collection to remove items from
	itemsDto := *openapiclient.NewModelsTBACollectionItemsDto([]string{"Accounts_example"}) // ModelsTBACollectionItemsDto | Details of the accounts to be removed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACollectionRemoveItemsEndpoint(context.Background(), collectionName).Authorization(authorization).ItemsDto(itemsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACollectionRemoveItemsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACollectionRemoveItemsEndpoint`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACollectionRemoveItemsEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** | Name of the TBA collection to remove items from | 

### Other Parameters

Other parameters are passed through a pointer to a apiTBACollectionRemoveItemsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **itemsDto** | [**ModelsTBACollectionItemsDto**](ModelsTBACollectionItemsDto.md) | Details of the accounts to be removed | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBACreationEndpoint

> ModelsTBACreationTask TBACreationEndpoint(ctx).Authorization(authorization).TbaCreateDto(tbaCreateDto).Execute()

Create a new token bound account (TBA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	tbaCreateDto := *openapiclient.NewModelsTBACreateDto("Chain_example", "TokenContract_example", "TokenId_example") // ModelsTBACreateDto | Token Bound Account creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBACreationEndpoint(context.Background()).Authorization(authorization).TbaCreateDto(tbaCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBACreationEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBACreationEndpoint`: ModelsTBACreationTask
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBACreationEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTBACreationEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **tbaCreateDto** | [**ModelsTBACreateDto**](ModelsTBACreateDto.md) | Token Bound Account creation details | 

### Return type

[**ModelsTBACreationTask**](ModelsTBACreationTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TBAQueryEndpoint

> ModelsTokenBoundAccountQueryResult TBAQueryEndpoint(ctx).Authorization(authorization).Chain(chain).Implementation(implementation).Contract(contract).TokenId(tokenId).Page(page).Limit(limit).Execute()

Get all token bound accounts on chain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	chain := "chain_example" // string | Chain name, should be either conflux or conflux_test
	implementation := "implementation_example" // string | The implementation contract address of the tba. Will throw an error if the implementation is not in whitelist (optional)
	contract := "contract_example" // string | Token contract address of the bounded NFT (optional)
	tokenId := "tokenId_example" // string | Token ID (optional)
	page := int32(56) // int32 | Page number (optional)
	limit := int32(56) // int32 | Limit per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TBAAPI.TBAQueryEndpoint(context.Background()).Authorization(authorization).Chain(chain).Implementation(implementation).Contract(contract).TokenId(tokenId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TBAAPI.TBAQueryEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TBAQueryEndpoint`: ModelsTokenBoundAccountQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TBAAPI.TBAQueryEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTBAQueryEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **chain** | **string** | Chain name, should be either conflux or conflux_test | 
 **implementation** | **string** | The implementation contract address of the tba. Will throw an error if the implementation is not in whitelist | 
 **contract** | **string** | Token contract address of the bounded NFT | 
 **tokenId** | **string** | Token ID | 
 **page** | **int32** | Page number | 
 **limit** | **int32** | Limit per page | 

### Return type

[**ModelsTokenBoundAccountQueryResult**](ModelsTokenBoundAccountQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

