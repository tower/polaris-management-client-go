# OAuthClientCredentialsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenUri** | Pointer to **string** | Token server URI | [optional] 
**ClientId** | Pointer to **string** | oauth client id | [optional] 
**ClientSecret** | Pointer to **string** | oauth client secret (input-only) | [optional] 
**Scopes** | Pointer to **[]string** | oauth scopes to specify when exchanging for a short-lived access token | [optional] 

## Methods

### NewOAuthClientCredentialsParameters

`func NewOAuthClientCredentialsParameters() *OAuthClientCredentialsParameters`

NewOAuthClientCredentialsParameters instantiates a new OAuthClientCredentialsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientCredentialsParametersWithDefaults

`func NewOAuthClientCredentialsParametersWithDefaults() *OAuthClientCredentialsParameters`

NewOAuthClientCredentialsParametersWithDefaults instantiates a new OAuthClientCredentialsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenUri

`func (o *OAuthClientCredentialsParameters) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *OAuthClientCredentialsParameters) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *OAuthClientCredentialsParameters) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *OAuthClientCredentialsParameters) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthClientCredentialsParameters) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClientCredentialsParameters) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClientCredentialsParameters) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthClientCredentialsParameters) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuthClientCredentialsParameters) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthClientCredentialsParameters) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthClientCredentialsParameters) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthClientCredentialsParameters) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *OAuthClientCredentialsParameters) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthClientCredentialsParameters) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthClientCredentialsParameters) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthClientCredentialsParameters) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


