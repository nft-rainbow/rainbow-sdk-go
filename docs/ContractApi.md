# \ContractApi

All URIs are relative to *http://api.nftrainbow.xyz/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployContract**](ContractApi.md#DeployContract) | **Post** /contracts/ | Deploy contract
[**GetContractInfo**](ContractApi.md#GetContractInfo) | **Get** /contracts/detail/{id} | Contract detail
[**GetContractSponsorInfo**](ContractApi.md#GetContractSponsorInfo) | **Get** /contracts/{address}/sponsor | Query sponsor
[**ListContracts**](ContractApi.md#ListContracts) | **Get** /contracts/ | Obtain contract list
[**SetContractSponsor**](ContractApi.md#SetContractSponsor) | **Post** /contracts/{address}/sponsor | Set sponsor



## DeployContract

> ModelsContract DeployContract(ctx).Authorization(authorization).ContractInfo(contractInfo).Execute()

Deploy contract



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
    contractInfo := *openapiclient.NewServicesContractDeployDto("Chain_example", "Name_example", "OwnerAddress_example", "Symbol_example", "Type_example") // ServicesContractDeployDto | contract_info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.DeployContract(context.Background()).Authorization(authorization).ContractInfo(contractInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.DeployContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployContract`: ModelsContract
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.DeployContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **contractInfo** | [**ServicesContractDeployDto**](ServicesContractDeployDto.md) | contract_info | 

### Return type

[**ModelsContract**](ModelsContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractInfo

> ModelsContract GetContractInfo(ctx, id).Authorization(authorization).Execute()

Contract detail



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
    resp, r, err := apiClient.ContractApi.GetContractInfo(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractInfo`: ModelsContract
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ModelsContract**](ModelsContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractSponsorInfo

> ServicesSponsorInfo GetContractSponsorInfo(ctx, address).Authorization(authorization).Chain(chain).Execute()

Query sponsor



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
    chain := "chain_example" // string | chain (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.GetContractSponsorInfo(context.Background(), address).Authorization(authorization).Chain(chain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractSponsorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractSponsorInfo`: ServicesSponsorInfo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractSponsorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSponsorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **chain** | **string** | chain | 

### Return type

[**ServicesSponsorInfo**](ServicesSponsorInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContracts

> ModelsContractTaskQueryResult ListContracts(ctx).Authorization(authorization).Page(page).Limit(limit).Execute()

Obtain contract list



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
    resp, r, err := apiClient.ContractApi.ListContracts(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ListContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContracts`: ModelsContractTaskQueryResult
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ListContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **page** | **int32** | page | 
 **limit** | **int32** | limit | 

### Return type

[**ModelsContractTaskQueryResult**](ModelsContractTaskQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetContractSponsor

> string SetContractSponsor(ctx, address).Authorization(authorization).Execute()

Set sponsor



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.SetContractSponsor(context.Background(), address).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.SetContractSponsor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetContractSponsor`: string
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.SetContractSponsor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContractSponsorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

