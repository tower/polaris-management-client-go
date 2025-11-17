# AzureStorageConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | the tenant id that the storage accounts belong to | 
**MultiTenantAppName** | Pointer to **string** | the name of the azure client application | [optional] 
**ConsentUrl** | Pointer to **string** | URL to the Azure permissions request page | [optional] 

## Methods

### NewAzureStorageConfigInfo

`func NewAzureStorageConfigInfo(tenantId string, ) *AzureStorageConfigInfo`

NewAzureStorageConfigInfo instantiates a new AzureStorageConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageConfigInfoWithDefaults

`func NewAzureStorageConfigInfoWithDefaults() *AzureStorageConfigInfo`

NewAzureStorageConfigInfoWithDefaults instantiates a new AzureStorageConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureStorageConfigInfo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureStorageConfigInfo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureStorageConfigInfo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMultiTenantAppName

`func (o *AzureStorageConfigInfo) GetMultiTenantAppName() string`

GetMultiTenantAppName returns the MultiTenantAppName field if non-nil, zero value otherwise.

### GetMultiTenantAppNameOk

`func (o *AzureStorageConfigInfo) GetMultiTenantAppNameOk() (*string, bool)`

GetMultiTenantAppNameOk returns a tuple with the MultiTenantAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenantAppName

`func (o *AzureStorageConfigInfo) SetMultiTenantAppName(v string)`

SetMultiTenantAppName sets MultiTenantAppName field to given value.

### HasMultiTenantAppName

`func (o *AzureStorageConfigInfo) HasMultiTenantAppName() bool`

HasMultiTenantAppName returns a boolean if a field has been set.

### GetConsentUrl

`func (o *AzureStorageConfigInfo) GetConsentUrl() string`

GetConsentUrl returns the ConsentUrl field if non-nil, zero value otherwise.

### GetConsentUrlOk

`func (o *AzureStorageConfigInfo) GetConsentUrlOk() (*string, bool)`

GetConsentUrlOk returns a tuple with the ConsentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentUrl

`func (o *AzureStorageConfigInfo) SetConsentUrl(v string)`

SetConsentUrl sets ConsentUrl field to given value.

### HasConsentUrl

`func (o *AzureStorageConfigInfo) HasConsentUrl() bool`

HasConsentUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


