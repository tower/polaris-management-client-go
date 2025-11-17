# PrincipalRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role | 
**Federated** | Pointer to **bool** | Whether the principal role is a federated role (that is, managed by an external identity provider) | [optional] [default to false]
**Properties** | Pointer to **map[string]string** |  | [optional] 
**CreateTimestamp** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**EntityVersion** | Pointer to **int32** | The version of the principal role object used to determine if the principal role metadata has changed | [optional] 

## Methods

### NewPrincipalRole

`func NewPrincipalRole(name string, ) *PrincipalRole`

NewPrincipalRole instantiates a new PrincipalRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalRoleWithDefaults

`func NewPrincipalRoleWithDefaults() *PrincipalRole`

NewPrincipalRoleWithDefaults instantiates a new PrincipalRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrincipalRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalRole) SetName(v string)`

SetName sets Name field to given value.


### GetFederated

`func (o *PrincipalRole) GetFederated() bool`

GetFederated returns the Federated field if non-nil, zero value otherwise.

### GetFederatedOk

`func (o *PrincipalRole) GetFederatedOk() (*bool, bool)`

GetFederatedOk returns a tuple with the Federated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederated

`func (o *PrincipalRole) SetFederated(v bool)`

SetFederated sets Federated field to given value.

### HasFederated

`func (o *PrincipalRole) HasFederated() bool`

HasFederated returns a boolean if a field has been set.

### GetProperties

`func (o *PrincipalRole) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PrincipalRole) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PrincipalRole) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PrincipalRole) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *PrincipalRole) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *PrincipalRole) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *PrincipalRole) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *PrincipalRole) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *PrincipalRole) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *PrincipalRole) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *PrincipalRole) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *PrincipalRole) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetEntityVersion

`func (o *PrincipalRole) GetEntityVersion() int32`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *PrincipalRole) GetEntityVersionOk() (*int32, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *PrincipalRole) SetEntityVersion(v int32)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *PrincipalRole) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


