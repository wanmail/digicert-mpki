/*
Enterprise PKI Manager REST API

 Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.  ## Base URL  The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.  Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.  ## Authentication  API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods: * Header-based API token authentication. * Authentication using a client authentication certificate.  ### API token  To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header: * A Service user token ID (**recommended**).     * An API token bound to a DigiCert ONE administrator.  **Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs.   **Service user token ID** * Service users are API tokens that are not associated with a specific user. * When you create a service user, you assign only the permissions needed for the API integration. * There are two ways to create a new service user:    * 1- Use the Account Manager in DigiCert ONE:      1. Navigate to the DigiCert ONE **Account Manager**.     2. Select **Access** from the left menu.     3. Select **Service User** from the left menu.     4. Select **Create service user** and follow the remaining prompts.    * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.   **Administrator API token**  * Standard users (administrators) can create API tokens in DigiCert ONE.  * API tokens have the same permissions and access scope as the administrator that created them.  * Actions linked to the API token are logged under the administrator's name in event audit logs. * To generate a new API token:     1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.   3. Select **Administrators**.   4. Select the administrator you wish to create an API token for from the list.   5. Scroll down to the **API Tokens** section on the administrator page.    6. Select **Create API token** and follow the remaining prompts.    ### Client authentication certificate  When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.  To generate a client authentication certificate:    1. Navigate to the DigiCert ONE **Account Manager**.   2. Select **Access** from the left menu.    3. Search for and select your Service user.    4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.   5. Enter a **Nickname** and select an **End date** for the certificate.    6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.   7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.   8. Review the permissions you wish to provide to the Service user that the certificate is associated with.      **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.    To use a client authentication certificate:  * Include the certificate in your API request. * In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com` * Omit the `x-api-key` header.  ## Requests  The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.  A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.  ### Method  The Enterprise PKI Manager API uses these standard HTTP methods:  * GET * POST * PUT * DELETE  ### Body and content type  All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.  GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).  ## Responses  Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.  **Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.  ### Headers  Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.  * HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation. * HTTP codes in the **400+** range describe an error.  Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.  **Example error response**  ```JSON {   \"error\": {     \"code\": \"no_access\",     \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"   } } ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LogSearchParameters struct for LogSearchParameters
type LogSearchParameters struct {
	AccountId *string `json:"account_id,omitempty"`
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	AdminId *string `json:"admin_id,omitempty"`
	ResourceId *string `json:"resource_id,omitempty"`
	ResourceType []string `json:"resource_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Action *string `json:"action,omitempty"`
	From *time.Time `json:"from,omitempty"`
	To *time.Time `json:"to,omitempty"`
}

// NewLogSearchParameters instantiates a new LogSearchParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchParameters() *LogSearchParameters {
	this := LogSearchParameters{}
	return &this
}

// NewLogSearchParametersWithDefaults instantiates a new LogSearchParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchParametersWithDefaults() *LogSearchParameters {
	this := LogSearchParameters{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LogSearchParameters) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LogSearchParameters) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LogSearchParameters) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *LogSearchParameters) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *LogSearchParameters) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *LogSearchParameters) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *LogSearchParameters) GetAdminId() string {
	if o == nil || o.AdminId == nil {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetAdminIdOk() (*string, bool) {
	if o == nil || o.AdminId == nil {
		return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *LogSearchParameters) HasAdminId() bool {
	if o != nil && o.AdminId != nil {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *LogSearchParameters) SetAdminId(v string) {
	o.AdminId = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *LogSearchParameters) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *LogSearchParameters) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *LogSearchParameters) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *LogSearchParameters) GetResourceType() []string {
	if o == nil || o.ResourceType == nil {
		var ret []string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetResourceTypeOk() ([]string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *LogSearchParameters) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given []string and assigns it to the ResourceType field.
func (o *LogSearchParameters) SetResourceType(v []string) {
	o.ResourceType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogSearchParameters) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogSearchParameters) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LogSearchParameters) SetStatus(v string) {
	o.Status = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *LogSearchParameters) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *LogSearchParameters) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *LogSearchParameters) SetAction(v string) {
	o.Action = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *LogSearchParameters) GetFrom() time.Time {
	if o == nil || o.From == nil {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetFromOk() (*time.Time, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LogSearchParameters) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *LogSearchParameters) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LogSearchParameters) GetTo() time.Time {
	if o == nil || o.To == nil {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchParameters) GetToOk() (*time.Time, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LogSearchParameters) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *LogSearchParameters) SetTo(v time.Time) {
	o.To = &v
}

func (o LogSearchParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.AdminId != nil {
		toSerialize["admin_id"] = o.AdminId
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableLogSearchParameters struct {
	value *LogSearchParameters
	isSet bool
}

func (v NullableLogSearchParameters) Get() *LogSearchParameters {
	return v.value
}

func (v *NullableLogSearchParameters) Set(val *LogSearchParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchParameters(val *LogSearchParameters) *NullableLogSearchParameters {
	return &NullableLogSearchParameters{value: val, isSet: true}
}

func (v NullableLogSearchParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


