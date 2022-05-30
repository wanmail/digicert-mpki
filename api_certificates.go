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


// CertificatesApiService CertificatesApi service
type CertificatesApiService service

type ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
	thumbprintSha256 *string
}

// Certificate thumbprint
func (r ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest) ThumbprintSha256(thumbprintSha256 string) ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest {
	r.thumbprintSha256 = &thumbprintSha256
	return r
}

func (r ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateEscrowApproveSerialNumberPostExecute(r)
}

/*
MpkiApiV1CertificateEscrowApproveSerialNumberPost Approve escrowed certificate recovery

Use this endpoint to approve a certificate recovery request for a given escrowed certificate serial number.

Certificate recovery approval requires the `RECOVER_EM_CERTIFICATE` user permission.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateEscrowApproveSerialNumberPost(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest {
	return ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
func (a *CertificatesApiService) MpkiApiV1CertificateEscrowApproveSerialNumberPostExecute(r ApiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateEscrowApproveSerialNumberPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/escrow/approve/{serial_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.thumbprintSha256 != nil {
		localVarQueryParams.Add("thumbprint_sha256", parameterToString(*r.thumbprintSha256, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
	includeChain *bool
}

// If true, include CA certificates in response
func (r ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest) IncludeChain(includeChain bool) ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest {
	r.includeChain = &includeChain
	return r
}

func (r ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest) Execute() (*CertificateResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateEscrowRecoverSerialNumberGetExecute(r)
}

/*
MpkiApiV1CertificateEscrowRecoverSerialNumberGet Recover escrowed certificate by serial number

Use this endpoint to recover an escrowed certificate by the given serial number. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateEscrowRecoverSerialNumberGet(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest {
	return ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
//  @return CertificateResponse
func (a *CertificatesApiService) MpkiApiV1CertificateEscrowRecoverSerialNumberGetExecute(r ApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest) (*CertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateEscrowRecoverSerialNumberGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/escrow/recover/{serial_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeChain != nil {
		localVarQueryParams.Add("include_chain", parameterToString(*r.includeChain, ""))
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

type ApiMpkiApiV1CertificateImportPostRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	trapSeatDuplicateError *bool
	mpkiApiV1CertificateImportPostRequestInner *[]MpkiApiV1CertificateImportPostRequestInner
}

// If true, automatically appends the current timestamp to the seat ID value. This feature facilitates duplicate seat ID error handling.  
func (r ApiMpkiApiV1CertificateImportPostRequest) TrapSeatDuplicateError(trapSeatDuplicateError bool) ApiMpkiApiV1CertificateImportPostRequest {
	r.trapSeatDuplicateError = &trapSeatDuplicateError
	return r
}

func (r ApiMpkiApiV1CertificateImportPostRequest) MpkiApiV1CertificateImportPostRequestInner(mpkiApiV1CertificateImportPostRequestInner []MpkiApiV1CertificateImportPostRequestInner) ApiMpkiApiV1CertificateImportPostRequest {
	r.mpkiApiV1CertificateImportPostRequestInner = &mpkiApiV1CertificateImportPostRequestInner
	return r
}

func (r ApiMpkiApiV1CertificateImportPostRequest) Execute() ([]MpkiApiV1CertificateImportPost200ResponseInner, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateImportPostExecute(r)
}

/*
MpkiApiV1CertificateImportPost Import certificate

Use this endpoint to import external certificates into your account. External certificates are issued by an issuing CA which does not originate from your DigiCert ONE account. 

Certificates issued by an external issuing CA which has also been imported to your account are categorized as imported certificates and can be managed accordingly.

Imported certificates can be downloaded, revoked, suspended/resumed, and can make use of OCSP/CRL validation services in your account. 

Certificates issued by an issuing CA which has not been imported to your account are categorized as unmanaged certificates.

Unmanaged certificates can be downloaded from your account, but no certificate lifecycle operations are available for unmanaged seats. 

Certificates imported with `revoked: true` and `reason: certificate_hold` will be imported with certificate_status `SUSPENDED`. 

Revoked certificates receive the certificate_status `REVOKED` when imported — any `reason` for revocation other than `certificate_hold` is replaced with `unspecified`. 

Unmanaged certificates cannot be suspended. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1CertificateImportPostRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateImportPost(ctx context.Context) ApiMpkiApiV1CertificateImportPostRequest {
	return ApiMpkiApiV1CertificateImportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MpkiApiV1CertificateImportPost200ResponseInner
func (a *CertificatesApiService) MpkiApiV1CertificateImportPostExecute(r ApiMpkiApiV1CertificateImportPostRequest) ([]MpkiApiV1CertificateImportPost200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MpkiApiV1CertificateImportPost200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateImportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate-import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.trapSeatDuplicateError != nil {
		localVarQueryParams.Add("trap_seat_duplicate_error", parameterToString(*r.trapSeatDuplicateError, ""))
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
	localVarPostBody = r.mpkiApiV1CertificateImportPostRequestInner
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

type ApiMpkiApiV1CertificatePostRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	mpkiApiV1CertificatePostRequest *MpkiApiV1CertificatePostRequest
}

func (r ApiMpkiApiV1CertificatePostRequest) MpkiApiV1CertificatePostRequest(mpkiApiV1CertificatePostRequest MpkiApiV1CertificatePostRequest) ApiMpkiApiV1CertificatePostRequest {
	r.mpkiApiV1CertificatePostRequest = &mpkiApiV1CertificatePostRequest
	return r
}

func (r ApiMpkiApiV1CertificatePostRequest) Execute() (*CertificateResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificatePostExecute(r)
}

/*
MpkiApiV1CertificatePost Issue certificate

Use this endpoint to issue a new certificate using the given certificate and profile details. 

Instant certificate issuance requires a certificate profile configured with `enrollment_method: rest_api` and `authentication_method: third_party_app`. 

All Subject DN and SAN fields configured in the certificate profile with `REST Request` selected as the field value source and `Required` checked must be included in your request body. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1CertificatePostRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificatePost(ctx context.Context) ApiMpkiApiV1CertificatePostRequest {
	return ApiMpkiApiV1CertificatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CertificateResponse
func (a *CertificatesApiService) MpkiApiV1CertificatePostExecute(r ApiMpkiApiV1CertificatePostRequest) (*CertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate"

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
	localVarPostBody = r.mpkiApiV1CertificatePostRequest
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

type ApiMpkiApiV1CertificateSearchGetRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	searchParameters *CertificateSearchParameters
	pagingParameters *PagingParameters
}

func (r ApiMpkiApiV1CertificateSearchGetRequest) SearchParameters(searchParameters CertificateSearchParameters) ApiMpkiApiV1CertificateSearchGetRequest {
	r.searchParameters = &searchParameters
	return r
}

func (r ApiMpkiApiV1CertificateSearchGetRequest) PagingParameters(pagingParameters PagingParameters) ApiMpkiApiV1CertificateSearchGetRequest {
	r.pagingParameters = &pagingParameters
	return r
}

func (r ApiMpkiApiV1CertificateSearchGetRequest) Execute() (*MpkiApiV1CertificateSearchGet200Response, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateSearchGetExecute(r)
}

/*
MpkiApiV1CertificateSearchGet Search certificates

Use this endpoint to get a detailed list of certificates in your account.

Results can be filtered using optional search and paging parameters appended to the endpoint URL as query tags (e.g. `?limit=10`). 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMpkiApiV1CertificateSearchGetRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateSearchGet(ctx context.Context) ApiMpkiApiV1CertificateSearchGetRequest {
	return ApiMpkiApiV1CertificateSearchGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MpkiApiV1CertificateSearchGet200Response
func (a *CertificatesApiService) MpkiApiV1CertificateSearchGetExecute(r ApiMpkiApiV1CertificateSearchGetRequest) (*MpkiApiV1CertificateSearchGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MpkiApiV1CertificateSearchGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateSearchGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate-search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchParameters != nil {
		localVarQueryParams.Add("Search parameters", parameterToString(*r.searchParameters, ""))
	}
	if r.pagingParameters != nil {
		localVarQueryParams.Add("Paging parameters", parameterToString(*r.pagingParameters, ""))
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

type ApiMpkiApiV1CertificateSerialNumberGetRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
}

func (r ApiMpkiApiV1CertificateSerialNumberGetRequest) Execute() (*PublicCertificateDetails, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateSerialNumberGetExecute(r)
}

/*
MpkiApiV1CertificateSerialNumberGet Get certificate by serial number

Use this endpoint to get the details of a given certificate serial number. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateSerialNumberGetRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberGet(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateSerialNumberGetRequest {
	return ApiMpkiApiV1CertificateSerialNumberGetRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
//  @return PublicCertificateDetails
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberGetExecute(r ApiMpkiApiV1CertificateSerialNumberGetRequest) (*PublicCertificateDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicCertificateDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateSerialNumberGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/{serial_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

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

type ApiMpkiApiV1CertificateSerialNumberRenewPostRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
	thumbprintSha256 *string
	mpkiApiV1CertificateSerialNumberRenewPostRequest *MpkiApiV1CertificateSerialNumberRenewPostRequest
}

// Certificate thumbprint
func (r ApiMpkiApiV1CertificateSerialNumberRenewPostRequest) ThumbprintSha256(thumbprintSha256 string) ApiMpkiApiV1CertificateSerialNumberRenewPostRequest {
	r.thumbprintSha256 = &thumbprintSha256
	return r
}

func (r ApiMpkiApiV1CertificateSerialNumberRenewPostRequest) MpkiApiV1CertificateSerialNumberRenewPostRequest(mpkiApiV1CertificateSerialNumberRenewPostRequest MpkiApiV1CertificateSerialNumberRenewPostRequest) ApiMpkiApiV1CertificateSerialNumberRenewPostRequest {
	r.mpkiApiV1CertificateSerialNumberRenewPostRequest = &mpkiApiV1CertificateSerialNumberRenewPostRequest
	return r
}

func (r ApiMpkiApiV1CertificateSerialNumberRenewPostRequest) Execute() (*CertificateResponse, *http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateSerialNumberRenewPostExecute(r)
}

/*
MpkiApiV1CertificateSerialNumberRenewPost Renew certificate

Use this endpoint to renew the certificate with the given serial number. 

Only certificates from certificate profiles configured with `enrollment_method: rest_api` and `authentication_method: third_party_app` can be renewed using this endpoint. 

Certificates must be within the renewal period configured in the certificate profile (default is 30 days before expiration) to be eligible for renewal. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateSerialNumberRenewPostRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRenewPost(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateSerialNumberRenewPostRequest {
	return ApiMpkiApiV1CertificateSerialNumberRenewPostRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
//  @return CertificateResponse
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRenewPostExecute(r ApiMpkiApiV1CertificateSerialNumberRenewPostRequest) (*CertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateSerialNumberRenewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/{serial_number}/renew"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.thumbprintSha256 != nil {
		localVarQueryParams.Add("thumbprint_sha256", parameterToString(*r.thumbprintSha256, ""))
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
	localVarPostBody = r.mpkiApiV1CertificateSerialNumberRenewPostRequest
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

type ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
}

func (r ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateSerialNumberRevokeDeleteExecute(r)
}

/*
MpkiApiV1CertificateSerialNumberRevokeDelete Resume suspended certificate

Use this endpoint to resume the previously suspended certificate with the given serial number. 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRevokeDelete(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest {
	return ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRevokeDeleteExecute(r ApiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateSerialNumberRevokeDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/{serial_number}/revoke"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMpkiApiV1CertificateSerialNumberRevokePutRequest struct {
	ctx context.Context
	ApiService *CertificatesApiService
	serialNumber string
	thumbprintSha256 *string
	mpkiApiV1CertificateSerialNumberRevokeDeleteRequest *MpkiApiV1CertificateSerialNumberRevokeDeleteRequest
}

// Certificate thumbprint
func (r ApiMpkiApiV1CertificateSerialNumberRevokePutRequest) ThumbprintSha256(thumbprintSha256 string) ApiMpkiApiV1CertificateSerialNumberRevokePutRequest {
	r.thumbprintSha256 = &thumbprintSha256
	return r
}

func (r ApiMpkiApiV1CertificateSerialNumberRevokePutRequest) MpkiApiV1CertificateSerialNumberRevokeDeleteRequest(mpkiApiV1CertificateSerialNumberRevokeDeleteRequest MpkiApiV1CertificateSerialNumberRevokeDeleteRequest) ApiMpkiApiV1CertificateSerialNumberRevokePutRequest {
	r.mpkiApiV1CertificateSerialNumberRevokeDeleteRequest = &mpkiApiV1CertificateSerialNumberRevokeDeleteRequest
	return r
}

func (r ApiMpkiApiV1CertificateSerialNumberRevokePutRequest) Execute() (*http.Response, error) {
	return r.ApiService.MpkiApiV1CertificateSerialNumberRevokePutExecute(r)
}

/*
MpkiApiV1CertificateSerialNumberRevokePut Revoke certificate

Use this endpoint to revoke a certificate by serial number. 

Include the `revocation_reason` in the request body — accepted values are: 

`unspecified`, `key_compromise`, `ca_compromise`, `affiliation_changed`, `superseded`, `cessation_of_operation`, `certificate_hold`, `remove_from_crl`, `privilege_withdrawn`, or `aa_compromise` . 

**Note:** If multiple certificates with the same serial number exist in your account, an error response with the message **"Collision occurred! Please specify additional parameter: thumbprint_sha256"** is returned. Retry the request with the SHA256 thumbprint of the certificate passed as a query tag (append `?thumbprint_sha256={your thumbprint}` to the endpoint URL).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serialNumber Certificate serial number
 @return ApiMpkiApiV1CertificateSerialNumberRevokePutRequest
*/
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRevokePut(ctx context.Context, serialNumber string) ApiMpkiApiV1CertificateSerialNumberRevokePutRequest {
	return ApiMpkiApiV1CertificateSerialNumberRevokePutRequest{
		ApiService: a,
		ctx: ctx,
		serialNumber: serialNumber,
	}
}

// Execute executes the request
func (a *CertificatesApiService) MpkiApiV1CertificateSerialNumberRevokePutExecute(r ApiMpkiApiV1CertificateSerialNumberRevokePutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.MpkiApiV1CertificateSerialNumberRevokePut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mpki/api/v1/certificate/{serial_number}/revoke"
	localVarPath = strings.Replace(localVarPath, "{"+"serial_number"+"}", url.PathEscape(parameterToString(r.serialNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.thumbprintSha256 != nil {
		localVarQueryParams.Add("thumbprint_sha256", parameterToString(*r.thumbprintSha256, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mpkiApiV1CertificateSerialNumberRevokeDeleteRequest
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
