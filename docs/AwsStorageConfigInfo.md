# AwsStorageConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **string** | the aws role arn that grants privileges on the S3 buckets | [optional] 
**ExternalId** | Pointer to **string** | an optional external id used to establish a trust relationship with AWS in the trust policy | [optional] 
**UserArn** | Pointer to **string** | the aws user arn used to assume the aws role | [optional] 
**Region** | Pointer to **string** | the aws region where data is stored | [optional] 
**Endpoint** | Pointer to **string** | endpoint for S3 requests (optional). Clients always see this value (if it is set). Polaris Servers may be configured to use a different endpoint URI via the &#x60;endpointInternal&#x60; property. | [optional] 
**StsEndpoint** | Pointer to **string** | endpoint for STS requests made by the Polaris Server (optional). If not set, defaults to &#39;endpointInternal&#39; (which in turn defaults to &#x60;endpoint&#x60;). | [optional] 
**StsUnavailable** | Pointer to **bool** | if set to &#x60;true&#x60;, instructs Polaris Servers to avoid using the STS endpoints when obtaining credentials for accessing data and metadata files within the related catalog. Setting this property to &#x60;true&#x60; effectively disables vending storage credentials to clients. This setting is intended for configuring catalogs with S3-compatible storage implementations that do not support STS. | [optional] 
**EndpointInternal** | Pointer to **string** | endpoint for S3 requests made by the Polaris Server (optional). If set, Polaris Service will use this value instead of &#x60;endpoint&#x60;. If not set, defaults to &#x60;endpoint&#x60;. Iceberg REST API clients never see this value. | [optional] 
**PathStyleAccess** | Pointer to **bool** | Whether S3 requests to files in this catalog should use &#39;path-style addressing for buckets&#39;. | [optional] [default to false]

## Methods

### NewAwsStorageConfigInfo

`func NewAwsStorageConfigInfo() *AwsStorageConfigInfo`

NewAwsStorageConfigInfo instantiates a new AwsStorageConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsStorageConfigInfoWithDefaults

`func NewAwsStorageConfigInfoWithDefaults() *AwsStorageConfigInfo`

NewAwsStorageConfigInfoWithDefaults instantiates a new AwsStorageConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsStorageConfigInfo) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsStorageConfigInfo) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsStorageConfigInfo) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AwsStorageConfigInfo) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetExternalId

`func (o *AwsStorageConfigInfo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsStorageConfigInfo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsStorageConfigInfo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsStorageConfigInfo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUserArn

`func (o *AwsStorageConfigInfo) GetUserArn() string`

GetUserArn returns the UserArn field if non-nil, zero value otherwise.

### GetUserArnOk

`func (o *AwsStorageConfigInfo) GetUserArnOk() (*string, bool)`

GetUserArnOk returns a tuple with the UserArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserArn

`func (o *AwsStorageConfigInfo) SetUserArn(v string)`

SetUserArn sets UserArn field to given value.

### HasUserArn

`func (o *AwsStorageConfigInfo) HasUserArn() bool`

HasUserArn returns a boolean if a field has been set.

### GetRegion

`func (o *AwsStorageConfigInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsStorageConfigInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsStorageConfigInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsStorageConfigInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEndpoint

`func (o *AwsStorageConfigInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AwsStorageConfigInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AwsStorageConfigInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AwsStorageConfigInfo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetStsEndpoint

`func (o *AwsStorageConfigInfo) GetStsEndpoint() string`

GetStsEndpoint returns the StsEndpoint field if non-nil, zero value otherwise.

### GetStsEndpointOk

`func (o *AwsStorageConfigInfo) GetStsEndpointOk() (*string, bool)`

GetStsEndpointOk returns a tuple with the StsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEndpoint

`func (o *AwsStorageConfigInfo) SetStsEndpoint(v string)`

SetStsEndpoint sets StsEndpoint field to given value.

### HasStsEndpoint

`func (o *AwsStorageConfigInfo) HasStsEndpoint() bool`

HasStsEndpoint returns a boolean if a field has been set.

### GetStsUnavailable

`func (o *AwsStorageConfigInfo) GetStsUnavailable() bool`

GetStsUnavailable returns the StsUnavailable field if non-nil, zero value otherwise.

### GetStsUnavailableOk

`func (o *AwsStorageConfigInfo) GetStsUnavailableOk() (*bool, bool)`

GetStsUnavailableOk returns a tuple with the StsUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUnavailable

`func (o *AwsStorageConfigInfo) SetStsUnavailable(v bool)`

SetStsUnavailable sets StsUnavailable field to given value.

### HasStsUnavailable

`func (o *AwsStorageConfigInfo) HasStsUnavailable() bool`

HasStsUnavailable returns a boolean if a field has been set.

### GetEndpointInternal

`func (o *AwsStorageConfigInfo) GetEndpointInternal() string`

GetEndpointInternal returns the EndpointInternal field if non-nil, zero value otherwise.

### GetEndpointInternalOk

`func (o *AwsStorageConfigInfo) GetEndpointInternalOk() (*string, bool)`

GetEndpointInternalOk returns a tuple with the EndpointInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointInternal

`func (o *AwsStorageConfigInfo) SetEndpointInternal(v string)`

SetEndpointInternal sets EndpointInternal field to given value.

### HasEndpointInternal

`func (o *AwsStorageConfigInfo) HasEndpointInternal() bool`

HasEndpointInternal returns a boolean if a field has been set.

### GetPathStyleAccess

`func (o *AwsStorageConfigInfo) GetPathStyleAccess() bool`

GetPathStyleAccess returns the PathStyleAccess field if non-nil, zero value otherwise.

### GetPathStyleAccessOk

`func (o *AwsStorageConfigInfo) GetPathStyleAccessOk() (*bool, bool)`

GetPathStyleAccessOk returns a tuple with the PathStyleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathStyleAccess

`func (o *AwsStorageConfigInfo) SetPathStyleAccess(v bool)`

SetPathStyleAccess sets PathStyleAccess field to given value.

### HasPathStyleAccess

`func (o *AwsStorageConfigInfo) HasPathStyleAccess() bool`

HasPathStyleAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


