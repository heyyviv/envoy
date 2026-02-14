#!/bin/bash

# Define variables
BINARY_NAME="envoy"
INSTALL_DIR="/usr/local/bin"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go (https://golang.org/dl/) and try again."
    exit 1
fi

echo "Building $BINARY_NAME..."
# Build the binary
go build -o "$BINARY_NAME" main.go
if [ $? -ne 0 ]; then
    echo "Error: Build failed."
    exit 1
fi

echo "Installing $BINARY_NAME to $INSTALL_DIR..."
# Move the binary to the installation directory
# usage of sudo might be required
if [ -w "$INSTALL_DIR" ]; then
    mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
else
    echo "Permission denied. Trying with sudo..."
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
fi

if [ $? -eq 0 ]; then
    echo "Success! $BINARY_NAME has been installed to $INSTALL_DIR."
    echo "You can now run '$BINARY_NAME start <filename>' from any directory."
    echo "Logs will be stored in ~/.envoy/"
else
    echo "Error: Installation failed."
    exit 1
fi
