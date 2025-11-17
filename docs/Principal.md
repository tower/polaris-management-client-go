# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ClientId** | Pointer to **string** | The output-only OAuth clientId associated with this principal if applicable | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**CreateTimestamp** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**EntityVersion** | Pointer to **int32** | The version of the principal object used to determine if the principal metadata has changed | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal(name string, ) *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Principal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Principal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Principal) SetName(v string)`

SetName sets Name field to given value.


### GetClientId

`func (o *Principal) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Principal) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Principal) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Principal) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProperties

`func (o *Principal) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Principal) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Principal) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Principal) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *Principal) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *Principal) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *Principal) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *Principal) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Principal) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Principal) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Principal) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Principal) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetEntityVersion

`func (o *Principal) GetEntityVersion() int32`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *Principal) GetEntityVersionOk() (*int32, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *Principal) SetEntityVersion(v int32)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *Principal) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


