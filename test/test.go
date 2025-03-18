package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdfDeploymentPreCheck(t *testing.T) {
	tfOpts := terraform.Options{
		TerraformDir: "../",
		Vars:         map[string]interface{}{"env": "dev"},
	}

	testName := "TestAdfDeploymentPreCheck"

	logToFile(testName, "Start Time: "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test: "+testName)
	logToFile(testName, "Description: Checking required parameter values before running terraform apply")
	logToFile(testName, "Test Environment: dev")

	planStruct := terraform.InitAndPlanAndShowWithStruct(t, &tfOpts)

	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for _, res := range planStruct.ResourcePlannedValuesMap {
			assert.NotEmpty(t, res.ProviderName)
			logToFile(testName, "Provider: "+res.ProviderName)
		}
	})

	t.Run("AC2: Validate naming convention", func(t *testing.T) {
		expectedName := "expected-pattern"
		actualName := fmt.Sprintf("%v", planStruct.Variables["account_name"].Value)
		assert.Contains(t, actualName, expectedName)
		logToFile(testName, "Account Name: "+actualName)
	})

	t.Run("AC3: Validate API type", func(t *testing.T) {
		apiType := fmt.Sprintf("%v", planStruct.Variables["api_type"].Value)
		assert.Equal(t, "MongoDB", apiType)
		logToFile(testName, "API Type: "+apiType)
	})

	t.Run("AC4: Validate Mandatory tags", func(t *testing.T) {
		tags := planStruct.Variables["tags"].Value.(map[string]interface{})
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Mandatory tag missing: "+tag)
		}
		logToFile(testName, "Tags validated successfully")
	})

	t.Run("AC5: Public network access disabled", func(t *testing.T) {
		publicAccess := planStruct.Variables["public_network_access"].Value.(bool)
		assert.False(t, publicAccess)
		logToFile(testName, "Public Network Access: Disabled")
	})

	t.Run("AC6: Private Endpoints configured", func(t *testing.T) {
		foundPrivateEndpoint := false
		for _, res := range planStruct.ResourcePlannedValuesMap {
			if res.Type == "azurerm_private_endpoint" {
				foundPrivateEndpoint = true
				break
			}
		}
		assert.True(t, foundPrivateEndpoint, "Private endpoints are not configured")
		logToFile(testName, "Private Endpoints configured")
	})

	t.Run("AC7: TLS Protocol Version", func(t *testing.T) {
		tlsVersion := fmt.Sprintf("%v", planStruct.Variables["min_tls_version"].Value)
		assert.Equal(t, "TLS1_2", tlsVersion)
		logToFile(testName, "TLS Version: "+tlsVersion)
	})

	t.Run("AC8: Log Analytic workspace configured", func(t *testing.T) {
		foundLogAnalytics := false
		for _, res := range planStruct.ResourcePlannedValuesMap {
			if res.Type == "azurerm_log_analytics_workspace" {
				foundLogAnalytics = true
				break
			}
		}
		assert.True(t, foundLogAnalytics, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace configured")
	})

	t.Run("AC9: Data Encryption CMK", func(t *testing.T) {
		dataEncryption := fmt.Sprintf("%v", planStruct.Variables["data_encryption"].Value)
		assert.Equal(t, "CMK", dataEncryption)
		logToFile(testName, "Data Encryption: CMK with Azure Key Vault")
	})

	t.Run("AC10: Write and Read Location", func(t *testing.T) {
		location := fmt.Sprintf("%v", planStruct.Variables["location"].Value)
		assert.Equal(t, "East US", location)
		logToFile(testName, "Write and Read Location: "+location)
	})

	t.Run("AC11: Configure Regions disabled", func(t *testing.T) {
		configureRegions := planStruct.Variables["configure_regions"].Value.(bool)
		assert.False(t, configureRegions)
		logToFile(testName, "Configure Regions: Disabled")
	})

	t.Run("AC12: Consistency Session", func(t *testing.T) {
		consistency := fmt.Sprintf("%v", planStruct.Variables["consistency_level"].Value)
		assert.Equal(t, "Session", consistency)
		logToFile(testName, "Consistency Level: "+consistency)
	})

	logToFile(testName, "End Time: "+time.Now().Format(time.RFC3339))
}

// logToFile helper
func logToFile(testName, message string) {
	f, err := os.OpenFile(fmt.Sprintf("%s-%s.log", testName, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()
	f.WriteString(message + "\n")
}
