# \BurnsApi

All URIs are relative to *http://api.nftrainbow.cn/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BurnBatch**](BurnsApi.md#BurnBatch) | **Post** /burns | Burn batch NFT
[**GetBurnDetail**](BurnsApi.md#GetBurnDetail) | **Get** /burns/{id} | Burn NFT detail
[**GetBurnList**](BurnsApi.md#GetBurnList) | **Get** /burns | Obtain the burned NFTs list



## BurnBatch

> []ModelsBurnTask BurnBatch(ctx).Authorization(authorization).BurnBatchDto(burnBatchDto).Execute()

Burn batch NFT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT
    burnBatchDto := *openapiclient.NewServicesBurnBatchDto("Chain_example", "ContractAddress_example", "ContractType_example", []openapiclient.ServicesBurnItemDto{*openapiclient.NewServicesBurnItemDto("TokenId_example")}) // ServicesBurnBatchDto | burn_batch_dto

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BurnsApi.BurnBatch(context.Background()).Authorization(authorization).BurnBatchDto(burnBatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BurnsApi.BurnBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BurnBatch`: []ModelsBurnTask
    fmt.Fprintf(os.Stdout, "Response from `BurnsApi.BurnBatch`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT
    id := int32(56) // int32 | id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BurnsApi.GetBurnDetail(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BurnsApi.GetBurnDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBurnDetail`: ModelsBurnTask
    fmt.Fprintf(os.Stdout, "Response from `BurnsApi.GetBurnDetail`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BurnsApi.GetBurnList(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BurnsApi.GetBurnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBurnList`: ModelsBurnTaskQueryResult
    fmt.Fprintf(os.Stdout, "Response from `BurnsApi.GetBurnList`: %v\n", resp)
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

