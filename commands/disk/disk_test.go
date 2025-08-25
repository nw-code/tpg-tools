package disk_test

import (
	"os"
	"testing"

	"github.com/nw-code/tpg-tools/commands/disk"
)

func TestDiskFree(t *testing.T) {
	t.Run("test from file", func(t *testing.T) {
		in, err := os.Open("testdata/df-h.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer in.Close()
		fs, err := disk.NewFileSystem("/dev/mapper/ubuntu--vg-root", in)
		if err != nil {
			t.Fatal(err)
		}
		got := fs.Usage()
		want := 93

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
