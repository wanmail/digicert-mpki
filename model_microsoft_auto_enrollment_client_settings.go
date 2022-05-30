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

// MicrosoftAutoEnrollmentClientSettings struct for MicrosoftAutoEnrollmentClientSettings
type MicrosoftAutoEnrollmentClientSettings struct {
	PrivateKeyExport *MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport `json:"private_key_export,omitempty"`
	EnrollmentMode *MicrosoftAutoEnrollmentClientSettingsEnrollmentMode `json:"enrollment_mode,omitempty"`
	// Describes KeySpec
	KeySpec *string `json:"key_spec,omitempty"`
	// Certificate template schema version
	TemplateSchemaVersion *int32 `json:"template_schema_version,omitempty"`
	// If true, certificate is used for Windows Hello for Business authentication
	WindowsHelloLogon *bool `json:"windows_hello_logon,omitempty"`
	// If true, certificate is enrolled on behalf of user
	EnrollOnBehalfOf *bool `json:"enroll_on_behalf_of,omitempty"`
	// If true, certificate is published to Active Directory
	PublishToAd *bool `json:"publish_to_ad,omitempty"`
	// Describes the type of End Entity Type
	EndEntityType *string `json:"end_entity_type,omitempty"`
	CryptoProviders *MicrosoftAutoEnrollmentClientSettingsCryptoProviders `json:"crypto_providers,omitempty"`
}

// NewMicrosoftAutoEnrollmentClientSettings instantiates a new MicrosoftAutoEnrollmentClientSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftAutoEnrollmentClientSettings() *MicrosoftAutoEnrollmentClientSettings {
	this := MicrosoftAutoEnrollmentClientSettings{}
	return &this
}

// NewMicrosoftAutoEnrollmentClientSettingsWithDefaults instantiates a new MicrosoftAutoEnrollmentClientSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftAutoEnrollmentClientSettingsWithDefaults() *MicrosoftAutoEnrollmentClientSettings {
	this := MicrosoftAutoEnrollmentClientSettings{}
	return &this
}

// GetPrivateKeyExport returns the PrivateKeyExport field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetPrivateKeyExport() MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport {
	if o == nil || o.PrivateKeyExport == nil {
		var ret MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport
		return ret
	}
	return *o.PrivateKeyExport
}

// GetPrivateKeyExportOk returns a tuple with the PrivateKeyExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetPrivateKeyExportOk() (*MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport, bool) {
	if o == nil || o.PrivateKeyExport == nil {
		return nil, false
	}
	return o.PrivateKeyExport, true
}

// HasPrivateKeyExport returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasPrivateKeyExport() bool {
	if o != nil && o.PrivateKeyExport != nil {
		return true
	}

	return false
}

// SetPrivateKeyExport gets a reference to the given MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport and assigns it to the PrivateKeyExport field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetPrivateKeyExport(v MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport) {
	o.PrivateKeyExport = &v
}

// GetEnrollmentMode returns the EnrollmentMode field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollmentMode() MicrosoftAutoEnrollmentClientSettingsEnrollmentMode {
	if o == nil || o.EnrollmentMode == nil {
		var ret MicrosoftAutoEnrollmentClientSettingsEnrollmentMode
		return ret
	}
	return *o.EnrollmentMode
}

// GetEnrollmentModeOk returns a tuple with the EnrollmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollmentModeOk() (*MicrosoftAutoEnrollmentClientSettingsEnrollmentMode, bool) {
	if o == nil || o.EnrollmentMode == nil {
		return nil, false
	}
	return o.EnrollmentMode, true
}

// HasEnrollmentMode returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasEnrollmentMode() bool {
	if o != nil && o.EnrollmentMode != nil {
		return true
	}

	return false
}

