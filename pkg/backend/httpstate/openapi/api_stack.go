/*
Pulumi Service API

The Pulumi Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type StackApi interface {

	/*
	CreateStack CreateStack creates a stack with the given cloud and stack name in the scope of the indicated project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return ApiCreateStackRequest
	*/
	CreateStack(ctx context.Context, organization string, project string) ApiCreateStackRequest

	// CreateStackExecute executes the request
	//  @return Stack
	CreateStackExecute(r ApiCreateStackRequest) (*Stack, *http.Response, error)

	/*
	DecryptValue DecryptValue decrypts a ciphertext value in the context of the indicated stack.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param stack
	@return ApiDecryptValueRequest
	*/
	DecryptValue(ctx context.Context, organization string, project string, stack string) ApiDecryptValueRequest

	// DecryptValueExecute executes the request
	//  @return DecryptValueResponse
	DecryptValueExecute(r ApiDecryptValueRequest) (*DecryptValueResponse, *http.Response, error)

	/*
	DeleteStack DeleteStack deletes the indicated stack. If force is true, the stack is deleted even if it contains resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param stack
	@return ApiDeleteStackRequest
	*/
	DeleteStack(ctx context.Context, organization string, project string, stack string) ApiDeleteStackRequest

	// DeleteStackExecute executes the request
	//  @return map[string]interface{}
	DeleteStackExecute(r ApiDeleteStackRequest) (map[string]interface{}, *http.Response, error)

	/*
	DoesProjectExist Returns true if a project with the given name exists, or false otherwise.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return ApiDoesProjectExistRequest
	*/
	DoesProjectExist(ctx context.Context, organization string, project string) ApiDoesProjectExistRequest

	// DoesProjectExistExecute executes the request
	//  @return map[string]interface{}
	DoesProjectExistExecute(r ApiDoesProjectExistRequest) (map[string]interface{}, *http.Response, error)

	/*
	EncryptValue EncryptValue encrypts a plaintext value in the context of the indicated stack.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param stack
	@return ApiEncryptValueRequest
	*/
	EncryptValue(ctx context.Context, organization string, project string, stack string) ApiEncryptValueRequest

	// EncryptValueExecute executes the request
	//  @return EncryptValueResponse
	EncryptValueExecute(r ApiEncryptValueRequest) (*EncryptValueResponse, *http.Response, error)

	/*
	GetStack GetStack retrieves the stack with the given name.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param stack
	@return ApiGetStackRequest
	*/
	GetStack(ctx context.Context, organization string, project string, stack string) ApiGetStackRequest

	// GetStackExecute executes the request
	//  @return Stack
	GetStackExecute(r ApiGetStackRequest) (*Stack, *http.Response, error)

	/*
	UpdateStackTags UpdateStackTags updates the stacks's tags, replacing all existing tags.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param stack
	@return ApiUpdateStackTagsRequest
	*/
	UpdateStackTags(ctx context.Context, organization string, project string, stack string) ApiUpdateStackTagsRequest

	// UpdateStackTagsExecute executes the request
	//  @return map[string]interface{}
	UpdateStackTagsExecute(r ApiUpdateStackTagsRequest) (map[string]interface{}, *http.Response, error)
}

// StackApiService StackApi service
type StackApiService service

type ApiCreateStackRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	createStackRequest *CreateStackRequest
}

// CreateStackRequest defines the request body for creating a new Stack
func (r ApiCreateStackRequest) CreateStackRequest(createStackRequest CreateStackRequest) ApiCreateStackRequest {
	r.createStackRequest = &createStackRequest
	return r
}

func (r ApiCreateStackRequest) Execute() (*Stack, *http.Response, error) {
	return r.ApiService.CreateStackExecute(r)
}

