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

// CertificateSearchParameters struct for CertificateSearchParameters
type CertificateSearchParameters struct {
	AccountId *string `json:"account_id,omitempty"`
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	BusinessUnitName *string `json:"business_unit_name,omitempty"`
	ProfileId *string `json:"profile_id,omitempty"`
	SeatId *string `json:"seat_id,omitempty"`
	SeatTypeId *string `json:"seat_type_id,omitempty"`
	// Limit results to certificates issued after specified date value
	ValidFrom *string `json:"valid_from,omitempty"`
	// Limit results to certificates expiring before specified date value
	ValidTo *string `json:"valid_to,omitempty"`
	SerialNumber **os.File `json:"serial_number,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	// Certificate Authority Key Identifier
	Aki *string `json:"aki,omitempty"`
	IssuingCaId *string `json:"issuing_ca_id,omitempty"`
	IssuingCaName *string `json:"issuing_ca_name,omitempty"`
	IssuingCaSerialNumber **os.File `json:"issuing_ca_serial_number,omitempty"`
	Status *string `json:"status,omitempty"`
	// SHA256 thumbprint of the certificate
	Thumbprint **os.File `json:"thumbprint,omitempty"`
}

// NewCertificateSearchParameters instantiates a new CertificateSearchParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateSearchParameters() *CertificateSearchParameters {
	this := CertificateSearchParameters{}
	return &this
}

// NewCertificateSearchParametersWithDefaults instantiates a new CertificateSearchParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateSearchParametersWithDefaults() *CertificateSearchParameters {
	this := CertificateSearchParameters{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CertificateSearchParameters) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *CertificateSearchParameters) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetBusinessUnitName returns the BusinessUnitName field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetBusinessUnitName() string {
	if o == nil || o.BusinessUnitName == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitName
}

// GetBusinessUnitNameOk returns a tuple with the BusinessUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetBusinessUnitNameOk() (*string, bool) {
	if o == nil || o.BusinessUnitName == nil {
		return nil, false
	}
	return o.BusinessUnitName, true
}

// HasBusinessUnitName returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasBusinessUnitName() bool {
	if o != nil && o.BusinessUnitName != nil {
		return true
	}

	return false
}

// SetBusinessUnitName gets a reference to the given string and assigns it to the BusinessUnitName field.
func (o *CertificateSearchParameters) SetBusinessUnitName(v string) {
	o.BusinessUnitName = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetProfileIdOk() (*string, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *CertificateSearchParameters) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetSeatId returns the SeatId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetSeatId() string {
	if o == nil || o.SeatId == nil {
		var ret string
		return ret
	}
	return *o.SeatId
}

// GetSeatIdOk returns a tuple with the SeatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetSeatIdOk() (*string, bool) {
	if o == nil || o.SeatId == nil {
		return nil, false
	}
	return o.SeatId, true
}

// HasSeatId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasSeatId() bool {
	if o != nil && o.SeatId != nil {
		return true
	}

	return false
}

// SetSeatId gets a reference to the given string and assigns it to the SeatId field.
func (o *CertificateSearchParameters) SetSeatId(v string) {
	o.SeatId = &v
}

// GetSeatTypeId returns the SeatTypeId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetSeatTypeId() string {
	if o == nil || o.SeatTypeId == nil {
		var ret string
		return ret
	}
	return *o.SeatTypeId
}

// GetSeatTypeIdOk returns a tuple with the SeatTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetSeatTypeIdOk() (*string, bool) {
	if o == nil || o.SeatTypeId == nil {
		return nil, false
	}
	return o.SeatTypeId, true
}

// HasSeatTypeId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasSeatTypeId() bool {
	if o != nil && o.SeatTypeId != nil {
		return true
	}

	return false
}

// SetSeatTypeId gets a reference to the given string and assigns it to the SeatTypeId field.
func (o *CertificateSearchParameters) SetSeatTypeId(v string) {
	o.SeatTypeId = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetValidFrom() string {
	if o == nil || o.ValidFrom == nil {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetValidFromOk() (*string, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *CertificateSearchParameters) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetValidTo() string {
	if o == nil || o.ValidTo == nil {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetValidToOk() (*string, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *CertificateSearchParameters) SetValidTo(v string) {
	o.ValidTo = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetSerialNumber() *os.File {
	if o == nil || o.SerialNumber == nil {
		var ret *os.File
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetSerialNumberOk() (**os.File, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given *os.File and assigns it to the SerialNumber field.
func (o *CertificateSearchParameters) SetSerialNumber(v *os.File) {
	o.SerialNumber = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *CertificateSearchParameters) SetCommonName(v string) {
	o.CommonName = &v
}

// GetAki returns the Aki field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetAki() string {
	if o == nil || o.Aki == nil {
		var ret string
		return ret
	}
	return *o.Aki
}

// GetAkiOk returns a tuple with the Aki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetAkiOk() (*string, bool) {
	if o == nil || o.Aki == nil {
		return nil, false
	}
	return o.Aki, true
}

// HasAki returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasAki() bool {
	if o != nil && o.Aki != nil {
		return true
	}

	return false
}

// SetAki gets a reference to the given string and assigns it to the Aki field.
func (o *CertificateSearchParameters) SetAki(v string) {
	o.Aki = &v
}

// GetIssuingCaId returns the IssuingCaId field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetIssuingCaId() string {
	if o == nil || o.IssuingCaId == nil {
		var ret string
		return ret
	}
	return *o.IssuingCaId
}

// GetIssuingCaIdOk returns a tuple with the IssuingCaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetIssuingCaIdOk() (*string, bool) {
	if o == nil || o.IssuingCaId == nil {
		return nil, false
	}
	return o.IssuingCaId, true
}

// HasIssuingCaId returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasIssuingCaId() bool {
	if o != nil && o.IssuingCaId != nil {
		return true
	}

	return false
}

// SetIssuingCaId gets a reference to the given string and assigns it to the IssuingCaId field.
func (o *CertificateSearchParameters) SetIssuingCaId(v string) {
	o.IssuingCaId = &v
}

// GetIssuingCaName returns the IssuingCaName field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetIssuingCaName() string {
	if o == nil || o.IssuingCaName == nil {
		var ret string
		return ret
	}
	return *o.IssuingCaName
}

// GetIssuingCaNameOk returns a tuple with the IssuingCaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetIssuingCaNameOk() (*string, bool) {
	if o == nil || o.IssuingCaName == nil {
		return nil, false
	}
	return o.IssuingCaName, true
}

// HasIssuingCaName returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasIssuingCaName() bool {
	if o != nil && o.IssuingCaName != nil {
		return true
	}

	return false
}

// SetIssuingCaName gets a reference to the given string and assigns it to the IssuingCaName field.
func (o *CertificateSearchParameters) SetIssuingCaName(v string) {
	o.IssuingCaName = &v
}

// GetIssuingCaSerialNumber returns the IssuingCaSerialNumber field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetIssuingCaSerialNumber() *os.File {
	if o == nil || o.IssuingCaSerialNumber == nil {
		var ret *os.File
		return ret
	}
	return *o.IssuingCaSerialNumber
}

// GetIssuingCaSerialNumberOk returns a tuple with the IssuingCaSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetIssuingCaSerialNumberOk() (**os.File, bool) {
	if o == nil || o.IssuingCaSerialNumber == nil {
		return nil, false
	}
	return o.IssuingCaSerialNumber, true
}

// HasIssuingCaSerialNumber returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasIssuingCaSerialNumber() bool {
	if o != nil && o.IssuingCaSerialNumber != nil {
		return true
	}

	return false
}

// SetIssuingCaSerialNumber gets a reference to the given *os.File and assigns it to the IssuingCaSerialNumber field.
func (o *CertificateSearchParameters) SetIssuingCaSerialNumber(v *os.File) {
	o.IssuingCaSerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateSearchParameters) SetStatus(v string) {
	o.Status = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *CertificateSearchParameters) GetThumbprint() *os.File {
	if o == nil || o.Thumbprint == nil {
		var ret *os.File
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSearchParameters) GetThumbprintOk() (**os.File, bool) {
	if o == nil || o.Thumbprint == nil {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *CertificateSearchParameters) HasThumbprint() bool {
	if o != nil && o.Thumbprint != nil {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given *os.File and assigns it to the Thumbprint field.
func (o *CertificateSearchParameters) SetThumbprint(v *os.File) {
	o.Thumbprint = &v
}

func (o CertificateSearchParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.BusinessUnitName != nil {
		toSerialize["business_unit_name"] = o.BusinessUnitName
	}
	if o.ProfileId != nil {
		toSerialize["profile_id"] = o.ProfileId
	}
	if o.SeatId != nil {
		toSerialize["seat_id"] = o.SeatId
	}
	if o.SeatTypeId != nil {
		toSerialize["seat_type_id"] = o.SeatTypeId
	}
	if o.ValidFrom != nil {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["valid_to"] = o.ValidTo
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Aki != nil {
		toSerialize["aki"] = o.Aki
	}
	if o.IssuingCaId != nil {
		toSerialize["issuing_ca_id"] = o.IssuingCaId
	}
	if o.IssuingCaName != nil {
		toSerialize["issuing_ca_name"] = o.IssuingCaName
	}
	if o.IssuingCaSerialNumber != nil {
		toSerialize["issuing_ca_serial_number"] = o.IssuingCaSerialNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Thumbprint != nil {
		toSerialize["thumbprint"] = o.Thumbprint
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateSearchParameters struct {
	value *CertificateSearchParameters
	isSet bool
}

func (v NullableCertificateSearchParameters) Get() *CertificateSearchParameters {
	return v.value
}

func (v *NullableCertificateSearchParameters) Set(val *CertificateSearchParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateSearchParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateSearchParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateSearchParameters(val *CertificateSearchParameters) *NullableCertificateSearchParameters {
	return &NullableCertificateSearchParameters{value: val, isSet: true}
}

func (v NullableCertificateSearchParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateSearchParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


