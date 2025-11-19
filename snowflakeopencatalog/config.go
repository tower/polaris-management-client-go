package snowflakeopencatalog

// Config contains configuration for connecting to Snowflake OpenCatalog (Polaris)
type Config struct {
	// Account is the Snowflake account identifier used for JWT authentication.
	// Format: account-identifier (without .snowflakecomputing.com suffix)
	// Example: "jhhgkvi-tower_sandbox_live" or "myorg-myaccount"
	Account string

	// User is the Snowflake user that has POLARIS_ACCOUNT_ADMIN role.
	// This user must have the public key configured in Snowflake.
	User string

	// PrivateKey is the PEM-encoded RSA private key for JWT authentication.
	// Supports both PKCS#1 and PKCS#8 formats.
	// The corresponding public key must be registered with the Snowflake user.
	PrivateKey string
}
