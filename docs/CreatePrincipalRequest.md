# CreatePrincipalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to [**Principal**](Principal.md) |  | [optional] 
**CredentialRotationRequired** | Pointer to **bool** | If true, the initial credentials can only be used to call rotateCredentials | [optional] 

## Methods

### NewCreatePrincipalRequest

`func NewCreatePrincipalRequest() *CreatePrincipalRequest`

NewCreatePrincipalRequest instantiates a new CreatePrincipalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrincipalRequestWithDefaults

`func NewCreatePrincipalRequestWithDefaults() *CreatePrincipalRequest`

NewCreatePrincipalRequestWithDefaults instantiates a new CreatePrincipalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *CreatePrincipalRequest) GetPrincipal() Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CreatePrincipalRequest) GetPrincipalOk() (*Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CreatePrincipalRequest) SetPrincipal(v Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *CreatePrincipalRequest) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetCredentialRotationRequired

`func (o *CreatePrincipalRequest) GetCredentialRotationRequired() bool`

GetCredentialRotationRequired returns the CredentialRotationRequired field if non-nil, zero value otherwise.

### GetCredentialRotationRequiredOk

`func (o *CreatePrincipalRequest) GetCredentialRotationRequiredOk() (*bool, bool)`

GetCredentialRotationRequiredOk returns a tuple with the CredentialRotationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRotationRequired

`func (o *CreatePrincipalRequest) SetCredentialRotationRequired(v bool)`

SetCredentialRotationRequired sets CredentialRotationRequired field to given value.

### HasCredentialRotationRequired

`func (o *CreatePrincipalRequest) HasCredentialRotationRequired() bool`

HasCredentialRotationRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


