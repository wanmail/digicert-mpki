# MicrosoftAutoEnrollmentClientSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyExport** | Pointer to [**MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport**](MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport.md) |  | [optional] 
**EnrollmentMode** | Pointer to [**MicrosoftAutoEnrollmentClientSettingsEnrollmentMode**](MicrosoftAutoEnrollmentClientSettingsEnrollmentMode.md) |  | [optional] 
**KeySpec** | Pointer to **string** | Describes KeySpec | [optional] 
**TemplateSchemaVersion** | Pointer to **int32** | Certificate template schema version | [optional] 
**WindowsHelloLogon** | Pointer to **bool** | If true, certificate is used for Windows Hello for Business authentication | [optional] 
**EnrollOnBehalfOf** | Pointer to **bool** | If true, certificate is enrolled on behalf of user | [optional] 
**PublishToAd** | Pointer to **bool** | If true, certificate is published to Active Directory | [optional] 
**EndEntityType** | Pointer to **string** | Describes the type of End Entity Type | [optional] 
**CryptoProviders** | Pointer to [**MicrosoftAutoEnrollmentClientSettingsCryptoProviders**](MicrosoftAutoEnrollmentClientSettingsCryptoProviders.md) |  | [optional] 

## Methods

### NewMicrosoftAutoEnrollmentClientSettings

`func NewMicrosoftAutoEnrollmentClientSettings() *MicrosoftAutoEnrollmentClientSettings`

NewMicrosoftAutoEnrollmentClientSettings instantiates a new MicrosoftAutoEnrollmentClientSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftAutoEnrollmentClientSettingsWithDefaults

`func NewMicrosoftAutoEnrollmentClientSettingsWithDefaults() *MicrosoftAutoEnrollmentClientSettings`

NewMicrosoftAutoEnrollmentClientSettingsWithDefaults instantiates a new MicrosoftAutoEnrollmentClientSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyExport

`func (o *MicrosoftAutoEnrollmentClientSettings) GetPrivateKeyExport() MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport`

GetPrivateKeyExport returns the PrivateKeyExport field if non-nil, zero value otherwise.

### GetPrivateKeyExportOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetPrivateKeyExportOk() (*MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport, bool)`

GetPrivateKeyExportOk returns a tuple with the PrivateKeyExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyExport

`func (o *MicrosoftAutoEnrollmentClientSettings) SetPrivateKeyExport(v MicrosoftAutoEnrollmentClientSettingsPrivateKeyExport)`

SetPrivateKeyExport sets PrivateKeyExport field to given value.

### HasPrivateKeyExport

`func (o *MicrosoftAutoEnrollmentClientSettings) HasPrivateKeyExport() bool`

HasPrivateKeyExport returns a boolean if a field has been set.

### GetEnrollmentMode

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollmentMode() MicrosoftAutoEnrollmentClientSettingsEnrollmentMode`

GetEnrollmentMode returns the EnrollmentMode field if non-nil, zero value otherwise.

### GetEnrollmentModeOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollmentModeOk() (*MicrosoftAutoEnrollmentClientSettingsEnrollmentMode, bool)`

GetEnrollmentModeOk returns a tuple with the EnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMode

`func (o *MicrosoftAutoEnrollmentClientSettings) SetEnrollmentMode(v MicrosoftAutoEnrollmentClientSettingsEnrollmentMode)`

SetEnrollmentMode sets EnrollmentMode field to given value.

### HasEnrollmentMode

`func (o *MicrosoftAutoEnrollmentClientSettings) HasEnrollmentMode() bool`

HasEnrollmentMode returns a boolean if a field has been set.

### GetKeySpec

`func (o *MicrosoftAutoEnrollmentClientSettings) GetKeySpec() string`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetKeySpecOk() (*string, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *MicrosoftAutoEnrollmentClientSettings) SetKeySpec(v string)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *MicrosoftAutoEnrollmentClientSettings) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.

### GetTemplateSchemaVersion

`func (o *MicrosoftAutoEnrollmentClientSettings) GetTemplateSchemaVersion() int32`

GetTemplateSchemaVersion returns the TemplateSchemaVersion field if non-nil, zero value otherwise.

### GetTemplateSchemaVersionOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetTemplateSchemaVersionOk() (*int32, bool)`

GetTemplateSchemaVersionOk returns a tuple with the TemplateSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSchemaVersion

