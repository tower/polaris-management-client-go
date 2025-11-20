package snowflakeopencatalog

import "errors"

// Sentinel errors for the snowflakeopencatalog package.
// Error messages MUST start with "snowflakeopencatalog:" as per maintainer conventions.
var (
	// Configuration errors
	ErrAccountNotConfigured    = errors.New("snowflakeopencatalog: account not configured")
	ErrUserNotConfigured       = errors.New("snowflakeopencatalog: user not configured")
	ErrPrivateKeyNotConfigured = errors.New("snowflakeopencatalog: private key not configured")
	ErrInvalidPrivateKey       = errors.New("snowflakeopencatalog: invalid private key format")

	// Authentication errors
	ErrFailedToGenerateJWT    = errors.New("snowflakeopencatalog: failed to generate JWT")
	ErrFailedToExchangeToken  = errors.New("snowflakeopencatalog: failed to exchange JWT for access token")
	ErrFailedToGetAccessToken = errors.New("snowflakeopencatalog: failed to get access token")
	ErrUnauthorizedClient     = errors.New("snowflakeopencatalog: unauthorized client")

	// API errors
	ErrFailedToListCatalogs         = errors.New("snowflakeopencatalog: failed to list catalogs")
	ErrFailedToCreateCatalog        = errors.New("snowflakeopencatalog: failed to create catalog")
	ErrFailedToGetCatalog           = errors.New("snowflakeopencatalog: failed to get catalog")
	ErrFailedToDeleteCatalog        = errors.New("snowflakeopencatalog: failed to delete catalog")
	ErrCatalogNotEmpty              = errors.New("snowflakeopencatalog: catalog is not empty")
	ErrFailedToListPrincipals       = errors.New("snowflakeopencatalog: failed to list principals")
	ErrFailedToCreatePrincipal      = errors.New("snowflakeopencatalog: failed to create principal")
	ErrFailedToGetPrincipal         = errors.New("snowflakeopencatalog: failed to get principal")
	ErrFailedToDeletePrincipal      = errors.New("snowflakeopencatalog: failed to delete principal")
	ErrFailedToListPrincipalRoles   = errors.New("snowflakeopencatalog: failed to list principal roles")
	ErrFailedToGetPrincipalRole     = errors.New("snowflakeopencatalog: failed to get principal role")
	ErrFailedToCreatePrincipalRole  = errors.New("snowflakeopencatalog: failed to create principal role")
	ErrFailedToCreateCatalogRole    = errors.New("snowflakeopencatalog: failed to create catalog role")
	ErrFailedToDeleteCatalogRole    = errors.New("snowflakeopencatalog: failed to delete catalog role")
	ErrFailedToAddGrant             = errors.New("snowflakeopencatalog: failed to add grant to catalog role")
	ErrFailedToAssignCatalogRole    = errors.New("snowflakeopencatalog: failed to assign catalog role to principal role")
	ErrFailedToAssignPrincipalRole  = errors.New("snowflakeopencatalog: failed to assign principal role to principal")
	ErrFailedToRotateCredentials    = errors.New("snowflakeopencatalog: failed to rotate credentials")
	ErrCatalogAlreadyExists         = errors.New("snowflakeopencatalog: catalog already exists")
	ErrCatalogNotFound              = errors.New("snowflakeopencatalog: catalog not found")
	ErrPrincipalAlreadyExists       = errors.New("snowflakeopencatalog: principal already exists")
	ErrPrincipalNotFound            = errors.New("snowflakeopencatalog: principal not found")

	// Public key errors
	ErrFailedToCalculateFingerprint = errors.New("snowflakeopencatalog: failed to calculate public key fingerprint")
	ErrFailedToMarshalPublicKey     = errors.New("snowflakeopencatalog: failed to marshal public key")
	ErrKeyNotRSA                    = errors.New("snowflakeopencatalog: key is not RSA")
	ErrFailedToDecodePEM            = errors.New("snowflakeopencatalog: failed to decode PEM block")
)
