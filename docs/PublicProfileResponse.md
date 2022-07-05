# PublicProfileResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** | Profile name | [optional] [default to null]
**Status** | **string** | Profile status | [optional] [default to null]
**SignatureAlgorithm** | **string** | Signature algorithm for certificates issued from the profile | [optional] [default to null]
**PublishToPublicDirectory** | **bool** | If true, certificate public key is published to the DigiCert PKI directory | [optional] [default to null]
**RenewalPeriodDays** | **int32** | Number of days before expiration in which the certificate can be renewed | [optional] [default to null]
**DuplicateCertPolicy** | **bool** | If true, certificate profile allows duplicate certificates | [optional] [default to null]
**OverrideCertValidityViaApi** | **bool** | If true, the default validity period set by the profile can be overridden using a request parameter | [optional] [default to null]
**CertificateDeliveryFormat** | **string** | Format for certificate delivery specified in the profile | [optional] [default to null]
**BusinessUnitId** | **string** | Business unit ID to which the profile belongs | [optional] [default to null]
**EnrollmentMethod** | [***ProfileEnrollmentMethod**](ProfileEnrollmentMethod.md) |  | [optional] [default to null]
**AuthenticationMethod** | [***ProfileAuthenticationMethod**](ProfileAuthenticationMethod.md) |  | [optional] [default to null]
**KeyEscrowPolicy** | [***KeyEscrowPolicy**](KeyEscrowPolicy.md) |  | [optional] [default to null]
**PrivateKeyAttributes** | [***PublicProfileResponsePrivateKeyAttributes**](PublicProfileResponse_private_key_attributes.md) |  | [optional] [default to null]
**ApiTokenBindingEnabled** | **bool** | If true, use of the certificate profile is restricted to a specific API token | [optional] [default to null]
**ApiTokenId** | **string** | API token ID to which the profile is bound if api_token_binding_enabled is true | [optional] [default to null]
**Certificate** | [***PublicProfileResponseCertificate**](PublicProfileResponse_certificate.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

