package disk_test

import (
	"os"
	"testing"

	"github.com/nw-code/tpg-tools/commands/disk"
)

func TestDiskUsage(t *testing.T) {
	t.Run("test from file", func(t *testing.T) {
		data, err := os.ReadFile("testdata/df-h.txt")
		if err != nil {
			t.Fatal(err)
		}
		got, err := disk.ParseDf(string(data), "/dev/mapper/ubuntu--vg-root")
		if err != nil {
			t.Fatal(err)
		}
		want := 93

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
