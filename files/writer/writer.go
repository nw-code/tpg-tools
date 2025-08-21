package writer

import (
	"flag"
	"os"
)

func WriteToFile(name string, b []byte) error {
	return os.WriteFile(name, b, 0600)
}

func Main() error {
	var b []byte
	size := flag.Int("size", 0, "provide number of zeros to initialise file")
	flag.Parse()
	path := flag.Args()[0]

	if *size > 0 {
		b = make([]byte, 1000)
	}
	return WriteToFile(path, b)
}
