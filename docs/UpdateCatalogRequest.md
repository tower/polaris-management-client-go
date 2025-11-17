# UpdateCatalogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentEntityVersion** | Pointer to **int32** | The version of the object onto which this update is applied; if the object changed, the update will fail and the caller should retry after fetching the latest version. | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**StorageConfigInfo** | Pointer to [**StorageConfigInfo**](StorageConfigInfo.md) |  | [optional] 

## Methods

### NewUpdateCatalogRequest

`func NewUpdateCatalogRequest() *UpdateCatalogRequest`

NewUpdateCatalogRequest instantiates a new UpdateCatalogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogRequestWithDefaults

`func NewUpdateCatalogRequestWithDefaults() *UpdateCatalogRequest`

NewUpdateCatalogRequestWithDefaults instantiates a new UpdateCatalogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentEntityVersion

`func (o *UpdateCatalogRequest) GetCurrentEntityVersion() int32`

GetCurrentEntityVersion returns the CurrentEntityVersion field if non-nil, zero value otherwise.

### GetCurrentEntityVersionOk

`func (o *UpdateCatalogRequest) GetCurrentEntityVersionOk() (*int32, bool)`

GetCurrentEntityVersionOk returns a tuple with the CurrentEntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEntityVersion

`func (o *UpdateCatalogRequest) SetCurrentEntityVersion(v int32)`

SetCurrentEntityVersion sets CurrentEntityVersion field to given value.

### HasCurrentEntityVersion

`func (o *UpdateCatalogRequest) HasCurrentEntityVersion() bool`

HasCurrentEntityVersion returns a boolean if a field has been set.

### GetProperties

`func (o *UpdateCatalogRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdateCatalogRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdateCatalogRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UpdateCatalogRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStorageConfigInfo

`func (o *UpdateCatalogRequest) GetStorageConfigInfo() StorageConfigInfo`

GetStorageConfigInfo returns the StorageConfigInfo field if non-nil, zero value otherwise.

### GetStorageConfigInfoOk

`func (o *UpdateCatalogRequest) GetStorageConfigInfoOk() (*StorageConfigInfo, bool)`

GetStorageConfigInfoOk returns a tuple with the StorageConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigInfo

`func (o *UpdateCatalogRequest) SetStorageConfigInfo(v StorageConfigInfo)`

SetStorageConfigInfo sets StorageConfigInfo field to given value.

### HasStorageConfigInfo

`func (o *UpdateCatalogRequest) HasStorageConfigInfo() bool`

HasStorageConfigInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


