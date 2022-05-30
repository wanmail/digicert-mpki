# ExternalEnrollmentReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 
**SeatType** | Pointer to **string** | Type of Seat to which this Enrollment is linked | [optional] 
**ProfileName** | Pointer to **string** | Name of Profile against which Enrollment was issued | [optional] 
**BuName** | Pointer to **string** | Name of Business unit to which Enrollment is linked | [optional] 
**Subject** | Pointer to **string** | Subject of Enrollment | [optional] 
**EnrollmentStatus** | Pointer to **string** | Current status of Enrollment | [optional] 
**CreateDate** | Pointer to **string** | Date when Enrollment was created in format YYYY-MM-DD | [optional] 
**ExpiryDate** | Pointer to **string** | Date when Enrollment should be (was) expired in format YYYY-MM-DD | [optional] 
**RedeemDate** | Pointer to **string** | Date when Enrollment was redeemed in format YYYY-MM-DD | [optional] 
**RejectDate** | Pointer to **string** | Date when Enrollment was rejected in format YYYY-MM-DD | [optional] 
**LastEditDate** | Pointer to **string** | Date when Enrollment was last edited in format YYYY-MM-DD | [optional] 
**LastEditAdmin** | Pointer to [**NameDto**](NameDto.md) |  | [optional] 

## Methods

### NewExternalEnrollmentReport

`func NewExternalEnrollmentReport() *ExternalEnrollmentReport`

NewExternalEnrollmentReport instantiates a new ExternalEnrollmentReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEnrollmentReportWithDefaults

`func NewExternalEnrollmentReportWithDefaults() *ExternalEnrollmentReport`

NewExternalEnrollmentReportWithDefaults instantiates a new ExternalEnrollmentReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *ExternalEnrollmentReport) GetSeat() NameDto`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *ExternalEnrollmentReport) GetSeatOk() (*NameDto, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *ExternalEnrollmentReport) SetSeat(v NameDto)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *ExternalEnrollmentReport) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetSeatType

`func (o *ExternalEnrollmentReport) GetSeatType() string`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *ExternalEnrollmentReport) GetSeatTypeOk() (*string, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *ExternalEnrollmentReport) SetSeatType(v string)`

SetSeatType sets SeatType field to given value.

### HasSeatType

`func (o *ExternalEnrollmentReport) HasSeatType() bool`

HasSeatType returns a boolean if a field has been set.

### GetProfileName

`func (o *ExternalEnrollmentReport) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ExternalEnrollmentReport) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ExternalEnrollmentReport) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ExternalEnrollmentReport) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetBuName

`func (o *ExternalEnrollmentReport) GetBuName() string`

GetBuName returns the BuName field if non-nil, zero value otherwise.

### GetBuNameOk

`func (o *ExternalEnrollmentReport) GetBuNameOk() (*string, bool)`

GetBuNameOk returns a tuple with the BuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuName

`func (o *ExternalEnrollmentReport) SetBuName(v string)`

SetBuName sets BuName field to given value.

### HasBuName

`func (o *ExternalEnrollmentReport) HasBuName() bool`

HasBuName returns a boolean if a field has been set.

### GetSubject

`func (o *ExternalEnrollmentReport) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExternalEnrollmentReport) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExternalEnrollmentReport) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ExternalEnrollmentReport) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetEnrollmentStatus

`func (o *ExternalEnrollmentReport) GetEnrollmentStatus() string`

GetEnrollmentStatus returns the EnrollmentStatus field if non-nil, zero value otherwise.

### GetEnrollmentStatusOk

`func (o *ExternalEnrollmentReport) GetEnrollmentStatusOk() (*string, bool)`

GetEnrollmentStatusOk returns a tuple with the EnrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentStatus

`func (o *ExternalEnrollmentReport) SetEnrollmentStatus(v string)`

SetEnrollmentStatus sets EnrollmentStatus field to given value.

### HasEnrollmentStatus

`func (o *ExternalEnrollmentReport) HasEnrollmentStatus() bool`

HasEnrollmentStatus returns a boolean if a field has been set.

### GetCreateDate

`func (o *ExternalEnrollmentReport) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ExternalEnrollmentReport) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ExternalEnrollmentReport) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ExternalEnrollmentReport) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *ExternalEnrollmentReport) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ExternalEnrollmentReport) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ExternalEnrollmentReport) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ExternalEnrollmentReport) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetRedeemDate

`func (o *ExternalEnrollmentReport) GetRedeemDate() string`

GetRedeemDate returns the RedeemDate field if non-nil, zero value otherwise.

### GetRedeemDateOk

`func (o *ExternalEnrollmentReport) GetRedeemDateOk() (*string, bool)`

GetRedeemDateOk returns a tuple with the RedeemDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemDate

`func (o *ExternalEnrollmentReport) SetRedeemDate(v string)`

SetRedeemDate sets RedeemDate field to given value.

### HasRedeemDate

`func (o *ExternalEnrollmentReport) HasRedeemDate() bool`

HasRedeemDate returns a boolean if a field has been set.

### GetRejectDate

`func (o *ExternalEnrollmentReport) GetRejectDate() string`

GetRejectDate returns the RejectDate field if non-nil, zero value otherwise.

### GetRejectDateOk

`func (o *ExternalEnrollmentReport) GetRejectDateOk() (*string, bool)`

GetRejectDateOk returns a tuple with the RejectDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectDate

`func (o *ExternalEnrollmentReport) SetRejectDate(v string)`

SetRejectDate sets RejectDate field to given value.

### HasRejectDate

`func (o *ExternalEnrollmentReport) HasRejectDate() bool`

HasRejectDate returns a boolean if a field has been set.

### GetLastEditDate

`func (o *ExternalEnrollmentReport) GetLastEditDate() string`

GetLastEditDate returns the LastEditDate field if non-nil, zero value otherwise.

### GetLastEditDateOk

`func (o *ExternalEnrollmentReport) GetLastEditDateOk() (*string, bool)`

GetLastEditDateOk returns a tuple with the LastEditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditDate

`func (o *ExternalEnrollmentReport) SetLastEditDate(v string)`

SetLastEditDate sets LastEditDate field to given value.

### HasLastEditDate

`func (o *ExternalEnrollmentReport) HasLastEditDate() bool`

HasLastEditDate returns a boolean if a field has been set.

### GetLastEditAdmin

`func (o *ExternalEnrollmentReport) GetLastEditAdmin() NameDto`

GetLastEditAdmin returns the LastEditAdmin field if non-nil, zero value otherwise.

### GetLastEditAdminOk

`func (o *ExternalEnrollmentReport) GetLastEditAdminOk() (*NameDto, bool)`

GetLastEditAdminOk returns a tuple with the LastEditAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditAdmin

`func (o *ExternalEnrollmentReport) SetLastEditAdmin(v NameDto)`

SetLastEditAdmin sets LastEditAdmin field to given value.

### HasLastEditAdmin

`func (o *ExternalEnrollmentReport) HasLastEditAdmin() bool`

HasLastEditAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


