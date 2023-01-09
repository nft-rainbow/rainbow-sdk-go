# \DefaultApi

All URIs are relative to *http://api.nftrainbow.cn/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MintsEasyFilesPost**](DefaultApi.md#MintsEasyFilesPost) | **Post** /mints/easy/files | 



## MintsEasyFilesPost

> ModelsMintTask MintsEasyFilesPost(ctx).EasyMintFileInfo(easyMintFileInfo).Execute()



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
    easyMintFileInfo := *openapiclient.NewServicesEasyMintFileDto("Chain_example", "Description_example", *openapiclient.NewMultipartFileHeader(), "MintToAddress_example", "Name_example") // ServicesEasyMintFileDto | easy_mint_file_info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MintsEasyFilesPost(context.Background()).EasyMintFileInfo(easyMintFileInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MintsEasyFilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintsEasyFilesPost`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MintsEasyFilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintsEasyFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easyMintFileInfo** | [**ServicesEasyMintFileDto**](ServicesEasyMintFileDto.md) | easy_mint_file_info | 

### Return type

[**ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

