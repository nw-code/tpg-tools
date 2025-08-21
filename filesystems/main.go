package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func main() {
	var count int
	fsys := os.DirFS("/tmp/tpg-tools")
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
