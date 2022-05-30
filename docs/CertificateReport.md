# CertificateReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **string** | Entity ID in UUID format | [optional] 
**Status** | Pointer to **string** | Current status of Certificate | [optional] 
**SeatType** | Pointer to **string** | Type of Seat for which Certificate was issued (User, Server, Device, etc) | [optional] 
**Issuer** | Pointer to **string** | ID of issuing CA | [optional] 
**Root** | Pointer to **string** | ID of root CA | [optional] 
**RevocationDate** | Pointer to **string** | Certificate revocation date in format YYYY-MM-DDTHH-MM-SS | [optional] 
**Attributes** | Pointer to **map[string]string** | Certificate attributes as key-value map | [optional] 

## Methods

### NewCertificateReport

`func NewCertificateReport() *CertificateReport`

NewCertificateReport instantiates a new CertificateReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateReportWithDefaults

`func NewCertificateReportWithDefaults() *CertificateReport`

NewCertificateReportWithDefaults instantiates a new CertificateReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *CertificateReport) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CertificateReport) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CertificateReport) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CertificateReport) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeatType

`func (o *CertificateReport) GetSeatType() string`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *CertificateReport) GetSeatTypeOk() (*string, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *CertificateReport) SetSeatType(v string)`

SetSeatType sets SeatType field to given value.

### HasSeatType

`func (o *CertificateReport) HasSeatType() bool`

HasSeatType returns a boolean if a field has been set.

### GetIssuer

`func (o *CertificateReport) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateReport) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateReport) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateReport) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRoot

`func (o *CertificateReport) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CertificateReport) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CertificateReport) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *CertificateReport) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetRevocationDate

`func (o *CertificateReport) GetRevocationDate() string`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *CertificateReport) GetRevocationDateOk() (*string, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *CertificateReport) SetRevocationDate(v string)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *CertificateReport) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### GetAttributes

`func (o *CertificateReport) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CertificateReport) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CertificateReport) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CertificateReport) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


