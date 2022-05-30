# EnrollmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeatId** | Pointer to **string** | Seat id | [optional] 
**ProfileId** | Pointer to **string** | Profile id | [optional] 
**ProfileName** | Pointer to **string** | Profile name | [optional] 
**EnrollmentCode** | Pointer to **string** | Enrollment code to be supplied to the certificate end user | [optional] 
**CreationDateTime** | Pointer to **string** | Creation date and time | [optional] 
**ExpiryDateTime** | Pointer to **string** | Expiration date and time | [optional] 
**Status** | Pointer to **string** | Enrollment status | [optional] 
**NumberOfBadAttempts** | Pointer to **float32** | Number of erroneous attempts to complete enrollment (such as with PUT requests sent to /mpki/api/v1/enrollment/{enrollmentCode}) | [optional] 

## Methods

### NewEnrollmentResponse

`func NewEnrollmentResponse() *EnrollmentResponse`

NewEnrollmentResponse instantiates a new EnrollmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentResponseWithDefaults

`func NewEnrollmentResponseWithDefaults() *EnrollmentResponse`

NewEnrollmentResponseWithDefaults instantiates a new EnrollmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeatId

`func (o *EnrollmentResponse) GetSeatId() string`

GetSeatId returns the SeatId field if non-nil, zero value otherwise.

### GetSeatIdOk

`func (o *EnrollmentResponse) GetSeatIdOk() (*string, bool)`

GetSeatIdOk returns a tuple with the SeatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatId

`func (o *EnrollmentResponse) SetSeatId(v string)`

SetSeatId sets SeatId field to given value.

### HasSeatId

`func (o *EnrollmentResponse) HasSeatId() bool`

HasSeatId returns a boolean if a field has been set.

### GetProfileId

`func (o *EnrollmentResponse) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EnrollmentResponse) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *EnrollmentResponse) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *EnrollmentResponse) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *EnrollmentResponse) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *EnrollmentResponse) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *EnrollmentResponse) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *EnrollmentResponse) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetEnrollmentCode

`func (o *EnrollmentResponse) GetEnrollmentCode() string`

GetEnrollmentCode returns the EnrollmentCode field if non-nil, zero value otherwise.

### GetEnrollmentCodeOk

`func (o *EnrollmentResponse) GetEnrollmentCodeOk() (*string, bool)`

GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCode

`func (o *EnrollmentResponse) SetEnrollmentCode(v string)`

SetEnrollmentCode sets EnrollmentCode field to given value.

### HasEnrollmentCode

`func (o *EnrollmentResponse) HasEnrollmentCode() bool`

HasEnrollmentCode returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *EnrollmentResponse) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *EnrollmentResponse) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *EnrollmentResponse) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *EnrollmentResponse) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *EnrollmentResponse) GetExpiryDateTime() string`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *EnrollmentResponse) GetExpiryDateTimeOk() (*string, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *EnrollmentResponse) SetExpiryDateTime(v string)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *EnrollmentResponse) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *EnrollmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrollmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrollmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrollmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNumberOfBadAttempts

`func (o *EnrollmentResponse) GetNumberOfBadAttempts() float32`

GetNumberOfBadAttempts returns the NumberOfBadAttempts field if non-nil, zero value otherwise.

### GetNumberOfBadAttemptsOk

`func (o *EnrollmentResponse) GetNumberOfBadAttemptsOk() (*float32, bool)`

GetNumberOfBadAttemptsOk returns a tuple with the NumberOfBadAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBadAttempts

`func (o *EnrollmentResponse) SetNumberOfBadAttempts(v float32)`

SetNumberOfBadAttempts sets NumberOfBadAttempts field to given value.

### HasNumberOfBadAttempts

`func (o *EnrollmentResponse) HasNumberOfBadAttempts() bool`

HasNumberOfBadAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


