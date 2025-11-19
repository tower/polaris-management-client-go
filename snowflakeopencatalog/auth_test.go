package snowflakeopencatalog

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateTestRSAKey generates a test RSA private key for testing
func generateTestRSAKey(t *testing.T) (*rsa.PrivateKey, string) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Encode to PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return privateKey, string(privateKeyPEM)
}

// mockOAuthServer creates a test server that mimics Snowflake's OAuth token endpoint
func mockOAuthServer(t *testing.T, expectedJWT string, returnUnauthorized bool) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify this is a POST request to the token endpoint
		if r.Method != http.MethodPost {
			t.Logf("Unexpected method: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check the path
		if !strings.HasSuffix(r.URL.Path, "/polaris/api/catalog/v1/oauth/tokens") {
			t.Logf("Unexpected path: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		// Parse form data
		if err := r.ParseForm(); err != nil {
			t.Logf("Failed to parse form: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Verify grant_type
		if grantType := r.FormValue("grant_type"); grantType != "client_credentials" {
			t.Logf("Invalid grant_type: %s", grantType)
			http.Error(w, "Invalid grant_type", http.StatusBadRequest)
			return
		}

		// Verify scope
		if scope := r.FormValue("scope"); scope != PolarisAccountAdminScope {
			t.Logf("Invalid scope: %s", scope)
			http.Error(w, "Invalid scope", http.StatusBadRequest)
			return
		}

		// Verify JWT token if expectedJWT is provided
		clientSecret := r.FormValue("client_secret")
		if expectedJWT != "" && clientSecret != expectedJWT {
			t.Logf("JWT mismatch. Expected prefix: %s..., Got: %s...",
				truncate(expectedJWT, 20), truncate(clientSecret, 20))
			http.Error(w, "Invalid JWT", http.StatusUnauthorized)
			return
		}

		// Return unauthorized if requested
		if returnUnauthorized {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error":             "invalid_client",
				"error_description": "Invalid JWT signature",
			})
			return
		}

		// Return successful token response
		response := TokenResponse{
			AccessToken: "test-access-token-12345",
			TokenType:   "Bearer",
			ExpiresIn:   3300, // 55 minutes
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Logf("Failed to encode response: %v", err)
		}
	}))
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

func testConfigWithKey(account, user, privateKeyPEM string) *Config {
	return &Config{
		Account:    account,
		User:       user,
		PrivateKey: privateKeyPEM,
	}
}

// TestNewAuthClient_Success tests successful creation of auth client
func TestNewAuthClient_Success(t *testing.T) {
	ctx := context.Background()
	_, privateKeyPEM := generateTestRSAKey(t)

	cfg := testConfigWithKey("test-account.snowflakecomputing.com", "test_user", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)

	require.NoError(t, err)
	require.NotNil(t, client)
	assert.NotNil(t, client.privateKey)
	assert.NotEmpty(t, client.publicKeyFingerprint)
	assert.True(t, strings.HasPrefix(client.publicKeyFingerprint, "SHA256:"))
}

// TestNewAuthClient_MissingConfig tests error handling for missing configuration
func TestNewAuthClient_MissingConfig(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name        string
		cfg         *Config
		expectedErr error
	}{
		{
			name: "missing account",
			cfg: &Config{
				Account:    "",
				User:       "test_user",
				PrivateKey: "test-key",
			},
			expectedErr: ErrAccountNotConfigured,
		},
		{
			name: "missing user",
			cfg: &Config{
				Account:    "test-account",
				User:       "",
				PrivateKey: "test-key",
			},
			expectedErr: ErrUserNotConfigured,
		},
		{
			name: "missing private key",
			cfg: &Config{
				Account:    "test-account",
				User:       "test_user",
				PrivateKey: "",
			},
			expectedErr: ErrPrivateKeyNotConfigured,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewAuthClient(ctx, tc.cfg)

			assert.Nil(t, client)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

// TestNewAuthClient_InvalidPrivateKey tests error handling for invalid private key
func TestNewAuthClient_InvalidPrivateKey(t *testing.T) {
	ctx := context.Background()

	cfg := testConfigWithKey("test-account", "test_user", "not-a-valid-key")

	client, err := NewAuthClient(ctx, cfg)

	assert.Nil(t, client)
	assert.ErrorIs(t, err, ErrInvalidPrivateKey)
}

// TestGenerateJWT tests JWT generation
func TestGenerateJWT(t *testing.T) {
	ctx := context.Background()
	privateKey, privateKeyPEM := generateTestRSAKey(t)

	cfg := testConfigWithKey("jhhgkvi-tower_sandbox_live.snowflakecomputing.com", "tower_sandbox_live", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)
	require.NoError(t, err)

	jwtToken, err := client.GenerateJWT(ctx)

	require.NoError(t, err)
	assert.NotEmpty(t, jwtToken)

	// Parse and verify the JWT
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &privateKey.PublicKey, nil
	})

	require.NoError(t, err)
	require.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok)

	// Verify claims
	assert.Equal(t, "JHHGKVI-TOWER_SANDBOX_LIVE.TOWER_SANDBOX_LIVE", claims["sub"])
	assert.True(t, strings.HasPrefix(claims["iss"].(string), "JHHGKVI-TOWER_SANDBOX_LIVE.TOWER_SANDBOX_LIVE.SHA256:"))
	assert.NotNil(t, claims["iat"])
	assert.NotNil(t, claims["exp"])
}

