# EnrollmentListParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**SeatId** | Pointer to **string** | Internal UUID of seat used for enrollment | [optional] 
**SeatIdentifier** | Pointer to **string** | Unique identifier of seat used for enrollment | [optional] 
**SeatName** | Pointer to **string** | Name of seat used for enrollment | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to **[]string** | List of enrollment statuses to limit search results to | [optional] 
**CreatedAtFrom** | Pointer to **string** | Limit results to enrollments created after provided date | [optional] 
**CreatedAtTo** | Pointer to **string** | Limit results to enrollments created before provided date | [optional] 
**ExpiresAtFrom** | Pointer to **string** | Limit results to enrollments expiring after provided date | [optional] 
**ExpiresAtTo** | Pointer to **string** | Limit results to enrollments expiring before provided date | [optional] 

## Methods

### NewEnrollmentListParameters

`func NewEnrollmentListParameters() *EnrollmentListParameters`

NewEnrollmentListParameters instantiates a new EnrollmentListParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentListParametersWithDefaults

`func NewEnrollmentListParametersWithDefaults() *EnrollmentListParameters`

NewEnrollmentListParametersWithDefaults instantiates a new EnrollmentListParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnrollmentListParameters) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnrollmentListParameters) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnrollmentListParameters) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *EnrollmentListParameters) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *EnrollmentListParameters) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *EnrollmentListParameters) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *EnrollmentListParameters) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *EnrollmentListParameters) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetProfileId

`func (o *EnrollmentListParameters) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EnrollmentListParameters) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *EnrollmentListParameters) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *EnrollmentListParameters) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetSeatId

`func (o *EnrollmentListParameters) GetSeatId() string`

GetSeatId returns the SeatId field if non-nil, zero value otherwise.

### GetSeatIdOk

`func (o *EnrollmentListParameters) GetSeatIdOk() (*string, bool)`

GetSeatIdOk returns a tuple with the SeatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatId

`func (o *EnrollmentListParameters) SetSeatId(v string)`

SetSeatId sets SeatId field to given value.

### HasSeatId

`func (o *EnrollmentListParameters) HasSeatId() bool`

HasSeatId returns a boolean if a field has been set.

### GetSeatIdentifier

`func (o *EnrollmentListParameters) GetSeatIdentifier() string`

GetSeatIdentifier returns the SeatIdentifier field if non-nil, zero value otherwise.

### GetSeatIdentifierOk

`func (o *EnrollmentListParameters) GetSeatIdentifierOk() (*string, bool)`

GetSeatIdentifierOk returns a tuple with the SeatIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIdentifier

`func (o *EnrollmentListParameters) SetSeatIdentifier(v string)`

SetSeatIdentifier sets SeatIdentifier field to given value.

### HasSeatIdentifier

`func (o *EnrollmentListParameters) HasSeatIdentifier() bool`

HasSeatIdentifier returns a boolean if a field has been set.

### GetSeatName

`func (o *EnrollmentListParameters) GetSeatName() string`

GetSeatName returns the SeatName field if non-nil, zero value otherwise.

### GetSeatNameOk

`func (o *EnrollmentListParameters) GetSeatNameOk() (*string, bool)`

GetSeatNameOk returns a tuple with the SeatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatName

`func (o *EnrollmentListParameters) SetSeatName(v string)`

SetSeatName sets SeatName field to given value.

### HasSeatName

`func (o *EnrollmentListParameters) HasSeatName() bool`

HasSeatName returns a boolean if a field has been set.

### GetEmail

`func (o *EnrollmentListParameters) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EnrollmentListParameters) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EnrollmentListParameters) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EnrollmentListParameters) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatuses

`func (o *EnrollmentListParameters) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *EnrollmentListParameters) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *EnrollmentListParameters) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *EnrollmentListParameters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetCreatedAtFrom

`func (o *EnrollmentListParameters) GetCreatedAtFrom() string`

GetCreatedAtFrom returns the CreatedAtFrom field if non-nil, zero value otherwise.

### GetCreatedAtFromOk

`func (o *EnrollmentListParameters) GetCreatedAtFromOk() (*string, bool)`

GetCreatedAtFromOk returns a tuple with the CreatedAtFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtFrom

`func (o *EnrollmentListParameters) SetCreatedAtFrom(v string)`

SetCreatedAtFrom sets CreatedAtFrom field to given value.

### HasCreatedAtFrom

`func (o *EnrollmentListParameters) HasCreatedAtFrom() bool`

HasCreatedAtFrom returns a boolean if a field has been set.

### GetCreatedAtTo

`func (o *EnrollmentListParameters) GetCreatedAtTo() string`

GetCreatedAtTo returns the CreatedAtTo field if non-nil, zero value otherwise.

### GetCreatedAtToOk

`func (o *EnrollmentListParameters) GetCreatedAtToOk() (*string, bool)`

GetCreatedAtToOk returns a tuple with the CreatedAtTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTo

`func (o *EnrollmentListParameters) SetCreatedAtTo(v string)`

SetCreatedAtTo sets CreatedAtTo field to given value.

### HasCreatedAtTo

`func (o *EnrollmentListParameters) HasCreatedAtTo() bool`

HasCreatedAtTo returns a boolean if a field has been set.

### GetExpiresAtFrom

`func (o *EnrollmentListParameters) GetExpiresAtFrom() string`

GetExpiresAtFrom returns the ExpiresAtFrom field if non-nil, zero value otherwise.

### GetExpiresAtFromOk

`func (o *EnrollmentListParameters) GetExpiresAtFromOk() (*string, bool)`

GetExpiresAtFromOk returns a tuple with the ExpiresAtFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtFrom

`func (o *EnrollmentListParameters) SetExpiresAtFrom(v string)`

SetExpiresAtFrom sets ExpiresAtFrom field to given value.

### HasExpiresAtFrom

`func (o *EnrollmentListParameters) HasExpiresAtFrom() bool`

HasExpiresAtFrom returns a boolean if a field has been set.

### GetExpiresAtTo

`func (o *EnrollmentListParameters) GetExpiresAtTo() string`

GetExpiresAtTo returns the ExpiresAtTo field if non-nil, zero value otherwise.

### GetExpiresAtToOk

`func (o *EnrollmentListParameters) GetExpiresAtToOk() (*string, bool)`

GetExpiresAtToOk returns a tuple with the ExpiresAtTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtTo

`func (o *EnrollmentListParameters) SetExpiresAtTo(v string)`

SetExpiresAtTo sets ExpiresAtTo field to given value.

### HasExpiresAtTo

`func (o *EnrollmentListParameters) HasExpiresAtTo() bool`

HasExpiresAtTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


