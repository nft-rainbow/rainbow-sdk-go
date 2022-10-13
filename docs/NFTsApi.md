# \NFTsApi

All URIs are relative to *http://api.nftrainbow.cn/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NFTInfo**](NFTsApi.md#NFTInfo) | **Get** /nft/{address}/{token_id} | Get NFT info, mainly owner and metadata



## NFTInfo

> ServicesNFT NFTInfo(ctx, address, tokenId).Authorization(authorization).Execute()

Get NFT info, mainly owner and metadata



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
    address := "address_example" // string | address
    tokenId := "tokenId_example" // string | token_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTsApi.NFTInfo(context.Background(), address, tokenId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTsApi.NFTInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NFTInfo`: ServicesNFT
    fmt.Fprintf(os.Stdout, "Response from `NFTsApi.NFTInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 
**tokenId** | **string** | token_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNFTInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 



### Return type

[**ServicesNFT**](ServicesNFT.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

