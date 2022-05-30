/*
Enterprise PKI Manager REST API

 Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.  ## Base URL  The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.  Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.  ## Authentication  API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods: * Header-based API token authentication. * Authentication using a client authentication certificate.  ### API token  To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header: * A Service user token ID (**recommended**).     * An API token bound to a DigiCert ONE administrator.  **Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs.   **Service user token ID** * Service users are API tokens that are not associated with a specific user. * When you create a service user, you assign only the permissions needed for the API integration. * There are two ways to create a new service user:    * 1- Use the Account Manager in DigiCert ONE:      1. Navigate to the DigiCert ONE **Account Manager**.     2. Select **Access** from the left menu.     3. Select **Service User** from the left menu.     4. Select **Create service user** and follow the remaining prompts.    * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.   **Administrator API token**  * Standard users (administrators) can create API tokens in DigiCert ONE.  * API tokens have the same permissions and access scope as the administrator that created them.  * Actions linked to the API token are logged under the administrator's name in event audit logs. * To generate a new API token:     1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.   3. Select **Administrators**.   4. Select the administrator you wish to create an API token for from the list.   5. Scroll down to the **API Tokens** section on the administrator page.    6. Select **Create API token** and follow the remaining prompts.    ### Client authentication certificate  When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.  To generate a client authentication certificate:    1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.    3. Search for and select your Service user.    4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.   5. Enter a **Nickname** and select an **End date** for the certificate.    6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.   7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.   8. Review the permissions you wish to provide to the Service user that the certificate is associated with.      **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.    To use a client authentication certificate:  * Include the certificate in your API request. * In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com` * Omit the `x-api-key` header.  ## Requests  The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.  A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.  ### Method  The Enterprise PKI Manager API uses these standard HTTP methods:  * GET * POST * PUT * DELETE  ### Body and content type  All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.  GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).  ## Responses  Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.  **Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.  ### Headers  Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.  * HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation. * HTTP codes in the **400+** range describe an error.  Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.  **Example error response**  ```JSON {   \"error\": {     \"code\": \"no_access\",     \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"   } } ``` 

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


// SeatsApiService SeatsApi service
type SeatsApiService service

type ApiMpkiApiV1SeatPostRequest struct {
	ctx context.Context
	ApiService *SeatsApiService
	mpkiApiV1SeatPostRequest *MpkiApiV1SeatPostRequest
}

func (r ApiMpkiApiV1SeatPostRequest) MpkiApiV1SeatPostRequest(mpkiApiV1SeatPostRequest MpkiApiV1SeatPostRequest) ApiMpkiApiV1SeatPostRequest {
	r.mpkiApiV1SeatPostRequest = &mpkiApiV1SeatPostRequest
	return r
}

func (r ApiMpkiApiV1SeatPostRequest) Execute() (*SeatDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1SeatPostExecute(r)
}

/*
MpkiApiV1SeatPost Create seat

Use this endpoint to create a new seat. Each seat decreases the number of allocated seats for the given seat type. This endpoint does not restore previously deleted seats—reusing deleted seat IDs results in the creation of a new seat with the previously deleted seat ID. 

Each unique seat ID represents only one active seat at any given time — duplicate seat IDs are not permitted. 

`seat_id` is the identifier for each seat — this must be unique within each business unit and for each seat type.

The `seat_type` parameter is optional. If no `seat_type` is specified, this defaults to `USER_SEAT`. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1SeatPostRequest
*/
func (a *SeatsApiService) MpkiApiV1SeatPost(ctx context.Context) ApiMpkiApiV1SeatPostRequest {
	return ApiMpkiApiV1SeatPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SeatDetails
func (a *SeatsApiService) MpkiApiV1SeatPostExecute(r ApiMpkiApiV1SeatPostRequest) (*SeatDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SeatDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SeatsApiService.MpkiApiV1SeatPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/seat"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=utf-8"}

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
	localVarPostBody = r.mpkiApiV1SeatPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
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

type ApiMpkiApiV1SeatSeatIdDeleteRequest struct {
	ctx context.Context
	ApiService *SeatsApiService
	seatId string
	seatTypeId *string
	businessUnitId *string
}

// Seat Type ID
func (r ApiMpkiApiV1SeatSeatIdDeleteRequest) SeatTypeId(seatTypeId string) ApiMpkiApiV1SeatSeatIdDeleteRequest {
	r.seatTypeId = &seatTypeId
	return r
}

// Business unit ID
func (r ApiMpkiApiV1SeatSeatIdDeleteRequest) BusinessUnitId(businessUnitId string) ApiMpkiApiV1SeatSeatIdDeleteRequest {
	r.businessUnitId = &businessUnitId
	return r
}

func (r ApiMpkiApiV1SeatSeatIdDeleteRequest) Execute() (*SeatDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1SeatSeatIdDeleteExecute(r)
}

/*
MpkiApiV1SeatSeatIdDelete Delete seat

Use this endpoint to delete a given seat ID. Deleting the seat ID restores an available seat of the same type in your account.
As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit.
The `seat_type` parameter is required. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seatId Seat ID
 @return ApiMpkiApiV1SeatSeatIdDeleteRequest
*/
func (a *SeatsApiService) MpkiApiV1SeatSeatIdDelete(ctx context.Context, seatId string) ApiMpkiApiV1SeatSeatIdDeleteRequest {
	return ApiMpkiApiV1SeatSeatIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		seatId: seatId,
	}
}

