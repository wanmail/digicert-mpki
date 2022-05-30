# LogSearchParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **string** |  | [optional] 
**AdminId** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLogSearchParameters

`func NewLogSearchParameters() *LogSearchParameters`

NewLogSearchParameters instantiates a new LogSearchParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchParametersWithDefaults

`func NewLogSearchParametersWithDefaults() *LogSearchParameters`

NewLogSearchParametersWithDefaults instantiates a new LogSearchParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LogSearchParameters) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LogSearchParameters) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LogSearchParameters) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LogSearchParameters) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *LogSearchParameters) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *LogSearchParameters) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *LogSearchParameters) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *LogSearchParameters) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetAdminId

`func (o *LogSearchParameters) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *LogSearchParameters) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *LogSearchParameters) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *LogSearchParameters) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetResourceId

`func (o *LogSearchParameters) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *LogSearchParameters) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *LogSearchParameters) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *LogSearchParameters) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *LogSearchParameters) GetResourceType() []string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *LogSearchParameters) GetResourceTypeOk() (*[]string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *LogSearchParameters) SetResourceType(v []string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *LogSearchParameters) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *LogSearchParameters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogSearchParameters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogSearchParameters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogSearchParameters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAction

`func (o *LogSearchParameters) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LogSearchParameters) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LogSearchParameters) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LogSearchParameters) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFrom

`func (o *LogSearchParameters) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogSearchParameters) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogSearchParameters) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LogSearchParameters) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *LogSearchParameters) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogSearchParameters) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogSearchParameters) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *LogSearchParameters) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


