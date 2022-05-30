# MpkiApiV1CertificateSerialNumberRenewPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csr** | Pointer to **string** | A CSR containing a new or existing public key | [optional] 
**Validity** | Pointer to [**Validity**](Validity.md) |  | [optional] 
**DeliveryFormat** | Pointer to [**DeliveryFormat**](DeliveryFormat.md) |  | [optional] 
**IncludeCaChain** | Pointer to **bool** |  | [optional] 

## Methods

### NewMpkiApiV1CertificateSerialNumberRenewPostRequest

`func NewMpkiApiV1CertificateSerialNumberRenewPostRequest() *MpkiApiV1CertificateSerialNumberRenewPostRequest`

NewMpkiApiV1CertificateSerialNumberRenewPostRequest instantiates a new MpkiApiV1CertificateSerialNumberRenewPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1CertificateSerialNumberRenewPostRequestWithDefaults

`func NewMpkiApiV1CertificateSerialNumberRenewPostRequestWithDefaults() *MpkiApiV1CertificateSerialNumberRenewPostRequest`

NewMpkiApiV1CertificateSerialNumberRenewPostRequestWithDefaults instantiates a new MpkiApiV1CertificateSerialNumberRenewPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsr

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetValidity

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetValidity() Validity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetValidityOk() (*Validity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) SetValidity(v Validity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetDeliveryFormat

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetDeliveryFormat() DeliveryFormat`

GetDeliveryFormat returns the DeliveryFormat field if non-nil, zero value otherwise.

### GetDeliveryFormatOk

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetDeliveryFormatOk() (*DeliveryFormat, bool)`

GetDeliveryFormatOk returns a tuple with the DeliveryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFormat

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) SetDeliveryFormat(v DeliveryFormat)`

SetDeliveryFormat sets DeliveryFormat field to given value.

### HasDeliveryFormat

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) HasDeliveryFormat() bool`

HasDeliveryFormat returns a boolean if a field has been set.

### GetIncludeCaChain

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetIncludeCaChain() bool`

GetIncludeCaChain returns the IncludeCaChain field if non-nil, zero value otherwise.

### GetIncludeCaChainOk

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) GetIncludeCaChainOk() (*bool, bool)`

GetIncludeCaChainOk returns a tuple with the IncludeCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCaChain

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) SetIncludeCaChain(v bool)`

SetIncludeCaChain sets IncludeCaChain field to given value.

### HasIncludeCaChain

`func (o *MpkiApiV1CertificateSerialNumberRenewPostRequest) HasIncludeCaChain() bool`

HasIncludeCaChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


