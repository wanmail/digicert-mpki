# PublicProfileResponseCertificateIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** | Issuer serial number | [optional] 
**SubjectDn** | Pointer to **string** | Issuer subject dn | [optional] 
**Certificate** | Pointer to **string** | Base-64 encoded issuer certificate | [optional] 
**Root** | Pointer to **bool** | Flag indicating whether its a root certificate | [optional] 
**Chain** | Pointer to [**[]CaResponse**](CaResponse.md) |  | [optional] 

## Methods

### NewPublicProfileResponseCertificateIssuer

`func NewPublicProfileResponseCertificateIssuer() *PublicProfileResponseCertificateIssuer`

NewPublicProfileResponseCertificateIssuer instantiates a new PublicProfileResponseCertificateIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProfileResponseCertificateIssuerWithDefaults

`func NewPublicProfileResponseCertificateIssuerWithDefaults() *PublicProfileResponseCertificateIssuer`

NewPublicProfileResponseCertificateIssuerWithDefaults instantiates a new PublicProfileResponseCertificateIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *PublicProfileResponseCertificateIssuer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PublicProfileResponseCertificateIssuer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PublicProfileResponseCertificateIssuer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PublicProfileResponseCertificateIssuer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubjectDn

`func (o *PublicProfileResponseCertificateIssuer) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *PublicProfileResponseCertificateIssuer) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *PublicProfileResponseCertificateIssuer) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.

### HasSubjectDn

`func (o *PublicProfileResponseCertificateIssuer) HasSubjectDn() bool`

HasSubjectDn returns a boolean if a field has been set.

### GetCertificate

`func (o *PublicProfileResponseCertificateIssuer) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PublicProfileResponseCertificateIssuer) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PublicProfileResponseCertificateIssuer) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PublicProfileResponseCertificateIssuer) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetRoot

`func (o *PublicProfileResponseCertificateIssuer) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *PublicProfileResponseCertificateIssuer) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *PublicProfileResponseCertificateIssuer) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *PublicProfileResponseCertificateIssuer) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetChain

`func (o *PublicProfileResponseCertificateIssuer) GetChain() []CaResponse`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *PublicProfileResponseCertificateIssuer) GetChainOk() (*[]CaResponse, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *PublicProfileResponseCertificateIssuer) SetChain(v []CaResponse)`

SetChain sets Chain field to given value.

### HasChain

`func (o *PublicProfileResponseCertificateIssuer) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


