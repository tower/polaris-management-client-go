# StorageConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | **string** | The cloud provider type this storage is built on. FILE is supported for testing purposes only | 
**AllowedLocations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStorageConfigInfo

`func NewStorageConfigInfo(storageType string, ) *StorageConfigInfo`

NewStorageConfigInfo instantiates a new StorageConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigInfoWithDefaults

`func NewStorageConfigInfoWithDefaults() *StorageConfigInfo`

NewStorageConfigInfoWithDefaults instantiates a new StorageConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *StorageConfigInfo) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *StorageConfigInfo) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *StorageConfigInfo) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetAllowedLocations

`func (o *StorageConfigInfo) GetAllowedLocations() []string`

GetAllowedLocations returns the AllowedLocations field if non-nil, zero value otherwise.

### GetAllowedLocationsOk

`func (o *StorageConfigInfo) GetAllowedLocationsOk() (*[]string, bool)`

GetAllowedLocationsOk returns a tuple with the AllowedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLocations

`func (o *StorageConfigInfo) SetAllowedLocations(v []string)`

SetAllowedLocations sets AllowedLocations field to given value.

### HasAllowedLocations

`func (o *StorageConfigInfo) HasAllowedLocations() bool`

HasAllowedLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


