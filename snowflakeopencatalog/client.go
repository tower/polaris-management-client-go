package snowflakeopencatalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	polarismgmt "github.com/tower/polaris-management-go"
	"golang.org/x/oauth2"
)

// staticTokenSource implements oauth2.TokenSource for a static bearer token
type staticTokenSource struct {
	token string
}

func (s *staticTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: s.token,
		TokenType:   "Bearer",
	}, nil
}

// Client wraps the Polaris Management API client with automatic token injection
type Client struct {
	authClient *AuthClient
	apiClient  *polarismgmt.APIClient
}

// NewClient creates a new Polaris Management API client
func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	authClient, err := NewAuthClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// Create configuration for the generated client
	apiCfg := polarismgmt.NewConfiguration()

	return &Client{
		authClient: authClient,
		apiClient:  polarismgmt.NewAPIClient(apiCfg),
	}, nil
}

// GetAPIClient returns the underlying Polaris Management API client
func (c *Client) GetAPIClient() *polarismgmt.APIClient {
	return c.apiClient
}

// logAPIError extracts and logs detailed error information from Polaris API errors
// including the response body, error message, and model details
func logAPIError(ctx context.Context, err error, operation string) {
	if err == nil {
		return
	}

	// Try to extract GenericOpenAPIError which contains body, error, and model
	if apiErr, ok := err.(*polarismgmt.GenericOpenAPIError); ok {
		body := apiErr.Body()
		errMsg := apiErr.Error()
		model := apiErr.Model()

		// Parse body as JSON for better readability
		var bodyJSON map[string]interface{}
		bodyStr := string(body)
		if len(body) > 0 && json.Unmarshal(body, &bodyJSON) == nil {
			prettyJSON, _ := json.MarshalIndent(bodyJSON, "", "  ")
			bodyStr = string(prettyJSON)
		}

		logger.WithContext(ctx).WithError(err).Errorf(
			"%s failed - Error: %s, Body: %s, Model: %+v",
			operation,
			errMsg,
			bodyStr,
			model,
		)
	} else {
		// Fallback for non-OpenAPI errors
		logger.WithContext(ctx).WithError(err).Errorf("%s failed", operation)
	}
}

// NewContext creates a context with OAuth token and server variables injected
func (c *Client) NewContext(ctx context.Context) (context.Context, error) {
	accessToken, err := c.authClient.GetAccessToken(ctx)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to get access token")
		return nil, ErrFailedToGetAccessToken
	}

	// Parse the base URL to extract scheme and host
	baseURL := c.authClient.GetManagementBaseURL()
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Errorf("failed to parse base URL: %s", baseURL)
		return nil, ErrFailedToGetAccessToken
	}

	// Extract host with /polaris prefix
	// For https://account.snowflakecomputing.com/polaris
	// We want: host = "account.snowflakecomputing.com/polaris"
	hostWithPath := parsedURL.Host
	if parsedURL.Path != "" && parsedURL.Path != "/" {
		hostWithPath = parsedURL.Host + parsedURL.Path
	}

	// Set server variables - the OpenAPI spec defines: "{scheme}://{host}/api/management/v1"
	ctx = context.WithValue(ctx, polarismgmt.ContextServerVariables, map[string]string{
		"scheme": parsedURL.Scheme,
		"host":   hostWithPath,
	})

	// Add OAuth2 token to context
	tokenSource := &staticTokenSource{token: accessToken}
	ctx = context.WithValue(ctx, polarismgmt.ContextOAuth2, tokenSource)

	return ctx, nil
}

