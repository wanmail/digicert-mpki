# PublicEnrollmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Enrollment internal id | [optional] [default to null]
**Status** | **string** | Status of the enrollment | [optional] [default to null]
**Email** | **string** | Email used to enroll | [optional] [default to null]
**CreatedAt** | **string** | Creation date/time of the enrollment | [optional] [default to null]
**ExpiresAt** | **string** | Expiry date/time of the enrollment | [optional] [default to null]
**Profile** | [***NameDto**](NameDto.md) |  | [optional] [default to null]
**BusinessUnit** | [***NameDto**](NameDto.md) |  | [optional] [default to null]
**MessageToUser** | **string** | A message sent to user by admin | [optional] [default to null]
**SeatIdMappingValue** | **string** | Value used as seat id mapping | [optional] [default to null]
**ApprovedByCurrentUser** | **bool** | Whether this particular enrollment has been approved by the current user | [optional] [default to null]
**Seat** | [***PublicEnrollmentDetailsSeat**](PublicEnrollmentDetails_seat.md) |  | [optional] [default to null]
**Comments** | **string** | Comments to the enrollment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

