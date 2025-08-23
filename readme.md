
Tasks
- [] terminal command to start a session
- [] save commands in a file (.txt for now)

# Setup 
## Bash user
 - Open your .bashrc file.
 - Ensure the following code is present, replacing ~/path/to/your/envoy with the correct path.

 ```
# Set a function to call the envoy logger for every command
function log_command() {
    # Check if the last command was empty, if so, do nothing
    if [ -z "$BASH_COMMAND" ]; then
        return
    fi
    # Execute the envoy tool, passing the command as an argument
    ~/path/to/your/envoy "$BASH_COMMAND"
}

# Use PROMPT_COMMAND to execute the function before each new prompt
PROMPT_COMMAND="log_command"
```

## Zsh users
- Open your .zshrc file.
-  Ensure the following code is present, replacing ~/path/to/your/envoy with the correct path.

```
# Add a function to log commands to the Go tool
function log_command() {
    # Check if the command is empty or is the envoy tool itself
    if [ -z "$1" ] || [[ "$1" == *envoy* ]]; then
        return
    fi
    # Execute the envoy tool with the command
    ~/path/to/your/envoy "$1"
}

# Use the preexec hook to execute the function before a command is run
preexec_functions+=("log_command")
```
