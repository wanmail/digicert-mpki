
/*
 * Enterprise PKI Manager REST API
 *
 *  Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.  ## Base URL  The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.  Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.  ## Authentication  API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods: * Header-based API token authentication. * Authentication using a client authentication certificate.  ### API token  To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header: * A Service user token ID (**recommended**).     * An API token bound to a DigiCert ONE administrator.  **Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs.   **Service user token ID** * Service users are API tokens that are not associated with a specific user. * When you create a service user, you assign only the permissions needed for the API integration. * There are two ways to create a new service user:    * 1- Use the Account Manager in DigiCert ONE:      1. Navigate to the DigiCert ONE **Account Manager**.     2. Select **Access** from the left menu.     3. Select **Service User** from the left menu.     4. Select **Create service user** and follow the remaining prompts.    * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.   **Administrator API token**  * Standard users (administrators) can create API tokens in DigiCert ONE.  * API tokens have the same permissions and access scope as the administrator that created them.  * Actions linked to the API token are logged under the administrator's name in event audit logs. * To generate a new API token:     1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.   3. Select **Administrators**.   4. Select the administrator you wish to create an API token for from the list.   5. Scroll down to the **API Tokens** section on the administrator page.    6. Select **Create API token** and follow the remaining prompts.    ### Client authentication certificate  When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.  To generate a client authentication certificate:    1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.    3. Search for and select your Service user.    4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.   5. Enter a **Nickname** and select an **End date** for the certificate.    6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.   7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.   8. Review the permissions you wish to provide to the Service user that the certificate is associated with.      **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.    To use a client authentication certificate:  * Include the certificate in your API request. * In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com` * Omit the `x-api-key` header.  ## Requests  The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.  A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.  ### Method  The Enterprise PKI Manager API uses these standard HTTP methods:  * GET * POST * PUT * DELETE  ### Body and content type  All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.  GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).  ## Responses  Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.  **Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.  ### Headers  Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.  * HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation. * HTTP codes in the **400+** range describe an error.  Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.  **Example error response**  ```JSON {   \"error\": {     \"code\": \"no_access\",     \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"   } } ``` 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ReportingApiService service
/*
ReportingApiService List audit logs
Use this endpoint to get all available audit log entries in your account.   Results can optionally be filtered using search and paging parameters appended to the endpoint URL (as query tags), e.g. &#x60;?create_date&#x3D;22-02-2022&#x60;.   Filters for audit log data fields can also be applied to results using this format: &#x60;?data[resource_type]&#x3D;enrollment&#x60;. For example, &#x60;?data[thumbprint]&#x3D;12143A4E7XXX2C589EE8AE3C86321B85CA7271328XXX1C9C836935D45XXXBE9X&#x60;. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReportingApiMpkiApiV1AuditLogGetOpts - Optional Parameters:
     * @param "PagingParameters" (optional.Interface of PagingParameters) - 
     * @param "SearchParameters" (optional.Interface of LogSearchParameters) - 
@return InlineResponse2005
*/

type ReportingApiMpkiApiV1AuditLogGetOpts struct {
    PagingParameters optional.Interface
    SearchParameters optional.Interface
}

func (a *ReportingApiService) MpkiApiV1AuditLogGet(ctx context.Context, localVarOptionals *ReportingApiMpkiApiV1AuditLogGetOpts) (InlineResponse2005, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2005
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mpki/api/v1/audit-log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PagingParameters.IsSet() {
		localVarQueryParams.Add("Paging parameters", parameterToString(localVarOptionals.PagingParameters.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchParameters.IsSet() {
		localVarQueryParams.Add("Search parameters", parameterToString(localVarOptionals.SearchParameters.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-Key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2005
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ReportingApiService Get audit log event by ID
Use this endpoint to get the details of a given audit log event ID.  
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Audit log event ID (displayed in the audit event details on UI, and in audit logs list response)
@return AuditLogItem
*/
func (a *ReportingApiService) MpkiApiV1AuditLogIdGet(ctx context.Context, id string) (AuditLogItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue AuditLogItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mpki/api/v1/audit-log/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-Key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v AuditLogItem
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ReportingApiService Get enrollment code report
Use this endpoint to get a report of all enrollment code events.   Results can optionally be filtered using search and paging parameters appended to the endpoint URL (as query tags), e.g. &#x60;?create_date&#x3D;22-02-2022&#x60;.  
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReportingApiMpkiApiV1ReportEnrollmentCodeGetOpts - Optional Parameters:
     * @param "PagingParameters" (optional.Interface of PagingParameters) - 
     * @param "SearchParameters" (optional.Interface of ReportSearchParams) - 
@return InlineResponse2004
*/

type ReportingApiMpkiApiV1ReportEnrollmentCodeGetOpts struct {
    PagingParameters optional.Interface
    SearchParameters optional.Interface
}

func (a *ReportingApiService) MpkiApiV1ReportEnrollmentCodeGet(ctx context.Context, localVarOptionals *ReportingApiMpkiApiV1ReportEnrollmentCodeGetOpts) (InlineResponse2004, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2004
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mpki/api/v1/report/enrollment-code"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PagingParameters.IsSet() {
		localVarQueryParams.Add("Paging parameters", parameterToString(localVarOptionals.PagingParameters.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchParameters.IsSet() {
		localVarQueryParams.Add("Search parameters", parameterToString(localVarOptionals.SearchParameters.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-Key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
