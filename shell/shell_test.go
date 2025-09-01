package shell_test

import (
	"bytes"
	"io"
	"os"
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

func TestNewSession(t *testing.T) {
	want := shell.Session{
		DryRun:     false,
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		Transcript: io.Discard,
	}
	got := shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != *got {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestSession_Run(t *testing.T) {
	in := bytes.NewBufferString("echo foo\n")
	out := &bytes.Buffer{}
	errs := &bytes.Buffer{}

	want := ":> echo foo\n:> Exiting...\n"
	session := shell.NewSession(in, out, errs)
	session.DryRun = true
	session.Run()

	if !cmp.Equal(out.String(), want) {
		t.Errorf(cmp.Diff(out.String(), want))
	}
}
