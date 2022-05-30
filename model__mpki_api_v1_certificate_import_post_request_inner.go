/*
Enterprise PKI Manager REST API

 Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.  ## Base URL  The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.  Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.  ## Authentication  API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods: * Header-based API token authentication. * Authentication using a client authentication certificate.  ### API token  To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header: * A Service user token ID (**recommended**).     * An API token bound to a DigiCert ONE administrator.  **Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs.   **Service user token ID** * Service users are API tokens that are not associated with a specific user. * When you create a service user, you assign only the permissions needed for the API integration. * There are two ways to create a new service user:    * 1- Use the Account Manager in DigiCert ONE:      1. Navigate to the DigiCert ONE **Account Manager**.     2. Select **Access** from the left menu.     3. Select **Service User** from the left menu.     4. Select **Create service user** and follow the remaining prompts.    * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.   **Administrator API token**  * Standard users (administrators) can create API tokens in DigiCert ONE.  * API tokens have the same permissions and access scope as the administrator that created them.  * Actions linked to the API token are logged under the administrator's name in event audit logs. * To generate a new API token:     1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.   3. Select **Administrators**.   4. Select the administrator you wish to create an API token for from the list.   5. Scroll down to the **API Tokens** section on the administrator page.    6. Select **Create API token** and follow the remaining prompts.    ### Client authentication certificate  When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.  To generate a client authentication certificate:    1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.    3. Search for and select your Service user.    4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.   5. Enter a **Nickname** and select an **End date** for the certificate.    6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.   7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.   8. Review the permissions you wish to provide to the Service user that the certificate is associated with.      **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.    To use a client authentication certificate:  * Include the certificate in your API request. * In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com` * Omit the `x-api-key` header.  ## Requests  The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.  A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.  ### Method  The Enterprise PKI Manager API uses these standard HTTP methods:  * GET * POST * PUT * DELETE  ### Body and content type  All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.  GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).  ## Responses  Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.  **Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.  ### Headers  Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.  * HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation. * HTTP codes in the **400+** range describe an error.  Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.  **Example error response**  ```JSON {   \"error\": {     \"code\": \"no_access\",     \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"   } } ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// MpkiApiV1CertificateImportPostRequestInner struct for MpkiApiV1CertificateImportPostRequestInner
type MpkiApiV1CertificateImportPostRequestInner struct {
	CertType *string `json:"cert_type,omitempty"`
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	SeatId *string `json:"seat_id,omitempty"`
	Cert **os.File `json:"cert,omitempty"`
	Notes *string `json:"notes,omitempty"`
	GroupName *string `json:"group_name,omitempty"`
	Revocation *MpkiApiV1CertificateImportPostRequestInnerRevocation `json:"revocation,omitempty"`
}

// NewMpkiApiV1CertificateImportPostRequestInner instantiates a new MpkiApiV1CertificateImportPostRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpkiApiV1CertificateImportPostRequestInner() *MpkiApiV1CertificateImportPostRequestInner {
	this := MpkiApiV1CertificateImportPostRequestInner{}
	return &this
}

// NewMpkiApiV1CertificateImportPostRequestInnerWithDefaults instantiates a new MpkiApiV1CertificateImportPostRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpkiApiV1CertificateImportPostRequestInnerWithDefaults() *MpkiApiV1CertificateImportPostRequestInner {
	this := MpkiApiV1CertificateImportPostRequestInner{}
	return &this
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertType() string {
	if o == nil || o.CertType == nil {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertTypeOk() (*string, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetCertType(v string) {
	o.CertType = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetSeatId returns the SeatId field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetSeatId() string {
	if o == nil || o.SeatId == nil {
		var ret string
		return ret
	}
	return *o.SeatId
}

// GetSeatIdOk returns a tuple with the SeatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetSeatIdOk() (*string, bool) {
	if o == nil || o.SeatId == nil {
		return nil, false
	}
	return o.SeatId, true
}

// HasSeatId returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasSeatId() bool {
	if o != nil && o.SeatId != nil {
		return true
	}

	return false
}

// SetSeatId gets a reference to the given string and assigns it to the SeatId field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetSeatId(v string) {
	o.SeatId = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetCert() *os.File {
	if o == nil || o.Cert == nil {
		var ret *os.File
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertOk() (**os.File, bool) {
	if o == nil || o.Cert == nil {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasCert() bool {
	if o != nil && o.Cert != nil {
		return true
	}

	return false
}

// SetCert gets a reference to the given *os.File and assigns it to the Cert field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetCert(v *os.File) {
	o.Cert = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetNotes(v string) {
	o.Notes = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetGroupName(v string) {
	o.GroupName = &v
}

// GetRevocation returns the Revocation field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetRevocation() MpkiApiV1CertificateImportPostRequestInnerRevocation {
	if o == nil || o.Revocation == nil {
		var ret MpkiApiV1CertificateImportPostRequestInnerRevocation
		return ret
	}
	return *o.Revocation
}

// GetRevocationOk returns a tuple with the Revocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) GetRevocationOk() (*MpkiApiV1CertificateImportPostRequestInnerRevocation, bool) {
	if o == nil || o.Revocation == nil {
		return nil, false
	}
	return o.Revocation, true
}

// HasRevocation returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPostRequestInner) HasRevocation() bool {
	if o != nil && o.Revocation != nil {
		return true
	}

	return false
}

// SetRevocation gets a reference to the given MpkiApiV1CertificateImportPostRequestInnerRevocation and assigns it to the Revocation field.
func (o *MpkiApiV1CertificateImportPostRequestInner) SetRevocation(v MpkiApiV1CertificateImportPostRequestInnerRevocation) {
	o.Revocation = &v
}

func (o MpkiApiV1CertificateImportPostRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertType != nil {
		toSerialize["cert_type"] = o.CertType
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.SeatId != nil {
		toSerialize["seat_id"] = o.SeatId
	}
	if o.Cert != nil {
		toSerialize["cert"] = o.Cert
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if o.Revocation != nil {
		toSerialize["revocation"] = o.Revocation
	}
	return json.Marshal(toSerialize)
}

type NullableMpkiApiV1CertificateImportPostRequestInner struct {
	value *MpkiApiV1CertificateImportPostRequestInner
	isSet bool
}

func (v NullableMpkiApiV1CertificateImportPostRequestInner) Get() *MpkiApiV1CertificateImportPostRequestInner {
	return v.value
}

func (v *NullableMpkiApiV1CertificateImportPostRequestInner) Set(val *MpkiApiV1CertificateImportPostRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMpkiApiV1CertificateImportPostRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMpkiApiV1CertificateImportPostRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpkiApiV1CertificateImportPostRequestInner(val *MpkiApiV1CertificateImportPostRequestInner) *NullableMpkiApiV1CertificateImportPostRequestInner {
	return &NullableMpkiApiV1CertificateImportPostRequestInner{value: val, isSet: true}
}

func (v NullableMpkiApiV1CertificateImportPostRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpkiApiV1CertificateImportPostRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


