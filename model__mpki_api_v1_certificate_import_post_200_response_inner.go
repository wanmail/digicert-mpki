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

// MpkiApiV1CertificateImportPost200ResponseInner struct for MpkiApiV1CertificateImportPost200ResponseInner
type MpkiApiV1CertificateImportPost200ResponseInner struct {
	CertType *string `json:"cert_type,omitempty"`
	IssuingCa *string `json:"issuing_ca,omitempty"`
	ImportStatus *string `json:"import_status,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Seat *MpkiApiV1CertificateImportPost200ResponseInnerSeat `json:"seat,omitempty"`
	BusinessUnitId *MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId `json:"business_unit_id,omitempty"`
	CertStatus *string `json:"cert_status,omitempty"`
	Revocation *MpkiApiV1CertificateImportPostRequestInnerRevocation `json:"revocation,omitempty"`
}

// NewMpkiApiV1CertificateImportPost200ResponseInner instantiates a new MpkiApiV1CertificateImportPost200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpkiApiV1CertificateImportPost200ResponseInner() *MpkiApiV1CertificateImportPost200ResponseInner {
	this := MpkiApiV1CertificateImportPost200ResponseInner{}
	return &this
}

// NewMpkiApiV1CertificateImportPost200ResponseInnerWithDefaults instantiates a new MpkiApiV1CertificateImportPost200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpkiApiV1CertificateImportPost200ResponseInnerWithDefaults() *MpkiApiV1CertificateImportPost200ResponseInner {
	this := MpkiApiV1CertificateImportPost200ResponseInner{}
	return &this
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertType() string {
	if o == nil || o.CertType == nil {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertTypeOk() (*string, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetCertType(v string) {
	o.CertType = &v
}

// GetIssuingCa returns the IssuingCa field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetIssuingCa() string {
	if o == nil || o.IssuingCa == nil {
		var ret string
		return ret
	}
	return *o.IssuingCa
}

// GetIssuingCaOk returns a tuple with the IssuingCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetIssuingCaOk() (*string, bool) {
	if o == nil || o.IssuingCa == nil {
		return nil, false
	}
	return o.IssuingCa, true
}

// HasIssuingCa returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasIssuingCa() bool {
	if o != nil && o.IssuingCa != nil {
		return true
	}

	return false
}

// SetIssuingCa gets a reference to the given string and assigns it to the IssuingCa field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetIssuingCa(v string) {
	o.IssuingCa = &v
}

// GetImportStatus returns the ImportStatus field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetImportStatus() string {
	if o == nil || o.ImportStatus == nil {
		var ret string
		return ret
	}
	return *o.ImportStatus
}

// GetImportStatusOk returns a tuple with the ImportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetImportStatusOk() (*string, bool) {
	if o == nil || o.ImportStatus == nil {
		return nil, false
	}
	return o.ImportStatus, true
}

// HasImportStatus returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasImportStatus() bool {
	if o != nil && o.ImportStatus != nil {
		return true
	}

	return false
}

// SetImportStatus gets a reference to the given string and assigns it to the ImportStatus field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetImportStatus(v string) {
	o.ImportStatus = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetSeat returns the Seat field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetSeat() MpkiApiV1CertificateImportPost200ResponseInnerSeat {
	if o == nil || o.Seat == nil {
		var ret MpkiApiV1CertificateImportPost200ResponseInnerSeat
		return ret
	}
	return *o.Seat
}

// GetSeatOk returns a tuple with the Seat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetSeatOk() (*MpkiApiV1CertificateImportPost200ResponseInnerSeat, bool) {
	if o == nil || o.Seat == nil {
		return nil, false
	}
	return o.Seat, true
}

// HasSeat returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasSeat() bool {
	if o != nil && o.Seat != nil {
		return true
	}

	return false
}

// SetSeat gets a reference to the given MpkiApiV1CertificateImportPost200ResponseInnerSeat and assigns it to the Seat field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetSeat(v MpkiApiV1CertificateImportPost200ResponseInnerSeat) {
	o.Seat = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetBusinessUnitId() MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId {
	if o == nil || o.BusinessUnitId == nil {
		var ret MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetBusinessUnitIdOk() (*MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId and assigns it to the BusinessUnitId field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetBusinessUnitId(v MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId) {
	o.BusinessUnitId = &v
}

// GetCertStatus returns the CertStatus field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertStatus() string {
	if o == nil || o.CertStatus == nil {
		var ret string
		return ret
	}
	return *o.CertStatus
}

// GetCertStatusOk returns a tuple with the CertStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertStatusOk() (*string, bool) {
	if o == nil || o.CertStatus == nil {
		return nil, false
	}
	return o.CertStatus, true
}

// HasCertStatus returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasCertStatus() bool {
	if o != nil && o.CertStatus != nil {
		return true
	}

	return false
}

// SetCertStatus gets a reference to the given string and assigns it to the CertStatus field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetCertStatus(v string) {
	o.CertStatus = &v
}

// GetRevocation returns the Revocation field value if set, zero value otherwise.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetRevocation() MpkiApiV1CertificateImportPostRequestInnerRevocation {
	if o == nil || o.Revocation == nil {
		var ret MpkiApiV1CertificateImportPostRequestInnerRevocation
		return ret
	}
	return *o.Revocation
}

// GetRevocationOk returns a tuple with the Revocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetRevocationOk() (*MpkiApiV1CertificateImportPostRequestInnerRevocation, bool) {
	if o == nil || o.Revocation == nil {
		return nil, false
	}
	return o.Revocation, true
}

// HasRevocation returns a boolean if a field has been set.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasRevocation() bool {
	if o != nil && o.Revocation != nil {
		return true
	}

	return false
}

// SetRevocation gets a reference to the given MpkiApiV1CertificateImportPostRequestInnerRevocation and assigns it to the Revocation field.
func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetRevocation(v MpkiApiV1CertificateImportPostRequestInnerRevocation) {
	o.Revocation = &v
}

func (o MpkiApiV1CertificateImportPost200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertType != nil {
		toSerialize["cert_type"] = o.CertType
	}
	if o.IssuingCa != nil {
		toSerialize["issuing_ca"] = o.IssuingCa
	}
	if o.ImportStatus != nil {
		toSerialize["import_status"] = o.ImportStatus
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Seat != nil {
		toSerialize["seat"] = o.Seat
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.CertStatus != nil {
		toSerialize["cert_status"] = o.CertStatus
	}
	if o.Revocation != nil {
		toSerialize["revocation"] = o.Revocation
	}
	return json.Marshal(toSerialize)
}

type NullableMpkiApiV1CertificateImportPost200ResponseInner struct {
	value *MpkiApiV1CertificateImportPost200ResponseInner
	isSet bool
}

func (v NullableMpkiApiV1CertificateImportPost200ResponseInner) Get() *MpkiApiV1CertificateImportPost200ResponseInner {
	return v.value
}

func (v *NullableMpkiApiV1CertificateImportPost200ResponseInner) Set(val *MpkiApiV1CertificateImportPost200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMpkiApiV1CertificateImportPost200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMpkiApiV1CertificateImportPost200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpkiApiV1CertificateImportPost200ResponseInner(val *MpkiApiV1CertificateImportPost200ResponseInner) *NullableMpkiApiV1CertificateImportPost200ResponseInner {
	return &NullableMpkiApiV1CertificateImportPost200ResponseInner{value: val, isSet: true}
}

func (v NullableMpkiApiV1CertificateImportPost200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpkiApiV1CertificateImportPost200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


