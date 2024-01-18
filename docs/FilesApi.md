# \FilesAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFiles**](FilesAPI.md#ListFiles) | **Get** /v1/files/ | Obtain file list
[**UploadFile**](FilesAPI.md#UploadFile) | **Post** /v1/files/ | Upload file
[**UploadFileToOss**](FilesAPI.md#UploadFileToOss) | **Post** /v1/files/oss | Upload file to OSS
[**UploadFolder**](FilesAPI.md#UploadFolder) | **Post** /v1/files/folder | Upload folder
[**UploadFolderToOSS**](FilesAPI.md#UploadFolderToOSS) | **Post** /v1/files/folder/oss | Upload folder to oss



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer openapi_token
	page := "page_example" // string | page (optional)
	limit := "limit_example" // string | limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.ListFiles(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ListFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFiles`: ModelsFilesQueryResult
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ListFiles`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer openapi_token
	file := os.NewFile(1234, "some_file") // *os.File | uploaded file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFile(context.Background()).Authorization(authorization).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: ServicesUploadFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFile`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer openapi_token
	file := os.NewFile(1234, "some_file") // *os.File | uploaded file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFileToOss(context.Background()).Authorization(authorization).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFileToOss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileToOss`: ServicesUploadFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFileToOss`: %v\n", resp)
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


## UploadFolder

> ServicesUploadFolderResponse UploadFolder(ctx).Authorization(authorization).Folder(folder).Execute()

Upload folder



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
	folder := os.NewFile(1234, "some_file") // *os.File | uploaded folder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFolder(context.Background()).Authorization(authorization).Folder(folder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFolder`: ServicesUploadFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **folder** | ***os.File** | uploaded folder | 

### Return type

[**ServicesUploadFolderResponse**](ServicesUploadFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFolderToOSS

> ServicesUploadFolderResponse UploadFolderToOSS(ctx).Authorization(authorization).Folder(folder).Execute()

Upload folder to oss



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
	folder := os.NewFile(1234, "some_file") // *os.File | uploaded folder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFolderToOSS(context.Background()).Authorization(authorization).Folder(folder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFolderToOSS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFolderToOSS`: ServicesUploadFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFolderToOSS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFolderToOSSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 
 **folder** | ***os.File** | uploaded folder | 

### Return type

[**ServicesUploadFolderResponse**](ServicesUploadFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

