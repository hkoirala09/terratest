package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestCosmosMongoDBDeploymentPreCheck(t *testing.T) {
	tfOpts := GetTerraformOptions()
	plan := terraform.InitAndPlanAndShowWithStruct(t, tfOpts)
	testName := "TestCosmosMongoDBDeploymentPreCheck"

	logToFile(testName, "Start Time: "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test: "+testName)
	logToFile(testName, "Description: Checking required parameter values before running terraform apply")
	logToFile(testName, "Test Environment: dev")

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		ValidateModuleAndProviderVersions(t, plan, testName)
	})

	t.Run("AC2: Validate the name as per naming convention", func(t *testing.T) {
		ValidateNamingConvention(t, plan, testName)
	})

	t.Run("AC3: Validate API type", func(t *testing.T) {
		ValidateAPIType(t, plan, testName)
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		ValidateMandatoryTags(t, plan, testName)
	})

	t.Run("AC5: Validate Public network access is disabled", func(t *testing.T) {
		ValidatePublicNetworkDisabled(t, plan, testName)
	})

	t.Run("AC6: Validate Private Endpoints are configured", func(t *testing.T) {
		ValidatePrivateEndpoints(t, plan, testName)
	})

	t.Run("AC7: Validate MTL Security Protocol Version is TLS 1.2", func(t *testing.T) {
		ValidateMTLSecurityProtocol(t, plan, testName)
	})

	t.Run("AC8: Validate Log Analytics workspace is configured", func(t *testing.T) {
		ValidateLogAnalyticsWorkspace(t, plan, testName)
	})

	t.Run("AC9: Validate data Encryption with CMK Azure Key Vault", func(t *testing.T) {
		ValidateDataEncryptionCMK(t, plan, testName)
	})

	t.Run("AC10: Validate Write and Read Location is East US", func(t *testing.T) {
		ValidateWriteReadLocation(t, plan, testName)
	})

	t.Run("AC11: Validate Configure Regions is disabled", func(t *testing.T) {
		ValidateConfigureRegionsDisabled(t, plan, testName)
	})

	t.Run("AC12: Validate Consistency used is Session", func(t *testing.T) {
		ValidateConsistencyLevel(t, plan, testName)
	})

	logToFile(testName, "End Time: "+time.Now().Format(time.RFC3339))
}

func TestCosmosMongoDBDeploymentPostCheck(t *testing.T) {
	testName := "TestCosmosMongoDBDeploymentPostCheck"
	logToFile(testName, "Start Time: "+time.Now().Format(time.RFC3339))

	assert.True(t, DeployResource(t, "cosmosdb", testName), "Deployment failed")
	defer DeleteResource(t, "cosmosdb", testName)

	logToFile(testName, "End Time: "+time.Now().Format(time.RFC3339))
}
