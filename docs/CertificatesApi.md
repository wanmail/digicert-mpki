# \CertificatesApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1CertificateEscrowApproveSerialNumberPost**](CertificatesApi.md#MpkiApiV1CertificateEscrowApproveSerialNumberPost) | **Post** /mpki/api/v1/certificate/escrow/approve/{serial_number} | Approve escrowed certificate recovery
[**MpkiApiV1CertificateEscrowRecoverSerialNumberGet**](CertificatesApi.md#MpkiApiV1CertificateEscrowRecoverSerialNumberGet) | **Get** /mpki/api/v1/certificate/escrow/recover/{serial_number} | Recover escrowed certificate by serial number
[**MpkiApiV1CertificateImportPost**](CertificatesApi.md#MpkiApiV1CertificateImportPost) | **Post** /mpki/api/v1/certificate-import | Import certificate
[**MpkiApiV1CertificatePost**](CertificatesApi.md#MpkiApiV1CertificatePost) | **Post** /mpki/api/v1/certificate | Issue certificate
[**MpkiApiV1CertificateSearchGet**](CertificatesApi.md#MpkiApiV1CertificateSearchGet) | **Get** /mpki/api/v1/certificate-search | Search certificates
[**MpkiApiV1CertificateSerialNumberGet**](CertificatesApi.md#MpkiApiV1CertificateSerialNumberGet) | **Get** /mpki/api/v1/certificate/{serial_number} | Get certificate by serial number
[**MpkiApiV1CertificateSerialNumberRenewPost**](CertificatesApi.md#MpkiApiV1CertificateSerialNumberRenewPost) | **Post** /mpki/api/v1/certificate/{serial_number}/renew | Renew certificate
[**MpkiApiV1CertificateSerialNumberRevokeDelete**](CertificatesApi.md#MpkiApiV1CertificateSerialNumberRevokeDelete) | **Delete** /mpki/api/v1/certificate/{serial_number}/revoke | Resume suspended certificate
[**MpkiApiV1CertificateSerialNumberRevokePut**](CertificatesApi.md#MpkiApiV1CertificateSerialNumberRevokePut) | **Put** /mpki/api/v1/certificate/{serial_number}/revoke | Revoke certificate



## MpkiApiV1CertificateEscrowApproveSerialNumberPost

> MpkiApiV1CertificateEscrowApproveSerialNumberPost(ctx, serialNumber).ThumbprintSha256(thumbprintSha256).Execute()

Approve escrowed certificate recovery



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number
    thumbprintSha256 := "ED9707A73630676168E3D531F622082B422CEEC1955706C1DC75AEEABDB80BCA" // string | Certificate thumbprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateEscrowApproveSerialNumberPost(context.Background(), serialNumber).ThumbprintSha256(thumbprintSha256).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateEscrowApproveSerialNumberPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateEscrowApproveSerialNumberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbprintSha256** | **string** | Certificate thumbprint | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificateEscrowRecoverSerialNumberGet

> CertificateResponse MpkiApiV1CertificateEscrowRecoverSerialNumberGet(ctx, serialNumber).IncludeChain(includeChain).Execute()

Recover escrowed certificate by serial number



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number
    includeChain := true // bool | If true, include CA certificates in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateEscrowRecoverSerialNumberGet(context.Background(), serialNumber).IncludeChain(includeChain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateEscrowRecoverSerialNumberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificateEscrowRecoverSerialNumberGet`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificateEscrowRecoverSerialNumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateEscrowRecoverSerialNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeChain** | **bool** | If true, include CA certificates in response | [default to false]

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificateImportPost

> []MpkiApiV1CertificateImportPost200ResponseInner MpkiApiV1CertificateImportPost(ctx).TrapSeatDuplicateError(trapSeatDuplicateError).MpkiApiV1CertificateImportPostRequestInner(mpkiApiV1CertificateImportPostRequestInner).Execute()

Import certificate



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
    trapSeatDuplicateError := true // bool | If true, automatically appends the current timestamp to the seat ID value. This feature facilitates duplicate seat ID error handling.   (optional)
    mpkiApiV1CertificateImportPostRequestInner := []openapiclient.MpkiApiV1CertificateImportPostRequestInner{*openapiclient.NewMpkiApiV1CertificateImportPostRequestInner()} // []MpkiApiV1CertificateImportPostRequestInner |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateImportPost(context.Background()).TrapSeatDuplicateError(trapSeatDuplicateError).MpkiApiV1CertificateImportPostRequestInner(mpkiApiV1CertificateImportPostRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificateImportPost`: []MpkiApiV1CertificateImportPost200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificateImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trapSeatDuplicateError** | **bool** | If true, automatically appends the current timestamp to the seat ID value. This feature facilitates duplicate seat ID error handling.   | 
 **mpkiApiV1CertificateImportPostRequestInner** | [**[]MpkiApiV1CertificateImportPostRequestInner**](MpkiApiV1CertificateImportPostRequestInner.md) |  | 

### Return type

[**[]MpkiApiV1CertificateImportPost200ResponseInner**](MpkiApiV1CertificateImportPost200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificatePost

> CertificateResponse MpkiApiV1CertificatePost(ctx).MpkiApiV1CertificatePostRequest(mpkiApiV1CertificatePostRequest).Execute()

Issue certificate



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
    mpkiApiV1CertificatePostRequest := *openapiclient.NewMpkiApiV1CertificatePostRequest() // MpkiApiV1CertificatePostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificatePost(context.Background()).MpkiApiV1CertificatePostRequest(mpkiApiV1CertificatePostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificatePost`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mpkiApiV1CertificatePostRequest** | [**MpkiApiV1CertificatePostRequest**](MpkiApiV1CertificatePostRequest.md) |  | 

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


## MpkiApiV1CertificateSearchGet

> MpkiApiV1CertificateSearchGet200Response MpkiApiV1CertificateSearchGet(ctx).SearchParameters(searchParameters).PagingParameters(pagingParameters).Execute()

Search certificates



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
    searchParameters := map[string][]openapiclient.CertificateSearchParameters{ ... } // CertificateSearchParameters |  (optional)
    pagingParameters := map[string][]openapiclient.PagingParameters{ ... } // PagingParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateSearchGet(context.Background()).SearchParameters(searchParameters).PagingParameters(pagingParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificateSearchGet`: MpkiApiV1CertificateSearchGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificateSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchParameters** | [**CertificateSearchParameters**](CertificateSearchParameters.md) |  | 
 **pagingParameters** | [**PagingParameters**](PagingParameters.md) |  | 

### Return type

[**MpkiApiV1CertificateSearchGet200Response**](MpkiApiV1CertificateSearchGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificateSerialNumberGet

> PublicCertificateDetails MpkiApiV1CertificateSerialNumberGet(ctx, serialNumber).Execute()

Get certificate by serial number



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateSerialNumberGet(context.Background(), serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateSerialNumberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificateSerialNumberGet`: PublicCertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificateSerialNumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateSerialNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicCertificateDetails**](PublicCertificateDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificateSerialNumberRenewPost

> CertificateResponse MpkiApiV1CertificateSerialNumberRenewPost(ctx, serialNumber).ThumbprintSha256(thumbprintSha256).MpkiApiV1CertificateSerialNumberRenewPostRequest(mpkiApiV1CertificateSerialNumberRenewPostRequest).Execute()

Renew certificate



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number
    thumbprintSha256 := "ED9707A73630676168E3D531F622082B422CEEC1955706C1DC75AEEABDB80BCA" // string | Certificate thumbprint (optional)
    mpkiApiV1CertificateSerialNumberRenewPostRequest := *openapiclient.NewMpkiApiV1CertificateSerialNumberRenewPostRequest() // MpkiApiV1CertificateSerialNumberRenewPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateSerialNumberRenewPost(context.Background(), serialNumber).ThumbprintSha256(thumbprintSha256).MpkiApiV1CertificateSerialNumberRenewPostRequest(mpkiApiV1CertificateSerialNumberRenewPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateSerialNumberRenewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1CertificateSerialNumberRenewPost`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.MpkiApiV1CertificateSerialNumberRenewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateSerialNumberRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbprintSha256** | **string** | Certificate thumbprint | 
 **mpkiApiV1CertificateSerialNumberRenewPostRequest** | [**MpkiApiV1CertificateSerialNumberRenewPostRequest**](MpkiApiV1CertificateSerialNumberRenewPostRequest.md) |  | 

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


## MpkiApiV1CertificateSerialNumberRevokeDelete

> MpkiApiV1CertificateSerialNumberRevokeDelete(ctx, serialNumber).Execute()

Resume suspended certificate



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateSerialNumberRevokeDelete(context.Background(), serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateSerialNumberRevokeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateSerialNumberRevokeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1CertificateSerialNumberRevokePut

> MpkiApiV1CertificateSerialNumberRevokePut(ctx, serialNumber).ThumbprintSha256(thumbprintSha256).MpkiApiV1CertificateSerialNumberRevokeDeleteRequest(mpkiApiV1CertificateSerialNumberRevokeDeleteRequest).Execute()

Revoke certificate



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
    serialNumber := "4099ABF9C1BB7C02BBE1CDC7836CA9D58A71C2A5" // string | Certificate serial number
    thumbprintSha256 := "ED9707A73630676168E3D531F622082B422CEEC1955706C1DC75AEEABDB80BCA" // string | Certificate thumbprint (optional)
    mpkiApiV1CertificateSerialNumberRevokeDeleteRequest := *openapiclient.NewMpkiApiV1CertificateSerialNumberRevokeDeleteRequest() // MpkiApiV1CertificateSerialNumberRevokeDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.MpkiApiV1CertificateSerialNumberRevokePut(context.Background(), serialNumber).ThumbprintSha256(thumbprintSha256).MpkiApiV1CertificateSerialNumberRevokeDeleteRequest(mpkiApiV1CertificateSerialNumberRevokeDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.MpkiApiV1CertificateSerialNumberRevokePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** | Certificate serial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1CertificateSerialNumberRevokePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbprintSha256** | **string** | Certificate thumbprint | 
 **mpkiApiV1CertificateSerialNumberRevokeDeleteRequest** | [**MpkiApiV1CertificateSerialNumberRevokeDeleteRequest**](MpkiApiV1CertificateSerialNumberRevokeDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

