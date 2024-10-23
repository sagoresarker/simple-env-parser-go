package main

import (
	"fmt"
	"log"

	"github.com/sagoresarker/simple-env-parser-go/envparser"
	"github.com/sagoresarker/simple-env-parser-go/utils"
)

func main() {

	// load .env file from the current directory
	if err := envparser.LoadTheEnvFileAndSetEnvVariablesOnCurrentOSProcess(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	variables := []string{"APP_NAME", "APP_ENV", "PORT"}

	fmt.Println("\nEnvironment Variables:")
	fmt.Println("---------------------")
	for _, v := range variables {
		value := utils.GetEnvVariableFromCurrentOSProcess(v, "not set")
		fmt.Printf("%s: %s\n", v, value)
	}
}
