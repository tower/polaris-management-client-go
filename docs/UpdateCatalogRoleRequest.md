# UpdateCatalogRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentEntityVersion** | **int32** | The version of the object onto which this update is applied; if the object changed, the update will fail and the caller should retry after fetching the latest version. | 
**Properties** | **map[string]string** |  | 

## Methods

### NewUpdateCatalogRoleRequest

`func NewUpdateCatalogRoleRequest(currentEntityVersion int32, properties map[string]string, ) *UpdateCatalogRoleRequest`

NewUpdateCatalogRoleRequest instantiates a new UpdateCatalogRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogRoleRequestWithDefaults

`func NewUpdateCatalogRoleRequestWithDefaults() *UpdateCatalogRoleRequest`

NewUpdateCatalogRoleRequestWithDefaults instantiates a new UpdateCatalogRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentEntityVersion

`func (o *UpdateCatalogRoleRequest) GetCurrentEntityVersion() int32`

GetCurrentEntityVersion returns the CurrentEntityVersion field if non-nil, zero value otherwise.

### GetCurrentEntityVersionOk

`func (o *UpdateCatalogRoleRequest) GetCurrentEntityVersionOk() (*int32, bool)`

GetCurrentEntityVersionOk returns a tuple with the CurrentEntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEntityVersion

`func (o *UpdateCatalogRoleRequest) SetCurrentEntityVersion(v int32)`

SetCurrentEntityVersion sets CurrentEntityVersion field to given value.


### GetProperties

`func (o *UpdateCatalogRoleRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdateCatalogRoleRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdateCatalogRoleRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


