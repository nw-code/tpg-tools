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
	Stdin          io.Reader
	Stdout, Stderr io.Writer
}

func NewSession(in io.Reader, out, err io.Writer) *Session {
	return &Session{
		Stdin:  in,
		Stdout: out,
		Stderr: err,
	}
}

func (s *Session) Run() {
	scn := bufio.NewScanner(s.Stdin)
	fmt.Fprint(s.Stdout, ":> ")
	for scn.Scan() {
		fmt.Fprint(s.Stdout, ":> ")
		cmd, err := CmdFromString(scn.Text())
		if err != nil {
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stdout, out, err)
		}
		fmt.Fprintf(s.Stdout, "%s", out)
		fmt.Fprint(s.Stdout, ":> ")
	}
	fmt.Fprintln(s.Stdout, "Exiting...")
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
	session.Run()
	return 0
}
