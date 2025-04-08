# \AccountsAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsertAccount**](AccountsAPI.md#InsertAccount) | **Post** /v1/accounts | Insert web3 account
[**QueryAccounts**](AccountsAPI.md#QueryAccounts) | **Get** /v1/accounts | Query web3 account



## InsertAccount

> []ModelsCustodialAccountDisplay InsertAccount(ctx).Authorization(authorization).InsertAccountReq(insertAccountReq).Execute()

Insert web3 account



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
	insertAccountReq := *openapiclient.NewServicesInsertAccountReq() // ServicesInsertAccountReq | insert_account_req

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.InsertAccount(context.Background()).Authorization(authorization).InsertAccountReq(insertAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.InsertAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertAccount`: []ModelsCustodialAccountDisplay
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.InsertAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **insertAccountReq** | [**ServicesInsertAccountReq**](ServicesInsertAccountReq.md) | insert_account_req | 

### Return type

[**[]ModelsCustodialAccountDisplay**](ModelsCustodialAccountDisplay.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAccounts

> []ModelsCustodialAccountDisplay QueryAccounts(ctx).Authorization(authorization).Phone(phone).Owned(owned).Execute()

Query web3 account



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
	phone := "phone_example" // string | 
	owned := true // bool | is created by user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.QueryAccounts(context.Background()).Authorization(authorization).Phone(phone).Owned(owned).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.QueryAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAccounts`: []ModelsCustodialAccountDisplay
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.QueryAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **phone** | **string** |  | 
 **owned** | **bool** | is created by user | 

### Return type

[**[]ModelsCustodialAccountDisplay**](ModelsCustodialAccountDisplay.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

