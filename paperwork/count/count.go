package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func (c *Counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}

	return lines
}

func NewCounter() *Counter {
	return &Counter{
		Input: os.Stdin,
	}
}

func Main() int {
	return NewCounter().Lines()
}
