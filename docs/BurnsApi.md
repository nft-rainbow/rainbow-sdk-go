# \BurnsAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BurnBatch**](BurnsAPI.md#BurnBatch) | **Post** /v1/burns/customizable/batch | Batch burn NFT
[**BurnNft**](BurnsAPI.md#BurnNft) | **Post** /v1/burns | Burn NFT
[**GetBurnDetail**](BurnsAPI.md#GetBurnDetail) | **Get** /v1/burns/{id} | Burn NFT detail
[**GetBurnList**](BurnsAPI.md#GetBurnList) | **Get** /v1/burns | Obtain the burned NFTs list



## BurnBatch

> []ModelsBurnTask BurnBatch(ctx).Authorization(authorization).BurnBatchDto(burnBatchDto).Execute()

Batch burn NFT



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
	burnBatchDto := *openapiclient.NewServicesBurnBatchDto("Chain_example", "ContractAddress_example", "ContractType_example", []openapiclient.ServicesBurnItemDto{*openapiclient.NewServicesBurnItemDto("TokenId_example")}) // ServicesBurnBatchDto | burn_batch_dto

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnsAPI.BurnBatch(context.Background()).Authorization(authorization).BurnBatchDto(burnBatchDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnsAPI.BurnBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BurnBatch`: []ModelsBurnTask
	fmt.Fprintf(os.Stdout, "Response from `BurnsAPI.BurnBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBurnBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **burnBatchDto** | [**ServicesBurnBatchDto**](ServicesBurnBatchDto.md) | burn_batch_dto | 

### Return type

[**[]ModelsBurnTask**](ModelsBurnTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BurnNft

> ModelsBurnTask BurnNft(ctx).Authorization(authorization).BurnDto(burnDto).Execute()

Burn NFT



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
	burnDto := *openapiclient.NewServicesBurnDto("Chain_example", "ContractAddress_example", "ContractType_example", "TokenId_example") // ServicesBurnDto | burn_dto

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnsAPI.BurnNft(context.Background()).Authorization(authorization).BurnDto(burnDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnsAPI.BurnNft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BurnNft`: ModelsBurnTask
	fmt.Fprintf(os.Stdout, "Response from `BurnsAPI.BurnNft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBurnNftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **burnDto** | [**ServicesBurnDto**](ServicesBurnDto.md) | burn_dto | 

### Return type

[**ModelsBurnTask**](ModelsBurnTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnDetail

> ModelsBurnTask GetBurnDetail(ctx, id).Authorization(authorization).Execute()

Burn NFT detail



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
	id := int32(56) // int32 | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnsAPI.GetBurnDetail(context.Background(), id).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnsAPI.GetBurnDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnDetail`: ModelsBurnTask
	fmt.Fprintf(os.Stdout, "Response from `BurnsAPI.GetBurnDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ModelsBurnTask**](ModelsBurnTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnList

> ModelsBurnTaskQueryResult GetBurnList(ctx).Authorization(authorization).Execute()

Obtain the burned NFTs list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnsAPI.GetBurnList(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnsAPI.GetBurnList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnList`: ModelsBurnTaskQueryResult
	fmt.Fprintf(os.Stdout, "Response from `BurnsAPI.GetBurnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

### Return type

[**ModelsBurnTaskQueryResult**](ModelsBurnTaskQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

