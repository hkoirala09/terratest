package test

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// logToFile creates structured logs with timestamps
func logToFile(testName, message string) {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("%s-%s.log", testName, time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	entry := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
	f.WriteString(entry)
}

// AssertEqualLogs logs and asserts equality
func AssertEqualLogs(t TestingT, expected, actual interface{}, message, testName string) {
	if expected != actual {
		logToFile(testName, "FAIL: "+message)
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	} else {
		logToFile(testName, "PASS: "+message)
	}
}

type TestingT interface {
	Errorf(format string, args ...interface{})
}
