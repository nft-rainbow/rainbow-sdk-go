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
	"os"
)


type FilesApi interface {

	/*
	ListFiles Obtain file list

	Get the file list containing the info of the files uploaded in the specified app

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListFilesRequest
	*/
	ListFiles(ctx context.Context) ApiListFilesRequest

	// ListFilesExecute executes the request
	//  @return ModelsFilesQueryResult
	ListFilesExecute(r ApiListFilesRequest) (*ModelsFilesQueryResult, *http.Response, error)

	/*
	UploadFile Upload file

	Upload a file which can be a video, an image and so on

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadFileRequest
	*/
	UploadFile(ctx context.Context) ApiUploadFileRequest

	// UploadFileExecute executes the request
	//  @return ServicesUploadFilesResponse
	UploadFileExecute(r ApiUploadFileRequest) (*ServicesUploadFilesResponse, *http.Response, error)

	/*
	UploadFileToOss Upload file to OSS

	Upload a file to OSS, which can be a video, an image and so on

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadFileToOssRequest
	*/
	UploadFileToOss(ctx context.Context) ApiUploadFileToOssRequest

	// UploadFileToOssExecute executes the request
	//  @return ServicesUploadFilesResponse
	UploadFileToOssExecute(r ApiUploadFileToOssRequest) (*ServicesUploadFilesResponse, *http.Response, error)

	/*
	UploadFolder Upload folder

	Upload a folder containing the files which can be a video, an image and so on

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadFolderRequest
	*/
	UploadFolder(ctx context.Context) ApiUploadFolderRequest

	// UploadFolderExecute executes the request
	//  @return ServicesUploadFolderResponse
	UploadFolderExecute(r ApiUploadFolderRequest) (*ServicesUploadFolderResponse, *http.Response, error)

	/*
	UploadFolderToOSS Upload folder to oss

	Upload a folder containing the files which can be a video, an image and so on, to oss

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadFolderToOSSRequest
	*/
	UploadFolderToOSS(ctx context.Context) ApiUploadFolderToOSSRequest

	// UploadFolderToOSSExecute executes the request
	//  @return ServicesUploadFolderResponse
	UploadFolderToOSSExecute(r ApiUploadFolderToOSSRequest) (*ServicesUploadFolderResponse, *http.Response, error)
}

// FilesApiService FilesApi service
type FilesApiService service

type ApiListFilesRequest struct {
	ctx context.Context
	ApiService FilesApi
	authorization *string
	page *string
	limit *string
}

// Bearer openapi_token
func (r ApiListFilesRequest) Authorization(authorization string) ApiListFilesRequest {
	r.authorization = &authorization
	return r
}

// page
func (r ApiListFilesRequest) Page(page string) ApiListFilesRequest {
	r.page = &page
	return r
}

// limit
func (r ApiListFilesRequest) Limit(limit string) ApiListFilesRequest {
	r.limit = &limit
	return r
}

func (r ApiListFilesRequest) Execute() (*ModelsFilesQueryResult, *http.Response, error) {
	return r.ApiService.ListFilesExecute(r)
}

