# Snowflake OpenCatalog (Polaris) Client

This package provides a Go client for interacting with Snowflake OpenCatalog (Polaris), implementing JWT authentication and OAuth token exchange for accessing the Polaris Management API.

## Overview

Snowflake OpenCatalog (internally called Polaris) is an Apache Iceberg catalog management service. This package provides:

- **JWT Authentication**: RS256 JWT generation with RSA key pair authentication
- **OAuth Token Exchange**: Client credentials flow to exchange JWT for access tokens
- **Token Caching**: Automatic token caching and refresh with mutex-protected storage
- **Management API Client**: Full wrapper around Polaris Management API operations

## Configuration

The package uses a standalone `Config` struct with no external dependencies:

```go
type Config struct {
    Account    string // Snowflake account identifier (e.g., "myorg-myaccount")
    User       string // Snowflake user with POLARIS_ACCOUNT_ADMIN role
    PrivateKey string // PEM-encoded RSA private key (PKCS#1 or PKCS#8)
}
```

### Example Usage

```go
package main

import (
    "context"
    "github.com/tower/tower-services/pkg/snowflakeopencatalog"
)

func main() {
    cfg := &snowflakeopencatalog.Config{
        Account:    "myorg-myaccount.snowflakecomputing.com",
        User:       "my_polaris_user",
        PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\n...",
    }

    ctx := context.Background()

    // Create client
    client, err := snowflakeopencatalog.NewClient(ctx, cfg)
    if err != nil {
        panic(err)
    }

    // List catalogs
    catalogs, err := client.ListCatalogs(ctx)
    if err != nil {
        panic(err)
    }

    // Use the client...
}
```

## Authentication Flow

1. **JWT Generation**:
   - Creates RS256-signed JWT with Snowflake-specific claims
   - Subject: `ACCOUNT_IDENTIFIER.USER`
   - Issuer: `ACCOUNT_IDENTIFIER.USER.SHA256:public_key_fingerprint`
   - Expiration: 59 minutes (max 1 hour)

2. **Token Exchange**:
   - Exchanges JWT for OAuth access token via `client_credentials` grant
   - Scope: `session:role:POLARIS_ACCOUNT_ADMIN`
   - Access token valid for ~1 hour

3. **Token Caching**:
   - Caches tokens with 5-minute refresh buffer
   - Thread-safe with RWMutex protection
   - Automatic refresh when expired

## API Operations

The client provides methods for all Polaris Management API operations:

### Catalog Operations
- `CreateCatalog()` - Create a new catalog
- `GetCatalog()` - Retrieve catalog by name
- `ListCatalogs()` - List all catalogs
- `DeleteCatalog()` - Delete a catalog

### Principal (Service Account) Operations
- `CreatePrincipal()` - Create a new service principal
- `GetPrincipal()` - Retrieve principal by name
- `ListPrincipals()` - List all principals
- `DeletePrincipal()` - Delete a principal
- `RotatePrincipalCredentials()` - Rotate OAuth credentials

### Principal Role Operations
- `CreatePrincipalRole()` - Create a new principal role
- `GetPrincipalRole()` - Retrieve principal role by name
- `ListPrincipalRoles()` - List all principal roles
- `AssignPrincipalRole()` - Assign role to principal

### Catalog Role Operations
- `CreateCatalogRole()` - Create a new catalog role
- `AddGrantToCatalogRole()` - Add privilege grant to catalog role
- `DeleteCatalogRole()` - Delete a catalog role
- `AssignCatalogRoleToPrincipalRole()` - Assign catalog role to principal role

## Testing

The package includes comprehensive test coverage:

- **Unit Tests** (`auth_test.go`):
  - JWT generation and verification
  - Token exchange and caching
  - Private key parsing (PKCS#1/PKCS#8)
  - Public key fingerprint calculation

- **Integration Tests** (`client_integration_test.go`):
  - Real API interactions with Snowflake OpenCatalog
  - Can be used as **sanity/connection tests** to verify:
    - Private key is correctly formatted and registered
    - Account identifier is correct
    - User has `POLARIS_ACCOUNT_ADMIN` role
    - JWT generation and OAuth token exchange work end-to-end

### Running Tests

```bash
# Unit tests only (no credentials required)
go test -v -run "^Test[^I]"

# Run all tests including integration (requires credentials)
export SNOWFLAKE_OC_ACCOUNT="myorg-myaccount.snowflakecomputing.com"
export SNOWFLAKE_OC_USER="my_polaris_user"
export SNOWFLAKE_OC_PRIVATE_KEY="$(cat /path/to/private_key.pem)"
go test -v
```

### Connection/Sanity Test

Use the integration tests to verify your Snowflake OpenCatalog setup:

```bash
# Set credentials
export SNOWFLAKE_OC_ACCOUNT="myorg-myaccount.snowflakecomputing.com"
export SNOWFLAKE_OC_USER="my_polaris_user"
export SNOWFLAKE_OC_PRIVATE_KEY="$(cat /path/to/private_key.pem)"

# Run integration tests
go test -v -run TestClientIntegration

# Or test just authentication
go test -v -run TestAuthClientIntegration
```

**Successful output:**
```
=== RUN   TestClientIntegration
=== RUN   TestClientIntegration/CreateClient
=== RUN   TestClientIntegration/ListCatalogs
    Found 5 catalogs
      - my_catalog_1
      - my_catalog_2
      ...
--- PASS: TestClientIntegration (2.14s)
```

**Common issues:**

- **401 Unauthorized**: Public key not registered or incorrect format
- **Connection errors**: Account identifier incorrect or network issues
- **403 Forbidden**: User doesn't have `POLARIS_ACCOUNT_ADMIN` role

## Interface for Testing

The package provides a `PolarisClient` interface to enable mocking in tests:

```go
type PolarisClient interface {
    CreateCatalog(ctx, req) (*Catalog, error)
    DeleteCatalog(ctx, name) error
    CreatePrincipal(ctx, req) (*PrincipalWithCredentials, error)
    // ... all API methods
}
```


## Private Key Setup

### Generate RSA Key Pair

```bash
# Generate private key (PKCS#1)
openssl genrsa -out private_key.pem 2048

# Extract public key
openssl rsa -in private_key.pem -pubout -out public_key.pem

# Or generate PKCS#8 format
openssl genpkey -algorithm RSA -out private_key_pkcs8.pem -pkeyopt rsa_keygen_bits:2048
```

### Register Public Key in Snowflake

```sql
ALTER USER my_polaris_user SET RSA_PUBLIC_KEY='MIIBIjANBgkqhki...';
GRANT ROLE POLARIS_ACCOUNT_ADMIN TO USER my_polaris_user;
```

## Dependencies

- `github.com/golang-jwt/jwt/v5` - JWT generation and signing
- `github.com/tower/polaris-management-go` - Generated Polaris API client
- `golang.org/x/oauth2` - OAuth2 token source interface
- Standard library: `crypto/*`, `encoding/*`, `net/http`

## Open Source Readiness

This package is designed to be extracted to a separate open-source repository (`github.com/tower/polaris-management-go`):

- **Zero Tower-specific logic**: Generic Snowflake OpenCatalog client
- **No external config dependencies**: Standalone `Config` struct
- **Comprehensive tests**: Both unit and integration test coverage
- **Clear documentation**: Usage examples and API reference
- **Standard interfaces**: Enables easy mocking and testing

The package can be moved to open source without any changes.
