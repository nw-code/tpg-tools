package main

import (
	"fmt"
	"os"
	"path"
)

func countGoFiles(dirPath string) int {
	var count int
	files, _ := os.ReadDir(dirPath)

	for _, f := range files {
		if f.IsDir() {
			subDir := path.Join(dirPath, f.Name())
			count += countGoFiles(subDir)
		} else if path.Ext(f.Name()) == ".go" {
			count++
		}
	}

	return count
}

func main() {
	dirPath := "/tmp/tpg-tools"
	tally := countGoFiles(dirPath)
	fmt.Println(tally)
}
