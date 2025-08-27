package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"envoy/logger"
)

// A variable to store the path where log files should be saved.
// If this is an empty string, it will default to the executable's directory.
var logDirectory = ""

const statusFile = ".envoy_log_status"

func main() {
	// Determine the executable's directory.
	execPath, err := os.Executable()
	if err != nil {
		log.Printf("ERROR: Could not get executable path: %v", err)
		return
	}
	execDir := filepath.Dir(execPath)

	// If the logDirectory variable is empty, default it to the executable's directory.
	if logDirectory == "" {
		logDirectory = execDir
	}

	// Check if a command was passed.
	if len(os.Args) < 2 {
		return
	}

	// Reconstruct the full command string.
	commandToLog := strings.Join(os.Args[1:], " ")

	// Check if the command is 'envoy start' or 'envoy stop'.
	commandParts := strings.Split(commandToLog, " ")
	if len(commandParts) >= 2 {
		action := commandParts[0]
		arg := commandParts[1]

		if action == "envoy" {
			if arg == "start" {
				// 'start' command: create or overwrite the status file with the new log file name.
				if len(commandParts) < 3 {
					// Log an error message to the status log.
					logger.LogStatus("ERROR: 'envoy start' requires a filename, e.g., 'envoy start mylog.txt'")
					return
				}
				logFileName := commandParts[2] + ".txt"

				statusFilePath := filepath.Join(logDirectory, statusFile)
				err := os.WriteFile(statusFilePath, []byte(logFileName), 0644)
				if err != nil {
					logger.LogStatus(fmt.Sprintf("ERROR writing status file: %v", err))
				}
				logger.LogStatus(fmt.Sprintf("Started logging to %s", logFileName))

			} else if arg == "stop" {
				// 'stop' command: delete the status file to stop logging.
				statusFilePath := filepath.Join(logDirectory, statusFile)
				err := os.Remove(statusFilePath)
				if err != nil && !os.IsNotExist(err) {
					logger.LogStatus(fmt.Sprintf("ERROR removing status file: %v", err))
				}
				logger.LogStatus("Stopped logging.")
			}
			return // Exit after handling start/stop commands.
		}
	}

	// For all other commands, check if logging is enabled.
	statusFilePath := filepath.Join(logDirectory, statusFile)
	if _, err := os.Stat(statusFilePath); err == nil {
		// Read the log file name from the status file.
		content, err := os.ReadFile(statusFilePath)
		if err != nil {
			logger.LogStatus(fmt.Sprintf("ERROR reading status file: %v", err))
			return
		}
		logFileName := strings.TrimSpace(string(content))
		fullLogPath := filepath.Join(logDirectory, logFileName)

		// Call the logger function to save the command to the specified file.
		logger.LogCommandToFile(commandToLog, fullLogPath)
	}
}
