//go:build tools

// Package main provides a tool to generate the Polaris Management Go client from OpenAPI spec.
//
// This tool wraps the openapi-generator-cli command with the specific configuration
// needed for this project.
//
// Usage:
//
//	go generate ./...
//
// Or run directly:
//
//	go run tools/generate.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	packageName   = "polarismgmt"
	goModuleName  = "github.com/tower/polaris-management-go"
	specFile      = "polaris-management-service.yml"
	generator     = "go"
	oldImportPath = "GIT_USER_ID/GIT_REPO_ID"
	newImportPath = "tower/polaris-management-go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Get the project root directory (parent of tools/)
	projectRoot, err := getProjectRoot()
	if err != nil {
		return fmt.Errorf("failed to get project root: %w", err)
	}

	specPath := filepath.Join(projectRoot, specFile)
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		return fmt.Errorf("OpenAPI spec file not found: %s", specPath)
	}

	// Check if openapi-generator-cli is available
	if _, err := exec.LookPath("openapi-generator-cli"); err != nil {
		return fmt.Errorf("openapi-generator-cli not found in PATH. Please install it first:\n" +
			"  npm install -g @openapitools/openapi-generator-cli")
	}

	fmt.Println("Generating Polaris Management Go client...")
	fmt.Printf("  Spec: %s\n", specFile)
	fmt.Printf("  Package: %s\n", packageName)
	fmt.Printf("  Module: %s\n", goModuleName)
	fmt.Println()

	// Build the openapi-generator-cli command
	cmd := buildGenerateCommand(projectRoot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the generator
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run openapi-generator-cli: %w", err)
	}

	fmt.Println("\nGeneration complete!")

	// Post-process: Replace tower/polaris-management-go with the correct path
	fmt.Println("\nPost-processing generated files...")
	if err := replaceImportPaths(projectRoot); err != nil {
		return fmt.Errorf("failed to replace import paths: %w", err)
	}

	fmt.Println("✓ Import paths updated")

	// Get testify dependency
	fmt.Println("\nInstalling test dependencies...")
	getCmd := exec.Command("go", "get", "github.com/stretchr/testify/assert@latest")
	getCmd.Dir = projectRoot
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr

	if err := getCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go get: %w", err)
	}

	fmt.Println("✓ Test dependencies installed")

	// Tidy up go.mod and go.sum
	fmt.Println("\nTidying go modules...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = projectRoot
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr

	if err := tidyCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %w", err)
	}

	fmt.Println("✓ Modules tidied")

	// Run tests
	fmt.Println("\nRunning tests...")
	testCmd := exec.Command("go", "test", "./...")
	testCmd.Dir = projectRoot
	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr

	if err := testCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go test: %w", err)
	}

	fmt.Println("✓ Tests passed")

	// Format the generated code
	fmt.Println("\nFormatting generated code...")
	fmtCmd := exec.Command("go", "fmt", "./...")
	fmtCmd.Dir = projectRoot
	fmtCmd.Stdout = os.Stdout
	fmtCmd.Stderr = os.Stderr

	if err := fmtCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go fmt: %w", err)
	}

	fmt.Println("✓ Code formatted")
	fmt.Println("\nDone!")

	return nil
}

func getProjectRoot() (string, error) {
	// Get the directory of this script
	ex, err := os.Executable()
	if err != nil {
		// Fallback: use current working directory
		return os.Getwd()
	}

	// tools/generate.go -> project root is parent directory
	toolsDir := filepath.Dir(ex)
	projectRoot := filepath.Dir(toolsDir)

	// If running with 'go run', the executable is in a temp directory
	// In that case, use the current working directory
	if strings.Contains(toolsDir, "go-build") {
		return os.Getwd()
	}

	return projectRoot, nil
}

func buildGenerateCommand(outputDir string) *exec.Cmd {
	additionalProps := strings.Join([]string{
		"packageName=" + packageName,
		"enumClassPrefix=true",
		"disallowAdditionalPropertiesIfNotPresent=false",
		"withGoMod=true",
		"goModuleName=" + goModuleName,
		"isGoSubmodule=false",
		"hideGenerationTimestamp=true",
	}, ",")

	args := []string{
		"generate",
		"-p", "generateInterfaces=true",
		"-i", specFile,
		"-g", generator,
		"-o", ".",
		"--additional-properties=" + additionalProps,
	}

	cmd := exec.Command("openapi-generator-cli", args...)
	cmd.Dir = outputDir

	return cmd
}

func replaceImportPaths(projectRoot string) error {
	// Find all .go files in the project
	var goFiles []string
	err := filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip vendor, .git, and other directories
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		// Process any file other than tools/generate.go
		if !strings.HasSuffix(info.Name(), "generate.go") {
			goFiles = append(goFiles, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	// Replace old import path with new one in each file
	for _, file := range goFiles {
		if err := replaceInFile(file, oldImportPath, newImportPath); err != nil {
			return fmt.Errorf("failed to process %s: %w", file, err)
		}
	}

	return nil
}

func replaceInFile(filePath, oldStr, newStr string) error {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Check if replacement is needed
	if !strings.Contains(string(content), oldStr) {
		return nil
	}

	// Replace all occurrences
	newContent := strings.ReplaceAll(string(content), oldStr, newStr)

	// Write back
	return os.WriteFile(filePath, []byte(newContent), 0644)
}
