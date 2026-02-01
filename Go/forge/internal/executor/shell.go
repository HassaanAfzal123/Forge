package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ConfirmAndRun shows the command, asks for confirmation, and executes it
// alreadyPrinted = true if AI streaming already printed the command
func ConfirmAndRun(command string, alreadyPrinted bool) error {
	if !isSafeCommand(command) {
		fmt.Println("Command blocked for safety!")
		return errors.New("unsafe command detected")
	}

	if !alreadyPrinted {
		fmt.Println("Proposed command:\n")
		fmt.Println(command)
	}

	fmt.Print("\nExecute this command? [y/N]: ")

	var response string
	fmt.Scanln(&response)
	if response != "y" && response != "Y" {
		fmt.Println("Aborted.")
		return nil
	}

	fmt.Println("\nExecuting...\n")

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func isSafeCommand(cmd string) bool {
	dangerPatterns := []string{
		"rm -rf /",
		"rm -rf /*",
		"sudo ",
		":(){:|:&};:",
	}

	lower := strings.ToLower(cmd)

	for _, pattern := range dangerPatterns {
		if strings.Contains(lower, pattern) {
			return false
		}
	}

	// Block unsafe multiline echo usage
	if strings.Contains(lower, "echo") && strings.Contains(lower, "\\n") {
		fmt.Println("Unsafe multiline file creation detected. Use cat << EOF instead.")
		return false
	}

	return true
}

