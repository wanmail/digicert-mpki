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

// EnrollmentListParameters struct for EnrollmentListParameters
type EnrollmentListParameters struct {
	AccountId *string `json:"account_id,omitempty"`
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	ProfileId *string `json:"profile_id,omitempty"`
	// Internal UUID of seat used for enrollment
	SeatId *string `json:"seat_id,omitempty"`
	// Unique identifier of seat used for enrollment
	SeatIdentifier *string `json:"seat_identifier,omitempty"`
	// Name of seat used for enrollment
	SeatName *string `json:"seat_name,omitempty"`
	Email *string `json:"email,omitempty"`
	// List of enrollment statuses to limit search results to
	Statuses []string `json:"statuses,omitempty"`
	// Limit results to enrollments created after provided date
	CreatedAtFrom *string `json:"created_at_from,omitempty"`
	// Limit results to enrollments created before provided date
	CreatedAtTo *string `json:"created_at_to,omitempty"`
	// Limit results to enrollments expiring after provided date
	ExpiresAtFrom *string `json:"expires_at_from,omitempty"`
	// Limit results to enrollments expiring before provided date
	ExpiresAtTo *string `json:"expires_at_to,omitempty"`
}

// NewEnrollmentListParameters instantiates a new EnrollmentListParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentListParameters() *EnrollmentListParameters {
	this := EnrollmentListParameters{}
	return &this
}

// NewEnrollmentListParametersWithDefaults instantiates a new EnrollmentListParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentListParametersWithDefaults() *EnrollmentListParameters {
	this := EnrollmentListParameters{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *EnrollmentListParameters) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *EnrollmentListParameters) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetProfileIdOk() (*string, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *EnrollmentListParameters) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetSeatId returns the SeatId field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetSeatId() string {
	if o == nil || o.SeatId == nil {
		var ret string
		return ret
	}
	return *o.SeatId
}

// GetSeatIdOk returns a tuple with the SeatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetSeatIdOk() (*string, bool) {
	if o == nil || o.SeatId == nil {
		return nil, false
	}
	return o.SeatId, true
}

// HasSeatId returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasSeatId() bool {
	if o != nil && o.SeatId != nil {
		return true
	}

	return false
}

// SetSeatId gets a reference to the given string and assigns it to the SeatId field.
func (o *EnrollmentListParameters) SetSeatId(v string) {
	o.SeatId = &v
}

// GetSeatIdentifier returns the SeatIdentifier field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetSeatIdentifier() string {
	if o == nil || o.SeatIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SeatIdentifier
}

// GetSeatIdentifierOk returns a tuple with the SeatIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetSeatIdentifierOk() (*string, bool) {
	if o == nil || o.SeatIdentifier == nil {
		return nil, false
	}
	return o.SeatIdentifier, true
}

// HasSeatIdentifier returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasSeatIdentifier() bool {
	if o != nil && o.SeatIdentifier != nil {
		return true
	}

	return false
}

// SetSeatIdentifier gets a reference to the given string and assigns it to the SeatIdentifier field.
func (o *EnrollmentListParameters) SetSeatIdentifier(v string) {
	o.SeatIdentifier = &v
}

// GetSeatName returns the SeatName field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetSeatName() string {
	if o == nil || o.SeatName == nil {
		var ret string
		return ret
	}
	return *o.SeatName
}

// GetSeatNameOk returns a tuple with the SeatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetSeatNameOk() (*string, bool) {
	if o == nil || o.SeatName == nil {
		return nil, false
	}
	return o.SeatName, true
}

// HasSeatName returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasSeatName() bool {
	if o != nil && o.SeatName != nil {
		return true
	}

	return false
}

