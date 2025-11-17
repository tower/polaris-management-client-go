# Client Generation Guide

This document explains how to regenerate the Polaris Management Go client from the OpenAPI specification.

## Prerequisites

### 1. Install OpenAPI Generator CLI

The client generation requires `openapi-generator-cli` to be installed globally:

```bash
npm install -g @openapitools/openapi-generator-cli
```

Verify installation:

```bash
openapi-generator-cli version
```

### 2. Go Toolchain

Ensure you have Go installed (1.19 or later recommended):

```bash
go version
```

## Regenerating the Client

To regenerate the client from the OpenAPI specification, run:

```bash
go generate ./...
```

This will automatically:

1. **Generate** the client code from `polaris-management-service.yml`
2. **Fix** import paths (replace placeholder paths with correct module path)
3. **Install** test dependencies (`github.com/stretchr/testify/assert`)
4. **Tidy** Go modules (`go mod tidy`)
5. **Run** all tests (`go test ./...`)
6. **Format** the generated code (`go fmt ./...`)

### Alternative: Run Generation Script Directly

You can also run the generation tool directly:

```bash
go run tools/generate.go
```

## What Gets Generated

The generator creates/updates:

- API client files (`api_*.go`)
- Model files (`model_*.go`)
- Client configuration (`client.go`, `configuration.go`)
- Documentation (`docs/`)
- Tests (`test/`)
- `README.md` (auto-generated, do not edit manually)
- `go.mod` and `go.sum`

## Configuration

The generation configuration is defined in `tools/generate.go`:

- **Package name**: `polarismgmt`
- **Go module**: `github.com/tower/polaris-management-go`
- **Generator**: `go` (OpenAPI Generator's Go client generator)
- **Additional properties**:
  - `generateInterfaces=true` - Generate interfaces for API clients
  - `enumClassPrefix=true` - Prefix enum classes
  - `withGoMod=true` - Generate with Go modules support
  - `hideGenerationTimestamp=true` - Don't include timestamps in generated files

## Updating the OpenAPI Specification

1. Update `polaris-management-service.yml` with your changes
2. Run `go generate ./...` to regenerate the client
3. Review the changes with `git diff`
4. Run tests to verify: `go test ./...`
5. Commit both the spec and generated code

## Troubleshooting

### Generator Not Found

If you see:
```
openapi-generator-cli not found in PATH
```

Install the generator CLI:
```bash
npm install -g @openapitools/openapi-generator-cli
```

### Tests Failing

If tests fail during generation, the script will stop and show the error. Fix the issues before proceeding:

1. Review test failures
2. Fix issues in the OpenAPI spec or generated code
3. Run `go generate ./...` again

### Import Path Issues

If you see import errors related to `GIT_USER_ID/GIT_REPO_ID`, ensure the post-processing step in `tools/generate.go` is working correctly. The script should automatically replace these placeholders with `tower/polaris-management-go`.

## Manual Generation (Not Recommended)

If you need to run the generation manually without the automation:

```bash
# Generate the client
openapi-generator-cli generate \
  -p generateInterfaces=true \
  -i polaris-management-service.yml \
  -g go \
  -o . \
  --additional-properties=\
packageName=polarismgmt,\
enumClassPrefix=true,\
disallowAdditionalPropertiesIfNotPresent=false,\
withGoMod=true,\
goModuleName=github.com/tower/polaris-management-go,\
isGoSubmodule=false,\
hideGenerationTimestamp=true

# Fix import paths manually (find/replace in all .go files)
# GIT_USER_ID/GIT_REPO_ID -> tower/polaris-management-go

# Install dependencies
go get github.com/stretchr/testify/assert@latest

# Tidy modules
go mod tidy

# Run tests
go test ./...

# Format code
go fmt ./...
```

**Note**: Using `go generate ./...` is strongly recommended as it handles all these steps automatically.