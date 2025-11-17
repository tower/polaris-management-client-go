# HiveConnectionConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouse** | Pointer to **string** | The warehouse location for the hive catalog. | [optional] 

## Methods

### NewHiveConnectionConfigInfo

`func NewHiveConnectionConfigInfo() *HiveConnectionConfigInfo`

NewHiveConnectionConfigInfo instantiates a new HiveConnectionConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveConnectionConfigInfoWithDefaults

`func NewHiveConnectionConfigInfoWithDefaults() *HiveConnectionConfigInfo`

NewHiveConnectionConfigInfoWithDefaults instantiates a new HiveConnectionConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouse

`func (o *HiveConnectionConfigInfo) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *HiveConnectionConfigInfo) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *HiveConnectionConfigInfo) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *HiveConnectionConfigInfo) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


