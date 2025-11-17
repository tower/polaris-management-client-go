# Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the type of catalog - internal or external | [default to "INTERNAL"]
**Name** | **string** | The name of the catalog | 
**Properties** | [**CatalogProperties**](CatalogProperties.md) |  | 
**CreateTimestamp** | Pointer to **int64** | The creation time represented as unix epoch timestamp in milliseconds | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** | The last update time represented as unix epoch timestamp in milliseconds | [optional] 
**EntityVersion** | Pointer to **int32** | The version of the catalog object used to determine if the catalog metadata has changed | [optional] 
**StorageConfigInfo** | [**StorageConfigInfo**](StorageConfigInfo.md) |  | 

## Methods

### NewCatalog

`func NewCatalog(type_ string, name string, properties CatalogProperties, storageConfigInfo StorageConfigInfo, ) *Catalog`

NewCatalog instantiates a new Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWithDefaults

`func NewCatalogWithDefaults() *Catalog`

NewCatalogWithDefaults instantiates a new Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Catalog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Catalog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Catalog) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Catalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Catalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Catalog) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *Catalog) GetProperties() CatalogProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Catalog) GetPropertiesOk() (*CatalogProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Catalog) SetProperties(v CatalogProperties)`

SetProperties sets Properties field to given value.


### GetCreateTimestamp

`func (o *Catalog) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *Catalog) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *Catalog) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *Catalog) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Catalog) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Catalog) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Catalog) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Catalog) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetEntityVersion

`func (o *Catalog) GetEntityVersion() int32`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *Catalog) GetEntityVersionOk() (*int32, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *Catalog) SetEntityVersion(v int32)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *Catalog) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.

### GetStorageConfigInfo

`func (o *Catalog) GetStorageConfigInfo() StorageConfigInfo`

GetStorageConfigInfo returns the StorageConfigInfo field if non-nil, zero value otherwise.

### GetStorageConfigInfoOk

`func (o *Catalog) GetStorageConfigInfoOk() (*StorageConfigInfo, bool)`

GetStorageConfigInfoOk returns a tuple with the StorageConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigInfo

`func (o *Catalog) SetStorageConfigInfo(v StorageConfigInfo)`

SetStorageConfigInfo sets StorageConfigInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


