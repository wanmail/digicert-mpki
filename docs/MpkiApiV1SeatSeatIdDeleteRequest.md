# MpkiApiV1SeatSeatIdDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeatName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**SeatType** | Pointer to [**MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType**](MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType.md) |  | [optional] 
**BusinessUnitId** | Pointer to **string** | Entity ID in UUID format | [optional] 

## Methods

### NewMpkiApiV1SeatSeatIdDeleteRequest

`func NewMpkiApiV1SeatSeatIdDeleteRequest() *MpkiApiV1SeatSeatIdDeleteRequest`

NewMpkiApiV1SeatSeatIdDeleteRequest instantiates a new MpkiApiV1SeatSeatIdDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1SeatSeatIdDeleteRequestWithDefaults

`func NewMpkiApiV1SeatSeatIdDeleteRequestWithDefaults() *MpkiApiV1SeatSeatIdDeleteRequest`

NewMpkiApiV1SeatSeatIdDeleteRequestWithDefaults instantiates a new MpkiApiV1SeatSeatIdDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeatName

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetSeatName() string`

GetSeatName returns the SeatName field if non-nil, zero value otherwise.

### GetSeatNameOk

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetSeatNameOk() (*string, bool)`

GetSeatNameOk returns a tuple with the SeatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatName

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) SetSeatName(v string)`

SetSeatName sets SeatName field to given value.

### HasSeatName

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) HasSeatName() bool`

HasSeatName returns a boolean if a field has been set.

### GetEmail

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSeatType

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetSeatType() MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetSeatTypeOk() (*MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) SetSeatType(v MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType)`

SetSeatType sets SeatType field to given value.

### HasSeatType

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) HasSeatType() bool`

HasSeatType returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *MpkiApiV1SeatSeatIdDeleteRequest) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


