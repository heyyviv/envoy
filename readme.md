# **Envoy - Command Logger**

Envoy is a lightweight, background utility that logs your terminal commands. It's designed to be a simple and unobtrusive way to keep a history of your shell usage, which can be useful for debugging, tracking work, or just remembering what you did.

## **Features**

- **Starts and stops on demand**: Only logs commands when you explicitly turn it on.  
- **Saves to a custom file**: You can specify the log file name.  
- **Cross-platform**: Works on both Linux and macOS using bash or zsh.  
- **Minimal overhead**: Runs in the background and has no noticeable impact on shell performance.

## **Installation**

### **Prerequisites**

- **Go**: You must have Go (version 1.20 or newer) installed.  
- **A Shell**: This utility is designed for bash or zsh.

### **Step-by-Step Guide**

1. **Clone the Repository**
   ```bash
   git clone https://github.com/heyyviv/envoy.git
   cd envoy
    ```

2. **Initialize Go Module**

   ```bash
   go mod init envoy
   ```

3. **Build the Executable**

   ```bash
   go build -o ./envoy main.go
   ```

4. **Configure Your Shell**

   * **For Zsh (macOS and Linux):**
     Add the following function to your `~/.zshrc` file:

     ```bash
     function preexec_go() {
       if [[ -n "$1" && "$1" != "preexec_go" ]]; then
         # CHANGE THIS PATH to the location of your 'envoy' executable
         ~/Desktop/projects/envoy/envoy "$1" >/dev/null 2>&1 & disown
       fi
     }
     preexec_functions+=(preexec_go)
     ```

   * **For Bash (Linux):**
     Add the following function to your `~/.bashrc` file:

     ```bash
     function preexec_go() {
       local command_line=$(history 1 | sed 's/^ *[0-9]* *//')
       if [[ -n "$command_line" && "$command_line" != "preexec_go" ]]; then
         # CHANGE THIS PATH to the location of your 'envoy' executable
         ( ~/Desktop/projects/envoy/envoy "$command_line" >/dev/null 2>&1 & ) disown
       fi
     }
     trap 'preexec_go' DEBUG
     ```

5. **Reload Your Shell**

   ```bash
   source ~/.zshrc   # or source ~/.bashrc
   ```

## **Usage**

### **Starting and Stopping the Logger**

* Start logging:

  ```bash
  envoy start mycommands
  ```

  This creates a file named `mycommands.txt` in the same directory as the executable.

* Stop logging:

  ```bash
  envoy stop
  ```

### **Log File Location**

All log files, including a status log file named `envoy_status.log` and a state file named `.envoy_log_status`, are stored in the same directory as the `envoy` executable.

For example:

```bash
~/Desktop/projects/envoy/
```


