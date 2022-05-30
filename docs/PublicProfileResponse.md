# PublicProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Entity ID in UUID format | [optional] 
**Name** | Pointer to **string** | Profile name | [optional] 
**Status** | Pointer to **string** | Profile status | [optional] 
**SignatureAlgorithm** | Pointer to **string** | Signature algorithm for certificates issued from the profile | [optional] 
**PublishToPublicDirectory** | Pointer to **bool** | If true, certificate public key is published to the DigiCert PKI directory | [optional] 
**RenewalPeriodDays** | Pointer to **int32** | Number of days before expiration in which the certificate can be renewed | [optional] 
**DuplicateCertPolicy** | Pointer to **bool** | If true, certificate profile allows duplicate certificates | [optional] 
**OverrideCertValidityViaApi** | Pointer to **bool** | If true, the default validity period set by the profile can be overridden using a request parameter | [optional] 
**CertificateDeliveryFormat** | Pointer to **string** | Format for certificate delivery specified in the profile | [optional] 
**BusinessUnitId** | Pointer to **string** | Business unit ID to which the profile belongs | [optional] 
**EnrollmentMethod** | Pointer to [**ProfileEnrollmentMethod**](ProfileEnrollmentMethod.md) |  | [optional] 
**AuthenticationMethod** | Pointer to [**ProfileAuthenticationMethod**](ProfileAuthenticationMethod.md) |  | [optional] 
**KeyEscrowPolicy** | Pointer to [**KeyEscrowPolicy**](KeyEscrowPolicy.md) |  | [optional] 
**PrivateKeyAttributes** | Pointer to [**PublicProfileResponsePrivateKeyAttributes**](PublicProfileResponsePrivateKeyAttributes.md) |  | [optional] 
**ApiTokenBindingEnabled** | Pointer to **bool** | If true, use of the certificate profile is restricted to a specific API token | [optional] 
**ApiTokenId** | Pointer to **string** | API token ID to which the profile is bound if api_token_binding_enabled is true | [optional] 
**Certificate** | Pointer to [**PublicProfileResponseCertificate**](PublicProfileResponseCertificate.md) |  | [optional] 

## Methods

### NewPublicProfileResponse

`func NewPublicProfileResponse() *PublicProfileResponse`

NewPublicProfileResponse instantiates a new PublicProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProfileResponseWithDefaults

`func NewPublicProfileResponseWithDefaults() *PublicProfileResponse`

NewPublicProfileResponseWithDefaults instantiates a new PublicProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicProfileResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicProfileResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PublicProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PublicProfileResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicProfileResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicProfileResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicProfileResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *PublicProfileResponse) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *PublicProfileResponse) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *PublicProfileResponse) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *PublicProfileResponse) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetPublishToPublicDirectory

`func (o *PublicProfileResponse) GetPublishToPublicDirectory() bool`

GetPublishToPublicDirectory returns the PublishToPublicDirectory field if non-nil, zero value otherwise.

### GetPublishToPublicDirectoryOk

`func (o *PublicProfileResponse) GetPublishToPublicDirectoryOk() (*bool, bool)`

GetPublishToPublicDirectoryOk returns a tuple with the PublishToPublicDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToPublicDirectory

`func (o *PublicProfileResponse) SetPublishToPublicDirectory(v bool)`

SetPublishToPublicDirectory sets PublishToPublicDirectory field to given value.

### HasPublishToPublicDirectory

`func (o *PublicProfileResponse) HasPublishToPublicDirectory() bool`

HasPublishToPublicDirectory returns a boolean if a field has been set.

### GetRenewalPeriodDays

`func (o *PublicProfileResponse) GetRenewalPeriodDays() int32`

GetRenewalPeriodDays returns the RenewalPeriodDays field if non-nil, zero value otherwise.

### GetRenewalPeriodDaysOk

`func (o *PublicProfileResponse) GetRenewalPeriodDaysOk() (*int32, bool)`

