package pipeline

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Reader io.Reader
	Output io.Writer
	Error  error
}

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	_, p.Error = io.Copy(p.Output, p.Reader)
}

func (p *Pipeline) String() (string, error) {
	if p.Error != nil {
		return "", p.Error
	}
	data, err := io.ReadAll(p.Reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FromString(str string) *Pipeline {
	p := New()
	p.Reader = bytes.NewBufferString(str)
	return p
}

func FromFile(pathname string) *Pipeline {
	p := New()
	b, err := os.ReadFile(pathname)
	p.Error = err
	p.Reader = bytes.NewBuffer(b)
	return p
}

func (p *Pipeline) Column(col int) *Pipeline {
	buf := new(bytes.Buffer)
	if p.Error != nil {
		p.Reader = buf
		return p
	}
	scanner := bufio.NewScanner(p.Reader)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		if col < 0 || col > len(cols)-1 {
			p.Error = errors.New("invalid column")
			p.Reader = new(bytes.Buffer)
			return p
		}
		buf.WriteString(cols[col-1] + "\n")
	}
	p.Reader = buf
	return p
}
