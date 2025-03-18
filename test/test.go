package test

import (
	"testing"
	"time"
	"fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdfDeploymentPreCheck(t *testing.T) {
	tfOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		NoColor:      true,
	})

	testName := "TestAdfDeploymentPreCheck"
	logToFile(testName, "Start Time: "+time.Now().Format(time.RFC3339))
	logToFile(testName, "Test: "+testName)
	logToFile(testName, "Description: Validating parameters before terraform apply")
	logToFile(testName, "Test Environment: dev")

	plan := terraform.InitAndPlanAndShowWithStruct(t, tfOpts)

	// AC1: Validate Module and Provider versions
	t.Run("AC1: Validate Module and Provider versions", func(t *testing.T) {
		for _, provider := range plan.RawPlan.ProviderFormats {
			assert.NotEmpty(t, provider, "Provider format version is missing")
			logToFile(testName, "Provider format version: "+provider)
		}
	})

	// AC2: Validate naming convention
	t.Run("AC2: Validate naming convention", func(t *testing.T) {
		actualName := plan.RawPlan.Variables["account_name"].Value.(string)
		expectedPrefix := "adf-eus-d"
		assert.Contains(t, actualName, expectedPrefix, "Name doesn't match convention")
		logToFile(testName, "Account Name: "+actualName)
	})

	// AC3: Validate API type
	t.Run("AC3: Validate API type", func(t *testing.T) {
		apiType := plan.RawPlan.Variables["api_type"].Value.(string)
		assert.Equal(t, "MongoDB", apiType, "API type is incorrect")
		logToFile(testName, "API Type: "+apiType)
	})

	// AC4: Validate mandatory tags
	t.Run("AC4: Validate mandatory tags", func(t *testing.T) {
		tags := plan.RawPlan.Variables["tags"].Value.(map[string]interface{})
		mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
		for _, tag := range mandatoryTags {
			_, exists := tags[tag]
			assert.True(t, exists, "Mandatory tag missing: "+tag)
		}
		logToFile(testName, "Mandatory tags verified")
	})

	// AC5: Validate public network access disabled
	t.Run("AC5: Validate Public Network Access", func(t *testing.T) {
		publicAccess := plan.RawPlan.Variables["public_network_access"].Value.(bool)
		assert.False(t, publicAccess, "Public network access should be disabled")
		logToFile(testName, fmt.Sprintf("Public Network Access: %v", publicAccess))
	})

	// AC6: Validate private endpoints configured
	t.Run("AC6: Validate Private Endpoints", func(t *testing.T) {
		foundPE := false
		for _, mod := range plan.ResourcePlannedValuesMap {
			if mod.Type == "azurerm_private_endpoint" {
				foundPE = true
				break
			}
		}
		assert.True(t, foundPE, "Private endpoints not configured")
		logToFile(testName, "Private endpoints configured")
	})

	// AC7: Validate TLS 1.2
	t.Run("AC7: Validate TLS version", func(t *testing.T) {
		tlsVersion := plan.RawPlan.Variables["min_tls_version"].Value.(string)
		assert.Equal(t, "TLS1_2", tlsVersion, "Minimum TLS should be 1.2")
		logToFile(testName, "TLS Version: "+tlsVersion)
	})

	// AC8: Validate Log Analytics workspace
	t.Run("AC8: Validate Log Analytics", func(t *testing.T) {
		foundLA := false
		for _, mod := range plan.ResourcePlannedValuesMap {
			if mod.Type == "azurerm_log_analytics_workspace" {
				foundLA = true
				break
			}
		}
		assert.True(t, foundLA, "Log Analytics workspace not configured")
		logToFile(testName, "Log Analytics workspace configured")
	})

	// AC9: Validate Data Encryption CMK
	t.Run("AC9: Validate CMK Encryption", func(t *testing.T) {
		encryptionType := plan.RawPlan.Variables["data_encryption"].Value.(string)
		assert.Equal(t, "CMK", encryptionType, "Encryption type must be CMK")
		logToFile(testName, "Data Encryption: "+encryptionType)
	})

	// AC10: Validate location East US
	t.Run("AC10: Validate Location", func(t *testing.T) {
		location := plan.RawPlan.Variables["location"].Value.(string)
		assert.Equal(t, "East US", location, "Incorrect location configured")
		logToFile(testName, "Location: "+location)
	})

	// AC11: Validate configure regions disabled
	t.Run("AC11: Validate configure regions disabled", func(t *testing.T) {
		configRegions := plan.RawPlan.Variables["configure_regions"].Value.(bool)
		assert.False(t, configRegions, "'Configure regions' should be disabled")
		logToFile(testName, "Configure Regions disabled")
	})

	// AC12: Validate consistency Session
	t.Run("AC12: Validate Consistency Level", func(t *testing.T) {
		consistency := plan.RawPlan.Variables["consistency_level"].Value.(string)
		assert.Equal(t, "Session", consistency, "Consistency level should be 'Session'")
		logToFile(testName, "Consistency Level: "+consistency)
	})

	logToFile(testName, "End Time: "+time.Now().Format(time.RFC3339))
}

// Logging utility
func logToFile(testName, message string) {
	logfile := fmt.Sprintf("%s-%s.log", testName, time.Now().Format("2006-01-02_15-04-05"))
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		panic(err)
	}
}