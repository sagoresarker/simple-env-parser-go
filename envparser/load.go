package envparser

import (
	"fmt"
	"os"
)

// This function reads the .env file, parses it, and sets environment variables
// on the current OS process
//
// N.B: I use declearative function name to make it more readable as it is for learning purpose
func LoadTheEnvFileAndSetEnvVariablesOnCurrentOSProcess(filename string) error {
	vars, err := ReadFromEnvFileAndParse(filename)
	if err != nil {
		return err
	}

	for _, v := range vars {
		if err := os.Setenv(v.Key, v.Value); err != nil {
			return fmt.Errorf("error setting environment variable %s: %w", v.Key, err)
		}
		fmt.Printf("Set %s=%s\n", v.Key, v.Value)
	}

	return nil
}
