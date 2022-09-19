# \TransactionApi

All URIs are relative to *http://api.nftrainbow.xyz/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionByID**](TransactionApi.md#GetTransactionByID) | **Get** /tx/{id} | Get transaction informantion by ID



## GetTransactionByID

> ServicesTxResp GetTransactionByID(ctx, id).Authorization(authorization).Execute()

Get transaction informantion by ID



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
    id := "id_example" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionApi.GetTransactionByID(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.GetTransactionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionByID`: ServicesTxResp
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.GetTransactionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ServicesTxResp**](ServicesTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

