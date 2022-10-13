# \ContractApi

All URIs are relative to *http://api.nftrainbow.cn/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContractSponsorWhitelist**](ContractApi.md#AddContractSponsorWhitelist) | **Post** /contracts/{address}/sponsor/whitelist | Add contract sponsored whitelist
[**DeployContract**](ContractApi.md#DeployContract) | **Post** /contracts/ | Deploy contract
[**GetContractAdmin**](ContractApi.md#GetContractAdmin) | **Get** /contracts/{address}/admin | Get administrator of contract, only work on conflux chain
[**GetContractInfo**](ContractApi.md#GetContractInfo) | **Get** /contracts/detail/{id} | Contract detail
[**GetContractSponsorInfo**](ContractApi.md#GetContractSponsorInfo) | **Get** /contracts/{address}/sponsor | Query sponsor
[**GetContractSponsoredWhitelist**](ContractApi.md#GetContractSponsoredWhitelist) | **Get** /contracts/{address}/sponsor/whitelist | Get contract sponsored whitelist
[**ListContracts**](ContractApi.md#ListContracts) | **Get** /contracts/ | Obtain contract list
[**RemoveContractSponsorWhitelist**](ContractApi.md#RemoveContractSponsorWhitelist) | **Delete** /contracts/{address}/sponsor/whitelist | Remove contract sponsored whitelist
[**SetContractSponsor**](ContractApi.md#SetContractSponsor) | **Post** /contracts/{address}/sponsor | Set sponsor
[**UpdateContractAdmin**](ContractApi.md#UpdateContractAdmin) | **Put** /contracts/{address}/admin | Update administrator of contract



## AddContractSponsorWhitelist

> ServicesSendTxResp AddContractSponsorWhitelist(ctx, address).Authorization(authorization).Users(users).Execute()

Add contract sponsored whitelist



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
    address := "address_example" // string | contract address
    users := []string{"Property_example"} // []string | Adding sponsor whitelist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.AddContractSponsorWhitelist(context.Background(), address).Authorization(authorization).Users(users).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.AddContractSponsorWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddContractSponsorWhitelist`: ServicesSendTxResp
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.AddContractSponsorWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContractSponsorWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **users** | **[]string** | Adding sponsor whitelist | 

### Return type

[**ServicesSendTxResp**](ServicesSendTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    contractInfo := *openapiclient.NewServicesContractDeployDto("Chain_example", false, "Name_example", "OwnerAddress_example", "Symbol_example", "Type_example") // ServicesContractDeployDto | contract_info

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


## GetContractAdmin

> string GetContractAdmin(ctx, address).Authorization(authorization).Execute()

Get administrator of contract, only work on conflux chain



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
    address := "address_example" // string | contract address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.GetContractAdmin(context.Background(), address).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractAdmin`: string
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractAdminRequest struct via the builder pattern


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


## GetContractSponsoredWhitelist

> []string GetContractSponsoredWhitelist(ctx, address).Authorization(authorization).Execute()

Get contract sponsored whitelist



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
    address := "address_example" // string | contract address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.GetContractSponsoredWhitelist(context.Background(), address).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractSponsoredWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractSponsoredWhitelist`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractSponsoredWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSponsoredWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

**[]string**

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


## RemoveContractSponsorWhitelist

> ServicesSendTxResp RemoveContractSponsorWhitelist(ctx, address).Authorization(authorization).Users(users).Execute()

Remove contract sponsored whitelist



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
    address := "address_example" // string | contract address
    users := []string{"Property_example"} // []string | Removing sponsor whitelist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.RemoveContractSponsorWhitelist(context.Background(), address).Authorization(authorization).Users(users).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.RemoveContractSponsorWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveContractSponsorWhitelist`: ServicesSendTxResp
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.RemoveContractSponsorWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContractSponsorWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **users** | **[]string** | Removing sponsor whitelist | 

### Return type

[**ServicesSendTxResp**](ServicesSendTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetContractSponsor

> ServicesSetSponsorResp SetContractSponsor(ctx, address).Authorization(authorization).Execute()

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
    // response from `SetContractSponsor`: ServicesSetSponsorResp
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

[**ServicesSetSponsorResp**](ServicesSetSponsorResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContractAdmin

> ServicesSendTxResp UpdateContractAdmin(ctx, address).Authorization(authorization).AdminInfo(adminInfo).Execute()

Update administrator of contract



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
    address := "address_example" // string | contract address
    adminInfo := *openapiclient.NewServicesContractAdminUpdateDto("AdminAddress_example") // ServicesContractAdminUpdateDto | contract admin update info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractApi.UpdateContractAdmin(context.Background(), address).Authorization(authorization).AdminInfo(adminInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.UpdateContractAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContractAdmin`: ServicesSendTxResp
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.UpdateContractAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContractAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **adminInfo** | [**ServicesContractAdminUpdateDto**](ServicesContractAdminUpdateDto.md) | contract admin update info | 

### Return type

[**ServicesSendTxResp**](ServicesSendTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

