package findgo

import (
	"flag"
	"io/fs"
	"os"
	"path"
)

type option func(*find)

type find struct {
	fsys fs.FS
}

func NewFind(opts ...option) *find {
	f := &find{}
	for _, opt := range opts {
		opt(f)
	}

	return f
}

func WithOptionFS(fsys fs.FS) func(*find) {
	return func(f *find) {
		f.fsys = fsys
	}
}

func WithOptionArg(p string) func(*find) {
	return func(f *find) {
		f.fsys = os.DirFS(p)
	}
}

func (f *find) Files() []string {
	var paths []string
	fs.WalkDir(f.fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() && path.Ext(d.Name()) == ".go" {
			paths = append(paths, p)
		}
		return nil
	})

	return paths
}

func Main() []string {
	pathArg := flag.String("path", "", "root path of search")
	flag.Parse()

	if *pathArg == "" {
		flag.Usage()
		os.Exit(1)
	}

	f := NewFind(
		WithOptionArg(*pathArg),
	)
	return f.Files()
}
