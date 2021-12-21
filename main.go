package main

import (
	"encoding/json"
	"fmt" // https://pkg.go.dev/fmt
	"io/ioutil"
	"os" // https://pkg.go.dev/os

	hello "github.com/docktermj/go-hello-world-module"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func main() {

	// Set fileUserName from file.

	fileUserName := "unknown"

	fileName := os.Getenv("XYZZY_FILE")
	if fileName != "" {

		jsonFile, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// Parse JSON.

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
		fileUserName = result["key"].(string)

	}

	// Print output.

	fmt.Println(hello.Hello())
	fmt.Printf("Hello, world! from %s (main) version %s-%s\n", programName, buildVersion, buildIteration)
	fmt.Printf("Hello, %s! from environment variable.\n", os.Getenv("XYZZY_NAME"))
	fmt.Printf("Hello, %s! from file.\n", fileUserName)
}
