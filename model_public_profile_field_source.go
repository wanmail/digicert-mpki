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

// PublicProfileFieldSource struct for PublicProfileFieldSource
type PublicProfileFieldSource struct {
	// If true, the source is mandatory
	Mandatory *bool `json:"mandatory,omitempty"`
	// Source of the attribute
	Source *string `json:"source,omitempty"`
	// Permanent attribute value if source is `FIXED_VALUE`
	Value map[string]interface{} `json:"value,omitempty"`
	// Permanent custom Other name value if source is `FIXED_VALUE`
	CustomValue map[string]interface{} `json:"custom_value,omitempty"`
	// SAML attribute mapping value
	SamlAssertionKey *string `json:"saml_assertion_key,omitempty"`
	// Active Directory attribute mapping value
	AdAttributeKey *string `json:"ad_attribute_key,omitempty"`
}

// NewPublicProfileFieldSource instantiates a new PublicProfileFieldSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicProfileFieldSource() *PublicProfileFieldSource {
	this := PublicProfileFieldSource{}
	return &this
}

// NewPublicProfileFieldSourceWithDefaults instantiates a new PublicProfileFieldSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicProfileFieldSourceWithDefaults() *PublicProfileFieldSource {
	this := PublicProfileFieldSource{}
	return &this
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetMandatory() bool {
	if o == nil || o.Mandatory == nil {
		var ret bool
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetMandatoryOk() (*bool, bool) {
	if o == nil || o.Mandatory == nil {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasMandatory() bool {
	if o != nil && o.Mandatory != nil {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given bool and assigns it to the Mandatory field.
func (o *PublicProfileFieldSource) SetMandatory(v bool) {
	o.Mandatory = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PublicProfileFieldSource) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *PublicProfileFieldSource) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetCustomValue returns the CustomValue field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetCustomValue() map[string]interface{} {
	if o == nil || o.CustomValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomValue
}

// GetCustomValueOk returns a tuple with the CustomValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetCustomValueOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomValue == nil {
		return nil, false
	}
	return o.CustomValue, true
}

// HasCustomValue returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasCustomValue() bool {
	if o != nil && o.CustomValue != nil {
		return true
	}

	return false
}

// SetCustomValue gets a reference to the given map[string]interface{} and assigns it to the CustomValue field.
func (o *PublicProfileFieldSource) SetCustomValue(v map[string]interface{}) {
	o.CustomValue = v
}

// GetSamlAssertionKey returns the SamlAssertionKey field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetSamlAssertionKey() string {
	if o == nil || o.SamlAssertionKey == nil {
		var ret string
		return ret
	}
	return *o.SamlAssertionKey
}

// GetSamlAssertionKeyOk returns a tuple with the SamlAssertionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetSamlAssertionKeyOk() (*string, bool) {
	if o == nil || o.SamlAssertionKey == nil {
		return nil, false
	}
	return o.SamlAssertionKey, true
}

// HasSamlAssertionKey returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasSamlAssertionKey() bool {
	if o != nil && o.SamlAssertionKey != nil {
		return true
	}

	return false
}

// SetSamlAssertionKey gets a reference to the given string and assigns it to the SamlAssertionKey field.
func (o *PublicProfileFieldSource) SetSamlAssertionKey(v string) {
	o.SamlAssertionKey = &v
}

// GetAdAttributeKey returns the AdAttributeKey field value if set, zero value otherwise.
func (o *PublicProfileFieldSource) GetAdAttributeKey() string {
	if o == nil || o.AdAttributeKey == nil {
		var ret string
		return ret
	}
	return *o.AdAttributeKey
}

// GetAdAttributeKeyOk returns a tuple with the AdAttributeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileFieldSource) GetAdAttributeKeyOk() (*string, bool) {
	if o == nil || o.AdAttributeKey == nil {
		return nil, false
	}
	return o.AdAttributeKey, true
}

// HasAdAttributeKey returns a boolean if a field has been set.
func (o *PublicProfileFieldSource) HasAdAttributeKey() bool {
	if o != nil && o.AdAttributeKey != nil {
		return true
	}

	return false
}

// SetAdAttributeKey gets a reference to the given string and assigns it to the AdAttributeKey field.
func (o *PublicProfileFieldSource) SetAdAttributeKey(v string) {
	o.AdAttributeKey = &v
}

func (o PublicProfileFieldSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mandatory != nil {
		toSerialize["mandatory"] = o.Mandatory
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.CustomValue != nil {
		toSerialize["custom_value"] = o.CustomValue
	}
	if o.SamlAssertionKey != nil {
		toSerialize["saml_assertion_key"] = o.SamlAssertionKey
	}
	if o.AdAttributeKey != nil {
		toSerialize["ad_attribute_key"] = o.AdAttributeKey
	}
	return json.Marshal(toSerialize)
}

type NullablePublicProfileFieldSource struct {
	value *PublicProfileFieldSource
	isSet bool
}

func (v NullablePublicProfileFieldSource) Get() *PublicProfileFieldSource {
	return v.value
}

func (v *NullablePublicProfileFieldSource) Set(val *PublicProfileFieldSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicProfileFieldSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicProfileFieldSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicProfileFieldSource(val *PublicProfileFieldSource) *NullablePublicProfileFieldSource {
	return &NullablePublicProfileFieldSource{value: val, isSet: true}
}

func (v NullablePublicProfileFieldSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicProfileFieldSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


