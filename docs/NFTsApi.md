# \NFTsAPI

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NFTHoldCount**](NFTsAPI.md#NFTHoldCount) | **Get** /v1/nft/count/{address}/{token_id}/{holder} | Get NFT hold count by address
[**NftInfo**](NFTsAPI.md#NftInfo) | **Get** /v1/nft/{address}/{token_id} | Get NFT info, mainly owner and metadata
[**UpdateNftTokenUri**](NFTsAPI.md#UpdateNftTokenUri) | **Put** /v1/nft/{address}/{token_id}/tokenUri | Update NFT token uri



## NFTHoldCount

> map[string]int32 NFTHoldCount(ctx, address, tokenId, holder).Authorization(authorization).Type_(type_).EpochHeight(epochHeight).Execute()

Get NFT hold count by address



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
	tokenId := "tokenId_example" // string | token_id
	holder := "holder_example" // string | holder address
	type_ := "type__example" // string | contract type erc721, erc1155, default is erc721 (optional)
	epochHeight := float32(8.14) // float32 | epoch height, default is latest_state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFTsAPI.NFTHoldCount(context.Background(), address, tokenId, holder).Authorization(authorization).Type_(type_).EpochHeight(epochHeight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTsAPI.NFTHoldCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NFTHoldCount`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `NFTsAPI.NFTHoldCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | contract address | 
**tokenId** | **string** | token_id | 
**holder** | **string** | holder address | 

### Other Parameters

Other parameters are passed through a pointer to a apiNFTHoldCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 



 **type_** | **string** | contract type erc721, erc1155, default is erc721 | 
 **epochHeight** | **float32** | epoch height, default is latest_state | 

### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NftInfo

> ServicesNFT NftInfo(ctx, address, tokenId).Authorization(authorization).Type_(type_).Execute()

Get NFT info, mainly owner and metadata



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
	tokenId := "tokenId_example" // string | token_id
	type_ := "type__example" // string | contract type erc721, erc1155, default is erc721 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFTsAPI.NftInfo(context.Background(), address, tokenId).Authorization(authorization).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTsAPI.NftInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NftInfo`: ServicesNFT
	fmt.Fprintf(os.Stdout, "Response from `NFTsAPI.NftInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 
**tokenId** | **string** | token_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNftInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


 **type_** | **string** | contract type erc721, erc1155, default is erc721 | 

### Return type

[**ServicesNFT**](ServicesNFT.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNftTokenUri

> ModelsTransaction UpdateNftTokenUri(ctx, address, tokenId).Authorization(authorization).Req(req).Execute()

Update NFT token uri



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
	tokenId := "tokenId_example" // string | token_id
	req := *openapiclient.NewRoutersUpdateTokenUriReq("Chain_example", "ContractType_example") // RoutersUpdateTokenUriReq | req

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFTsAPI.UpdateNftTokenUri(context.Background(), address, tokenId).Authorization(authorization).Req(req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTsAPI.UpdateNftTokenUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNftTokenUri`: ModelsTransaction
	fmt.Fprintf(os.Stdout, "Response from `NFTsAPI.UpdateNftTokenUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | address | 
**tokenId** | **string** | token_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNftTokenUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer Open_JWT | 


 **req** | [**RoutersUpdateTokenUriReq**](RoutersUpdateTokenUriReq.md) | req | 

### Return type

[**ModelsTransaction**](ModelsTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

