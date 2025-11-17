# ConnectionConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The type of remote catalog service represented by this connection | 
**Uri** | Pointer to **string** | URI to the remote catalog service | [optional] 
**AuthenticationParameters** | Pointer to [**AuthenticationParameters**](AuthenticationParameters.md) |  | [optional] 
**ServiceIdentity** | Pointer to [**ServiceIdentityInfo**](ServiceIdentityInfo.md) |  | [optional] 

## Methods

### NewConnectionConfigInfo

`func NewConnectionConfigInfo(connectionType string, ) *ConnectionConfigInfo`

NewConnectionConfigInfo instantiates a new ConnectionConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionConfigInfoWithDefaults

`func NewConnectionConfigInfoWithDefaults() *ConnectionConfigInfo`

NewConnectionConfigInfoWithDefaults instantiates a new ConnectionConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ConnectionConfigInfo) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ConnectionConfigInfo) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ConnectionConfigInfo) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetUri

`func (o *ConnectionConfigInfo) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConnectionConfigInfo) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConnectionConfigInfo) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConnectionConfigInfo) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetAuthenticationParameters

`func (o *ConnectionConfigInfo) GetAuthenticationParameters() AuthenticationParameters`

GetAuthenticationParameters returns the AuthenticationParameters field if non-nil, zero value otherwise.

### GetAuthenticationParametersOk

`func (o *ConnectionConfigInfo) GetAuthenticationParametersOk() (*AuthenticationParameters, bool)`

GetAuthenticationParametersOk returns a tuple with the AuthenticationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationParameters

`func (o *ConnectionConfigInfo) SetAuthenticationParameters(v AuthenticationParameters)`

SetAuthenticationParameters sets AuthenticationParameters field to given value.

### HasAuthenticationParameters

`func (o *ConnectionConfigInfo) HasAuthenticationParameters() bool`

HasAuthenticationParameters returns a boolean if a field has been set.

### GetServiceIdentity

`func (o *ConnectionConfigInfo) GetServiceIdentity() ServiceIdentityInfo`

GetServiceIdentity returns the ServiceIdentity field if non-nil, zero value otherwise.

### GetServiceIdentityOk

`func (o *ConnectionConfigInfo) GetServiceIdentityOk() (*ServiceIdentityInfo, bool)`

GetServiceIdentityOk returns a tuple with the ServiceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIdentity

`func (o *ConnectionConfigInfo) SetServiceIdentity(v ServiceIdentityInfo)`

SetServiceIdentity sets ServiceIdentity field to given value.

### HasServiceIdentity

`func (o *ConnectionConfigInfo) HasServiceIdentity() bool`

HasServiceIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


