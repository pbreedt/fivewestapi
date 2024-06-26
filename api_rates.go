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
)


// RatesAPIService RatesAPI service
type RatesAPIService service

type ApiReadLatestApiV1RatesLatestGetRequest struct {
	ctx context.Context
	ApiService *RatesAPIService
	base *string
	quote *string
	symbol *string
	exchange *Exchange
	start *float32
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Base(base string) ApiReadLatestApiV1RatesLatestGetRequest {
	r.base = &base
	return r
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Quote(quote string) ApiReadLatestApiV1RatesLatestGetRequest {
	r.quote = &quote
	return r
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Symbol(symbol string) ApiReadLatestApiV1RatesLatestGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Exchange(exchange Exchange) ApiReadLatestApiV1RatesLatestGetRequest {
	r.exchange = &exchange
	return r
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Start(start float32) ApiReadLatestApiV1RatesLatestGetRequest {
	r.start = &start
	return r
}

func (r ApiReadLatestApiV1RatesLatestGetRequest) Execute() (*map[string][]interface{}, *http.Response, error) {
	return r.ApiService.ReadLatestApiV1RatesLatestGetExecute(r)
}

/*
ReadLatestApiV1RatesLatestGet Read Latest

Read the latest rate from the database.

Query Parameters
-----------
- `base` **(optional)**: Base symbol of the pair.

- `quote` **(optional)**: Quote symbol of the pair.

- `symbol` **(optional)**: Symbol for the pair for the selected exchange.

- `exchange` **(optional)**: Name of exchange.

- `start` **(optional)**: Start timestamp (most recent rate after this timestamp will be returned).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReadLatestApiV1RatesLatestGetRequest
*/
func (a *RatesAPIService) ReadLatestApiV1RatesLatestGet(ctx context.Context) ApiReadLatestApiV1RatesLatestGetRequest {
	return ApiReadLatestApiV1RatesLatestGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string][]interface{}
func (a *RatesAPIService) ReadLatestApiV1RatesLatestGetExecute(r ApiReadLatestApiV1RatesLatestGetRequest) (*map[string][]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RatesAPIService.ReadLatestApiV1RatesLatestGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/api/v1/rates/latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.base != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base", r.base, "")
	}
	if r.quote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote", r.quote, "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
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

type ApiStreamLatestApiV1RatesStreamLatestGetRequest struct {
	ctx context.Context
	ApiService *RatesAPIService
	base *string
	quote *string
	exchange *string
}

func (r ApiStreamLatestApiV1RatesStreamLatestGetRequest) Base(base string) ApiStreamLatestApiV1RatesStreamLatestGetRequest {
	r.base = &base
	return r
}

func (r ApiStreamLatestApiV1RatesStreamLatestGetRequest) Quote(quote string) ApiStreamLatestApiV1RatesStreamLatestGetRequest {
	r.quote = &quote
	return r
}

func (r ApiStreamLatestApiV1RatesStreamLatestGetRequest) Exchange(exchange string) ApiStreamLatestApiV1RatesStreamLatestGetRequest {
	r.exchange = &exchange
	return r
}

func (r ApiStreamLatestApiV1RatesStreamLatestGetRequest) Execute() (*MarketRate, *http.Response, error) {
	return r.ApiService.StreamLatestApiV1RatesStreamLatestGetExecute(r)
}

/*
StreamLatestApiV1RatesStreamLatestGet Stream Latest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStreamLatestApiV1RatesStreamLatestGetRequest
*/
func (a *RatesAPIService) StreamLatestApiV1RatesStreamLatestGet(ctx context.Context) ApiStreamLatestApiV1RatesStreamLatestGetRequest {
	return ApiStreamLatestApiV1RatesStreamLatestGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MarketRate
func (a *RatesAPIService) StreamLatestApiV1RatesStreamLatestGetExecute(r ApiStreamLatestApiV1RatesStreamLatestGetRequest) (*MarketRate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketRate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RatesAPIService.StreamLatestApiV1RatesStreamLatestGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/api/v1/rates/stream/latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.base == nil {
		return localVarReturnValue, nil, reportError("base is required and must be specified")
	}
	if r.quote == nil {
		return localVarReturnValue, nil, reportError("quote is required and must be specified")
	}
	if r.exchange == nil {
		return localVarReturnValue, nil, reportError("exchange is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "base", r.base, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "quote", r.quote, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "")
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
