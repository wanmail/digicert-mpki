# {{classname}}

All URIs are relative to *https://one.digicert.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1EnrollmentDetailsGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentDetailsGet) | **Get** /mpki/api/v1/enrollment-details/ | List enrollments
[**MpkiApiV1EnrollmentDetailsIdentifierGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentDetailsIdentifierGet) | **Get** /mpki/api/v1/enrollment-details/{identifier} | Get enrollment details by ID
[**MpkiApiV1EnrollmentEnrollmentCodeGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentEnrollmentCodeGet) | **Get** /mpki/api/v1/enrollment/{enrollmentCode} | Get enrollment details by enrollment code
[**MpkiApiV1EnrollmentPost**](EnrollmentsApi.md#MpkiApiV1EnrollmentPost) | **Post** /mpki/api/v1/enrollment | Create enrollment
[**MpkiApiV1EnrollmentRedeemPost**](EnrollmentsApi.md#MpkiApiV1EnrollmentRedeemPost) | **Post** /mpki/api/v1/enrollment/redeem | Redeem enrollment

# **MpkiApiV1EnrollmentDetailsGet**
> []PublicEnrollmentDetails MpkiApiV1EnrollmentDetailsGet(ctx, optional)
List enrollments

Use this endpoint to list all pending certificate enrollments in your account.   Results can be filtered using optional search and paging parameters, added as query tags. Append `?{query_parameter}=value` to the endpoint URL, separate multiple parameters with `&`. For example, `/mpki/api/v1/enrollment-details?account_id={account ID}&business_unit_id={business unit ID}` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnrollmentsApiMpkiApiV1EnrollmentDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentsApiMpkiApiV1EnrollmentDetailsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchParameters** | [**optional.Interface of EnrollmentListParameters**](.md)|  | 

### Return type

[**[]PublicEnrollmentDetails**](PublicEnrollmentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1EnrollmentDetailsIdentifierGet**
> PublicEnrollmentDetails MpkiApiV1EnrollmentDetailsIdentifierGet(ctx, identifier)
Get enrollment details by ID

Use this endpoint to get enrollment details for a given enrollment ID.   The unique enrollment ID is returned when the enrollment is created. The enrollment ID can also be retrieved from the enrollment details page in your account, or a list enrollments GET request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Enrollment ID | 

### Return type

[**PublicEnrollmentDetails**](PublicEnrollmentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1EnrollmentEnrollmentCodeGet**
> EnrollmentResponse MpkiApiV1EnrollmentEnrollmentCodeGet(ctx, enrollmentCode, seatId)
Get enrollment details by enrollment code

Use this endpoint to get enrollment details for a given enrollment code and seat ID.   `seat_id` is a required query parameter to be included in requests to this endpoint. Append `?seat_id={seat ID}` to the endpoint URL, replace `{seat ID}` with the seat ID to be retrieved.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enrollmentCode** | **string**| Enrollment code received in enrollment creation response | 
  **seatId** | **string**| Seat ID | 

### Return type

[**EnrollmentResponse**](EnrollmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1EnrollmentPost**
> EnrollmentResponse MpkiApiV1EnrollmentPost(ctx, optional)
Create enrollment

Use this endpoint to create a pending certificate enrollment, to be completed and submitted by the certificate end user.   The following enrollment details are required fields in the request body:  `profile` — Certificate profile ID   `seat` — Seat ID and seat email address   `attributes` — Enrollment attributes, including certificate subject DN information, validity period, and CSR (if applicable)  A custom `enrollment_code` can optionally be included in the request body. If no enrollment code is provided in the request, one will be generated and returned in the response body.   `expiry_date` is an optional request body field. If no expiry_date is specified in the request, and the certificate profile has no set enrollment code validity period, the default value is 10 days.   **Note:** This endpoint only accepts enrollment requests for certificate profiles with the `ENROLLMENT_CODE` authentication method.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnrollmentsApiMpkiApiV1EnrollmentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentsApiMpkiApiV1EnrollmentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1EnrollmentBody**](V1EnrollmentBody.md)|  | 

### Return type

[**EnrollmentResponse**](EnrollmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1EnrollmentRedeemPost**
> CertificateResponse MpkiApiV1EnrollmentRedeemPost(ctx, optional)
Redeem enrollment

Use this endpoint to redeem a pending certificate enrollment.   The issued certificate is returned in the response body.  The following enrollment details are required fields in the request body:  `profile` — Certificate profile ID   `seat` — Seat ID  `enrollment_code` — The enrollment code which was provided by the administrator, or generated automatically, when the enrollment was created. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnrollmentsApiMpkiApiV1EnrollmentRedeemPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnrollmentsApiMpkiApiV1EnrollmentRedeemPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EnrollmentRedeemBody**](EnrollmentRedeemBody.md)|  | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