// ListCatalogs lists all catalogs in the Polaris account
func (c *Client) ListCatalogs(ctx context.Context) (*polarismgmt.Catalogs, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.ListCatalogs(ctx).Execute()
	if err != nil {
		logAPIError(ctx, err, "list catalogs")
		return nil, ErrFailedToListCatalogs
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// CreateCatalog creates a new catalog in Polaris
func (c *Client) CreateCatalog(ctx context.Context, req polarismgmt.CreateCatalogRequest) (*polarismgmt.Catalog, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.CreateCatalog(ctx).CreateCatalogRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, "create catalog")
		return nil, ErrFailedToCreateCatalog
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// GetCatalog retrieves a catalog by name
func (c *Client) GetCatalog(ctx context.Context, catalogName string) (*polarismgmt.Catalog, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.GetCatalog(ctx, catalogName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("get catalog: %s", catalogName))
		return nil, ErrFailedToGetCatalog
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// DeleteCatalog deletes a catalog by name
func (c *Client) DeleteCatalog(ctx context.Context, catalogName string) error {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return err
	}

	httpResp, err := c.apiClient.DefaultAPI.DeleteCatalog(ctx, catalogName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("delete catalog: %s", catalogName))

		// Check if the error is due to catalog not being empty to return a specific error.
		// Snowflake OpenCatalog returns: "Catalog '...' cannot be dropped, it is not empty"
		if apiErr, ok := err.(*polarismgmt.GenericOpenAPIError); ok {
			body := apiErr.Body()
			errMsg := strings.ToLower(string(body))

			// Check for "not empty" error message from Polaris
			if strings.Contains(errMsg, "not empty") || strings.Contains(errMsg, "cannot be dropped") {
				logger.WithContext(ctx).Error("catalog is not empty - must delete all tables/namespaces first")
				return ErrCatalogNotEmpty
			}
		}

		return ErrFailedToDeleteCatalog
	}
	defer httpResp.Body.Close()

	return nil
}

// ListPrincipals lists all service principals
func (c *Client) ListPrincipals(ctx context.Context) (*polarismgmt.Principals, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.ListPrincipals(ctx).Execute()
	if err != nil {
		logAPIError(ctx, err, "list principals")
		return nil, ErrFailedToListPrincipals
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// CreatePrincipal creates a new service principal
func (c *Client) CreatePrincipal(ctx context.Context, req polarismgmt.CreatePrincipalRequest) (*polarismgmt.PrincipalWithCredentials, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.CreatePrincipal(ctx).CreatePrincipalRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, "create principal")
		return nil, ErrFailedToCreatePrincipal
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// GetPrincipal retrieves a principal by name
func (c *Client) GetPrincipal(ctx context.Context, principalName string) (*polarismgmt.Principal, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.GetPrincipal(ctx, principalName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("get principal: %s", principalName))
		return nil, ErrFailedToGetPrincipal
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// DeletePrincipal deletes a principal by name
func (c *Client) DeletePrincipal(ctx context.Context, principalName string) error {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return err
	}

	httpResp, err := c.apiClient.DefaultAPI.DeletePrincipal(ctx, principalName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("delete principal: %s", principalName))
		return ErrFailedToDeletePrincipal
	}
	defer httpResp.Body.Close()

	return nil
}

// ListPrincipalRoles lists all principal roles
func (c *Client) ListPrincipalRoles(ctx context.Context) (*polarismgmt.PrincipalRoles, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.ListPrincipalRoles(ctx).Execute()
	if err != nil {
		logAPIError(ctx, err, "list principal roles")
		return nil, ErrFailedToListPrincipalRoles
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// RotatePrincipalCredentials rotates the credentials for a service principal
func (c *Client) RotatePrincipalCredentials(ctx context.Context, principalName string) (*polarismgmt.PrincipalWithCredentials, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.RotateCredentials(ctx, principalName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("rotate credentials for principal: %s", principalName))
		return nil, ErrFailedToRotateCredentials
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// GetPrincipalRole retrieves a principal role by name
func (c *Client) GetPrincipalRole(ctx context.Context, roleName string) (*polarismgmt.PrincipalRole, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.GetPrincipalRole(ctx, roleName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("get principal role: %s", roleName))
		return nil, ErrFailedToGetPrincipalRole
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// CreatePrincipalRole creates a new principal role
func (c *Client) CreatePrincipalRole(ctx context.Context, req polarismgmt.CreatePrincipalRoleRequest) (*polarismgmt.PrincipalRole, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.CreatePrincipalRole(ctx).CreatePrincipalRoleRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, "create principal role")
		return nil, ErrFailedToCreatePrincipalRole
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// CreateCatalogRole creates a new catalog role for a specific catalog
func (c *Client) CreateCatalogRole(ctx context.Context, catalogName string, req polarismgmt.CreateCatalogRoleRequest) (*polarismgmt.CatalogRole, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, httpResp, err := c.apiClient.DefaultAPI.CreateCatalogRole(ctx, catalogName).CreateCatalogRoleRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("create catalog role for catalog: %s", catalogName))
		return nil, ErrFailedToCreateCatalogRole
	}
	defer httpResp.Body.Close()

	return resp, nil
}

// AddGrantToCatalogRole adds a grant (privilege) to a catalog role
func (c *Client) AddGrantToCatalogRole(ctx context.Context, catalogName string, catalogRoleName string, req polarismgmt.AddGrantRequest) (*http.Response, error) {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return nil, err
	}

	httpResp, err := c.apiClient.DefaultAPI.AddGrantToCatalogRole(ctx, catalogName, catalogRoleName).AddGrantRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("add grant to catalog role %s in catalog %s", catalogRoleName, catalogName))
		return nil, ErrFailedToAddGrant
	}

	return httpResp, nil
}

// AssignCatalogRoleToPrincipalRole assigns a catalog role to a principal role
func (c *Client) AssignCatalogRoleToPrincipalRole(ctx context.Context, principalRoleName string, catalogName string, req polarismgmt.GrantCatalogRoleRequest) error {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return err
	}

	httpResp, err := c.apiClient.DefaultAPI.AssignCatalogRoleToPrincipalRole(ctx, principalRoleName, catalogName).GrantCatalogRoleRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("assign catalog role to principal role %s", principalRoleName))
		return ErrFailedToAssignCatalogRole
	}
	defer httpResp.Body.Close()

	return nil
}

