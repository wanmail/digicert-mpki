# PublicProfileField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Attribute type | [optional] 
**Mandatory** | Pointer to **bool** | If true, the attribute is mandatory | [optional] 
**Multiple** | Pointer to **bool** | If true, the attribute can have multiple values | [optional] 
**Sources** | Pointer to [**PublicProfileFieldSource**](PublicProfileFieldSource.md) |  | [optional] 

## Methods

### NewPublicProfileField

`func NewPublicProfileField() *PublicProfileField`

NewPublicProfileField instantiates a new PublicProfileField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProfileFieldWithDefaults

`func NewPublicProfileFieldWithDefaults() *PublicProfileField`

NewPublicProfileFieldWithDefaults instantiates a new PublicProfileField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicProfileField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicProfileField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicProfileField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicProfileField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMandatory

`func (o *PublicProfileField) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *PublicProfileField) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *PublicProfileField) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *PublicProfileField) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetMultiple

`func (o *PublicProfileField) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *PublicProfileField) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *PublicProfileField) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *PublicProfileField) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetSources

`func (o *PublicProfileField) GetSources() PublicProfileFieldSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PublicProfileField) GetSourcesOk() (*PublicProfileFieldSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PublicProfileField) SetSources(v PublicProfileFieldSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PublicProfileField) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


