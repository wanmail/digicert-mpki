# \SeatsApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1SeatPost**](SeatsApi.md#MpkiApiV1SeatPost) | **Post** /mpki/api/v1/seat | Create seat
[**MpkiApiV1SeatSeatIdDelete**](SeatsApi.md#MpkiApiV1SeatSeatIdDelete) | **Delete** /mpki/api/v1/seat/{seat_id} | Delete seat
[**MpkiApiV1SeatSeatIdGet**](SeatsApi.md#MpkiApiV1SeatSeatIdGet) | **Get** /mpki/api/v1/seat/{seat_id} | Get seat by ID
[**MpkiApiV1SeatSeatIdPut**](SeatsApi.md#MpkiApiV1SeatSeatIdPut) | **Put** /mpki/api/v1/seat/{seat_id} | Update seat
[**MpkiApiV1SeatTypesGet**](SeatsApi.md#MpkiApiV1SeatTypesGet) | **Get** /mpki/api/v1/seat-types | Get available seat types



## MpkiApiV1SeatPost

> SeatDetails MpkiApiV1SeatPost(ctx).MpkiApiV1SeatPostRequest(mpkiApiV1SeatPostRequest).Execute()

Create seat



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
    mpkiApiV1SeatPostRequest := *openapiclient.NewMpkiApiV1SeatPostRequest() // MpkiApiV1SeatPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeatsApi.MpkiApiV1SeatPost(context.Background()).MpkiApiV1SeatPostRequest(mpkiApiV1SeatPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.MpkiApiV1SeatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1SeatPost`: SeatDetails
    fmt.Fprintf(os.Stdout, "Response from `SeatsApi.MpkiApiV1SeatPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1SeatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mpkiApiV1SeatPostRequest** | [**MpkiApiV1SeatPostRequest**](MpkiApiV1SeatPostRequest.md) |  | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1SeatSeatIdDelete

> SeatDetails MpkiApiV1SeatSeatIdDelete(ctx, seatId).SeatTypeId(seatTypeId).BusinessUnitId(businessUnitId).Execute()

Delete seat



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
    seatId := "some-seat-id@example.com" // string | Seat ID
    seatTypeId := "USER_SEAT" // string | Seat Type ID
    businessUnitId := "8f3764a4-678c-469d-8699-0308552ec256" // string | Business unit ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeatsApi.MpkiApiV1SeatSeatIdDelete(context.Background(), seatId).SeatTypeId(seatTypeId).BusinessUnitId(businessUnitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.MpkiApiV1SeatSeatIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1SeatSeatIdDelete`: SeatDetails
    fmt.Fprintf(os.Stdout, "Response from `SeatsApi.MpkiApiV1SeatSeatIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seatId** | **string** | Seat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1SeatSeatIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seatTypeId** | **string** | Seat Type ID | 
 **businessUnitId** | **string** | Business unit ID | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1SeatSeatIdGet

> MpkiApiV1SeatSeatIdDelete200Response MpkiApiV1SeatSeatIdGet(ctx, seatId).SeatTypeId(seatTypeId).BusinessUnitId(businessUnitId).Execute()

Get seat by ID



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
    seatId := "some-seat-id@example.com" // string | Seat ID
    seatTypeId := "seatTypeId_example" // string | Seat Type ID
    businessUnitId := "8caedac5-8669-4f9a-8c42-44825d154577" // string | Business unit ID.  As identical seat IDs may exist across various business units of the same account, the `business_unit_id` for the given seat must be specified as a query parameter if the account contains more than one business unit.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeatsApi.MpkiApiV1SeatSeatIdGet(context.Background(), seatId).SeatTypeId(seatTypeId).BusinessUnitId(businessUnitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.MpkiApiV1SeatSeatIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1SeatSeatIdGet`: MpkiApiV1SeatSeatIdDelete200Response
    fmt.Fprintf(os.Stdout, "Response from `SeatsApi.MpkiApiV1SeatSeatIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seatId** | **string** | Seat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1SeatSeatIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seatTypeId** | **string** | Seat Type ID | 
 **businessUnitId** | **string** | Business unit ID.  As identical seat IDs may exist across various business units of the same account, the &#x60;business_unit_id&#x60; for the given seat must be specified as a query parameter if the account contains more than one business unit.  | 

### Return type

[**MpkiApiV1SeatSeatIdDelete200Response**](MpkiApiV1SeatSeatIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1SeatSeatIdPut

> SeatDetails MpkiApiV1SeatSeatIdPut(ctx, seatId).MpkiApiV1SeatSeatIdDeleteRequest(mpkiApiV1SeatSeatIdDeleteRequest).Execute()

Update seat



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
    seatId := "some-seat-id@example.com" // string | Seat ID
    mpkiApiV1SeatSeatIdDeleteRequest := *openapiclient.NewMpkiApiV1SeatSeatIdDeleteRequest() // MpkiApiV1SeatSeatIdDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeatsApi.MpkiApiV1SeatSeatIdPut(context.Background(), seatId).MpkiApiV1SeatSeatIdDeleteRequest(mpkiApiV1SeatSeatIdDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.MpkiApiV1SeatSeatIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1SeatSeatIdPut`: SeatDetails
    fmt.Fprintf(os.Stdout, "Response from `SeatsApi.MpkiApiV1SeatSeatIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seatId** | **string** | Seat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1SeatSeatIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mpkiApiV1SeatSeatIdDeleteRequest** | [**MpkiApiV1SeatSeatIdDeleteRequest**](MpkiApiV1SeatSeatIdDeleteRequest.md) |  | 

### Return type

[**SeatDetails**](SeatDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1SeatTypesGet

> []MpkiApiV1SeatTypesGet200ResponseInner MpkiApiV1SeatTypesGet(ctx).Execute()

Get available seat types



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
    resp, r, err := apiClient.SeatsApi.MpkiApiV1SeatTypesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.MpkiApiV1SeatTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1SeatTypesGet`: []MpkiApiV1SeatTypesGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SeatsApi.MpkiApiV1SeatTypesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1SeatTypesGetRequest struct via the builder pattern


### Return type

[**[]MpkiApiV1SeatTypesGet200ResponseInner**](MpkiApiV1SeatTypesGet200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

