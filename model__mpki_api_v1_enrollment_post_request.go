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

// MpkiApiV1EnrollmentPostRequest struct for MpkiApiV1EnrollmentPostRequest
type MpkiApiV1EnrollmentPostRequest struct {
	// Entity ID in UUID format
	Profile string `json:"profile"`
	Seat MpkiApiV1EnrollmentPostRequestSeat `json:"seat"`
	Attributes map[string]interface{} `json:"attributes"`
	EnrollmentCode *string `json:"enrollment_code,omitempty"`
	// This parameter can only accept \"yyyy-MM-dd\" and \"yyyy-MM-dd HH:mm:ss\" formats
	ExpiryDate *string `json:"expiry_date,omitempty"`
}

// NewMpkiApiV1EnrollmentPostRequest instantiates a new MpkiApiV1EnrollmentPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpkiApiV1EnrollmentPostRequest(profile string, seat MpkiApiV1EnrollmentPostRequestSeat, attributes map[string]interface{}) *MpkiApiV1EnrollmentPostRequest {
	this := MpkiApiV1EnrollmentPostRequest{}
	this.Profile = profile
	this.Seat = seat
	this.Attributes = attributes
	return &this
}

// NewMpkiApiV1EnrollmentPostRequestWithDefaults instantiates a new MpkiApiV1EnrollmentPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpkiApiV1EnrollmentPostRequestWithDefaults() *MpkiApiV1EnrollmentPostRequest {
	this := MpkiApiV1EnrollmentPostRequest{}
	return &this
}

// GetProfile returns the Profile field value
func (o *MpkiApiV1EnrollmentPostRequest) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1EnrollmentPostRequest) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *MpkiApiV1EnrollmentPostRequest) SetProfile(v string) {
	o.Profile = v
}

// GetSeat returns the Seat field value
func (o *MpkiApiV1EnrollmentPostRequest) GetSeat() MpkiApiV1EnrollmentPostRequestSeat {
	if o == nil {
		var ret MpkiApiV1EnrollmentPostRequestSeat
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1EnrollmentPostRequest) GetSeatOk() (*MpkiApiV1EnrollmentPostRequestSeat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *MpkiApiV1EnrollmentPostRequest) SetSeat(v MpkiApiV1EnrollmentPostRequestSeat) {
	o.Seat = v
}

// GetAttributes returns the Attributes field value
func (o *MpkiApiV1EnrollmentPostRequest) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MpkiApiV1EnrollmentPostRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *MpkiApiV1EnrollmentPostRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetEnrollmentCode returns the EnrollmentCode field value if set, zero value otherwise.
func (o *MpkiApiV1EnrollmentPostRequest) GetEnrollmentCode() string {
	if o == nil || o.EnrollmentCode == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentCode
}

// GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1EnrollmentPostRequest) GetEnrollmentCodeOk() (*string, bool) {
	if o == nil || o.EnrollmentCode == nil {
		return nil, false
	}
	return o.EnrollmentCode, true
}

// HasEnrollmentCode returns a boolean if a field has been set.
func (o *MpkiApiV1EnrollmentPostRequest) HasEnrollmentCode() bool {
	if o != nil && o.EnrollmentCode != nil {
		return true
	}

	return false
}

// SetEnrollmentCode gets a reference to the given string and assigns it to the EnrollmentCode field.
func (o *MpkiApiV1EnrollmentPostRequest) SetEnrollmentCode(v string) {
	o.EnrollmentCode = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *MpkiApiV1EnrollmentPostRequest) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpkiApiV1EnrollmentPostRequest) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *MpkiApiV1EnrollmentPostRequest) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *MpkiApiV1EnrollmentPostRequest) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

func (o MpkiApiV1EnrollmentPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["seat"] = o.Seat
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.EnrollmentCode != nil {
		toSerialize["enrollment_code"] = o.EnrollmentCode
	}
	if o.ExpiryDate != nil {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	return json.Marshal(toSerialize)
}

type NullableMpkiApiV1EnrollmentPostRequest struct {
	value *MpkiApiV1EnrollmentPostRequest
	isSet bool
}

func (v NullableMpkiApiV1EnrollmentPostRequest) Get() *MpkiApiV1EnrollmentPostRequest {
	return v.value
}

func (v *NullableMpkiApiV1EnrollmentPostRequest) Set(val *MpkiApiV1EnrollmentPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMpkiApiV1EnrollmentPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMpkiApiV1EnrollmentPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpkiApiV1EnrollmentPostRequest(val *MpkiApiV1EnrollmentPostRequest) *NullableMpkiApiV1EnrollmentPostRequest {
	return &NullableMpkiApiV1EnrollmentPostRequest{value: val, isSet: true}
}

func (v NullableMpkiApiV1EnrollmentPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpkiApiV1EnrollmentPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


