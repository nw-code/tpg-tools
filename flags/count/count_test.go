package count_test

import (
	"os"
	"strings"
	"testing"

	"github.com/nw-code/tpg-tools/flags/count"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count.Main,
	}))
}

func TestCount(t *testing.T) {
	t.Run("count lines from stdin", func(t *testing.T) {
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

	t.Run("count words from stdin", func(t *testing.T) {
		t.Parallel()
		src := strings.NewReader("line 1\nline 2\nline 3")
		counter, err := count.NewCounter(
			count.WithInput(src),
		)

		if err != nil {
			t.Fatal(err)
		}

		got := counter.Words()
		want := 6

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

	t.Run("count lines from file", func(t *testing.T) {
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

	t.Run("count lines from multiple files", func(t *testing.T) {
		t.Parallel()
		args := []string{"testdata/three_lines.txt", "testdata/four_lines.txt"}
		counter, err := count.NewCounter(
			count.WithInputFromArgs(args),
		)

		if err != nil {
			t.Fatal(err)
		}

		got := counter.Lines()
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
