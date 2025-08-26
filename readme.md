
Tasks
- [ ] terminal command to start a session
- [ ] save commands in a file (.txt for now)


# Setup 
## Bash user
 - Open your .bashrc file.
 - Ensure the following code is present, replacing ~/path/to/your/envoy with the correct path.

```
# This function is executed before every command.
# It uses 'history 1' to get the last command entered.
function preexec_go() {
  # Get the command from the history.
  # Using history 1 ensures we get the most recent command.
  local command_line=$(history 1 | sed 's/^ *[0-9]* *//')

  # Check if the command line is empty or is the 'go_preexec' command itself to avoid a recursive loop.
  if [[ -n "$command_line" && "$command_line" != "preexec_go" ]]; then
    # Run the Go program and pass the command line as an argument.
    # The path to your Go executable must be correct.
    # We run it in the background with '&' to not block the shell.
    # We also redirect output to /dev/null and use 'disown' to prevent job completion messages.
    /path/to/your/go_watcher "$command_line" >/dev/null 2>&1 & disown
  fi
}

# The 'trap DEBUG' command is the key.
# It executes the specified command (in this case, our function) before every command.
trap 'preexec_go' DEBUG
```

## Zsh users
- Open your .zshrc file.
-  Ensure the following code is present, replacing ~/path/to/your/envoy with the correct path.

```
( /Users/vivekdas/Desktop/projects/envoy/main "$1" &>/dev/null & )
# This function is executed just before a command is executed.
# Zsh provides the command string as the first argument to the preexec hook.
function preexec_go() {
  # Check if the command line is empty or is the 'go_preexec' command itself to avoid a recursive loop.
  if [[ -n "$1" && "$1" != "preexec_go" ]]; then
    # Run the Go program and pass the command line as an argument.
    # The path to your Go executable must be correct.
    # We run it in a subshell with 'disown' to prevent job start notifications.
     ( /Users/vivekdas/Desktop/projects/envoy/main "$1" &>/dev/null & )
  fi
}

# The 'preexec' hook is a special Zsh feature.
# It automatically runs the specified function just before a command is executed.
# We append our function to the list of functions to run.
preexec_functions+=(preexec_go)

```
