package magic

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newCursor(t *testing.T) {
	src := bufio.NewScanner(strings.NewReader(""))
	
	t.Run("Correct Cursor Initialization", func(t *testing.T) {
		c := newCursor(src)

		assert.NotNil(t, c)
		assert.Equal(t, src, c.src)
		assert.Equal(t, -1, c.lBufIdx)
	})
}

func Test_nextBufPanic(t *testing.T) {
	src := bufio.NewScanner(strings.NewReader(""))

	t.Run("nextBuf Panic when empty", func(t *testing.T) {
		c := newCursor(src)

		assert.Panics(t, func() {
			c.next()
		})
	})
}
