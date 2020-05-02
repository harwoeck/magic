package magic

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type stringSplitDecoder struct {
	r   *bufio.Scanner
	sep string
}

func NewStringSplitDecoder(sep string) Decoder {
	return &stringSplitDecoder{
		sep: sep,
	}
}

func (dec *stringSplitDecoder) FromReader(src io.Reader) error {
	dec.r = bufio.NewScanner(src)

	return nil
}

func (dec *stringSplitDecoder) Read() (record []string, err error) {
	if !dec.r.Scan() {
		return nil, fmt.Errorf("stringSplit: unable to scan next line. no one left")
	}

	splitted := strings.Split(dec.r.Text(), dec.sep)

	return splitted, nil
}
