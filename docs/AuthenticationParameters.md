# AuthenticationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** | The type of authentication to use when connecting to the remote rest service | 

## Methods

### NewAuthenticationParameters

`func NewAuthenticationParameters(authenticationType string, ) *AuthenticationParameters`

NewAuthenticationParameters instantiates a new AuthenticationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationParametersWithDefaults

`func NewAuthenticationParametersWithDefaults() *AuthenticationParameters`

NewAuthenticationParametersWithDefaults instantiates a new AuthenticationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AuthenticationParameters) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AuthenticationParameters) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AuthenticationParameters) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


