//go:build integration

package disk_test

import (
	"os/exec"
	"testing"

	"github.com/nw-code/tpg-tools/commands/disk"
)

func TestGetDfOutput(t *testing.T) {
	err := exec.Command("df", "-h").Run()
	if err != nil {
		t.Skipf("unable to run 'df' command: %v", err)
	}

	data, err := disk.GetDfOutput()
	if err != nil {
		t.Fatal(err)
	}

	fs := "/dev/mapper/ubuntu--vg-root"
	use, err := disk.ParseDf(data, fs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Filesystem %q is at %d%% usage", fs, use)
}
