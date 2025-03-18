package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetTerraformOptions(env string) *terraform.Options {
	return &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"env": env,
		},
		NoColor: true,
	}
}

// Helper function for detailed logging
func LogTest(testName string, message string) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Unable to create log directory: %v", err)
	}

	logFileName := fmt.Sprintf("%s-%s.log", testName, time.Now().Format("2006-01-02_15-04-05"))
	logPath := filepath.Join(logDir, logFileName)

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(message)
}