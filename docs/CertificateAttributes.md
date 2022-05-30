# CertificateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**PublicProfileResponseCertificateSubject**](PublicProfileResponseCertificateSubject.md) |  | [optional] 
**Extensions** | Pointer to [**CertificateExtensions**](CertificateExtensions.md) |  | [optional] 

## Methods

### NewCertificateAttributes

`func NewCertificateAttributes() *CertificateAttributes`

NewCertificateAttributes instantiates a new CertificateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAttributesWithDefaults

`func NewCertificateAttributesWithDefaults() *CertificateAttributes`

NewCertificateAttributesWithDefaults instantiates a new CertificateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CertificateAttributes) GetSubject() PublicProfileResponseCertificateSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateAttributes) GetSubjectOk() (*PublicProfileResponseCertificateSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateAttributes) SetSubject(v PublicProfileResponseCertificateSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateAttributes) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetExtensions

`func (o *CertificateAttributes) GetExtensions() CertificateExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CertificateAttributes) GetExtensionsOk() (*CertificateExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CertificateAttributes) SetExtensions(v CertificateExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CertificateAttributes) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


