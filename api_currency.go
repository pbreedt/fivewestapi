/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CurrencyAPIService CurrencyAPI service
type CurrencyAPIService service

type ApiGetAllApiV1CurrencyGetRequest struct {
	ctx context.Context
	ApiService *CurrencyAPIService
	page *int32
	pageSize *int32
	symbol *string
	network *string
	fiat *bool
	defaultsOnly *bool
}

func (r ApiGetAllApiV1CurrencyGetRequest) Page(page int32) ApiGetAllApiV1CurrencyGetRequest {
	r.page = &page
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) PageSize(pageSize int32) ApiGetAllApiV1CurrencyGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) Symbol(symbol string) ApiGetAllApiV1CurrencyGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) Network(network string) ApiGetAllApiV1CurrencyGetRequest {
	r.network = &network
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) Fiat(fiat bool) ApiGetAllApiV1CurrencyGetRequest {
	r.fiat = &fiat
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) DefaultsOnly(defaultsOnly bool) ApiGetAllApiV1CurrencyGetRequest {
	r.defaultsOnly = &defaultsOnly
	return r
}

func (r ApiGetAllApiV1CurrencyGetRequest) Execute() ([]CurrencyResponse, *http.Response, error) {
	return r.ApiService.GetAllApiV1CurrencyGetExecute(r)
}

/*
GetAllApiV1CurrencyGet Get All

Get all supported currency information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllApiV1CurrencyGetRequest
*/
func (a *CurrencyAPIService) GetAllApiV1CurrencyGet(ctx context.Context) ApiGetAllApiV1CurrencyGetRequest {
	return ApiGetAllApiV1CurrencyGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []CurrencyResponse
func (a *CurrencyAPIService) GetAllApiV1CurrencyGetExecute(r ApiGetAllApiV1CurrencyGetRequest) ([]CurrencyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CurrencyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrencyAPIService.GetAllApiV1CurrencyGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments/api/v1/currency"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	} else {
		var defaultValue int32 = 1
		r.pageSize = &defaultValue
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "")
	}
	if r.network != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "")
	}
	if r.fiat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fiat", r.fiat, "")
	}
	if r.defaultsOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "defaults_only", r.defaultsOnly, "")
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
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

type ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest struct {
	ctx context.Context
	ApiService *CurrencyAPIService
	page *int32
	pageSize *int32
	symbol *string
	network *string
	fiat *bool
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) Page(page int32) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	r.page = &page
	return r
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) PageSize(pageSize int32) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) Symbol(symbol string) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) Network(network string) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	r.network = &network
	return r
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) Fiat(fiat bool) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	r.fiat = &fiat
	return r
}

func (r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) Execute() ([]Currency, *http.Response, error) {
	return r.ApiService.GetAllSupportedCurrencyInformationApiV1CurrencyGetExecute(r)
}

/*
GetAllSupportedCurrencyInformationApiV1CurrencyGet Get All Supported Currency Information

Gets all supported currencies with pagination

Query Parameters
-----------
- `page` **(optional)**: Page number/window to query for (defaults to 0 for the first page).
- `page_size` **(optional)**: Size of the page returned (defaults to 1, maximum size = 1000 records).
- `symbol` **(optional)**: Name of currency to filter on (e.g. BTC).
- `network` **(optional)**: Name of the blockchain network to filter on.
- `fiat` **(optional)**: Filter on fiat currencies or cryptocurrencies. If null, will return all records.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest
*/
func (a *CurrencyAPIService) GetAllSupportedCurrencyInformationApiV1CurrencyGet(ctx context.Context) ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest {
	return ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Currency
func (a *CurrencyAPIService) GetAllSupportedCurrencyInformationApiV1CurrencyGetExecute(r ApiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest) ([]Currency, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Currency
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrencyAPIService.GetAllSupportedCurrencyInformationApiV1CurrencyGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallet/api/v1/currency"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	} else {
		var defaultValue int32 = 1
		r.pageSize = &defaultValue
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "")
	}
	if r.network != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "")
	}
	if r.fiat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fiat", r.fiat, "")
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
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

type ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest struct {
	ctx context.Context
	ApiService *CurrencyAPIService
	network *string
	fiat *bool
}

func (r ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest) Network(network string) ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest {
	r.network = &network
	return r
}

func (r ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest) Fiat(fiat bool) ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest {
	r.fiat = &fiat
	return r
}

func (r ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetAllSupportedSymbolsApiV1CurrencySymbolsGetExecute(r)
}

/*
GetAllSupportedSymbolsApiV1CurrencySymbolsGet Get All Supported Symbols

Gets all supported symbols

Query Parameters
-----------
- `network` **(optional)**: Name of the blockchain network to filter on.
- `fiat` **(optional)**: Filter on fiat currencies or cryptocurrencies. If null, will return all records.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest
*/
func (a *CurrencyAPIService) GetAllSupportedSymbolsApiV1CurrencySymbolsGet(ctx context.Context) ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest {
	return ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *CurrencyAPIService) GetAllSupportedSymbolsApiV1CurrencySymbolsGetExecute(r ApiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrencyAPIService.GetAllSupportedSymbolsApiV1CurrencySymbolsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallet/api/v1/currency/symbols"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.network != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "")
	}
	if r.fiat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fiat", r.fiat, "")
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
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

type ApiGetApiV1CurrencyIdGetRequest struct {
	ctx context.Context
	ApiService *CurrencyAPIService
	id string
}

func (r ApiGetApiV1CurrencyIdGetRequest) Execute() (*CurrencyResponse, *http.Response, error) {
	return r.ApiService.GetApiV1CurrencyIdGetExecute(r)
}

/*
GetApiV1CurrencyIdGet Get

Get all supported currency information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetApiV1CurrencyIdGetRequest
*/
func (a *CurrencyAPIService) GetApiV1CurrencyIdGet(ctx context.Context, id string) ApiGetApiV1CurrencyIdGetRequest {
	return ApiGetApiV1CurrencyIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CurrencyResponse
func (a *CurrencyAPIService) GetApiV1CurrencyIdGetExecute(r ApiGetApiV1CurrencyIdGetRequest) (*CurrencyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrencyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrencyAPIService.GetApiV1CurrencyIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments/api/v1/currency/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
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

type ApiSymbolsApiV1CurrencySymbolsGetRequest struct {
	ctx context.Context
	ApiService *CurrencyAPIService
	network *string
	fiat *bool
}

func (r ApiSymbolsApiV1CurrencySymbolsGetRequest) Network(network string) ApiSymbolsApiV1CurrencySymbolsGetRequest {
	r.network = &network
	return r
}

func (r ApiSymbolsApiV1CurrencySymbolsGetRequest) Fiat(fiat bool) ApiSymbolsApiV1CurrencySymbolsGetRequest {
	r.fiat = &fiat
	return r
}

func (r ApiSymbolsApiV1CurrencySymbolsGetRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SymbolsApiV1CurrencySymbolsGetExecute(r)
}

/*
SymbolsApiV1CurrencySymbolsGet Symbols

Get all supported symbols.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSymbolsApiV1CurrencySymbolsGetRequest
*/
func (a *CurrencyAPIService) SymbolsApiV1CurrencySymbolsGet(ctx context.Context) ApiSymbolsApiV1CurrencySymbolsGetRequest {
	return ApiSymbolsApiV1CurrencySymbolsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *CurrencyAPIService) SymbolsApiV1CurrencySymbolsGetExecute(r ApiSymbolsApiV1CurrencySymbolsGetRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrencyAPIService.SymbolsApiV1CurrencySymbolsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments/api/v1/currency/symbols"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.network != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "")
	}
	if r.fiat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fiat", r.fiat, "")
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
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