// TestExchangeJWTForToken_Success tests successful token exchange
func TestExchangeJWTForToken_Success(t *testing.T) {
	ctx := context.Background()
	_, privateKeyPEM := generateTestRSAKey(t)

	// Create mock OAuth server
	server := mockOAuthServer(t, "", false)
	defer server.Close()

	// Use server URL directly (it includes http://)
	cfg := testConfigWithKey(server.URL, "test_user", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)
	require.NoError(t, err)

	// Generate JWT
	jwtToken, err := client.GenerateJWT(ctx)
	require.NoError(t, err)

	// Exchange for access token using test server URL
	testOAuthURL := server.URL + "/polaris/api/catalog/v1/oauth/tokens"
	tokenResp, err := client.ExchangeJWTForTokenWithURL(ctx, jwtToken, testOAuthURL)

	require.NoError(t, err)
	require.NotNil(t, tokenResp)
	assert.Equal(t, "test-access-token-12345", tokenResp.AccessToken)
	assert.Equal(t, "Bearer", tokenResp.TokenType)
	assert.Equal(t, 3300, tokenResp.ExpiresIn)
}

// TestExchangeJWTForToken_Unauthorized tests handling of unauthorized response
func TestExchangeJWTForToken_Unauthorized(t *testing.T) {
	ctx := context.Background()
	_, privateKeyPEM := generateTestRSAKey(t)

	// Create mock OAuth server that returns unauthorized
	server := mockOAuthServer(t, "", true)
	defer server.Close()

	cfg := testConfigWithKey(server.URL, "test_user", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)
	require.NoError(t, err)

	jwtToken, err := client.GenerateJWT(ctx)
	require.NoError(t, err)

	// Exchange should fail with unauthorized
	testOAuthURL := server.URL + "/polaris/api/catalog/v1/oauth/tokens"
	tokenResp, err := client.ExchangeJWTForTokenWithURL(ctx, jwtToken, testOAuthURL)

	assert.Nil(t, tokenResp)
	assert.ErrorIs(t, err, ErrUnauthorizedClient)
}

// TestGetAccessToken_Caching tests token caching behavior
func TestGetAccessToken_Caching(t *testing.T) {
	ctx := context.Background()
	_, privateKeyPEM := generateTestRSAKey(t)

	server := mockOAuthServer(t, "", false)
	defer server.Close()

	cfg := testConfigWithKey(server.URL, "test_user", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)
	require.NoError(t, err)

	// Set test OAuth URL to use the mock server
	client.testOAuthURL = server.URL + "/polaris/api/catalog/v1/oauth/tokens"

	// First call should fetch a new token
	token1, err := client.GetAccessToken(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, token1)
	assert.Equal(t, "test-access-token-12345", token1)

	// Second call should return cached token (same value)
	token2, err := client.GetAccessToken(ctx)
	require.NoError(t, err)
	assert.Equal(t, token1, token2)

	// Verify the cached token exists
	assert.NotNil(t, client.cachedToken)
	assert.False(t, client.expiresAt.IsZero())
}