// AssignPrincipalRole assigns a principal role to a principal (service account)
func (c *Client) AssignPrincipalRole(ctx context.Context, principalName string, req polarismgmt.GrantPrincipalRoleRequest) error {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return err
	}

	httpResp, err := c.apiClient.DefaultAPI.AssignPrincipalRole(ctx, principalName).GrantPrincipalRoleRequest(req).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("assign principal role to principal %s", principalName))
		return ErrFailedToAssignPrincipalRole
	}
	defer httpResp.Body.Close()

	return nil
}

// DeleteCatalogRole deletes a catalog role from a catalog
func (c *Client) DeleteCatalogRole(ctx context.Context, catalogName string, catalogRoleName string) error {
	ctx, err := c.NewContext(ctx)
	if err != nil {
		return err
	}

	httpResp, err := c.apiClient.DefaultAPI.DeleteCatalogRole(ctx, catalogName, catalogRoleName).Execute()
	if err != nil {
		logAPIError(ctx, err, fmt.Sprintf("delete catalog role %s from catalog %s", catalogRoleName, catalogName))
		return ErrFailedToDeleteCatalogRole
	}
	defer httpResp.Body.Close()

	return nil
}

// SanitizeCatalogName converts a Tower catalog name to a valid Polaris catalog name
// Polaris catalog names must match: ^[a-zA-Z0-9_-]+$
func SanitizeCatalogName(name string) string {
	sanitized := strings.ReplaceAll(name, " ", "_")
	sanitized = strings.ReplaceAll(sanitized, ".", "_")

	result := make([]rune, 0, len(sanitized))
	for _, r := range sanitized {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			result = append(result, r)
		}
	}

	return string(result)
}
