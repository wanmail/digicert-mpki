# AuditLogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AuditLogItemAccount**](AuditLogItemAccount.md) |  | [optional] 
**BusinessUnit** | Pointer to [**AuditLogItemAccount**](AuditLogItemAccount.md) |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId**](MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditLogItem

`func NewAuditLogItem() *AuditLogItem`

NewAuditLogItem instantiates a new AuditLogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogItemWithDefaults

`func NewAuditLogItemWithDefaults() *AuditLogItem`

NewAuditLogItemWithDefaults instantiates a new AuditLogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLogItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AuditLogItem) GetAccount() AuditLogItemAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AuditLogItem) GetAccountOk() (*AuditLogItemAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AuditLogItem) SetAccount(v AuditLogItemAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AuditLogItem) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *AuditLogItem) GetBusinessUnit() AuditLogItemAccount`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *AuditLogItem) GetBusinessUnitOk() (*AuditLogItemAccount, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *AuditLogItem) SetBusinessUnit(v AuditLogItemAccount)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *AuditLogItem) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetResourceType

`func (o *AuditLogItem) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AuditLogItem) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AuditLogItem) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AuditLogItem) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *AuditLogItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuditLogItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuditLogItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuditLogItem) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetAction

`func (o *AuditLogItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLogItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLogItem) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLogItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUser

`func (o *AuditLogItem) GetUser() MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditLogItem) GetUserOk() (*MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditLogItem) SetUser(v MpkiApiV1CertificateImportPost200ResponseInnerBusinessUnitId)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditLogItem) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuditLogItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLogItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *AuditLogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditLogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditLogItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditLogItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *AuditLogItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditLogItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditLogItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditLogItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *AuditLogItem) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditLogItem) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditLogItem) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *AuditLogItem) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


