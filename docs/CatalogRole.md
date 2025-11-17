# CatalogRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**CreateTimestamp** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**EntityVersion** | Pointer to **int32** | The version of the catalog role object used to determine if the catalog role metadata has changed | [optional] 

## Methods

### NewCatalogRole

`func NewCatalogRole(name string, ) *CatalogRole`

NewCatalogRole instantiates a new CatalogRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogRoleWithDefaults

`func NewCatalogRoleWithDefaults() *CatalogRole`

NewCatalogRoleWithDefaults instantiates a new CatalogRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogRole) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *CatalogRole) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CatalogRole) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CatalogRole) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CatalogRole) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *CatalogRole) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *CatalogRole) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *CatalogRole) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *CatalogRole) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *CatalogRole) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *CatalogRole) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *CatalogRole) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *CatalogRole) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetEntityVersion

`func (o *CatalogRole) GetEntityVersion() int32`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *CatalogRole) GetEntityVersionOk() (*int32, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *CatalogRole) SetEntityVersion(v int32)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *CatalogRole) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


