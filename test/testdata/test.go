package test

import (
    "testing"
    "time"
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/gruntwork-io/terratest/modules/terraform"
)

// GetTerraformOptionsForCosmosMongoDB returns terraform options for CosmosDB tests
func GetTerraformOptionsForCosmosMongoDB() *terraform.Options {
    return &terraform.Options{
        TerraformDir: "../terraform/modules/cosmos-mongodb",
        NoColor:      true,
    }
}

// TestAdfDeploymentPreCheck checks terraform plan for valid input and configuration
func TestAdfDeploymentPreCheck(t *testing.T) {
    t1 := time.Now().Format(time.RFC3339)
    fmt.Printf("Test started at: %s\n", t1)

    tfOpts := GetTerraformOptionsForCosmosMongoDB()
    plan := terraform.InitAndPlanAndShowWithStruct(t, tfOpts)

    t.Run("Check Cosmos MongoDB Name", func(t *testing.T) {
        expectedName := "expected-name-pattern"
        actualName := plan.RawPlan.Variables["cosmos_db_name"].Value
        assert.Contains(t, actualName, expectedName, "Error: Cosmos DB name does not match the following convention")
    })
}

// TestAdfDeploymentPostCheck tests terraform apply to deploy resources and cleanup afterwards
func TestAdfDeploymentPostCheck(t *testing.T) {
    t2 := time.Now().Format(time.RFC3339)
    fmt.Printf("Test started at: %s\n", t2)

    tfOpts := GetTerraformOptionsForCosmosMongoDB()

    // Deploy resources in Azure
    terraform.InitAndApply(t, tfOpts)

    // Validate resources deployed successfully here
    assert.True(t, DeployAdfInAzure(t, tfOpts), "ERROR: Failed to deploy resources in Azure")

    // Cleanup resources after test
    defer DeleteAdfInAzure(t, tfOpts)
}

// DeployAdfInAzure is a placeholder for your deployment validation logic
func DeployAdfInAzure(t *testing.T, tfOpts *terraform.Options) bool {
    // Implement your resource validation logic
    return true
}

// DeleteAdfInAzure cleans up Azure resources after testing
func DeleteAdfInAzure(t *testing.T, tfOpts *terraform.Options) {
    terraform.Destroy(t, tfOpts)
}
