# \ContractAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContractSponsorWhitelist**](ContractAPI.md#AddContractSponsorWhitelist) | **Post** /v1/contracts/{address}/sponsor/whitelist | Add contract sponsored whitelist
[**DeployContract**](ContractAPI.md#DeployContract) | **Post** /v1/contracts/ | Deploy contract
[**DeployContract_0**](ContractAPI.md#DeployContract_0) | **Post** /v1/contracts/ | Deploy contract
[**GetContractAdmin**](ContractAPI.md#GetContractAdmin) | **Get** /v1/contracts/{address}/admin | Get administrator of contract, only work on conflux chain
[**GetContractAutoSponsor**](ContractAPI.md#GetContractAutoSponsor) | **Get** /v1/contracts/{address}/config/auto-sponsor | Get contract auto sponsor config
[**GetContractInfo**](ContractAPI.md#GetContractInfo) | **Get** /v1/contracts/detail/{id} | Contract detail
[**GetContractProfile**](ContractAPI.md#GetContractProfile) | **Get** /v1/contracts/{address}/profile | Get contract runtime profile
[**GetContractSponsorInfo**](ContractAPI.md#GetContractSponsorInfo) | **Get** /v1/contracts/{address}/sponsor | Query sponsor
[**GetContractSponsoredWhitelist**](ContractAPI.md#GetContractSponsoredWhitelist) | **Get** /v1/contracts/{address}/sponsor/whitelist | Get contract sponsored whitelist
[**ListContracts**](ContractAPI.md#ListContracts) | **Get** /v1/contracts/ | Obtain contract list
[**RemoveContractSponsorWhitelist**](ContractAPI.md#RemoveContractSponsorWhitelist) | **Delete** /v1/contracts/{address}/sponsor/whitelist | Remove contract sponsored whitelist
[**SetContractAutoSponsor**](ContractAPI.md#SetContractAutoSponsor) | **Post** /v1/contracts/{address}/config/auto-sponsor | Set contract auto sponsor config
[**SetContractSponsor**](ContractAPI.md#SetContractSponsor) | **Post** /v1/contracts/{address}/sponsor | Set sponsor
[**SetContractTransaferable**](ContractAPI.md#SetContractTransaferable) | **Post** /v1/contracts/{address}/config/transaferable | Set Contract Transaferable Config
[**UpdateContractAdmin**](ContractAPI.md#UpdateContractAdmin) | **Put** /v1/contracts/{address}/admin | Update administrator of contract



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | contract address
	users := []string{"Property_example"} // []string | Adding sponsor whitelist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.AddContractSponsorWhitelist(context.Background(), address).Authorization(authorization).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.AddContractSponsorWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContractSponsorWhitelist`: ServicesSendTxResp
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.AddContractSponsorWhitelist`: %v\n", resp)
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

> ModelsContract DeployContract(ctx).Authorization(authorization).Authorization2(authorization2).ContractInfo(contractInfo).Execute()

Deploy contract



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
	contractInfo := *openapiclient.NewServicesContractDeployDto("Chain_example", "Name_example", "Symbol_example", "Type_example") // ServicesContractDeployDto | contract_info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.DeployContract(context.Background()).Authorization(authorization).Authorization2(authorization2).ContractInfo(contractInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.DeployContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployContract`: ModelsContract
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.DeployContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **authorization2** | **string** | Bearer Open_JWT | 
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


## DeployContract_0

> ModelsContract DeployContract_0(ctx).Authorization(authorization).Authorization2(authorization2).ContractInfo(contractInfo).Execute()

Deploy contract



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
	contractInfo := *openapiclient.NewServicesContractDeployDto("Chain_example", "Name_example", "Symbol_example", "Type_example") // ServicesContractDeployDto | contract_info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.DeployContract_0(context.Background()).Authorization(authorization).Authorization2(authorization2).ContractInfo(contractInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.DeployContract_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployContract_0`: ModelsContract
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.DeployContract_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployContract_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **authorization2** | **string** | Bearer Open_JWT | 
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | contract address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractAdmin(context.Background(), address).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractAdmin`: string
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractAdmin`: %v\n", resp)
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


## GetContractAutoSponsor

> ModelsAutoSponsorContract GetContractAutoSponsor(ctx, address).Authorization(authorization).Execute()

Get contract auto sponsor config



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
	address := "address_example" // string | contract address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractAutoSponsor(context.Background(), address).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractAutoSponsor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractAutoSponsor`: ModelsAutoSponsorContract
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractAutoSponsor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractAutoSponsorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ModelsAutoSponsorContract**](ModelsAutoSponsorContract.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	id := int32(56) // int32 | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractInfo(context.Background(), id).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractInfo`: ModelsContract
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractInfo`: %v\n", resp)
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


## GetContractProfile

> ModelsContractRuntimeProfile GetContractProfile(ctx, address).Authorization(authorization).IgnoreTokenIds(ignoreTokenIds).Execute()

Get contract runtime profile



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
	ignoreTokenIds := "ignoreTokenIds_example" // string | the returned max token id will ignore the token ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractProfile(context.Background(), address).Authorization(authorization).IgnoreTokenIds(ignoreTokenIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractProfile`: ModelsContractRuntimeProfile
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **ignoreTokenIds** | **string** | the returned max token id will ignore the token ids | 

### Return type

[**ModelsContractRuntimeProfile**](ModelsContractRuntimeProfile.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | address
	chain := "chain_example" // string | chain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractSponsorInfo(context.Background(), address).Authorization(authorization).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractSponsorInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractSponsorInfo`: ServicesSponsorInfo
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractSponsorInfo`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | contract address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetContractSponsoredWhitelist(context.Background(), address).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetContractSponsoredWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractSponsoredWhitelist`: []string
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetContractSponsoredWhitelist`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	page := int32(56) // int32 | page (optional)
	limit := int32(56) // int32 | limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.ListContracts(context.Background()).Authorization(authorization).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.ListContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContracts`: ModelsContractTaskQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.ListContracts`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | contract address
	users := []string{"Property_example"} // []string | Removing sponsor whitelist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.RemoveContractSponsorWhitelist(context.Background(), address).Authorization(authorization).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.RemoveContractSponsorWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContractSponsorWhitelist`: ServicesSendTxResp
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.RemoveContractSponsorWhitelist`: %v\n", resp)
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


## SetContractAutoSponsor

> map[string]map[string]interface{} SetContractAutoSponsor(ctx, address).Authorization(authorization).AutoSponsorReq(autoSponsorReq).Execute()

Set contract auto sponsor config



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
	address := "address_example" // string | contract address
	autoSponsorReq := *openapiclient.NewServicesContractAutoSponsorReq() // ServicesContractAutoSponsorReq | contract auto sponsor config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.SetContractAutoSponsor(context.Background(), address).Authorization(authorization).AutoSponsorReq(autoSponsorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.SetContractAutoSponsor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetContractAutoSponsor`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.SetContractAutoSponsor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContractAutoSponsorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **autoSponsorReq** | [**ServicesContractAutoSponsorReq**](ServicesContractAutoSponsorReq.md) | contract auto sponsor config | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetContractSponsor

> ServicesSetSponsorResp SetContractSponsor(ctx, address).Authorization(authorization).Chain(chain).Gas(gas).Storage(storage).AutoSponsor(autoSponsor).Execute()

Set sponsor



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
	address := "address_example" // string | Contract address
	chain := "chain_example" // string | chain: conflux, conflux_test(default) (optional)
	gas := "gas_example" // string | gas: default value is 1 (optional)
	storage := "storage_example" // string | storage: default value is by user default setting (optional)
	autoSponsor := true // bool | Open auto sponsor or not, for mainnet contract keep user account have enough balance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.SetContractSponsor(context.Background(), address).Authorization(authorization).Chain(chain).Gas(gas).Storage(storage).AutoSponsor(autoSponsor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.SetContractSponsor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetContractSponsor`: ServicesSetSponsorResp
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.SetContractSponsor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContractSponsorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 

 **chain** | **string** | chain: conflux, conflux_test(default) | 
 **gas** | **string** | gas: default value is 1 | 
 **storage** | **string** | storage: default value is by user default setting | 
 **autoSponsor** | **bool** | Open auto sponsor or not, for mainnet contract keep user account have enough balance | 

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


## SetContractTransaferable

> ServicesSendTxResp SetContractTransaferable(ctx, address).Authorization(authorization).Execute()

Set Contract Transaferable Config



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
	address := "address_example" // string | contract address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.SetContractTransaferable(context.Background(), address).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.SetContractTransaferable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetContractTransaferable`: ServicesSendTxResp
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.SetContractTransaferable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContractTransaferableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Bearer Open_JWT
	address := "address_example" // string | contract address
	adminInfo := *openapiclient.NewServicesContractAdminUpdateDto("AdminAddress_example") // ServicesContractAdminUpdateDto | contract admin update info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.UpdateContractAdmin(context.Background(), address).Authorization(authorization).AdminInfo(adminInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.UpdateContractAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContractAdmin`: ServicesSendTxResp
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.UpdateContractAdmin`: %v\n", resp)
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

