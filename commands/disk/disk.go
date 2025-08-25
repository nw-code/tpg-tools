package disk

import (
	//	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type filesystem struct {
	use  int
	name string
}

func NewFileSystem(name string, in io.Reader) (*filesystem, error) {
	reg := regexp.MustCompile(fmt.Sprintf("(?m)^%s.+\\s(\\d+)%%", name))

	b, _ := io.ReadAll(in)
	txt := string(b)

	matches := reg.FindStringSubmatch(txt)
	if len(matches) == 0 {
		return nil, fmt.Errorf("unable to find filesystem %s", name)
	}

	use, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}
	f := new(filesystem)
	f.use = use
	f.name = name

	return f, nil
}

func (f *filesystem) Usage() int {
	return f.use
}

func Main() {
	fs := flag.String("filesystem", "", "filesystem to query")
	flag.Parse()

	if *fs == "" {
		fmt.Fprint(os.Stderr, "filesystem flag is mandatory")
		flag.Usage()
		os.Exit(1)
	}

	var buf bytes.Buffer
	df := exec.Command("df", "-h")
	df.Stdout = &buf
	df.Run()
	f, _ := NewFileSystem(*fs, &buf)

	fmt.Println(f.Usage())
}
