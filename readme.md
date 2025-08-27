# **Envoy \- Command Logger**

Envoy is a lightweight, background utility that logs your terminal commands. It's designed to be a simple and unobtrusive way to keep a history of your shell usage, which can be useful for debugging, tracking work, or just remembering what you did.

## **Features**

* **Starts and stops on demand**: Only logs commands when you explicitly turn it on.  
* **Saves to a custom file**: You can specify the log file name.  
* **Cross-platform**: Works on both Linux and macOS using bash or zsh.  
* **Minimal overhead**: Runs in the background and has no noticeable impact on shell performance.

## **Installation**

### **Prerequisites**

* **Go**: You must have Go (version 1.20 or newer) installed.  
* **A Shell**: This utility is designed for bash or zsh.

### **Step-by-Step Guide**

1. Clone the Repository:  
   Clone the project to your local machine. 
   ``` 
   git clone https://github.com/heyyviv/envoy.git  
   cd envoy
   ```

2. Initialize Go Module:  
   In the envoy directory, run this command to create the Go module.  
   ``` go mod init envoy ```

3. Build the Executable:  
   Build the Go program. This command compiles main.go and logger/logger.go into a single executable named envoy and places it in the correct location.  
   ``` go build \-o ./envoy main.go ```

4. Configure Your Shell:  
   Now, you need to add the shell hook to your profile so that Envoy runs with every command.  
   * For Zsh (macOS and Linux):  
   ```
     Add the following function to your \~/.zshrc file.  
     function preexec\_go() {  
       if \[\[ \-n "$1" && "$1" \!= "preexec\_go" \]\]; then  
         \# CHANGE THIS PATH to the location of your 'envoy' executable  
         \~/Desktop/projects/envoy/envoy "$1" \>/dev/null 2\>&1 & disown  
       fi  
     }  
     preexec\_functions+=(preexec\_go)
     ```

   * For Bash (Linux):  
   ```
     Add the following function to your \~/.bashrc file.  
     function preexec\_go() {  
       local command\_line=$(history 1 | sed 's/^ \*\[0-9\]\* \*//')  
       if \[\[ \-n "$command\_line" && "$command\_line" \!= "preexec\_go" \]\]; then  
         \# CHANGE THIS PATH to the location of your 'envoy' executable  
         ( \~/Desktop/projects/envoy/envoy "$command\_line" \>/dev/null 2\>&1 & ) disown  
       fi  
     }  
     trap 'preexec\_go' DEBUG
     ```

5. Reload Your Shell:  
   After saving the changes to your .zshrc or .bashrc file, run the following command to apply the changes without restarting your terminal.  
   ``` source \~/.zshrc  \# or source \~/.bashrc ```

## **Usage**

### **Starting and Stopping the Logger**

* To start logging, run the envoy start command followed by the desired filename (without the .txt extension).  
  ``` envoy start mycommands ```

  This will create a file named mycommands.txt in the same directory as the executable.  
* To stop logging, simply run the envoy stop command.  
  ``` envoy stop ```

### **Log File Location**

All log files, including a status log file named envoy\_status.log and a state file named .envoy\_log\_status, are stored in the same directory as the envoy executable. In your case, this is \~/Desktop/projects/envoy/.
