# SigV4AuthenticationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | The aws IAM role arn assumed by polaris userArn when signing requests | 
**RoleSessionName** | Pointer to **string** | The role session name to be used by the SigV4 protocol for signing requests | [optional] 
**ExternalId** | Pointer to **string** | An optional external id used to establish a trust relationship with AWS in the trust policy | [optional] 
**SigningRegion** | **string** | Region to be used by the SigV4 protocol for signing requests | 
**SigningName** | Pointer to **string** | The service name to be used by the SigV4 protocol for signing requests, the default signing name is \&quot;execute-api\&quot; is if not provided | [optional] 

## Methods

### NewSigV4AuthenticationParameters

`func NewSigV4AuthenticationParameters(roleArn string, signingRegion string, ) *SigV4AuthenticationParameters`

NewSigV4AuthenticationParameters instantiates a new SigV4AuthenticationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigV4AuthenticationParametersWithDefaults

`func NewSigV4AuthenticationParametersWithDefaults() *SigV4AuthenticationParameters`

NewSigV4AuthenticationParametersWithDefaults instantiates a new SigV4AuthenticationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *SigV4AuthenticationParameters) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SigV4AuthenticationParameters) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SigV4AuthenticationParameters) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRoleSessionName

`func (o *SigV4AuthenticationParameters) GetRoleSessionName() string`

GetRoleSessionName returns the RoleSessionName field if non-nil, zero value otherwise.

### GetRoleSessionNameOk

`func (o *SigV4AuthenticationParameters) GetRoleSessionNameOk() (*string, bool)`

GetRoleSessionNameOk returns a tuple with the RoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSessionName

`func (o *SigV4AuthenticationParameters) SetRoleSessionName(v string)`

SetRoleSessionName sets RoleSessionName field to given value.

### HasRoleSessionName

`func (o *SigV4AuthenticationParameters) HasRoleSessionName() bool`

HasRoleSessionName returns a boolean if a field has been set.

### GetExternalId

`func (o *SigV4AuthenticationParameters) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SigV4AuthenticationParameters) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SigV4AuthenticationParameters) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SigV4AuthenticationParameters) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSigningRegion

`func (o *SigV4AuthenticationParameters) GetSigningRegion() string`

GetSigningRegion returns the SigningRegion field if non-nil, zero value otherwise.

### GetSigningRegionOk

`func (o *SigV4AuthenticationParameters) GetSigningRegionOk() (*string, bool)`

GetSigningRegionOk returns a tuple with the SigningRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRegion

`func (o *SigV4AuthenticationParameters) SetSigningRegion(v string)`

SetSigningRegion sets SigningRegion field to given value.


### GetSigningName

`func (o *SigV4AuthenticationParameters) GetSigningName() string`

GetSigningName returns the SigningName field if non-nil, zero value otherwise.

### GetSigningNameOk

`func (o *SigV4AuthenticationParameters) GetSigningNameOk() (*string, bool)`

GetSigningNameOk returns a tuple with the SigningName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningName

`func (o *SigV4AuthenticationParameters) SetSigningName(v string)`

SetSigningName sets SigningName field to given value.

### HasSigningName

`func (o *SigV4AuthenticationParameters) HasSigningName() bool`

HasSigningName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


