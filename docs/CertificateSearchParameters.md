# CertificateSearchParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **string** |  | [optional] 
**BusinessUnitName** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**SeatId** | Pointer to **string** |  | [optional] 
**SeatTypeId** | Pointer to **string** |  | [optional] 
**ValidFrom** | Pointer to **string** | Limit results to certificates issued after specified date value | [optional] 
**ValidTo** | Pointer to **string** | Limit results to certificates expiring before specified date value | [optional] 
**SerialNumber** | Pointer to ***os.File** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Aki** | Pointer to **string** | Certificate Authority Key Identifier | [optional] 
**IssuingCaId** | Pointer to **string** |  | [optional] 
**IssuingCaName** | Pointer to **string** |  | [optional] 
**IssuingCaSerialNumber** | Pointer to ***os.File** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Thumbprint** | Pointer to ***os.File** | SHA256 thumbprint of the certificate | [optional] 

## Methods

### NewCertificateSearchParameters

`func NewCertificateSearchParameters() *CertificateSearchParameters`

NewCertificateSearchParameters instantiates a new CertificateSearchParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearchParametersWithDefaults

`func NewCertificateSearchParametersWithDefaults() *CertificateSearchParameters`

NewCertificateSearchParametersWithDefaults instantiates a new CertificateSearchParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CertificateSearchParameters) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CertificateSearchParameters) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CertificateSearchParameters) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CertificateSearchParameters) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *CertificateSearchParameters) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *CertificateSearchParameters) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *CertificateSearchParameters) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *CertificateSearchParameters) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetBusinessUnitName

`func (o *CertificateSearchParameters) GetBusinessUnitName() string`

GetBusinessUnitName returns the BusinessUnitName field if non-nil, zero value otherwise.

### GetBusinessUnitNameOk

`func (o *CertificateSearchParameters) GetBusinessUnitNameOk() (*string, bool)`

GetBusinessUnitNameOk returns a tuple with the BusinessUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitName

`func (o *CertificateSearchParameters) SetBusinessUnitName(v string)`

SetBusinessUnitName sets BusinessUnitName field to given value.

### HasBusinessUnitName

`func (o *CertificateSearchParameters) HasBusinessUnitName() bool`

HasBusinessUnitName returns a boolean if a field has been set.

### GetProfileId

`func (o *CertificateSearchParameters) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CertificateSearchParameters) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CertificateSearchParameters) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *CertificateSearchParameters) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetSeatId

`func (o *CertificateSearchParameters) GetSeatId() string`

GetSeatId returns the SeatId field if non-nil, zero value otherwise.

### GetSeatIdOk

`func (o *CertificateSearchParameters) GetSeatIdOk() (*string, bool)`

GetSeatIdOk returns a tuple with the SeatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatId

`func (o *CertificateSearchParameters) SetSeatId(v string)`

SetSeatId sets SeatId field to given value.

### HasSeatId

`func (o *CertificateSearchParameters) HasSeatId() bool`

HasSeatId returns a boolean if a field has been set.

### GetSeatTypeId

`func (o *CertificateSearchParameters) GetSeatTypeId() string`

GetSeatTypeId returns the SeatTypeId field if non-nil, zero value otherwise.

### GetSeatTypeIdOk

`func (o *CertificateSearchParameters) GetSeatTypeIdOk() (*string, bool)`

GetSeatTypeIdOk returns a tuple with the SeatTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatTypeId

`func (o *CertificateSearchParameters) SetSeatTypeId(v string)`

SetSeatTypeId sets SeatTypeId field to given value.

### HasSeatTypeId

`func (o *CertificateSearchParameters) HasSeatTypeId() bool`

HasSeatTypeId returns a boolean if a field has been set.

### GetValidFrom

`func (o *CertificateSearchParameters) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CertificateSearchParameters) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CertificateSearchParameters) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *CertificateSearchParameters) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *CertificateSearchParameters) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CertificateSearchParameters) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CertificateSearchParameters) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *CertificateSearchParameters) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateSearchParameters) GetSerialNumber() *os.File`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateSearchParameters) GetSerialNumberOk() (**os.File, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateSearchParameters) SetSerialNumber(v *os.File)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateSearchParameters) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificateSearchParameters) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateSearchParameters) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateSearchParameters) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateSearchParameters) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetAki

`func (o *CertificateSearchParameters) GetAki() string`

GetAki returns the Aki field if non-nil, zero value otherwise.

### GetAkiOk

`func (o *CertificateSearchParameters) GetAkiOk() (*string, bool)`

GetAkiOk returns a tuple with the Aki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAki

`func (o *CertificateSearchParameters) SetAki(v string)`

SetAki sets Aki field to given value.

### HasAki

`func (o *CertificateSearchParameters) HasAki() bool`

HasAki returns a boolean if a field has been set.

### GetIssuingCaId

`func (o *CertificateSearchParameters) GetIssuingCaId() string`

GetIssuingCaId returns the IssuingCaId field if non-nil, zero value otherwise.

### GetIssuingCaIdOk

`func (o *CertificateSearchParameters) GetIssuingCaIdOk() (*string, bool)`

GetIssuingCaIdOk returns a tuple with the IssuingCaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCaId

`func (o *CertificateSearchParameters) SetIssuingCaId(v string)`

SetIssuingCaId sets IssuingCaId field to given value.

### HasIssuingCaId

`func (o *CertificateSearchParameters) HasIssuingCaId() bool`

HasIssuingCaId returns a boolean if a field has been set.

### GetIssuingCaName

`func (o *CertificateSearchParameters) GetIssuingCaName() string`

GetIssuingCaName returns the IssuingCaName field if non-nil, zero value otherwise.

### GetIssuingCaNameOk

`func (o *CertificateSearchParameters) GetIssuingCaNameOk() (*string, bool)`

GetIssuingCaNameOk returns a tuple with the IssuingCaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCaName

`func (o *CertificateSearchParameters) SetIssuingCaName(v string)`

SetIssuingCaName sets IssuingCaName field to given value.

### HasIssuingCaName

`func (o *CertificateSearchParameters) HasIssuingCaName() bool`

HasIssuingCaName returns a boolean if a field has been set.

### GetIssuingCaSerialNumber

`func (o *CertificateSearchParameters) GetIssuingCaSerialNumber() *os.File`

GetIssuingCaSerialNumber returns the IssuingCaSerialNumber field if non-nil, zero value otherwise.

### GetIssuingCaSerialNumberOk

`func (o *CertificateSearchParameters) GetIssuingCaSerialNumberOk() (**os.File, bool)`

GetIssuingCaSerialNumberOk returns a tuple with the IssuingCaSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCaSerialNumber

`func (o *CertificateSearchParameters) SetIssuingCaSerialNumber(v *os.File)`

SetIssuingCaSerialNumber sets IssuingCaSerialNumber field to given value.

### HasIssuingCaSerialNumber

`func (o *CertificateSearchParameters) HasIssuingCaSerialNumber() bool`

HasIssuingCaSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateSearchParameters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateSearchParameters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateSearchParameters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateSearchParameters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprint

`func (o *CertificateSearchParameters) GetThumbprint() *os.File`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificateSearchParameters) GetThumbprintOk() (**os.File, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificateSearchParameters) SetThumbprint(v *os.File)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificateSearchParameters) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


