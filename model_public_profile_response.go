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

// PublicProfileResponse struct for PublicProfileResponse
type PublicProfileResponse struct {
	// Entity ID in UUID format
	Id *string `json:"id,omitempty"`
	// Profile name
	Name *string `json:"name,omitempty"`
	// Profile status
	Status *string `json:"status,omitempty"`
	// Signature algorithm for certificates issued from the profile
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`
	// If true, certificate public key is published to the DigiCert PKI directory
	PublishToPublicDirectory *bool `json:"publish_to_public_directory,omitempty"`
	// Number of days before expiration in which the certificate can be renewed
	RenewalPeriodDays *int32 `json:"renewal_period_days,omitempty"`
	// If true, certificate profile allows duplicate certificates
	DuplicateCertPolicy *bool `json:"duplicate_cert_policy,omitempty"`
	// If true, the default validity period set by the profile can be overridden using a request parameter
	OverrideCertValidityViaApi *bool `json:"override_cert_validity_via_api,omitempty"`
	// Format for certificate delivery specified in the profile
	CertificateDeliveryFormat *string `json:"certificate_delivery_format,omitempty"`
	// Business unit ID to which the profile belongs
	BusinessUnitId *string `json:"business_unit_id,omitempty"`
	EnrollmentMethod *ProfileEnrollmentMethod `json:"enrollment_method,omitempty"`
	AuthenticationMethod *ProfileAuthenticationMethod `json:"authentication_method,omitempty"`
	KeyEscrowPolicy *KeyEscrowPolicy `json:"key_escrow_policy,omitempty"`
	PrivateKeyAttributes *PublicProfileResponsePrivateKeyAttributes `json:"private_key_attributes,omitempty"`
	// If true, use of the certificate profile is restricted to a specific API token
	ApiTokenBindingEnabled *bool `json:"api_token_binding_enabled,omitempty"`
	// API token ID to which the profile is bound if api_token_binding_enabled is true
	ApiTokenId *string `json:"api_token_id,omitempty"`
	Certificate *PublicProfileResponseCertificate `json:"certificate,omitempty"`
}

// NewPublicProfileResponse instantiates a new PublicProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicProfileResponse() *PublicProfileResponse {
	this := PublicProfileResponse{}
	return &this
}

// NewPublicProfileResponseWithDefaults instantiates a new PublicProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicProfileResponseWithDefaults() *PublicProfileResponse {
	this := PublicProfileResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicProfileResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicProfileResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PublicProfileResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *PublicProfileResponse) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetPublishToPublicDirectory returns the PublishToPublicDirectory field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetPublishToPublicDirectory() bool {
	if o == nil || o.PublishToPublicDirectory == nil {
		var ret bool
		return ret
	}
	return *o.PublishToPublicDirectory
}

// GetPublishToPublicDirectoryOk returns a tuple with the PublishToPublicDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetPublishToPublicDirectoryOk() (*bool, bool) {
	if o == nil || o.PublishToPublicDirectory == nil {
		return nil, false
	}
	return o.PublishToPublicDirectory, true
}

// HasPublishToPublicDirectory returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasPublishToPublicDirectory() bool {
	if o != nil && o.PublishToPublicDirectory != nil {
		return true
	}

	return false
}

// SetPublishToPublicDirectory gets a reference to the given bool and assigns it to the PublishToPublicDirectory field.
func (o *PublicProfileResponse) SetPublishToPublicDirectory(v bool) {
	o.PublishToPublicDirectory = &v
}

// GetRenewalPeriodDays returns the RenewalPeriodDays field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetRenewalPeriodDays() int32 {
	if o == nil || o.RenewalPeriodDays == nil {
		var ret int32
		return ret
	}
	return *o.RenewalPeriodDays
}

// GetRenewalPeriodDaysOk returns a tuple with the RenewalPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetRenewalPeriodDaysOk() (*int32, bool) {
	if o == nil || o.RenewalPeriodDays == nil {
		return nil, false
	}
	return o.RenewalPeriodDays, true
}

// HasRenewalPeriodDays returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasRenewalPeriodDays() bool {
	if o != nil && o.RenewalPeriodDays != nil {
		return true
	}

	return false
}

// SetRenewalPeriodDays gets a reference to the given int32 and assigns it to the RenewalPeriodDays field.
func (o *PublicProfileResponse) SetRenewalPeriodDays(v int32) {
	o.RenewalPeriodDays = &v
}

// GetDuplicateCertPolicy returns the DuplicateCertPolicy field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetDuplicateCertPolicy() bool {
	if o == nil || o.DuplicateCertPolicy == nil {
		var ret bool
		return ret
	}
	return *o.DuplicateCertPolicy
}

// GetDuplicateCertPolicyOk returns a tuple with the DuplicateCertPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetDuplicateCertPolicyOk() (*bool, bool) {
	if o == nil || o.DuplicateCertPolicy == nil {
		return nil, false
	}
	return o.DuplicateCertPolicy, true
}

// HasDuplicateCertPolicy returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasDuplicateCertPolicy() bool {
	if o != nil && o.DuplicateCertPolicy != nil {
		return true
	}

	return false
}

// SetDuplicateCertPolicy gets a reference to the given bool and assigns it to the DuplicateCertPolicy field.
func (o *PublicProfileResponse) SetDuplicateCertPolicy(v bool) {
	o.DuplicateCertPolicy = &v
}

// GetOverrideCertValidityViaApi returns the OverrideCertValidityViaApi field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetOverrideCertValidityViaApi() bool {
	if o == nil || o.OverrideCertValidityViaApi == nil {
		var ret bool
		return ret
	}
	return *o.OverrideCertValidityViaApi
}

// GetOverrideCertValidityViaApiOk returns a tuple with the OverrideCertValidityViaApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetOverrideCertValidityViaApiOk() (*bool, bool) {
	if o == nil || o.OverrideCertValidityViaApi == nil {
		return nil, false
	}
	return o.OverrideCertValidityViaApi, true
}

// HasOverrideCertValidityViaApi returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasOverrideCertValidityViaApi() bool {
	if o != nil && o.OverrideCertValidityViaApi != nil {
		return true
	}

	return false
}

// SetOverrideCertValidityViaApi gets a reference to the given bool and assigns it to the OverrideCertValidityViaApi field.
func (o *PublicProfileResponse) SetOverrideCertValidityViaApi(v bool) {
	o.OverrideCertValidityViaApi = &v
}

// GetCertificateDeliveryFormat returns the CertificateDeliveryFormat field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetCertificateDeliveryFormat() string {
	if o == nil || o.CertificateDeliveryFormat == nil {
		var ret string
		return ret
	}
	return *o.CertificateDeliveryFormat
}

// GetCertificateDeliveryFormatOk returns a tuple with the CertificateDeliveryFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetCertificateDeliveryFormatOk() (*string, bool) {
	if o == nil || o.CertificateDeliveryFormat == nil {
		return nil, false
	}
	return o.CertificateDeliveryFormat, true
}

// HasCertificateDeliveryFormat returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasCertificateDeliveryFormat() bool {
	if o != nil && o.CertificateDeliveryFormat != nil {
		return true
	}

	return false
}

// SetCertificateDeliveryFormat gets a reference to the given string and assigns it to the CertificateDeliveryFormat field.
func (o *PublicProfileResponse) SetCertificateDeliveryFormat(v string) {
	o.CertificateDeliveryFormat = &v
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetBusinessUnitId() string {
	if o == nil || o.BusinessUnitId == nil {
		var ret string
		return ret
	}
	return *o.BusinessUnitId
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetBusinessUnitIdOk() (*string, bool) {
	if o == nil || o.BusinessUnitId == nil {
		return nil, false
	}
	return o.BusinessUnitId, true
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId != nil {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given string and assigns it to the BusinessUnitId field.
func (o *PublicProfileResponse) SetBusinessUnitId(v string) {
	o.BusinessUnitId = &v
}

// GetEnrollmentMethod returns the EnrollmentMethod field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetEnrollmentMethod() ProfileEnrollmentMethod {
	if o == nil || o.EnrollmentMethod == nil {
		var ret ProfileEnrollmentMethod
		return ret
	}
	return *o.EnrollmentMethod
}

// GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetEnrollmentMethodOk() (*ProfileEnrollmentMethod, bool) {
	if o == nil || o.EnrollmentMethod == nil {
		return nil, false
	}
	return o.EnrollmentMethod, true
}

// HasEnrollmentMethod returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasEnrollmentMethod() bool {
	if o != nil && o.EnrollmentMethod != nil {
		return true
	}

	return false
}

// SetEnrollmentMethod gets a reference to the given ProfileEnrollmentMethod and assigns it to the EnrollmentMethod field.
func (o *PublicProfileResponse) SetEnrollmentMethod(v ProfileEnrollmentMethod) {
	o.EnrollmentMethod = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetAuthenticationMethod() ProfileAuthenticationMethod {
	if o == nil || o.AuthenticationMethod == nil {
		var ret ProfileAuthenticationMethod
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetAuthenticationMethodOk() (*ProfileAuthenticationMethod, bool) {
	if o == nil || o.AuthenticationMethod == nil {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasAuthenticationMethod() bool {
	if o != nil && o.AuthenticationMethod != nil {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given ProfileAuthenticationMethod and assigns it to the AuthenticationMethod field.
func (o *PublicProfileResponse) SetAuthenticationMethod(v ProfileAuthenticationMethod) {
	o.AuthenticationMethod = &v
}

// GetKeyEscrowPolicy returns the KeyEscrowPolicy field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetKeyEscrowPolicy() KeyEscrowPolicy {
	if o == nil || o.KeyEscrowPolicy == nil {
		var ret KeyEscrowPolicy
		return ret
	}
	return *o.KeyEscrowPolicy
}

// GetKeyEscrowPolicyOk returns a tuple with the KeyEscrowPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetKeyEscrowPolicyOk() (*KeyEscrowPolicy, bool) {
	if o == nil || o.KeyEscrowPolicy == nil {
		return nil, false
	}
	return o.KeyEscrowPolicy, true
}

// HasKeyEscrowPolicy returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasKeyEscrowPolicy() bool {
	if o != nil && o.KeyEscrowPolicy != nil {
		return true
	}

	return false
}

// SetKeyEscrowPolicy gets a reference to the given KeyEscrowPolicy and assigns it to the KeyEscrowPolicy field.
func (o *PublicProfileResponse) SetKeyEscrowPolicy(v KeyEscrowPolicy) {
	o.KeyEscrowPolicy = &v
}

// GetPrivateKeyAttributes returns the PrivateKeyAttributes field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetPrivateKeyAttributes() PublicProfileResponsePrivateKeyAttributes {
	if o == nil || o.PrivateKeyAttributes == nil {
		var ret PublicProfileResponsePrivateKeyAttributes
		return ret
	}
	return *o.PrivateKeyAttributes
}

// GetPrivateKeyAttributesOk returns a tuple with the PrivateKeyAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetPrivateKeyAttributesOk() (*PublicProfileResponsePrivateKeyAttributes, bool) {
	if o == nil || o.PrivateKeyAttributes == nil {
		return nil, false
	}
	return o.PrivateKeyAttributes, true
}

// HasPrivateKeyAttributes returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasPrivateKeyAttributes() bool {
	if o != nil && o.PrivateKeyAttributes != nil {
		return true
	}

	return false
}

// SetPrivateKeyAttributes gets a reference to the given PublicProfileResponsePrivateKeyAttributes and assigns it to the PrivateKeyAttributes field.
func (o *PublicProfileResponse) SetPrivateKeyAttributes(v PublicProfileResponsePrivateKeyAttributes) {
	o.PrivateKeyAttributes = &v
}

// GetApiTokenBindingEnabled returns the ApiTokenBindingEnabled field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetApiTokenBindingEnabled() bool {
	if o == nil || o.ApiTokenBindingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ApiTokenBindingEnabled
}

// GetApiTokenBindingEnabledOk returns a tuple with the ApiTokenBindingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetApiTokenBindingEnabledOk() (*bool, bool) {
	if o == nil || o.ApiTokenBindingEnabled == nil {
		return nil, false
	}
	return o.ApiTokenBindingEnabled, true
}

// HasApiTokenBindingEnabled returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasApiTokenBindingEnabled() bool {
	if o != nil && o.ApiTokenBindingEnabled != nil {
		return true
	}

	return false
}

// SetApiTokenBindingEnabled gets a reference to the given bool and assigns it to the ApiTokenBindingEnabled field.
func (o *PublicProfileResponse) SetApiTokenBindingEnabled(v bool) {
	o.ApiTokenBindingEnabled = &v
}

// GetApiTokenId returns the ApiTokenId field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetApiTokenId() string {
	if o == nil || o.ApiTokenId == nil {
		var ret string
		return ret
	}
	return *o.ApiTokenId
}

// GetApiTokenIdOk returns a tuple with the ApiTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetApiTokenIdOk() (*string, bool) {
	if o == nil || o.ApiTokenId == nil {
		return nil, false
	}
	return o.ApiTokenId, true
}

// HasApiTokenId returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasApiTokenId() bool {
	if o != nil && o.ApiTokenId != nil {
		return true
	}

	return false
}

// SetApiTokenId gets a reference to the given string and assigns it to the ApiTokenId field.
func (o *PublicProfileResponse) SetApiTokenId(v string) {
	o.ApiTokenId = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *PublicProfileResponse) GetCertificate() PublicProfileResponseCertificate {
	if o == nil || o.Certificate == nil {
		var ret PublicProfileResponseCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfileResponse) GetCertificateOk() (*PublicProfileResponseCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *PublicProfileResponse) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given PublicProfileResponseCertificate and assigns it to the Certificate field.
func (o *PublicProfileResponse) SetCertificate(v PublicProfileResponseCertificate) {
	o.Certificate = &v
}

func (o PublicProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.PublishToPublicDirectory != nil {
		toSerialize["publish_to_public_directory"] = o.PublishToPublicDirectory
	}
	if o.RenewalPeriodDays != nil {
		toSerialize["renewal_period_days"] = o.RenewalPeriodDays
	}
	if o.DuplicateCertPolicy != nil {
		toSerialize["duplicate_cert_policy"] = o.DuplicateCertPolicy
	}
	if o.OverrideCertValidityViaApi != nil {
		toSerialize["override_cert_validity_via_api"] = o.OverrideCertValidityViaApi
	}
	if o.CertificateDeliveryFormat != nil {
		toSerialize["certificate_delivery_format"] = o.CertificateDeliveryFormat
	}
	if o.BusinessUnitId != nil {
		toSerialize["business_unit_id"] = o.BusinessUnitId
	}
	if o.EnrollmentMethod != nil {
		toSerialize["enrollment_method"] = o.EnrollmentMethod
	}
	if o.AuthenticationMethod != nil {
		toSerialize["authentication_method"] = o.AuthenticationMethod
	}
	if o.KeyEscrowPolicy != nil {
		toSerialize["key_escrow_policy"] = o.KeyEscrowPolicy
	}
	if o.PrivateKeyAttributes != nil {
		toSerialize["private_key_attributes"] = o.PrivateKeyAttributes
	}
	if o.ApiTokenBindingEnabled != nil {
		toSerialize["api_token_binding_enabled"] = o.ApiTokenBindingEnabled
	}
	if o.ApiTokenId != nil {
		toSerialize["api_token_id"] = o.ApiTokenId
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullablePublicProfileResponse struct {
	value *PublicProfileResponse
	isSet bool
}

func (v NullablePublicProfileResponse) Get() *PublicProfileResponse {
	return v.value
}

func (v *NullablePublicProfileResponse) Set(val *PublicProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicProfileResponse(val *PublicProfileResponse) *NullablePublicProfileResponse {
	return &NullablePublicProfileResponse{value: val, isSet: true}
}

func (v NullablePublicProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


