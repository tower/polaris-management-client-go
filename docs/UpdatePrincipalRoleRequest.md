# UpdatePrincipalRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentEntityVersion** | **int32** | The version of the object onto which this update is applied; if the object changed, the update will fail and the caller should retry after fetching the latest version. | 
**Properties** | **map[string]string** |  | 

## Methods

### NewUpdatePrincipalRoleRequest

`func NewUpdatePrincipalRoleRequest(currentEntityVersion int32, properties map[string]string, ) *UpdatePrincipalRoleRequest`

NewUpdatePrincipalRoleRequest instantiates a new UpdatePrincipalRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrincipalRoleRequestWithDefaults

`func NewUpdatePrincipalRoleRequestWithDefaults() *UpdatePrincipalRoleRequest`

NewUpdatePrincipalRoleRequestWithDefaults instantiates a new UpdatePrincipalRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentEntityVersion

`func (o *UpdatePrincipalRoleRequest) GetCurrentEntityVersion() int32`

GetCurrentEntityVersion returns the CurrentEntityVersion field if non-nil, zero value otherwise.

### GetCurrentEntityVersionOk

`func (o *UpdatePrincipalRoleRequest) GetCurrentEntityVersionOk() (*int32, bool)`

GetCurrentEntityVersionOk returns a tuple with the CurrentEntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEntityVersion

`func (o *UpdatePrincipalRoleRequest) SetCurrentEntityVersion(v int32)`

SetCurrentEntityVersion sets CurrentEntityVersion field to given value.


### GetProperties

`func (o *UpdatePrincipalRoleRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdatePrincipalRoleRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdatePrincipalRoleRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


