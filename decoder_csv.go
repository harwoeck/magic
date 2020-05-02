package magic

import (
	"encoding/csv"
	"fmt"
	"io"
)

type csvDecoder struct {
	r         *csv.Reader
	comma     rune
	skipFirst bool
}

func NewCSVDecoder(comma rune, skipFirst bool) Decoder {
	return &csvDecoder{
		comma:     comma,
		skipFirst: skipFirst,
	}
}

func (d *csvDecoder) FromReader(r io.Reader) error {
	d.r = csv.NewReader(r)
	d.r.Comma = d.comma

	if d.skipFirst {
		if _, err := d.Read(); err != nil {
			return fmt.Errorf("csv: unable to skip first line: %w", err)
		}
	}

	return nil
}

func (d *csvDecoder) Read() (record []string, err error) {
	return d.r.Read()
}
