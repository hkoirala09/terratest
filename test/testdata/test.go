package test

import (
	"testing"
	"time"
	"os"
	"log"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global logger instance
var logger *log.Logger

// Initialize logging to a file
func init() {
	logFile, err := os.OpenFile("test_results.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ERROR: Unable to open log file: %v", err)
	}
	logger = log.New(logFile, "TEST LOG: ", log.LstdFlags|log.Lshortfile)
	logger.Println("========== Terraform Test Suite Started ==========")
}

// Function to Get Terraform Options
func GetTerraformOptions(terraformDir string, planFile string) *terraform.Options {
	return &terraform.Options{
		TerraformDir: terraformDir,
		PlanFilePath: planFile,
		NoColor:      true,
	}
}

// Validate Terraform Plan and File Creation
func TestTerraformPlan(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	logger.Println("Running Terraform Init and Plan...")
	terraform.InitAndPlan(t, tfOpts)
	terraform.SavePlanFile(t, tfOpts, "cosmosdb_plan.out")

	// Ensure plan file exists
	if _, err := os.Stat("cosmosdb_plan.out"); os.IsNotExist(err) {
		t.Fatalf("ERROR: Terraform plan file was not created")
	} else {
		logger.Println("✅ Terraform plan file created: cosmosdb_plan.out")
	}
}

// Validate Public Network Access is Disabled
func TestPublicNetworkAccess(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	publicAccess := terraform.OutputBool(t, tfOpts, "public_network_enabled")
	logger.Println("Validating Public Network Access...")

	assert.False(t, publicAccess, "ERROR: Public network access should be disabled")
}

// Validate Encryption Settings
func TestEncryptionSettings(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	encryption := terraform.Output(t, tfOpts, "data_encryption")
	logger.Println("Checking if encryption is set to CMK...")

	assert.Equal(t, "CMK", encryption, "ERROR: Encryption must be set to CMK with Azure Key Vault")
}

// Validate Terraform Output Values
func TestTerraformOutputs(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	// Validate CosmosDB Account ID
	cosmosdbID := terraform.Output(t, tfOpts, "cosmosdb_id")
	logger.Println("Validating CosmosDB ID Output...")
	assert.NotEmpty(t, cosmosdbID, "ERROR: CosmosDB ID should not be empty")

	// Validate CosmosDB Endpoint
	cosmosdbEndpoint := terraform.Output(t, tfOpts, "cosmosdb_endpoint")
	logger.Println("Validating CosmosDB Endpoint Output...")
	assert.NotEmpty(t, cosmosdbEndpoint, "ERROR: CosmosDB Endpoint should not be empty")
}

// Apply Terraform and Validate Deployment
func TestTerraformApply(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	logger.Println("Applying Terraform deployment...")
	terraform.ApplyAndIdempotent(t, tfOpts)

	// Validate Deployment Success
	assert.True(t, terraform.OutputExists(t, tfOpts, "cosmosdb_id"), "ERROR: CosmosDB instance not deployed")
}

// Destroy Terraform Resources
func TestTerraformDestroy(t *testing.T) {
	tfOpts := GetTerraformOptions("../terraform/modules/cosmosdb-mongodb", "cosmosdb_plan.out")

	logger.Println("Destroying Terraform resources...")
	terraform.Destroy(t, tfOpts)
	logger.Println("✅ Terraform resources successfully destroyed.")
}

// Full End-to-End Test (Plan, Apply, Validate, Destroy)
func TestTerraformFullRun(t *testing.T) {
	t.Run("TestTerraformPlan", TestTerraformPlan)
	t.Run("TestTerraformApply", TestTerraformApply)
	t.Run("TestTerraformOutputs", TestTerraformOutputs)
	t.Run("TestPublicNetworkAccess", TestPublicNetworkAccess)
	t.Run("TestEncryptionSettings", TestEncryptionSettings)
	t.Run("TestTerraformDestroy", TestTerraformDestroy)
	logger.Println("✅ Full Terraform Test Execution Completed.")
}
