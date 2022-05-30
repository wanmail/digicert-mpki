# Validity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Validity unit | [optional] 
**Duration** | Pointer to **int32** | Number of validity unit | [optional] 

## Methods

### NewValidity

`func NewValidity() *Validity`

NewValidity instantiates a new Validity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidityWithDefaults

`func NewValidityWithDefaults() *Validity`

NewValidityWithDefaults instantiates a new Validity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Validity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Validity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Validity) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Validity) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDuration

`func (o *Validity) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Validity) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Validity) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Validity) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


