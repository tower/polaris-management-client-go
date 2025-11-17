# \DefaultAPI

All URIs are relative to *https://localhost/api/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGrantToCatalogRole**](DefaultAPI.md#AddGrantToCatalogRole) | **Put** /catalogs/{catalogName}/catalog-roles/{catalogRoleName}/grants | 
[**AssignCatalogRoleToPrincipalRole**](DefaultAPI.md#AssignCatalogRoleToPrincipalRole) | **Put** /principal-roles/{principalRoleName}/catalog-roles/{catalogName} | 
[**AssignPrincipalRole**](DefaultAPI.md#AssignPrincipalRole) | **Put** /principals/{principalName}/principal-roles | 
[**CreateCatalog**](DefaultAPI.md#CreateCatalog) | **Post** /catalogs | 
[**CreateCatalogRole**](DefaultAPI.md#CreateCatalogRole) | **Post** /catalogs/{catalogName}/catalog-roles | 
[**CreatePrincipal**](DefaultAPI.md#CreatePrincipal) | **Post** /principals | 
[**CreatePrincipalRole**](DefaultAPI.md#CreatePrincipalRole) | **Post** /principal-roles | 
[**DeleteCatalog**](DefaultAPI.md#DeleteCatalog) | **Delete** /catalogs/{catalogName} | 
[**DeleteCatalogRole**](DefaultAPI.md#DeleteCatalogRole) | **Delete** /catalogs/{catalogName}/catalog-roles/{catalogRoleName} | 
[**DeletePrincipal**](DefaultAPI.md#DeletePrincipal) | **Delete** /principals/{principalName} | 
[**DeletePrincipalRole**](DefaultAPI.md#DeletePrincipalRole) | **Delete** /principal-roles/{principalRoleName} | 
[**GetCatalog**](DefaultAPI.md#GetCatalog) | **Get** /catalogs/{catalogName} | 
[**GetCatalogRole**](DefaultAPI.md#GetCatalogRole) | **Get** /catalogs/{catalogName}/catalog-roles/{catalogRoleName} | 
[**GetPrincipal**](DefaultAPI.md#GetPrincipal) | **Get** /principals/{principalName} | 
[**GetPrincipalRole**](DefaultAPI.md#GetPrincipalRole) | **Get** /principal-roles/{principalRoleName} | 
[**ListAssigneePrincipalRolesForCatalogRole**](DefaultAPI.md#ListAssigneePrincipalRolesForCatalogRole) | **Get** /catalogs/{catalogName}/catalog-roles/{catalogRoleName}/principal-roles | 
[**ListAssigneePrincipalsForPrincipalRole**](DefaultAPI.md#ListAssigneePrincipalsForPrincipalRole) | **Get** /principal-roles/{principalRoleName}/principals | 
[**ListCatalogRoles**](DefaultAPI.md#ListCatalogRoles) | **Get** /catalogs/{catalogName}/catalog-roles | 
[**ListCatalogRolesForPrincipalRole**](DefaultAPI.md#ListCatalogRolesForPrincipalRole) | **Get** /principal-roles/{principalRoleName}/catalog-roles/{catalogName} | 
[**ListCatalogs**](DefaultAPI.md#ListCatalogs) | **Get** /catalogs | 
[**ListGrantsForCatalogRole**](DefaultAPI.md#ListGrantsForCatalogRole) | **Get** /catalogs/{catalogName}/catalog-roles/{catalogRoleName}/grants | 
[**ListPrincipalRoles**](DefaultAPI.md#ListPrincipalRoles) | **Get** /principal-roles | 
[**ListPrincipalRolesAssigned**](DefaultAPI.md#ListPrincipalRolesAssigned) | **Get** /principals/{principalName}/principal-roles | 
[**ListPrincipals**](DefaultAPI.md#ListPrincipals) | **Get** /principals | 
[**ResetCredentials**](DefaultAPI.md#ResetCredentials) | **Post** /principals/{principalName}/reset | 
[**RevokeCatalogRoleFromPrincipalRole**](DefaultAPI.md#RevokeCatalogRoleFromPrincipalRole) | **Delete** /principal-roles/{principalRoleName}/catalog-roles/{catalogName}/{catalogRoleName} | 
[**RevokeGrantFromCatalogRole**](DefaultAPI.md#RevokeGrantFromCatalogRole) | **Post** /catalogs/{catalogName}/catalog-roles/{catalogRoleName}/grants | 
[**RevokePrincipalRole**](DefaultAPI.md#RevokePrincipalRole) | **Delete** /principals/{principalName}/principal-roles/{principalRoleName} | 
[**RotateCredentials**](DefaultAPI.md#RotateCredentials) | **Post** /principals/{principalName}/rotate | 
[**UpdateCatalog**](DefaultAPI.md#UpdateCatalog) | **Put** /catalogs/{catalogName} | 
[**UpdateCatalogRole**](DefaultAPI.md#UpdateCatalogRole) | **Put** /catalogs/{catalogName}/catalog-roles/{catalogRoleName} | 
[**UpdatePrincipal**](DefaultAPI.md#UpdatePrincipal) | **Put** /principals/{principalName} | 
[**UpdatePrincipalRole**](DefaultAPI.md#UpdatePrincipalRole) | **Put** /principal-roles/{principalRoleName} | 



## AddGrantToCatalogRole

> AddGrantToCatalogRole(ctx, catalogName, catalogRoleName).AddGrantRequest(addGrantRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog where the role will receive the grant
	catalogRoleName := "catalogRoleName_example" // string | The name of the role receiving the grant (must exist)
	addGrantRequest := *openapiclient.NewAddGrantRequest() // AddGrantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddGrantToCatalogRole(context.Background(), catalogName, catalogRoleName).AddGrantRequest(addGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGrantToCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog where the role will receive the grant | 
**catalogRoleName** | **string** | The name of the role receiving the grant (must exist) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGrantToCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addGrantRequest** | [**AddGrantRequest**](AddGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignCatalogRoleToPrincipalRole

> AssignCatalogRoleToPrincipalRole(ctx, principalRoleName, catalogName).GrantCatalogRoleRequest(grantCatalogRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name
	catalogName := "catalogName_example" // string | The name of the catalog where the catalogRoles reside
	grantCatalogRoleRequest := *openapiclient.NewGrantCatalogRoleRequest() // GrantCatalogRoleRequest | The principal to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AssignCatalogRoleToPrincipalRole(context.Background(), principalRoleName, catalogName).GrantCatalogRoleRequest(grantCatalogRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssignCatalogRoleToPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 
**catalogName** | **string** | The name of the catalog where the catalogRoles reside | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignCatalogRoleToPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **grantCatalogRoleRequest** | [**GrantCatalogRoleRequest**](GrantCatalogRoleRequest.md) | The principal to create | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignPrincipalRole

> AssignPrincipalRole(ctx, principalName).GrantPrincipalRoleRequest(grantPrincipalRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The name of the target principal
	grantPrincipalRoleRequest := *openapiclient.NewGrantPrincipalRoleRequest() // GrantPrincipalRoleRequest | The principal role to assign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AssignPrincipalRole(context.Background(), principalName).GrantPrincipalRoleRequest(grantPrincipalRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssignPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The name of the target principal | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantPrincipalRoleRequest** | [**GrantPrincipalRoleRequest**](GrantPrincipalRoleRequest.md) | The principal role to assign | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCatalog

> Catalog CreateCatalog(ctx).CreateCatalogRequest(createCatalogRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	createCatalogRequest := *openapiclient.NewCreateCatalogRequest(*openapiclient.NewCatalog("Type_example", "Name_example", *openapiclient.NewCatalogProperties("DefaultBaseLocation_example"), *openapiclient.NewStorageConfigInfo("StorageType_example"))) // CreateCatalogRequest | The Catalog to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCatalog(context.Background()).CreateCatalogRequest(createCatalogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalog`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCatalogRequest** | [**CreateCatalogRequest**](CreateCatalogRequest.md) | The Catalog to create | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCatalogRole

> CatalogRole CreateCatalogRole(ctx, catalogName).CreateCatalogRoleRequest(createCatalogRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The catalog for which we are reading/updating roles
	createCatalogRoleRequest := *openapiclient.NewCreateCatalogRoleRequest() // CreateCatalogRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCatalogRole(context.Background(), catalogName).CreateCatalogRoleRequest(createCatalogRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalogRole`: CatalogRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCatalogRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The catalog for which we are reading/updating roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCatalogRoleRequest** | [**CreateCatalogRoleRequest**](CreateCatalogRoleRequest.md) |  | 

### Return type

[**CatalogRole**](CatalogRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrincipal

> PrincipalWithCredentials CreatePrincipal(ctx).CreatePrincipalRequest(createPrincipalRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	createPrincipalRequest := *openapiclient.NewCreatePrincipalRequest() // CreatePrincipalRequest | The principal to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePrincipal(context.Background()).CreatePrincipalRequest(createPrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrincipal`: PrincipalWithCredentials
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPrincipalRequest** | [**CreatePrincipalRequest**](CreatePrincipalRequest.md) | The principal to create | 

### Return type

[**PrincipalWithCredentials**](PrincipalWithCredentials.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrincipalRole

> PrincipalRole CreatePrincipalRole(ctx).CreatePrincipalRoleRequest(createPrincipalRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	createPrincipalRoleRequest := *openapiclient.NewCreatePrincipalRoleRequest() // CreatePrincipalRoleRequest | The principal to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePrincipalRole(context.Background()).CreatePrincipalRoleRequest(createPrincipalRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrincipalRole`: PrincipalRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePrincipalRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPrincipalRoleRequest** | [**CreatePrincipalRoleRequest**](CreatePrincipalRoleRequest.md) | The principal to create | 

### Return type

[**PrincipalRole**](PrincipalRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCatalog

> DeleteCatalog(ctx, catalogName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteCatalog(context.Background(), catalogName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCatalogRole

> DeleteCatalogRole(ctx, catalogName, catalogRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The catalog for which we are retrieving roles
	catalogRoleName := "catalogRoleName_example" // string | The name of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteCatalogRole(context.Background(), catalogName, catalogRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The catalog for which we are retrieving roles | 
**catalogRoleName** | **string** | The name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrincipal

> DeletePrincipal(ctx, principalName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The principal name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrincipal(context.Background(), principalName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The principal name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrincipalRole

> DeletePrincipalRole(ctx, principalRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrincipalRole(context.Background(), principalRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalog

> Catalog GetCatalog(ctx, catalogName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCatalog(context.Background(), catalogName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalog`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Catalog**](Catalog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogRole

> CatalogRole GetCatalogRole(ctx, catalogName, catalogRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The catalog for which we are retrieving roles
	catalogRoleName := "catalogRoleName_example" // string | The name of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCatalogRole(context.Background(), catalogName, catalogRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogRole`: CatalogRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCatalogRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The catalog for which we are retrieving roles | 
**catalogRoleName** | **string** | The name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CatalogRole**](CatalogRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipal

> Principal GetPrincipal(ctx, principalName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The principal name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrincipal(context.Background(), principalName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipal`: Principal
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The principal name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Principal**](Principal.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalRole

> PrincipalRole GetPrincipalRole(ctx, principalRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrincipalRole(context.Background(), principalRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalRole`: PrincipalRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrincipalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalRole**](PrincipalRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssigneePrincipalRolesForCatalogRole

> PrincipalRoles ListAssigneePrincipalRolesForCatalogRole(ctx, catalogName, catalogRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog where the catalog role resides
	catalogRoleName := "catalogRoleName_example" // string | The name of the catalog role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListAssigneePrincipalRolesForCatalogRole(context.Background(), catalogName, catalogRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListAssigneePrincipalRolesForCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssigneePrincipalRolesForCatalogRole`: PrincipalRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListAssigneePrincipalRolesForCatalogRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog where the catalog role resides | 
**catalogRoleName** | **string** | The name of the catalog role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssigneePrincipalRolesForCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrincipalRoles**](PrincipalRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssigneePrincipalsForPrincipalRole

> Principals ListAssigneePrincipalsForPrincipalRole(ctx, principalRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListAssigneePrincipalsForPrincipalRole(context.Background(), principalRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListAssigneePrincipalsForPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssigneePrincipalsForPrincipalRole`: Principals
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListAssigneePrincipalsForPrincipalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssigneePrincipalsForPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Principals**](Principals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogRoles

> CatalogRoles ListCatalogRoles(ctx, catalogName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The catalog for which we are reading/updating roles

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListCatalogRoles(context.Background(), catalogName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListCatalogRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCatalogRoles`: CatalogRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListCatalogRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The catalog for which we are reading/updating roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogRoles**](CatalogRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogRolesForPrincipalRole

> CatalogRoles ListCatalogRolesForPrincipalRole(ctx, principalRoleName, catalogName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name
	catalogName := "catalogName_example" // string | The name of the catalog where the catalogRoles reside

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListCatalogRolesForPrincipalRole(context.Background(), principalRoleName, catalogName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListCatalogRolesForPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCatalogRolesForPrincipalRole`: CatalogRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListCatalogRolesForPrincipalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 
**catalogName** | **string** | The name of the catalog where the catalogRoles reside | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogRolesForPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CatalogRoles**](CatalogRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogs

> Catalogs ListCatalogs(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListCatalogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCatalogs`: Catalogs
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListCatalogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogsRequest struct via the builder pattern


### Return type

[**Catalogs**](Catalogs.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrantsForCatalogRole

> GrantResources ListGrantsForCatalogRole(ctx, catalogName, catalogRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog where the role will receive the grant
	catalogRoleName := "catalogRoleName_example" // string | The name of the role receiving the grant (must exist)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGrantsForCatalogRole(context.Background(), catalogName, catalogRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGrantsForCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrantsForCatalogRole`: GrantResources
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGrantsForCatalogRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog where the role will receive the grant | 
**catalogRoleName** | **string** | The name of the role receiving the grant (must exist) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsForCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GrantResources**](GrantResources.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrincipalRoles

> PrincipalRoles ListPrincipalRoles(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPrincipalRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPrincipalRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrincipalRoles`: PrincipalRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPrincipalRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalRolesRequest struct via the builder pattern


### Return type

[**PrincipalRoles**](PrincipalRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrincipalRolesAssigned

> PrincipalRoles ListPrincipalRolesAssigned(ctx, principalName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The name of the target principal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPrincipalRolesAssigned(context.Background(), principalName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPrincipalRolesAssigned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrincipalRolesAssigned`: PrincipalRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPrincipalRolesAssigned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The name of the target principal | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalRolesAssignedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalRoles**](PrincipalRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrincipals

> Principals ListPrincipals(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPrincipals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrincipals`: Principals
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPrincipals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalsRequest struct via the builder pattern


### Return type

[**Principals**](Principals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCredentials

> PrincipalWithCredentials ResetCredentials(ctx, principalName).ResetPrincipalRequest(resetPrincipalRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The principal's name
	resetPrincipalRequest := *openapiclient.NewResetPrincipalRequest() // ResetPrincipalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ResetCredentials(context.Background(), principalName).ResetPrincipalRequest(resetPrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResetCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetCredentials`: PrincipalWithCredentials
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ResetCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The principal&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetPrincipalRequest** | [**ResetPrincipalRequest**](ResetPrincipalRequest.md) |  | 

### Return type

[**PrincipalWithCredentials**](PrincipalWithCredentials.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCatalogRoleFromPrincipalRole

> RevokeCatalogRoleFromPrincipalRole(ctx, principalRoleName, catalogName, catalogRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name
	catalogName := "catalogName_example" // string | The name of the catalog that contains the role to revoke
	catalogRoleName := "catalogRoleName_example" // string | The name of the catalog role that should be revoked

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokeCatalogRoleFromPrincipalRole(context.Background(), principalRoleName, catalogName, catalogRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokeCatalogRoleFromPrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 
**catalogName** | **string** | The name of the catalog that contains the role to revoke | 
**catalogRoleName** | **string** | The name of the catalog role that should be revoked | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCatalogRoleFromPrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeGrantFromCatalogRole

> RevokeGrantFromCatalogRole(ctx, catalogName, catalogRoleName).Cascade(cascade).RevokeGrantRequest(revokeGrantRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog where the role will receive the grant
	catalogRoleName := "catalogRoleName_example" // string | The name of the role receiving the grant (must exist)
	cascade := true // bool | If true, the grant revocation cascades to all subresources. (optional) (default to false)
	revokeGrantRequest := *openapiclient.NewRevokeGrantRequest() // RevokeGrantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokeGrantFromCatalogRole(context.Background(), catalogName, catalogRoleName).Cascade(cascade).RevokeGrantRequest(revokeGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokeGrantFromCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog where the role will receive the grant | 
**catalogRoleName** | **string** | The name of the role receiving the grant (must exist) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeGrantFromCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascade** | **bool** | If true, the grant revocation cascades to all subresources. | [default to false]
 **revokeGrantRequest** | [**RevokeGrantRequest**](RevokeGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrincipalRole

> RevokePrincipalRole(ctx, principalName, principalRoleName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The name of the target principal
	principalRoleName := "principalRoleName_example" // string | The name of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokePrincipalRole(context.Background(), principalName, principalRoleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokePrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The name of the target principal | 
**principalRoleName** | **string** | The name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateCredentials

> PrincipalWithCredentials RotateCredentials(ctx, principalName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The user name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RotateCredentials(context.Background(), principalName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RotateCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateCredentials`: PrincipalWithCredentials
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RotateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The user name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalWithCredentials**](PrincipalWithCredentials.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalog

> Catalog UpdateCatalog(ctx, catalogName).UpdateCatalogRequest(updateCatalogRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The name of the catalog
	updateCatalogRequest := *openapiclient.NewUpdateCatalogRequest() // UpdateCatalogRequest | The catalog details to use in the update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateCatalog(context.Background(), catalogName).UpdateCatalogRequest(updateCatalogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalog`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The name of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCatalogRequest** | [**UpdateCatalogRequest**](UpdateCatalogRequest.md) | The catalog details to use in the update | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogRole

> CatalogRole UpdateCatalogRole(ctx, catalogName, catalogRoleName).UpdateCatalogRoleRequest(updateCatalogRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	catalogName := "catalogName_example" // string | The catalog for which we are retrieving roles
	catalogRoleName := "catalogRoleName_example" // string | The name of the role
	updateCatalogRoleRequest := *openapiclient.NewUpdateCatalogRoleRequest(int32(123), map[string]string{"key": "Inner_example"}) // UpdateCatalogRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateCatalogRole(context.Background(), catalogName, catalogRoleName).UpdateCatalogRoleRequest(updateCatalogRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCatalogRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogRole`: CatalogRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCatalogRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogName** | **string** | The catalog for which we are retrieving roles | 
**catalogRoleName** | **string** | The name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCatalogRoleRequest** | [**UpdateCatalogRoleRequest**](UpdateCatalogRoleRequest.md) |  | 

### Return type

[**CatalogRole**](CatalogRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrincipal

> Principal UpdatePrincipal(ctx, principalName).UpdatePrincipalRequest(updatePrincipalRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalName := "principalName_example" // string | The principal name
	updatePrincipalRequest := *openapiclient.NewUpdatePrincipalRequest(int32(123), map[string]string{"key": "Inner_example"}) // UpdatePrincipalRequest | The principal details to use in the update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrincipal(context.Background(), principalName).UpdatePrincipalRequest(updatePrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrincipal`: Principal
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalName** | **string** | The principal name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePrincipalRequest** | [**UpdatePrincipalRequest**](UpdatePrincipalRequest.md) | The principal details to use in the update | 

### Return type

[**Principal**](Principal.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrincipalRole

> PrincipalRole UpdatePrincipalRole(ctx, principalRoleName).UpdatePrincipalRoleRequest(updatePrincipalRoleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tower/polaris-management-client-go"
)

func main() {
	principalRoleName := "principalRoleName_example" // string | The principal role name
	updatePrincipalRoleRequest := *openapiclient.NewUpdatePrincipalRoleRequest(int32(123), map[string]string{"key": "Inner_example"}) // UpdatePrincipalRoleRequest | The principalRole details to use in the update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrincipalRole(context.Background(), principalRoleName).UpdatePrincipalRoleRequest(updatePrincipalRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrincipalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrincipalRole`: PrincipalRole
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrincipalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRoleName** | **string** | The principal role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePrincipalRoleRequest** | [**UpdatePrincipalRoleRequest**](UpdatePrincipalRoleRequest.md) | The principalRole details to use in the update | 

### Return type

[**PrincipalRole**](PrincipalRole.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

