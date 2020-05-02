package magic

import "io"

// Decoder can be used to decode an input from a io.Reader and provide it to
// a magic.Manager and it's underlying cursor for decoded record reading.
type Decoder interface {
	FromReader(src io.Reader) error
	Read() (record []string, err error)
}
