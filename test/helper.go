package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetTerraformOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: "../terraform/modules/cosmos-mongodb",
		NoColor:      true,
	}
}

func DeployResource(t *testing.T, resourceType, testName string) (string, error) {
	opts := GetTerraformOptions()
	return terraform.ApplyE(t, opts)
}

func DeleteResource(t *testing.T, resourceType, testName string) (string, error) {
	opts := GetTerraformOptions()
	return terraform.DestroyE(t, opts)
}
