# \ReportingApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1AuditLogGet**](ReportingApi.md#MpkiApiV1AuditLogGet) | **Get** /mpki/api/v1/audit-log | List audit logs
[**MpkiApiV1AuditLogIdGet**](ReportingApi.md#MpkiApiV1AuditLogIdGet) | **Get** /mpki/api/v1/audit-log/{id} | Get audit log event by ID
[**MpkiApiV1ReportEnrollmentCodeGet**](ReportingApi.md#MpkiApiV1ReportEnrollmentCodeGet) | **Get** /mpki/api/v1/report/enrollment-code | Get enrollment code report



## MpkiApiV1AuditLogGet

> MpkiApiV1AuditLogGet200Response MpkiApiV1AuditLogGet(ctx).PagingParameters(pagingParameters).SearchParameters(searchParameters).Execute()

List audit logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pagingParameters := map[string][]openapiclient.PagingParameters{ ... } // PagingParameters |  (optional)
    searchParameters := map[string][]openapiclient.LogSearchParameters{ ... } // LogSearchParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.MpkiApiV1AuditLogGet(context.Background()).PagingParameters(pagingParameters).SearchParameters(searchParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.MpkiApiV1AuditLogGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1AuditLogGet`: MpkiApiV1AuditLogGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.MpkiApiV1AuditLogGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1AuditLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagingParameters** | [**PagingParameters**](PagingParameters.md) |  | 
 **searchParameters** | [**LogSearchParameters**](LogSearchParameters.md) |  | 

### Return type

[**MpkiApiV1AuditLogGet200Response**](MpkiApiV1AuditLogGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1AuditLogIdGet

> AuditLogItem MpkiApiV1AuditLogIdGet(ctx, id).Execute()

Get audit log event by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "efcd5341-0fca-4fa3-9bc8-52e63c37e345" // string | Audit log event ID (displayed in the audit event details on UI, and in audit logs list response)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.MpkiApiV1AuditLogIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.MpkiApiV1AuditLogIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1AuditLogIdGet`: AuditLogItem
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.MpkiApiV1AuditLogIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Audit log event ID (displayed in the audit event details on UI, and in audit logs list response) | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1AuditLogIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLogItem**](AuditLogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1ReportEnrollmentCodeGet

> MpkiApiV1ReportEnrollmentCodeGet200Response MpkiApiV1ReportEnrollmentCodeGet(ctx).PagingParameters(pagingParameters).SearchParameters(searchParameters).Execute()

Get enrollment code report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pagingParameters := map[string][]openapiclient.PagingParameters{ ... } // PagingParameters |  (optional)
    searchParameters := map[string][]openapiclient.ReportSearchParams{ ... } // ReportSearchParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.MpkiApiV1ReportEnrollmentCodeGet(context.Background()).PagingParameters(pagingParameters).SearchParameters(searchParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.MpkiApiV1ReportEnrollmentCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1ReportEnrollmentCodeGet`: MpkiApiV1ReportEnrollmentCodeGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.MpkiApiV1ReportEnrollmentCodeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1ReportEnrollmentCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagingParameters** | [**PagingParameters**](PagingParameters.md) |  | 
 **searchParameters** | [**ReportSearchParams**](ReportSearchParams.md) |  | 

### Return type

[**MpkiApiV1ReportEnrollmentCodeGet200Response**](MpkiApiV1ReportEnrollmentCodeGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

