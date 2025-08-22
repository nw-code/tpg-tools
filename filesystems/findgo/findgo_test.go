package findgo_test

import (
	"strings"
	"testing"
	"testing/fstest"
	"archive/zip"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/filesystems/findgo"
)

func TestFindGo(t *testing.T) {
	want := []string{
		"testdata/tree/0.go",
		"testdata/tree/1/1_1/1.go",
		"testdata/tree/2/2.go",
		"testdata/tree/3/4.go",
	}
	t.Run("lists correct files [in-memory]", func(t *testing.T) {
		fsys := fstest.MapFS{
			"testdata/tree/0.go":       {},
			"testdata/tree/1/1.txt":    {},
			"testdata/tree/1/1_1/1.go": {},
			"testdata/tree/2/2.go":     {},
			"testdata/tree/3/3.txt":    {},
			"testdata/tree/3/4.go":     {},
		}
		finder := findgo.NewFind(
			findgo.WithOptionFS(fsys),
		)
		got := finder.Files()

		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(got, want))
		}
	})
	t.Run("lists correct files [zip archive]", func(t *testing.T) {
		fsys, err := zip.OpenReader("files.zip")
		if err != nil {
			t.Fatal(err)
		}
		finder := findgo.NewFind(
			findgo.WithOptionFS(fsys),
		)
		got := finder.Files()

		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(got, want))
		}
	})
	t.Run("lists correct files [on-disk]", func(t *testing.T) {
		finder := findgo.NewFind(
			findgo.WithOptionArg("testdata/tree"),
		)
		got := finder.Files()
		for i := range want {
			path, _ := strings.CutPrefix(want[i], "testdata/tree/")
			want[i] = path
		}
		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(got, want))
		}
	})
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"testdata/tree/0.go":       {},
		"testdata/tree/1/1.txt":    {},
		"testdata/tree/1/1_1/1.go": {},
		"testdata/tree/2/2.go":     {},
		"testdata/tree/3/3.txt":    {},
		"testdata/tree/3/4.go":     {},
	}
	finder := findgo.NewFind(
		findgo.WithOptionFS(fsys),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		finder.Files()
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	finder := findgo.NewFind(
		findgo.WithOptionArg("testdata/tree"),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		finder.Files()
	}
}
