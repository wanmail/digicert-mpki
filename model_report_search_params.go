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

// ReportSearchParams struct for ReportSearchParams
type ReportSearchParams struct {
	AccountId *string `json:"account_id,omitempty"`
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	ProfileId *string `json:"profile_id,omitempty"`
	SeatId *string `json:"seat_id,omitempty"`
	AdminId *string `json:"admin_id,omitempty"`
	CreateDate *string `json:"create_date,omitempty"`
	RedeemDate *string `json:"redeem_date,omitempty"`
	RejectDate *string `json:"reject_date,omitempty"`
	ResetDate *string `json:"reset_date,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
}

// NewReportSearchParams instantiates a new ReportSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportSearchParams() *ReportSearchParams {
	this := ReportSearchParams{}
	return &this
}

// NewReportSearchParamsWithDefaults instantiates a new ReportSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportSearchParamsWithDefaults() *ReportSearchParams {
	this := ReportSearchParams{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ReportSearchParams) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ReportSearchParams) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ReportSearchParams) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *ReportSearchParams) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *ReportSearchParams) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *ReportSearchParams) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ReportSearchParams) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetProfileIdOk() (*string, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ReportSearchParams) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *ReportSearchParams) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetSeatId returns the SeatId field value if set, zero value otherwise.
func (o *ReportSearchParams) GetSeatId() string {
	if o == nil || o.SeatId == nil {
		var ret string
		return ret
	}
	return *o.SeatId
}

// GetSeatIdOk returns a tuple with the SeatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetSeatIdOk() (*string, bool) {
	if o == nil || o.SeatId == nil {
		return nil, false
	}
	return o.SeatId, true
}

// HasSeatId returns a boolean if a field has been set.
func (o *ReportSearchParams) HasSeatId() bool {
	if o != nil && o.SeatId != nil {
		return true
	}

	return false
}

// SetSeatId gets a reference to the given string and assigns it to the SeatId field.
func (o *ReportSearchParams) SetSeatId(v string) {
	o.SeatId = &v
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *ReportSearchParams) GetAdminId() string {
	if o == nil || o.AdminId == nil {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetAdminIdOk() (*string, bool) {
	if o == nil || o.AdminId == nil {
		return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *ReportSearchParams) HasAdminId() bool {
	if o != nil && o.AdminId != nil {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *ReportSearchParams) SetAdminId(v string) {
	o.AdminId = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *ReportSearchParams) GetCreateDate() string {
	if o == nil || o.CreateDate == nil {
		var ret string
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetCreateDateOk() (*string, bool) {
	if o == nil || o.CreateDate == nil {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *ReportSearchParams) HasCreateDate() bool {
	if o != nil && o.CreateDate != nil {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given string and assigns it to the CreateDate field.
func (o *ReportSearchParams) SetCreateDate(v string) {
	o.CreateDate = &v
}

// GetRedeemDate returns the RedeemDate field value if set, zero value otherwise.
func (o *ReportSearchParams) GetRedeemDate() string {
	if o == nil || o.RedeemDate == nil {
		var ret string
		return ret
	}
	return *o.RedeemDate
}

// GetRedeemDateOk returns a tuple with the RedeemDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetRedeemDateOk() (*string, bool) {
	if o == nil || o.RedeemDate == nil {
		return nil, false
	}
	return o.RedeemDate, true
}

// HasRedeemDate returns a boolean if a field has been set.
func (o *ReportSearchParams) HasRedeemDate() bool {
	if o != nil && o.RedeemDate != nil {
		return true
	}

	return false
}

// SetRedeemDate gets a reference to the given string and assigns it to the RedeemDate field.
func (o *ReportSearchParams) SetRedeemDate(v string) {
	o.RedeemDate = &v
}

// GetRejectDate returns the RejectDate field value if set, zero value otherwise.
func (o *ReportSearchParams) GetRejectDate() string {
	if o == nil || o.RejectDate == nil {
		var ret string
		return ret
	}
	return *o.RejectDate
}

// GetRejectDateOk returns a tuple with the RejectDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetRejectDateOk() (*string, bool) {
	if o == nil || o.RejectDate == nil {
		return nil, false
	}
	return o.RejectDate, true
}

// HasRejectDate returns a boolean if a field has been set.
func (o *ReportSearchParams) HasRejectDate() bool {
	if o != nil && o.RejectDate != nil {
		return true
	}

	return false
}

// SetRejectDate gets a reference to the given string and assigns it to the RejectDate field.
func (o *ReportSearchParams) SetRejectDate(v string) {
	o.RejectDate = &v
}

// GetResetDate returns the ResetDate field value if set, zero value otherwise.
func (o *ReportSearchParams) GetResetDate() string {
	if o == nil || o.ResetDate == nil {
		var ret string
		return ret
	}
	return *o.ResetDate
}

// GetResetDateOk returns a tuple with the ResetDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetResetDateOk() (*string, bool) {
	if o == nil || o.ResetDate == nil {
		return nil, false
	}
	return o.ResetDate, true
}

// HasResetDate returns a boolean if a field has been set.
func (o *ReportSearchParams) HasResetDate() bool {
	if o != nil && o.ResetDate != nil {
		return true
	}

	return false
}

// SetResetDate gets a reference to the given string and assigns it to the ResetDate field.
func (o *ReportSearchParams) SetResetDate(v string) {
	o.ResetDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ReportSearchParams) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSearchParams) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ReportSearchParams) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ReportSearchParams) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

func (o ReportSearchParams) MarshalJSON() ([]byte, error) {
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
	if o.AdminId != nil {
		toSerialize["admin_id"] = o.AdminId
	}
	if o.CreateDate != nil {
		toSerialize["create_date"] = o.CreateDate
	}
	if o.RedeemDate != nil {
		toSerialize["redeem_date"] = o.RedeemDate
	}
	if o.RejectDate != nil {
		toSerialize["reject_date"] = o.RejectDate
	}
	if o.ResetDate != nil {
		toSerialize["reset_date"] = o.ResetDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	return json.Marshal(toSerialize)
}

type NullableReportSearchParams struct {
	value *ReportSearchParams
	isSet bool
}

func (v NullableReportSearchParams) Get() *ReportSearchParams {
	return v.value
}

func (v *NullableReportSearchParams) Set(val *ReportSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableReportSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableReportSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportSearchParams(val *ReportSearchParams) *NullableReportSearchParams {
	return &NullableReportSearchParams{value: val, isSet: true}
}

func (v NullableReportSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


