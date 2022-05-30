# \EnrollmentsApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1EnrollmentDetailsGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentDetailsGet) | **Get** /mpki/api/v1/enrollment-details/ | List enrollments
[**MpkiApiV1EnrollmentDetailsIdentifierGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentDetailsIdentifierGet) | **Get** /mpki/api/v1/enrollment-details/{identifier} | Get enrollment details by ID
[**MpkiApiV1EnrollmentEnrollmentCodeGet**](EnrollmentsApi.md#MpkiApiV1EnrollmentEnrollmentCodeGet) | **Get** /mpki/api/v1/enrollment/{enrollmentCode} | Get enrollment details by enrollment code
[**MpkiApiV1EnrollmentPost**](EnrollmentsApi.md#MpkiApiV1EnrollmentPost) | **Post** /mpki/api/v1/enrollment | Create enrollment
[**MpkiApiV1EnrollmentRedeemPost**](EnrollmentsApi.md#MpkiApiV1EnrollmentRedeemPost) | **Post** /mpki/api/v1/enrollment/redeem | Redeem enrollment



## MpkiApiV1EnrollmentDetailsGet

> []PublicEnrollmentDetails MpkiApiV1EnrollmentDetailsGet(ctx).SearchParameters(searchParameters).Execute()

List enrollments



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
    searchParameters := map[string][]openapiclient.EnrollmentListParameters{ ... } // EnrollmentListParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.MpkiApiV1EnrollmentDetailsGet(context.Background()).SearchParameters(searchParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.MpkiApiV1EnrollmentDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1EnrollmentDetailsGet`: []PublicEnrollmentDetails
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.MpkiApiV1EnrollmentDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1EnrollmentDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchParameters** | [**EnrollmentListParameters**](EnrollmentListParameters.md) |  | 

### Return type

[**[]PublicEnrollmentDetails**](PublicEnrollmentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1EnrollmentDetailsIdentifierGet

> PublicEnrollmentDetails MpkiApiV1EnrollmentDetailsIdentifierGet(ctx, identifier).Execute()

Get enrollment details by ID



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
    identifier := "d6743d9a-fb67-4a64-a33e-8a3229d28cc4" // string | Enrollment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.MpkiApiV1EnrollmentDetailsIdentifierGet(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.MpkiApiV1EnrollmentDetailsIdentifierGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1EnrollmentDetailsIdentifierGet`: PublicEnrollmentDetails
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.MpkiApiV1EnrollmentDetailsIdentifierGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Enrollment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1EnrollmentDetailsIdentifierGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicEnrollmentDetails**](PublicEnrollmentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1EnrollmentEnrollmentCodeGet

> EnrollmentResponse MpkiApiV1EnrollmentEnrollmentCodeGet(ctx, enrollmentCode).SeatId(seatId).Execute()

Get enrollment details by enrollment code



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
    enrollmentCode := "662641677" // string | Enrollment code received in enrollment creation response
    seatId := "some-seat-id@example.com" // string | Seat ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.MpkiApiV1EnrollmentEnrollmentCodeGet(context.Background(), enrollmentCode).SeatId(seatId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.MpkiApiV1EnrollmentEnrollmentCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1EnrollmentEnrollmentCodeGet`: EnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.MpkiApiV1EnrollmentEnrollmentCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enrollmentCode** | **string** | Enrollment code received in enrollment creation response | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1EnrollmentEnrollmentCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seatId** | **string** | Seat ID | 

### Return type

[**EnrollmentResponse**](EnrollmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1EnrollmentPost

> EnrollmentResponse MpkiApiV1EnrollmentPost(ctx).MpkiApiV1EnrollmentPostRequest(mpkiApiV1EnrollmentPostRequest).Execute()

Create enrollment



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
    mpkiApiV1EnrollmentPostRequest := *openapiclient.NewMpkiApiV1EnrollmentPostRequest("Profile_example", *openapiclient.NewMpkiApiV1EnrollmentPostRequestSeat("SeatId_example", "Email_example"), map[string]interface{}(123)) // MpkiApiV1EnrollmentPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.MpkiApiV1EnrollmentPost(context.Background()).MpkiApiV1EnrollmentPostRequest(mpkiApiV1EnrollmentPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.MpkiApiV1EnrollmentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1EnrollmentPost`: EnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.MpkiApiV1EnrollmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1EnrollmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mpkiApiV1EnrollmentPostRequest** | [**MpkiApiV1EnrollmentPostRequest**](MpkiApiV1EnrollmentPostRequest.md) |  | 

### Return type

[**EnrollmentResponse**](EnrollmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1EnrollmentRedeemPost

> CertificateResponse MpkiApiV1EnrollmentRedeemPost(ctx).MpkiApiV1EnrollmentRedeemPostRequest(mpkiApiV1EnrollmentRedeemPostRequest).Execute()

Redeem enrollment



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
    mpkiApiV1EnrollmentRedeemPostRequest := *openapiclient.NewMpkiApiV1EnrollmentRedeemPostRequest() // MpkiApiV1EnrollmentRedeemPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.MpkiApiV1EnrollmentRedeemPost(context.Background()).MpkiApiV1EnrollmentRedeemPostRequest(mpkiApiV1EnrollmentRedeemPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.MpkiApiV1EnrollmentRedeemPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1EnrollmentRedeemPost`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.MpkiApiV1EnrollmentRedeemPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1EnrollmentRedeemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mpkiApiV1EnrollmentRedeemPostRequest** | [**MpkiApiV1EnrollmentRedeemPostRequest**](MpkiApiV1EnrollmentRedeemPostRequest.md) |  | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

