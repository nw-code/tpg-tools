package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"testing/fstest"
)

func main() {
	var count int
	/*
	   In Go, when initializing a map with struct literals as values,
	   the type of the struct literal can be omitted if the type is explicitly
	   defined for the map's value type. This is a specific case of
	   composite literal type omission.
	*/
	//type MapFS map[string]*MapFile

	fsys := fstest.MapFS{
		"tmp/tpg-tools/0.go":     &fstest.MapFile{},
		"tmp/tpg-tools/1/1.go":   {},
		"tmp/tpg-tools/1_1/1.go": {},
		"tmp/tpg-tools/2/2.go":   {},
		"tmp/tpg-tools/3/3.go":   {},
		"tmp/tpg-tools/4/4.go":   {},
	}

	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() && path.Ext(p) == ".go" {
			count++
		}
		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(count)
}
