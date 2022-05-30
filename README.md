# Go API client for openapi


Welcome to the DigiCert ONE Enterprise PKI Manager REST API. The Enterprise PKI Manager API service provides operations for managing enrollments and certificates for users and devices, retrieving certificate profiles and seats, and account event reporting.

## Base URL

The base URL path for endpoints in the Enterprise PKI Manager REST API is: `{server}/mpki/api/v1`.

Replace `{server}` with the hostname of your DigiCert ONE instance. For example, if you are using the DigiCert ONE hosted cloud solution, your `{server}` is `https://one.digicert.com`.

## Authentication

API clients can authenticate to endpoints in the Enterprise PKI Manager REST API using these methods:
* Header-based API token authentication.
* Authentication using a client authentication certificate.

### API token

To authenticate with an API token, include the custom HTTP header `x-api-key` in your request. Use one of these values in the `x-api-key` header:
* A Service user token ID (**recommended**).    
* An API token bound to a DigiCert ONE administrator.

**Note:** We recommend that you dedicate a Service user token ID to API operations as this distinguishes API requests from administrator actions in your account audit logs. 

**Service user token ID**
* Service users are API tokens that are not associated with a specific user.
* When you create a service user, you assign only the permissions needed for the API integration.
* There are two ways to create a new service user:

  * 1- Use the Account Manager in DigiCert ONE: 
    1. Navigate to the DigiCert ONE **Account Manager**.
    2. Select **Access** from the left menu.
    3. Select **Service User** from the left menu.
    4. Select **Create service user** and follow the remaining prompts.

  * 2- Make a request to the `/account/api/v1/user` endpoint in the DigiCert ONE Account Manager API service.


**Administrator API token**

* Standard users (administrators) can create API tokens in DigiCert ONE.

* API tokens have the same permissions and access scope as the administrator that created them.

* Actions linked to the API token are logged under the administrator's name in event audit logs.
* To generate a new API token: 

  1. Navigate to the DigiCert ONE **Account Manager**.
  2. Select **Access** from the left menu.
  3. Select **Administrators**.
  4. Select the administrator you wish to create an API token for from the list.
  5. Scroll down to the **API Tokens** section on the administrator page. 
  6. Select **Create API token** and follow the remaining prompts. 


### Client authentication certificate

When authenticating with a client authentication certificate, you present a trusted certificate in your request. Both DigiCert ONE administrators and service users can use client authentication certificates.

To generate a client authentication certificate:

  1. Navigate to the DigiCert ONE **Account Manager**.
  2. Select **Access** from the left menu. 
  3. Search for and select your Service user. 
  4. Scroll to the **Authentication certificates** section and select **Create authentication certificate**.
  5. Enter a **Nickname** and select an **End date** for the certificate. 
  6. Select **Generate certificate**. Copy the password that is displayed (this is only displayed once and is required to decrypt the PKCS12 certificate file) and select **Download certificate**.
  7. Confirm that you have downloaded the certificate file (**Certificate_pkcs12.p12**) and that you can successfully decrypt it with the provided password, then click **Close** in the certificate download dialog box.
  8. Review the permissions you wish to provide to the Service user that the certificate is associated with.
  
  **Note**: We recommend that you assign permissions according to the principle of least privilege, i.e. you assign the minimum permissions needed to perform the required tasks.  

To use a client authentication certificate:

* Include the certificate in your API request.
* In the base URL for the endpoint path, add the prefix `clientauth`. For example: `https://clientauth.one.digicert.com`
* Omit the `x-api-key` header.

## Requests

The Enterprise PKI Manager API service accepts REST calls on HTTP/HTTPS ports 80/443. All requests are submitted using RESTful URLs and REST features, including header-based authentication and JSON request types. The data character set encoding for requests is UTF-8.

A well-formed request uses port 443 and specifies the user-agent and content-length HTTP headers. Each request consists of a method and an endpoint. Some requests also include a request payload if relevant to the operation being performed.

### Method

The Enterprise PKI Manager API uses these standard HTTP methods:

* GET
* POST
* PUT
* DELETE

### Body and content type

All requests that accept a body require passing in JSON formatted data with the `Content-Type` header set to `application/json`.

GET requests do not require passing formatted data in the request payload. However, some GET operations allow you to filter the results by providing additional path parameters or URL query strings (appended to the endpoint URL, e.g. `?limit=50`).

## Responses

Each response consists of a header and a body. The body is formatted based on the content type requested in the `Accept` header.

**Note:** Enterprise PKI Manager API endpoints only support responses with a content type of `application/json`. Requests that use the `Accept` header to specify a different content type will fail.

### Headers

Each response includes a header with a response code based on [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) specifications.

* HTTP codes in the **200-399** range describe a successful request. Response bodies for HTTP codes in this range include the response data associated with the operation.
* HTTP codes in the **400+** range describe an error.

