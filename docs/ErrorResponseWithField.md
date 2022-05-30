# ErrorResponseWithField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Field** | **string** |  | 

## Methods

### NewErrorResponseWithField

`func NewErrorResponseWithField(message string, field string, ) *ErrorResponseWithField`

NewErrorResponseWithField instantiates a new ErrorResponseWithField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithFieldWithDefaults

`func NewErrorResponseWithFieldWithDefaults() *ErrorResponseWithField`

NewErrorResponseWithFieldWithDefaults instantiates a new ErrorResponseWithField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseWithField) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseWithField) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseWithField) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponseWithField) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseWithField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseWithField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseWithField) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetField

`func (o *ErrorResponseWithField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ErrorResponseWithField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ErrorResponseWithField) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


