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
	"strings"
)


type MetadataAPI interface {

	/*
	CreateMetadata Create NFT metadata

	Create NFT metadata by providing the info including name, description and so on

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateMetadataRequest
	*/
	CreateMetadata(ctx context.Context) ApiCreateMetadataRequest

	// CreateMetadataExecute executes the request
	//  @return ModelsExposedMetadata
	CreateMetadataExecute(r ApiCreateMetadataRequest) (*ModelsExposedMetadata, *http.Response, error)

	/*
	GetMetadatInfo Query metadata

	Query the metadata according to metadata_id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param metadataId metadata_id
	@return ApiGetMetadatInfoRequest
	*/
	GetMetadatInfo(ctx context.Context, metadataId string) ApiGetMetadatInfoRequest

	// GetMetadatInfoExecute executes the request
	//  @return ModelsExposedMetadata
	GetMetadatInfoExecute(r ApiGetMetadatInfoRequest) (*ModelsExposedMetadata, *http.Response, error)

	/*
	ListMetadatas Obtain metadata list

	Get the metadata list containing the info of the metadata created in the specified app

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListMetadatasRequest
	*/
	ListMetadatas(ctx context.Context) ApiListMetadatasRequest

	// ListMetadatasExecute executes the request
	//  @return ModelsExposedMetadataQueryResult
	ListMetadatasExecute(r ApiListMetadatasRequest) (*ModelsExposedMetadataQueryResult, *http.Response, error)
}

// MetadataAPIService MetadataAPI service
type MetadataAPIService service

type ApiCreateMetadataRequest struct {
	ctx context.Context
	ApiService MetadataAPI
	authorization *string
	metadataInfo *ServicesMetadataDto
}

// Bearer openapi_token
func (r ApiCreateMetadataRequest) Authorization(authorization string) ApiCreateMetadataRequest {
	r.authorization = &authorization
	return r
}

// metadata_info
func (r ApiCreateMetadataRequest) MetadataInfo(metadataInfo ServicesMetadataDto) ApiCreateMetadataRequest {
	r.metadataInfo = &metadataInfo
	return r
}

func (r ApiCreateMetadataRequest) Execute() (*ModelsExposedMetadata, *http.Response, error) {
	return r.ApiService.CreateMetadataExecute(r)
}

/*
CreateMetadata Create NFT metadata

Create NFT metadata by providing the info including name, description and so on

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateMetadataRequest
*/
func (a *MetadataAPIService) CreateMetadata(ctx context.Context) ApiCreateMetadataRequest {
	return ApiCreateMetadataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsExposedMetadata
func (a *MetadataAPIService) CreateMetadataExecute(r ApiCreateMetadataRequest) (*ModelsExposedMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsExposedMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataAPIService.CreateMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.metadataInfo == nil {
		return localVarReturnValue, nil, reportError("metadataInfo is required and must be specified")
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
	localVarPostBody = r.metadataInfo
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

type ApiGetMetadatInfoRequest struct {
	ctx context.Context
	ApiService MetadataAPI
	authorization *string
	metadataId string
}

// Bearer openapi_token
func (r ApiGetMetadatInfoRequest) Authorization(authorization string) ApiGetMetadatInfoRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetMetadatInfoRequest) Execute() (*ModelsExposedMetadata, *http.Response, error) {
	return r.ApiService.GetMetadatInfoExecute(r)
}

/*
GetMetadatInfo Query metadata

Query the metadata according to metadata_id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metadataId metadata_id
 @return ApiGetMetadatInfoRequest
*/
func (a *MetadataAPIService) GetMetadatInfo(ctx context.Context, metadataId string) ApiGetMetadatInfoRequest {
	return ApiGetMetadatInfoRequest{
		ApiService: a,
		ctx: ctx,
		metadataId: metadataId,
	}
}

// Execute executes the request
//  @return ModelsExposedMetadata
func (a *MetadataAPIService) GetMetadatInfoExecute(r ApiGetMetadatInfoRequest) (*ModelsExposedMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsExposedMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataAPIService.GetMetadatInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata/{metadata_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"metadata_id"+"}", url.PathEscape(parameterValueToString(r.metadataId, "metadataId")), -1)

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

type ApiListMetadatasRequest struct {
	ctx context.Context
	ApiService MetadataAPI
	authorization *string
	page *string
	limit *string
}

// Bearer openapi_token
func (r ApiListMetadatasRequest) Authorization(authorization string) ApiListMetadatasRequest {
	r.authorization = &authorization
	return r
}

// page
func (r ApiListMetadatasRequest) Page(page string) ApiListMetadatasRequest {
	r.page = &page
	return r
}

// limit
func (r ApiListMetadatasRequest) Limit(limit string) ApiListMetadatasRequest {
	r.limit = &limit
	return r
}

func (r ApiListMetadatasRequest) Execute() (*ModelsExposedMetadataQueryResult, *http.Response, error) {
	return r.ApiService.ListMetadatasExecute(r)
}

/*
ListMetadatas Obtain metadata list

Get the metadata list containing the info of the metadata created in the specified app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMetadatasRequest
*/
func (a *MetadataAPIService) ListMetadatas(ctx context.Context) ApiListMetadatasRequest {
	return ApiListMetadatasRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsExposedMetadataQueryResult
func (a *MetadataAPIService) ListMetadatasExecute(r ApiListMetadatasRequest) (*ModelsExposedMetadataQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsExposedMetadataQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataAPIService.ListMetadatas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
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
