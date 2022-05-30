# MpkiApiV1CertificateImportPostRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **string** |  | [optional] 
**SeatId** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to ***os.File** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Revocation** | Pointer to [**MpkiApiV1CertificateImportPostRequestInnerRevocation**](MpkiApiV1CertificateImportPostRequestInnerRevocation.md) |  | [optional] 

## Methods

### NewMpkiApiV1CertificateImportPostRequestInner

`func NewMpkiApiV1CertificateImportPostRequestInner() *MpkiApiV1CertificateImportPostRequestInner`

NewMpkiApiV1CertificateImportPostRequestInner instantiates a new MpkiApiV1CertificateImportPostRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1CertificateImportPostRequestInnerWithDefaults

`func NewMpkiApiV1CertificateImportPostRequestInnerWithDefaults() *MpkiApiV1CertificateImportPostRequestInner`

NewMpkiApiV1CertificateImportPostRequestInnerWithDefaults instantiates a new MpkiApiV1CertificateImportPostRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetSeatId

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetSeatId() string`

GetSeatId returns the SeatId field if non-nil, zero value otherwise.

### GetSeatIdOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetSeatIdOk() (*string, bool)`

GetSeatIdOk returns a tuple with the SeatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatId

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetSeatId(v string)`

SetSeatId sets SeatId field to given value.

### HasSeatId

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasSeatId() bool`

HasSeatId returns a boolean if a field has been set.

### GetCert

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetCert() *os.File`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetCertOk() (**os.File, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetCert(v *os.File)`

SetCert sets Cert field to given value.

### HasCert

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetNotes

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetGroupName

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetRevocation

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetRevocation() MpkiApiV1CertificateImportPostRequestInnerRevocation`

GetRevocation returns the Revocation field if non-nil, zero value otherwise.

### GetRevocationOk

`func (o *MpkiApiV1CertificateImportPostRequestInner) GetRevocationOk() (*MpkiApiV1CertificateImportPostRequestInnerRevocation, bool)`

GetRevocationOk returns a tuple with the Revocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocation

`func (o *MpkiApiV1CertificateImportPostRequestInner) SetRevocation(v MpkiApiV1CertificateImportPostRequestInnerRevocation)`

SetRevocation sets Revocation field to given value.

### HasRevocation

`func (o *MpkiApiV1CertificateImportPostRequestInner) HasRevocation() bool`

HasRevocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


