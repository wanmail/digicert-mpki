# PublicProfileFieldSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mandatory** | Pointer to **bool** | If true, the source is mandatory | [optional] 
**Source** | Pointer to **string** | Source of the attribute | [optional] 
**Value** | Pointer to **map[string]interface{}** | Permanent attribute value if source is &#x60;FIXED_VALUE&#x60; | [optional] 
**CustomValue** | Pointer to **map[string]interface{}** | Permanent custom Other name value if source is &#x60;FIXED_VALUE&#x60; | [optional] 
**SamlAssertionKey** | Pointer to **string** | SAML attribute mapping value | [optional] 
**AdAttributeKey** | Pointer to **string** | Active Directory attribute mapping value | [optional] 

## Methods

### NewPublicProfileFieldSource

`func NewPublicProfileFieldSource() *PublicProfileFieldSource`

NewPublicProfileFieldSource instantiates a new PublicProfileFieldSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProfileFieldSourceWithDefaults

`func NewPublicProfileFieldSourceWithDefaults() *PublicProfileFieldSource`

NewPublicProfileFieldSourceWithDefaults instantiates a new PublicProfileFieldSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMandatory

`func (o *PublicProfileFieldSource) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *PublicProfileFieldSource) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *PublicProfileFieldSource) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *PublicProfileFieldSource) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetSource

`func (o *PublicProfileFieldSource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublicProfileFieldSource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublicProfileFieldSource) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PublicProfileFieldSource) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *PublicProfileFieldSource) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PublicProfileFieldSource) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PublicProfileFieldSource) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PublicProfileFieldSource) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCustomValue

`func (o *PublicProfileFieldSource) GetCustomValue() map[string]interface{}`

GetCustomValue returns the CustomValue field if non-nil, zero value otherwise.

### GetCustomValueOk

`func (o *PublicProfileFieldSource) GetCustomValueOk() (*map[string]interface{}, bool)`

GetCustomValueOk returns a tuple with the CustomValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue

`func (o *PublicProfileFieldSource) SetCustomValue(v map[string]interface{})`

SetCustomValue sets CustomValue field to given value.

### HasCustomValue

`func (o *PublicProfileFieldSource) HasCustomValue() bool`

HasCustomValue returns a boolean if a field has been set.

### GetSamlAssertionKey

`func (o *PublicProfileFieldSource) GetSamlAssertionKey() string`

GetSamlAssertionKey returns the SamlAssertionKey field if non-nil, zero value otherwise.

### GetSamlAssertionKeyOk

`func (o *PublicProfileFieldSource) GetSamlAssertionKeyOk() (*string, bool)`

GetSamlAssertionKeyOk returns a tuple with the SamlAssertionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAssertionKey

`func (o *PublicProfileFieldSource) SetSamlAssertionKey(v string)`

SetSamlAssertionKey sets SamlAssertionKey field to given value.

### HasSamlAssertionKey

`func (o *PublicProfileFieldSource) HasSamlAssertionKey() bool`

HasSamlAssertionKey returns a boolean if a field has been set.

### GetAdAttributeKey

`func (o *PublicProfileFieldSource) GetAdAttributeKey() string`

GetAdAttributeKey returns the AdAttributeKey field if non-nil, zero value otherwise.

### GetAdAttributeKeyOk

`func (o *PublicProfileFieldSource) GetAdAttributeKeyOk() (*string, bool)`

GetAdAttributeKeyOk returns a tuple with the AdAttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAttributeKey

`func (o *PublicProfileFieldSource) SetAdAttributeKey(v string)`

SetAdAttributeKey sets AdAttributeKey field to given value.

### HasAdAttributeKey

`func (o *PublicProfileFieldSource) HasAdAttributeKey() bool`

HasAdAttributeKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


