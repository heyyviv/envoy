package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// This function saves the given command to a log file.
// It is designed to be called by the main Go program.
func LogCommand(command string) {
	// Specify the path to the log file.
	// You can change this to any path you want.
	logFilePath := "command_log.txt"

	// Open the file in append mode. This ensures new commands are added
	// to the end of the file instead of overwriting existing content.
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Log the error if the file cannot be opened.
		log.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	// Get the current timestamp in a specific format.
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Format the log entry with the timestamp and the command.
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, command)

	// Write the formatted entry to the file.
	if _, err := file.WriteString(logEntry); err != nil {
		// Log the error if the writing fails.
		log.Printf("Error writing to log file: %v\n", err)
	}
}
