package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// The path for internal status and error logs.
const statusLogFile = "envoy_status.log"

// LogCommandToFile saves the given command to a specified log file.
// It is designed to be called by the main Go program.
func LogCommandToFile(command, filePath string) {
	// Open the file in append mode. This ensures new commands are added.
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		LogStatus(fmt.Sprintf("ERROR: Could not open command log file %s: %v", filePath, err))
		return
	}
	defer file.Close()

	// Get the current timestamp.
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Format the log entry.
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, command)

	// Write the entry to the file.
	if _, err := file.WriteString(logEntry); err != nil {
		LogStatus(fmt.Sprintf("ERROR: Could not write to log file %s: %v", filePath, err))
	}
}

// LogStatus saves a status or error message to a dedicated log file.
func LogStatus(message string) {
	file, err := os.OpenFile(statusLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// If we can't open the status log, fall back to the standard logger.
		log.Printf("ERROR: Could not open status log file %s: %v", statusLogFile, err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)

	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("ERROR: Could not write to status log file %s: %v", statusLogFile, err)
	}
}
