package disk

import (
	"bufio"
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
	free int
	name string
}

func NewFileSystem(name string, in io.Reader) (*filesystem, error) {
	var match string
	f := new(filesystem)
	reg, err := regexp.Compile(fmt.Sprintf("^%s.*", name))
	if err != nil {
		return nil, fmt.Errorf("could not compile regex %v", err)
	}
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		match = reg.FindString(scanner.Text())
		if match != "" {
			break
		}
	}

	if match == "" {
		return nil, fmt.Errorf("could not file filesystem %q", name)
	}

	err = f.parse(match)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *filesystem) parse(attr string) error {
	reg := regexp.MustCompile(`^([^\s]+).+\s(\d+)%`)
	matches := reg.FindStringSubmatch(attr)
	if len(matches) == 0 {
		return fmt.Errorf("unable to parse attributes from %s", attr)
	}

	f.name = matches[1]
	free, err := strconv.Atoi(matches[2])
	if err != nil {
		return err
	}
	f.free = free

	return nil
}

func (f *filesystem) Usage() int {
	return f.free
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
