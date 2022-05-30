# \ServiceStatusApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1HealthExtensiveGet**](ServiceStatusApi.md#MpkiApiV1HealthExtensiveGet) | **Get** /mpki/api/v1/health/extensive | Query service health
[**MpkiApiV1HelloGet**](ServiceStatusApi.md#MpkiApiV1HelloGet) | **Get** /mpki/api/v1/hello | Query service status



## MpkiApiV1HealthExtensiveGet

> string MpkiApiV1HealthExtensiveGet(ctx).Execute()

Query service health



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceStatusApi.MpkiApiV1HealthExtensiveGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceStatusApi.MpkiApiV1HealthExtensiveGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1HealthExtensiveGet`: string
    fmt.Fprintf(os.Stdout, "Response from `ServiceStatusApi.MpkiApiV1HealthExtensiveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1HealthExtensiveGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1HelloGet

> string MpkiApiV1HelloGet(ctx).Execute()

Query service status



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceStatusApi.MpkiApiV1HelloGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceStatusApi.MpkiApiV1HelloGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1HelloGet`: string
    fmt.Fprintf(os.Stdout, "Response from `ServiceStatusApi.MpkiApiV1HelloGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1HelloGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

