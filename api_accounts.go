/*
Rainbow-API

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rainbowsdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type AccountsAPI interface {

	/*
	InsertAccount Insert web3 account

	Insert web3 account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInsertAccountRequest
	*/
	InsertAccount(ctx context.Context) ApiInsertAccountRequest

	// InsertAccountExecute executes the request
	//  @return []ModelsCustodialAccountDisplay
	InsertAccountExecute(r ApiInsertAccountRequest) ([]ModelsCustodialAccountDisplay, *http.Response, error)

	/*
	QueryAccounts Query web3 account

	Query web3 account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryAccountsRequest
	*/
	QueryAccounts(ctx context.Context) ApiQueryAccountsRequest

	// QueryAccountsExecute executes the request
	//  @return []ModelsCustodialAccountDisplay
	QueryAccountsExecute(r ApiQueryAccountsRequest) ([]ModelsCustodialAccountDisplay, *http.Response, error)
}

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiInsertAccountRequest struct {
	ctx context.Context
	ApiService AccountsAPI
	authorization *string
	insertAccountReq *ServicesInsertAccountReq
}

// Bearer Open_JWT
func (r ApiInsertAccountRequest) Authorization(authorization string) ApiInsertAccountRequest {
	r.authorization = &authorization
	return r
}

// insert_account_req
func (r ApiInsertAccountRequest) InsertAccountReq(insertAccountReq ServicesInsertAccountReq) ApiInsertAccountRequest {
	r.insertAccountReq = &insertAccountReq
	return r
}

func (r ApiInsertAccountRequest) Execute() ([]ModelsCustodialAccountDisplay, *http.Response, error) {
	return r.ApiService.InsertAccountExecute(r)
}

/*
InsertAccount Insert web3 account

Insert web3 account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInsertAccountRequest
*/
func (a *AccountsAPIService) InsertAccount(ctx context.Context) ApiInsertAccountRequest {
	return ApiInsertAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelsCustodialAccountDisplay
func (a *AccountsAPIService) InsertAccountExecute(r ApiInsertAccountRequest) ([]ModelsCustodialAccountDisplay, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelsCustodialAccountDisplay
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.InsertAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.insertAccountReq == nil {
		return localVarReturnValue, nil, reportError("insertAccountReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "", "")
	// body params
	localVarPostBody = r.insertAccountReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RainbowErrorsRainbowErrorDetailInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RainbowErrorsRainbowErrorDetailInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryAccountsRequest struct {
	ctx context.Context
	ApiService AccountsAPI
	authorization *string
	phone *string
	owned *bool
}

// Bearer Open_JWT
func (r ApiQueryAccountsRequest) Authorization(authorization string) ApiQueryAccountsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiQueryAccountsRequest) Phone(phone string) ApiQueryAccountsRequest {
	r.phone = &phone
	return r
}

// is created by user
func (r ApiQueryAccountsRequest) Owned(owned bool) ApiQueryAccountsRequest {
	r.owned = &owned
	return r
}

func (r ApiQueryAccountsRequest) Execute() ([]ModelsCustodialAccountDisplay, *http.Response, error) {
	return r.ApiService.QueryAccountsExecute(r)
}

/*
QueryAccounts Query web3 account

Query web3 account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryAccountsRequest
*/
func (a *AccountsAPIService) QueryAccounts(ctx context.Context) ApiQueryAccountsRequest {
	return ApiQueryAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelsCustodialAccountDisplay
func (a *AccountsAPIService) QueryAccountsExecute(r ApiQueryAccountsRequest) ([]ModelsCustodialAccountDisplay, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelsCustodialAccountDisplay
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.QueryAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.phone == nil {
		return localVarReturnValue, nil, reportError("phone is required and must be specified")
	}

	if r.owned != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "owned", r.owned, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "phone", r.phone, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RainbowErrorsRainbowErrorDetailInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RainbowErrorsRainbowErrorDetailInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