// SetEnrollmentMode gets a reference to the given MicrosoftAutoEnrollmentClientSettingsEnrollmentMode and assigns it to the EnrollmentMode field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetEnrollmentMode(v MicrosoftAutoEnrollmentClientSettingsEnrollmentMode) {
	o.EnrollmentMode = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetKeySpec() string {
	if o == nil || o.KeySpec == nil {
		var ret string
		return ret
	}
	return *o.KeySpec
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetKeySpecOk() (*string, bool) {
	if o == nil || o.KeySpec == nil {
		return nil, false
	}
	return o.KeySpec, true
}

// HasKeySpec returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasKeySpec() bool {
	if o != nil && o.KeySpec != nil {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given string and assigns it to the KeySpec field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetKeySpec(v string) {
	o.KeySpec = &v
}

// GetTemplateSchemaVersion returns the TemplateSchemaVersion field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetTemplateSchemaVersion() int32 {
	if o == nil || o.TemplateSchemaVersion == nil {
		var ret int32
		return ret
	}
	return *o.TemplateSchemaVersion
}

// GetTemplateSchemaVersionOk returns a tuple with the TemplateSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetTemplateSchemaVersionOk() (*int32, bool) {
	if o == nil || o.TemplateSchemaVersion == nil {
		return nil, false
	}
	return o.TemplateSchemaVersion, true
}

// HasTemplateSchemaVersion returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasTemplateSchemaVersion() bool {
	if o != nil && o.TemplateSchemaVersion != nil {
		return true
	}

	return false
}

// SetTemplateSchemaVersion gets a reference to the given int32 and assigns it to the TemplateSchemaVersion field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetTemplateSchemaVersion(v int32) {
	o.TemplateSchemaVersion = &v
}

// GetWindowsHelloLogon returns the WindowsHelloLogon field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetWindowsHelloLogon() bool {
	if o == nil || o.WindowsHelloLogon == nil {
		var ret bool
		return ret
	}
	return *o.WindowsHelloLogon
}

// GetWindowsHelloLogonOk returns a tuple with the WindowsHelloLogon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetWindowsHelloLogonOk() (*bool, bool) {
	if o == nil || o.WindowsHelloLogon == nil {
		return nil, false
	}
	return o.WindowsHelloLogon, true
}

// HasWindowsHelloLogon returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasWindowsHelloLogon() bool {
	if o != nil && o.WindowsHelloLogon != nil {
		return true
	}

	return false
}

// SetWindowsHelloLogon gets a reference to the given bool and assigns it to the WindowsHelloLogon field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetWindowsHelloLogon(v bool) {
	o.WindowsHelloLogon = &v
}

// GetEnrollOnBehalfOf returns the EnrollOnBehalfOf field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollOnBehalfOf() bool {
	if o == nil || o.EnrollOnBehalfOf == nil {
		var ret bool
		return ret
	}
	return *o.EnrollOnBehalfOf
}

// GetEnrollOnBehalfOfOk returns a tuple with the EnrollOnBehalfOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollOnBehalfOfOk() (*bool, bool) {
	if o == nil || o.EnrollOnBehalfOf == nil {
		return nil, false
	}
	return o.EnrollOnBehalfOf, true
}

// HasEnrollOnBehalfOf returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasEnrollOnBehalfOf() bool {
	if o != nil && o.EnrollOnBehalfOf != nil {
		return true
	}

	return false
}

// SetEnrollOnBehalfOf gets a reference to the given bool and assigns it to the EnrollOnBehalfOf field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetEnrollOnBehalfOf(v bool) {
	o.EnrollOnBehalfOf = &v
}

// GetPublishToAd returns the PublishToAd field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetPublishToAd() bool {
	if o == nil || o.PublishToAd == nil {
		var ret bool
		return ret
	}
	return *o.PublishToAd
}

// GetPublishToAdOk returns a tuple with the PublishToAd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetPublishToAdOk() (*bool, bool) {
	if o == nil || o.PublishToAd == nil {
		return nil, false
	}
	return o.PublishToAd, true
}

// HasPublishToAd returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasPublishToAd() bool {
	if o != nil && o.PublishToAd != nil {
		return true
	}

	return false
}

// SetPublishToAd gets a reference to the given bool and assigns it to the PublishToAd field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetPublishToAd(v bool) {
	o.PublishToAd = &v
}

