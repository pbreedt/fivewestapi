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
	"time"
)


// ViewTradesAPIService ViewTradesAPI service
type ViewTradesAPIService service

type ApiGetAllTradesApiV1TradeGetRequest struct {
	ctx context.Context
	ApiService *ViewTradesAPIService
	baseCurrency *string
	quoteCurrency *string
	type_ *TradeType
	mode *Mode
	startTime *time.Time
	endTime *time.Time
	page *int32
	pageSize *int32
}

func (r ApiGetAllTradesApiV1TradeGetRequest) BaseCurrency(baseCurrency string) ApiGetAllTradesApiV1TradeGetRequest {
	r.baseCurrency = &baseCurrency
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) QuoteCurrency(quoteCurrency string) ApiGetAllTradesApiV1TradeGetRequest {
	r.quoteCurrency = &quoteCurrency
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) Type_(type_ TradeType) ApiGetAllTradesApiV1TradeGetRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) Mode(mode Mode) ApiGetAllTradesApiV1TradeGetRequest {
	r.mode = &mode
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) StartTime(startTime time.Time) ApiGetAllTradesApiV1TradeGetRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) EndTime(endTime time.Time) ApiGetAllTradesApiV1TradeGetRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) Page(page int32) ApiGetAllTradesApiV1TradeGetRequest {
	r.page = &page
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) PageSize(pageSize int32) ApiGetAllTradesApiV1TradeGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetAllTradesApiV1TradeGetRequest) Execute() (*PageTrade, *http.Response, error) {
	return r.ApiService.GetAllTradesApiV1TradeGetExecute(r)
}

/*
GetAllTradesApiV1TradeGet Get All Trades

Gets trade history of user.

Query Parameters
-----------
- `base_currency` **(optional)**: Filter trades by a particular base currency.
- `quote_currency` **(optional)**: Filter trades by a particular quote currency.
- `type` **(optional)**: Filter on type of trade - spot, pts (Post-Trade Settlement) or margin (currently unavailable).
- `mode` **(optional)**: Filter by the trade mode - otc (Over-The-Counter), arb (Arbitrage), rfq (Request For Quote).
- `start_time` **(optional)**: Set a starting datetime for the trade query range.
- `end_time` **(optional)**: Set an ending datetime for the trade query range.
- `page` **(optional)**: Page pagination (defaults to 0).
- `page_size` **(optional)**: Page size pagination (defaults to 1).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllTradesApiV1TradeGetRequest
*/
func (a *ViewTradesAPIService) GetAllTradesApiV1TradeGet(ctx context.Context) ApiGetAllTradesApiV1TradeGetRequest {
	return ApiGetAllTradesApiV1TradeGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageTrade
func (a *ViewTradesAPIService) GetAllTradesApiV1TradeGetExecute(r ApiGetAllTradesApiV1TradeGetRequest) (*PageTrade, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageTrade
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewTradesAPIService.GetAllTradesApiV1TradeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/trade"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.baseCurrency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_currency", r.baseCurrency, "")
	}
	if r.quoteCurrency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_currency", r.quoteCurrency, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "")
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime, "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime, "")
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

type ApiGetTradeByIDApiV1TradeIdGetRequest struct {
	ctx context.Context
	ApiService *ViewTradesAPIService
	id string
}

func (r ApiGetTradeByIDApiV1TradeIdGetRequest) Execute() ([]Trade, *http.Response, error) {
	return r.ApiService.GetTradeByIDApiV1TradeIdGetExecute(r)
}

/*
GetTradeByIDApiV1TradeIdGet Get Trade By Id

Gets a specific trade by ID.

Query Parameters
-----------
- `id` **(optional)**: ID of the trade to query for.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetTradeByIDApiV1TradeIdGetRequest
*/
func (a *ViewTradesAPIService) GetTradeByIDApiV1TradeIdGet(ctx context.Context, id string) ApiGetTradeByIDApiV1TradeIdGetRequest {
	return ApiGetTradeByIDApiV1TradeIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []Trade
func (a *ViewTradesAPIService) GetTradeByIDApiV1TradeIdGetExecute(r ApiGetTradeByIDApiV1TradeIdGetRequest) ([]Trade, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Trade
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewTradesAPIService.GetTradeByIDApiV1TradeIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trading/api/v1/trade/{id}"
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