package main

import (
	stdctx "context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"errors"

	"forge/internal/ai"
	"forge/internal/context"
	"forge/internal/executor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: forge \"your instruction here\"")
		os.Exit(1)
	}

	instruction := strings.Join(os.Args[1:], " ")
	sys := context.Detect()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	prompt := fmt.Sprintf(`
You are a terminal assistant.
Return ONLY shell commands.
Do not explain anything.
Commands should be run relative to the current working directory unless the user specifies an absolute path.
When creating multi-line files, ALWAYS use cat << 'EOF' ... EOF.
Never use echo for multi-line content.

System:
OS: %s
Shell: %s
Home: %s
Current working directory: %s

User request:
"%s"
`, sys.OS, sys.Shell, sys.HomeDir, cwd, instruction)

	fmt.Println("Sending prompt to Mistral...\n")

	ctx, cancel := signal.NotifyContext(stdctx.Background(), os.Interrupt)
	defer cancel()

	response, err := ai.GenerateStream(ctx, prompt)
if err != nil {
    if errors.Is(err, stdctx.Canceled) {
        fmt.Println("\nCanceled.")
        return
    }

    fmt.Println("Error:", err)
    return
}


	command := cleanCommand(response)
	command = rewritePaths(command, cwd, sys.HomeDir)

	err = executor.ConfirmAndRun(command, true)
	if err != nil {
		fmt.Println("Execution error:", err)
	}
}

// cleanCommand trims Markdown-style code fences from AI output
func cleanCommand(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```bash")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}

// rewritePaths replaces absolute home paths with current working directory
func rewritePaths(command, cwd, home string) string {
	return strings.ReplaceAll(command, home, cwd)
}

