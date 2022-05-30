# PagingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewPagingParameters

`func NewPagingParameters() *PagingParameters`

NewPagingParameters instantiates a new PagingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingParametersWithDefaults

`func NewPagingParametersWithDefaults() *PagingParameters`

NewPagingParametersWithDefaults instantiates a new PagingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *PagingParameters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagingParameters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagingParameters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PagingParameters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PagingParameters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagingParameters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagingParameters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagingParameters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