GetRenewalPeriodDaysOk returns a tuple with the RenewalPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriodDays

`func (o *PublicProfileResponse) SetRenewalPeriodDays(v int32)`

SetRenewalPeriodDays sets RenewalPeriodDays field to given value.

### HasRenewalPeriodDays

`func (o *PublicProfileResponse) HasRenewalPeriodDays() bool`

HasRenewalPeriodDays returns a boolean if a field has been set.

### GetDuplicateCertPolicy

`func (o *PublicProfileResponse) GetDuplicateCertPolicy() bool`

GetDuplicateCertPolicy returns the DuplicateCertPolicy field if non-nil, zero value otherwise.

### GetDuplicateCertPolicyOk

`func (o *PublicProfileResponse) GetDuplicateCertPolicyOk() (*bool, bool)`

GetDuplicateCertPolicyOk returns a tuple with the DuplicateCertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateCertPolicy

`func (o *PublicProfileResponse) SetDuplicateCertPolicy(v bool)`

SetDuplicateCertPolicy sets DuplicateCertPolicy field to given value.

### HasDuplicateCertPolicy

`func (o *PublicProfileResponse) HasDuplicateCertPolicy() bool`

HasDuplicateCertPolicy returns a boolean if a field has been set.

### GetOverrideCertValidityViaApi

`func (o *PublicProfileResponse) GetOverrideCertValidityViaApi() bool`

GetOverrideCertValidityViaApi returns the OverrideCertValidityViaApi field if non-nil, zero value otherwise.

### GetOverrideCertValidityViaApiOk

`func (o *PublicProfileResponse) GetOverrideCertValidityViaApiOk() (*bool, bool)`

GetOverrideCertValidityViaApiOk returns a tuple with the OverrideCertValidityViaApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCertValidityViaApi

`func (o *PublicProfileResponse) SetOverrideCertValidityViaApi(v bool)`

SetOverrideCertValidityViaApi sets OverrideCertValidityViaApi field to given value.

### HasOverrideCertValidityViaApi

`func (o *PublicProfileResponse) HasOverrideCertValidityViaApi() bool`

HasOverrideCertValidityViaApi returns a boolean if a field has been set.

### GetCertificateDeliveryFormat

`func (o *PublicProfileResponse) GetCertificateDeliveryFormat() string`

GetCertificateDeliveryFormat returns the CertificateDeliveryFormat field if non-nil, zero value otherwise.

### GetCertificateDeliveryFormatOk

`func (o *PublicProfileResponse) GetCertificateDeliveryFormatOk() (*string, bool)`

GetCertificateDeliveryFormatOk returns a tuple with the CertificateDeliveryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateDeliveryFormat

`func (o *PublicProfileResponse) SetCertificateDeliveryFormat(v string)`

SetCertificateDeliveryFormat sets CertificateDeliveryFormat field to given value.

### HasCertificateDeliveryFormat

`func (o *PublicProfileResponse) HasCertificateDeliveryFormat() bool`

HasCertificateDeliveryFormat returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *PublicProfileResponse) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicProfileResponse) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicProfileResponse) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicProfileResponse) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetEnrollmentMethod

`func (o *PublicProfileResponse) GetEnrollmentMethod() ProfileEnrollmentMethod`

GetEnrollmentMethod returns the EnrollmentMethod field if non-nil, zero value otherwise.

### GetEnrollmentMethodOk

`func (o *PublicProfileResponse) GetEnrollmentMethodOk() (*ProfileEnrollmentMethod, bool)`

GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethod

`func (o *PublicProfileResponse) SetEnrollmentMethod(v ProfileEnrollmentMethod)`

SetEnrollmentMethod sets EnrollmentMethod field to given value.

### HasEnrollmentMethod

`func (o *PublicProfileResponse) HasEnrollmentMethod() bool`

