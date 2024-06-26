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


// RequestForQuoteAPIService RequestForQuoteAPI service
type RequestForQuoteAPIService service

type ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest struct {
	ctx context.Context
	ApiService *RequestForQuoteAPIService
	id string
}

func (r ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest) Execute() (*Rfq, *http.Response, error) {
	return r.ApiService.AcceptQuoteApiV1RfqAcceptQuoteIdPostExecute(r)
}

/*
AcceptQuoteApiV1RfqAcceptQuoteIdPost Accept Quote

Accept a generated quote by ID. Generated quotes may only be accepted within a short window of less than 10 seconds.
If a quote has already been accepted or has expired, a 400 status code is returned.

Body Parameters
-----------
- `id` **(required)**: ID of the quote to be accepted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest
*/
func (a *RequestForQuoteAPIService) AcceptQuoteApiV1RfqAcceptQuoteIdPost(ctx context.Context, id string) ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest {
	return ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Rfq
func (a *RequestForQuoteAPIService) AcceptQuoteApiV1RfqAcceptQuoteIdPostExecute(r ApiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest) (*Rfq, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rfq
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestForQuoteAPIService.AcceptQuoteApiV1RfqAcceptQuoteIdPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/rfq/accept_quote/{id}"
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

type ApiGenerateQuoteApiV1RfqPostRequest struct {
	ctx context.Context
	ApiService *RequestForQuoteAPIService
	rfqRequest *RfqRequest
}

func (r ApiGenerateQuoteApiV1RfqPostRequest) RfqRequest(rfqRequest RfqRequest) ApiGenerateQuoteApiV1RfqPostRequest {
	r.rfqRequest = &rfqRequest
	return r
}

func (r ApiGenerateQuoteApiV1RfqPostRequest) Execute() (*Rfq, *http.Response, error) {
	return r.ApiService.GenerateQuoteApiV1RfqPostExecute(r)
}

/*
GenerateQuoteApiV1RfqPost Generate Quote

Request a quote to trade a in a specific currency pair or exchange two currencies.
The user should specify both a 'from_currency' and a 'to_currency' and only one of 'from_amount' and 'to_amount'.
Note that if both 'from_amount' and 'to_amount' are specified, the requested amount will default to the 'from_amount'.
The 'from_amount' argument should detail an amount to sell - denominated in the 'from_currency' currency.
The 'to_amount' argument should detail and amount to buy - denominated in the 'to_currency' currency.

Body Parameters
-----------
- `from_amount` **(optional)**: An amount to sell - denominated in 'from_currency'
- `from_currency` **(required)**: A currency to trade out of.
- `to_amount` **(optional)**: An amount to buy - denominated in 'to_currency'
- `to_currency` **(required)**: A currency to trade into.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateQuoteApiV1RfqPostRequest
*/
func (a *RequestForQuoteAPIService) GenerateQuoteApiV1RfqPost(ctx context.Context) ApiGenerateQuoteApiV1RfqPostRequest {
	return ApiGenerateQuoteApiV1RfqPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Rfq
func (a *RequestForQuoteAPIService) GenerateQuoteApiV1RfqPostExecute(r ApiGenerateQuoteApiV1RfqPostRequest) (*Rfq, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rfq
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestForQuoteAPIService.GenerateQuoteApiV1RfqPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/rfq"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rfqRequest == nil {
		return localVarReturnValue, nil, reportError("rfqRequest is required and must be specified")
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
	localVarPostBody = r.rfqRequest
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

type ApiGetAllQuoteRequestsApiV1RfqGetRequest struct {
	ctx context.Context
	ApiService *RequestForQuoteAPIService
	quoteStatus *QuoteStatus
	page *int32
	pageSize *int32
}

func (r ApiGetAllQuoteRequestsApiV1RfqGetRequest) QuoteStatus(quoteStatus QuoteStatus) ApiGetAllQuoteRequestsApiV1RfqGetRequest {
	r.quoteStatus = &quoteStatus
	return r
}

func (r ApiGetAllQuoteRequestsApiV1RfqGetRequest) Page(page int32) ApiGetAllQuoteRequestsApiV1RfqGetRequest {
	r.page = &page
	return r
}

func (r ApiGetAllQuoteRequestsApiV1RfqGetRequest) PageSize(pageSize int32) ApiGetAllQuoteRequestsApiV1RfqGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetAllQuoteRequestsApiV1RfqGetRequest) Execute() (*PageRfq, *http.Response, error) {
	return r.ApiService.GetAllQuoteRequestsApiV1RfqGetExecute(r)
}

/*
GetAllQuoteRequestsApiV1RfqGet Get All Quote Requests

Get all RFQ requests with filter parameters and pagination.

Query Parameters
-----------
- `quote_status` **(optional)**: Filter by the status of the quote (either 'accepted' or 'unaccepted'). If unspecified, all quotes will be returned.
- `page` **(optional)**: Page number/window to query for (defaults to 0 for the first page).
- `page_size` **(optional)**: Size of the page returned (defaults to 1, maximum size = 1000 records).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllQuoteRequestsApiV1RfqGetRequest
*/
func (a *RequestForQuoteAPIService) GetAllQuoteRequestsApiV1RfqGet(ctx context.Context) ApiGetAllQuoteRequestsApiV1RfqGetRequest {
	return ApiGetAllQuoteRequestsApiV1RfqGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageRfq
func (a *RequestForQuoteAPIService) GetAllQuoteRequestsApiV1RfqGetExecute(r ApiGetAllQuoteRequestsApiV1RfqGetRequest) (*PageRfq, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageRfq
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestForQuoteAPIService.GetAllQuoteRequestsApiV1RfqGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/rfq"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.quoteStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_status", r.quoteStatus, "")
	}
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

type ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest struct {
	ctx context.Context
	ApiService *RequestForQuoteAPIService
}

func (r ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest) Execute() ([]BidAsk, *http.Response, error) {
	return r.ApiService.GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetExecute(r)
}

/*
GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet Get All Tradable Currency Pairs

Get bid and asks for all active rfq pairs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest
*/
func (a *RequestForQuoteAPIService) GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet(ctx context.Context) ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest {
	return ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []BidAsk
func (a *RequestForQuoteAPIService) GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetExecute(r ApiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest) ([]BidAsk, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []BidAsk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestForQuoteAPIService.GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/rfq/pairs_bid_ask"

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

type ApiGetQuoteRequestApiV1RfqIdGetRequest struct {
	ctx context.Context
	ApiService *RequestForQuoteAPIService
	id string
}

func (r ApiGetQuoteRequestApiV1RfqIdGetRequest) Execute() (*Rfq, *http.Response, error) {
	return r.ApiService.GetQuoteRequestApiV1RfqIdGetExecute(r)
}

/*
GetQuoteRequestApiV1RfqIdGet Get Quote Request

Get a specific quote

Query Parameters
-----------
- `id` **(required)**: ID of the quote to fetch. If a quote with that ID for that user does not exist, a 404 status code is returned

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetQuoteRequestApiV1RfqIdGetRequest
*/
func (a *RequestForQuoteAPIService) GetQuoteRequestApiV1RfqIdGet(ctx context.Context, id string) ApiGetQuoteRequestApiV1RfqIdGetRequest {
	return ApiGetQuoteRequestApiV1RfqIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Rfq
func (a *RequestForQuoteAPIService) GetQuoteRequestApiV1RfqIdGetExecute(r ApiGetQuoteRequestApiV1RfqIdGetRequest) (*Rfq, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rfq
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestForQuoteAPIService.GetQuoteRequestApiV1RfqIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/rfq/{id}"
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
