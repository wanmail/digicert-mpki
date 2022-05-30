# CaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** | Issuer serial number | [optional] 
**SubjectDn** | Pointer to **string** | Issuer subject dn | [optional] 
**Certificate** | Pointer to **string** | Base-64 encoded issuer certificate | [optional] 
**Root** | Pointer to **bool** | Flag indicating whether its a root certificate | [optional] 
**Chain** | Pointer to [**[]CaResponse**](CaResponse.md) |  | [optional] 

## Methods

### NewCaResponse

`func NewCaResponse() *CaResponse`

NewCaResponse instantiates a new CaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaResponseWithDefaults

`func NewCaResponseWithDefaults() *CaResponse`

NewCaResponseWithDefaults instantiates a new CaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *CaResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CaResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CaResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CaResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubjectDn

`func (o *CaResponse) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *CaResponse) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *CaResponse) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.

### HasSubjectDn

`func (o *CaResponse) HasSubjectDn() bool`

HasSubjectDn returns a boolean if a field has been set.

### GetCertificate

`func (o *CaResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CaResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CaResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CaResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetRoot

`func (o *CaResponse) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CaResponse) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CaResponse) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *CaResponse) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetChain

`func (o *CaResponse) GetChain() []CaResponse`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CaResponse) GetChainOk() (*[]CaResponse, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CaResponse) SetChain(v []CaResponse)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CaResponse) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


