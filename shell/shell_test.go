package shell_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/shell"
)

func TestCmdFromString(t *testing.T) {
	shellCmd := "/usr/bin/df"
	shellArgs := []string{"-h"}
	cmd := shell.CmdFromString(shellCmd, shellArgs...)

	if cmd.Path != shellCmd {
		t.Errorf("got %q, want %q", cmd.Path, shellCmd)
	}

	if !cmp.Equal(cmd.Args[1:], shellArgs) {
		t.Errorf(cmp.Diff(cmd.Args[1:], shellArgs))
	}
}
