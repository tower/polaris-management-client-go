# IcebergRestConnectionConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteCatalogName** | Pointer to **string** | The name of a remote catalog instance within the remote catalog service; in some older systems this is specified as the &#39;warehouse&#39; when multiple logical catalogs are served under the same base uri, and often translates into a &#39;prefix&#39; added to all REST resource paths | [optional] 

## Methods

### NewIcebergRestConnectionConfigInfo

`func NewIcebergRestConnectionConfigInfo() *IcebergRestConnectionConfigInfo`

NewIcebergRestConnectionConfigInfo instantiates a new IcebergRestConnectionConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIcebergRestConnectionConfigInfoWithDefaults

`func NewIcebergRestConnectionConfigInfoWithDefaults() *IcebergRestConnectionConfigInfo`

NewIcebergRestConnectionConfigInfoWithDefaults instantiates a new IcebergRestConnectionConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteCatalogName

`func (o *IcebergRestConnectionConfigInfo) GetRemoteCatalogName() string`

GetRemoteCatalogName returns the RemoteCatalogName field if non-nil, zero value otherwise.

### GetRemoteCatalogNameOk

`func (o *IcebergRestConnectionConfigInfo) GetRemoteCatalogNameOk() (*string, bool)`

GetRemoteCatalogNameOk returns a tuple with the RemoteCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCatalogName

`func (o *IcebergRestConnectionConfigInfo) SetRemoteCatalogName(v string)`

SetRemoteCatalogName sets RemoteCatalogName field to given value.

### HasRemoteCatalogName

`func (o *IcebergRestConnectionConfigInfo) HasRemoteCatalogName() bool`

HasRemoteCatalogName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


