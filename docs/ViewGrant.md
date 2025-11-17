# ViewGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** |  | 
**ViewName** | **string** |  | 
**Privilege** | [**ViewPrivilege**](ViewPrivilege.md) |  | 

## Methods

### NewViewGrant

`func NewViewGrant(namespace []string, viewName string, privilege ViewPrivilege, ) *ViewGrant`

NewViewGrant instantiates a new ViewGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewGrantWithDefaults

`func NewViewGrantWithDefaults() *ViewGrant`

NewViewGrantWithDefaults instantiates a new ViewGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ViewGrant) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ViewGrant) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ViewGrant) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetViewName

`func (o *ViewGrant) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewGrant) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewGrant) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### GetPrivilege

`func (o *ViewGrant) GetPrivilege() ViewPrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ViewGrant) GetPrivilegeOk() (*ViewPrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ViewGrant) SetPrivilege(v ViewPrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


