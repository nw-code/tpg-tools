package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	DryRun         bool
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	Transcript     io.Writer
}

func NewSession(in io.Reader, out, err io.Writer) *Session {
	return &Session{
		DryRun:     false,
		Stdin:      in,
		Stdout:     out,
		Stderr:     err,
		Transcript: io.Discard,
	}
}

func (s *Session) Run() {
	stdout := io.MultiWriter(s.Stdout, s.Transcript)
	stderr := io.MultiWriter(s.Stderr, s.Transcript)
	scn := bufio.NewScanner(s.Stdin)
	fmt.Fprint(stdout, ":> ")
	for scn.Scan() {
		if s.DryRun {
			fmt.Fprintln(stdout, scn.Text())
		} else {
			cmd, err := CmdFromString(scn.Text())
			if err != nil {
				fmt.Fprint(stdout, ":> ")
				continue
			}
			fmt.Fprintln(s.Transcript, scn.Text())
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(stderr, out, err)
			}
			fmt.Fprintf(stdout, "%s", out)
		}
		fmt.Fprint(stdout, ":> ")
	}
	fmt.Fprintln(stdout, "Exiting...")
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil, errors.New("empty input")
	}

	return exec.Command(args[0], args[1:]...), nil
}

func Main() int {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	transcript, err := os.Create("transcript.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer transcript.Close()
	session.Transcript = transcript
	session.Run()
	return 0
}
