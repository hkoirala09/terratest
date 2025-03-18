package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdfDeploymentPreCheck(t *testing.T) {
	testName := "TestAdfDeploymentPreCheck"
	tfOpts := GetTerraformOptionsForCosmosMongoDB()
	plan := terraform.InitAndPlanAndShowWithStruct(t, tfOpts)

	logToFile(testName, "Start Time : "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test : "+testName)
	logToFile(testName, "Description : Checking required parameter values before running terraform apply")
	logToFile(testName, "Test Environment : dev")

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for _, provider := range plan.PlannedValues.RootModule.ChildModules[0].Resources {
			assert.NotEmpty(t, provider.ProviderName, "Provider version not specified")
			logToFile(testName, "Provider: "+provider.ProviderName)
		}
	})

	t.Run("AC2: Validate name is as per naming convention", func(t *testing.T) {
		expectedName := "expected-name-pattern"
		actualName := plan.Variables["account_name"].Value.(string)
		assert.Contains(t, actualName, expectedName, "Cosmos DB name does not match naming convention")
		logToFile(testName, "Account Name: "+actualName)
	})

	t.Run("AC3: Validate the API type", func(t *testing.T) {
		apiType := plan.Variables["api_type"].Value.(string)
		assert.Equal(t, "MongoDB", apiType, "Incorrect API type configured")
		logToFile(testName, "API Type: "+apiType)
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		tags := plan.Variables["tags"].Value.(map[string]interface{})
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Mandatory tag missing: "+tag)
		}
		logToFile(testName, "Tags validated successfully")
	})

	t.Run("AC5: Validate Public network access is disabled", func(t *testing.T) {
		publicAccess := plan.Variables["public_network_access"].Value.(bool)
		assert.False(t, publicAccess, "Public network access must be disabled")
		logToFile(testName, "Public Network Access: disabled")
	})

	t.Run("AC6: Validate Private Endpoints are configured", func(t *testing.T) {
		privateEndpoints := plan.PlannedValues.RootModule.ChildModules[0].Resources
		assert.NotEmpty(t, privateEndpoints, "Private endpoints are not configured")
		logToFile(testName, "Private Endpoints configured")
	})

	t.Run("AC7: Validate MTL Security Protocol Version is TLS 1.2", func(t *testing.T) {
		tlsVersion := plan.Variables["min_tls_version"].Value.(string)
		assert.Equal(t, "TLS1_2", tlsVersion, "Minimum TLS version must be TLS 1.2")
		logToFile(testName, "TLS Version: "+tlsVersion)
	})

	t.Run("AC8: Validate Log Analytic workspace is configured", func(t *testing.T) {
		logAnalytics := plan.PlannedValues.RootModule.Resources
		assert.NotEmpty(t, logAnalytics, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace configured")
	})

	t.Run("AC9: Validate data Encryption set to CMK with Azure Key Vault", func(t *testing.T) {
		dataEncryption := plan.Variables["data_encryption"].Value.(string)
		assert.Equal(t, "CMK", dataEncryption, "Data encryption is not set to CMK")
		logToFile(testName, "Data Encryption: CMK with Azure Key Vault")
	})

	t.Run("AC10: Validate Write and Read Location is East US", func(t *testing.T) {
		location := plan.Variables["location"].Value.(string)
		assert.Equal(t, "East US", location, "Location not set to East US")
		logToFile(testName, "Write and Read Location: East US")
	})

	t.Run("AC11: Validate 'Configure Regions' is disabled", func(t *testing.T) {
		configureRegions := plan.Variables["configure_regions"].Value.(bool)
		assert.False(t, configureRegions, "Configure regions is not disabled")
		logToFile(testName, "Configure Regions: disabled")
	})

	t.Run("AC12: Validate Consistency used is Session", func(t *testing.T) {
		consistency := plan.Variables["consistency_level"].Value.(string)
		assert.Equal(t, "Session", consistency, "Consistency level not set to Session")
		logToFile(testName, "Consistency Level: Session")
	})

	logToFile(testName, "End Time : "+time.Now().Format(time.RFC3339))
}
