package disk

import (
	"bytes"
	"fmt"
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
	var buf = new(bytes.Buffer)

	df := exec.Command("df", "-h")
	df.Stdout = buf
	err := df.Run()

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
