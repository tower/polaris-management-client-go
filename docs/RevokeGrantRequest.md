# RevokeGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grant** | Pointer to [**GrantResource**](GrantResource.md) |  | [optional] 

## Methods

### NewRevokeGrantRequest

`func NewRevokeGrantRequest() *RevokeGrantRequest`

NewRevokeGrantRequest instantiates a new RevokeGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeGrantRequestWithDefaults

`func NewRevokeGrantRequestWithDefaults() *RevokeGrantRequest`

NewRevokeGrantRequestWithDefaults instantiates a new RevokeGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrant

`func (o *RevokeGrantRequest) GetGrant() GrantResource`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *RevokeGrantRequest) GetGrantOk() (*GrantResource, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *RevokeGrantRequest) SetGrant(v GrantResource)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *RevokeGrantRequest) HasGrant() bool`

HasGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


