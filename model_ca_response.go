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

// CaResponse struct for CaResponse
type CaResponse struct {
	// Issuer serial number
	SerialNumber *string `json:"serial_number,omitempty"`
	// Issuer subject dn
	SubjectDn *string `json:"subject_dn,omitempty"`
	// Base-64 encoded issuer certificate
	Certificate *string `json:"certificate,omitempty"`
	// Flag indicating whether its a root certificate
	Root *bool `json:"root,omitempty"`
	Chain []CaResponse `json:"chain,omitempty"`
}

// NewCaResponse instantiates a new CaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaResponse() *CaResponse {
	this := CaResponse{}
	return &this
}

// NewCaResponseWithDefaults instantiates a new CaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaResponseWithDefaults() *CaResponse {
	this := CaResponse{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CaResponse) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CaResponse) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *CaResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSubjectDn returns the SubjectDn field value if set, zero value otherwise.
func (o *CaResponse) GetSubjectDn() string {
	if o == nil || o.SubjectDn == nil {
		var ret string
		return ret
	}
	return *o.SubjectDn
}

// GetSubjectDnOk returns a tuple with the SubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaResponse) GetSubjectDnOk() (*string, bool) {
	if o == nil || o.SubjectDn == nil {
		return nil, false
	}
	return o.SubjectDn, true
}

// HasSubjectDn returns a boolean if a field has been set.
func (o *CaResponse) HasSubjectDn() bool {
	if o != nil && o.SubjectDn != nil {
		return true
	}

	return false
}

// SetSubjectDn gets a reference to the given string and assigns it to the SubjectDn field.
func (o *CaResponse) SetSubjectDn(v string) {
	o.SubjectDn = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CaResponse) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaResponse) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CaResponse) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CaResponse) SetCertificate(v string) {
	o.Certificate = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *CaResponse) GetRoot() bool {
	if o == nil || o.Root == nil {
		var ret bool
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaResponse) GetRootOk() (*bool, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *CaResponse) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given bool and assigns it to the Root field.
func (o *CaResponse) SetRoot(v bool) {
	o.Root = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CaResponse) GetChain() []CaResponse {
	if o == nil || o.Chain == nil {
		var ret []CaResponse
		return ret
	}
	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaResponse) GetChainOk() ([]CaResponse, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CaResponse) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given []CaResponse and assigns it to the Chain field.
func (o *CaResponse) SetChain(v []CaResponse) {
	o.Chain = v
}

func (o CaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.SubjectDn != nil {
		toSerialize["subject_dn"] = o.SubjectDn
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCaResponse struct {
	value *CaResponse
	isSet bool
}

func (v NullableCaResponse) Get() *CaResponse {
	return v.value
}

func (v *NullableCaResponse) Set(val *CaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaResponse(val *CaResponse) *NullableCaResponse {
	return &NullableCaResponse{value: val, isSet: true}
}

func (v NullableCaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