/*
CreateStack CreateStack creates a stack with the given cloud and stack name in the scope of the indicated project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return ApiCreateStackRequest
*/
func (a *StackApiService) CreateStack(ctx context.Context, organization string, project string) ApiCreateStackRequest {
	return ApiCreateStackRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return Stack
func (a *StackApiService) CreateStackExecute(r ApiCreateStackRequest) (*Stack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Stack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.CreateStack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createStackRequest == nil {
		return localVarReturnValue, nil, reportError("createStackRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createStackRequest
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

type ApiDecryptValueRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	stack string
	decryptValueRequest *DecryptValueRequest
}

func (r ApiDecryptValueRequest) DecryptValueRequest(decryptValueRequest DecryptValueRequest) ApiDecryptValueRequest {
	r.decryptValueRequest = &decryptValueRequest
	return r
}

func (r ApiDecryptValueRequest) Execute() (*DecryptValueResponse, *http.Response, error) {
	return r.ApiService.DecryptValueExecute(r)
}

/*
DecryptValue DecryptValue decrypts a ciphertext value in the context of the indicated stack.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param stack
 @return ApiDecryptValueRequest
*/
func (a *StackApiService) DecryptValue(ctx context.Context, organization string, project string, stack string) ApiDecryptValueRequest {
	return ApiDecryptValueRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		stack: stack,
	}
}

// Execute executes the request
//  @return DecryptValueResponse
func (a *StackApiService) DecryptValueExecute(r ApiDecryptValueRequest) (*DecryptValueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DecryptValueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.DecryptValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}/{stack}/decrypt"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stack"+"}", url.PathEscape(parameterToString(r.stack, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.decryptValueRequest == nil {
		return localVarReturnValue, nil, reportError("decryptValueRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.decryptValueRequest
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

type ApiDeleteStackRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	stack string
	force *bool
}

func (r ApiDeleteStackRequest) Force(force bool) ApiDeleteStackRequest {
	r.force = &force
	return r
}

func (r ApiDeleteStackRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteStackExecute(r)
}

/*
DeleteStack DeleteStack deletes the indicated stack. If force is true, the stack is deleted even if it contains resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param stack
 @return ApiDeleteStackRequest
*/
func (a *StackApiService) DeleteStack(ctx context.Context, organization string, project string, stack string) ApiDeleteStackRequest {
	return ApiDeleteStackRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		stack: stack,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *StackApiService) DeleteStackExecute(r ApiDeleteStackRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.DeleteStack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}/{stack}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stack"+"}", url.PathEscape(parameterToString(r.stack, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
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

type ApiDoesProjectExistRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
}

func (r ApiDoesProjectExistRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DoesProjectExistExecute(r)
}

/*
DoesProjectExist Returns true if a project with the given name exists, or false otherwise.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return ApiDoesProjectExistRequest
*/
func (a *StackApiService) DoesProjectExist(ctx context.Context, organization string, project string) ApiDoesProjectExistRequest {
	return ApiDoesProjectExistRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *StackApiService) DoesProjectExistExecute(r ApiDoesProjectExistRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.DoesProjectExist")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stacks/{organization}/{project}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiEncryptValueRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	stack string
	encryptValueRequest *EncryptValueRequest
}

func (r ApiEncryptValueRequest) EncryptValueRequest(encryptValueRequest EncryptValueRequest) ApiEncryptValueRequest {
	r.encryptValueRequest = &encryptValueRequest
	return r
}

func (r ApiEncryptValueRequest) Execute() (*EncryptValueResponse, *http.Response, error) {
	return r.ApiService.EncryptValueExecute(r)
}

/*
EncryptValue EncryptValue encrypts a plaintext value in the context of the indicated stack.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param stack
 @return ApiEncryptValueRequest
*/
func (a *StackApiService) EncryptValue(ctx context.Context, organization string, project string, stack string) ApiEncryptValueRequest {
	return ApiEncryptValueRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		stack: stack,
	}
}

// Execute executes the request
//  @return EncryptValueResponse
func (a *StackApiService) EncryptValueExecute(r ApiEncryptValueRequest) (*EncryptValueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EncryptValueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.EncryptValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}/{stack}/encrypt"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stack"+"}", url.PathEscape(parameterToString(r.stack, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.encryptValueRequest == nil {
		return localVarReturnValue, nil, reportError("encryptValueRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.encryptValueRequest
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

type ApiGetStackRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	stack string
}

func (r ApiGetStackRequest) Execute() (*Stack, *http.Response, error) {
	return r.ApiService.GetStackExecute(r)
}

/*
GetStack GetStack retrieves the stack with the given name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param stack
 @return ApiGetStackRequest
*/
func (a *StackApiService) GetStack(ctx context.Context, organization string, project string, stack string) ApiGetStackRequest {
	return ApiGetStackRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		stack: stack,
	}
}

// Execute executes the request
//  @return Stack
func (a *StackApiService) GetStackExecute(r ApiGetStackRequest) (*Stack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Stack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.GetStack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}/{stack}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stack"+"}", url.PathEscape(parameterToString(r.stack, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiUpdateStackTagsRequest struct {
	ctx context.Context
	ApiService StackApi
	organization string
	project string
	stack string
	requestBody *map[string]string
}

func (r ApiUpdateStackTagsRequest) RequestBody(requestBody map[string]string) ApiUpdateStackTagsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiUpdateStackTagsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateStackTagsExecute(r)
}

/*
UpdateStackTags UpdateStackTags updates the stacks's tags, replacing all existing tags.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param stack
 @return ApiUpdateStackTagsRequest
*/
func (a *StackApiService) UpdateStackTags(ctx context.Context, organization string, project string, stack string) ApiUpdateStackTagsRequest {
	return ApiUpdateStackTagsRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		stack: stack,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *StackApiService) UpdateStackTagsExecute(r ApiUpdateStackTagsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackApiService.UpdateStackTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{organization}/{project}/{stack}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stack"+"}", url.PathEscape(parameterToString(r.stack, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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
