package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	files  []io.Reader
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func (c *counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}

	for _, f := range c.files {
		f.(io.Closer).Close()
	}

	return lines
}

func (c *counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}

	for _, f := range c.files {
		f.(io.Closer).Close()
	}

	return words
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithInput(in io.Reader) option {
	return func(c *counter) error {
		if in == nil {
			return errors.New("nil input reader")
		}
		c.input = in
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}

		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}

		c.input = io.MultiReader(c.files...)

		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func MainLines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0
}

func MainWords() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Words())
	return 0
}
