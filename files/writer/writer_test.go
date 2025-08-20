package writer_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/nw-code/tpg-tools/files/writer"
)

func TestWriter(t *testing.T) {
	t.Run("validates file contents", func(t *testing.T) {
		path := t.TempDir() + "/write_test.txt"
		want := []byte{1, 2, 3}
		err := writer.WriteToFile(path, want)
		if err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("error accessing contents of file %q: %v", path, err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
	t.Run("returns error for unwritable file", func(t *testing.T) {
		path := "missing-path/write_test.txt"
		err := writer.WriteToFile(path, []byte{})

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("clobbers existing file", func(t *testing.T) {
		path := t.TempDir() + "/write_test.txt"
		err := os.WriteFile(path, []byte("Some test data"), 0600)
		if err != nil {
			t.Fatal(err)
		}
		want := []byte{1, 2, 3}
		err = writer.WriteToFile(path, want)
		if err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("ensures correct permissions on file", func(t *testing.T) {
		path := t.TempDir() + "/write_test.txt"
		err := writer.WriteToFile(path, []byte{})
		if err != nil {
			t.Fatal(err)
		}
		stat, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		got := int(stat.Mode())
		want := 0600

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
