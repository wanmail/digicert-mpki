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

// PublicEnrollmentDetails struct for PublicEnrollmentDetails
type PublicEnrollmentDetails struct {
	// Enrollment internal id
	Id *string `json:"id,omitempty"`
	// Status of the enrollment
	Status *string `json:"status,omitempty"`
	// Email used to enroll
	Email *string `json:"email,omitempty"`
	// Creation date/time of the enrollment
	CreatedAt *string `json:"created_at,omitempty"`
	// Expiry date/time of the enrollment
	ExpiresAt *string `json:"expires_at,omitempty"`
	Profile *NameDto `json:"profile,omitempty"`
	BusinessUnit *NameDto `json:"business_unit,omitempty"`
	// A message sent to user by admin
	MessageToUser *string `json:"message_to_user,omitempty"`
	// Value used as seat id mapping
	SeatIdMappingValue *string `json:"seat_id_mapping_value,omitempty"`
	// Whether this particular enrollment has been approved by the current user
	ApprovedByCurrentUser *bool `json:"approved_by_current_user,omitempty"`
	Seat *PublicEnrollmentDetailsSeat `json:"seat,omitempty"`
	// Comments to the enrollment
	Comments *string `json:"comments,omitempty"`
}

// NewPublicEnrollmentDetails instantiates a new PublicEnrollmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicEnrollmentDetails() *PublicEnrollmentDetails {
	this := PublicEnrollmentDetails{}
	return &this
}

// NewPublicEnrollmentDetailsWithDefaults instantiates a new PublicEnrollmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicEnrollmentDetailsWithDefaults() *PublicEnrollmentDetails {
	this := PublicEnrollmentDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicEnrollmentDetails) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PublicEnrollmentDetails) SetStatus(v string) {
	o.Status = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PublicEnrollmentDetails) SetEmail(v string) {
	o.Email = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PublicEnrollmentDetails) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PublicEnrollmentDetails) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetProfile() NameDto {
	if o == nil || o.Profile == nil {
		var ret NameDto
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetProfileOk() (*NameDto, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NameDto and assigns it to the Profile field.
func (o *PublicEnrollmentDetails) SetProfile(v NameDto) {
	o.Profile = &v
}

// GetBusinessUnit returns the BusinessUnit field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetBusinessUnit() NameDto {
	if o == nil || o.BusinessUnit == nil {
		var ret NameDto
		return ret
	}
	return *o.BusinessUnit
}

// GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetBusinessUnitOk() (*NameDto, bool) {
	if o == nil || o.BusinessUnit == nil {
		return nil, false
	}
	return o.BusinessUnit, true
}

// HasBusinessUnit returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasBusinessUnit() bool {
	if o != nil && o.BusinessUnit != nil {
		return true
	}

	return false
}

// SetBusinessUnit gets a reference to the given NameDto and assigns it to the BusinessUnit field.
func (o *PublicEnrollmentDetails) SetBusinessUnit(v NameDto) {
	o.BusinessUnit = &v
}

// GetMessageToUser returns the MessageToUser field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetMessageToUser() string {
	if o == nil || o.MessageToUser == nil {
		var ret string
		return ret
	}
	return *o.MessageToUser
}

// GetMessageToUserOk returns a tuple with the MessageToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetMessageToUserOk() (*string, bool) {
	if o == nil || o.MessageToUser == nil {
		return nil, false
	}
	return o.MessageToUser, true
}

// HasMessageToUser returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasMessageToUser() bool {
	if o != nil && o.MessageToUser != nil {
		return true
	}

	return false
}

// SetMessageToUser gets a reference to the given string and assigns it to the MessageToUser field.
func (o *PublicEnrollmentDetails) SetMessageToUser(v string) {
	o.MessageToUser = &v
}

// GetSeatIdMappingValue returns the SeatIdMappingValue field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetSeatIdMappingValue() string {
	if o == nil || o.SeatIdMappingValue == nil {
		var ret string
		return ret
	}
	return *o.SeatIdMappingValue
}

