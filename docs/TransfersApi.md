# \TransfersApi

All URIs are relative to *http://api.nftrainbow.xyz/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchTransferNft**](TransfersApi.md#BatchTransferNft) | **Post** /transfers/customizable/batch | Batch Transfer NFTs
[**GetTransferDetail**](TransfersApi.md#GetTransferDetail) | **Get** /transfers/{id} | Transfer NFT detail
[**ListTransfer**](TransfersApi.md#ListTransfer) | **Get** /transfers/ | Obtain the transferred NFTs list
[**TransferNft**](TransfersApi.md#TransferNft) | **Post** /transfers/customizable | Transfer NFT



## BatchTransferNft

> []ModelsTransferTask BatchTransferNft(ctx).Authorization(authorization).TransferBatchDto(transferBatchDto).Execute()

Batch Transfer NFTs



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
    transferBatchDto := *openapiclient.NewServicesTransferBatchDto("Chain_example", "ContractAddress_example", "ContractType_example", []openapiclient.ServicesTransferItemDto{*openapiclient.NewServicesTransferItemDto("TokenId_example", "TransferFromAddress_example", "TransferToAddress_example")}) // ServicesTransferBatchDto | transfer_batch_dto

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.BatchTransferNft(context.Background()).Authorization(authorization).TransferBatchDto(transferBatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.BatchTransferNft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchTransferNft`: []ModelsTransferTask
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.BatchTransferNft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchTransferNftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **transferBatchDto** | [**ServicesTransferBatchDto**](ServicesTransferBatchDto.md) | transfer_batch_dto | 

### Return type

[**[]ModelsTransferTask**](ModelsTransferTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferDetail

> ModelsTransferTask GetTransferDetail(ctx, id).Authorization(authorization).Execute()

Transfer NFT detail



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
    resp, r, err := apiClient.TransfersApi.GetTransferDetail(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetTransferDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransferDetail`: ModelsTransferTask
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetTransferDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ModelsTransferTask**](ModelsTransferTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransfer

> ModelsTransferTaskQueryResult ListTransfer(ctx).Authorization(authorization).Page(page).Limit(limit).Execute()

Obtain the transferred NFTs list



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
    page := int32(56) // int32 | page (optional)
    limit := int32(56) // int32 | limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.ListTransfer(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.ListTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransfer`: ModelsTransferTaskQueryResult
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.ListTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **page** | **int32** | page | 
 **limit** | **int32** | limit | 

### Return type

[**ModelsTransferTaskQueryResult**](ModelsTransferTaskQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferNft

> ModelsTransferTask TransferNft(ctx).Authorization(authorization).TransferDto(transferDto).Execute()

Transfer NFT



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
    transferDto := *openapiclient.NewServicesTransferDto("Chain_example", "ContractAddress_example", "ContractType_example", "TokenId_example", "TransferFromAddress_example", "TransferToAddress_example") // ServicesTransferDto | transfer_dto

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.TransferNft(context.Background()).Authorization(authorization).TransferDto(transferDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.TransferNft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferNft`: ModelsTransferTask
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.TransferNft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferNftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **transferDto** | [**ServicesTransferDto**](ServicesTransferDto.md) | transfer_dto | 

### Return type

[**ModelsTransferTask**](ModelsTransferTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

