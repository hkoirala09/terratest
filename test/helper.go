package test

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetTerraformOptionsForCosmosMongoDB() *terraform.Options {
	env := os.Getenv("ENV")
	return &terraform.Options{
		TerraformDir: filepath.Join("..", env),
		Vars: map[string]interface{}{
			"account_name":          "cosmos-dev-001",
			"api_type":              "MongoDB",
			"min_tls_version":       "TLS1_2",
			"data_encryption":       "CMK",
			"location":              "East US",
			"public_network_access": false,
			"configure_regions":     false,
			"consistency_level":     "Session",
			"tags": map[string]string{
				"CreatorID":     "test",
				"ProjectName":   "cosmosdb",
				"RunID":         "run001",
				"WorkspaceName": "dev",
			},
		},
	}
}

func logToFile(testName, message string) {
	filename := fmt.Sprintf("%s.log", testName)
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s\n", message))
}