`func (o *MicrosoftAutoEnrollmentClientSettings) SetTemplateSchemaVersion(v int32)`

SetTemplateSchemaVersion sets TemplateSchemaVersion field to given value.

### HasTemplateSchemaVersion

`func (o *MicrosoftAutoEnrollmentClientSettings) HasTemplateSchemaVersion() bool`

HasTemplateSchemaVersion returns a boolean if a field has been set.

### GetWindowsHelloLogon

`func (o *MicrosoftAutoEnrollmentClientSettings) GetWindowsHelloLogon() bool`

GetWindowsHelloLogon returns the WindowsHelloLogon field if non-nil, zero value otherwise.

### GetWindowsHelloLogonOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetWindowsHelloLogonOk() (*bool, bool)`

GetWindowsHelloLogonOk returns a tuple with the WindowsHelloLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsHelloLogon

`func (o *MicrosoftAutoEnrollmentClientSettings) SetWindowsHelloLogon(v bool)`

SetWindowsHelloLogon sets WindowsHelloLogon field to given value.

### HasWindowsHelloLogon

`func (o *MicrosoftAutoEnrollmentClientSettings) HasWindowsHelloLogon() bool`

HasWindowsHelloLogon returns a boolean if a field has been set.

### GetEnrollOnBehalfOf

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollOnBehalfOf() bool`

GetEnrollOnBehalfOf returns the EnrollOnBehalfOf field if non-nil, zero value otherwise.

### GetEnrollOnBehalfOfOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEnrollOnBehalfOfOk() (*bool, bool)`

GetEnrollOnBehalfOfOk returns a tuple with the EnrollOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollOnBehalfOf

`func (o *MicrosoftAutoEnrollmentClientSettings) SetEnrollOnBehalfOf(v bool)`

SetEnrollOnBehalfOf sets EnrollOnBehalfOf field to given value.

### HasEnrollOnBehalfOf

`func (o *MicrosoftAutoEnrollmentClientSettings) HasEnrollOnBehalfOf() bool`

HasEnrollOnBehalfOf returns a boolean if a field has been set.

### GetPublishToAd

`func (o *MicrosoftAutoEnrollmentClientSettings) GetPublishToAd() bool`

GetPublishToAd returns the PublishToAd field if non-nil, zero value otherwise.

### GetPublishToAdOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetPublishToAdOk() (*bool, bool)`

GetPublishToAdOk returns a tuple with the PublishToAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToAd

`func (o *MicrosoftAutoEnrollmentClientSettings) SetPublishToAd(v bool)`

SetPublishToAd sets PublishToAd field to given value.

### HasPublishToAd

`func (o *MicrosoftAutoEnrollmentClientSettings) HasPublishToAd() bool`

HasPublishToAd returns a boolean if a field has been set.

### GetEndEntityType

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEndEntityType() string`

GetEndEntityType returns the EndEntityType field if non-nil, zero value otherwise.

### GetEndEntityTypeOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetEndEntityTypeOk() (*string, bool)`

GetEndEntityTypeOk returns a tuple with the EndEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityType

`func (o *MicrosoftAutoEnrollmentClientSettings) SetEndEntityType(v string)`

SetEndEntityType sets EndEntityType field to given value.

### HasEndEntityType

`func (o *MicrosoftAutoEnrollmentClientSettings) HasEndEntityType() bool`

HasEndEntityType returns a boolean if a field has been set.

### GetCryptoProviders

`func (o *MicrosoftAutoEnrollmentClientSettings) GetCryptoProviders() MicrosoftAutoEnrollmentClientSettingsCryptoProviders`

GetCryptoProviders returns the CryptoProviders field if non-nil, zero value otherwise.

### GetCryptoProvidersOk

`func (o *MicrosoftAutoEnrollmentClientSettings) GetCryptoProvidersOk() (*MicrosoftAutoEnrollmentClientSettingsCryptoProviders, bool)`

GetCryptoProvidersOk returns a tuple with the CryptoProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProviders

`func (o *MicrosoftAutoEnrollmentClientSettings) SetCryptoProviders(v MicrosoftAutoEnrollmentClientSettingsCryptoProviders)`

SetCryptoProviders sets CryptoProviders field to given value.

### HasCryptoProviders

`func (o *MicrosoftAutoEnrollmentClientSettings) HasCryptoProviders() bool`

HasCryptoProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


