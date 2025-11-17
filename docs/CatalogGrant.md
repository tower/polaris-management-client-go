# CatalogGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privilege** | [**CatalogPrivilege**](CatalogPrivilege.md) |  | 

## Methods

### NewCatalogGrant

`func NewCatalogGrant(privilege CatalogPrivilege, ) *CatalogGrant`

NewCatalogGrant instantiates a new CatalogGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogGrantWithDefaults

`func NewCatalogGrantWithDefaults() *CatalogGrant`

NewCatalogGrantWithDefaults instantiates a new CatalogGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilege

`func (o *CatalogGrant) GetPrivilege() CatalogPrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *CatalogGrant) GetPrivilegeOk() (*CatalogPrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *CatalogGrant) SetPrivilege(v CatalogPrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


