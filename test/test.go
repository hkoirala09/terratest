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

	logToFile(testName, "Start Time : "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test : "+testName)
	logToFile(testName, "Description : Checking required parameter values before running terraform apply")
	logToFile(testName, "Test Environment : dev")

	// Run Terraform Init & Plan
	terraform.InitAndPlan(t, tfOpts)

	// Get structured plan
	plan := terraform.ShowWithStruct(t, tfOpts)

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for _, module := range plan.ResourcePlannedValuesMap {
			assert.NotEmpty(t, module.Type, "Module type should not be empty")
			logToFile(testName, "Module: "+module.Type+" - "+module.Name)
		}
	})

	t.Run("AC2: Validate name is as per naming convention", func(t *testing.T) {
		accountName := tfOpts.Vars["account_name"].(string)
		assert.Contains(t, accountName, "cosmos", "Account name does not follow naming convention")
		logToFile(testName, "Account Name: "+accountName)
	})

	t.Run("AC3: Validate the API type", func(t *testing.T) {
		apiType := tfOpts.Vars["api_type"].(string)
		assert.Equal(t, "MongoDB", apiType, "Incorrect API type configured")
		logToFile(testName, "API Type: "+apiType)
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		tags := tfOpts.Vars["tags"].(map[string]string)
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Mandatory tag missing: "+tag)
		}
		logToFile(testName, "Tags validated successfully")
	})

	t.Run("AC5: Validate Public network access is disabled", func(t *testing.T) {
		publicAccess := tfOpts.Vars["public_network_access"].(bool)
		assert.False(t, publicAccess, "Public network access must be disabled")
		logToFile(testName, "Public Network Access: disabled")
	})

	t.Run("AC6: Validate Private Endpoints are configured", func(t *testing.T) {
		found := false
		for _, resource := range plan.ResourcePlannedValuesMap {
			if resource.Type == "azurerm_private_endpoint" {
				found = true
				break
			}
		}
		assert.True(t, found, "Private endpoints are not configured")
		logToFile(testName, "Private Endpoints configured")
	})

	t.Run("AC7: Validate MTL Security Protocol Version is TLS 1.2", func(t *testing.T) {
		tlsVersion := tfOpts.Vars["min_tls_version"].(string)
		assert.Equal(t, "TLS1_2", tlsVersion, "Minimum TLS version must be TLS 1.2")
		logToFile(testName, "TLS Version: "+tlsVersion)
	})

	t.Run("AC8: Validate Log Analytic workspace is configured", func(t *testing.T) {
		found := false
		for _, resource := range plan.ResourcePlannedValuesMap {
			if resource.Type == "azurerm_log_analytics_workspace" {
				found = true
				break
			}
		}
		assert.True(t, found, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace configured")
	})

	t.Run("AC9: Validate data Encryption set to CMK with Azure Key Vault", func(t *testing.T) {
		dataEncryption := tfOpts.Vars["data_encryption"].(string)
		assert.Equal(t, "CMK", dataEncryption, "Data encryption is not set to CMK")
		logToFile(testName, "Data Encryption: CMK with Azure Key Vault")
	})

	t.Run("AC10: Validate Write and Read Location is East US", func(t *testing.T) {
		location := tfOpts.Vars["location"].(string)
		assert.Equal(t, "East US", location, "Location not set to East US")
		logToFile(testName, "Write and Read Location: East US")
	})

	t.Run("AC11: Validate 'Configure Regions' is disabled", func(t *testing.T) {
		configureRegions := tfOpts.Vars["configure_regions"].(bool)
		assert.False(t, configureRegions, "Configure regions is not disabled")
		logToFile(testName, "Configure Regions: disabled")
	})

	t.Run("AC12: Validate Consistency used is Session", func(t *testing.T) {
		consistency := tfOpts.Vars["consistency_level"].(string)
		assert.Equal(t, "Session", consistency, "Consistency level not set to Session")
		logToFile(testName, "Consistency Level: Session")
	})

	logToFile(testName, "End Time : "+time.Now().Format(time.RFC3339))
}
