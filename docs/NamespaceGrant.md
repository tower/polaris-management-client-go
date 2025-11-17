# NamespaceGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** |  | 
**Privilege** | [**NamespacePrivilege**](NamespacePrivilege.md) |  | 

## Methods

### NewNamespaceGrant

`func NewNamespaceGrant(namespace []string, privilege NamespacePrivilege, ) *NamespaceGrant`

NewNamespaceGrant instantiates a new NamespaceGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceGrantWithDefaults

`func NewNamespaceGrantWithDefaults() *NamespaceGrant`

NewNamespaceGrantWithDefaults instantiates a new NamespaceGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *NamespaceGrant) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NamespaceGrant) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NamespaceGrant) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetPrivilege

`func (o *NamespaceGrant) GetPrivilege() NamespacePrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *NamespaceGrant) GetPrivilegeOk() (*NamespacePrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *NamespaceGrant) SetPrivilege(v NamespacePrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


