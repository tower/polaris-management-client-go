package snowflakeopencatalog

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

const (
	// PolarisAccountAdminScope is the OAuth scope required for catalog management
	PolarisAccountAdminScope = "session:role:POLARIS_ACCOUNT_ADMIN"

	// TokenRefreshBuffer is how long before expiry we should refresh the token
	TokenRefreshBuffer = 5 * time.Minute
)

var logger = logrus.WithField("component", "snowflakeopencatalog")

// TokenResponse represents the OAuth token response from Snowflake
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// AuthClient handles JWT generation and OAuth token exchange for Snowflake OpenCatalog
type AuthClient struct {
	cfg                  *Config
	privateKey           *rsa.PrivateKey
	publicKeyFingerprint string

	// Token caching
	mu          sync.RWMutex
	cachedToken *TokenResponse
	expiresAt   time.Time

	// Optional OAuth URL override
	// If set, this URL will be used instead of the default Snowflake OAuth endpoint
	testOAuthURL string
}

// NewAuthClient creates a new authentication client
func NewAuthClient(ctx context.Context, cfg *Config) (*AuthClient, error) {
	if cfg.Account == "" {
		return nil, ErrAccountNotConfigured
	}

	if cfg.User == "" {
		return nil, ErrUserNotConfigured
	}

	if cfg.PrivateKey == "" {
		return nil, ErrPrivateKeyNotConfigured
	}

	privateKey, err := parsePrivateKey(cfg.PrivateKey)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to parse private key")
		return nil, ErrInvalidPrivateKey
	}

	fingerprint, err := calculatePublicKeyFingerprint(privateKey)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to calculate public key fingerprint")
		return nil, ErrFailedToCalculateFingerprint
	}

	return &AuthClient{
		cfg:                  cfg,
		privateKey:           privateKey,
		publicKeyFingerprint: fingerprint,
	}, nil
}

// parsePrivateKey parses a PEM-encoded RSA private key
func parsePrivateKey(pemKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, ErrFailedToDecodePEM
	}

	// Try PKCS8 format first
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, ErrKeyNotRSA
	}

	// Fall back to PKCS1 format
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// calculatePublicKeyFingerprint calculates the SHA256 fingerprint of the public key
func calculatePublicKeyFingerprint(privateKey *rsa.PrivateKey) (string, error) {
	publicKey := &privateKey.PublicKey

	// Marshal the public key to DER format (PKIX)
	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", ErrFailedToMarshalPublicKey
	}

	// Calculate SHA256 hash
	hash := sha256.Sum256(publicKeyDER)

	// Encode as base64
	fingerprint := base64.StdEncoding.EncodeToString(hash[:])

	return fmt.Sprintf("SHA256:%s", fingerprint), nil
}

// GenerateJWT creates a JWT token signed with the private key
func (c *AuthClient) GenerateJWT(ctx context.Context) (string, error) {
	now := time.Now()

	// Snowflake requires specific claims format
	// Account identifier: Extract first part before .snowflakecomputing.com and convert to uppercase
	// For "jhhgkvi-tower_sandbox_live.snowflakecomputing.com" -> "JHHGKVI-TOWER_SANDBOX_LIVE"
	accountParts := strings.Split(c.cfg.Account, ".")
	accountIdentifier := strings.ToUpper(accountParts[0])
	username := strings.ToUpper(c.cfg.User)

	// Subject: ACCOUNT_IDENTIFIER.USER
	subject := fmt.Sprintf("%s.%s", accountIdentifier, username)

	// Issuer: ACCOUNT_IDENTIFIER.USER.SHA256:public_key_fingerprint
	issuer := fmt.Sprintf("%s.%s", subject, c.publicKeyFingerprint)

	claims := jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(59 * time.Minute)), // Max 1 hour, use 59 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to sign JWT")
		return "", ErrFailedToGenerateJWT
	}

	return signedToken, nil
}

// ExchangeJWTForToken exchanges a JWT for an OAuth access token
func (c *AuthClient) ExchangeJWTForToken(ctx context.Context, jwtToken string) (*TokenResponse, error) {
	return c.ExchangeJWTForTokenWithURL(ctx, jwtToken, c.getOAuthURL())
}

// ExchangeJWTForTokenWithURL exchanges a JWT for an OAuth access token using a custom OAuth URL
func (c *AuthClient) ExchangeJWTForTokenWithURL(ctx context.Context, jwtToken string, oauthURL string) (*TokenResponse, error) {

	// Prepare form data
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_secret", jwtToken)
	formData.Set("scope", PolarisAccountAdminScope)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, oauthURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to create OAuth request")
		return nil, ErrFailedToExchangeToken
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to execute OAuth request")
		return nil, ErrFailedToExchangeToken
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to read OAuth response")
		return nil, ErrFailedToExchangeToken
	}

	if resp.StatusCode != http.StatusOK {
		logger.WithContext(ctx).Errorf("token exchange failed with status %d: %s", resp.StatusCode, string(body))
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorizedClient
		}
		return nil, ErrFailedToExchangeToken
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to parse token response")
		return nil, ErrFailedToExchangeToken
	}

	return &tokenResp, nil
}

// GetAccessToken returns a valid access token, refreshing if necessary
func (c *AuthClient) GetAccessToken(ctx context.Context) (string, error) {
	c.mu.RLock()
	if c.cachedToken != nil && time.Now().Before(c.expiresAt.Add(-TokenRefreshBuffer)) {
		token := c.cachedToken.AccessToken
		c.mu.RUnlock()
		return token, nil
	}
	c.mu.RUnlock()

	// Need to refresh
	c.mu.Lock()
	defer c.mu.Unlock()

	// Double-check after acquiring write lock
	if c.cachedToken != nil && time.Now().Before(c.expiresAt.Add(-TokenRefreshBuffer)) {
		return c.cachedToken.AccessToken, nil
	}

	// Generate JWT
	jwtToken, err := c.GenerateJWT(ctx)
	if err != nil {
		return "", err
	}

	// Exchange for access token
	tokenResp, err := c.ExchangeJWTForToken(ctx, jwtToken)
	if err != nil {
		return "", err
	}

	// Cache the token
	c.cachedToken = tokenResp

	// Snowflake returns expires_in=0, but the access token is valid for 1 hour
	// Use a default expiration of 55 minutes to be safe
	expiresIn := tokenResp.ExpiresIn
	if expiresIn == 0 {
		expiresIn = 55 * 60 // 55 minutes in seconds
	}
	c.expiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)

	logger.WithContext(ctx).Debugf("snowflakeopencatalog: obtained access token, caching for %d seconds", expiresIn)

	return tokenResp.AccessToken, nil
}

// getOAuthURL constructs the OAuth token endpoint URL
func (c *AuthClient) getOAuthURL() string {
	// Use test URL if set (for testing purposes)
	if c.testOAuthURL != "" {
		return c.testOAuthURL
	}
	return fmt.Sprintf("https://%s/polaris/api/catalog/v1/oauth/tokens", c.cfg.Account)
}

// GetManagementBaseURL returns the base URL for the Polaris Management API
func (c *AuthClient) GetManagementBaseURL() string {
	return fmt.Sprintf("https://%s/polaris", c.cfg.Account)
}
