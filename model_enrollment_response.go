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

// EnrollmentResponse struct for EnrollmentResponse
type EnrollmentResponse struct {
	// Seat id
	SeatId *string `json:"seat_id,omitempty"`
	// Profile id
	ProfileId *string `json:"profile_id,omitempty"`
	// Profile name
	ProfileName *string `json:"profile_name,omitempty"`
	// Enrollment code to be supplied to the certificate end user
	EnrollmentCode *string `json:"enrollment_code,omitempty"`
	// Creation date and time
	CreationDateTime *string `json:"creation_date_time,omitempty"`
	// Expiration date and time
	ExpiryDateTime *string `json:"expiry_date_time,omitempty"`
	// Enrollment status
	Status *string `json:"status,omitempty"`
	// Number of erroneous attempts to complete enrollment (such as with PUT requests sent to /mpki/api/v1/enrollment/{enrollmentCode})
	NumberOfBadAttempts *float32 `json:"number_of_bad_attempts,omitempty"`
}

// NewEnrollmentResponse instantiates a new EnrollmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentResponse() *EnrollmentResponse {
	this := EnrollmentResponse{}
	return &this
}

// NewEnrollmentResponseWithDefaults instantiates a new EnrollmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentResponseWithDefaults() *EnrollmentResponse {
	this := EnrollmentResponse{}
	return &this
}

// GetSeatId returns the SeatId field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetSeatId() string {
	if o == nil || o.SeatId == nil {
		var ret string
		return ret
	}
	return *o.SeatId
}

// GetSeatIdOk returns a tuple with the SeatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetSeatIdOk() (*string, bool) {
	if o == nil || o.SeatId == nil {
		return nil, false
	}
	return o.SeatId, true
}

// HasSeatId returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasSeatId() bool {
	if o != nil && o.SeatId != nil {
		return true
	}

	return false
}

// SetSeatId gets a reference to the given string and assigns it to the SeatId field.
func (o *EnrollmentResponse) SetSeatId(v string) {
	o.SeatId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetProfileIdOk() (*string, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *EnrollmentResponse) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *EnrollmentResponse) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetEnrollmentCode returns the EnrollmentCode field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetEnrollmentCode() string {
	if o == nil || o.EnrollmentCode == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentCode
}

// GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetEnrollmentCodeOk() (*string, bool) {
	if o == nil || o.EnrollmentCode == nil {
		return nil, false
	}
	return o.EnrollmentCode, true
}

// HasEnrollmentCode returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasEnrollmentCode() bool {
	if o != nil && o.EnrollmentCode != nil {
		return true
	}

	return false
}

// SetEnrollmentCode gets a reference to the given string and assigns it to the EnrollmentCode field.
func (o *EnrollmentResponse) SetEnrollmentCode(v string) {
	o.EnrollmentCode = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetCreationDateTime() string {
	if o == nil || o.CreationDateTime == nil {
		var ret string
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetCreationDateTimeOk() (*string, bool) {
	if o == nil || o.CreationDateTime == nil {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasCreationDateTime() bool {
	if o != nil && o.CreationDateTime != nil {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given string and assigns it to the CreationDateTime field.
func (o *EnrollmentResponse) SetCreationDateTime(v string) {
	o.CreationDateTime = &v
}

// GetExpiryDateTime returns the ExpiryDateTime field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetExpiryDateTime() string {
	if o == nil || o.ExpiryDateTime == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDateTime
}

// GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetExpiryDateTimeOk() (*string, bool) {
	if o == nil || o.ExpiryDateTime == nil {
		return nil, false
	}
	return o.ExpiryDateTime, true
}

// HasExpiryDateTime returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasExpiryDateTime() bool {
	if o != nil && o.ExpiryDateTime != nil {
		return true
	}

	return false
}

// SetExpiryDateTime gets a reference to the given string and assigns it to the ExpiryDateTime field.
func (o *EnrollmentResponse) SetExpiryDateTime(v string) {
	o.ExpiryDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnrollmentResponse) SetStatus(v string) {
	o.Status = &v
}

// GetNumberOfBadAttempts returns the NumberOfBadAttempts field value if set, zero value otherwise.
func (o *EnrollmentResponse) GetNumberOfBadAttempts() float32 {
	if o == nil || o.NumberOfBadAttempts == nil {
		var ret float32
		return ret
	}
	return *o.NumberOfBadAttempts
}

// GetNumberOfBadAttemptsOk returns a tuple with the NumberOfBadAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentResponse) GetNumberOfBadAttemptsOk() (*float32, bool) {
	if o == nil || o.NumberOfBadAttempts == nil {
		return nil, false
	}
	return o.NumberOfBadAttempts, true
}

// HasNumberOfBadAttempts returns a boolean if a field has been set.
func (o *EnrollmentResponse) HasNumberOfBadAttempts() bool {
	if o != nil && o.NumberOfBadAttempts != nil {
		return true
	}

	return false
}

// SetNumberOfBadAttempts gets a reference to the given float32 and assigns it to the NumberOfBadAttempts field.
func (o *EnrollmentResponse) SetNumberOfBadAttempts(v float32) {
	o.NumberOfBadAttempts = &v
}

func (o EnrollmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SeatId != nil {
		toSerialize["seat_id"] = o.SeatId
	}
	if o.ProfileId != nil {
		toSerialize["profile_id"] = o.ProfileId
	}
	if o.ProfileName != nil {
		toSerialize["profile_name"] = o.ProfileName
	}
	if o.EnrollmentCode != nil {
		toSerialize["enrollment_code"] = o.EnrollmentCode
	}
	if o.CreationDateTime != nil {
		toSerialize["creation_date_time"] = o.CreationDateTime
	}
	if o.ExpiryDateTime != nil {
		toSerialize["expiry_date_time"] = o.ExpiryDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.NumberOfBadAttempts != nil {
		toSerialize["number_of_bad_attempts"] = o.NumberOfBadAttempts
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentResponse struct {
	value *EnrollmentResponse
	isSet bool
}

func (v NullableEnrollmentResponse) Get() *EnrollmentResponse {
	return v.value
}

func (v *NullableEnrollmentResponse) Set(val *EnrollmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentResponse(val *EnrollmentResponse) *NullableEnrollmentResponse {
	return &NullableEnrollmentResponse{value: val, isSet: true}
}

func (v NullableEnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


