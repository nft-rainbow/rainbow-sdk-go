# \MintsApi

All URIs are relative to *http://api.nftrainbow.xyz/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCustomMint**](MintsApi.md#BatchCustomMint) | **Post** /mints/customizable/batch | Batch Mint NFTs
[**CustomMint**](MintsApi.md#CustomMint) | **Post** /mints/ | Mint NFT
[**EasyMintByFile**](MintsApi.md#EasyMintByFile) | **Post** /mints/easy/files | Mint NFT with file
[**EasyMintByMetadata**](MintsApi.md#EasyMintByMetadata) | **Post** /mints/easy/urls | Mint NFT with metadata
[**GetMintDetail**](MintsApi.md#GetMintDetail) | **Get** /mints/{id} | Mint NFT detail
[**ListMints**](MintsApi.md#ListMints) | **Get** /mints/ | Obtain NFT list



## BatchCustomMint

> []ModelsMintTask BatchCustomMint(ctx).Authorization(authorization).CustomMintBatchDto(customMintBatchDto).Execute()

Batch Mint NFTs



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
    customMintBatchDto := *openapiclient.NewServicesCustomMintBatchDto("Chain_example", "ContractAddress_example", "ContractType_example", []openapiclient.ServicesMintItemDto{*openapiclient.NewServicesMintItemDto("MetadataUri_example", "MintToAddress_example")}) // ServicesCustomMintBatchDto | custom_mint_batch_dto

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.BatchCustomMint(context.Background()).Authorization(authorization).CustomMintBatchDto(customMintBatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.BatchCustomMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCustomMint`: []ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.BatchCustomMint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCustomMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **customMintBatchDto** | [**ServicesCustomMintBatchDto**](ServicesCustomMintBatchDto.md) | custom_mint_batch_dto | 

### Return type

[**[]ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomMint

> ModelsMintTask CustomMint(ctx).Authorization(authorization).CustomMintDto(customMintDto).Execute()

Mint NFT



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
    customMintDto := *openapiclient.NewServicesCustomMintDto("Chain_example", "ContractAddress_example", "ContractType_example", "MetadataUri_example", "MintToAddress_example") // ServicesCustomMintDto | custom_mint_dto

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.CustomMint(context.Background()).Authorization(authorization).CustomMintDto(customMintDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.CustomMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomMint`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.CustomMint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **customMintDto** | [**ServicesCustomMintDto**](ServicesCustomMintDto.md) | custom_mint_dto | 

### Return type

[**ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EasyMintByFile

> ModelsMintTask EasyMintByFile(ctx).Authorization(authorization).File(file).Chain(chain).Description(description).MintToAddress(mintToAddress).Name(name).Execute()

Mint NFT with file



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
    file := os.NewFile(1234, "some_file") // *os.File | file
    chain := "chain_example" // string | 
    description := "description_example" // string | 
    mintToAddress := "mintToAddress_example" // string | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.EasyMintByFile(context.Background()).Authorization(authorization).File(file).Chain(chain).Description(description).MintToAddress(mintToAddress).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.EasyMintByFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EasyMintByFile`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.EasyMintByFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEasyMintByFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **file** | ***os.File** | file | 
 **chain** | **string** |  | 
 **description** | **string** |  | 
 **mintToAddress** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EasyMintByMetadata

> ModelsMintTask EasyMintByMetadata(ctx).Authorization(authorization).EasyMintMetaInfo(easyMintMetaInfo).Execute()

Mint NFT with metadata



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
    easyMintMetaInfo := *openapiclient.NewServicesEasyMintMetaDto("Chain_example", "Description_example", "FileUrl_example", "MintToAddress_example", "Name_example") // ServicesEasyMintMetaDto | easy_mint_meta_info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.EasyMintByMetadata(context.Background()).Authorization(authorization).EasyMintMetaInfo(easyMintMetaInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.EasyMintByMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EasyMintByMetadata`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.EasyMintByMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEasyMintByMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **easyMintMetaInfo** | [**ServicesEasyMintMetaDto**](ServicesEasyMintMetaDto.md) | easy_mint_meta_info | 

### Return type

[**ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMintDetail

> ModelsMintTask GetMintDetail(ctx, id).Authorization(authorization).Execute()

Mint NFT detail



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
    resp, r, err := apiClient.MintsApi.GetMintDetail(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.GetMintDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMintDetail`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.GetMintDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ModelsMintTask**](ModelsMintTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMints

> ModelsMintTaskQueryResult ListMints(ctx).Authorization(authorization).Page(page).Limit(limit).Execute()

Obtain NFT list



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
    resp, r, err := apiClient.MintsApi.ListMints(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.ListMints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMints`: ModelsMintTaskQueryResult
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.ListMints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **page** | **int32** | page | 
 **limit** | **int32** | limit | 

### Return type

[**ModelsMintTaskQueryResult**](ModelsMintTaskQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