// SetSeatName gets a reference to the given string and assigns it to the SeatName field.
func (o *EnrollmentListParameters) SetSeatName(v string) {
	o.SeatName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EnrollmentListParameters) SetEmail(v string) {
	o.Email = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetStatuses() []string {
	if o == nil || o.Statuses == nil {
		var ret []string
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetStatusesOk() ([]string, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasStatuses() bool {
	if o != nil && o.Statuses != nil {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []string and assigns it to the Statuses field.
func (o *EnrollmentListParameters) SetStatuses(v []string) {
	o.Statuses = v
}

// GetCreatedAtFrom returns the CreatedAtFrom field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetCreatedAtFrom() string {
	if o == nil || o.CreatedAtFrom == nil {
		var ret string
		return ret
	}
	return *o.CreatedAtFrom
}

// GetCreatedAtFromOk returns a tuple with the CreatedAtFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetCreatedAtFromOk() (*string, bool) {
	if o == nil || o.CreatedAtFrom == nil {
		return nil, false
	}
	return o.CreatedAtFrom, true
}

// HasCreatedAtFrom returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasCreatedAtFrom() bool {
	if o != nil && o.CreatedAtFrom != nil {
		return true
	}

	return false
}

// SetCreatedAtFrom gets a reference to the given string and assigns it to the CreatedAtFrom field.
func (o *EnrollmentListParameters) SetCreatedAtFrom(v string) {
	o.CreatedAtFrom = &v
}

// GetCreatedAtTo returns the CreatedAtTo field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetCreatedAtTo() string {
	if o == nil || o.CreatedAtTo == nil {
		var ret string
		return ret
	}
	return *o.CreatedAtTo
}

// GetCreatedAtToOk returns a tuple with the CreatedAtTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetCreatedAtToOk() (*string, bool) {
	if o == nil || o.CreatedAtTo == nil {
		return nil, false
	}
	return o.CreatedAtTo, true
}

// HasCreatedAtTo returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasCreatedAtTo() bool {
	if o != nil && o.CreatedAtTo != nil {
		return true
	}

	return false
}

// SetCreatedAtTo gets a reference to the given string and assigns it to the CreatedAtTo field.
func (o *EnrollmentListParameters) SetCreatedAtTo(v string) {
	o.CreatedAtTo = &v
}

// GetExpiresAtFrom returns the ExpiresAtFrom field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetExpiresAtFrom() string {
	if o == nil || o.ExpiresAtFrom == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAtFrom
}

// GetExpiresAtFromOk returns a tuple with the ExpiresAtFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetExpiresAtFromOk() (*string, bool) {
	if o == nil || o.ExpiresAtFrom == nil {
		return nil, false
	}
	return o.ExpiresAtFrom, true
}

// HasExpiresAtFrom returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasExpiresAtFrom() bool {
	if o != nil && o.ExpiresAtFrom != nil {
		return true
	}

	return false
}

// SetExpiresAtFrom gets a reference to the given string and assigns it to the ExpiresAtFrom field.
func (o *EnrollmentListParameters) SetExpiresAtFrom(v string) {
	o.ExpiresAtFrom = &v
}

// GetExpiresAtTo returns the ExpiresAtTo field value if set, zero value otherwise.
func (o *EnrollmentListParameters) GetExpiresAtTo() string {
	if o == nil || o.ExpiresAtTo == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAtTo
}

// GetExpiresAtToOk returns a tuple with the ExpiresAtTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentListParameters) GetExpiresAtToOk() (*string, bool) {
	if o == nil || o.ExpiresAtTo == nil {
		return nil, false
	}
	return o.ExpiresAtTo, true
}

// HasExpiresAtTo returns a boolean if a field has been set.
func (o *EnrollmentListParameters) HasExpiresAtTo() bool {
	if o != nil && o.ExpiresAtTo != nil {
		return true
	}

	return false
}

// SetExpiresAtTo gets a reference to the given string and assigns it to the ExpiresAtTo field.
func (o *EnrollmentListParameters) SetExpiresAtTo(v string) {
	o.ExpiresAtTo = &v
}

func (o EnrollmentListParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.ProfileId != nil {
		toSerialize["profile_id"] = o.ProfileId
	}
	if o.SeatId != nil {
		toSerialize["seat_id"] = o.SeatId
	}
	if o.SeatIdentifier != nil {
		toSerialize["seat_identifier"] = o.SeatIdentifier
	}
	if o.SeatName != nil {
		toSerialize["seat_name"] = o.SeatName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.CreatedAtFrom != nil {
		toSerialize["created_at_from"] = o.CreatedAtFrom
	}
	if o.CreatedAtTo != nil {
		toSerialize["created_at_to"] = o.CreatedAtTo
	}
	if o.ExpiresAtFrom != nil {
		toSerialize["expires_at_from"] = o.ExpiresAtFrom
	}
	if o.ExpiresAtTo != nil {
		toSerialize["expires_at_to"] = o.ExpiresAtTo
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentListParameters struct {
	value *EnrollmentListParameters
	isSet bool
}

func (v NullableEnrollmentListParameters) Get() *EnrollmentListParameters {
	return v.value
}

func (v *NullableEnrollmentListParameters) Set(val *EnrollmentListParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentListParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentListParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentListParameters(val *EnrollmentListParameters) *NullableEnrollmentListParameters {
	return &NullableEnrollmentListParameters{value: val, isSet: true}
}

func (v NullableEnrollmentListParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentListParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


