# MpkiApiV1CertificateSearchGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Offset** | **int32** |  | 
**Limit** | **int32** |  | 
**Items** | [**[]PublicCertificateDetails**](PublicCertificateDetails.md) |  | 

## Methods

### NewMpkiApiV1CertificateSearchGet200Response

`func NewMpkiApiV1CertificateSearchGet200Response(total int32, offset int32, limit int32, items []PublicCertificateDetails, ) *MpkiApiV1CertificateSearchGet200Response`

NewMpkiApiV1CertificateSearchGet200Response instantiates a new MpkiApiV1CertificateSearchGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpkiApiV1CertificateSearchGet200ResponseWithDefaults

`func NewMpkiApiV1CertificateSearchGet200ResponseWithDefaults() *MpkiApiV1CertificateSearchGet200Response`

NewMpkiApiV1CertificateSearchGet200ResponseWithDefaults instantiates a new MpkiApiV1CertificateSearchGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MpkiApiV1CertificateSearchGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MpkiApiV1CertificateSearchGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MpkiApiV1CertificateSearchGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *MpkiApiV1CertificateSearchGet200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MpkiApiV1CertificateSearchGet200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MpkiApiV1CertificateSearchGet200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *MpkiApiV1CertificateSearchGet200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MpkiApiV1CertificateSearchGet200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MpkiApiV1CertificateSearchGet200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetItems

`func (o *MpkiApiV1CertificateSearchGet200Response) GetItems() []PublicCertificateDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MpkiApiV1CertificateSearchGet200Response) GetItemsOk() (*[]PublicCertificateDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MpkiApiV1CertificateSearchGet200Response) SetItems(v []PublicCertificateDetails)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


