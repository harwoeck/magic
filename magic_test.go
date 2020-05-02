package magic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	t.Run("Correct Manager Initialization", func(t *testing.T) {
		src := FromString("")
		dec := NewStringSplitDecoder("")

		m := NewManager(src, dec, nil)
		assert.NotNil(t, m)
		assert.NotNil(t, m.cursor)
		assert.Equal(t, dec, m.cursor.dec)
	})
}
