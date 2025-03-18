package test

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func logToFile(testName, message string) {
	logDir := filepath.Join("logs")
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating logs directory:", err)
		return
	}

	timeStamp := time.Now().Format("2006-01-02T15_04_05")
	logFile := filepath.Join(logDir, fmt.Sprintf("%s-%s.log", testName, timeStamp))

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", message)); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}
