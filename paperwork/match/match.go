package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type matcher struct {
	input  io.Reader
	output io.Writer
}

type option func(*matcher) error

func NewMatcher(opts ...option) (*matcher, error) {
	m := &matcher{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *matcher) Find(searchString string) {
	scanner := bufio.NewScanner(m.input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			fmt.Fprintln(m.output, line)
		}
	}
}

func WithInput(input io.Reader) option {
	return func(m *matcher) error {
		if input == nil {
			return errors.New("nil reader")
		}

		m.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(m *matcher) error {
		if output == nil {
			return errors.New("nil reader")
		}

		m.output = output
		return nil
	}
}

func Main(searchString string) {
	m, err := NewMatcher()
	if err != nil {
		panic(err)
	}
	m.Find(searchString)
}
