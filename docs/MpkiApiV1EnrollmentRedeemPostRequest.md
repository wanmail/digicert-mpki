# MpkiApiV1EnrollmentRedeemPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | Pointer to [**MpkiApiV1EnrollmentRedeemPostRequestSeat**](MpkiApiV1EnrollmentRedeemPostRequestSeat.md) |  | [optional] 
**Profile** | Pointer to [**MpkiApiV1EnrollmentRedeemPostRequestSeat**](MpkiApiV1EnrollmentRedeemPostRequestSeat.md) |  | [optional] 
**EnrollmentCode** | Pointer to **string** |  | [optional] 
**Csr** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Validity** | Pointer to [**Validity**](Validity.md) |  | [optional] 

## Methods

### NewMpkiApiV1EnrollmentRedeemPostRequest

`func NewMpkiApiV1EnrollmentRedeemPostRequest() *MpkiApiV1EnrollmentRedeemPostRequest`

NewMpkiApiV1EnrollmentRedeemPostRequest instantiates a new MpkiApiV1EnrollmentRedeemPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1EnrollmentRedeemPostRequestWithDefaults

`func NewMpkiApiV1EnrollmentRedeemPostRequestWithDefaults() *MpkiApiV1EnrollmentRedeemPostRequest`

NewMpkiApiV1EnrollmentRedeemPostRequestWithDefaults instantiates a new MpkiApiV1EnrollmentRedeemPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetSeat() MpkiApiV1EnrollmentRedeemPostRequestSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetSeatOk() (*MpkiApiV1EnrollmentRedeemPostRequestSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetSeat(v MpkiApiV1EnrollmentRedeemPostRequestSeat)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetProfile

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetProfile() MpkiApiV1EnrollmentRedeemPostRequestSeat`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetProfileOk() (*MpkiApiV1EnrollmentRedeemPostRequestSeat, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetProfile(v MpkiApiV1EnrollmentRedeemPostRequestSeat)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetEnrollmentCode

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetEnrollmentCode() string`

GetEnrollmentCode returns the EnrollmentCode field if non-nil, zero value otherwise.

### GetEnrollmentCodeOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetEnrollmentCodeOk() (*string, bool)`

GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCode

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetEnrollmentCode(v string)`

SetEnrollmentCode sets EnrollmentCode field to given value.

### HasEnrollmentCode

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasEnrollmentCode() bool`

HasEnrollmentCode returns a boolean if a field has been set.

### GetCsr

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetAttributes

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetValidity

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetValidity() Validity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) GetValidityOk() (*Validity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) SetValidity(v Validity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *MpkiApiV1EnrollmentRedeemPostRequest) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


