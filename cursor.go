package magic

import (
	"fmt"
	"os"
)

type cursor struct {
	dec     Decoder
	lBuf    []string
	lBufIdx int
}

func newCursor(src Decoder) *cursor {
	c := &cursor{
		dec:     src,
		lBufIdx: -1,
	}

	return c
}

func (c *cursor) nextBuf() {
	lBuf, err := c.dec.Read()
	if err != nil {
		fmt.Printf("[cursor] unable to read nextBuf (e.g. record): %v\n", err)
		os.Exit(-1)
	}

	c.lBuf = lBuf
	c.lBufIdx = -1
}

// next returns the next string from the cursor
func (c *cursor) next() string {
	c.lBufIdx++

	if c.lBufIdx == len(c.lBuf) {
		c.nextBuf()
		return c.next()
	}

	return c.lBuf[c.lBufIdx]
}
