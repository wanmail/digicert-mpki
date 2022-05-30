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


// EnrollmentsApiService EnrollmentsApi service
type EnrollmentsApiService service

type ApiMpkiApiV1EnrollmentDetailsGetRequest struct {
	ctx context.Context
	ApiService *EnrollmentsApiService
	searchParameters *EnrollmentListParameters
}

func (r ApiMpkiApiV1EnrollmentDetailsGetRequest) SearchParameters(searchParameters EnrollmentListParameters) ApiMpkiApiV1EnrollmentDetailsGetRequest {
	r.searchParameters = &searchParameters
	return r
}

func (r ApiMpkiApiV1EnrollmentDetailsGetRequest) Execute() ([]PublicEnrollmentDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1EnrollmentDetailsGetExecute(r)
}

/*
MpkiApiV1EnrollmentDetailsGet List enrollments

Use this endpoint to list all pending certificate enrollments in your account. 

Results can be filtered using optional search and paging parameters, added as query tags. Append `?{query_parameter}=value` to the endpoint URL, separate multiple parameters with `&`. For example, `/mpki/api/v1/enrollment-details?account_id={account ID}&business_unit_id={business unit ID}`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1EnrollmentDetailsGetRequest
*/
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentDetailsGet(ctx context.Context) ApiMpkiApiV1EnrollmentDetailsGetRequest {
	return ApiMpkiApiV1EnrollmentDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PublicEnrollmentDetails
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentDetailsGetExecute(r ApiMpkiApiV1EnrollmentDetailsGetRequest) ([]PublicEnrollmentDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PublicEnrollmentDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnrollmentsApiService.MpkiApiV1EnrollmentDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/enrollment-details/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchParameters != nil {
		localVarQueryParams.Add("Search parameters", parameterToString(*r.searchParameters, ""))
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

type ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest struct {
	ctx context.Context
	ApiService *EnrollmentsApiService
	identifier string
}

func (r ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest) Execute() (*PublicEnrollmentDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1EnrollmentDetailsIdentifierGetExecute(r)
}

/*
MpkiApiV1EnrollmentDetailsIdentifierGet Get enrollment details by ID

Use this endpoint to get enrollment details for a given enrollment ID. 

The unique enrollment ID is returned when the enrollment is created. The enrollment ID can also be retrieved from the enrollment details page in your account, or a list enrollments GET request.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identifier Enrollment ID
 @return ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest
*/
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentDetailsIdentifierGet(ctx context.Context, identifier string) ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest {
	return ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest{
		ApiService: a,
		ctx: ctx,
		identifier: identifier,
	}
}

// Execute executes the request
//  @return PublicEnrollmentDetails
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentDetailsIdentifierGetExecute(r ApiMpkiApiV1EnrollmentDetailsIdentifierGetRequest) (*PublicEnrollmentDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicEnrollmentDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnrollmentsApiService.MpkiApiV1EnrollmentDetailsIdentifierGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/enrollment-details/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterToString(r.identifier, "")), -1)

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

type ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest struct {
	ctx context.Context
	ApiService *EnrollmentsApiService
	enrollmentCode string
	seatId *string
}

// Seat ID
func (r ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest) SeatId(seatId string) ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest {
	r.seatId = &seatId
	return r
}

func (r ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest) Execute() (*EnrollmentResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1EnrollmentEnrollmentCodeGetExecute(r)
}

/*
MpkiApiV1EnrollmentEnrollmentCodeGet Get enrollment details by enrollment code

Use this endpoint to get enrollment details for a given enrollment code and seat ID. 

`seat_id` is a required query parameter to be included in requests to this endpoint. Append `?seat_id={seat ID}` to the endpoint URL, replace `{seat ID}` with the seat ID to be retrieved. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param enrollmentCode Enrollment code received in enrollment creation response
 @return ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest
*/
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentEnrollmentCodeGet(ctx context.Context, enrollmentCode string) ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest {
	return ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest{
		ApiService: a,
		ctx: ctx,
		enrollmentCode: enrollmentCode,
	}
}

// Execute executes the request
//  @return EnrollmentResponse
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentEnrollmentCodeGetExecute(r ApiMpkiApiV1EnrollmentEnrollmentCodeGetRequest) (*EnrollmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnrollmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnrollmentsApiService.MpkiApiV1EnrollmentEnrollmentCodeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/enrollment/{enrollmentCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"enrollmentCode"+"}", url.PathEscape(parameterToString(r.enrollmentCode, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.seatId == nil {
		return localVarReturnValue, nil, reportError("seatId is required and must be specified")
	}

	localVarQueryParams.Add("seat_id", parameterToString(*r.seatId, ""))
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

type ApiMpkiApiV1EnrollmentPostRequest struct {
	ctx context.Context
	ApiService *EnrollmentsApiService
	mpkiApiV1EnrollmentPostRequest *MpkiApiV1EnrollmentPostRequest
}

func (r ApiMpkiApiV1EnrollmentPostRequest) MpkiApiV1EnrollmentPostRequest(mpkiApiV1EnrollmentPostRequest MpkiApiV1EnrollmentPostRequest) ApiMpkiApiV1EnrollmentPostRequest {
	r.mpkiApiV1EnrollmentPostRequest = &mpkiApiV1EnrollmentPostRequest
	return r
}

func (r ApiMpkiApiV1EnrollmentPostRequest) Execute() (*EnrollmentResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1EnrollmentPostExecute(r)
}

/*
MpkiApiV1EnrollmentPost Create enrollment

Use this endpoint to create a pending certificate enrollment, to be completed and submitted by the certificate end user. 

The following enrollment details are required fields in the request body:

`profile` — Certificate profile ID 

`seat` — Seat ID and seat email address 

`attributes` — Enrollment attributes, including certificate subject DN information, validity period, and CSR (if applicable)

A custom `enrollment_code` can optionally be included in the request body. If no enrollment code is provided in the request, one will be generated and returned in the response body. 

`expiry_date` is an optional request body field. If no expiry_date is specified in the request, and the certificate profile has no set enrollment code validity period, the default value is 10 days. 

**Note:** This endpoint only accepts enrollment requests for certificate profiles with the `ENROLLMENT_CODE` authentication method. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1EnrollmentPostRequest
*/
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentPost(ctx context.Context) ApiMpkiApiV1EnrollmentPostRequest {
	return ApiMpkiApiV1EnrollmentPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EnrollmentResponse
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentPostExecute(r ApiMpkiApiV1EnrollmentPostRequest) (*EnrollmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnrollmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnrollmentsApiService.MpkiApiV1EnrollmentPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/enrollment"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.mpkiApiV1EnrollmentPostRequest
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

type ApiMpkiApiV1EnrollmentRedeemPostRequest struct {
	ctx context.Context
	ApiService *EnrollmentsApiService
	mpkiApiV1EnrollmentRedeemPostRequest *MpkiApiV1EnrollmentRedeemPostRequest
}

func (r ApiMpkiApiV1EnrollmentRedeemPostRequest) MpkiApiV1EnrollmentRedeemPostRequest(mpkiApiV1EnrollmentRedeemPostRequest MpkiApiV1EnrollmentRedeemPostRequest) ApiMpkiApiV1EnrollmentRedeemPostRequest {
	r.mpkiApiV1EnrollmentRedeemPostRequest = &mpkiApiV1EnrollmentRedeemPostRequest
	return r
}

func (r ApiMpkiApiV1EnrollmentRedeemPostRequest) Execute() (*CertificateResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1EnrollmentRedeemPostExecute(r)
}

/*
MpkiApiV1EnrollmentRedeemPost Redeem enrollment

Use this endpoint to redeem a pending certificate enrollment. 

The issued certificate is returned in the response body.

The following enrollment details are required fields in the request body:

`profile` — Certificate profile ID 

`seat` — Seat ID

`enrollment_code` — The enrollment code which was provided by the administrator, or generated automatically, when the enrollment was created.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1EnrollmentRedeemPostRequest
*/
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentRedeemPost(ctx context.Context) ApiMpkiApiV1EnrollmentRedeemPostRequest {
	return ApiMpkiApiV1EnrollmentRedeemPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CertificateResponse
func (a *EnrollmentsApiService) MpkiApiV1EnrollmentRedeemPostExecute(r ApiMpkiApiV1EnrollmentRedeemPostRequest) (*CertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnrollmentsApiService.MpkiApiV1EnrollmentRedeemPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/enrollment/redeem"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.mpkiApiV1EnrollmentRedeemPostRequest
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
