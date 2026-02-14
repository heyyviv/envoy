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

### **Quick Install**

1. **Clone the Repository**
   ```bash
   git clone https://github.com/heyyviv/envoy.git
   cd envoy
   ```

2. **Run the Install Script**
   ```bash
   chmod +x install.sh
   ./install.sh
   ```
   *Note: You may be asked for your password to install the binary to `/usr/local/bin`.*

### **Manual Installation**

If you prefer to install manually:

1. Initialize Go Module: `go mod init envoy`
2. Build: `go build -o envoy main.go`
3. Move `envoy` to a directory in your PATH (e.g., `/usr/local/bin`).

---

## **Configuration**

Add the shell hook to your profile so that Envoy runs with every command.

* **For Zsh (macOS and Linux):**
  Add to `~/.zshrc`:

  ```bash
  function preexec_go() {
    if [[ -n "$1" && "$1" != "preexec_go" ]]; then
      envoy "$1" >/dev/null 2>&1 & disown
    fi
  }
  preexec_functions+=(preexec_go)
  ```

* **For Bash (Linux):**
  Add to `~/.bashrc`:

  ```bash
  function preexec_go() {
    local command_line=$(history 1 | sed 's/^ *[0-9]* *//')
    if [[ -n "$command_line" && "$command_line" != "preexec_go" ]]; then
      ( envoy "$command_line" >/dev/null 2>&1 & ) disown
    fi
  }
  trap 'preexec_go' DEBUG
  ```

Reload your shell: `source ~/.zshrc` or `source ~/.bashrc`.

## **Usage**

### **Starting and Stopping the Logger**

* **Start logging:**
  ```bash
  envoy start mycommands
  ```
  This starts logging commands to `~/.envoy/mycommands.txt`.

* **Stop logging:**
  ```bash
  envoy stop
  ```

### **Log File Location**
All logs are stored in `~/.envoy/`.


