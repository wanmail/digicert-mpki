# {{classname}}

All URIs are relative to *https://one.digicert.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1SeatPost**](SeatsApi.md#MpkiApiV1SeatPost) | **Post** /mpki/api/v1/seat | Create seat
[**MpkiApiV1SeatSeatIdDelete**](SeatsApi.md#MpkiApiV1SeatSeatIdDelete) | **Delete** /mpki/api/v1/seat/{seat_id} | Delete seat
[**MpkiApiV1SeatSeatIdGet**](SeatsApi.md#MpkiApiV1SeatSeatIdGet) | **Get** /mpki/api/v1/seat/{seat_id} | Get seat by ID
[**MpkiApiV1SeatSeatIdPut**](SeatsApi.md#MpkiApiV1SeatSeatIdPut) | **Put** /mpki/api/v1/seat/{seat_id} | Update seat
[**MpkiApiV1SeatTypesGet**](SeatsApi.md#MpkiApiV1SeatTypesGet) | **Get** /mpki/api/v1/seat-types | Get available seat types

# **MpkiApiV1SeatPost**
> SeatDetails MpkiApiV1SeatPost(ctx, optional)
Create seat

Use this endpoint to create a new seat. Each seat decreases the number of allocated seats for the given seat type. This endpoint does not restore previously deleted seats—reusing deleted seat IDs results in the creation of a new seat with the previously deleted seat ID.   Each unique seat ID represents only one active seat at any given time — duplicate seat IDs are not permitted.   `seat_id` is the identifier for each seat — this must be unique within each business unit and for each seat type.  The `seat_type` parameter is optional. If no `seat_type` is specified, this defaults to `USER_SEAT`. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SeatsApiMpkiApiV1SeatPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeatsApiMpkiApiV1SeatPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1SeatBody**](V1SeatBody.md)|  | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1SeatSeatIdDelete**
> SeatDetails MpkiApiV1SeatSeatIdDelete(ctx, seatId, seatTypeId, optional)
Delete seat

Use this endpoint to delete a given seat ID. Deleting the seat ID restores an available seat of the same type in your account. As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit. The `seat_type` parameter is required. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **seatId** | **string**| Seat ID | 
  **seatTypeId** | **string**| Seat Type ID | 
 **optional** | ***SeatsApiMpkiApiV1SeatSeatIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeatsApiMpkiApiV1SeatSeatIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **businessUnitId** | [**optional.Interface of string**](.md)| Business unit ID | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1SeatSeatIdGet**
> InlineResponse200 MpkiApiV1SeatSeatIdGet(ctx, seatId, seatTypeId, optional)
Get seat by ID

Use this endpoint to retrieve the details of a given seat ID. Seat IDs are unique for each `business_unit_id` and `seat_type`. As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit. The `seat_type` parameter is required. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **seatId** | **string**| Seat ID | 
  **seatTypeId** | **string**| Seat Type ID | 
 **optional** | ***SeatsApiMpkiApiV1SeatSeatIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeatsApiMpkiApiV1SeatSeatIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **businessUnitId** | [**optional.Interface of string**](.md)| Business unit ID.  As identical seat IDs may exist across various business units of the same account, the &#x60;business_unit_id&#x60; for the given seat must be specified as a query parameter if the account contains more than one business unit.  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1SeatSeatIdPut**
> SeatDetails MpkiApiV1SeatSeatIdPut(ctx, seatId, optional)
Update seat

Use this endpoint to update the details for a given seat ID. Include the seat details to be updated in the request body. As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit. The `seat_type` parameter is optional. If no `seat_type` is specified, this defaults to `USER_SEAT`. Possible values for `seat_type.id` are `USER_SEAT`, `DEVICE_SEAT` and `SERVER_SEAT`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **seatId** | **string**| Seat ID | 
 **optional** | ***SeatsApiMpkiApiV1SeatSeatIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeatsApiMpkiApiV1SeatSeatIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SeatSeatIdBody**](SeatSeatIdBody.md)|  | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1SeatTypesGet**
> []InlineResponse2001 MpkiApiV1SeatTypesGet(ctx, )
Get available seat types

Use this endpoint to fetch details about available seat types for the account. Returns an array containing seat type details.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

