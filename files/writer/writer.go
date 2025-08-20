package writer

import (
	"os"
)

func WriteToFile(name string, b []byte) error {
	return os.WriteFile(name, b, 0600)
}
