# {{classname}}

All URIs are relative to *https://one.digicert.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpkiApiV1HealthExtensiveGet**](ServiceStatusApi.md#MpkiApiV1HealthExtensiveGet) | **Get** /mpki/api/v1/health/extensive | Query service health
[**MpkiApiV1HelloGet**](ServiceStatusApi.md#MpkiApiV1HelloGet) | **Get** /mpki/api/v1/hello | Query service status

# **MpkiApiV1HealthExtensiveGet**
> string MpkiApiV1HealthExtensiveGet(ctx, )
Query service health

Use this endpoint to perform a health check to query service status. An authorized request returns either HTTP status 200 - \"UP\" or 500 - \"DOWN\". 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1HelloGet**
> string MpkiApiV1HelloGet(ctx, )
Query service status

Use this endpoint to query service status and test connectivity to the API. A successful response indicates that the service is live. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

