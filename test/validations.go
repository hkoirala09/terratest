package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func ValidateModuleAndProviderVersions(t *testing.T, plan *terraform.PlanStruct, testName string) {
	version := plan.RawPlan.ProviderFormatVersion
	assert.NotEmpty(t, version, "Module and provider version missing")
	logToFile(testName, "Module and Provider Version: "+version)
}

func ValidateNamingConvention(t *testing.T, plan *terraform.PlanStruct, testName string) {
	expectedPattern := "expected-name-pattern"
	actual := plan.RawPlan.Variables["account_name"].Value
	assert.Contains(t, actual, expectedPattern)
	logToFile(testName, "Naming convention verified: "+actual.(string))
}

func ValidateAPIType(t *testing.T, plan *terraform.PlanStruct, testName string) {
	apiType := plan.RawPlan.Variables["api_type"].Value
	assert.Equal(t, "MongoDB", apiType, "Incorrect API Type")
	logToFile(testName, "API Type validated: "+apiType.(string))
}

func ValidateMandatoryTags(t *testing.T, plan *terraform.PlanStruct, testName string) {
	tags := plan.RawPlan.Variables["tags"].Value.(map[string]interface{})
	requiredTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
	for _, tag := range requiredTags {
		assert.Contains(t, tags, tag, "Mandatory tag missing: "+tag)
	}
	logToFile(testName, "Mandatory tags validated successfully")
}

func ValidatePublicNetworkDisabled(t *testing.T, plan *terraform.PlanStruct, testName string) {
	publicAccess := plan.RawPlan.Variables["public_network_access"].Value
	assert.False(t, publicAccess.(bool), "Public network access must be disabled")
	logToFile(testName, "Public network access validated as disabled")
}

func ValidatePrivateEndpoints(t *testing.T, plan *terraform.PlanStruct, testName string) {
	endpoints := plan.ResourcePlannedValuesMap["private_endpoint"].AttributeValues
	assert.NotEmpty(t, endpoints, "No private endpoints found")
	logToFile(testName, "Private endpoints configuration validated")
}

func ValidateMTLSecurityProtocol(t *testing.T, plan *terraform.PlanStruct, testName string) {
	tlsVersion := plan.RawPlan.Variables["min_tls_version"].Value
	assert.Equal(t, "TLS1_2", tlsVersion, "TLS version should be TLS1_2")
	logToFile(testName, "TLS version validated as TLS1_2")
}

func ValidateLogAnalyticsWorkspace(t *testing.T, plan *terraform.PlanStruct, testName string) {
	workspace := plan.ResourcePlannedValuesMap["log_analytics_workspace"].AttributeValues
	assert.NotEmpty(t, workspace, "Log Analytics workspace not configured")
	logToFile(testName, "Log Analytics workspace validated successfully")
}

func ValidateDataEncryptionCMK(t *testing.T, plan *terraform.PlanStruct, testName string) {
	encryption := plan.RawPlan.Variables["encryption_key"].Value
	assert.NotEmpty(t, encryption, "CMK encryption not configured")
	logToFile(testName, "Data encryption with CMK validated")
}

func ValidateWriteReadLocation(t *testing.T, plan *terraform.PlanStruct, testName string) {
	writeLoc := plan.RawPlan.Variables["write_location"].Value
	readLoc := plan.RawPlan.Variables["read_location"].Value
	assert.Equal(t, "East US", writeLoc)
	assert.Equal(t, "East US", readLoc)
	logToFile(testName, "Write and read location validated as East US")
}

func ValidateConfigureRegionsDisabled(t *testing.T, plan *terraform.PlanStruct, testName string) {
	configRegions := plan.RawPlan.Variables["configure_regions"].Value
	assert.False(t, configRegions.(bool), "Configure regions must be disabled")
	logToFile(testName, "Configure regions validated as disabled")
}

func ValidateConsistencyLevel(t *testing.T, plan *terraform.PlanStruct, testName string) {
	consistency := plan.RawPlan.Variables["consistency_level"].Value
	assert.Equal(t, "Session", consistency, "Consistency level must be Session")
	logToFile(testName, "Consistency level validated as Session")
}
