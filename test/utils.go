package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Load expected values from JSON file
func LoadExpectedValues(t *testing.T, env string) map[string]interface{} {
	filePath := "testdata/" + env + "/expected_values.json"
	data, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err, "Error reading expected values file")

	var expected map[string]interface{}
	err = json.Unmarshal(data, &expected)
	assert.NoError(t, err, "Error unmarshalling expected values JSON")

	return expected
}

// Helper to validate tags
func ValidateMandatoryTags(t *testing.T, actualTags map[string]interface{}) {
	mandatoryTags := []string{"CreatorID", "ProjectName", "RunID", "WorkspaceName"}
	for _, tag := range mandatoryTags {
		_, exists := actualTags[tag]
		assert.True(t, exists, "Missing mandatory tag: "+tag)
	}
}

// Helper to validate naming convention
func ValidateNamingConvention(t *testing.T, actualName, expectedPattern string) {
	assert.Contains(t, actualName, expectedPattern, "Resource name does not follow naming convention")
}

// Helper to check boolean conditions clearly
func ValidateBooleanValue(t *testing.T, actual, expected bool, message string) {
	assert.Equal(t, expected, actual, message)
}

// Helper to validate string values clearly
func ValidateStringValue(t *testing.T, actual, expected, message string) {
	assert.Equal(t, expected, actual, message)
}
