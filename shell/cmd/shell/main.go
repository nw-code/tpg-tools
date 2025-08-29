package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nw-code/tpg-tools/shell"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Print(":> ")
	for scn.Scan() {
		fmt.Print(":> ")
		cmd, err := shell.CmdFromString(scn.Text())
		if err != nil {
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", out)
		fmt.Print(":> ")
	}
	fmt.Println("Exiting...")
}
