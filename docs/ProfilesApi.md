# {{classname}}

All URIs are relative to *https://one.digicert.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1ProfileGet**](ProfilesApi.md#MpkiApiV1ProfileGet) | **Get** /mpki/api/v1/profile | List profiles
[**MpkiApiV1ProfileProfileIdGet**](ProfilesApi.md#MpkiApiV1ProfileProfileIdGet) | **Get** /mpki/api/v1/profile/{profile_id} | Get profile by ID

# **MpkiApiV1ProfileGet**
> [][]PublicProfileResponse MpkiApiV1ProfileGet(ctx, optional)
List profiles

Use this endpoint to list all available certificate profiles in the account.   Results can optionally be limited to profiles bound to a specific API token — append the `api_token_id` to the endpoint URL as a query tag (`?api_token_id={id}`). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProfilesApiMpkiApiV1ProfileGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProfilesApiMpkiApiV1ProfileGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTokenId** | **optional.String**| Limit results to profiles bound to the given API token ID | 

### Return type

[**[][]PublicProfileResponse**](array.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1ProfileProfileIdGet**
> PublicProfileResponse MpkiApiV1ProfileProfileIdGet(ctx, profileId)
Get profile by ID

Use this endpoint to get the details for a given certificate profile ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Certificate profile ID | 

### Return type

[**PublicProfileResponse**](PublicProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

