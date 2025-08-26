package disk

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func ParseDf(data, fs string) (int, error) {
	reg := regexp.MustCompile(fmt.Sprintf("(?m)^%s.+\\s(\\d+)%%", fs))
	matches := reg.FindStringSubmatch(data)
	if len(matches) == 0 {
		return 0, fmt.Errorf("unable to find filesystem %s", fs)
	}

	use, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	return use, nil
}

func GetDfOutput() (string, error) {
	data, err := exec.Command("df", "-h").CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Main() {
	fs := flag.String("filesystem", "", "Filesystem usage")
	flag.Parse()

	if *fs == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := GetDfOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	use, err := ParseDf(data, *fs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("Filesystem %q is %d%% full\n", *fs, use)
}