// Execute executes the request
//  @return SeatDetails
func (a *SeatsApiService) MpkiApiV1SeatSeatIdDeleteExecute(r ApiMpkiApiV1SeatSeatIdDeleteRequest) (*SeatDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SeatDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SeatsApiService.MpkiApiV1SeatSeatIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/seat/{seat_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"seat_id"+"}", url.PathEscape(parameterToString(r.seatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.seatTypeId == nil {
		return localVarReturnValue, nil, reportError("seatTypeId is required and must be specified")
	}

	if r.businessUnitId != nil {
		localVarQueryParams.Add("business_unit_id", parameterToString(*r.businessUnitId, ""))
	}
	localVarQueryParams.Add("seat_type_id", parameterToString(*r.seatTypeId, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
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

type ApiMpkiApiV1SeatSeatIdGetRequest struct {
	ctx context.Context
	ApiService *SeatsApiService
	seatId string
	seatTypeId *string
	businessUnitId *string
}

// Seat Type ID
func (r ApiMpkiApiV1SeatSeatIdGetRequest) SeatTypeId(seatTypeId string) ApiMpkiApiV1SeatSeatIdGetRequest {
	r.seatTypeId = &seatTypeId
	return r
}

// Business unit ID.  As identical seat IDs may exist across various business units of the same account, the &#x60;business_unit_id&#x60; for the given seat must be specified as a query parameter if the account contains more than one business unit. 
func (r ApiMpkiApiV1SeatSeatIdGetRequest) BusinessUnitId(businessUnitId string) ApiMpkiApiV1SeatSeatIdGetRequest {
	r.businessUnitId = &businessUnitId
	return r
}

func (r ApiMpkiApiV1SeatSeatIdGetRequest) Execute() (*MpkiApiV1SeatSeatIdDelete200Response, *http.Response, error) {
	return r.ApiService.MpkiApiV1SeatSeatIdGetExecute(r)
}

/*
MpkiApiV1SeatSeatIdGet Get seat by ID

Use this endpoint to retrieve the details of a given seat ID. Seat IDs are unique for each `business_unit_id` and `seat_type`.
As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit.
The `seat_type` parameter is required. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seatId Seat ID
 @return ApiMpkiApiV1SeatSeatIdGetRequest
*/
func (a *SeatsApiService) MpkiApiV1SeatSeatIdGet(ctx context.Context, seatId string) ApiMpkiApiV1SeatSeatIdGetRequest {
	return ApiMpkiApiV1SeatSeatIdGetRequest{
		ApiService: a,
		ctx: ctx,
		seatId: seatId,
	}
}

// Execute executes the request
//  @return MpkiApiV1SeatSeatIdDelete200Response
func (a *SeatsApiService) MpkiApiV1SeatSeatIdGetExecute(r ApiMpkiApiV1SeatSeatIdGetRequest) (*MpkiApiV1SeatSeatIdDelete200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MpkiApiV1SeatSeatIdDelete200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SeatsApiService.MpkiApiV1SeatSeatIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/seat/{seat_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"seat_id"+"}", url.PathEscape(parameterToString(r.seatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.seatTypeId == nil {
		return localVarReturnValue, nil, reportError("seatTypeId is required and must be specified")
	}

	if r.businessUnitId != nil {
		localVarQueryParams.Add("business_unit_id", parameterToString(*r.businessUnitId, ""))
	}
	localVarQueryParams.Add("seat_type_id", parameterToString(*r.seatTypeId, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
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

type ApiMpkiApiV1SeatSeatIdPutRequest struct {
	ctx context.Context
	ApiService *SeatsApiService
	seatId string
	mpkiApiV1SeatSeatIdDeleteRequest *MpkiApiV1SeatSeatIdDeleteRequest
}

func (r ApiMpkiApiV1SeatSeatIdPutRequest) MpkiApiV1SeatSeatIdDeleteRequest(mpkiApiV1SeatSeatIdDeleteRequest MpkiApiV1SeatSeatIdDeleteRequest) ApiMpkiApiV1SeatSeatIdPutRequest {
	r.mpkiApiV1SeatSeatIdDeleteRequest = &mpkiApiV1SeatSeatIdDeleteRequest
	return r
}

func (r ApiMpkiApiV1SeatSeatIdPutRequest) Execute() (*SeatDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1SeatSeatIdPutExecute(r)
}

/*
MpkiApiV1SeatSeatIdPut Update seat

Use this endpoint to update the details for a given seat ID. Include the seat details to be updated in the request body.
As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit.
The `seat_type` parameter is optional. If no `seat_type` is specified, this defaults to `USER_SEAT`. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seatId Seat ID
 @return ApiMpkiApiV1SeatSeatIdPutRequest
*/
func (a *SeatsApiService) MpkiApiV1SeatSeatIdPut(ctx context.Context, seatId string) ApiMpkiApiV1SeatSeatIdPutRequest {
	return ApiMpkiApiV1SeatSeatIdPutRequest{
		ApiService: a,
		ctx: ctx,
		seatId: seatId,
	}
}

// Execute executes the request
//  @return SeatDetails
func (a *SeatsApiService) MpkiApiV1SeatSeatIdPutExecute(r ApiMpkiApiV1SeatSeatIdPutRequest) (*SeatDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SeatDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SeatsApiService.MpkiApiV1SeatSeatIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/seat/{seat_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"seat_id"+"}", url.PathEscape(parameterToString(r.seatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=utf-8"}

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
	localVarPostBody = r.mpkiApiV1SeatSeatIdDeleteRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
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

type ApiMpkiApiV1SeatTypesGetRequest struct {
	ctx context.Context
	ApiService *SeatsApiService
}

func (r ApiMpkiApiV1SeatTypesGetRequest) Execute() ([]MpkiApiV1SeatTypesGet200ResponseInner, *http.Response, error) {
	return r.ApiService.MpkiApiV1SeatTypesGetExecute(r)
}

/*
MpkiApiV1SeatTypesGet Get available seat types

Use this endpoint to fetch details about available seat types for the account. Returns an array containing seat type details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1SeatTypesGetRequest
*/
func (a *SeatsApiService) MpkiApiV1SeatTypesGet(ctx context.Context) ApiMpkiApiV1SeatTypesGetRequest {
	return ApiMpkiApiV1SeatTypesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MpkiApiV1SeatTypesGet200ResponseInner
func (a *SeatsApiService) MpkiApiV1SeatTypesGetExecute(r ApiMpkiApiV1SeatTypesGetRequest) ([]MpkiApiV1SeatTypesGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MpkiApiV1SeatTypesGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SeatsApiService.MpkiApiV1SeatTypesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/seat-types"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
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
