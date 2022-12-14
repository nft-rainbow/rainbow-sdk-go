# \LoginApi

All URIs are relative to *http://api.nftrainbow.cn/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginApp**](LoginApi.md#LoginApp) | **Post** /login | App login
[**RefreshAuth**](LoginApi.md#RefreshAuth) | **Get** /refresh_token | Refresh JWT



## LoginApp

> MiddlewaresLoginResp LoginApp(ctx).AppLoginInfo(appLoginInfo).Execute()

App login



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
    appLoginInfo := *openapiclient.NewMiddlewaresAppLogin("AppId_example", "AppSecret_example") // MiddlewaresAppLogin | login info, contain app_id and app_secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.LoginApp(context.Background()).AppLoginInfo(appLoginInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.LoginApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginApp`: MiddlewaresLoginResp
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.LoginApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appLoginInfo** | [**MiddlewaresAppLogin**](MiddlewaresAppLogin.md) | login info, contain app_id and app_secret | 

### Return type

[**MiddlewaresLoginResp**](MiddlewaresLoginResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAuth

> MiddlewaresLoginResp RefreshAuth(ctx).Authorization(authorization).Execute()

Refresh JWT



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
    authorization := "authorization_example" // string | Bearer openapi_token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.RefreshAuth(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.RefreshAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAuth`: MiddlewaresLoginResp
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.RefreshAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer openapi_token | 

### Return type

[**MiddlewaresLoginResp**](MiddlewaresLoginResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

