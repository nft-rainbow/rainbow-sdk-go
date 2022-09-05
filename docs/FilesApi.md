# \FilesApi

All URIs are relative to *http://api.nftrainbow.xyz/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFiles**](FilesApi.md#ListFiles) | **Get** /files/ | Obtain file list
[**UploadFile**](FilesApi.md#UploadFile) | **Post** /files/ | Upload file
[**UploadFileToOss**](FilesApi.md#UploadFileToOss) | **Post** /files/oss | Upload file to OSS



## ListFiles

> ModelsFilesQueryResult ListFiles(ctx).Authorization(authorization).Page(page).Limit(limit).Execute()

Obtain file list



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
    authorization := "authorization_example" // string | Bearer openapi_token
    page := "page_example" // string | page (optional)
    limit := "limit_example" // string | limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.ListFiles(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.ListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiles`: ModelsFilesQueryResult
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.ListFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **page** | **string** | page | 
 **limit** | **string** | limit | 

### Return type

[**ModelsFilesQueryResult**](ModelsFilesQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> ServicesUploadFilesResponse UploadFile(ctx).Authorization(authorization).File(file).Execute()

Upload file



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
    authorization := "authorization_example" // string | Bearer openapi_token
    file := os.NewFile(1234, "some_file") // *os.File | uploaded file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.UploadFile(context.Background()).Authorization(authorization).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: ServicesUploadFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **file** | ***os.File** | uploaded file | 

### Return type

[**ServicesUploadFilesResponse**](ServicesUploadFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToOss

> ServicesUploadFilesResponse UploadFileToOss(ctx).Authorization(authorization).File(file).Execute()

Upload file to OSS



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
    authorization := "authorization_example" // string | Bearer openapi_token
    file := os.NewFile(1234, "some_file") // *os.File | uploaded file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.UploadFileToOss(context.Background()).Authorization(authorization).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.UploadFileToOss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileToOss`: ServicesUploadFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.UploadFileToOss`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToOssRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **file** | ***os.File** | uploaded file | 

### Return type

[**ServicesUploadFilesResponse**](ServicesUploadFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

