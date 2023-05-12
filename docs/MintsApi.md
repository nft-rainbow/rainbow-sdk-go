# \MintsApi

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppBatchMintByMetaUri**](MintsApi.md#AppBatchMintByMetaUri) | **Post** /dashboard/apps/{id}/nft/batch/by-meta-uri | Batch Mint NFT with metadata uri
[**AppBatchMintNFT**](MintsApi.md#AppBatchMintNFT) | **Post** /dashboard/apps/{id}/nft/batch/by-meta-parts | Batch Mint NFT with metadata parts
[**BatchCustomMint**](MintsApi.md#BatchCustomMint) | **Post** /v1/mints/customizable/batch | Batch Mint NFTs
[**CustomMint**](MintsApi.md#CustomMint) | **Post** /v1/mints/ | Mint NFT
[**EasyMintByFile**](MintsApi.md#EasyMintByFile) | **Post** /v1/mints/easy/files | Mint NFT with file
[**EasyMintByMetadata**](MintsApi.md#EasyMintByMetadata) | **Post** /v1/mints/easy/urls | Mint NFT with metadata
[**GetMintDetail**](MintsApi.md#GetMintDetail) | **Get** /v1/mints/{id} | Mint NFT detail
[**ListMints**](MintsApi.md#ListMints) | **Get** /v1/mints/ | Obtain NFT list
[**ListMints_0**](MintsApi.md#ListMints_0) | **Get** /v1/mints/ | Obtain NFT list
[**ReMintNFT**](MintsApi.md#ReMintNFT) | **Post** /v1/mints/{id}/reMint | Reset mint task status to init so that it can be minted again



## AppBatchMintByMetaUri

> []int32 AppBatchMintByMetaUri(ctx, id).Authorization(authorization).BatchMintRequest(batchMintRequest).Execute()

Batch Mint NFT with metadata uri



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
    authorization := "authorization_example" // string | Bearer JWT
    id := int32(56) // int32 | id
    batchMintRequest := *openapiclient.NewServicesAppBatchMintByMetaUriDto("Chain_example", "ContractAddress_example", []openapiclient.ServicesMintItemDto{*openapiclient.NewServicesMintItemDto("MintToAddress_example")}) // ServicesAppBatchMintByMetaUriDto | batch_mint_request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.AppBatchMintByMetaUri(context.Background(), id).Authorization(authorization).BatchMintRequest(batchMintRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.AppBatchMintByMetaUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppBatchMintByMetaUri`: []int32
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.AppBatchMintByMetaUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBatchMintByMetaUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer JWT | 

 **batchMintRequest** | [**ServicesAppBatchMintByMetaUriDto**](ServicesAppBatchMintByMetaUriDto.md) | batch_mint_request | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBatchMintNFT

> []ModelsMintTask AppBatchMintNFT(ctx, id).Authorization(authorization).AppBatchMintMetaInfo(appBatchMintMetaInfo).Execute()

Batch Mint NFT with metadata parts



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
    authorization := "authorization_example" // string | Bearer JWT
    id := int32(56) // int32 | id
    appBatchMintMetaInfo := []openapiclient.ServicesAppMintByMetaPartsDto{*openapiclient.NewServicesAppMintByMetaPartsDto("Chain_example", "FileUrl_example", "MintToAddress_example", "Name_example")} // []ServicesAppMintByMetaPartsDto | mint_meta

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.AppBatchMintNFT(context.Background(), id).Authorization(authorization).AppBatchMintMetaInfo(appBatchMintMetaInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.AppBatchMintNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppBatchMintNFT`: []ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.AppBatchMintNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBatchMintNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer JWT | 

 **appBatchMintMetaInfo** | [**[]ServicesAppMintByMetaPartsDto**](ServicesAppMintByMetaPartsDto.md) | mint_meta | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT
    customMintBatchDto := *openapiclient.NewServicesCustomMintBatchDto("Chain_example", "ContractAddress_example", []openapiclient.ServicesMintItemDto{*openapiclient.NewServicesMintItemDto("MintToAddress_example")}) // ServicesCustomMintBatchDto | custom_mint_batch_dto

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT
    customMintDto := *openapiclient.NewServicesCustomMintDto("Chain_example", "ContractAddress_example", "MintToAddress_example") // ServicesCustomMintDto | custom_mint_dto

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

> ModelsMintTask EasyMintByFile(ctx).Authorization(authorization).Chain(chain).Description(description).MintToAddress(mintToAddress).Name(name).File(file).Execute()

Mint NFT with file



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
    chain := "chain_example" // string | 
    description := "description_example" // string | 
    mintToAddress := "mintToAddress_example" // string | 
    name := "name_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.EasyMintByFile(context.Background()).Authorization(authorization).Chain(chain).Description(description).MintToAddress(mintToAddress).Name(name).File(file).Execute()
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
 **chain** | **string** |  | 
 **description** | **string** |  | 
 **mintToAddress** | **string** |  | 
 **name** | **string** |  | 
 **file** | ***os.File** | file | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "authorization_example" // string | Bearer Open_JWT
    easyMintMetaInfo := *openapiclient.NewServicesEasyMintMetaPartsDto("Chain_example", "FileUrl_example", "MintToAddress_example", "Name_example") // ServicesEasyMintMetaPartsDto | easy_mint_meta_info

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
 **easyMintMetaInfo** | [**ServicesEasyMintMetaPartsDto**](ServicesEasyMintMetaPartsDto.md) | easy_mint_meta_info | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

> ModelsMintTaskQueryResult ListMints(ctx).Authorization(authorization).Authorization2(authorization2).Page(page).Limit(limit).Contract(contract).MintTo(mintTo).Status(status).Chain(chain).Page2(page2).Limit2(limit2).Contract2(contract2).MintTo2(mintTo2).Status2(status2).Chain2(chain2).Execute()

Obtain NFT list



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
    authorization2 := "authorization_example" // string | Bearer Open_JWT
    page := int32(56) // int32 | page (optional)
    limit := int32(56) // int32 | limit (optional)
    contract := "contract_example" // string | contract (optional)
    mintTo := "mintTo_example" // string | mint_to (optional)
    status := int32(56) // int32 | status (optional)
    chain := "chain_example" // string | chain (optional)
    page2 := int32(56) // int32 | page (optional)
    limit2 := int32(56) // int32 | limit (optional)
    contract2 := "contract_example" // string | contract (optional)
    mintTo2 := "mintTo_example" // string | mint_to (optional)
    status2 := int32(56) // int32 | status (optional)
    chain2 := "chain_example" // string | chain (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.ListMints(context.Background()).Authorization(authorization).Authorization2(authorization2).Page(page).Limit(limit).Contract(contract).MintTo(mintTo).Status(status).Chain(chain).Page2(page2).Limit2(limit2).Contract2(contract2).MintTo2(mintTo2).Status2(status2).Chain2(chain2).Execute()
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
 **authorization2** | **string** | Bearer Open_JWT | 
 **page** | **int32** | page | 
 **limit** | **int32** | limit | 
 **contract** | **string** | contract | 
 **mintTo** | **string** | mint_to | 
 **status** | **int32** | status | 
 **chain** | **string** | chain | 
 **page2** | **int32** | page | 
 **limit2** | **int32** | limit | 
 **contract2** | **string** | contract | 
 **mintTo2** | **string** | mint_to | 
 **status2** | **int32** | status | 
 **chain2** | **string** | chain | 

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


## ListMints_0

> ModelsMintTaskQueryResult ListMints_0(ctx).Authorization(authorization).Authorization2(authorization2).Page(page).Limit(limit).Contract(contract).MintTo(mintTo).Status(status).Chain(chain).Page2(page2).Limit2(limit2).Contract2(contract2).MintTo2(mintTo2).Status2(status2).Chain2(chain2).Execute()

Obtain NFT list



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
    authorization2 := "authorization_example" // string | Bearer Open_JWT
    page := int32(56) // int32 | page (optional)
    limit := int32(56) // int32 | limit (optional)
    contract := "contract_example" // string | contract (optional)
    mintTo := "mintTo_example" // string | mint_to (optional)
    status := int32(56) // int32 | status (optional)
    chain := "chain_example" // string | chain (optional)
    page2 := int32(56) // int32 | page (optional)
    limit2 := int32(56) // int32 | limit (optional)
    contract2 := "contract_example" // string | contract (optional)
    mintTo2 := "mintTo_example" // string | mint_to (optional)
    status2 := int32(56) // int32 | status (optional)
    chain2 := "chain_example" // string | chain (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.ListMints_0(context.Background()).Authorization(authorization).Authorization2(authorization2).Page(page).Limit(limit).Contract(contract).MintTo(mintTo).Status(status).Chain(chain).Page2(page2).Limit2(limit2).Contract2(contract2).MintTo2(mintTo2).Status2(status2).Chain2(chain2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.ListMints_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMints_0`: ModelsMintTaskQueryResult
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.ListMints_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMints_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **authorization2** | **string** | Bearer Open_JWT | 
 **page** | **int32** | page | 
 **limit** | **int32** | limit | 
 **contract** | **string** | contract | 
 **mintTo** | **string** | mint_to | 
 **status** | **int32** | status | 
 **chain** | **string** | chain | 
 **page2** | **int32** | page | 
 **limit2** | **int32** | limit | 
 **contract2** | **string** | contract | 
 **mintTo2** | **string** | mint_to | 
 **status2** | **int32** | status | 
 **chain2** | **string** | chain | 

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


## ReMintNFT

> ModelsMintTask ReMintNFT(ctx, id).Authorization(authorization).Execute()

Reset mint task status to init so that it can be minted again



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
    id := int32(56) // int32 | id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.ReMintNFT(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.ReMintNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReMintNFT`: ModelsMintTask
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.ReMintNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReMintNFTRequest struct via the builder pattern


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

