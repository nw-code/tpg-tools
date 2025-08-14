package hello_test

import (
	"bytes"
	"testing"

	"github.com/nw-code/tpg-tools/paperwork/hello"
)

func TestHello(t *testing.T) {
	var buf bytes.Buffer
	want := "Hello, world\n"
	//buf := &bytes.Buffer{}
	p := hello.NewPrinter()
	p.Output = &buf
	p.Print()
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
