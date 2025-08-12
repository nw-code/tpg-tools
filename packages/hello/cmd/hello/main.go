package main

import (
	"os"

	"github.com/nw-code/tpg-tools/packages/hello"
)

func main() {
	hello.PrintTo(os.Stdout)
}
