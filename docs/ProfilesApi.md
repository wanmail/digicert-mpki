# \ProfilesApi

All URIs are relative to *https://one.digicert.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1ProfileGet**](ProfilesApi.md#MpkiApiV1ProfileGet) | **Get** /mpki/api/v1/profile | List profiles
[**MpkiApiV1ProfileProfileIdGet**](ProfilesApi.md#MpkiApiV1ProfileProfileIdGet) | **Get** /mpki/api/v1/profile/{profile_id} | Get profile by ID



## MpkiApiV1ProfileGet

> [][]PublicProfileResponse MpkiApiV1ProfileGet(ctx).ApiTokenId(apiTokenId).Execute()

List profiles



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
    apiTokenId := "apiTokenId_example" // string | Limit results to profiles bound to the given API token ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.MpkiApiV1ProfileGet(context.Background()).ApiTokenId(apiTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.MpkiApiV1ProfileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1ProfileGet`: [][]PublicProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.MpkiApiV1ProfileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1ProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTokenId** | **string** | Limit results to profiles bound to the given API token ID | 

### Return type

[**[][]PublicProfileResponse**](array.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MpkiApiV1ProfileProfileIdGet

> PublicProfileResponse MpkiApiV1ProfileProfileIdGet(ctx, profileId).Execute()

Get profile by ID



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
    profileId := "d6743d9a-fb67-4a64-a33e-8a3229d28cc4" // string | Certificate profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.MpkiApiV1ProfileProfileIdGet(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.MpkiApiV1ProfileProfileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MpkiApiV1ProfileProfileIdGet`: PublicProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.MpkiApiV1ProfileProfileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Certificate profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMpkiApiV1ProfileProfileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicProfileResponse**](PublicProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

