# Forge CLI – AI-Powered Terminal Assistant

Forge CLI is a Go-based command-line tool that converts natural language instructions into executable shell commands using the Mistral AI model (via Ollama). It is designed to safely generate and run terminal commands in your current working environment.

Features

Converts human-readable instructions into shell commands.

Detects system context: OS, shell, home directory, username, and editor.

Cleans AI output to produce valid shell commands.

Rewrites paths to ensure commands run relative to the current working directory.

Safety checks to prevent dangerous commands (rm -rf /, unsafe multi-line echo, etc.).

Dry-run and confirmation before executing any command.

Modular code structure for AI interaction, context detection, and command execution.


Getting Started
Requirements

Go 1.21+ installed

Ollama Mistral API running locally (http://localhost:11434/api/generate)

Bash shell (or compatible shell)

Installation

Clone the repository:

git clone <your-repo-url>
cd forge
go mod tidy

Usage

Run the CLI with a natural language instruction:

go run ./cmd/forge "create a folder named forge-polish-test"


Example AI-generated command:

mkdir forge-polish-test


Dry-run confirmation:

Execute this command? [y/N]:


Type y to execute, N to abort.

Multi-line file creation

Forge CLI automatically instructs the AI to use cat << 'EOF' for multi-line files:

go run ./cmd/forge "create a file main.cpp with hello world in C++ and execute it"


Example AI output:

cat << 'EOF' > main.cpp
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
EOF

g++ main.cpp -o main && ./main

Safety Features

Prevents execution of destructive commands:

rm -rf /

Unsafe sudo commands

Multi-line echo for file creation

Prompts for user confirmation before executing AI-generated commands.

Design & Architecture

AI Module (internal/ai): Handles API requests and parses AI responses.

Context Module (internal/context): Detects system OS, shell, and paths.

Executor Module (internal/executor): Validates commands and executes safely.

CLI (cmd/forge): Orchestrates user input, AI interaction, and execution.

Future Improvements

Add streaming of AI output for real-time command preview.

Introduce partial safety checks while AI is generating commands.

Support multiple shells and Windows environments.

Allow AI to handle more complex instructions with step-by-step command generation.

Skills Demonstrated

Go modules, packages, and struct design.

HTTP requests and JSON parsing (encoding/json).

CLI argument parsing and OS-level operations.

Safe shell command execution with os/exec.

Error handling, user interaction, and code modularity.
