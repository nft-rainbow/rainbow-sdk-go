# \AccountsAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsertAccount**](AccountsAPI.md#InsertAccount) | **Post** /v1/accounts | Insert web3 account
[**QueryAccount**](AccountsAPI.md#QueryAccount) | **Get** /v1/accounts | Query web3 account



## InsertAccount

> []ModelsAccountDisplayPart InsertAccount(ctx).Authorization(authorization).InsertAccountReq(insertAccountReq).Execute()

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
	// response from `InsertAccount`: []ModelsAccountDisplayPart
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

[**[]ModelsAccountDisplayPart**](ModelsAccountDisplayPart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAccount

> []ModelsAccountDisplayPart QueryAccount(ctx).Authorization(authorization).Phone(phone).Execute()

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
	phone := "phone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.QueryAccount(context.Background()).Authorization(authorization).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.QueryAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAccount`: []ModelsAccountDisplayPart
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.QueryAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **phone** | **string** |  | 

### Return type

[**[]ModelsAccountDisplayPart**](ModelsAccountDisplayPart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

