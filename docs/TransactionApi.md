# \TransactionAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildNftApproveTx**](TransactionAPI.md#BuildNftApproveTx) | **Post** /v1/tx/build/nft/approve | Build NFT approve transaction
[**BuildNftSetApprovalForAllTx**](TransactionAPI.md#BuildNftSetApprovalForAllTx) | **Post** /v1/tx/build/nft/set-approval-for-all | Build NFT set approval for all transaction
[**BuildNftTransferTx**](TransactionAPI.md#BuildNftTransferTx) | **Post** /v1/tx/build/nft/transfer | Build NFT transfer transaction
[**GetTransactionByID**](TransactionAPI.md#GetTransactionByID) | **Get** /v1/tx/{id} | Get transaction informantion by ID
[**SendTransaction**](TransactionAPI.md#SendTransaction) | **Post** /v1/tx | Send a transaction



## BuildNftApproveTx

> ServicesBuildTxResp BuildNftApproveTx(ctx).Authorization(authorization).ApproveTxReq(approveTxReq).Execute()

Build NFT approve transaction



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
	approveTxReq := *openapiclient.NewServicesBuildNftApproveTxReq("Chain_example", "ContractAddress_example", "ContractType_example", "From_example", "To_example", "TokenId_example") // ServicesBuildNftApproveTxReq | Approve transaction request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.BuildNftApproveTx(context.Background()).Authorization(authorization).ApproveTxReq(approveTxReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.BuildNftApproveTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildNftApproveTx`: ServicesBuildTxResp
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.BuildNftApproveTx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildNftApproveTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **approveTxReq** | [**ServicesBuildNftApproveTxReq**](ServicesBuildNftApproveTxReq.md) | Approve transaction request | 

### Return type

[**ServicesBuildTxResp**](ServicesBuildTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildNftSetApprovalForAllTx

> ServicesBuildTxResp BuildNftSetApprovalForAllTx(ctx).Authorization(authorization).SetApprovalForAllTxReq(setApprovalForAllTxReq).Execute()

Build NFT set approval for all transaction



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
	setApprovalForAllTxReq := *openapiclient.NewServicesBuildNftSetApprovalForAllTxReq(false, "Chain_example", "ContractAddress_example", "ContractType_example", "From_example", "Operator_example") // ServicesBuildNftSetApprovalForAllTxReq | Set approval for all transaction request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.BuildNftSetApprovalForAllTx(context.Background()).Authorization(authorization).SetApprovalForAllTxReq(setApprovalForAllTxReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.BuildNftSetApprovalForAllTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildNftSetApprovalForAllTx`: ServicesBuildTxResp
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.BuildNftSetApprovalForAllTx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildNftSetApprovalForAllTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **setApprovalForAllTxReq** | [**ServicesBuildNftSetApprovalForAllTxReq**](ServicesBuildNftSetApprovalForAllTxReq.md) | Set approval for all transaction request | 

### Return type

[**ServicesBuildTxResp**](ServicesBuildTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildNftTransferTx

> ServicesBuildTxResp BuildNftTransferTx(ctx).Authorization(authorization).TransferDto(transferDto).Execute()

Build NFT transfer transaction



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
	transferDto := *openapiclient.NewServicesTransferDto("Chain_example", "ContractAddress_example", "ContractType_example", "TokenId_example", "TransferFromAddress_example", "TransferToAddress_example") // ServicesTransferDto | Transfer DTO

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.BuildNftTransferTx(context.Background()).Authorization(authorization).TransferDto(transferDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.BuildNftTransferTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildNftTransferTx`: ServicesBuildTxResp
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.BuildNftTransferTx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildNftTransferTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **transferDto** | [**ServicesTransferDto**](ServicesTransferDto.md) | Transfer DTO | 

### Return type

[**ServicesBuildTxResp**](ServicesBuildTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionByID

> ServicesTxResp GetTransactionByID(ctx, id).Authorization(authorization).Execute()

Get transaction informantion by ID



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
	id := "id_example" // string | Transaction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.GetTransactionByID(context.Background(), id).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.GetTransactionByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionByID`: ServicesTxResp
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.GetTransactionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


### Return type

[**ServicesTxResp**](ServicesTxResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTransaction

> ServicesSendTxResp SendTransaction(ctx).Authorization(authorization).SendTxReq(sendTxReq).Execute()

Send a transaction



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
	sendTxReq := *openapiclient.NewServicesSendTxReq("Chain_example", "From_example", "To_example") // ServicesSendTxReq | Send transaction request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.SendTransaction(context.Background()).Authorization(authorization).SendTxReq(sendTxReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.SendTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTransaction`: ServicesSendTxResp
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.SendTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 
 **sendTxReq** | [**ServicesSendTxReq**](ServicesSendTxReq.md) | Send transaction request | 

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

