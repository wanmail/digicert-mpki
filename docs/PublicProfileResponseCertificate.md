# PublicProfileResponseCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**PublicProfileResponseCertificateSubject**](PublicProfileResponseCertificateSubject.md) |  | [optional] 
**Validity** | Pointer to [**Validity**](Validity.md) |  | [optional] 
**Extensions** | Pointer to [**PublicProfileResponseCertificateExtensions**](PublicProfileResponseCertificateExtensions.md) |  | [optional] 
**Issuer** | Pointer to [**PublicProfileResponseCertificateIssuer**](PublicProfileResponseCertificateIssuer.md) |  | [optional] 

## Methods

### NewPublicProfileResponseCertificate

`func NewPublicProfileResponseCertificate() *PublicProfileResponseCertificate`

NewPublicProfileResponseCertificate instantiates a new PublicProfileResponseCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProfileResponseCertificateWithDefaults

`func NewPublicProfileResponseCertificateWithDefaults() *PublicProfileResponseCertificate`

NewPublicProfileResponseCertificateWithDefaults instantiates a new PublicProfileResponseCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PublicProfileResponseCertificate) GetSubject() PublicProfileResponseCertificateSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicProfileResponseCertificate) GetSubjectOk() (*PublicProfileResponseCertificateSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicProfileResponseCertificate) SetSubject(v PublicProfileResponseCertificateSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicProfileResponseCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidity

`func (o *PublicProfileResponseCertificate) GetValidity() Validity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *PublicProfileResponseCertificate) GetValidityOk() (*Validity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *PublicProfileResponseCertificate) SetValidity(v Validity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *PublicProfileResponseCertificate) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetExtensions

`func (o *PublicProfileResponseCertificate) GetExtensions() PublicProfileResponseCertificateExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *PublicProfileResponseCertificate) GetExtensionsOk() (*PublicProfileResponseCertificateExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *PublicProfileResponseCertificate) SetExtensions(v PublicProfileResponseCertificateExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *PublicProfileResponseCertificate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetIssuer

`func (o *PublicProfileResponseCertificate) GetIssuer() PublicProfileResponseCertificateIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PublicProfileResponseCertificate) GetIssuerOk() (*PublicProfileResponseCertificateIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PublicProfileResponseCertificate) SetIssuer(v PublicProfileResponseCertificateIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PublicProfileResponseCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


