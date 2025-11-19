package snowflakeopencatalog

import (
	"context"
	"net/http"

	polarismgmt "github.com/tower/polaris-management-go"
)

// PolarisClient defines the interface for interacting with Snowflake OpenCatalog/Polaris
// This interface allows for easier testing by enabling mock implementations
type PolarisClient interface {
	// Catalog operations
	CreateCatalog(ctx context.Context, req polarismgmt.CreateCatalogRequest) (*polarismgmt.Catalog, error)
	GetCatalog(ctx context.Context, catalogName string) (*polarismgmt.Catalog, error)
	DeleteCatalog(ctx context.Context, catalogName string) error

	// Principal operations
	CreatePrincipal(ctx context.Context, req polarismgmt.CreatePrincipalRequest) (*polarismgmt.PrincipalWithCredentials, error)
	GetPrincipal(ctx context.Context, principalName string) (*polarismgmt.Principal, error)
	DeletePrincipal(ctx context.Context, principalName string) error

	// Principal Role operations
	GetPrincipalRole(ctx context.Context, roleName string) (*polarismgmt.PrincipalRole, error)
	CreatePrincipalRole(ctx context.Context, req polarismgmt.CreatePrincipalRoleRequest) (*polarismgmt.PrincipalRole, error)
	AssignPrincipalRole(ctx context.Context, principalName string, req polarismgmt.GrantPrincipalRoleRequest) error

	// Catalog Role operations
	CreateCatalogRole(ctx context.Context, catalogName string, req polarismgmt.CreateCatalogRoleRequest) (*polarismgmt.CatalogRole, error)
	AddGrantToCatalogRole(ctx context.Context, catalogName string, catalogRoleName string, req polarismgmt.AddGrantRequest) (*http.Response, error)
	DeleteCatalogRole(ctx context.Context, catalogName string, roleName string) error
	AssignCatalogRoleToPrincipalRole(ctx context.Context, principalRoleName string, catalogName string, req polarismgmt.GrantCatalogRoleRequest) error
}

// Ensure Client implements PolarisClient
var _ PolarisClient = (*Client)(nil)
