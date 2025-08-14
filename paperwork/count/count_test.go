package count_test

import (
	//"bytes"
	"strings"
	"testing"

	"github.com/nw-code/tpg-tools/paperwork/count"
)

func TestCount(t *testing.T) {
	t.Parallel()
	src := strings.NewReader("line1\nline2\nline3")
	//src := bytes.NewBuffer([]byte("line1\nline2\nline3"))
	counter := count.NewCounter()
	counter.Input = src
	got := counter.Lines()
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
