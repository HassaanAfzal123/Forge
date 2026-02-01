package context

import (
	"os"
	"os/user"
	"runtime"
)

type SystemContext struct {
	OS       string
	Shell    string
	HomeDir  string
	Username string
	Editor   string
}

func Detect() SystemContext {
	u, _ := user.Current()
	home, _ := os.UserHomeDir()

	return SystemContext{
		OS:       runtime.GOOS,
		Shell:    os.Getenv("SHELL"),
		HomeDir:  home,
		Username: u.Username,
		Editor:   os.Getenv("EDITOR"),
	}
}
