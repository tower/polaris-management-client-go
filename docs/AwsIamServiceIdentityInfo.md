# AwsIamServiceIdentityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamArn** | **string** | The ARN of the IAM user or IAM role Polaris uses to assume roles and then access external resources. | 

## Methods

### NewAwsIamServiceIdentityInfo

`func NewAwsIamServiceIdentityInfo(iamArn string, ) *AwsIamServiceIdentityInfo`

NewAwsIamServiceIdentityInfo instantiates a new AwsIamServiceIdentityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIamServiceIdentityInfoWithDefaults

`func NewAwsIamServiceIdentityInfoWithDefaults() *AwsIamServiceIdentityInfo`

NewAwsIamServiceIdentityInfoWithDefaults instantiates a new AwsIamServiceIdentityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIamArn

`func (o *AwsIamServiceIdentityInfo) GetIamArn() string`

GetIamArn returns the IamArn field if non-nil, zero value otherwise.

### GetIamArnOk

`func (o *AwsIamServiceIdentityInfo) GetIamArnOk() (*string, bool)`

GetIamArnOk returns a tuple with the IamArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamArn

`func (o *AwsIamServiceIdentityInfo) SetIamArn(v string)`

SetIamArn sets IamArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


