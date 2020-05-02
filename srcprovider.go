package magic

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func FromString(s string) io.Reader {
	return strings.NewReader(s)
}

func FromFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("[manager] unable to open file: %v\n", err)
	}

	return f
}
