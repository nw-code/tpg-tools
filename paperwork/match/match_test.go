package match_test

import (
	"bytes"
	"testing"

	"github.com/nw-code/tpg-tools/paperwork/match"
)

func TestMatch(t *testing.T) {
	t.Run("single match", func(t *testing.T) {
		src := bytes.NewBufferString("line 1\nhello world!\nline 3")
		dst := new(bytes.Buffer)
		searchString := "hello"

		matcher, err := match.NewMatcher(
			match.WithInput(src),
			match.WithOutput(dst),
		)

		if err != nil {
			t.Fatal(err)
		}

		matcher.Find(searchString)
		got := dst.String()
		want := "hello world!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("multiple match", func(t *testing.T) {
		src := bytes.NewBufferString("line 1\nhello world!\nline 3")
		dst := new(bytes.Buffer)
		searchString := "line"

		matcher, err := match.NewMatcher(
			match.WithInput(src),
			match.WithOutput(dst),
		)

		if err != nil {
			t.Fatal(err)
		}

		matcher.Find(searchString)
		got := dst.String()
		want := "line 1\nline 3\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
