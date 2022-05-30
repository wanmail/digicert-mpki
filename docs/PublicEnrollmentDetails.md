# PublicEnrollmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Enrollment internal id | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**Email** | Pointer to **string** | Email used to enroll | [optional] 
**CreatedAt** | Pointer to **string** | Creation date/time of the enrollment | [optional] 
**ExpiresAt** | Pointer to **string** | Expiry date/time of the enrollment | [optional] 
**Profile** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 
**BusinessUnit** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 
**MessageToUser** | Pointer to **string** | A message sent to user by admin | [optional] 
**SeatIdMappingValue** | Pointer to **string** | Value used as seat id mapping | [optional] 
**ApprovedByCurrentUser** | Pointer to **bool** | Whether this particular enrollment has been approved by the current user | [optional] 
**Seat** | Pointer to [**PublicEnrollmentDetailsSeat**](PublicEnrollmentDetailsSeat.md) |  | [optional] 
**Comments** | Pointer to **string** | Comments to the enrollment | [optional] 

## Methods

### NewPublicEnrollmentDetails

`func NewPublicEnrollmentDetails() *PublicEnrollmentDetails`

NewPublicEnrollmentDetails instantiates a new PublicEnrollmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEnrollmentDetailsWithDefaults

`func NewPublicEnrollmentDetailsWithDefaults() *PublicEnrollmentDetails`

NewPublicEnrollmentDetailsWithDefaults instantiates a new PublicEnrollmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicEnrollmentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicEnrollmentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicEnrollmentDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicEnrollmentDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PublicEnrollmentDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicEnrollmentDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicEnrollmentDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicEnrollmentDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmail

`func (o *PublicEnrollmentDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicEnrollmentDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicEnrollmentDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PublicEnrollmentDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicEnrollmentDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicEnrollmentDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicEnrollmentDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicEnrollmentDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PublicEnrollmentDetails) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicEnrollmentDetails) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicEnrollmentDetails) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PublicEnrollmentDetails) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProfile

`func (o *PublicEnrollmentDetails) GetProfile() NameDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PublicEnrollmentDetails) GetProfileOk() (*NameDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PublicEnrollmentDetails) SetProfile(v NameDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PublicEnrollmentDetails) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *PublicEnrollmentDetails) GetBusinessUnit() NameDto`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *PublicEnrollmentDetails) GetBusinessUnitOk() (*NameDto, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *PublicEnrollmentDetails) SetBusinessUnit(v NameDto)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *PublicEnrollmentDetails) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetMessageToUser

`func (o *PublicEnrollmentDetails) GetMessageToUser() string`

GetMessageToUser returns the MessageToUser field if non-nil, zero value otherwise.

### GetMessageToUserOk

`func (o *PublicEnrollmentDetails) GetMessageToUserOk() (*string, bool)`

GetMessageToUserOk returns a tuple with the MessageToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageToUser

`func (o *PublicEnrollmentDetails) SetMessageToUser(v string)`

SetMessageToUser sets MessageToUser field to given value.

### HasMessageToUser

`func (o *PublicEnrollmentDetails) HasMessageToUser() bool`

HasMessageToUser returns a boolean if a field has been set.

### GetSeatIdMappingValue

`func (o *PublicEnrollmentDetails) GetSeatIdMappingValue() string`

GetSeatIdMappingValue returns the SeatIdMappingValue field if non-nil, zero value otherwise.

### GetSeatIdMappingValueOk

`func (o *PublicEnrollmentDetails) GetSeatIdMappingValueOk() (*string, bool)`

GetSeatIdMappingValueOk returns a tuple with the SeatIdMappingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIdMappingValue

`func (o *PublicEnrollmentDetails) SetSeatIdMappingValue(v string)`

SetSeatIdMappingValue sets SeatIdMappingValue field to given value.

### HasSeatIdMappingValue

`func (o *PublicEnrollmentDetails) HasSeatIdMappingValue() bool`

HasSeatIdMappingValue returns a boolean if a field has been set.

### GetApprovedByCurrentUser

`func (o *PublicEnrollmentDetails) GetApprovedByCurrentUser() bool`

GetApprovedByCurrentUser returns the ApprovedByCurrentUser field if non-nil, zero value otherwise.

### GetApprovedByCurrentUserOk

`func (o *PublicEnrollmentDetails) GetApprovedByCurrentUserOk() (*bool, bool)`

GetApprovedByCurrentUserOk returns a tuple with the ApprovedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByCurrentUser

`func (o *PublicEnrollmentDetails) SetApprovedByCurrentUser(v bool)`

SetApprovedByCurrentUser sets ApprovedByCurrentUser field to given value.

### HasApprovedByCurrentUser

`func (o *PublicEnrollmentDetails) HasApprovedByCurrentUser() bool`

HasApprovedByCurrentUser returns a boolean if a field has been set.

### GetSeat

`func (o *PublicEnrollmentDetails) GetSeat() PublicEnrollmentDetailsSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *PublicEnrollmentDetails) GetSeatOk() (*PublicEnrollmentDetailsSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *PublicEnrollmentDetails) SetSeat(v PublicEnrollmentDetailsSeat)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *PublicEnrollmentDetails) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetComments

`func (o *PublicEnrollmentDetails) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PublicEnrollmentDetails) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PublicEnrollmentDetails) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PublicEnrollmentDetails) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


