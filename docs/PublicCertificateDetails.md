# PublicCertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 
**Seat** | Pointer to [**MpkiApiV1CertificatePostRequestSeat**](MpkiApiV1CertificatePostRequestSeat.md) |  | [optional] 
**Account** | Pointer to [**IdDto**](IdDto.md) |  | [optional] 
**BusinessUnit** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**ValidFrom** | Pointer to **string** |  | [optional] 
**ValidTo** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Revocation** | Pointer to [**PublicCertificateDetailsRevocation**](PublicCertificateDetailsRevocation.md) |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicCertificateDetails

`func NewPublicCertificateDetails() *PublicCertificateDetails`

NewPublicCertificateDetails instantiates a new PublicCertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCertificateDetailsWithDefaults

`func NewPublicCertificateDetailsWithDefaults() *PublicCertificateDetails`

NewPublicCertificateDetailsWithDefaults instantiates a new PublicCertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *PublicCertificateDetails) GetProfile() NameDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PublicCertificateDetails) GetProfileOk() (*NameDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PublicCertificateDetails) SetProfile(v NameDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PublicCertificateDetails) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeat

`func (o *PublicCertificateDetails) GetSeat() MpkiApiV1CertificatePostRequestSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *PublicCertificateDetails) GetSeatOk() (*MpkiApiV1CertificatePostRequestSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *PublicCertificateDetails) SetSeat(v MpkiApiV1CertificatePostRequestSeat)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *PublicCertificateDetails) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetAccount

`func (o *PublicCertificateDetails) GetAccount() IdDto`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PublicCertificateDetails) GetAccountOk() (*IdDto, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PublicCertificateDetails) SetAccount(v IdDto)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PublicCertificateDetails) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *PublicCertificateDetails) GetBusinessUnit() NameDto`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *PublicCertificateDetails) GetBusinessUnitOk() (*NameDto, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *PublicCertificateDetails) SetBusinessUnit(v NameDto)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *PublicCertificateDetails) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetCertificate

`func (o *PublicCertificateDetails) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PublicCertificateDetails) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PublicCertificateDetails) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PublicCertificateDetails) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCommonName

`func (o *PublicCertificateDetails) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PublicCertificateDetails) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PublicCertificateDetails) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PublicCertificateDetails) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetStatus

`func (o *PublicCertificateDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicCertificateDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicCertificateDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicCertificateDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PublicCertificateDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PublicCertificateDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PublicCertificateDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PublicCertificateDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetValidFrom

`func (o *PublicCertificateDetails) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *PublicCertificateDetails) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *PublicCertificateDetails) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *PublicCertificateDetails) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *PublicCertificateDetails) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *PublicCertificateDetails) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *PublicCertificateDetails) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *PublicCertificateDetails) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetNotes

`func (o *PublicCertificateDetails) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PublicCertificateDetails) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PublicCertificateDetails) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PublicCertificateDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRevocation

`func (o *PublicCertificateDetails) GetRevocation() PublicCertificateDetailsRevocation`

GetRevocation returns the Revocation field if non-nil, zero value otherwise.

### GetRevocationOk

`func (o *PublicCertificateDetails) GetRevocationOk() (*PublicCertificateDetailsRevocation, bool)`

GetRevocationOk returns a tuple with the Revocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocation

`func (o *PublicCertificateDetails) SetRevocation(v PublicCertificateDetailsRevocation)`

SetRevocation sets Revocation field to given value.

### HasRevocation

`func (o *PublicCertificateDetails) HasRevocation() bool`

HasRevocation returns a boolean if a field has been set.

### GetThumbprint

`func (o *PublicCertificateDetails) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *PublicCertificateDetails) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *PublicCertificateDetails) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *PublicCertificateDetails) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


