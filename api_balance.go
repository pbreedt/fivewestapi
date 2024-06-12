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


// BalanceAPIService BalanceAPI service
type BalanceAPIService service

type ApiGetAllWalletBalancesApiV1BalanceGetRequest struct {
	ctx context.Context
	ApiService *BalanceAPIService
	page *int32
	pageSize *int32
	symbol *string
}

func (r ApiGetAllWalletBalancesApiV1BalanceGetRequest) Page(page int32) ApiGetAllWalletBalancesApiV1BalanceGetRequest {
	r.page = &page
	return r
}

func (r ApiGetAllWalletBalancesApiV1BalanceGetRequest) PageSize(pageSize int32) ApiGetAllWalletBalancesApiV1BalanceGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetAllWalletBalancesApiV1BalanceGetRequest) Symbol(symbol string) ApiGetAllWalletBalancesApiV1BalanceGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetAllWalletBalancesApiV1BalanceGetRequest) Execute() ([]Balance, *http.Response, error) {
	return r.ApiService.GetAllWalletBalancesApiV1BalanceGetExecute(r)
}

/*
GetAllWalletBalancesApiV1BalanceGet Get All Wallet Balances

Gets wallet balances with pagination.

Query Parameters
-----------
- `page` **(optional)**: Page number/window to query for (defaults to 0 for the first page).
- `page_size` **(optional)**: Size of the page returned (defaults to 1, maximum size = 1000 records).
- `symbol` **(optional)**: Name of currency to filter on (e.g. BTC).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllWalletBalancesApiV1BalanceGetRequest
*/
func (a *BalanceAPIService) GetAllWalletBalancesApiV1BalanceGet(ctx context.Context) ApiGetAllWalletBalancesApiV1BalanceGetRequest {
	return ApiGetAllWalletBalancesApiV1BalanceGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Balance
func (a *BalanceAPIService) GetAllWalletBalancesApiV1BalanceGetExecute(r ApiGetAllWalletBalancesApiV1BalanceGetRequest) ([]Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalanceAPIService.GetAllWalletBalancesApiV1BalanceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallet/api/v1/balance"

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

type ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest struct {
	ctx context.Context
	ApiService *BalanceAPIService
	symbol string
}

func (r ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest) Execute() (*Balance, *http.Response, error) {
	return r.ApiService.GetBalanceForSymbolApiV1BalanceSymbolGetExecute(r)
}

/*
GetBalanceForSymbolApiV1BalanceSymbolGet Get Balance For Symbol

Gets balance for a specific symbol.

If a balance for that symbol is zero or doesn't exist, null will be returned.

Query Parameters
-----------
- `symbol` **(required)**: Name of currency to filter on (e.g. BTC).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param symbol
 @return ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest
*/
func (a *BalanceAPIService) GetBalanceForSymbolApiV1BalanceSymbolGet(ctx context.Context, symbol string) ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest {
	return ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

// Execute executes the request
//  @return Balance
func (a *BalanceAPIService) GetBalanceForSymbolApiV1BalanceSymbolGetExecute(r ApiGetBalanceForSymbolApiV1BalanceSymbolGetRequest) (*Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalanceAPIService.GetBalanceForSymbolApiV1BalanceSymbolGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallet/api/v1/balance/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", url.PathEscape(parameterValueToString(r.symbol, "symbol")), -1)

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