Unsuccessful requests return a list with one or more `errors`. Each error object includes a `code` and a `message` describing the problem with the request.

**Example error response**

```JSON
{
  \"error\": {
    \"code\": \"no_access\",
    \"message\": \"User has no access to the Business unit with id = xxxxx or such Business unit does not exist\"
  }
}
```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://one.digicert.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CertificatesApi* | [**MpkiApiV1CertificateEscrowApproveSerialNumberPost**](docs/CertificatesApi.md#mpkiapiv1certificateescrowapproveserialnumberpost) | **Post** /mpki/api/v1/certificate/escrow/approve/{serial_number} | Approve escrowed certificate recovery
*CertificatesApi* | [**MpkiApiV1CertificateEscrowRecoverSerialNumberGet**](docs/CertificatesApi.md#mpkiapiv1certificateescrowrecoverserialnumberget) | **Get** /mpki/api/v1/certificate/escrow/recover/{serial_number} | Recover escrowed certificate by serial number
*CertificatesApi* | [**MpkiApiV1CertificateImportPost**](docs/CertificatesApi.md#mpkiapiv1certificateimportpost) | **Post** /mpki/api/v1/certificate-import | Import certificate
*CertificatesApi* | [**MpkiApiV1CertificatePost**](docs/CertificatesApi.md#mpkiapiv1certificatepost) | **Post** /mpki/api/v1/certificate | Issue certificate
*CertificatesApi* | [**MpkiApiV1CertificateSearchGet**](docs/CertificatesApi.md#mpkiapiv1certificatesearchget) | **Get** /mpki/api/v1/certificate-search | Search certificates
*CertificatesApi* | [**MpkiApiV1CertificateSerialNumberGet**](docs/CertificatesApi.md#mpkiapiv1certificateserialnumberget) | **Get** /mpki/api/v1/certificate/{serial_number} | Get certificate by serial number
*CertificatesApi* | [**MpkiApiV1CertificateSerialNumberRenewPost**](docs/CertificatesApi.md#mpkiapiv1certificateserialnumberrenewpost) | **Post** /mpki/api/v1/certificate/{serial_number}/renew | Renew certificate
*CertificatesApi* | [**MpkiApiV1CertificateSerialNumberRevokeDelete**](docs/CertificatesApi.md#mpkiapiv1certificateserialnumberrevokedelete) | **Delete** /mpki/api/v1/certificate/{serial_number}/revoke | Resume suspended certificate
*CertificatesApi* | [**MpkiApiV1CertificateSerialNumberRevokePut**](docs/CertificatesApi.md#mpkiapiv1certificateserialnumberrevokeput) | **Put** /mpki/api/v1/certificate/{serial_number}/revoke | Revoke certificate
*EnrollmentsApi* | [**MpkiApiV1EnrollmentDetailsGet**](docs/EnrollmentsApi.md#mpkiapiv1enrollmentdetailsget) | **Get** /mpki/api/v1/enrollment-details/ | List enrollments
*EnrollmentsApi* | [**MpkiApiV1EnrollmentDetailsIdentifierGet**](docs/EnrollmentsApi.md#mpkiapiv1enrollmentdetailsidentifierget) | **Get** /mpki/api/v1/enrollment-details/{identifier} | Get enrollment details by ID
*EnrollmentsApi* | [**MpkiApiV1EnrollmentEnrollmentCodeGet**](docs/EnrollmentsApi.md#mpkiapiv1enrollmentenrollmentcodeget) | **Get** /mpki/api/v1/enrollment/{enrollmentCode} | Get enrollment details by enrollment code
*EnrollmentsApi* | [**MpkiApiV1EnrollmentPost**](docs/EnrollmentsApi.md#mpkiapiv1enrollmentpost) | **Post** /mpki/api/v1/enrollment | Create enrollment
*EnrollmentsApi* | [**MpkiApiV1EnrollmentRedeemPost**](docs/EnrollmentsApi.md#mpkiapiv1enrollmentredeempost) | **Post** /mpki/api/v1/enrollment/redeem | Redeem enrollment
*ProfilesApi* | [**MpkiApiV1ProfileGet**](docs/ProfilesApi.md#mpkiapiv1profileget) | **Get** /mpki/api/v1/profile | List profiles
*ProfilesApi* | [**MpkiApiV1ProfileProfileIdGet**](docs/ProfilesApi.md#mpkiapiv1profileprofileidget) | **Get** /mpki/api/v1/profile/{profile_id} | Get profile by ID
*ReportingApi* | [**MpkiApiV1AuditLogGet**](docs/ReportingApi.md#mpkiapiv1auditlogget) | **Get** /mpki/api/v1/audit-log | List audit logs
*ReportingApi* | [**MpkiApiV1AuditLogIdGet**](docs/ReportingApi.md#mpkiapiv1auditlogidget) | **Get** /mpki/api/v1/audit-log/{id} | Get audit log event by ID
*ReportingApi* | [**MpkiApiV1ReportEnrollmentCodeGet**](docs/ReportingApi.md#mpkiapiv1reportenrollmentcodeget) | **Get** /mpki/api/v1/report/enrollment-code | Get enrollment code report
*SeatsApi* | [**MpkiApiV1SeatPost**](docs/SeatsApi.md#mpkiapiv1seatpost) | **Post** /mpki/api/v1/seat | Create seat
*SeatsApi* | [**MpkiApiV1SeatSeatIdDelete**](docs/SeatsApi.md#mpkiapiv1seatseatiddelete) | **Delete** /mpki/api/v1/seat/{seat_id} | Delete seat
*SeatsApi* | [**MpkiApiV1SeatSeatIdGet**](docs/SeatsApi.md#mpkiapiv1seatseatidget) | **Get** /mpki/api/v1/seat/{seat_id} | Get seat by ID
*SeatsApi* | [**MpkiApiV1SeatSeatIdPut**](docs/SeatsApi.md#mpkiapiv1seatseatidput) | **Put** /mpki/api/v1/seat/{seat_id} | Update seat
*SeatsApi* | [**MpkiApiV1SeatTypesGet**](docs/SeatsApi.md#mpkiapiv1seattypesget) | **Get** /mpki/api/v1/seat-types | Get available seat types
*ServiceStatusApi* | [**MpkiApiV1HealthExtensiveGet**](docs/ServiceStatusApi.md#mpkiapiv1healthextensiveget) | **Get** /mpki/api/v1/health/extensive | Query service health
*ServiceStatusApi* | [**MpkiApiV1HelloGet**](docs/ServiceStatusApi.md#mpkiapiv1helloget) | **Get** /mpki/api/v1/hello | Query service status


## Documentation For Models

 - [AeConfigResponse](docs/AeConfigResponse.md)
 - [AuditLogItem](docs/AuditLogItem.md)
 - [AuditLogItemAccount](docs/AuditLogItemAccount.md)
 - [CaResponse](docs/CaResponse.md)
 - [CertificateAttributes](docs/CertificateAttributes.md)
 - [CertificateExtensions](docs/CertificateExtensions.md)
 - [CertificateExtensionsSan](docs/CertificateExtensionsSan.md)
 - [CertificateReport](docs/CertificateReport.md)
 - [CertificateResponse](docs/CertificateResponse.md)
 - [CertificateSearchParameters](docs/CertificateSearchParameters.md)
 - [DeliveryFormat](docs/DeliveryFormat.md)
 - [EnrollmentListParameters](docs/EnrollmentListParameters.md)
 - [EnrollmentReport](docs/EnrollmentReport.md)
 - [EnrollmentResponse](docs/EnrollmentResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseWithField](docs/ErrorResponseWithField.md)
 - [ExternalEnrollmentReport](docs/ExternalEnrollmentReport.md)
 - [IdDto](docs/IdDto.md)
 - [KeyEscrowPolicy](docs/KeyEscrowPolicy.md)
 - [LogSearchParameters](docs/LogSearchParameters.md)
 - [MicrosoftAutoEnrollmentClientSettings](docs/MicrosoftAutoEnrollmentClientSettings.md)
 - [MicrosoftAutoEnrollmentClientSettingsCryptoProviders](docs/MicrosoftAutoEnrollmentClientSettingsCryptoProviders.md)
 - [MicrosoftAutoEnrollmentClientSettingsEnrollmentMode](docs/MicrosoftAutoEnrollmentClientSettingsEnrollmentMode.md)
 - [MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport](docs/MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport.md)
 - [MpkiApiV1AuditLogGet200Response](docs/MpkiApiV1AuditLogGet200Response.md)
 - [MpkiApiV1AuditLogGet200ResponseAllOf](docs/MpkiApiV1AuditLogGet200ResponseAllOf.md)
 - [MpkiApiV1CertificateImportPost200ResponseInner](docs/MpkiApiV1CertificateImportPost200ResponseInner.md)
 - [MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId](docs/MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId.md)
 - [MpkiApiV1CertificateImportPost200ResponseInnerSeat](docs/MpkiApiV1CertificateImportPost200ResponseInnerSeat.md)
 - [MpkiApiV1CertificateImportPostRequestInner](docs/MpkiApiV1CertificateImportPostRequestInner.md)
 - [MpkiApiV1CertificateImportPostRequestInnerRevocation](docs/MpkiApiV1CertificateImportPostRequestInnerRevocation.md)
 - [MpkiApiV1CertificatePostRequest](docs/MpkiApiV1CertificatePostRequest.md)
 - [MpkiApiV1CertificatePostRequestAttributes](docs/MpkiApiV1CertificatePostRequestAttributes.md)
 - [MpkiApiV1CertificatePostRequestSeat](docs/MpkiApiV1CertificatePostRequestSeat.md)
 - [MpkiApiV1CertificateSearchGet200Response](docs/MpkiApiV1CertificateSearchGet200Response.md)
 - [MpkiApiV1CertificateSearchGet200ResponseAllOf](docs/MpkiApiV1CertificateSearchGet200ResponseAllOf.md)
 - [MpkiApiV1CertificateSerialNumberRenewPostRequest](docs/MpkiApiV1CertificateSerialNumberRenewPostRequest.md)
 - [MpkiApiV1CertificateSerialNumberRevokeDeleteRequest](docs/MpkiApiV1CertificateSerialNumberRevokeDeleteRequest.md)
 - [MpkiApiV1EnrollmentPostRequest](docs/MpkiApiV1EnrollmentPostRequest.md)
 - [MpkiApiV1EnrollmentPostRequestSeat](docs/MpkiApiV1EnrollmentPostRequestSeat.md)
 - [MpkiApiV1EnrollmentRedeemPostRequest](docs/MpkiApiV1EnrollmentRedeemPostRequest.md)
 - [MpkiApiV1EnrollmentRedeemPostRequestSeat](docs/MpkiApiV1EnrollmentRedeemPostRequestSeat.md)
 - [MpkiApiV1ReportEnrollmentCodeGet200Response](docs/MpkiApiV1ReportEnrollmentCodeGet200Response.md)
 - [MpkiApiV1ReportEnrollmentCodeGet200ResponseAllOf](docs/MpkiApiV1ReportEnrollmentCodeGet200ResponseAllOf.md)
 - [MpkiApiV1SeatPostRequest](docs/MpkiApiV1SeatPostRequest.md)
 - [MpkiApiV1SeatPostRequestAllOf](docs/MpkiApiV1SeatPostRequestAllOf.md)
 - [MpkiApiV1SeatSeatIdDelete200Response](docs/MpkiApiV1SeatSeatIdDelete200Response.md)
 - [MpkiApiV1SeatSeatIdDelete200ResponseAllOf](docs/MpkiApiV1SeatSeatIdDelete200ResponseAllOf.md)
 - [MpkiApiV1SeatSeatIdDeleteRequest](docs/MpkiApiV1SeatSeatIdDeleteRequest.md)
 - [MpkiApiV1SeatSeatIdDeleteRequestAllOf](docs/MpkiApiV1SeatSeatIdDeleteRequestAllOf.md)
 - [MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType](docs/MpkiApiV1SeatSeatIdDeleteRequestAllOfSeatType.md)
 - [MpkiApiV1SeatTypesGet200ResponseInner](docs/MpkiApiV1SeatTypesGet200ResponseInner.md)
 - [NameDto](docs/NameDto.md)
 - [PaginationResult](docs/PaginationResult.md)
 - [PagingParameters](docs/PagingParameters.md)
 - [ProfileAuthenticationMethod](docs/ProfileAuthenticationMethod.md)
 - [ProfileEnrollmentMethod](docs/ProfileEnrollmentMethod.md)
 - [PublicCertificateDetails](docs/PublicCertificateDetails.md)
 - [PublicCertificateDetailsRevocation](docs/PublicCertificateDetailsRevocation.md)
 - [PublicEnrollmentDetails](docs/PublicEnrollmentDetails.md)
 - [PublicEnrollmentDetailsSeat](docs/PublicEnrollmentDetailsSeat.md)
 - [PublicProfileField](docs/PublicProfileField.md)
 - [PublicProfileFieldSource](docs/PublicProfileFieldSource.md)
 - [PublicProfileResponse](docs/PublicProfileResponse.md)
 - [PublicProfileResponseCertificate](docs/PublicProfileResponseCertificate.md)
 - [PublicProfileResponseCertificateExtensions](docs/PublicProfileResponseCertificateExtensions.md)
 - [PublicProfileResponseCertificateExtensionsSan](docs/PublicProfileResponseCertificateExtensionsSan.md)
 - [PublicProfileResponseCertificateIssuer](docs/PublicProfileResponseCertificateIssuer.md)
 - [PublicProfileResponseCertificateIssuerAllOf](docs/PublicProfileResponseCertificateIssuerAllOf.md)
 - [PublicProfileResponseCertificateSubject](docs/PublicProfileResponseCertificateSubject.md)
 - [PublicProfileResponsePrivateKeyAttributes](docs/PublicProfileResponsePrivateKeyAttributes.md)
 - [ReportSearchParams](docs/ReportSearchParams.md)
 - [SeatDetails](docs/SeatDetails.md)
 - [Validity](docs/Validity.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



