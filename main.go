package main

import (
	"fmt"
	"os"
	"strings"

	// We import the logger package to use its functions.
	// The path should be relative to your Go module.
	"envoy/logger"
)

func main() {
	// Check if a command was passed as a command-line argument.
	// os.Args[0] is the program name, so we need at least 2 arguments.
	if len(os.Args) < 2 {
		return
	}

	// Join the command-line arguments to reconstruct the full command string.
	commandToLog := strings.Join(os.Args[1:], " ")
	fmt.Println(commandToLog)

	// Call the function from our imported logger package.
	// The `logger.LogCommand` function will handle writing the command to the file.
	logger.LogCommand(commandToLog)
}
