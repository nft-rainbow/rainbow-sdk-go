# \MetadataApi

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetadata**](MetadataApi.md#CreateMetadata) | **Post** /v1/metadata/ | Create NFT metadata
[**GetMetadatInfo**](MetadataApi.md#GetMetadatInfo) | **Get** /v1/metadata/{metadata_id} | Query metadata
[**ListMetadatas**](MetadataApi.md#ListMetadatas) | **Get** /v1/metadata/ | Obtain metadata list



## CreateMetadata

> ModelsExposedMetadata CreateMetadata(ctx).Authorization(authorization).MetadataInfo(metadataInfo).Execute()

Create NFT metadata



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
    authorization := "authorization_example" // string | Bearer openapi_token
    metadataInfo := *openapiclient.NewServicesMetadataDto("Description_example", "Image_example", "Name_example") // ServicesMetadataDto | metadata_info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.CreateMetadata(context.Background()).Authorization(authorization).MetadataInfo(metadataInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.CreateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetadata`: ModelsExposedMetadata
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.CreateMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **metadataInfo** | [**ServicesMetadataDto**](ServicesMetadataDto.md) | metadata_info | 

### Return type

[**ModelsExposedMetadata**](ModelsExposedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadatInfo

> ModelsExposedMetadata GetMetadatInfo(ctx, metadataId).Authorization(authorization).Execute()

Query metadata



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
    authorization := "authorization_example" // string | Bearer openapi_token
    metadataId := "metadataId_example" // string | metadata_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetMetadatInfo(context.Background(), metadataId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetMetadatInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadatInfo`: ModelsExposedMetadata
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetMetadatInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataId** | **string** | metadata_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadatInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 


### Return type

[**ModelsExposedMetadata**](ModelsExposedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetadatas

> ModelsExposedMetadataQueryResult ListMetadatas(ctx).Authorization(authorization).Page(page).Limit(limit).Execute()

Obtain metadata list



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
    authorization := "authorization_example" // string | Bearer openapi_token
    page := "page_example" // string | page (optional)
    limit := "limit_example" // string | limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.ListMetadatas(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.ListMetadatas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetadatas`: ModelsExposedMetadataQueryResult
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.ListMetadatas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetadatasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **page** | **string** | page | 
 **limit** | **string** | limit | 

### Return type

[**ModelsExposedMetadataQueryResult**](ModelsExposedMetadataQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