/*
ListFiles Obtain file list

Get the file list containing the info of the files uploaded in the specified app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListFilesRequest
*/
func (a *FilesApiService) ListFiles(ctx context.Context) ApiListFilesRequest {
	return ApiListFilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsFilesQueryResult
func (a *FilesApiService) ListFilesExecute(r ApiListFilesRequest) (*ModelsFilesQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsFilesQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.ListFiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/files/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
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

type ApiUploadFileRequest struct {
	ctx context.Context
	ApiService FilesApi
	authorization *string
	file *os.File
}

// Bearer openapi_token
func (r ApiUploadFileRequest) Authorization(authorization string) ApiUploadFileRequest {
	r.authorization = &authorization
	return r
}

// uploaded file
func (r ApiUploadFileRequest) File(file *os.File) ApiUploadFileRequest {
	r.file = file
	return r
}

func (r ApiUploadFileRequest) Execute() (*ServicesUploadFilesResponse, *http.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

/*
UploadFile Upload file

Upload a file which can be a video, an image and so on

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadFileRequest
*/
func (a *FilesApiService) UploadFile(ctx context.Context) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServicesUploadFilesResponse
func (a *FilesApiService) UploadFileExecute(r ApiUploadFileRequest) (*ServicesUploadFilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServicesUploadFilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/files/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
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

type ApiUploadFileToOssRequest struct {
	ctx context.Context
	ApiService FilesApi
	authorization *string
	file *os.File
}

// Bearer openapi_token
func (r ApiUploadFileToOssRequest) Authorization(authorization string) ApiUploadFileToOssRequest {
	r.authorization = &authorization
	return r
}

// uploaded file
func (r ApiUploadFileToOssRequest) File(file *os.File) ApiUploadFileToOssRequest {
	r.file = file
	return r
}

func (r ApiUploadFileToOssRequest) Execute() (*ServicesUploadFilesResponse, *http.Response, error) {
	return r.ApiService.UploadFileToOssExecute(r)
}

/*
UploadFileToOss Upload file to OSS

Upload a file to OSS, which can be a video, an image and so on

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadFileToOssRequest
*/
func (a *FilesApiService) UploadFileToOss(ctx context.Context) ApiUploadFileToOssRequest {
	return ApiUploadFileToOssRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServicesUploadFilesResponse
func (a *FilesApiService) UploadFileToOssExecute(r ApiUploadFileToOssRequest) (*ServicesUploadFilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServicesUploadFilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.UploadFileToOss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/files/oss"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
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

type ApiUploadFolderRequest struct {
	ctx context.Context
	ApiService FilesApi
	authorization *string
	folder *os.File
}

// Bearer openapi_token
func (r ApiUploadFolderRequest) Authorization(authorization string) ApiUploadFolderRequest {
	r.authorization = &authorization
	return r
}

// uploaded folder
func (r ApiUploadFolderRequest) Folder(folder *os.File) ApiUploadFolderRequest {
	r.folder = folder
	return r
}

func (r ApiUploadFolderRequest) Execute() (*ServicesUploadFolderResponse, *http.Response, error) {
	return r.ApiService.UploadFolderExecute(r)
}

/*
UploadFolder Upload folder

Upload a folder containing the files which can be a video, an image and so on

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadFolderRequest
*/
func (a *FilesApiService) UploadFolder(ctx context.Context) ApiUploadFolderRequest {
	return ApiUploadFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServicesUploadFolderResponse
func (a *FilesApiService) UploadFolderExecute(r ApiUploadFolderRequest) (*ServicesUploadFolderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServicesUploadFolderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.UploadFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/files/folder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.folder == nil {
		return localVarReturnValue, nil, reportError("folder is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	var folderLocalVarFormFileName string
	var folderLocalVarFileName     string
	var folderLocalVarFileBytes    []byte

	folderLocalVarFormFileName = "folder"


	folderLocalVarFile := r.folder

	if folderLocalVarFile != nil {
		fbs, _ := io.ReadAll(folderLocalVarFile)

		folderLocalVarFileBytes = fbs
		folderLocalVarFileName = folderLocalVarFile.Name()
		folderLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: folderLocalVarFileBytes, fileName: folderLocalVarFileName, formFileName: folderLocalVarFormFileName})
	}
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

type ApiUploadFolderToOSSRequest struct {
	ctx context.Context
	ApiService FilesApi
	authorization *string
	folder *os.File
}

// Bearer openapi_token
func (r ApiUploadFolderToOSSRequest) Authorization(authorization string) ApiUploadFolderToOSSRequest {
	r.authorization = &authorization
	return r
}

// uploaded folder
func (r ApiUploadFolderToOSSRequest) Folder(folder *os.File) ApiUploadFolderToOSSRequest {
	r.folder = folder
	return r
}

func (r ApiUploadFolderToOSSRequest) Execute() (*ServicesUploadFolderResponse, *http.Response, error) {
	return r.ApiService.UploadFolderToOSSExecute(r)
}

/*
UploadFolderToOSS Upload folder to oss

Upload a folder containing the files which can be a video, an image and so on, to oss

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadFolderToOSSRequest
*/
func (a *FilesApiService) UploadFolderToOSS(ctx context.Context) ApiUploadFolderToOSSRequest {
	return ApiUploadFolderToOSSRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServicesUploadFolderResponse
func (a *FilesApiService) UploadFolderToOSSExecute(r ApiUploadFolderToOSSRequest) (*ServicesUploadFolderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServicesUploadFolderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.UploadFolderToOSS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/files/folder/oss"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.folder == nil {
		return localVarReturnValue, nil, reportError("folder is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	var folderLocalVarFormFileName string
	var folderLocalVarFileName     string
	var folderLocalVarFileBytes    []byte

	folderLocalVarFormFileName = "folder"


	folderLocalVarFile := r.folder

	if folderLocalVarFile != nil {
		fbs, _ := io.ReadAll(folderLocalVarFile)

		folderLocalVarFileBytes = fbs
		folderLocalVarFileName = folderLocalVarFile.Name()
		folderLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: folderLocalVarFileBytes, fileName: folderLocalVarFileName, formFileName: folderLocalVarFormFileName})
	}
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
