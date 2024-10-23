package envparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EnvVar represents a single environment variable
type EnvVar struct {
	Key   string
	Value string
}

// ReadFromEnvFileAndParse reads a file and returns a slice of environment variables
func ReadFromEnvFileAndParse(filename string) ([]EnvVar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var vars []EnvVar
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split on first = sign, we use 2 because we want to split on the first = sign
		// the strings.SplitN function mainly splits on the separator and returns a slice of strings
		// suppose we have a line like this: APP_NAME=SimpleEnvParser
		// then strings.SplitN(line, "=", 2) will return ["APP_NAME", "SimpleEnvParser"]
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Warning: Line %d is not a valid key=value pair, skipping\n", lineNum)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		vars = append(vars, EnvVar{
			Key:   key,
			Value: value,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return vars, nil
}