// TestGetAccessToken_Refresh tests token refresh when expired
func TestGetAccessToken_Refresh(t *testing.T) {
	ctx := context.Background()
	_, privateKeyPEM := generateTestRSAKey(t)

	server := mockOAuthServer(t, "", false)
	defer server.Close()

	cfg := testConfigWithKey(server.URL, "test_user", privateKeyPEM)

	client, err := NewAuthClient(ctx, cfg)
	require.NoError(t, err)

	// Set test OAuth URL to use the mock server
	client.testOAuthURL = server.URL + "/polaris/api/catalog/v1/oauth/tokens"

	// Get initial token
	token1, err := client.GetAccessToken(ctx)
	require.NoError(t, err)
	assert.Equal(t, "test-access-token-12345", token1)

	// Verify token is cached and expiry is set
	assert.NotNil(t, client.cachedToken)
	originalExpiresAt := client.expiresAt
	assert.False(t, originalExpiresAt.IsZero())

	// Manually expire the token by setting expiresAt far in the past
	// This should trigger a refresh on the next GetAccessToken call
	client.mu.Lock()
	client.expiresAt = time.Now().Add(-1 * time.Hour)
	client.mu.Unlock()

	// Next call should refresh and get a new token
	// (in real scenario it would be different, but our mock always returns the same token)
	token2, err := client.GetAccessToken(ctx)
	require.NoError(t, err)
	assert.Equal(t, "test-access-token-12345", token2) // Same in mock

	// Verify expiresAt was updated to a future time
	assert.True(t, client.expiresAt.After(time.Now()),
		"expiresAt should be updated to a future time after refresh")
	assert.True(t, client.expiresAt.After(originalExpiresAt),
		"expiresAt should be later than the original value")
}

// TestParsePrivateKey_PKCS8 tests parsing PKCS8 format
func TestParsePrivateKey_PKCS8(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Encode as PKCS8
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	require.NoError(t, err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	parsed, err := parsePrivateKey(string(privateKeyPEM))

	require.NoError(t, err)
	assert.Equal(t, privateKey.N, parsed.N)
}

// TestParsePrivateKey_PKCS1 tests parsing PKCS1 format
func TestParsePrivateKey_PKCS1(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Encode as PKCS1
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	parsed, err := parsePrivateKey(string(privateKeyPEM))

	require.NoError(t, err)
	assert.Equal(t, privateKey.N, parsed.N)
}

// TestParsePrivateKey_Invalid tests error handling for invalid keys
func TestParsePrivateKey_Invalid(t *testing.T) {
	testCases := []struct {
		name        string
		key         string
		expectedErr error
	}{
		{
			name:        "not PEM",
			key:         "not-a-pem-key",
			expectedErr: ErrFailedToDecodePEM,
		},
		{
			name: "invalid PEM content",
			key: `-----BEGIN PRIVATE KEY-----
invalid-base64-content
-----END PRIVATE KEY-----`,
			expectedErr: nil, // Will fail at parse stage
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parsePrivateKey(tc.key)
			assert.Error(t, err)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
			}
		})
	}
}

// TestCalculatePublicKeyFingerprint tests fingerprint calculation
func TestCalculatePublicKeyFingerprint(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	fingerprint, err := calculatePublicKeyFingerprint(privateKey)

	require.NoError(t, err)
	assert.NotEmpty(t, fingerprint)
	assert.True(t, strings.HasPrefix(fingerprint, "SHA256:"))
	// Base64 encoded SHA256 should be 44 characters + "SHA256:" prefix
	assert.Equal(t, 51, len(fingerprint))
}

// TestGetOAuthURL tests OAuth URL construction
func TestGetOAuthURL(t *testing.T) {
	_, privateKeyPEM := generateTestRSAKey(t)
	cfg := testConfigWithKey("test-account.snowflakecomputing.com", "test_user", privateKeyPEM)

	client, err := NewAuthClient(context.Background(), cfg)
	require.NoError(t, err)

	url := client.getOAuthURL()

	assert.Equal(t, "https://test-account.snowflakecomputing.com/polaris/api/catalog/v1/oauth/tokens", url)
}

// TestGetManagementBaseURL tests management API base URL construction
func TestGetManagementBaseURL(t *testing.T) {
	_, privateKeyPEM := generateTestRSAKey(t)
	cfg := testConfigWithKey("test-account.snowflakecomputing.com", "test_user", privateKeyPEM)

	client, err := NewAuthClient(context.Background(), cfg)
	require.NoError(t, err)

	url := client.GetManagementBaseURL()

	assert.Equal(t, "https://test-account.snowflakecomputing.com/polaris", url)
}
