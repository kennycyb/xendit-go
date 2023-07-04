/*
Xendit API

Customer API description

API version: 2.0.0
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package customer

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	
	common "github.com/kennycyb/xendit-go/common"
	utils "github.com/kennycyb/xendit-go/utils"
	"strings"
)


type CustomersApi interface {

	/*
	GetCustomer Get Customer

	Retrieves a single customer object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Xendit-generated Customer ID.  Will start with cust-...
	@return ApiGetCustomerRequest
	*/
	GetCustomer(ctx context.Context, id string) ApiGetCustomerRequest

	// GetCustomerExecute executes the request
	//  @return Customer
	GetCustomerExecute(r ApiGetCustomerRequest) (*Customer, *http.Response, error)
}

// CustomersApiService CustomersApi service
type CustomersApiService struct {
	client common.IClient
}

func NewCustomersApi (client common.IClient) CustomersApi {
	return &CustomersApiService{
		client: client,
	}
}

type ApiGetCustomerRequest struct {
	ctx context.Context
	ApiService CustomersApi
	id string
	forUserId *string
	xIdempotencyKey *string
}

// For User ID
func (r ApiGetCustomerRequest) ForUserId(forUserId string) ApiGetCustomerRequest {
	r.forUserId = &forUserId
	return r
}

// Idempotency Key
func (r ApiGetCustomerRequest) XIdempotencyKey(xIdempotencyKey string) ApiGetCustomerRequest {
	r.xIdempotencyKey = &xIdempotencyKey
	return r
}

func (r ApiGetCustomerRequest) Execute() (*Customer, *http.Response, error) {
	return r.ApiService.GetCustomerExecute(r)
}

/*
GetCustomer Get Customer

Retrieves a single customer object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Xendit-generated Customer ID.  Will start with cust-...
 @return ApiGetCustomerRequest
*/
func (a *CustomersApiService) GetCustomer(ctx context.Context, id string) ApiGetCustomerRequest {
	return ApiGetCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomersApiService) GetCustomerExecute(r ApiGetCustomerRequest) (*Customer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "CustomersApiService.GetCustomer")
	if err != nil {
		return localVarReturnValue, nil, common.NewGenericOpenAPIError(nil, err.Error(), nil)
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(utils.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	if r.xIdempotencyKey != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "x-idempotency-key", r.xIdempotencyKey, "")
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := utils.NewGenericOpenAPIError(localVarBody, localVarHTTPResponse.Status, nil)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := utils.NewGenericOpenAPIError(localVarBody, err.Error(), nil)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
