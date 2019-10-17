package magic

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	src := bufio.NewScanner(strings.NewReader(""))

	t.Run("Correct Manager Initialization", func(t *testing.T) {
		m := NewManager(src)
		assert.NotNil(t, m)
		assert.NotNil(t, m.cursor)
	})
}
