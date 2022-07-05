# {{classname}}

All URIs are relative to *https://one.digicert.com/*

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

# **MpkiApiV1CertificateEscrowApproveSerialNumberPost**
> MpkiApiV1CertificateEscrowApproveSerialNumberPost(ctx, serialNumber, optional)
Approve escrowed certificate recovery

Use this endpoint to approve a certificate recovery request for a given escrowed certificate serial number.  Certificate recovery approval requires the `RECOVER_EM_CERTIFICATE` user permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 
 **optional** | ***CertificatesApiMpkiApiV1CertificateEscrowApproveSerialNumberPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateEscrowApproveSerialNumberPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbprintSha256** | **optional.String**| Certificate thumbprint | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateEscrowRecoverSerialNumberGet**
> CertificateResponse MpkiApiV1CertificateEscrowRecoverSerialNumberGet(ctx, serialNumber, optional)
Recover escrowed certificate by serial number

Use this endpoint to recover an escrowed certificate by the given serial number.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 
 **optional** | ***CertificatesApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateEscrowRecoverSerialNumberGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeChain** | **optional.Bool**| If true, include CA certificates in response | [default to false]

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateImportPost**
> []InlineResponse2002 MpkiApiV1CertificateImportPost(ctx, optional)
Import certificate

Use this endpoint to import external certificates into your account. External certificates are issued by an issuing CA which does not originate from your DigiCert ONE account.   Certificates issued by an external issuing CA which has also been imported to your account are categorized as imported certificates and can be managed accordingly.  Imported certificates can be downloaded, revoked, suspended/resumed, and can make use of OCSP/CRL validation services in your account.   Certificates issued by an issuing CA which has not been imported to your account are categorized as unmanaged certificates.  Unmanaged certificates can be downloaded from your account, but no certificate lifecycle operations are available for unmanaged seats.   Certificates imported with `revoked: true` and `reason: certificate_hold` will be imported with certificate_status `SUSPENDED`.   Revoked certificates receive the certificate_status `REVOKED` when imported — any `reason` for revocation other than `certificate_hold` is replaced with `unspecified`.   Unmanaged certificates cannot be suspended.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiMpkiApiV1CertificateImportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateImportPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []V1CertificateimportBody**](v1_certificateimport_body.md)|  | 
 **trapSeatDuplicateError** | **optional.**| If true, automatically appends the current timestamp to the seat ID value. This feature facilitates duplicate seat ID error handling.   | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificatePost**
> CertificateResponse MpkiApiV1CertificatePost(ctx, optional)
Issue certificate

Use this endpoint to issue a new certificate using the given certificate and profile details.   Instant certificate issuance requires a certificate profile configured with `enrollment_method: rest_api` and `authentication_method: third_party_app`.   All Subject DN and SAN fields configured in the certificate profile with `REST Request` selected as the field value source and `Required` checked must be included in your request body.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiMpkiApiV1CertificatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1CertificateBody**](V1CertificateBody.md)|  | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateSearchGet**
> InlineResponse2003 MpkiApiV1CertificateSearchGet(ctx, optional)
Search certificates

Use this endpoint to get a detailed list of certificates in your account.  Results can be filtered using optional search and paging parameters appended to the endpoint URL as query tags (e.g. `?limit=10`).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiMpkiApiV1CertificateSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateSearchGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchParameters** | [**optional.Interface of CertificateSearchParameters**](.md)|  | 
 **pagingParameters** | [**optional.Interface of PagingParameters**](.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateSerialNumberGet**
> PublicCertificateDetails MpkiApiV1CertificateSerialNumberGet(ctx, serialNumber)
Get certificate by serial number

Use this endpoint to get the details of a given certificate serial number.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 

### Return type

[**PublicCertificateDetails**](PublicCertificateDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateSerialNumberRenewPost**
> CertificateResponse MpkiApiV1CertificateSerialNumberRenewPost(ctx, serialNumber, optional)
Renew certificate

Use this endpoint to renew the certificate with the given serial number.   Only certificates from certificate profiles configured with `enrollment_method: rest_api` and `authentication_method: third_party_app` can be renewed using this endpoint.   Certificates must be within the renewal period configured in the certificate profile (default is 30 days before expiration) to be eligible for renewal.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 
 **optional** | ***CertificatesApiMpkiApiV1CertificateSerialNumberRenewPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateSerialNumberRenewPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SerialNumberRenewBody**](SerialNumberRenewBody.md)|  | 
 **thumbprintSha256** | **optional.**| Certificate thumbprint | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateSerialNumberRevokeDelete**
> MpkiApiV1CertificateSerialNumberRevokeDelete(ctx, serialNumber)
Resume suspended certificate

Use this endpoint to resume the previously suspended certificate with the given serial number.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpkiApiV1CertificateSerialNumberRevokePut**
> MpkiApiV1CertificateSerialNumberRevokePut(ctx, serialNumber, optional)
Revoke certificate

Use this endpoint to revoke a certificate by serial number.   Include the `revocation_reason` in the request body — accepted values are:   `unspecified`, `key_compromise`, `ca_compromise`, `affiliation_changed`, `superseded`, `cessation_of_operation`, `certificate_hold`, `remove_from_crl`, `privilege_withdrawn`, or `aa_compromise` .   **Note:** If multiple certificates with the same serial number exist in your account, an error response with the message **\"Collision occurred! Please specify additional parameter: thumbprint_sha256\"** is returned. Retry the request with the SHA256 thumbprint of the certificate passed as a query tag (append `?thumbprint_sha256={your thumbprint}` to the endpoint URL). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialNumber** | **string**| Certificate serial number | 
 **optional** | ***CertificatesApiMpkiApiV1CertificateSerialNumberRevokePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiMpkiApiV1CertificateSerialNumberRevokePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SerialNumberRevokeBody**](SerialNumberRevokeBody.md)|  | 
 **thumbprintSha256** | **optional.**| Certificate thumbprint | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

