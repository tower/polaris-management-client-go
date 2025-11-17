# HadoopConnectionConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouse** | Pointer to **string** | The file path to where this catalog should store tables | [optional] 

## Methods

### NewHadoopConnectionConfigInfo

`func NewHadoopConnectionConfigInfo() *HadoopConnectionConfigInfo`

NewHadoopConnectionConfigInfo instantiates a new HadoopConnectionConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHadoopConnectionConfigInfoWithDefaults

`func NewHadoopConnectionConfigInfoWithDefaults() *HadoopConnectionConfigInfo`

NewHadoopConnectionConfigInfoWithDefaults instantiates a new HadoopConnectionConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouse

`func (o *HadoopConnectionConfigInfo) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *HadoopConnectionConfigInfo) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *HadoopConnectionConfigInfo) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *HadoopConnectionConfigInfo) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


