# CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**DeliveryFormat** | Pointer to [**DeliveryFormat**](DeliveryFormat.md) |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**Pkcs12Password** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateResponse

`func NewCertificateResponse() *CertificateResponse`

NewCertificateResponse instantiates a new CertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateResponseWithDefaults

`func NewCertificateResponseWithDefaults() *CertificateResponse`

NewCertificateResponseWithDefaults instantiates a new CertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *CertificateResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDeliveryFormat

`func (o *CertificateResponse) GetDeliveryFormat() DeliveryFormat`

GetDeliveryFormat returns the DeliveryFormat field if non-nil, zero value otherwise.

### GetDeliveryFormatOk

`func (o *CertificateResponse) GetDeliveryFormatOk() (*DeliveryFormat, bool)`

GetDeliveryFormatOk returns a tuple with the DeliveryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFormat

`func (o *CertificateResponse) SetDeliveryFormat(v DeliveryFormat)`

SetDeliveryFormat sets DeliveryFormat field to given value.

### HasDeliveryFormat

`func (o *CertificateResponse) HasDeliveryFormat() bool`

HasDeliveryFormat returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPkcs12Password

`func (o *CertificateResponse) GetPkcs12Password() string`

GetPkcs12Password returns the Pkcs12Password field if non-nil, zero value otherwise.

### GetPkcs12PasswordOk

`func (o *CertificateResponse) GetPkcs12PasswordOk() (*string, bool)`

GetPkcs12PasswordOk returns a tuple with the Pkcs12Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12Password

`func (o *CertificateResponse) SetPkcs12Password(v string)`

SetPkcs12Password sets Pkcs12Password field to given value.

### HasPkcs12Password

`func (o *CertificateResponse) HasPkcs12Password() bool`

HasPkcs12Password returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


