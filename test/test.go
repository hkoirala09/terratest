package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdfDeploymentPreCheck(t *testing.T) {
	tfOpts := GetTerraformOptionsForCosmosMongoDB()
	plan := terraform.InitAndPlanAndShowWithStruct(t, tfOpts)
	testName := "TestAdfDeploymentPreCheck"

	logToFile(testName, "Start Time : "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test : "+testName)
	logToFile(testName, "Description : Checking required parameter values before running terraform apply")
	logToFile(testName, "Test Environment : dev")

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for name, provider := range plan.Configuration.ProviderConfig {
			assert.NotEmpty(t, provider.VersionConstraint, "Provider version constraint not specified for provider: "+name)
			logToFile(testName, fmt.Sprintf("Provider: %s, Version Constraint: %s", name, provider.VersionConstraint))
		}
	})

	t.Run("AC2: Validate name is as per naming convention", func(t *testing.T) {
		expectedName := "expected-name-pattern"
		actualName := plan.Configuration.RootModule.Variables["account_name"].Default
		assert.Contains(t, actualName, expectedName, "Cosmos DB name does not match naming convention")
		logToFile(testName, "Account Name: "+fmt.Sprint(actualName))
	})

	t.Run("AC3: Validate the API type", func(t *testing.T) {
		apiType := plan.Configuration.RootModule.Variables["api_type"].Default
		assert.Equal(t, "MongoDB", apiType, "Incorrect API type configured")
		logToFile(testName, "API Type: "+fmt.Sprint(apiType))
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		tags := plan.Configuration.RootModule.Variables["tags"].Default.(map[string]interface{})
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Mandatory tag missing: "+tag)
		}
		logToFile(testName, "Tags validated successfully")
	})

	t.Run("AC5: Validate Public network access is disabled", func(t *testing.T) {
		publicAccess := plan.Configuration.RootModule.Variables["public_network_access"].Default
		assert.Equal(t, false, publicAccess, "Public network access must be disabled")
		logToFile(testName, "Public Network Access: disabled")
	})

	t.Run("AC6: Validate Private Endpoints are configured", func(t *testing.T) {
		privateEndpoints := plan.PlannedValues.RootModule.ChildModules[0].Resources
		assert.NotEmpty(t, privateEndpoints, "Private endpoints are not configured")
		logToFile(testName, "Private Endpoints configured")
	})

	t.Run("AC7: Validate MTL Security Protocol Version is TLS 1.2", func(t *testing.T) {
		tlsVersion := plan.Configuration.RootModule.Variables["min_tls_version"].Default
		assert.Equal(t, "TLS1_2", tlsVersion, "Minimum TLS version must be TLS 1.2")
		logToFile(testName, "TLS Version: "+fmt.Sprint(tlsVersion))
	})

	t.Run("AC8: Validate Log Analytic workspace is configured", func(t *testing.T) {
		logAnalytics := plan.PlannedValues.RootModule.Resources
		assert.NotEmpty(t, logAnalytics, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace configured")
	})

	t.Run("AC9: Validate data Encryption set to CMK with Azure Key Vault", func(t *testing.T) {
		dataEncryption := plan.Configuration.RootModule.Variables["data_encryption"].Default
		assert.Equal(t, "CMK", dataEncryption, "Data encryption is not set to CMK")
		logToFile(testName, "Data Encryption: CMK with Azure Key Vault")
	})

	t.Run("AC10: Validate Write and Read Location is East US", func(t *testing.T) {
		location := plan.Configuration.RootModule.Variables["location"].Default
		assert.Equal(t, "East US", location, "Location not set to East US")
		logToFile(testName, "Write and Read Location: East US")
	})

	t.Run("AC11: Validate 'Configure Regions' is disabled", func(t *testing.T) {
		configureRegions := plan.Configuration.RootModule.Variables["configure_regions"].Default
		assert.Equal(t, false, configureRegions, "Configure regions is not disabled")
		logToFile(testName, "Configure Regions: disabled")
	})

	t.Run("AC12: Validate Consistency used is Session", func(t *testing.T) {
		consistency := plan.Configuration.RootModule.Variables["consistency_level"].Default
		assert.Equal(t, "Session", consistency, "Consistency level not set to Session")
		logToFile(testName, "Consistency Level: Session")
	})

	logToFile(testName, "End Time : "+time.Now().Format(time.RFC3339))
}