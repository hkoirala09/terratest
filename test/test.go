package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdfDeploymentPreCheck(t *testing.T) {
	testName := "TestAdfDeploymentPreCheck"
	tfOpts := terraform.Options{
		TerraformDir: "../",
		Vars:         map[string]interface{}{"env": "dev"},
	}

	logToFile(testName, "Start Time: "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test: "+testName)
	logToFile(testName, "Description: Checking required parameter values before terraform apply")
	logToFile(testName, "Test Environment: dev")

	planStruct := terraform.InitAndPlanAndShowWithStruct(t, &tfOpts)

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for _, res := range planStruct.ResourcePlannedValuesMap {
			assert.NotEmpty(t, res.ProviderName, "Provider name/version is empty")
			logToFile(testName, "Provider: "+res.ProviderName)
		}
	})

	t.Run("AC2: Validate name is as per naming convention", func(t *testing.T) {
		expectedNamePattern := "cosmos-mongo"
		actualName := planStruct.Variables["account_name"].Value.(string)
		assert.Contains(t, actualName, expectedNamePattern, "Account name does not match naming convention")
		logToFile(testName, "Account Name: "+actualName)
	})

	t.Run("AC3: Validate the API type", func(t *testing.T) {
		apiType := planStruct.Variables["api_type"].Value.(string)
		assert.Equal(t, "MongoDB", apiType, "API type mismatch")
		logToFile(testName, "API Type: "+apiType)
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		tags := planStruct.Variables["tags"].Value.(map[string]interface{})
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Missing mandatory tag: "+tag)
		}
		logToFile(testName, "Mandatory Tags: Validated successfully")
	})

	t.Run("AC5: Validate Public network access is disabled", func(t *testing.T) {
		publicAccess := planStruct.Variables["public_network_access"].Value.(bool)
		assert.False(t, publicAccess, "Public network access must be disabled")
		logToFile(testName, "Public Network Access: Disabled")
	})

	t.Run("AC6: Validate Private Endpoints are configured", func(t *testing.T) {
		privateEndpoints := false
		for _, res := range planStruct.ResourcePlannedValuesMap {
			if res.Type == "azurerm_private_endpoint" {
				privateEndpoints = true
				break
			}
		}
		assert.True(t, privateEndpoints, "Private endpoints not configured")
		logToFile(testName, "Private Endpoints: Configured")
	})

	t.Run("AC7: Validate MTL Security Protocol Version is TLS 1.2", func(t *testing.T) {
		tlsVersion := planStruct.Variables["min_tls_version"].Value.(string)
		assert.Equal(t, "TLS1_2", tlsVersion, "MTL security protocol mismatch")
		logToFile(testName, "TLS Version: "+tlsVersion)
	})

	t.Run("AC8: Validate Log Analytics workspace is configured", func(t *testing.T) {
		logAnalyticsConfigured := false
		for _, res := range planStruct.ResourcePlannedValuesMap {
			if res.Type == "azurerm_log_analytics_workspace" {
				logAnalyticsConfigured = true
				break
			}
		}
		assert.True(t, logAnalyticsConfigured, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace: Configured")
	})

	t.Run("AC9: Validate data Encryption set to CMK with Azure Key Vault", func(t *testing.T) {
		dataEncryption := planStruct.Variables["data_encryption"].Value.(string)
		assert.Equal(t, "CMK", dataEncryption, "Data encryption is not set to CMK")
		logToFile(testName, "Data Encryption: CMK with Azure Key Vault")
	})

	t.Run("AC10: Validate Write and Read Location is East US", func(t *testing.T) {
		location := planStruct.Variables["location"].Value.(string)
		assert.Equal(t, "East US", location, "Location mismatch")
		logToFile(testName, "Write and Read Location: "+location)
	})

	t.Run("AC11: Validate 'Configure Regions' is disabled", func(t *testing.T) {
		configureRegions := planStruct.Variables["configure_regions"].Value.(bool)
		assert.False(t, configureRegions, "Configure regions should be disabled")
		logToFile(testName, "Configure Regions: Disabled")
	})

	t.Run("AC12: Validate Consistency used is Session", func(t *testing.T) {
		consistency := planStruct.Variables["consistency_level"].Value.(string)
		assert.Equal(t, "Session", consistency, "Consistency level mismatch")
		logToFile(testName, "Consistency Level: "+consistency)
	})

	logToFile(testName, "End Time: "+time.Now().Format(time.RFC3339))
}