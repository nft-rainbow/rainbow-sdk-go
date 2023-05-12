# \NFTsApi

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NftInfo**](NFTsApi.md#NftInfo) | **Get** /v1/nft/{address}/{token_id} | Get NFT info, mainly owner and metadata
[**UpdateNftTokenUri**](NFTsApi.md#UpdateNftTokenUri) | **Put** /v1/nft/{address}/{token_id}/tokenUri | Update NFT token uri



## NftInfo

> ServicesNFT NftInfo(ctx, address, tokenId).Authorization(authorization).Execute()

Get NFT info, mainly owner and metadata



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
    address := "address_example" // string | address
    tokenId := "tokenId_example" // string | token_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTsApi.NftInfo(context.Background(), address, tokenId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTsApi.NftInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NftInfo`: ServicesNFT
    fmt.Fprintf(os.Stdout, "Response from `NFTsApi.NftInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 
**tokenId** | **string** | token_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNftInfoRequest struct via the builder pattern


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


## UpdateNftTokenUri

> ModelsTransaction UpdateNftTokenUri(ctx, address, tokenId).Authorization(authorization).Req(req).Execute()

Update NFT token uri



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
    address := "address_example" // string | address
    tokenId := "tokenId_example" // string | token_id
    req := *openapiclient.NewRoutersUpdateTokenUriReq("Chain_example", "ContractType_example") // RoutersUpdateTokenUriReq | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTsApi.UpdateNftTokenUri(context.Background(), address, tokenId).Authorization(authorization).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTsApi.UpdateNftTokenUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNftTokenUri`: ModelsTransaction
    fmt.Fprintf(os.Stdout, "Response from `NFTsApi.UpdateNftTokenUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 
**tokenId** | **string** | token_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNftTokenUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


 **req** | [**RoutersUpdateTokenUriReq**](RoutersUpdateTokenUriReq.md) | req | 

### Return type

[**ModelsTransaction**](ModelsTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

