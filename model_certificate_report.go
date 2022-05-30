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

// CertificateReport struct for CertificateReport
type CertificateReport struct {
	// Entity ID in UUID format
	CertificateId *string `json:"certificateId,omitempty"`
	// Current status of Certificate
	Status *string `json:"status,omitempty"`
	// Type of Seat for which Certificate was issued (User, Server, Device, etc)
	SeatType *string `json:"seatType,omitempty"`
	// ID of issuing CA
	Issuer *string `json:"issuer,omitempty"`
	// ID of root CA
	Root *string `json:"root,omitempty"`
	// Certificate revocation date in format YYYY-MM-DDTHH-MM-SS
	RevocationDate *string `json:"revocationDate,omitempty"`
	// Certificate attributes as key-value map
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// NewCertificateReport instantiates a new CertificateReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateReport() *CertificateReport {
	this := CertificateReport{}
	return &this
}

// NewCertificateReportWithDefaults instantiates a new CertificateReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateReportWithDefaults() *CertificateReport {
	this := CertificateReport{}
	return &this
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *CertificateReport) GetCertificateId() string {
	if o == nil || o.CertificateId == nil {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetCertificateIdOk() (*string, bool) {
	if o == nil || o.CertificateId == nil {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *CertificateReport) HasCertificateId() bool {
	if o != nil && o.CertificateId != nil {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *CertificateReport) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateReport) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateReport) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateReport) SetStatus(v string) {
	o.Status = &v
}

// GetSeatType returns the SeatType field value if set, zero value otherwise.
func (o *CertificateReport) GetSeatType() string {
	if o == nil || o.SeatType == nil {
		var ret string
		return ret
	}
	return *o.SeatType
}

// GetSeatTypeOk returns a tuple with the SeatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetSeatTypeOk() (*string, bool) {
	if o == nil || o.SeatType == nil {
		return nil, false
	}
	return o.SeatType, true
}

// HasSeatType returns a boolean if a field has been set.
func (o *CertificateReport) HasSeatType() bool {
	if o != nil && o.SeatType != nil {
		return true
	}

	return false
}

// SetSeatType gets a reference to the given string and assigns it to the SeatType field.
func (o *CertificateReport) SetSeatType(v string) {
	o.SeatType = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CertificateReport) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CertificateReport) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CertificateReport) SetIssuer(v string) {
	o.Issuer = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *CertificateReport) GetRoot() string {
	if o == nil || o.Root == nil {
		var ret string
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetRootOk() (*string, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *CertificateReport) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given string and assigns it to the Root field.
func (o *CertificateReport) SetRoot(v string) {
	o.Root = &v
}

// GetRevocationDate returns the RevocationDate field value if set, zero value otherwise.
func (o *CertificateReport) GetRevocationDate() string {
	if o == nil || o.RevocationDate == nil {
		var ret string
		return ret
	}
	return *o.RevocationDate
}

// GetRevocationDateOk returns a tuple with the RevocationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetRevocationDateOk() (*string, bool) {
	if o == nil || o.RevocationDate == nil {
		return nil, false
	}
	return o.RevocationDate, true
}

// HasRevocationDate returns a boolean if a field has been set.
func (o *CertificateReport) HasRevocationDate() bool {
	if o != nil && o.RevocationDate != nil {
		return true
	}

	return false
}

// SetRevocationDate gets a reference to the given string and assigns it to the RevocationDate field.
func (o *CertificateReport) SetRevocationDate(v string) {
	o.RevocationDate = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CertificateReport) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateReport) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CertificateReport) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *CertificateReport) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o CertificateReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateId != nil {
		toSerialize["certificateId"] = o.CertificateId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SeatType != nil {
		toSerialize["seatType"] = o.SeatType
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	if o.RevocationDate != nil {
		toSerialize["revocationDate"] = o.RevocationDate
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateReport struct {
	value *CertificateReport
	isSet bool
}

func (v NullableCertificateReport) Get() *CertificateReport {
	return v.value
}

func (v *NullableCertificateReport) Set(val *CertificateReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateReport(val *CertificateReport) *NullableCertificateReport {
	return &NullableCertificateReport{value: val, isSet: true}
}

func (v NullableCertificateReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


