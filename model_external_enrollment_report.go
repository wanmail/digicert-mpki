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

// ExternalEnrollmentReport struct for ExternalEnrollmentReport
type ExternalEnrollmentReport struct {
	Seat *NameDto `json:"seat,omitempty"`
	// Type of Seat to which this Enrollment is linked
	SeatType *string `json:"seat_type,omitempty"`
	// Name of Profile against which Enrollment was issued
	ProfileName *string `json:"profile_name,omitempty"`
	// Name of Business unit to which Enrollment is linked
	BuName *string `json:"bu_name,omitempty"`
	// Subject of Enrollment
	Subject *string `json:"subject,omitempty"`
	// Current status of Enrollment
	EnrollmentStatus *string `json:"enrollment_status,omitempty"`
	// Date when Enrollment was created in format YYYY-MM-DD
	CreateDate *string `json:"create_date,omitempty"`
	// Date when Enrollment should be (was) expired in format YYYY-MM-DD
	ExpiryDate *string `json:"expiry_date,omitempty"`
	// Date when Enrollment was redeemed in format YYYY-MM-DD
	RedeemDate *string `json:"redeem_date,omitempty"`
	// Date when Enrollment was rejected in format YYYY-MM-DD
	RejectDate *string `json:"reject_date,omitempty"`
	// Date when Enrollment was last edited in format YYYY-MM-DD
	LastEditDate *string `json:"last_edit_date,omitempty"`
	LastEditAdmin *NameDto `json:"last_edit_admin,omitempty"`
}

// NewExternalEnrollmentReport instantiates a new ExternalEnrollmentReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEnrollmentReport() *ExternalEnrollmentReport {
	this := ExternalEnrollmentReport{}
	return &this
}

// NewExternalEnrollmentReportWithDefaults instantiates a new ExternalEnrollmentReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEnrollmentReportWithDefaults() *ExternalEnrollmentReport {
	this := ExternalEnrollmentReport{}
	return &this
}

// GetSeat returns the Seat field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetSeat() NameDto {
	if o == nil || o.Seat == nil {
		var ret NameDto
		return ret
	}
	return *o.Seat
}

// GetSeatOk returns a tuple with the Seat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetSeatOk() (*NameDto, bool) {
	if o == nil || o.Seat == nil {
		return nil, false
	}
	return o.Seat, true
}

// HasSeat returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasSeat() bool {
	if o != nil && o.Seat != nil {
		return true
	}

	return false
}

// SetSeat gets a reference to the given NameDto and assigns it to the Seat field.
func (o *ExternalEnrollmentReport) SetSeat(v NameDto) {
	o.Seat = &v
}

// GetSeatType returns the SeatType field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetSeatType() string {
	if o == nil || o.SeatType == nil {
		var ret string
		return ret
	}
	return *o.SeatType
}

// GetSeatTypeOk returns a tuple with the SeatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetSeatTypeOk() (*string, bool) {
	if o == nil || o.SeatType == nil {
		return nil, false
	}
	return o.SeatType, true
}

// HasSeatType returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasSeatType() bool {
	if o != nil && o.SeatType != nil {
		return true
	}

	return false
}

// SetSeatType gets a reference to the given string and assigns it to the SeatType field.
func (o *ExternalEnrollmentReport) SetSeatType(v string) {
	o.SeatType = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *ExternalEnrollmentReport) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetBuName returns the BuName field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetBuName() string {
	if o == nil || o.BuName == nil {
		var ret string
		return ret
	}
	return *o.BuName
}

// GetBuNameOk returns a tuple with the BuName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetBuNameOk() (*string, bool) {
	if o == nil || o.BuName == nil {
		return nil, false
	}
	return o.BuName, true
}

// HasBuName returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasBuName() bool {
	if o != nil && o.BuName != nil {
		return true
	}

	return false
}

// SetBuName gets a reference to the given string and assigns it to the BuName field.
func (o *ExternalEnrollmentReport) SetBuName(v string) {
	o.BuName = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ExternalEnrollmentReport) SetSubject(v string) {
	o.Subject = &v
}

// GetEnrollmentStatus returns the EnrollmentStatus field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetEnrollmentStatus() string {
	if o == nil || o.EnrollmentStatus == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentStatus
}

// GetEnrollmentStatusOk returns a tuple with the EnrollmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetEnrollmentStatusOk() (*string, bool) {
	if o == nil || o.EnrollmentStatus == nil {
		return nil, false
	}
	return o.EnrollmentStatus, true
}

// HasEnrollmentStatus returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasEnrollmentStatus() bool {
	if o != nil && o.EnrollmentStatus != nil {
		return true
	}

	return false
}

