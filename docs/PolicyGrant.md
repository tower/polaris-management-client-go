# PolicyGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** |  | 
**PolicyName** | **string** |  | 
**Privilege** | [**PolicyPrivilege**](PolicyPrivilege.md) |  | 

## Methods

### NewPolicyGrant

`func NewPolicyGrant(namespace []string, policyName string, privilege PolicyPrivilege, ) *PolicyGrant`

NewPolicyGrant instantiates a new PolicyGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGrantWithDefaults

`func NewPolicyGrantWithDefaults() *PolicyGrant`

NewPolicyGrantWithDefaults instantiates a new PolicyGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *PolicyGrant) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PolicyGrant) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PolicyGrant) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetPolicyName

`func (o *PolicyGrant) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyGrant) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyGrant) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetPrivilege

`func (o *PolicyGrant) GetPrivilege() PolicyPrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *PolicyGrant) GetPrivilegeOk() (*PolicyPrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *PolicyGrant) SetPrivilege(v PolicyPrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


