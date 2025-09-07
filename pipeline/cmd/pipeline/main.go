package main

import (
	"github.com/nw-code/tpg-tools/pipeline"
)

func main() {
	pipeline.FromString("hello, world\n").Stdout()
}
