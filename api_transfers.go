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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type TransfersApi interface {

	/*
	BatchTransferNft Batch Transfer NFTs

	Transfer several NFTs once

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBatchTransferNftRequest
	*/
	BatchTransferNft(ctx context.Context) ApiBatchTransferNftRequest

	// BatchTransferNftExecute executes the request
	//  @return []ModelsTransferTask
	BatchTransferNftExecute(r ApiBatchTransferNftRequest) ([]ModelsTransferTask, *http.Response, error)

	/*
	GetTransferDetail Transfer NFT detail

	Get NFT Transfer detail info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id
	@return ApiGetTransferDetailRequest
	*/
	GetTransferDetail(ctx context.Context, id int32) ApiGetTransferDetailRequest

	// GetTransferDetailExecute executes the request
	//  @return ModelsTransferTask
	GetTransferDetailExecute(r ApiGetTransferDetailRequest) (*ModelsTransferTask, *http.Response, error)

	/*
	ListTransfer Obtain the transferred NFTs list

	Get the NFT list containing the transferred NFT information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTransferRequest
	*/
	ListTransfer(ctx context.Context) ApiListTransferRequest

	// ListTransferExecute executes the request
	//  @return ModelsTransferTaskQueryResult
	ListTransferExecute(r ApiListTransferRequest) (*ModelsTransferTaskQueryResult, *http.Response, error)

	/*
	TransferNft Transfer NFT

	Transfer NFT by admin

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTransferNftRequest
	*/
	TransferNft(ctx context.Context) ApiTransferNftRequest

	// TransferNftExecute executes the request
	//  @return ModelsTransferTask
	TransferNftExecute(r ApiTransferNftRequest) (*ModelsTransferTask, *http.Response, error)
}

// TransfersApiService TransfersApi service
type TransfersApiService service

type ApiBatchTransferNftRequest struct {
	ctx context.Context
	ApiService TransfersApi
	authorization *string
	transferBatchDto *ServicesTransferBatchDto
}

// Bearer Open_JWT
func (r ApiBatchTransferNftRequest) Authorization(authorization string) ApiBatchTransferNftRequest {
	r.authorization = &authorization
	return r
}

// transfer_batch_dto
func (r ApiBatchTransferNftRequest) TransferBatchDto(transferBatchDto ServicesTransferBatchDto) ApiBatchTransferNftRequest {
	r.transferBatchDto = &transferBatchDto
	return r
}

func (r ApiBatchTransferNftRequest) Execute() ([]ModelsTransferTask, *http.Response, error) {
	return r.ApiService.BatchTransferNftExecute(r)
}

/*
BatchTransferNft Batch Transfer NFTs

Transfer several NFTs once

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBatchTransferNftRequest
*/
func (a *TransfersApiService) BatchTransferNft(ctx context.Context) ApiBatchTransferNftRequest {
	return ApiBatchTransferNftRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelsTransferTask
func (a *TransfersApiService) BatchTransferNftExecute(r ApiBatchTransferNftRequest) ([]ModelsTransferTask, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelsTransferTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.BatchTransferNft")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transfers/customizable/batch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.transferBatchDto == nil {
		return localVarReturnValue, nil, reportError("transferBatchDto is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	// body params
	localVarPostBody = r.transferBatchDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetTransferDetailRequest struct {
	ctx context.Context
	ApiService TransfersApi
	authorization *string
	id int32
}

// Bearer Open_JWT
func (r ApiGetTransferDetailRequest) Authorization(authorization string) ApiGetTransferDetailRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetTransferDetailRequest) Execute() (*ModelsTransferTask, *http.Response, error) {
	return r.ApiService.GetTransferDetailExecute(r)
}

/*
GetTransferDetail Transfer NFT detail

Get NFT Transfer detail info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id
 @return ApiGetTransferDetailRequest
*/
func (a *TransfersApiService) GetTransferDetail(ctx context.Context, id int32) ApiGetTransferDetailRequest {
	return ApiGetTransferDetailRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsTransferTask
func (a *TransfersApiService) GetTransferDetailExecute(r ApiGetTransferDetailRequest) (*ModelsTransferTask, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsTransferTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.GetTransferDetail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transfers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListTransferRequest struct {
	ctx context.Context
	ApiService TransfersApi
	authorization *string
	page *int32
	limit *int32
}

// Bearer Open_JWT
func (r ApiListTransferRequest) Authorization(authorization string) ApiListTransferRequest {
	r.authorization = &authorization
	return r
}

// page
func (r ApiListTransferRequest) Page(page int32) ApiListTransferRequest {
	r.page = &page
	return r
}

// limit
func (r ApiListTransferRequest) Limit(limit int32) ApiListTransferRequest {
	r.limit = &limit
	return r
}

func (r ApiListTransferRequest) Execute() (*ModelsTransferTaskQueryResult, *http.Response, error) {
	return r.ApiService.ListTransferExecute(r)
}

/*
ListTransfer Obtain the transferred NFTs list

Get the NFT list containing the transferred NFT information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListTransferRequest
*/
func (a *TransfersApiService) ListTransfer(ctx context.Context) ApiListTransferRequest {
	return ApiListTransferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsTransferTaskQueryResult
func (a *TransfersApiService) ListTransferExecute(r ApiListTransferRequest) (*ModelsTransferTaskQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsTransferTaskQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.ListTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transfers/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiTransferNftRequest struct {
	ctx context.Context
	ApiService TransfersApi
	authorization *string
	transferDto *ServicesTransferDto
}

// Bearer Open_JWT
func (r ApiTransferNftRequest) Authorization(authorization string) ApiTransferNftRequest {
	r.authorization = &authorization
	return r
}

// transfer_dto
func (r ApiTransferNftRequest) TransferDto(transferDto ServicesTransferDto) ApiTransferNftRequest {
	r.transferDto = &transferDto
	return r
}

func (r ApiTransferNftRequest) Execute() (*ModelsTransferTask, *http.Response, error) {
	return r.ApiService.TransferNftExecute(r)
}

/*
TransferNft Transfer NFT

Transfer NFT by admin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTransferNftRequest
*/
func (a *TransfersApiService) TransferNft(ctx context.Context) ApiTransferNftRequest {
	return ApiTransferNftRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsTransferTask
func (a *TransfersApiService) TransferNftExecute(r ApiTransferNftRequest) (*ModelsTransferTask, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsTransferTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.TransferNft")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transfers/customizable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.transferDto == nil {
		return localVarReturnValue, nil, reportError("transferDto is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	// body params
	localVarPostBody = r.transferDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