// SetEnrollmentStatus gets a reference to the given string and assigns it to the EnrollmentStatus field.
func (o *ExternalEnrollmentReport) SetEnrollmentStatus(v string) {
	o.EnrollmentStatus = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetCreateDate() string {
	if o == nil || o.CreateDate == nil {
		var ret string
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetCreateDateOk() (*string, bool) {
	if o == nil || o.CreateDate == nil {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasCreateDate() bool {
	if o != nil && o.CreateDate != nil {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given string and assigns it to the CreateDate field.
func (o *ExternalEnrollmentReport) SetCreateDate(v string) {
	o.CreateDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ExternalEnrollmentReport) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetRedeemDate returns the RedeemDate field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetRedeemDate() string {
	if o == nil || o.RedeemDate == nil {
		var ret string
		return ret
	}
	return *o.RedeemDate
}

// GetRedeemDateOk returns a tuple with the RedeemDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetRedeemDateOk() (*string, bool) {
	if o == nil || o.RedeemDate == nil {
		return nil, false
	}
	return o.RedeemDate, true
}

// HasRedeemDate returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasRedeemDate() bool {
	if o != nil && o.RedeemDate != nil {
		return true
	}

	return false
}

// SetRedeemDate gets a reference to the given string and assigns it to the RedeemDate field.
func (o *ExternalEnrollmentReport) SetRedeemDate(v string) {
	o.RedeemDate = &v
}

// GetRejectDate returns the RejectDate field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetRejectDate() string {
	if o == nil || o.RejectDate == nil {
		var ret string
		return ret
	}
	return *o.RejectDate
}

// GetRejectDateOk returns a tuple with the RejectDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetRejectDateOk() (*string, bool) {
	if o == nil || o.RejectDate == nil {
		return nil, false
	}
	return o.RejectDate, true
}

// HasRejectDate returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasRejectDate() bool {
	if o != nil && o.RejectDate != nil {
		return true
	}

	return false
}

// SetRejectDate gets a reference to the given string and assigns it to the RejectDate field.
func (o *ExternalEnrollmentReport) SetRejectDate(v string) {
	o.RejectDate = &v
}

// GetLastEditDate returns the LastEditDate field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetLastEditDate() string {
	if o == nil || o.LastEditDate == nil {
		var ret string
		return ret
	}
	return *o.LastEditDate
}

// GetLastEditDateOk returns a tuple with the LastEditDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetLastEditDateOk() (*string, bool) {
	if o == nil || o.LastEditDate == nil {
		return nil, false
	}
	return o.LastEditDate, true
}

// HasLastEditDate returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasLastEditDate() bool {
	if o != nil && o.LastEditDate != nil {
		return true
	}

	return false
}

// SetLastEditDate gets a reference to the given string and assigns it to the LastEditDate field.
func (o *ExternalEnrollmentReport) SetLastEditDate(v string) {
	o.LastEditDate = &v
}

// GetLastEditAdmin returns the LastEditAdmin field value if set, zero value otherwise.
func (o *ExternalEnrollmentReport) GetLastEditAdmin() NameDto {
	if o == nil || o.LastEditAdmin == nil {
		var ret NameDto
		return ret
	}
	return *o.LastEditAdmin
}

// GetLastEditAdminOk returns a tuple with the LastEditAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEnrollmentReport) GetLastEditAdminOk() (*NameDto, bool) {
	if o == nil || o.LastEditAdmin == nil {
		return nil, false
	}
	return o.LastEditAdmin, true
}

// HasLastEditAdmin returns a boolean if a field has been set.
func (o *ExternalEnrollmentReport) HasLastEditAdmin() bool {
	if o != nil && o.LastEditAdmin != nil {
		return true
	}

	return false
}

// SetLastEditAdmin gets a reference to the given NameDto and assigns it to the LastEditAdmin field.
func (o *ExternalEnrollmentReport) SetLastEditAdmin(v NameDto) {
	o.LastEditAdmin = &v
}

func (o ExternalEnrollmentReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Seat != nil {
		toSerialize["seat"] = o.Seat
	}
	if o.SeatType != nil {
		toSerialize["seat_type"] = o.SeatType
	}
	if o.ProfileName != nil {
		toSerialize["profile_name"] = o.ProfileName
	}
	if o.BuName != nil {
		toSerialize["bu_name"] = o.BuName
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.EnrollmentStatus != nil {
		toSerialize["enrollment_status"] = o.EnrollmentStatus
	}
	if o.CreateDate != nil {
		toSerialize["create_date"] = o.CreateDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if o.RedeemDate != nil {
		toSerialize["redeem_date"] = o.RedeemDate
	}
	if o.RejectDate != nil {
		toSerialize["reject_date"] = o.RejectDate
	}
	if o.LastEditDate != nil {
		toSerialize["last_edit_date"] = o.LastEditDate
	}
	if o.LastEditAdmin != nil {
		toSerialize["last_edit_admin"] = o.LastEditAdmin
	}
	return json.Marshal(toSerialize)
}

type NullableExternalEnrollmentReport struct {
	value *ExternalEnrollmentReport
	isSet bool
}

func (v NullableExternalEnrollmentReport) Get() *ExternalEnrollmentReport {
	return v.value
}

func (v *NullableExternalEnrollmentReport) Set(val *ExternalEnrollmentReport) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEnrollmentReport) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEnrollmentReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEnrollmentReport(val *ExternalEnrollmentReport) *NullableExternalEnrollmentReport {
	return &NullableExternalEnrollmentReport{value: val, isSet: true}
}

func (v NullableExternalEnrollmentReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEnrollmentReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


