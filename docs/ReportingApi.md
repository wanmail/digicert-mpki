# {{classname}}

All URIs are relative to *https://one.digicert.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1AuditLogGet**](ReportingApi.md#MpkiApiV1AuditLogGet) | **Get** /mpki/api/v1/audit-log | List audit logs
[**MpkiApiV1AuditLogIdGet**](ReportingApi.md#MpkiApiV1AuditLogIdGet) | **Get** /mpki/api/v1/audit-log/{id} | Get audit log event by ID
[**MpkiApiV1ReportEnrollmentCodeGet**](ReportingApi.md#MpkiApiV1ReportEnrollmentCodeGet) | **Get** /mpki/api/v1/report/enrollment-code | Get enrollment code report

# **MpkiApiV1AuditLogGet**
> InlineResponse2005 MpkiApiV1AuditLogGet(ctx, optional)
List audit logs

Use this endpoint to get all available audit log entries in your account.   Results can optionally be filtered using search and paging parameters appended to the endpoint URL (as query tags), e.g. `?create_date=22-02-2022`.   Filters for audit log data fields can also be applied to results using this format: `?data[resource_type]=enrollment`. For example, `?data[thumbprint]=12143A4E7XXX2C589EE8AE3C86321B85CA7271328XXX1C9C836935D45XXXBE9X`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiMpkiApiV1AuditLogGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiMpkiApiV1AuditLogGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagingParameters** | [**optional.Interface of PagingParameters**](.md)|  | 
 **searchParameters** | [**optional.Interface of LogSearchParameters**](.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1AuditLogIdGet**
> AuditLogItem MpkiApiV1AuditLogIdGet(ctx, id)
Get audit log event by ID

Use this endpoint to get the details of a given audit log event ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Audit log event ID (displayed in the audit event details on UI, and in audit logs list response) | 

### Return type

[**AuditLogItem**](AuditLogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1ReportEnrollmentCodeGet**
> InlineResponse2004 MpkiApiV1ReportEnrollmentCodeGet(ctx, optional)
Get enrollment code report

Use this endpoint to get a report of all enrollment code events.   Results can optionally be filtered using search and paging parameters appended to the endpoint URL (as query tags), e.g. `?create_date=22-02-2022`.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiMpkiApiV1ReportEnrollmentCodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiMpkiApiV1ReportEnrollmentCodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagingParameters** | [**optional.Interface of PagingParameters**](.md)|  | 
 **searchParameters** | [**optional.Interface of ReportSearchParams**](.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

