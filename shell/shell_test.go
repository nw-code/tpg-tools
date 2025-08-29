package shell_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/shell"
)

func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	input := "/usr/bin/df -h"
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/usr/bin/df", "-h"}

	if !cmp.Equal(cmd.Args, want) {
		t.Errorf(cmp.Diff(cmd.Args, want))
	}
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	input := ""
	_, err := shell.CmdFromString(input)
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}
