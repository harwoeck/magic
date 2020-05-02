package magic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newCursor(t *testing.T) {
	dec := NewStringSplitDecoder("")

	t.Run("Correct Cursor Initialization", func(t *testing.T) {
		c := newCursor(dec)

		assert.NotNil(t, c)
		assert.Equal(t, dec, c.dec)
		assert.Equal(t, -1, c.lBufIdx)
	})
}

func Test_nextBufPanic(t *testing.T) {
	dec := NewStringSplitDecoder("")

	t.Run("nextBuf Panic when empty", func(t *testing.T) {
		c := newCursor(dec)

		assert.Panics(t, func() {
			c.next()
		})
	})
}
