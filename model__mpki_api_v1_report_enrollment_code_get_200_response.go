/*
Enterprise PKI Manager REST API

 Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.  ## Base URL  The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.  Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.  ## Authentication  API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods: * Header-based API token authentication. * Authentication using a client authentication certificate.  ### API token  To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header: * A Service user token ID (**recommended**).     * An API token bound to a DigiCert ONE administrator.  **Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs.   **Service user token ID** * Service users are API tokens that are not associated with a specific user. * When you create a service user, you assign only the permissions needed for the API integration. * There are two ways to create a new service user:    * 1- Use the Account Manager in DigiCert ONE:      1. Navigate to the DigiCert ONE **Account Manager**.     2. Select **Access** from the left menu.     3. Select **Service User** from the left menu.     4. Select **Create service user** and follow the remaining prompts.    * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.   **Administrator API token**  * Standard users (administrators) can create API tokens in DigiCert ONE.  * API tokens have the same permissions and access scope as the administrator that created them.  * Actions linked to the API token are logged under the administrator's name in event audit logs. * To generate a new API token:     1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.   3. Select **Administrators**.   4. Select the administrator you wish to create an API token for from the list.   5. Scroll down to the **API Tokens** section on the administrator page.    6. Select **Create API token** and follow the remaining prompts.    ### Client authentication certificate  When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.  To generate a client authentication certificate:    1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.    3. Search for and select your Service user.    4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.   5. Enter a **Nickname** and select an **End date** for the certificate.    6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.   7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.   8. Review the permissions you wish to provide to the Service user that the certificate is associated with.      **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.    To use a client authentication certificate:  * Include the certificate in your API request. * In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com` * Omit the `x-api-key` header.  ## Requests  The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.  A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.  ### Method  The Enterprise PKI Manager API uses these standard HTTP methods:  * GET * POST * PUT * DELETE  ### Body and content type  All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.  GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).  ## Responses  Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.  **Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.  ### Headers  Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.  * HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation. * HTTP codes in the **400+** range describe an error.  Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.  **Example error response**  ```JSON {   \"error\": {     \"code\": \"no_access\",     \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"   } } ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MpkiApiV1ReportEnrollmentCodeGet200Response struct for MpkiApiV1ReportEnrollmentCodeGet200Response
type MpkiApiV1ReportEnrollmentCodeGet200Response struct {
	Total int32 `json:"total"`
	Offset int32 `json:"offset"`
	Limit int32 `json:"limit"`
	Items []ExternalEnrollmentReport `json:"items"`
}

// NewMpkiApiV1ReportEnrollmentCodeGet200Response instantiates a new MpkiApiV1ReportEnrollmentCodeGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpkiApiV1ReportEnrollmentCodeGet200Response(total int32, offset int32, limit int32, items []ExternalEnrollmentReport) *MpkiApiV1ReportEnrollmentCodeGet200Response {
	this := MpkiApiV1ReportEnrollmentCodeGet200Response{}
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	this.Items = items
	return &this
}

// NewMpkiApiV1ReportEnrollmentCodeGet200ResponseWithDefaults instantiates a new MpkiApiV1ReportEnrollmentCodeGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpkiApiV1ReportEnrollmentCodeGet200ResponseWithDefaults() *MpkiApiV1ReportEnrollmentCodeGet200Response {
	this := MpkiApiV1ReportEnrollmentCodeGet200Response{}
	return &this
}

// GetTotal returns the Total field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) SetTotal(v int32) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetItems returns the Items field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetItems() []ExternalEnrollmentReport {
	if o == nil {
		var ret []ExternalEnrollmentReport
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) GetItemsOk() ([]ExternalEnrollmentReport, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *MpkiApiV1ReportEnrollmentCodeGet200Response) SetItems(v []ExternalEnrollmentReport) {
	o.Items = v
}

func (o MpkiApiV1ReportEnrollmentCodeGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableMpkiApiV1ReportEnrollmentCodeGet200Response struct {
	value *MpkiApiV1ReportEnrollmentCodeGet200Response
	isSet bool
}

func (v NullableMpkiApiV1ReportEnrollmentCodeGet200Response) Get() *MpkiApiV1ReportEnrollmentCodeGet200Response {
	return v.value
}

func (v *NullableMpkiApiV1ReportEnrollmentCodeGet200Response) Set(val *MpkiApiV1ReportEnrollmentCodeGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMpkiApiV1ReportEnrollmentCodeGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMpkiApiV1ReportEnrollmentCodeGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpkiApiV1ReportEnrollmentCodeGet200Response(val *MpkiApiV1ReportEnrollmentCodeGet200Response) *NullableMpkiApiV1ReportEnrollmentCodeGet200Response {
	return &NullableMpkiApiV1ReportEnrollmentCodeGet200Response{value: val, isSet: true}
}

func (v NullableMpkiApiV1ReportEnrollmentCodeGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpkiApiV1ReportEnrollmentCodeGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


