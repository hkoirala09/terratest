package test

import (
	"log"
	"os"
	"testing"
	"time"
	"strings"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global logger instance
var logger *log.Logger

// Initialize logging for test results
func init() {
	logFile, err := os.OpenFile("test_results.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ERROR: Unable to open log file: %v", err)
	}
	logger = log.New(logFile, "TEST LOG: ", log.LstdFlags|log.Lshortfile)
	logger.Println("========== Terraform Test Suite Started ==========")
}

// Get Terraform options for tests
func GetTerraformOptions(terraformDir string, planFile string) *terraform.Options {
	return &terraform.Options{
		TerraformDir: terraformDir,
		PlanFilePath: planFile,
		NoColor:      true,
	}
}

// Save Terraform Plan and Validate File Exists
func ValidateTerraformPlan(t *testing.T, tfOpts *terraform.Options, planFile string) {
	terraform.InitAndPlan(t, tfOpts)
	terraform.SavePlanFile(t, tfOpts, planFile)

	if _, err := os.Stat(planFile); os.IsNotExist(err) {
		t.Fatalf("ERROR: Terraform plan file was not created")
	} else {
		logger.Println("✅ Terraform plan file created: " + planFile)
	}
}

// Validate a specific Terraform Output
func ValidateTerraformOutput(t *testing.T, tfOpts *terraform.Options, outputName string, expectedValue string) {
	actualValue := terraform.Output(t, tfOpts, outputName)
	logger.Println("Checking Terraform Output: " + outputName)

	assert.Equal(t, expectedValue, actualValue, "ERROR: "+outputName+" does not match expected value")
}

// Validate Public Network Access is Disabled
func ValidatePublicNetworkAccess(t *testing.T, tfOpts *terraform.Options) {
	publicAccess := terraform.OutputBool(t, tfOpts, "public_network_enabled")
	logger.Println("Validating Public Network Access...")

	assert.False(t, publicAccess, "ERROR: Public network access should be disabled")
}

// Validate CosmosDB Encryption Settings
func ValidateEncryptionSettings(t *testing.T, tfOpts *terraform.Options) {
	encryption := terraform.Output(t, tfOpts, "data_encryption")
	logger.Println("Checking if encryption is set to CMK...")

	assert.Equal(t, "CMK", encryption, "ERROR: Encryption must be set to CMK with Azure Key Vault")
}

// Run and Clean Up Terraform Test
func RunTerraformTest(t *testing.T, tfOpts *terraform.Options) {
	logger.Println("Applying Terraform Deployment...")
	terraform.ApplyAndIdempotent(t, tfOpts)

	// Validate Terraform Output Exists
	assert.True(t, terraform.OutputExists(t, tfOpts, "cosmosdb_id"), "ERROR: CosmosDB instance not deployed")

	logger.Println("✅ Deployment validation complete. Destroying resources...")
	defer terraform.Destroy(t, tfOpts)
	logger.Println("========== Terraform Test Completed ==========")
}