// GetSeatIdMappingValueOk returns a tuple with the SeatIdMappingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetSeatIdMappingValueOk() (*string, bool) {
	if o == nil || o.SeatIdMappingValue == nil {
		return nil, false
	}
	return o.SeatIdMappingValue, true
}

// HasSeatIdMappingValue returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasSeatIdMappingValue() bool {
	if o != nil && o.SeatIdMappingValue != nil {
		return true
	}

	return false
}

// SetSeatIdMappingValue gets a reference to the given string and assigns it to the SeatIdMappingValue field.
func (o *PublicEnrollmentDetails) SetSeatIdMappingValue(v string) {
	o.SeatIdMappingValue = &v
}

// GetApprovedByCurrentUser returns the ApprovedByCurrentUser field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetApprovedByCurrentUser() bool {
	if o == nil || o.ApprovedByCurrentUser == nil {
		var ret bool
		return ret
	}
	return *o.ApprovedByCurrentUser
}

// GetApprovedByCurrentUserOk returns a tuple with the ApprovedByCurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetApprovedByCurrentUserOk() (*bool, bool) {
	if o == nil || o.ApprovedByCurrentUser == nil {
		return nil, false
	}
	return o.ApprovedByCurrentUser, true
}

// HasApprovedByCurrentUser returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasApprovedByCurrentUser() bool {
	if o != nil && o.ApprovedByCurrentUser != nil {
		return true
	}

	return false
}

// SetApprovedByCurrentUser gets a reference to the given bool and assigns it to the ApprovedByCurrentUser field.
func (o *PublicEnrollmentDetails) SetApprovedByCurrentUser(v bool) {
	o.ApprovedByCurrentUser = &v
}

// GetSeat returns the Seat field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetSeat() PublicEnrollmentDetailsSeat {
	if o == nil || o.Seat == nil {
		var ret PublicEnrollmentDetailsSeat
		return ret
	}
	return *o.Seat
}

// GetSeatOk returns a tuple with the Seat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetSeatOk() (*PublicEnrollmentDetailsSeat, bool) {
	if o == nil || o.Seat == nil {
		return nil, false
	}
	return o.Seat, true
}

// HasSeat returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasSeat() bool {
	if o != nil && o.Seat != nil {
		return true
	}

	return false
}

// SetSeat gets a reference to the given PublicEnrollmentDetailsSeat and assigns it to the Seat field.
func (o *PublicEnrollmentDetails) SetSeat(v PublicEnrollmentDetailsSeat) {
	o.Seat = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PublicEnrollmentDetails) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEnrollmentDetails) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PublicEnrollmentDetails) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PublicEnrollmentDetails) SetComments(v string) {
	o.Comments = &v
}

func (o PublicEnrollmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.BusinessUnit != nil {
		toSerialize["business_unit"] = o.BusinessUnit
	}
	if o.MessageToUser != nil {
		toSerialize["message_to_user"] = o.MessageToUser
	}
	if o.SeatIdMappingValue != nil {
		toSerialize["seat_id_mapping_value"] = o.SeatIdMappingValue
	}
	if o.ApprovedByCurrentUser != nil {
		toSerialize["approved_by_current_user"] = o.ApprovedByCurrentUser
	}
	if o.Seat != nil {
		toSerialize["seat"] = o.Seat
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	return json.Marshal(toSerialize)
}

type NullablePublicEnrollmentDetails struct {
	value *PublicEnrollmentDetails
	isSet bool
}

func (v NullablePublicEnrollmentDetails) Get() *PublicEnrollmentDetails {
	return v.value
}

func (v *NullablePublicEnrollmentDetails) Set(val *PublicEnrollmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicEnrollmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicEnrollmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicEnrollmentDetails(val *PublicEnrollmentDetails) *NullablePublicEnrollmentDetails {
	return &NullablePublicEnrollmentDetails{value: val, isSet: true}
}

func (v NullablePublicEnrollmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicEnrollmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


