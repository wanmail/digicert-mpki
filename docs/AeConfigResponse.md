# AeConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of AE Config, used for filename | [optional] 
**ProfileIds** | Pointer to **[]string** | Array of Profile IDs included in AE Config | [optional] 
**GenerationDateTime** | Pointer to **string** | Generation date/time of the AE Config | [optional] 
**AeConfig** | Pointer to **string** | AE Config with Windows style line endings (\\r\\n) | [optional] 

## Methods

### NewAeConfigResponse

`func NewAeConfigResponse() *AeConfigResponse`

NewAeConfigResponse instantiates a new AeConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeConfigResponseWithDefaults

`func NewAeConfigResponseWithDefaults() *AeConfigResponse`

NewAeConfigResponseWithDefaults instantiates a new AeConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AeConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AeConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AeConfigResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AeConfigResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileIds

`func (o *AeConfigResponse) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *AeConfigResponse) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *AeConfigResponse) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *AeConfigResponse) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetGenerationDateTime

`func (o *AeConfigResponse) GetGenerationDateTime() string`

GetGenerationDateTime returns the GenerationDateTime field if non-nil, zero value otherwise.

### GetGenerationDateTimeOk

`func (o *AeConfigResponse) GetGenerationDateTimeOk() (*string, bool)`

GetGenerationDateTimeOk returns a tuple with the GenerationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationDateTime

`func (o *AeConfigResponse) SetGenerationDateTime(v string)`

SetGenerationDateTime sets GenerationDateTime field to given value.

### HasGenerationDateTime

`func (o *AeConfigResponse) HasGenerationDateTime() bool`

HasGenerationDateTime returns a boolean if a field has been set.

### GetAeConfig

`func (o *AeConfigResponse) GetAeConfig() string`

GetAeConfig returns the AeConfig field if non-nil, zero value otherwise.

### GetAeConfigOk

`func (o *AeConfigResponse) GetAeConfigOk() (*string, bool)`

GetAeConfigOk returns a tuple with the AeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeConfig

`func (o *AeConfigResponse) SetAeConfig(v string)`

SetAeConfig sets AeConfig field to given value.

### HasAeConfig

`func (o *AeConfigResponse) HasAeConfig() bool`

HasAeConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


