# TableGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** |  | 
**TableName** | **string** |  | 
**Privilege** | [**TablePrivilege**](TablePrivilege.md) |  | 

## Methods

### NewTableGrant

`func NewTableGrant(namespace []string, tableName string, privilege TablePrivilege, ) *TableGrant`

NewTableGrant instantiates a new TableGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableGrantWithDefaults

`func NewTableGrantWithDefaults() *TableGrant`

NewTableGrantWithDefaults instantiates a new TableGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *TableGrant) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TableGrant) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TableGrant) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetTableName

`func (o *TableGrant) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableGrant) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableGrant) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetPrivilege

`func (o *TableGrant) GetPrivilege() TablePrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *TableGrant) GetPrivilegeOk() (*TablePrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *TableGrant) SetPrivilege(v TablePrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


