# CertificateExtensionsSan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **bool** |  | [optional] [default to false]
**Attributes** | Pointer to [**[]PublicProfileField**](PublicProfileField.md) |  | [optional] 

## Methods

### NewCertificateExtensionsSan

`func NewCertificateExtensionsSan() *CertificateExtensionsSan`

NewCertificateExtensionsSan instantiates a new CertificateExtensionsSan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateExtensionsSanWithDefaults

`func NewCertificateExtensionsSanWithDefaults() *CertificateExtensionsSan`

NewCertificateExtensionsSanWithDefaults instantiates a new CertificateExtensionsSan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *CertificateExtensionsSan) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *CertificateExtensionsSan) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *CertificateExtensionsSan) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *CertificateExtensionsSan) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetAttributes

`func (o *CertificateExtensionsSan) GetAttributes() []PublicProfileField`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CertificateExtensionsSan) GetAttributesOk() (*[]PublicProfileField, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CertificateExtensionsSan) SetAttributes(v []PublicProfileField)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CertificateExtensionsSan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


