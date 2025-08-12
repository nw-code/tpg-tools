package hello_test

import (
	"bytes"
	"testing"

	"github.com/nw-code/tpg-tools/packages/hello"
)

func TestHello(t *testing.T) {
	want := "Hello, world\n"
	buf := bytes.Buffer{}
	hello.PrintTo(&buf)
	got := buf.String()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
