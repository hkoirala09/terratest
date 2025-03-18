package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// LogToFile logs the given message into a log file
func LogToFile(testName string, message string) {
	currentTime := time.Now().Format("2006-01-02T15_04_05")
	logFileName := fmt.Sprintf("%s-%s.log", testName, currentTime)

	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer f.Close()

	logger := fmt.Sprintf("%s\n", message)
	f.WriteString(logger)
}

// LoadExpectedValues loads expected values from JSON file
func LoadExpectedValues(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var values map[string]interface{}
	if err = json.Unmarshal(data, &values); err != nil {
		return nil, err
	}

	return values, nil
}

// GetTerraformOptions returns Terraform options for the given module
func GetTerraformOptions(t *testing.T, terraformDir string, vars map[string]interface{}) *terraform.Options {
	return &terraform.Options{
		TerraformDir: terraformDir,
		Vars:         vars,
		NoColor:      true,
	}
}
