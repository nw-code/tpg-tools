package main

import (
	"os"

	"github.com/nw-code/tpg-tools/paperwork/match"
)

func main() {
	searchString := os.Args[1]
	match.Main(searchString)
}
