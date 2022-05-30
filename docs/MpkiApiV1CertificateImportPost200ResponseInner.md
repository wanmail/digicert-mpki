# MpkiApiV1CertificateImportPost200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** |  | [optional] 
**IssuingCa** | Pointer to **string** |  | [optional] 
**ImportStatus** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Seat** | Pointer to [**MpkiApiV1CertificateImportPost200ResponseInnerSeat**](MpkiApiV1CertificateImportPost200ResponseInnerSeat.md) |  | [optional] 
**BusinessUnitId** | Pointer to [**MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId**](MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId.md) |  | [optional] 
**CertStatus** | Pointer to **string** |  | [optional] 
**Revocation** | Pointer to [**MpkiApiV1CertificateImportPostRequestInnerRevocation**](MpkiApiV1CertificateImportPostRequestInnerRevocation.md) |  | [optional] 

## Methods

### NewMpkiApiV1CertificateImportPost200ResponseInner

`func NewMpkiApiV1CertificateImportPost200ResponseInner() *MpkiApiV1CertificateImportPost200ResponseInner`

NewMpkiApiV1CertificateImportPost200ResponseInner instantiates a new MpkiApiV1CertificateImportPost200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1CertificateImportPost200ResponseInnerWithDefaults

`func NewMpkiApiV1CertificateImportPost200ResponseInnerWithDefaults() *MpkiApiV1CertificateImportPost200ResponseInner`

NewMpkiApiV1CertificateImportPost200ResponseInnerWithDefaults instantiates a new MpkiApiV1CertificateImportPost200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetIssuingCa

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetIssuingCa() string`

GetIssuingCa returns the IssuingCa field if non-nil, zero value otherwise.

### GetIssuingCaOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetIssuingCaOk() (*string, bool)`

GetIssuingCaOk returns a tuple with the IssuingCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCa

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetIssuingCa(v string)`

SetIssuingCa sets IssuingCa field to given value.

### HasIssuingCa

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasIssuingCa() bool`

HasIssuingCa returns a boolean if a field has been set.

### GetImportStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetImportStatus() string`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetImportStatusOk() (*string, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetImportStatus(v string)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSeat

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetSeat() MpkiApiV1CertificateImportPost200ResponseInnerSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetSeatOk() (*MpkiApiV1CertificateImportPost200ResponseInnerSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetSeat(v MpkiApiV1CertificateImportPost200ResponseInnerSeat)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetBusinessUnitId() MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetBusinessUnitIdOk() (*MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetBusinessUnitId(v MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetCertStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertStatus() string`

GetCertStatus returns the CertStatus field if non-nil, zero value otherwise.

### GetCertStatusOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetCertStatusOk() (*string, bool)`

GetCertStatusOk returns a tuple with the CertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetCertStatus(v string)`

SetCertStatus sets CertStatus field to given value.

### HasCertStatus

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasCertStatus() bool`

HasCertStatus returns a boolean if a field has been set.

### GetRevocation

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetRevocation() MpkiApiV1CertificateImportPostRequestInnerRevocation`

GetRevocation returns the Revocation field if non-nil, zero value otherwise.

### GetRevocationOk

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) GetRevocationOk() (*MpkiApiV1CertificateImportPostRequestInnerRevocation, bool)`

GetRevocationOk returns a tuple with the Revocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocation

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) SetRevocation(v MpkiApiV1CertificateImportPostRequestInnerRevocation)`

SetRevocation sets Revocation field to given value.

### HasRevocation

`func (o *MpkiApiV1CertificateImportPost200ResponseInner) HasRevocation() bool`

HasRevocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


