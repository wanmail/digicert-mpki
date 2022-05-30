# MpkiApiV1CertificatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**IdDto**](IdDto.md) |  | [optional] 
**Seat** | Pointer to [**MpkiApiV1CertificatePostRequestSeat**](MpkiApiV1CertificatePostRequestSeat.md) |  | [optional] 
**Csr** | Pointer to **string** |  | [optional] 
**Validity** | Pointer to [**Validity**](Validity.md) |  | [optional] 
**DeliveryFormat** | Pointer to [**DeliveryFormat**](DeliveryFormat.md) |  | [optional] 
**IncludeCaChain** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to [**MpkiApiV1CertificatePostRequestAttributes**](MpkiApiV1CertificatePostRequestAttributes.md) |  | [optional] 

## Methods

### NewMpkiApiV1CertificatePostRequest

`func NewMpkiApiV1CertificatePostRequest() *MpkiApiV1CertificatePostRequest`

NewMpkiApiV1CertificatePostRequest instantiates a new MpkiApiV1CertificatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1CertificatePostRequestWithDefaults

`func NewMpkiApiV1CertificatePostRequestWithDefaults() *MpkiApiV1CertificatePostRequest`

NewMpkiApiV1CertificatePostRequestWithDefaults instantiates a new MpkiApiV1CertificatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *MpkiApiV1CertificatePostRequest) GetProfile() IdDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MpkiApiV1CertificatePostRequest) GetProfileOk() (*IdDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MpkiApiV1CertificatePostRequest) SetProfile(v IdDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MpkiApiV1CertificatePostRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeat

`func (o *MpkiApiV1CertificatePostRequest) GetSeat() MpkiApiV1CertificatePostRequestSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *MpkiApiV1CertificatePostRequest) GetSeatOk() (*MpkiApiV1CertificatePostRequestSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *MpkiApiV1CertificatePostRequest) SetSeat(v MpkiApiV1CertificatePostRequestSeat)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *MpkiApiV1CertificatePostRequest) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetCsr

`func (o *MpkiApiV1CertificatePostRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *MpkiApiV1CertificatePostRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *MpkiApiV1CertificatePostRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *MpkiApiV1CertificatePostRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetValidity

`func (o *MpkiApiV1CertificatePostRequest) GetValidity() Validity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *MpkiApiV1CertificatePostRequest) GetValidityOk() (*Validity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *MpkiApiV1CertificatePostRequest) SetValidity(v Validity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *MpkiApiV1CertificatePostRequest) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetDeliveryFormat

`func (o *MpkiApiV1CertificatePostRequest) GetDeliveryFormat() DeliveryFormat`

GetDeliveryFormat returns the DeliveryFormat field if non-nil, zero value otherwise.

### GetDeliveryFormatOk

`func (o *MpkiApiV1CertificatePostRequest) GetDeliveryFormatOk() (*DeliveryFormat, bool)`

GetDeliveryFormatOk returns a tuple with the DeliveryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFormat

`func (o *MpkiApiV1CertificatePostRequest) SetDeliveryFormat(v DeliveryFormat)`

SetDeliveryFormat sets DeliveryFormat field to given value.

### HasDeliveryFormat

`func (o *MpkiApiV1CertificatePostRequest) HasDeliveryFormat() bool`

HasDeliveryFormat returns a boolean if a field has been set.

### GetIncludeCaChain

`func (o *MpkiApiV1CertificatePostRequest) GetIncludeCaChain() bool`

GetIncludeCaChain returns the IncludeCaChain field if non-nil, zero value otherwise.

### GetIncludeCaChainOk

`func (o *MpkiApiV1CertificatePostRequest) GetIncludeCaChainOk() (*bool, bool)`

GetIncludeCaChainOk returns a tuple with the IncludeCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCaChain

`func (o *MpkiApiV1CertificatePostRequest) SetIncludeCaChain(v bool)`

SetIncludeCaChain sets IncludeCaChain field to given value.

### HasIncludeCaChain

`func (o *MpkiApiV1CertificatePostRequest) HasIncludeCaChain() bool`

HasIncludeCaChain returns a boolean if a field has been set.

### GetAttributes

`func (o *MpkiApiV1CertificatePostRequest) GetAttributes() MpkiApiV1CertificatePostRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MpkiApiV1CertificatePostRequest) GetAttributesOk() (*MpkiApiV1CertificatePostRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MpkiApiV1CertificatePostRequest) SetAttributes(v MpkiApiV1CertificatePostRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MpkiApiV1CertificatePostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


