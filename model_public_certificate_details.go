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

// PublicCertificateDetails struct for PublicCertificateDetails
type PublicCertificateDetails struct {
	Profile *NameDto `json:"profile,omitempty"`
	Seat *MpkiApiV1CertificatePostRequestSeat `json:"seat,omitempty"`
	Account *IdDto `json:"account,omitempty"`
	BusinessUnit *NameDto `json:"business_unit,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Status *string `json:"status,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	ValidFrom *string `json:"valid_from,omitempty"`
	ValidTo *string `json:"valid_to,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Revocation *PublicCertificateDetailsRevocation `json:"revocation,omitempty"`
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// NewPublicCertificateDetails instantiates a new PublicCertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicCertificateDetails() *PublicCertificateDetails {
	this := PublicCertificateDetails{}
	return &this
}

// NewPublicCertificateDetailsWithDefaults instantiates a new PublicCertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicCertificateDetailsWithDefaults() *PublicCertificateDetails {
	this := PublicCertificateDetails{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetProfile() NameDto {
	if o == nil || o.Profile == nil {
		var ret NameDto
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetProfileOk() (*NameDto, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NameDto and assigns it to the Profile field.
func (o *PublicCertificateDetails) SetProfile(v NameDto) {
	o.Profile = &v
}

// GetSeat returns the Seat field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetSeat() MpkiApiV1CertificatePostRequestSeat {
	if o == nil || o.Seat == nil {
		var ret MpkiApiV1CertificatePostRequestSeat
		return ret
	}
	return *o.Seat
}

// GetSeatOk returns a tuple with the Seat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetSeatOk() (*MpkiApiV1CertificatePostRequestSeat, bool) {
	if o == nil || o.Seat == nil {
		return nil, false
	}
	return o.Seat, true
}

// HasSeat returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasSeat() bool {
	if o != nil && o.Seat != nil {
		return true
	}

	return false
}

// SetSeat gets a reference to the given MpkiApiV1CertificatePostRequestSeat and assigns it to the Seat field.
func (o *PublicCertificateDetails) SetSeat(v MpkiApiV1CertificatePostRequestSeat) {
	o.Seat = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetAccount() IdDto {
	if o == nil || o.Account == nil {
		var ret IdDto
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetAccountOk() (*IdDto, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IdDto and assigns it to the Account field.
func (o *PublicCertificateDetails) SetAccount(v IdDto) {
	o.Account = &v
}

// GetBusinessUnit returns the BusinessUnit field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetBusinessUnit() NameDto {
	if o == nil || o.BusinessUnit == nil {
		var ret NameDto
		return ret
	}
	return *o.BusinessUnit
}

// GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetBusinessUnitOk() (*NameDto, bool) {
	if o == nil || o.BusinessUnit == nil {
		return nil, false
	}
	return o.BusinessUnit, true
}

// HasBusinessUnit returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasBusinessUnit() bool {
	if o != nil && o.BusinessUnit != nil {
		return true
	}

	return false
}

// SetBusinessUnit gets a reference to the given NameDto and assigns it to the BusinessUnit field.
func (o *PublicCertificateDetails) SetBusinessUnit(v NameDto) {
	o.BusinessUnit = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *PublicCertificateDetails) SetCertificate(v string) {
	o.Certificate = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PublicCertificateDetails) SetCommonName(v string) {
	o.CommonName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PublicCertificateDetails) SetStatus(v string) {
	o.Status = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PublicCertificateDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetValidFrom() string {
	if o == nil || o.ValidFrom == nil {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetValidFromOk() (*string, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *PublicCertificateDetails) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetValidTo() string {
	if o == nil || o.ValidTo == nil {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetValidToOk() (*string, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *PublicCertificateDetails) SetValidTo(v string) {
	o.ValidTo = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *PublicCertificateDetails) SetNotes(v string) {
	o.Notes = &v
}

// GetRevocation returns the Revocation field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetRevocation() PublicCertificateDetailsRevocation {
	if o == nil || o.Revocation == nil {
		var ret PublicCertificateDetailsRevocation
		return ret
	}
	return *o.Revocation
}

// GetRevocationOk returns a tuple with the Revocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetRevocationOk() (*PublicCertificateDetailsRevocation, bool) {
	if o == nil || o.Revocation == nil {
		return nil, false
	}
	return o.Revocation, true
}

// HasRevocation returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasRevocation() bool {
	if o != nil && o.Revocation != nil {
		return true
	}

	return false
}

// SetRevocation gets a reference to the given PublicCertificateDetailsRevocation and assigns it to the Revocation field.
func (o *PublicCertificateDetails) SetRevocation(v PublicCertificateDetailsRevocation) {
	o.Revocation = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *PublicCertificateDetails) GetThumbprint() string {
	if o == nil || o.Thumbprint == nil {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicCertificateDetails) GetThumbprintOk() (*string, bool) {
	if o == nil || o.Thumbprint == nil {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *PublicCertificateDetails) HasThumbprint() bool {
	if o != nil && o.Thumbprint != nil {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *PublicCertificateDetails) SetThumbprint(v string) {
	o.Thumbprint = &v
}

func (o PublicCertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Seat != nil {
		toSerialize["seat"] = o.Seat
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.BusinessUnit != nil {
		toSerialize["business_unit"] = o.BusinessUnit
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.ValidFrom != nil {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["valid_to"] = o.ValidTo
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Revocation != nil {
		toSerialize["revocation"] = o.Revocation
	}
	if o.Thumbprint != nil {
		toSerialize["thumbprint"] = o.Thumbprint
	}
	return json.Marshal(toSerialize)
}

type NullablePublicCertificateDetails struct {
	value *PublicCertificateDetails
	isSet bool
}

func (v NullablePublicCertificateDetails) Get() *PublicCertificateDetails {
	return v.value
}

func (v *NullablePublicCertificateDetails) Set(val *PublicCertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicCertificateDetails(val *PublicCertificateDetails) *NullablePublicCertificateDetails {
	return &NullablePublicCertificateDetails{value: val, isSet: true}
}

func (v NullablePublicCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


