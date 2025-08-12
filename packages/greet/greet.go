package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintTo(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	fmt.Fprintf(w, "Hello, %s\n", scanner.Text())
}

func Main() {
	PrintTo(os.Stdin, os.Stdout)
}
