package utils

import "os"

// Get retrieves an environment variable with a default value
func GetEnvVariableFromCurrentOSProcess(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