// GetEndEntityType returns the EndEntityType field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEndEntityType() string {
	if o == nil || o.EndEntityType == nil {
		var ret string
		return ret
	}
	return *o.EndEntityType
}

// GetEndEntityTypeOk returns a tuple with the EndEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetEndEntityTypeOk() (*string, bool) {
	if o == nil || o.EndEntityType == nil {
		return nil, false
	}
	return o.EndEntityType, true
}

// HasEndEntityType returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasEndEntityType() bool {
	if o != nil && o.EndEntityType != nil {
		return true
	}

	return false
}

// SetEndEntityType gets a reference to the given string and assigns it to the EndEntityType field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetEndEntityType(v string) {
	o.EndEntityType = &v
}

// GetCryptoProviders returns the CryptoProviders field value if set, zero value otherwise.
func (o *MicrosoftAutoEnrollmentClientSettings) GetCryptoProviders() MicrosoftAutoEnrollmentClientSettingsCryptoProviders {
	if o == nil || o.CryptoProviders == nil {
		var ret MicrosoftAutoEnrollmentClientSettingsCryptoProviders
		return ret
	}
	return *o.CryptoProviders
}

// GetCryptoProvidersOk returns a tuple with the CryptoProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) GetCryptoProvidersOk() (*MicrosoftAutoEnrollmentClientSettingsCryptoProviders, bool) {
	if o == nil || o.CryptoProviders == nil {
		return nil, false
	}
	return o.CryptoProviders, true
}

// HasCryptoProviders returns a boolean if a field has been set.
func (o *MicrosoftAutoEnrollmentClientSettings) HasCryptoProviders() bool {
	if o != nil && o.CryptoProviders != nil {
		return true
	}

	return false
}

// SetCryptoProviders gets a reference to the given MicrosoftAutoEnrollmentClientSettingsCryptoProviders and assigns it to the CryptoProviders field.
func (o *MicrosoftAutoEnrollmentClientSettings) SetCryptoProviders(v MicrosoftAutoEnrollmentClientSettingsCryptoProviders) {
	o.CryptoProviders = &v
}

func (o MicrosoftAutoEnrollmentClientSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrivateKeyExport != nil {
		toSerialize["private_key_export"] = o.PrivateKeyExport
	}
	if o.EnrollmentMode != nil {
		toSerialize["enrollment_mode"] = o.EnrollmentMode
	}
	if o.KeySpec != nil {
		toSerialize["key_spec"] = o.KeySpec
	}
	if o.TemplateSchemaVersion != nil {
		toSerialize["template_schema_version"] = o.TemplateSchemaVersion
	}
	if o.WindowsHelloLogon != nil {
		toSerialize["windows_hello_logon"] = o.WindowsHelloLogon
	}
	if o.EnrollOnBehalfOf != nil {
		toSerialize["enroll_on_behalf_of"] = o.EnrollOnBehalfOf
	}
	if o.PublishToAd != nil {
		toSerialize["publish_to_ad"] = o.PublishToAd
	}
	if o.EndEntityType != nil {
		toSerialize["end_entity_type"] = o.EndEntityType
	}
	if o.CryptoProviders != nil {
		toSerialize["crypto_providers"] = o.CryptoProviders
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftAutoEnrollmentClientSettings struct {
	value *MicrosoftAutoEnrollmentClientSettings
	isSet bool
}

func (v NullableMicrosoftAutoEnrollmentClientSettings) Get() *MicrosoftAutoEnrollmentClientSettings {
	return v.value
}

func (v *NullableMicrosoftAutoEnrollmentClientSettings) Set(val *MicrosoftAutoEnrollmentClientSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftAutoEnrollmentClientSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftAutoEnrollmentClientSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftAutoEnrollmentClientSettings(val *MicrosoftAutoEnrollmentClientSettings) *NullableMicrosoftAutoEnrollmentClientSettings {
	return &NullableMicrosoftAutoEnrollmentClientSettings{value: val, isSet: true}
}

func (v NullableMicrosoftAutoEnrollmentClientSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftAutoEnrollmentClientSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