HasEnrollmentMethod returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *PublicProfileResponse) GetAuthenticationMethod() ProfileAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PublicProfileResponse) GetAuthenticationMethodOk() (*ProfileAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PublicProfileResponse) SetAuthenticationMethod(v ProfileAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *PublicProfileResponse) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetKeyEscrowPolicy

`func (o *PublicProfileResponse) GetKeyEscrowPolicy() KeyEscrowPolicy`

GetKeyEscrowPolicy returns the KeyEscrowPolicy field if non-nil, zero value otherwise.

### GetKeyEscrowPolicyOk

`func (o *PublicProfileResponse) GetKeyEscrowPolicyOk() (*KeyEscrowPolicy, bool)`

GetKeyEscrowPolicyOk returns a tuple with the KeyEscrowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEscrowPolicy

`func (o *PublicProfileResponse) SetKeyEscrowPolicy(v KeyEscrowPolicy)`

SetKeyEscrowPolicy sets KeyEscrowPolicy field to given value.

### HasKeyEscrowPolicy

`func (o *PublicProfileResponse) HasKeyEscrowPolicy() bool`

HasKeyEscrowPolicy returns a boolean if a field has been set.

### GetPrivateKeyAttributes

`func (o *PublicProfileResponse) GetPrivateKeyAttributes() PublicProfileResponsePrivateKeyAttributes`

GetPrivateKeyAttributes returns the PrivateKeyAttributes field if non-nil, zero value otherwise.

### GetPrivateKeyAttributesOk

`func (o *PublicProfileResponse) GetPrivateKeyAttributesOk() (*PublicProfileResponsePrivateKeyAttributes, bool)`

GetPrivateKeyAttributesOk returns a tuple with the PrivateKeyAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAttributes

`func (o *PublicProfileResponse) SetPrivateKeyAttributes(v PublicProfileResponsePrivateKeyAttributes)`

SetPrivateKeyAttributes sets PrivateKeyAttributes field to given value.

### HasPrivateKeyAttributes

`func (o *PublicProfileResponse) HasPrivateKeyAttributes() bool`

HasPrivateKeyAttributes returns a boolean if a field has been set.

### GetApiTokenBindingEnabled

`func (o *PublicProfileResponse) GetApiTokenBindingEnabled() bool`

GetApiTokenBindingEnabled returns the ApiTokenBindingEnabled field if non-nil, zero value otherwise.

### GetApiTokenBindingEnabledOk

`func (o *PublicProfileResponse) GetApiTokenBindingEnabledOk() (*bool, bool)`

GetApiTokenBindingEnabledOk returns a tuple with the ApiTokenBindingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenBindingEnabled

`func (o *PublicProfileResponse) SetApiTokenBindingEnabled(v bool)`

SetApiTokenBindingEnabled sets ApiTokenBindingEnabled field to given value.

### HasApiTokenBindingEnabled

`func (o *PublicProfileResponse) HasApiTokenBindingEnabled() bool`

HasApiTokenBindingEnabled returns a boolean if a field has been set.

### GetApiTokenId

`func (o *PublicProfileResponse) GetApiTokenId() string`

GetApiTokenId returns the ApiTokenId field if non-nil, zero value otherwise.

### GetApiTokenIdOk

`func (o *PublicProfileResponse) GetApiTokenIdOk() (*string, bool)`

GetApiTokenIdOk returns a tuple with the ApiTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenId

`func (o *PublicProfileResponse) SetApiTokenId(v string)`

SetApiTokenId sets ApiTokenId field to given value.

### HasApiTokenId

`func (o *PublicProfileResponse) HasApiTokenId() bool`

HasApiTokenId returns a boolean if a field has been set.

### GetCertificate

`func (o *PublicProfileResponse) GetCertificate() PublicProfileResponseCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PublicProfileResponse) GetCertificateOk() (*PublicProfileResponseCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PublicProfileResponse) SetCertificate(v PublicProfileResponseCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PublicProfileResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


