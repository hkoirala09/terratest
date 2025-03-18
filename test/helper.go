package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// GetTerraformOptions retrieves terraform options
func GetTerraformOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: "../terraform/modules/cosmos-mongodb",
		NoColor:      true,
	}
}

// DeployResource executes terraform apply
func DeployResource(t *testing.T, resourceType, testName string) bool {
	tfOpts := GetTerraformOptions()
	logToFile(testName, "Terraform Apply Started: "+resourceType)

	if _, err := terraform.InitAndApplyE(t, tfOpts); err != nil {
		logToFile(testName, "Terraform Apply Error: "+err.Error())
		return false
	}

	logToFile(testName, "Terraform Apply Completed: "+resourceType)
	return true
}

// DeleteResource executes terraform destroy
func DeleteResource(t *testing.T, resourceType, testName string) {
	tfOpts := GetTerraformOptions()
	logToFile(testName, "Terraform Destroy Started: "+resourceType)

	if err := terraform.DestroyE(t, tfOpts); err != nil {
		logToFile(testName, "Terraform Destroy Error: "+err.Error())
	} else {
		logToFile(testName, "Terraform Destroy Completed: "+resourceType)
	}
}
