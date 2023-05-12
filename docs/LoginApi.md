# \LoginApi

All URIs are relative to *http://api.nftrainbow.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginApp**](LoginApi.md#LoginApp) | **Post** /v1/login | App login
[**RefreshAppAuth**](LoginApi.md#RefreshAppAuth) | **Get** /v1/refresh_token | Refresh JWT
[**RefreshUserAuth**](LoginApi.md#RefreshUserAuth) | **Get** /dashboard/refresh_token | Refresh JWT
[**UserLogin**](LoginApi.md#UserLogin) | **Post** /dashboard/login | User login



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appLoginInfo := *openapiclient.NewMiddlewaresAppLoginInfo("AppId_example", "AppSecret_example") // MiddlewaresAppLoginInfo | login info, contain app_id and app_secret

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
 **appLoginInfo** | [**MiddlewaresAppLoginInfo**](MiddlewaresAppLoginInfo.md) | login info, contain app_id and app_secret | 

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


## RefreshAppAuth

> MiddlewaresLoginResp RefreshAppAuth(ctx).Authorization(authorization).Execute()

Refresh JWT



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
    authorization := "authorization_example" // string | Bearer openapi_token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.RefreshAppAuth(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.RefreshAppAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAppAuth`: MiddlewaresLoginResp
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.RefreshAppAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAppAuthRequest struct via the builder pattern


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


## RefreshUserAuth

> MiddlewaresLoginResp RefreshUserAuth(ctx).Authorization(authorization).Execute()

Refresh JWT



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
    authorization := "authorization_example" // string | Bearer openapi_token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.RefreshUserAuth(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.RefreshUserAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshUserAuth`: MiddlewaresLoginResp
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.RefreshUserAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshUserAuthRequest struct via the builder pattern


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


## UserLogin

> MiddlewaresLoginResp UserLogin(ctx).UserLoginInfo(userLoginInfo).Execute()

User login



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
    userLoginInfo := *openapiclient.NewMiddlewaresUserLoginInfo("Email_example", "Password_example") // MiddlewaresUserLoginInfo | login info, contain app_id and app_secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.UserLogin(context.Background()).UserLoginInfo(userLoginInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.UserLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserLogin`: MiddlewaresLoginResp
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.UserLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLoginInfo** | [**MiddlewaresUserLoginInfo**](MiddlewaresUserLoginInfo.md) | login info, contain app_id and app_secret | 

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

