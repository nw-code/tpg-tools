package shell

import (
	"os/exec"
)

func CmdFromString(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
