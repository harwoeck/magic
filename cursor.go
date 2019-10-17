package magic

import (
	"strings"
)

type cursor struct {
	src     SrcProvider
	lBuf    []string
	lBufIdx int
}

func newCursor(src SrcProvider) *cursor {
	c := &cursor{
		src:     src,
		lBufIdx: -1,
	}
	return c
}

func (c *cursor) nextBuf() {
	if !c.src.Scan() {
		panic("end of input reached")
	}

	buf := c.src.Text()
	c.lBuf = strings.Split(buf, " ")
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
