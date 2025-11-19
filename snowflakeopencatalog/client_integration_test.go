// Integration tests for Snowflake OpenCatalog (Polaris) client.
//
// These tests verify the client works correctly with a real Snowflake OpenCatalog instance.
// They can be used as a sanity/connection test to verify:
//   - Private key is correctly formatted and registered with Snowflake user
//   - Account identifier is correct
//   - User has POLARIS_ACCOUNT_ADMIN role
//   - JWT generation and OAuth token exchange work end-to-end
//
// The tests are automatically skipped if credentials are not provided via environment variables.
package snowflakeopencatalog

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Integration tests for Snowflake OpenCatalog (Polaris) client.
// These tests verify the client can successfully authenticate and interact with a real Polaris instance.
//
// To run these tests, set the following environment variables:
//
//	SNOWFLAKE_OC_ACCOUNT     - Snowflake account identifier (e.g., "myorg-myaccount.snowflakecomputing.com")
//	SNOWFLAKE_OC_USER        - Snowflake user with POLARIS_ACCOUNT_ADMIN role
//	SNOWFLAKE_OC_PRIVATE_KEY - PEM-encoded RSA private key (the public key must be registered with the user)
//
// Example:
//
//	export SNOWFLAKE_OC_ACCOUNT="myorg-myaccount.snowflakecomputing.com"
//	export SNOWFLAKE_OC_USER="my_polaris_user"
//	export SNOWFLAKE_OC_PRIVATE_KEY="$(cat /path/to/private_key.pem)"
//	go test -v -run TestClientIntegration
//
// These tests are automatically skipped if credentials are not configured.

func TestClientIntegration(t *testing.T) {
	account := os.Getenv("SNOWFLAKE_OC_ACCOUNT")
	user := os.Getenv("SNOWFLAKE_OC_USER")
	privateKey := os.Getenv("SNOWFLAKE_OC_PRIVATE_KEY")

	if account == "" || user == "" || privateKey == "" {
		t.Skip("Skipping integration test: Set SNOWFLAKE_OC_ACCOUNT, SNOWFLAKE_OC_USER, and SNOWFLAKE_OC_PRIVATE_KEY to run")
	}

	cfg := &Config{
		Account:    account,
		User:       user,
		PrivateKey: privateKey,
	}

	ctx := context.Background()

	t.Run("CreateClient", func(t *testing.T) {
		client, err := NewClient(ctx, cfg)
		require.NoError(t, err)
		require.NotNil(t, client)
	})

	t.Run("ListCatalogs", func(t *testing.T) {
		client, err := NewClient(ctx, cfg)
		require.NoError(t, err)

		resp, err := client.ListCatalogs(ctx)
		require.NoError(t, err)
		require.NotNil(t, resp)

		catalogs := resp.GetCatalogs()
		t.Logf("Found %d catalogs", len(catalogs))
		for _, cat := range catalogs {
			t.Logf("  - %s", cat.GetName())
		}
	})

	t.Run("ListPrincipals", func(t *testing.T) {
		client, err := NewClient(ctx, cfg)
		require.NoError(t, err)

		resp, err := client.ListPrincipals(ctx)
		require.NoError(t, err)
		require.NotNil(t, resp)

		principals := resp.GetPrincipals()
		t.Logf("Found %d principals", len(principals))
		for _, p := range principals {
			t.Logf("  - %s", p.GetName())
		}
	})

	t.Run("ListPrincipalRoles", func(t *testing.T) {
		client, err := NewClient(ctx, cfg)
		require.NoError(t, err)

		resp, err := client.ListPrincipalRoles(ctx)
		require.NoError(t, err)
		require.NotNil(t, resp)

		roles := resp.GetRoles()
		t.Logf("Found %d principal roles", len(roles))
		for _, r := range roles {
			t.Logf("  - %s", r.GetName())
		}
	})
}

// TestAuthClientIntegration tests the authentication flow with a real Snowflake OpenCatalog instance.
// This is useful for verifying JWT generation and OAuth token exchange work correctly.
func TestAuthClientIntegration(t *testing.T) {
	account := os.Getenv("SNOWFLAKE_OC_ACCOUNT")
	user := os.Getenv("SNOWFLAKE_OC_USER")
	privateKey := os.Getenv("SNOWFLAKE_OC_PRIVATE_KEY")

	if account == "" || user == "" || privateKey == "" {
		t.Skip("Skipping integration test: Set SNOWFLAKE_OC_ACCOUNT, SNOWFLAKE_OC_USER, and SNOWFLAKE_OC_PRIVATE_KEY to run")
	}

	cfg := &Config{
		Account:    account,
		User:       user,
		PrivateKey: privateKey,
	}

	ctx := context.Background()

	t.Run("GenerateJWT", func(t *testing.T) {
		authClient, err := NewAuthClient(ctx, cfg)
		require.NoError(t, err)

		jwtToken, err := authClient.GenerateJWT(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, jwtToken)
		t.Logf("Generated JWT: %s...", jwtToken[:50])
	})

	t.Run("ExchangeJWTForToken", func(t *testing.T) {
		authClient, err := NewAuthClient(ctx, cfg)
		require.NoError(t, err)

		jwtToken, err := authClient.GenerateJWT(ctx)
		require.NoError(t, err)

		tokenResp, err := authClient.ExchangeJWTForToken(ctx, jwtToken)
		require.NoError(t, err)
		require.NotNil(t, tokenResp)
		require.NotEmpty(t, tokenResp.AccessToken)
		require.Contains(t, []string{"Bearer", "bearer"}, tokenResp.TokenType)
		// Note: Snowflake returns expires_in=0, expiration is in the JWT itself
		t.Logf("Access token received (expires_in=%d)", tokenResp.ExpiresIn)
	})

	t.Run("GetAccessToken", func(t *testing.T) {
		authClient, err := NewAuthClient(ctx, cfg)
		require.NoError(t, err)

		// First call should generate a new token
		token1, err := authClient.GetAccessToken(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, token1)
		t.Logf("First token: %s...", token1[:50])

		// Second call should use cached token
		token2, err := authClient.GetAccessToken(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, token2)
		t.Logf("Second token: %s...", token2[:50])

		// Tokens should be the same (cached)
		require.Equal(t, token1, token2, "Should return cached token")
	})
}

func TestSanitizeCatalogName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-catalog", "my-catalog"},
		{"my_catalog", "my_catalog"},
		{"My Catalog", "My_Catalog"},
		{"catalog.name", "catalog_name"},
		{"catalog@123", "catalog123"},
		{"test-catalog_v2", "test-catalog_v2"},
		{"hello world!", "hello_world"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := SanitizeCatalogName(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
