package greet_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/nw-code/tpg-tools/packages/greet"
)

func TestGreet(t *testing.T) {
	name := "Nick"
	r := strings.NewReader("Nick")
	buf := bytes.Buffer{}
	want := fmt.Sprintf("Hello, %s\n", name)
	greet.PrintTo(r, &buf)
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
