# MpkiApiV1EnrollmentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **string** | Entity ID in UUID format | 
**Seat** | [**MpkiApiV1EnrollmentPostRequestSeat**](MpkiApiV1EnrollmentPostRequestSeat.md) |  | 
**Attributes** | **map[string]interface{}** |  | 
**EnrollmentCode** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** | This parameter can only accept \&quot;yyyy-MM-dd\&quot; and \&quot;yyyy-MM-dd HH:mm:ss\&quot; formats | [optional] 

## Methods

### NewMpkiApiV1EnrollmentPostRequest

`func NewMpkiApiV1EnrollmentPostRequest(profile string, seat MpkiApiV1EnrollmentPostRequestSeat, attributes map[string]interface{}, ) *MpkiApiV1EnrollmentPostRequest`

NewMpkiApiV1EnrollmentPostRequest instantiates a new MpkiApiV1EnrollmentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1EnrollmentPostRequestWithDefaults

`func NewMpkiApiV1EnrollmentPostRequestWithDefaults() *MpkiApiV1EnrollmentPostRequest`

NewMpkiApiV1EnrollmentPostRequestWithDefaults instantiates a new MpkiApiV1EnrollmentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *MpkiApiV1EnrollmentPostRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MpkiApiV1EnrollmentPostRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MpkiApiV1EnrollmentPostRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetSeat

`func (o *MpkiApiV1EnrollmentPostRequest) GetSeat() MpkiApiV1EnrollmentPostRequestSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *MpkiApiV1EnrollmentPostRequest) GetSeatOk() (*MpkiApiV1EnrollmentPostRequestSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *MpkiApiV1EnrollmentPostRequest) SetSeat(v MpkiApiV1EnrollmentPostRequestSeat)`

SetSeat sets Seat field to given value.


### GetAttributes

`func (o *MpkiApiV1EnrollmentPostRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MpkiApiV1EnrollmentPostRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MpkiApiV1EnrollmentPostRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetEnrollmentCode

`func (o *MpkiApiV1EnrollmentPostRequest) GetEnrollmentCode() string`

GetEnrollmentCode returns the EnrollmentCode field if non-nil, zero value otherwise.

### GetEnrollmentCodeOk

`func (o *MpkiApiV1EnrollmentPostRequest) GetEnrollmentCodeOk() (*string, bool)`

GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCode

`func (o *MpkiApiV1EnrollmentPostRequest) SetEnrollmentCode(v string)`

SetEnrollmentCode sets EnrollmentCode field to given value.

### HasEnrollmentCode

`func (o *MpkiApiV1EnrollmentPostRequest) HasEnrollmentCode() bool`

HasEnrollmentCode returns a boolean if a field has been set.

### GetExpiryDate

`func (o *MpkiApiV1EnrollmentPostRequest) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MpkiApiV1EnrollmentPostRequest) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MpkiApiV1EnrollmentPostRequest) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MpkiApiV1EnrollmentPostRequest) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


