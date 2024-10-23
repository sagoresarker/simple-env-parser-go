package envparser

import (
	"os"
	"testing"
)

// TestLoadTheEnvFileAndSetEnvVariablesOnCurrentOSProcess tests the function that loads
// the environment variables from a file and sets them on the current OS process
func TestLoadTheEnvFileAndSetEnvVariablesOnCurrentOSProcess(t *testing.T) {
	tempFile, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the temp .env file
	envContent := "APP_NAME=TestApp\nPORT=3000\n# This is a comment\nEMPTY_VALUE=\n"
	if _, err := tempFile.WriteString(envContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	if err := LoadTheEnvFileAndSetEnvVariablesOnCurrentOSProcess(tempFile.Name()); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Test if the environment variables are set correctly
	tests := []struct {
		key           string
		expectedValue string
	}{
		{"APP_NAME", "TestApp"},
		{"PORT", "3000"},
		{"EMPTY_VALUE", ""},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			value := os.Getenv(tt.key)
			if value != tt.expectedValue {
				t.Errorf("Expected %s for key %s, got %s", tt.expectedValue, tt.key, value)
			}
		})
	}
}
