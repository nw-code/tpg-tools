package count_test

import (
	"strings"
	"testing"

	"github.com/nw-code/tpg-tools/arguments/count"
)

func TestCount(t *testing.T) {
	t.Run("from stdin", func(t *testing.T) {
		t.Parallel()
		src := strings.NewReader("line1\nline2\nline3")
		counter, err := count.NewCounter(
			count.WithInput(src),
		)

		if err != nil {
			t.Fatal(err)
		}

		got := counter.Lines()
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("ensure no panic for zero args", func(t *testing.T) {
		t.Parallel()
		src := strings.NewReader("1\n2\n3")
		counter, err := count.NewCounter(
			count.WithInput(src),
			count.WithInputFromArgs([]string{}),
		)

		if err != nil {
			t.Fatal(err)
		}

		got := counter.Lines()
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("from file", func(t *testing.T) {
		t.Parallel()
		args := []string{"testdata/three_lines.txt"}
		counter, err := count.NewCounter(
			count.WithInputFromArgs(args),
		)

		if err != nil {
			t.Fatal(err)
		}

		got := counter.Lines()
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
